package handler

import (
	"context"
	"time"

	"comm/db/mysql"
	"comm/logger"
	"comm/timer"
	"comm/util/deep"
	"{{.name}}-service/srv/types"
	"proto/{{.name}}"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = {{.name}}.{{toTitle .entity}}{}
var _ = logger.Info

// {{toTitle .entity}}Add defined TODO
func (h *Handler) {{toTitle .entity}}Add(ctx context.Context, req *{{.name}}.{{toTitle .entity}}AddReq, rsp *{{.name}}.{{toTitle .entity}}AddRes) error {
	logger.Info("Received {{toTitle .entity}}Add request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .entity}}Add")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	data := types.{{toTitle .entity}}{}
	err = deep.Copy(&data, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .entity}}AddDB")
	data.CreateTime = time.Now()
	data.UpdateTime = time.Now()
	err = h.{{toTitle .entity}}AddDB(ctx, db, &data)
	if err != nil {
		logger.Errorf("{{toTitle .entity}}AddDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

// {{toTitle .entity}}Del defined TODO
func (h *Handler) {{toTitle .entity}}Del(ctx context.Context, req *{{.name}}.{{toTitle .entity}}DelReq, rsp *{{.name}}.{{toTitle .entity}}DelRes) error {
	logger.Info("Received {{toTitle .entity}}Del request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .entity}}Del")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.{{toTitle .entity}}{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .entity}}DelDB")
	err = h.{{toTitle .entity}}DelDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf("{{toTitle .entity}}DelDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

// {{toTitle .entity}}Update defined TODO
func (h *Handler) {{toTitle .entity}}Update(ctx context.Context, req *{{.name}}.{{toTitle .entity}}UpdateReq, rsp *{{.name}}.{{toTitle .entity}}UpdateRes) error {
	logger.Info("Received {{toTitle .entity}}Update request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .entity}}Update")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.{{toTitle .entity}}{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .entity}}UpdateDB")
	err = h.{{toTitle .entity}}UpdateDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf("{{toTitle .entity}}UpdateDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

// {{toTitle .entity}}One defined TODO
func (h *Handler) {{toTitle .entity}}One(ctx context.Context, req *{{.name}}.{{toTitle .entity}}OneReq, rsp *{{.name}}.{{toTitle .entity}}OneRes) error {
	logger.Info("Received {{toTitle .entity}}One request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .entity}}One")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, data := types.{{toTitle .entity}}{}, types.{{toTitle .entity}}{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .entity}}OneDB")
	err = h.{{toTitle .entity}}OneDB(ctx, db, &filter, &data)
	if err != nil {
		logger.Errorf("{{toTitle .entity}}OneDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

// {{toTitle .entity}}Page defined TODO
func (h *Handler) {{toTitle .entity}}Page(ctx context.Context, req *{{.name}}.{{toTitle .entity}}PageReq, rsp *{{.name}}.{{toTitle .entity}}PageRes) error {
	logger.Info("Received {{toTitle .entity}}Page request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .entity}}Page")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, list := types.{{toTitle .entity}}{}, []types.{{toTitle .entity}}{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .entity}}PageDB")
	var totalRecord int64
	var totalPage int64
	err = h.{{toTitle .entity}}PageDB(ctx, db, &filter, &list, &totalRecord)
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
