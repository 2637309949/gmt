package handler

import (
	"comm/logger"
	"context"
	"{{.name}}-service/srv/types"

	"github.com/xormplus/xorm"
)

// {{toTitle .proto.Name}}AddDB defined TODO
func (h *Handler) {{toTitle .proto.Name}}AddDB(ctx context.Context, db *xorm.Engine, item *types.{{toTitle .proto.Name}}) error {
	logger.Info(ctx, "Received {{toTitle .proto.Name}}AddDB request")

	// item.CreateTime = time.Now()
	ret, err := db.Insert(item)
	if err != nil {
		return err
	}
	item.ID = uint64(ret)
	return err
}

// {{toTitle .proto.Name}}DelDB defined TODO
func (h *Handler) {{toTitle .proto.Name}}DelDB(ctx context.Context, db *xorm.Engine, where *types.{{toTitle .proto.Name}}) error {
	logger.Info(ctx, "Received {{toTitle .proto.Name}}DelDB request")

	ret, err := db.ID(where.ID).Update(&types.{{toTitle .proto.Name}}{
		// UpdateTime: time.Now(),
		// IsDelete:   1,
	})
	if err != nil {
		return err
	}
	where.ID = uint64(ret)
	return err
}

// {{toTitle .proto.Name}}UpdateDB defined TODO
func (h *Handler) {{toTitle .proto.Name}}UpdateDB(ctx context.Context, db *xorm.Engine, item *types.{{toTitle .proto.Name}}) error {
	logger.Info(ctx, "Received {{toTitle .proto.Name}}UpdateDB request")

	_, err := db.ID(item.ID).Update(item)
	if err != nil {
		return err
	}
	return err
}

// {{toTitle .proto.Name}}OneDB defined TODO
func (h *Handler) {{toTitle .proto.Name}}OneDB(ctx context.Context, db *xorm.Engine, where *types.{{toTitle .proto.Name}}, item *types.{{toTitle .proto.Name}}) error {
	logger.Info(ctx, "Received {{toTitle .proto.Name}}OneDB request")

	ext, err := db.Get(item)
	if err != nil {
		return err
	}
	if !ext {
		return xorm.ErrNotExist
	}
	return err
}

// {{toTitle .proto.Name}}PageDB defined TODO
func (h *Handler) {{toTitle .proto.Name}}PageDB(ctx context.Context, db *xorm.Engine, where *types.{{toTitle .proto.Name}}, list *[]types.{{toTitle .proto.Name}}, totalRecord ...*uint64) error {
	logger.Info(ctx, "Received {{toTitle .proto.Name}}PageDB request")

	err := db.Find(list, where)
	if err != nil {
		return err
	}
	if len(totalRecord) > 0 {
		ct, err := db.Count(where)
		if err != nil {
			return err
		}
		*totalRecord[0] = uint64(ct)
	}
	return nil
}
