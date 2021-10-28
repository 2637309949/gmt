package handler

import (
	"comm/logger"
	"context"
	"{{.name}}-service/srv/types"
	"time"

	"github.com/xormplus/xorm"
)

// {{toTitle .entity}}AddDB defined TODO
func (h *Handler) {{toTitle .entity}}AddDB(ctx context.Context, db *xorm.Engine, item *types.{{toTitle .entity}}) error {
	logger.Info("Received {{toTitle .entity}}AddDB request")

	item.CreateTime = time.Now()
	ret, err := db.Insert(item)
	if err != nil {
		return err
	}
	item.ID = ret
	return err
}

// {{toTitle .entity}}DelDB defined TODO
func (h *Handler) {{toTitle .entity}}DelDB(ctx context.Context, db *xorm.Engine, where *types.{{toTitle .entity}}) error {
	logger.Info("Received {{toTitle .entity}}DelDB request")

	ret, err := db.ID(where.ID).Update(&types.{{toTitle .entity}}{
		UpdateTime: time.Now(),
		IsDelete:   1,
	})
	if err != nil {
		return err
	}
	where.ID = ret
	return err
}

// {{toTitle .entity}}UpdateDB defined TODO
func (h *Handler) {{toTitle .entity}}UpdateDB(ctx context.Context, db *xorm.Engine, item *types.{{toTitle .entity}}) error {
	logger.Info("Received {{toTitle .entity}}UpdateDB request")

	_, err := db.ID(item.ID).Update(item)
	if err != nil {
		return err
	}
	return err
}

// {{toTitle .entity}}OneDB defined TODO
func (h *Handler) {{toTitle .entity}}OneDB(ctx context.Context, db *xorm.Engine, where *types.{{toTitle .entity}}, item *types.{{toTitle .entity}}) error {
	logger.Info("Received {{toTitle .entity}}OneDB request")

	ext, err := db.Get(item)
	if err != nil {
		return err
	}
	if !ext {
		return xorm.ErrNotExist
	}
	return err
}

// {{toTitle .entity}}PageDB defined TODO
func (h *Handler) {{toTitle .entity}}PageDB(ctx context.Context, db *xorm.Engine, where *types.{{toTitle .entity}}, list *[]types.{{toTitle .entity}}, totalRecord ...*int64) error {
	logger.Info("Received {{toTitle .entity}}PageDB request")

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
