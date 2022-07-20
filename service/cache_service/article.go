package cache_service

import (
	"encoding/json"

	"github.com/EDDYCJY/go-gin-example/service/cache_service"
	"github.com/planet-i/gin-project/models"
	"github.com/planet-i/gin-project/pkg/gredis"
	"github.com/planet-i/gin-project/pkg/logging"
)

type Article struct {
	ID    int
	TagID int
	State int

	PageNum  int
	PageSize int
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, article, 3600)
	return article, nil
}
