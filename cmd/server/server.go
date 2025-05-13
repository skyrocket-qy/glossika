package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"glossika/api"
	"glossika/internal/boot"
	"glossika/internal/controller"
	authcontroller "glossika/internal/controller/auth"
	maincontroller "glossika/internal/controller/main"
	"glossika/internal/domain"
	"glossika/internal/service/database"
	"glossika/internal/service/redis"
	authucase "glossika/internal/usecase/auth"
	mainusecase "glossika/internal/usecase/main"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "The main service command",
	Long:  ``,
	Run:   RunServer,
}

func RunServer(cmd *cobra.Command, args []string) {
	if err := boot.Run(); err != nil {
		log.Fatal().Msgf("Initialization failed: %v", err)
	}

	db, err := database.New()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
	defer database.Close(db)

	redisSvc, err := redis.New()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	handlers, err := newHandlers(redisSvc, db)
	if err != nil {
		log.Error().Err(err).Msg("Error creating handlers")
		return
	}

	router := newRouter()
	api.RegisterAPIHandlers(router, handlers)
	port, _ := cmd.Flags().GetString("port")
	server := newHTTPServer(router, port)

	go func() {
		log.Info().Msgf("Starting server on port %s...", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Server failed")
			return
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server forced to shutdown: %v", err)
	} else {
		log.Info().Msg("Server shut down gracefully.")
	}
}

func init() {
	Cmd.Flags().StringP("port", "p", "8080", "port")
	Cmd.Flags().
		StringVarP(&domain.Database, `database`, "d", "postgres", `"postgres", "mysql"`)
	Cmd.Flags().
		StringVarP(&domain.Env, `env`, "e", "local", `"local", "dev", "prod"`)
	Cmd.Flags().
		StringVarP(&domain.LogTo, `log`, "l", "stdout", `"stdout", "loki"`)

	Cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		validDatabases := map[string]bool{"postgres": true, "mysql": true}
		if !validDatabases[domain.Database] {
			return fmt.Errorf("invalid database value: %s. Must be one of: postgres, mysql",
				domain.Database)
		}

		validEnvs := map[string]bool{"local": true, "dev": true, "prod": true}
		if !validEnvs[domain.Env] {
			return fmt.Errorf(
				"invalid environment value: %s. Must be one of: dev, prod",
				domain.Env,
			)
		}

		port, _ := cmd.Flags().GetString("port")
		if _, err := strconv.Atoi(port); err != nil || port == "" {
			return fmt.Errorf("invalid port value: %s. Must be a valid number", port)
		}

		validLogs := map[string]bool{"stdout": true, "loki": true}
		if !validLogs[domain.LogTo] {
			return fmt.Errorf(
				"invalid log value: %s. Must be one of: stdout, loki",
				domain.LogTo,
			)
		}

		return nil
	}
}

func newHandlers(redisSvc *redis.RedisService, db *gorm.DB) (*controller.Handlers, error) {
	authUsecase := authucase.New(redisSvc, db)
	mainUsecase := mainusecase.New(redisSvc, db)

	authHandler := authcontroller.NewHandler(authUsecase)
	MainHandler := maincontroller.NewHandler(mainUsecase)

	handlers := controller.NewHandlers(authHandler, MainHandler)

	return handlers, nil
}

func newHTTPServer(router *gin.Engine, port string) *http.Server {
	return &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func newRouter() *gin.Engine {
	if domain.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	return gin.Default()
}
