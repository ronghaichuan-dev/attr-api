package dao

import (
	"god-help-service/internal/dao/internal"
)

// attrClickDao is the data access object for the table attr_click.
type attrClickDao struct {
	*internal.AttrClickDao
}

var (
	// AttrClick is a globally accessible object for table attr_click operations.
	AttrClick = attrClickDao{internal.NewAttrClickDao()}
)
