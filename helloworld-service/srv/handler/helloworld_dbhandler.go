package handler

import (
	"comm/logger"
	"context"
	"helloworld/srv/types"
	"time"

	"github.com/guregu/null"
	"github.com/xormplus/xorm"
)

func (h *Handler) ArticleAddDB(ctx context.Context, db *xorm.Engine, data *types.Article) error {
	logger.Info("Received ArticleAddDB request")
	ret, err := db.Insert(data)
	if err != nil {
		return err
	}
	data.ID.Int64 = ret
	return err
}

func (h *Handler) ArticleDelDB(ctx context.Context, db *xorm.Engine, filter *types.Article) error {
	logger.Info("Received ArticleDelDB request")
	ret, err := db.ID(filter.ID.Int64).Update(&types.Article{
		UpdateTime: null.TimeFrom(time.Now()),
		IsDelete:   null.IntFrom(1),
	})
	if err != nil {
		return err
	}
	filter.ID.Int64 = ret
	return err
}

func (h *Handler) ArticleUpdateDB(ctx context.Context, db *xorm.Engine, data *types.Article) error {
	logger.Info("Received ArticleUpdateDB request")
	ret, err := db.ID(data.ID.Int64).Update(data)
	if err != nil {
		return err
	}
	data.ID.Int64 = ret
	return err
}

func (h *Handler) ArticleOneDB(ctx context.Context, db *xorm.Engine, filter *types.Article, data *types.Article) error {
	logger.Info("Received ArticleOneDB request")
	_, err := db.Get(data)
	if err != nil {
		return err
	}
	return err
}

func (h *Handler) ArticlePageDB(ctx context.Context, db *xorm.Engine, filter *types.Article, list *[]types.Article, totalRecord ...*int64) error {
	logger.Info("Received ArticlePageDB request")
	err := db.Find(list, filter)
	if err != nil {
		return err
	}
	ct, err := db.Count(filter)
	if err != nil {
		return err
	}
	if len(totalRecord) > 0 {
		*totalRecord[0] = ct
	}
	return nil
}
