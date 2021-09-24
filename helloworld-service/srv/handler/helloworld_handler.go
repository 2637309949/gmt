package handler

import (
	"context"
	"time"

	"comm/db/mysql"
	"comm/logger"
	"comm/timer"
	"comm/util/deep"
	"helloworld/srv/types"
	"proto/helloworld"
)

func (h *Handler) ArticleAdd(ctx context.Context, req *helloworld.Article, rsp *helloworld.Article) error {
	logger.Info("Received ArticleAdd request")

	marker := timer.NewMark()
	defer marker.Init("ArticleAdd")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	data := types.Article{}
	err = deep.Copy(&data, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticleAddDB")
	data.CreateTime = time.Now()
	data.UpdateTime = time.Now()
	err = h.ArticleAddDB(ctx, db, &data)
	if err != nil {
		logger.Errorf("ArticleAddDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

func (h *Handler) ArticleDel(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.Article) error {
	logger.Info("Received ArticleDel request")

	marker := timer.NewMark()
	defer marker.Init("ArticleDel")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.Article{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticleDelDB")
	err = h.ArticleDelDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf("ArticleDelDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

func (h *Handler) ArticleUpdate(ctx context.Context, req *helloworld.Article, rsp *helloworld.Article) error {
	logger.Info("Received ArticleUpdate request")

	marker := timer.NewMark()
	defer marker.Init("ArticleUpdate")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.Article{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticleUpdateDB")
	err = h.ArticleUpdateDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf("ArticleUpdateDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

func (h *Handler) ArticleOne(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.Article) error {
	logger.Info("Received ArticleOne request")

	marker := timer.NewMark()
	defer marker.Init("ArticleOne")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, data := types.Article{}, types.Article{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticleOneDB")
	err = h.ArticleOneDB(ctx, db, &filter, &data)
	if err != nil {
		logger.Errorf("ArticleOneDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

func (h *Handler) ArticlePage(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.ArticleList) error {
	logger.Info("Received ArticlePage request")

	marker := timer.NewMark()
	defer marker.Init("ArticlePage")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, list := types.Article{}, []types.Article{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticlePageDB")
	var totalRecord int64
	var totalPage int64
	err = h.ArticlePageDB(ctx, db, &filter, &list, &totalRecord)
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy rsp")
	err = deep.Copy(&rsp.Data, &list)
	if totalRecord < req.Size {
		totalPage = 1
	} else if totalRecord%req.Size == 0 {
		totalPage = totalRecord / req.Size
	} else {
		totalPage = totalRecord/req.Size + 1
	}

	rsp.TotalRecord = totalRecord
	rsp.TotalPage = totalPage
	rsp.Size = req.Size
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}
	return nil
}
