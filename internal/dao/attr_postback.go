package dao

import (
	"god-help-service/internal/dao/internal"
)

// attrPostbackDao is the data access object for the table attr_postback.
type attrPostbackDao struct {
	*internal.AttrPostbackDao
}

var (
	// AttrPostback is a globally accessible object for table attr_postback operations.
	AttrPostback = attrPostbackDao{internal.NewAttrPostbackDao()}
)
