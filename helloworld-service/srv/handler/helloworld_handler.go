package handler

import (
	"context"

	"comm/db/mysql"
	"comm/logger"
	"comm/timer"
	"comm/util/reflect"
	"helloworld/srv/types"
	"proto/helloworld"
)

func (h *Handler) ArticleAdd(ctx context.Context, req *helloworld.Article, rsp *helloworld.Article) error {
	logger.Info("Received ArticleAdd request")

	marker := timer.NewMark()
	defer marker.Init("ArticleAdd")

	marker.Mark("data_source")
	db, err := mysql.InitDb("data_source")
	if err != nil {
		logger.Errorf("InitDb %v", err)
		return err
	}

	marker.Mark("Copy req")
	data := types.Article{}
	err = reflect.Copy(&data, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticleAddDb")
	err = h.ArticleAddDb(ctx, db, &data)
	if err != nil {
		logger.Errorf("ArticleAddDb %v", err)
		return err
	}
	rsp.Id = data.ID.Int64
	return nil
}

func (h *Handler) ArticleDel(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.Article) error {
	logger.Info("Received ArticleDel request")

	marker := timer.NewMark()
	defer marker.Init("ArticleDel")

	marker.Mark("data_source")
	db, err := mysql.InitDb("data_source")
	if err != nil {
		logger.Errorf("InitDb %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.Article{}
	err = reflect.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticleDelDb")
	err = h.ArticleDelDb(ctx, db, &filter)
	if err != nil {
		logger.Errorf("ArticleDelDb %v", err)
		return err
	}
	rsp.Id = filter.ID.Int64
	return nil
}

func (h *Handler) ArticleUpdate(ctx context.Context, req *helloworld.Article, rsp *helloworld.Article) error {
	logger.Info("Received ArticleUpdate request")

	marker := timer.NewMark()
	defer marker.Init("ArticleUpdate")

	marker.Mark("data_source")
	db, err := mysql.InitDb("data_source")
	if err != nil {
		logger.Errorf("InitDb %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.Article{}
	err = reflect.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticleUpdateDb")
	err = h.ArticleUpdateDb(ctx, db, &filter)
	if err != nil {
		logger.Errorf("ArticleUpdateDb %v", err)
		return err
	}
	rsp.Id = filter.ID.Int64
	return nil
}

func (h *Handler) ArticleOne(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.Article) error {
	logger.Info("Received ArticleOne request")

	marker := timer.NewMark()
	defer marker.Init("ArticleOne")

	marker.Mark("data_source")
	db, err := mysql.InitDb("data_source")
	if err != nil {
		logger.Errorf("InitDb %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, data := types.Article{}, types.Article{}
	err = reflect.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticleOneDb")
	err = h.ArticleOneDb(ctx, db, &filter, &data)
	if err != nil {
		logger.Errorf("ArticleOneDb %v", err)
		return err
	}
	rsp.Id = data.ID.Int64
	return nil
}

func (h *Handler) ArticlePage(ctx context.Context, req *helloworld.ArticleFilter, rsp *helloworld.ArticleList) error {
	logger.Info("Received ArticlePage request")

	marker := timer.NewMark()
	defer marker.Init("ArticlePage")

	marker.Mark("data_source")
	db, err := mysql.InitDb("data_source")
	if err != nil {
		logger.Errorf("InitDb %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, list := types.Article{}, []types.Article{}
	err = reflect.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("ArticlePageDb")
	var totalRecord int32
	var totalPage int32
	err = h.ArticlePageDb(ctx, db, &filter, &list, &totalRecord)
	if err != nil {
		logger.Errorf("InitDb %v", err)
		return err
	}

	marker.Mark("Copy rsp")
	err = reflect.Copy(&rsp.Data, &list)
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
