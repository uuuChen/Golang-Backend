package products

import (
	"encoding/json"
	"fmt"
	"glossika/internal/db"
	"time"

	"github.com/go-redis/redis"
)

const redisKeyRecommendations = "recommendations"

func (helper *productHelper) ListRecommendations() ([]db.Product, error) {
	val, err := helper.redisClient.Get(redisKeyRecommendations).Result()
	if err == redis.Nil {
		ret, err := helper.listAndCacheRecommendationsFromDB()
		if err != nil {
			return nil, fmt.Errorf("failed to list and cache recommendations")
		}
		return ret, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get recommendations")
	} else {
		var ret []db.Product
		err = json.Unmarshal([]byte(val), &ret)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal recommendations")
		}
		return ret, nil
	}
}

func (helper *productHelper) listAndCacheRecommendationsFromDB() ([]db.Product, error) {
	ret, err := helper.db.ListRecommendations()
	if err != nil {
		return nil, fmt.Errorf("failed to list recommendations")
	}

	retJSON, err := json.Marshal(ret)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal recommendations")
	}

	const cacheExpirationTime = 10 * time.Minute

	err = helper.redisClient.Set(redisKeyRecommendations, retJSON, cacheExpirationTime).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to store recommendations in cache")
	}

	return ret, nil
}
