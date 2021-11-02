package handler

import (
	"context"
	"time"

	"comm/db/mysql"
	"comm/logger"
	"comm/timer"
	"comm/util/deep"
	"helloworld-service/srv/types"
	"proto/helloworld"
)

// ArticleAdd defined TODO
func (h *Handler) ArticleAdd(ctx context.Context, req *helloworld.ArticleAddReq, rsp *helloworld.ArticleAddRes) error {
	logger.Info(ctx, "Received ArticleAdd request")

	marker := timer.NewMark()
	defer marker.Init(ctx, "ArticleAdd")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf(ctx, "InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	data := types.Article{}
	err = deep.Copy(&data, req)
	if err != nil {
		logger.Errorf(ctx, "Copy %v", err)
		return err
	}

	marker.Mark("ArticleAddDB")
	data.CreateTime = time.Now()
	data.UpdateTime = time.Now()
	err = h.ArticleAddDB(ctx, db, &data)
	if err != nil {
		logger.Errorf(ctx, "ArticleAddDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

// ArticleDel defined TODO
func (h *Handler) ArticleDel(ctx context.Context, req *helloworld.ArticleDelReq, rsp *helloworld.ArticleDelRes) error {
	logger.Info(ctx, "Received ArticleDel request")

	marker := timer.NewMark()
	defer marker.Init(ctx, "ArticleDel")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf(ctx, "InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.Article{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf(ctx, "Copy %v", err)
		return err
	}

	marker.Mark("ArticleDelDB")
	err = h.ArticleDelDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf(ctx, "ArticleDelDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

// ArticleUpdate defined TODO
func (h *Handler) ArticleUpdate(ctx context.Context, req *helloworld.ArticleUpdateReq, rsp *helloworld.ArticleUpdateRes) error {
	logger.Info(ctx, "Received ArticleUpdate request")

	marker := timer.NewMark()
	defer marker.Init(ctx, "ArticleUpdate")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf(ctx, "InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.Article{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf(ctx, "Copy %v", err)
		return err
	}

	marker.Mark("ArticleUpdateDB")
	err = h.ArticleUpdateDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf(ctx, "ArticleUpdateDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

// ArticleOne defined TODO
func (h *Handler) ArticleOne(ctx context.Context, req *helloworld.ArticleOneReq, rsp *helloworld.ArticleOneRes) error {
	logger.Info(ctx, "Received ArticleOne request")

	marker := timer.NewMark()
	defer marker.Init(ctx, "ArticleOne")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf(ctx, "InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, data := types.Article{}, types.Article{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf(ctx, "Copy %v", err)
		return err
	}

	marker.Mark("ArticleOneDB")
	err = h.ArticleOneDB(ctx, db, &filter, &data)
	if err != nil {
		logger.Errorf(ctx, "ArticleOneDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

// ArticlePage defined TODO
func (h *Handler) ArticlePage(ctx context.Context, req *helloworld.ArticlePageReq, rsp *helloworld.ArticlePageRes) error {
	logger.Info(ctx, "Received ArticlePage request")

	marker := timer.NewMark()
	defer marker.Init(ctx, "ArticlePage")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf(ctx, "InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, list := types.Article{}, []types.Article{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf(ctx, "Copy %v", err)
		return err
	}

	marker.Mark("ArticlePageDB")
	var totalRecord int64
	var totalPage int64
	err = h.ArticlePageDB(ctx, db, &filter, &list, &totalRecord)
	if err != nil {
		logger.Errorf(ctx, "InitDB %v", err)
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
		logger.Errorf(ctx, "Copy %v", err)
		return err
	}
	return nil
}
