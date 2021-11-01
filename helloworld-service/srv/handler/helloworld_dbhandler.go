package handler

import (
	"comm/logger"
	"context"
	"helloworld-service/srv/types"
	"time"

	"github.com/xormplus/xorm"
)

// ArticleAddDB defined TODO
func (h *Handler) ArticleAddDB(ctx context.Context, db *xorm.Engine, item *types.Article) error {
	logger.Info("Received ArticleAddDB request")

	item.CreateTime = time.Now()
	ret, err := db.Insert(item)
	if err != nil {
		return err
	}
	item.ID = uint64(ret)
	return err
}

// ArticleDelDB defined TODO
func (h *Handler) ArticleDelDB(ctx context.Context, db *xorm.Engine, where *types.Article) error {
	logger.Info("Received ArticleDelDB request")

	ret, err := db.ID(where.ID).Update(&types.Article{
		UpdateTime: time.Now(),
		IsDelete:   1,
	})
	if err != nil {
		return err
	}
	where.ID = uint64(ret)
	return err
}

// ArticleUpdateDB defined TODO
func (h *Handler) ArticleUpdateDB(ctx context.Context, db *xorm.Engine, item *types.Article) error {
	logger.Info("Received ArticleUpdateDB request")

	_, err := db.ID(item.ID).Update(item)
	if err != nil {
		return err
	}
	return err
}

// ArticleOneDB defined TODO
func (h *Handler) ArticleOneDB(ctx context.Context, db *xorm.Engine, where *types.Article, item *types.Article) error {
	logger.Info("Received ArticleOneDB request")

	ext, err := db.Get(item)
	if err != nil {
		return err
	}
	if !ext {
		return xorm.ErrNotExist
	}
	return err
}

// ArticlePageDB defined TODO
func (h *Handler) ArticlePageDB(ctx context.Context, db *xorm.Engine, where *types.Article, list *[]types.Article, totalRecord ...*int64) error {
	logger.Info("Received ArticlePageDB request")

	err := db.Find(list, where)
	if err != nil {
		return err
	}
	if len(totalRecord) > 0 {
		ct, err := db.Count(where)
		if err != nil {
			return err
		}
		*totalRecord[0] = ct
	}
	return nil
}
