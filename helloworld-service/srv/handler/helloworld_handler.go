package handler

import (
	"context"

	"comm/logger"
	"proto/helloworld"
)

func (h *Handler) ArticleAdd(ctx context.Context, req *helloworld.Article, rsp *helloworld.Article) error {
	logger.Info("Received srv.helloworld.ArticleAdd request")
	return nil
}

func (h *Handler) ArticleDel(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.Article) error {
	logger.Info("Received srv.helloworld.ArticleDel request")
	return nil
}

func (h *Handler) ArticleUpdate(ctx context.Context, req *helloworld.Article, rsp *helloworld.Article) error {
	logger.Info("Received srv.helloworld.ArticleUpdate request")
	return nil
}

func (h *Handler) ArticleOne(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.Article) error {
	logger.Info("Received srv.helloworld.ArticleOne request")
	return nil
}

func (h *Handler) ArticlePage(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.ArticleList) error {
	logger.Info("Received srv.helloworld.ArticlePage request")
	return nil
}
