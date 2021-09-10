package handler

import (
	"comm/logger"
	"context"
	"github.com/guregu/null"
	"github.com/xormplus/xorm"
	"helloworld/srv/types"
)

func (h *Handler) ArticleAddDb(ctx context.Context, db *xorm.Engine, data *types.Article) error {
	logger.Info("Received ArticleAddDb request")
	return nil
}

func (h *Handler) ArticleDelDb(ctx context.Context, db *xorm.Engine, filter *types.Article) error {
	logger.Info("Received ArticleDelDb request")
	return nil
}

func (h *Handler) ArticleUpdateDb(ctx context.Context, db *xorm.Engine, data *types.Article) error {
	logger.Info("Received ArticleUpdateDb request")
	return nil
}

func (h *Handler) ArticleOneDb(ctx context.Context, db *xorm.Engine, filter *types.Article, data *types.Article) error {
	logger.Info("Received ArticleOneDb request")
	return nil
}

func (h *Handler) ArticlePageDb(ctx context.Context, db *xorm.Engine, filter *types.Article, list *[]types.Article, totalRecord ...*int32) error {
	logger.Info("Received ArticlePageDb request")
	*list = []types.Article{types.Article{Name: null.StringFrom("hello")}}
	if len(totalRecord) > 0 {
		*totalRecord[0] = 10 
	}
	return nil
}
