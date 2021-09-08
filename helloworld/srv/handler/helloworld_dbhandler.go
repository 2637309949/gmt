package handler

import (
	"comm/logger"
	"context"
	"github.com/xormplus/xorm"
	"helloworld/srv/types"
)

func (h *Handler) ArticleDbAdd(ctx context.Context, db *xorm.Engine, data *types.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}

func (h *Handler) ArticleDbDel(ctx context.Context, db *xorm.Engine, filter *types.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}

func (h *Handler) ArticleDbUpdate(ctx context.Context, db *xorm.Engine, data *types.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}

func (h *Handler) ArticleDbOne(ctx context.Context, db *xorm.Engine, filter *types.Article, data *types.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}

func (h *Handler) ArticleDbPage(ctx context.Context, db *xorm.Engine, filter *types.Article, list *[]types.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}
