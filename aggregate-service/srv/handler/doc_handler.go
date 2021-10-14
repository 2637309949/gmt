package handler

import (
	"comm/define"
	"comm/logger"
	"comm/timer"
	"context"
	"fmt"
	"proto/aggregate"
)

func (s *Handler) UploadDoc(ctx context.Context, req *aggregate.UploadDocReq, resp *aggregate.UploadDocRes) error {
	logger.Info("Received UploadDoc request")

	marker := timer.NewMark()
	defer marker.Init("UploadDoc")

	doc := *req
	logger.Infof("%+v", doc)
	ersp, err := s.EsClient.Update().
		Index(define.DocIndex).
		Type(define.DocIndex).
		Id(fmt.Sprintf("%v", req.DocId)).
		Doc(doc).
		DocAsUpsert(true).Do(ctx)
	if err != nil {
		logger.Errorf("UploadDoc fail. %s %+v", err.Error(), ersp)
		return fmt.Errorf("UploadDoc fail. %s %+v", err.Error(), ersp)
	}
	return nil
}
