// Code generated by ent, DO NOT EDIT.

package cart

import (
	"time"
)

const (
	// Label holds the string label denoting the cart type in the database.
	Label = "cart"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodsID holds the string denoting the goods_id field in the database.
	FieldGoodsID = "goods_id"
	// FieldGoodsSkuID holds the string denoting the goods_sku_id field in the database.
	FieldGoodsSkuID = "goods_sku_id"
	// FieldGoodsSkuDesc holds the string denoting the goods_sku_desc field in the database.
	FieldGoodsSkuDesc = "goods_sku_desc"
	// FieldNum holds the string denoting the num field in the database.
	FieldNum = "num"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the cart in the database.
	Table = "carts"
)

// Columns holds all SQL columns for cart fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldGoodsID,
	FieldGoodsSkuID,
	FieldGoodsSkuDesc,
	FieldNum,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
