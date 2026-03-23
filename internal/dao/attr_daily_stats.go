package dao

import (
	"god-help-service/internal/dao/internal"
)

// attrDailyStatsDao is the data access object for the table attr_daily_stats.
type attrDailyStatsDao struct {
	*internal.AttrDailyStatsDao
}

var (
	// AttrDailyStats is a globally accessible object for table attr_daily_stats operations.
	AttrDailyStats = attrDailyStatsDao{internal.NewAttrDailyStatsDao()}
)
