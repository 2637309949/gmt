package handler

import (
	"comm/logger"
	"context"
	"helloworld/srv/types"
	"time"

	"github.com/guregu/null"
	"github.com/xormplus/xorm"
)

func (h *Handler) ArticleAddDb(ctx context.Context, db *xorm.Engine, data *types.Article) error {
	logger.Info("Received ArticleAddDb request")
	ret, err := db.Insert(data)
	data.ID.Int64 = ret
	return err
}

func (h *Handler) ArticleDelDb(ctx context.Context, db *xorm.Engine, filter *types.Article) error {
	logger.Info("Received ArticleDelDb request")
	ret, err := db.ID(filter.ID.Int64).Update(&types.Article{
		UpdateTime: null.TimeFrom(time.Now()),
		IsDelete:   null.IntFrom(1),
	})
	filter.ID.Int64 = ret
	return err
}

func (h *Handler) ArticleUpdateDb(ctx context.Context, db *xorm.Engine, data *types.Article) error {
	logger.Info("Received ArticleUpdateDb request")
	ret, err := db.ID(data.ID.Int64).Update(data)
	data.ID.Int64 = ret
	return err
}

func (h *Handler) ArticleOneDb(ctx context.Context, db *xorm.Engine, filter *types.Article, data *types.Article) error {
	logger.Info("Received ArticleOneDb request")
	_, err := db.Get(data)
	return err
}

func (h *Handler) ArticlePageDb(ctx context.Context, db *xorm.Engine, filter *types.Article, list *[]types.Article, totalRecord ...*int64) error {
	logger.Info("Received ArticlePageDb request")
	err := db.Find(list, filter)
	if err != nil {
		return err
	}
	ct, err := db.Count(filter)
	if len(totalRecord) > 0 {
		*totalRecord[0] = ct
	}
	return nil
}
