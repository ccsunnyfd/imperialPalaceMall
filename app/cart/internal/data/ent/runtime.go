// Code generated by ent, DO NOT EDIT.

package ent

import (
	"imperialPalaceMall/app/cart/internal/data/ent/cart"
	"imperialPalaceMall/app/cart/internal/data/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cartFields := schema.Cart{}.Fields()
	_ = cartFields
	// cartDescCreatedAt is the schema descriptor for created_at field.
	cartDescCreatedAt := cartFields[6].Descriptor()
	// cart.DefaultCreatedAt holds the default value on creation for the created_at field.
	cart.DefaultCreatedAt = cartDescCreatedAt.Default.(func() time.Time)
	// cartDescUpdatedAt is the schema descriptor for updated_at field.
	cartDescUpdatedAt := cartFields[7].Descriptor()
	// cart.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cart.DefaultUpdatedAt = cartDescUpdatedAt.Default.(func() time.Time)
}
