package handler

import (
	"comm/logger"
	"context"
	"test-service/srv/types"
	"time"

	"github.com/xormplus/xorm"
)

// EntityAddDB defined TODO
func (h *Handler) EntityAddDB(ctx context.Context, db *xorm.Engine, item *types.Entity) error {
	logger.Info("Received EntityAddDB request")

	item.CreateTime = time.Now()
	ret, err := db.Insert(item)
	if err != nil {
		return err
	}
	item.ID = ret
	return err
}

// EntityDelDB defined TODO
func (h *Handler) EntityDelDB(ctx context.Context, db *xorm.Engine, where *types.Entity) error {
	logger.Info("Received EntityDelDB request")

	ret, err := db.ID(where.ID).Update(&types.Entity{
		UpdateTime: time.Now(),
		IsDelete:   1,
	})
	if err != nil {
		return err
	}
	where.ID = ret
	return err
}

// EntityUpdateDB defined TODO
func (h *Handler) EntityUpdateDB(ctx context.Context, db *xorm.Engine, item *types.Entity) error {
	logger.Info("Received EntityUpdateDB request")

	_, err := db.ID(item.ID).Update(item)
	if err != nil {
		return err
	}
	return err
}

// EntityOneDB defined TODO
func (h *Handler) EntityOneDB(ctx context.Context, db *xorm.Engine, where *types.Entity, item *types.Entity) error {
	logger.Info("Received EntityOneDB request")

	ext, err := db.Get(item)
	if err != nil {
		return err
	}
	if !ext {
		return xorm.ErrNotExist
	}
	return err
}

// EntityPageDB defined TODO
func (h *Handler) EntityPageDB(ctx context.Context, db *xorm.Engine, where *types.Entity, list *[]types.Entity, totalRecord ...*int64) error {
	logger.Info("Received EntityPageDB request")

	err := db.Find(list, where)
	if err != nil {
		return err
	}
	if len(totalRecord) > 0 {
		ct, err := db.Count(where)
		if err != nil {
			return err
		}
		*totalRecord[0] = ct
	}
	return nil
}
