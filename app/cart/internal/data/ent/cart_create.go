// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imperialPalaceMall/app/cart/internal/data/ent/cart"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CartCreate is the builder for creating a Cart entity.
type CartCreate struct {
	config
	mutation *CartMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (cc *CartCreate) SetUserID(i int64) *CartCreate {
	cc.mutation.SetUserID(i)
	return cc
}

// SetGoodsID sets the "goods_id" field.
func (cc *CartCreate) SetGoodsID(i int64) *CartCreate {
	cc.mutation.SetGoodsID(i)
	return cc
}

// SetGoodsSkuID sets the "goods_sku_id" field.
func (cc *CartCreate) SetGoodsSkuID(i int64) *CartCreate {
	cc.mutation.SetGoodsSkuID(i)
	return cc
}

// SetGoodsSkuDesc sets the "goods_sku_desc" field.
func (cc *CartCreate) SetGoodsSkuDesc(s string) *CartCreate {
	cc.mutation.SetGoodsSkuDesc(s)
	return cc
}

// SetNum sets the "num" field.
func (cc *CartCreate) SetNum(i int32) *CartCreate {
	cc.mutation.SetNum(i)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CartCreate) SetCreatedAt(t time.Time) *CartCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CartCreate) SetNillableCreatedAt(t *time.Time) *CartCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CartCreate) SetUpdatedAt(t time.Time) *CartCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CartCreate) SetNillableUpdatedAt(t *time.Time) *CartCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CartCreate) SetID(i int64) *CartCreate {
	cc.mutation.SetID(i)
	return cc
}

// Mutation returns the CartMutation object of the builder.
func (cc *CartCreate) Mutation() *CartMutation {
	return cc.mutation
}

// Save creates the Cart in the database.
func (cc *CartCreate) Save(ctx context.Context) (*Cart, error) {
	var (
		err  error
		node *Cart
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cart)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CartMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CartCreate) SaveX(ctx context.Context) *Cart {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CartCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CartCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CartCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := cart.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := cart.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CartCreate) check() error {
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Cart.user_id"`)}
	}
	if _, ok := cc.mutation.GoodsID(); !ok {
		return &ValidationError{Name: "goods_id", err: errors.New(`ent: missing required field "Cart.goods_id"`)}
	}
	if _, ok := cc.mutation.GoodsSkuID(); !ok {
		return &ValidationError{Name: "goods_sku_id", err: errors.New(`ent: missing required field "Cart.goods_sku_id"`)}
	}
	if _, ok := cc.mutation.GoodsSkuDesc(); !ok {
		return &ValidationError{Name: "goods_sku_desc", err: errors.New(`ent: missing required field "Cart.goods_sku_desc"`)}
	}
	if _, ok := cc.mutation.Num(); !ok {
		return &ValidationError{Name: "num", err: errors.New(`ent: missing required field "Cart.num"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Cart.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Cart.updated_at"`)}
	}
	return nil
}

func (cc *CartCreate) sqlSave(ctx context.Context) (*Cart, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (cc *CartCreate) createSpec() (*Cart, *sqlgraph.CreateSpec) {
	var (
		_node = &Cart{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cart.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: cart.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := cc.mutation.GoodsID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsID,
		})
		_node.GoodsID = value
	}
	if value, ok := cc.mutation.GoodsSkuID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsSkuID,
		})
		_node.GoodsSkuID = value
	}
	if value, ok := cc.mutation.GoodsSkuDesc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cart.FieldGoodsSkuDesc,
		})
		_node.GoodsSkuDesc = value
	}
	if value, ok := cc.mutation.Num(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: cart.FieldNum,
		})
		_node.Num = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cart.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cart.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// CartCreateBulk is the builder for creating many Cart entities in bulk.
type CartCreateBulk struct {
	config
	builders []*CartCreate
}

// Save creates the Cart entities in the database.
func (ccb *CartCreateBulk) Save(ctx context.Context) ([]*Cart, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cart, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CartMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CartCreateBulk) SaveX(ctx context.Context) []*Cart {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CartCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CartCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
