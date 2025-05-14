package mainusecase

import (
	"context"
	"encoding/json"
	"time"

	"recsvc/internal/domain/er"
	"recsvc/internal/model"

	"github.com/redis/go-redis/v9"
)

const CacheKey = "recommendation_cache"

type GetRecommendationOut struct {
	Recommendations []Recommendation
}

type Recommendation struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (u *Usecase) GetRecommendation(c context.Context) (*GetRecommendationOut, error) {
	val, err := u.redisSvc.Cli.Get(c, CacheKey).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, er.W(err)
		}

		recommendations, err := u.getRecommendationFromDB(c)
		if err != nil {
			return nil, er.W(err)
		}

		recommendationsBytes, err := json.Marshal(recommendations)
		if err != nil {
			return nil, er.W(err)
		}

		if err := u.redisSvc.Cli.Set(c, CacheKey, recommendationsBytes, 10*time.Minute).Err(); err != nil {
			return nil, er.W(err)
		}

		return &GetRecommendationOut{
			Recommendations: recommendations,
		}, nil
	}

	var recommendations []Recommendation
	if err := json.Unmarshal([]byte(val), &recommendations); err != nil {
		return nil, er.W(err)
	}

	return &GetRecommendationOut{Recommendations: recommendations}, nil
}

func (u *Usecase) getRecommendationFromDB(c context.Context) ([]Recommendation, error) {
	// mock slow db query
	MinResponseTime := time.Now().Add(3 * time.Second)
	defer time.Sleep(time.Until(MinResponseTime))

	var Recommendations []Recommendation

	if err := u.db.WithContext(c).
		Model(model.Merchandise{}).
		Order("visit_count DESC").
		Limit(10).
		Scan(&Recommendations).Error; err != nil {
		return nil, err
	}

	return Recommendations, nil
}
