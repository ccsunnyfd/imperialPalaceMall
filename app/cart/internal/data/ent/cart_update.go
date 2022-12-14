// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imperialPalaceMall/app/cart/internal/data/ent/cart"
	"imperialPalaceMall/app/cart/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CartUpdate is the builder for updating Cart entities.
type CartUpdate struct {
	config
	hooks    []Hook
	mutation *CartMutation
}

// Where appends a list predicates to the CartUpdate builder.
func (cu *CartUpdate) Where(ps ...predicate.Cart) *CartUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CartUpdate) SetUserID(i int64) *CartUpdate {
	cu.mutation.ResetUserID()
	cu.mutation.SetUserID(i)
	return cu
}

// AddUserID adds i to the "user_id" field.
func (cu *CartUpdate) AddUserID(i int64) *CartUpdate {
	cu.mutation.AddUserID(i)
	return cu
}

// SetGoodsID sets the "goods_id" field.
func (cu *CartUpdate) SetGoodsID(i int64) *CartUpdate {
	cu.mutation.ResetGoodsID()
	cu.mutation.SetGoodsID(i)
	return cu
}

// AddGoodsID adds i to the "goods_id" field.
func (cu *CartUpdate) AddGoodsID(i int64) *CartUpdate {
	cu.mutation.AddGoodsID(i)
	return cu
}

// SetGoodsSkuID sets the "goods_sku_id" field.
func (cu *CartUpdate) SetGoodsSkuID(i int64) *CartUpdate {
	cu.mutation.ResetGoodsSkuID()
	cu.mutation.SetGoodsSkuID(i)
	return cu
}

// AddGoodsSkuID adds i to the "goods_sku_id" field.
func (cu *CartUpdate) AddGoodsSkuID(i int64) *CartUpdate {
	cu.mutation.AddGoodsSkuID(i)
	return cu
}

// SetGoodsSkuDesc sets the "goods_sku_desc" field.
func (cu *CartUpdate) SetGoodsSkuDesc(s string) *CartUpdate {
	cu.mutation.SetGoodsSkuDesc(s)
	return cu
}

// SetNum sets the "num" field.
func (cu *CartUpdate) SetNum(i int32) *CartUpdate {
	cu.mutation.ResetNum()
	cu.mutation.SetNum(i)
	return cu
}

// AddNum adds i to the "num" field.
func (cu *CartUpdate) AddNum(i int32) *CartUpdate {
	cu.mutation.AddNum(i)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CartUpdate) SetCreatedAt(t time.Time) *CartUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CartUpdate) SetNillableCreatedAt(t *time.Time) *CartUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CartUpdate) SetUpdatedAt(t time.Time) *CartUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *CartUpdate) SetNillableUpdatedAt(t *time.Time) *CartUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// Mutation returns the CartMutation object of the builder.
func (cu *CartUpdate) Mutation() *CartMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CartUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CartUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CartUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CartUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cart.Table,
			Columns: cart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: cart.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldUserID,
		})
	}
	if value, ok := cu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldUserID,
		})
	}
	if value, ok := cu.mutation.GoodsID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsID,
		})
	}
	if value, ok := cu.mutation.AddedGoodsID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsID,
		})
	}
	if value, ok := cu.mutation.GoodsSkuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsSkuID,
		})
	}
	if value, ok := cu.mutation.AddedGoodsSkuID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsSkuID,
		})
	}
	if value, ok := cu.mutation.GoodsSkuDesc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cart.FieldGoodsSkuDesc,
		})
	}
	if value, ok := cu.mutation.Num(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: cart.FieldNum,
		})
	}
	if value, ok := cu.mutation.AddedNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: cart.FieldNum,
		})
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cart.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cart.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CartUpdateOne is the builder for updating a single Cart entity.
type CartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartMutation
}

// SetUserID sets the "user_id" field.
func (cuo *CartUpdateOne) SetUserID(i int64) *CartUpdateOne {
	cuo.mutation.ResetUserID()
	cuo.mutation.SetUserID(i)
	return cuo
}

// AddUserID adds i to the "user_id" field.
func (cuo *CartUpdateOne) AddUserID(i int64) *CartUpdateOne {
	cuo.mutation.AddUserID(i)
	return cuo
}

// SetGoodsID sets the "goods_id" field.
func (cuo *CartUpdateOne) SetGoodsID(i int64) *CartUpdateOne {
	cuo.mutation.ResetGoodsID()
	cuo.mutation.SetGoodsID(i)
	return cuo
}

// AddGoodsID adds i to the "goods_id" field.
func (cuo *CartUpdateOne) AddGoodsID(i int64) *CartUpdateOne {
	cuo.mutation.AddGoodsID(i)
	return cuo
}

// SetGoodsSkuID sets the "goods_sku_id" field.
func (cuo *CartUpdateOne) SetGoodsSkuID(i int64) *CartUpdateOne {
	cuo.mutation.ResetGoodsSkuID()
	cuo.mutation.SetGoodsSkuID(i)
	return cuo
}

// AddGoodsSkuID adds i to the "goods_sku_id" field.
func (cuo *CartUpdateOne) AddGoodsSkuID(i int64) *CartUpdateOne {
	cuo.mutation.AddGoodsSkuID(i)
	return cuo
}

// SetGoodsSkuDesc sets the "goods_sku_desc" field.
func (cuo *CartUpdateOne) SetGoodsSkuDesc(s string) *CartUpdateOne {
	cuo.mutation.SetGoodsSkuDesc(s)
	return cuo
}

// SetNum sets the "num" field.
func (cuo *CartUpdateOne) SetNum(i int32) *CartUpdateOne {
	cuo.mutation.ResetNum()
	cuo.mutation.SetNum(i)
	return cuo
}

// AddNum adds i to the "num" field.
func (cuo *CartUpdateOne) AddNum(i int32) *CartUpdateOne {
	cuo.mutation.AddNum(i)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CartUpdateOne) SetCreatedAt(t time.Time) *CartUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableCreatedAt(t *time.Time) *CartUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CartUpdateOne) SetUpdatedAt(t time.Time) *CartUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableUpdatedAt(t *time.Time) *CartUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// Mutation returns the CartMutation object of the builder.
func (cuo *CartUpdateOne) Mutation() *CartMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CartUpdateOne) Select(field string, fields ...string) *CartUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cart entity.
func (cuo *CartUpdateOne) Save(ctx context.Context) (*Cart, error) {
	var (
		err  error
		node *Cart
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (cuo *CartUpdateOne) SaveX(ctx context.Context) *Cart {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CartUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CartUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CartUpdateOne) sqlSave(ctx context.Context) (_node *Cart, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cart.Table,
			Columns: cart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: cart.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cart.FieldID)
		for _, f := range fields {
			if !cart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldUserID,
		})
	}
	if value, ok := cuo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldUserID,
		})
	}
	if value, ok := cuo.mutation.GoodsID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsID,
		})
	}
	if value, ok := cuo.mutation.AddedGoodsID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsID,
		})
	}
	if value, ok := cuo.mutation.GoodsSkuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsSkuID,
		})
	}
	if value, ok := cuo.mutation.AddedGoodsSkuID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cart.FieldGoodsSkuID,
		})
	}
	if value, ok := cuo.mutation.GoodsSkuDesc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cart.FieldGoodsSkuDesc,
		})
	}
	if value, ok := cuo.mutation.Num(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: cart.FieldNum,
		})
	}
	if value, ok := cuo.mutation.AddedNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: cart.FieldNum,
		})
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cart.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cart.FieldUpdatedAt,
		})
	}
	_node = &Cart{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
