package handler

import (
	"context"

	"comm/db/mysql"
	"comm/logger"
	"comm/timer"
	"comm/util/deep"
	"{{.name}}-service/srv/types"
	"proto/{{.name}}"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = {{.name}}.{{toTitle .proto.Name}}{}
var _ = logger.Info

// {{toTitle .proto.Name}}Add defined TODO
func (h *Handler) {{toTitle .proto.Name}}Add(ctx context.Context, req *{{.name}}.{{toTitle .proto.Name}}AddReq, rsp *{{.name}}.{{toTitle .proto.Name}}AddRes) error {
	logger.Info("Received {{toTitle .proto.Name}}Add request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .proto.Name}}Add")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	data := types.{{toTitle .proto.Name}}{}
	err = deep.Copy(&data, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .proto.Name}}AddDB")
	// data.CreateTime = time.Now()
	// data.UpdateTime = time.Now()
	err = h.{{toTitle .proto.Name}}AddDB(ctx, db, &data)
	if err != nil {
		logger.Errorf("{{toTitle .proto.Name}}AddDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

// {{toTitle .proto.Name}}Del defined TODO
func (h *Handler) {{toTitle .proto.Name}}Del(ctx context.Context, req *{{.name}}.{{toTitle .proto.Name}}DelReq, rsp *{{.name}}.{{toTitle .proto.Name}}DelRes) error {
	logger.Info("Received {{toTitle .proto.Name}}Del request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .proto.Name}}Del")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.{{toTitle .proto.Name}}{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .proto.Name}}DelDB")
	err = h.{{toTitle .proto.Name}}DelDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf("{{toTitle .proto.Name}}DelDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

// {{toTitle .proto.Name}}Update defined TODO
func (h *Handler) {{toTitle .proto.Name}}Update(ctx context.Context, req *{{.name}}.{{toTitle .proto.Name}}UpdateReq, rsp *{{.name}}.{{toTitle .proto.Name}}UpdateRes) error {
	logger.Info("Received {{toTitle .proto.Name}}Update request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .proto.Name}}Update")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter := types.{{toTitle .proto.Name}}{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .proto.Name}}UpdateDB")
	err = h.{{toTitle .proto.Name}}UpdateDB(ctx, db, &filter)
	if err != nil {
		logger.Errorf("{{toTitle .proto.Name}}UpdateDB %v", err)
		return err
	}
	rsp.Id = filter.ID
	return nil
}

// {{toTitle .proto.Name}}One defined TODO
func (h *Handler) {{toTitle .proto.Name}}One(ctx context.Context, req *{{.name}}.{{toTitle .proto.Name}}OneReq, rsp *{{.name}}.{{toTitle .proto.Name}}OneRes) error {
	logger.Info("Received {{toTitle .proto.Name}}One request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .proto.Name}}One")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, data := types.{{toTitle .proto.Name}}{}, types.{{toTitle .proto.Name}}{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .proto.Name}}OneDB")
	err = h.{{toTitle .proto.Name}}OneDB(ctx, db, &filter, &data)
	if err != nil {
		logger.Errorf("{{toTitle .proto.Name}}OneDB %v", err)
		return err
	}
	rsp.Id = data.ID
	return nil
}

// {{toTitle .proto.Name}}Page defined TODO
func (h *Handler) {{toTitle .proto.Name}}Page(ctx context.Context, req *{{.name}}.{{toTitle .proto.Name}}PageReq, rsp *{{.name}}.{{toTitle .proto.Name}}PageRes) error {
	logger.Info("Received {{toTitle .proto.Name}}Page request")

	marker := timer.NewMark()
	defer marker.Init("{{toTitle .proto.Name}}Page")

	marker.Mark("data_source")
	db, err := mysql.InitDB("data_source")
	if err != nil {
		logger.Errorf("InitDB %v", err)
		return err
	}

	marker.Mark("Copy req")
	filter, list := types.{{toTitle .proto.Name}}{}, []types.{{toTitle .proto.Name}}{}
	err = deep.Copy(&filter, req)
	if err != nil {
		logger.Errorf("Copy %v", err)
		return err
	}

	marker.Mark("{{toTitle .proto.Name}}PageDB")
	var totalRecord uint64
	var totalPage uint64
	err = h.{{toTitle .proto.Name}}PageDB(ctx, db, &filter, &list, &totalRecord)
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
