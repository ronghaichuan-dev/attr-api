package sharding

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func ShardModel(ctx context.Context, table string, year int, alias ...string) *gdb.Model {
	if len(alias) > 0 {
		return g.DB().Ctx(ctx).Model()
	}
	return g.DB().Ctx(ctx).Model(fmt.Sprintf("%s_%d", table, year))
}

func UnionAll(ctx context.Context, model ...*gdb.Model) *gdb.Model {
	return g.DB().UnionAll(model...).Ctx(ctx)
}
