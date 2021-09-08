package handler

import (
	"context"

	"comm/logger"
	"proto/helloworld"
)

func (h *Handler) ArticleAdd(ctx context.Context, req *helloworld.Article, rsp *helloworld.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}

func (h *Handler) ArticleDel(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}

func (h *Handler) ArticleUpdate(ctx context.Context, req *helloworld.Article, rsp *helloworld.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}

func (h *Handler) ArticleOne(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.Article) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}

func (h *Handler) ArticlePage(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.ArticleList) error {
	logger.Info("Received Helloworld.Call request")
	return nil
}
