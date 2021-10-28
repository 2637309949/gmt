package handler

import (
	"context"
	"time"

	"comm/db/mysql"
	"comm/logger"
	"comm/timer"
	"comm/util/deep"
	"test-service/srv/types"
	"proto/test"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = test.Entity{}
var _ = logger.Info

// EntityAdd defined TODO
func (h *Handler) EntityAdd(ctx context.Context, req *test.EntityAddReq, rsp *test.EntityAddRes) error {
	logger.Info("Received EntityAdd request")

	marker := timer.NewMark()
	defer marker.Init("EntityAdd")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	data := types.Entity{}
	err = deep.Copy(&data, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("EntityAddDB")
	data.CreateTime = time.Now()
	data.UpdateTime = time.Now()
	err = h.EntityAddDB(ctx, db, &data)
	if err != nil {
		logger.Errorf("EntityAddDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

// EntityDel defined TODO
func (h *Handler) EntityDel(ctx context.Context, req *test.EntityDelReq, rsp *test.EntityDelRes) error {
	logger.Info("Received EntityDel request")

	marker := timer.NewMark()
	defer marker.Init("EntityDel")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.Entity{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("EntityDelDB")
	err = h.EntityDelDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf("EntityDelDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

// EntityUpdate defined TODO
func (h *Handler) EntityUpdate(ctx context.Context, req *test.EntityUpdateReq, rsp *test.EntityUpdateRes) error {
	logger.Info("Received EntityUpdate request")

	marker := timer.NewMark()
	defer marker.Init("EntityUpdate")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.Entity{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("EntityUpdateDB")
	err = h.EntityUpdateDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf("EntityUpdateDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

// EntityOne defined TODO
func (h *Handler) EntityOne(ctx context.Context, req *test.EntityOneReq, rsp *test.EntityOneRes) error {
	logger.Info("Received EntityOne request")

	marker := timer.NewMark()
	defer marker.Init("EntityOne")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, data := types.Entity{}, types.Entity{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("EntityOneDB")
	err = h.EntityOneDB(ctx, db, &filter, &data)
	if err != nil {
		logger.Errorf("EntityOneDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

// EntityPage defined TODO
func (h *Handler) EntityPage(ctx context.Context, req *test.EntityPageReq, rsp *test.EntityPageRes) error {
	logger.Info("Received EntityPage request")

	marker := timer.NewMark()
	defer marker.Init("EntityPage")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, list := types.Entity{}, []types.Entity{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("EntityPageDB")
	var totalRecord int64
	var totalPage int64
	err = h.EntityPageDB(ctx, db, &filter, &list, &totalRecord)
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
