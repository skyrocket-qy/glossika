package database

import (
	"fmt"

	"recsvc/internal/domain"
	"recsvc/internal/domain/er"
	"recsvc/internal/model"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (db *gorm.DB, err error) {
	log.Info().Msg("New db")
	config := gorm.Config{}

	switch domain.Database {
	case "mysql":
		log.Info().Msg("Connecting to MySQL")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			viper.GetString("db.user"),
			viper.GetString("db.password"),
			viper.GetString("db.host"),
			viper.GetInt("db.port"),
			viper.GetString("db.dbname"),
			viper.GetString("db.timezone"),
		)

		db, err = gorm.Open(mysql.Open(dsn), &config)
	}

	if err != nil {
		err = er.W(err, er.DBUnavailable)
		return
	}

	db.AutoMigrate(&model.User{}, &model.Merchandise{})

	if initFakeData(db) != nil {
		return nil, er.W(err)
	}

	return db, err
}

func Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func initFakeData(db *gorm.DB) error {
	// check if fake data already exists
	if tx := db.Take(&model.Merchandise{}); tx.RowsAffected > 0 {
		return nil
	}

	fakeMerchandises := []model.Merchandise{
		{Name: "Merchandise 1", VisitCount: 1},
		{Name: "Merchandise 2", VisitCount: 1},
		{Name: "Merchandise 3", VisitCount: 1},
		{Name: "Merchandise 4", VisitCount: 1},
		{Name: "Merchandise 5", VisitCount: 1},
		{Name: "Merchandise 6", VisitCount: 1},
		{Name: "Merchandise 7", VisitCount: 1},
		{Name: "Merchandise 8", VisitCount: 1},
		{Name: "Merchandise 9", VisitCount: 1},
		{Name: "Merchandise 10", VisitCount: 0},
		{Name: "Merchandise 11", VisitCount: 1},
	}
	if err := db.Create(&fakeMerchandises).Error; err != nil {
		return err
	}
	return nil
}
