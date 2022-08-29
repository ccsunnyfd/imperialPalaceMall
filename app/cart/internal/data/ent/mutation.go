// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imperialPalaceMall/app/cart/internal/data/ent/cart"
	"imperialPalaceMall/app/cart/internal/data/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCart = "Cart"
)

// CartMutation represents an operation that mutates the Cart nodes in the graph.
type CartMutation struct {
	config
	op              Op
	typ             string
	id              *int64
	user_id         *int64
	adduser_id      *int64
	goods_id        *int64
	addgoods_id     *int64
	goods_sku_id    *int64
	addgoods_sku_id *int64
	goods_sku_desc  *string
	num             *int32
	addnum          *int32
	created_at      *time.Time
	updated_at      *time.Time
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*Cart, error)
	predicates      []predicate.Cart
}

var _ ent.Mutation = (*CartMutation)(nil)

// cartOption allows management of the mutation configuration using functional options.
type cartOption func(*CartMutation)

// newCartMutation creates new mutation for the Cart entity.
func newCartMutation(c config, op Op, opts ...cartOption) *CartMutation {
	m := &CartMutation{
		config:        c,
		op:            op,
		typ:           TypeCart,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCartID sets the ID field of the mutation.
func withCartID(id int64) cartOption {
	return func(m *CartMutation) {
		var (
			err   error
			once  sync.Once
			value *Cart
		)
		m.oldValue = func(ctx context.Context) (*Cart, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Cart.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCart sets the old Cart of the mutation.
func withCart(node *Cart) cartOption {
	return func(m *CartMutation) {
		m.oldValue = func(context.Context) (*Cart, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CartMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CartMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Cart entities.
func (m *CartMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CartMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CartMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Cart.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserID sets the "user_id" field.
func (m *CartMutation) SetUserID(i int64) {
	m.user_id = &i
	m.adduser_id = nil
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *CartMutation) UserID() (r int64, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Cart entity.
// If the Cart object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CartMutation) OldUserID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// AddUserID adds i to the "user_id" field.
func (m *CartMutation) AddUserID(i int64) {
	if m.adduser_id != nil {
		*m.adduser_id += i
	} else {
		m.adduser_id = &i
	}
}

// AddedUserID returns the value that was added to the "user_id" field in this mutation.
func (m *CartMutation) AddedUserID() (r int64, exists bool) {
	v := m.adduser_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserID resets all changes to the "user_id" field.
func (m *CartMutation) ResetUserID() {
	m.user_id = nil
	m.adduser_id = nil
}

// SetGoodsID sets the "goods_id" field.
func (m *CartMutation) SetGoodsID(i int64) {
	m.goods_id = &i
	m.addgoods_id = nil
}

// GoodsID returns the value of the "goods_id" field in the mutation.
func (m *CartMutation) GoodsID() (r int64, exists bool) {
	v := m.goods_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGoodsID returns the old "goods_id" field's value of the Cart entity.
// If the Cart object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CartMutation) OldGoodsID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGoodsID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGoodsID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGoodsID: %w", err)
	}
	return oldValue.GoodsID, nil
}

// AddGoodsID adds i to the "goods_id" field.
func (m *CartMutation) AddGoodsID(i int64) {
	if m.addgoods_id != nil {
		*m.addgoods_id += i
	} else {
		m.addgoods_id = &i
	}
}

// AddedGoodsID returns the value that was added to the "goods_id" field in this mutation.
func (m *CartMutation) AddedGoodsID() (r int64, exists bool) {
	v := m.addgoods_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetGoodsID resets all changes to the "goods_id" field.
func (m *CartMutation) ResetGoodsID() {
	m.goods_id = nil
	m.addgoods_id = nil
}

// SetGoodsSkuID sets the "goods_sku_id" field.
func (m *CartMutation) SetGoodsSkuID(i int64) {
	m.goods_sku_id = &i
	m.addgoods_sku_id = nil
}

// GoodsSkuID returns the value of the "goods_sku_id" field in the mutation.
func (m *CartMutation) GoodsSkuID() (r int64, exists bool) {
	v := m.goods_sku_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGoodsSkuID returns the old "goods_sku_id" field's value of the Cart entity.
// If the Cart object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CartMutation) OldGoodsSkuID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGoodsSkuID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGoodsSkuID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGoodsSkuID: %w", err)
	}
	return oldValue.GoodsSkuID, nil
}

// AddGoodsSkuID adds i to the "goods_sku_id" field.
func (m *CartMutation) AddGoodsSkuID(i int64) {
	if m.addgoods_sku_id != nil {
		*m.addgoods_sku_id += i
	} else {
		m.addgoods_sku_id = &i
	}
}

// AddedGoodsSkuID returns the value that was added to the "goods_sku_id" field in this mutation.
func (m *CartMutation) AddedGoodsSkuID() (r int64, exists bool) {
	v := m.addgoods_sku_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetGoodsSkuID resets all changes to the "goods_sku_id" field.
func (m *CartMutation) ResetGoodsSkuID() {
	m.goods_sku_id = nil
	m.addgoods_sku_id = nil
}

// SetGoodsSkuDesc sets the "goods_sku_desc" field.
func (m *CartMutation) SetGoodsSkuDesc(s string) {
	m.goods_sku_desc = &s
}

// GoodsSkuDesc returns the value of the "goods_sku_desc" field in the mutation.
func (m *CartMutation) GoodsSkuDesc() (r string, exists bool) {
	v := m.goods_sku_desc
	if v == nil {
		return
	}
	return *v, true
}

// OldGoodsSkuDesc returns the old "goods_sku_desc" field's value of the Cart entity.
// If the Cart object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CartMutation) OldGoodsSkuDesc(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGoodsSkuDesc is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGoodsSkuDesc requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGoodsSkuDesc: %w", err)
	}
	return oldValue.GoodsSkuDesc, nil
}

// ResetGoodsSkuDesc resets all changes to the "goods_sku_desc" field.
func (m *CartMutation) ResetGoodsSkuDesc() {
	m.goods_sku_desc = nil
}

// SetNum sets the "num" field.
func (m *CartMutation) SetNum(i int32) {
	m.num = &i
	m.addnum = nil
}

// Num returns the value of the "num" field in the mutation.
func (m *CartMutation) Num() (r int32, exists bool) {
	v := m.num
	if v == nil {
		return
	}
	return *v, true
}

// OldNum returns the old "num" field's value of the Cart entity.
// If the Cart object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CartMutation) OldNum(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNum is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNum requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNum: %w", err)
	}
	return oldValue.Num, nil
}

// AddNum adds i to the "num" field.
func (m *CartMutation) AddNum(i int32) {
	if m.addnum != nil {
		*m.addnum += i
	} else {
		m.addnum = &i
	}
}

// AddedNum returns the value that was added to the "num" field in this mutation.
func (m *CartMutation) AddedNum() (r int32, exists bool) {
	v := m.addnum
	if v == nil {
		return
	}
	return *v, true
}

// ResetNum resets all changes to the "num" field.
func (m *CartMutation) ResetNum() {
	m.num = nil
	m.addnum = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *CartMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *CartMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Cart entity.
// If the Cart object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CartMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *CartMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *CartMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *CartMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Cart entity.
// If the Cart object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CartMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *CartMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the CartMutation builder.
func (m *CartMutation) Where(ps ...predicate.Cart) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CartMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Cart).
func (m *CartMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CartMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.user_id != nil {
		fields = append(fields, cart.FieldUserID)
	}
	if m.goods_id != nil {
		fields = append(fields, cart.FieldGoodsID)
	}
	if m.goods_sku_id != nil {
		fields = append(fields, cart.FieldGoodsSkuID)
	}
	if m.goods_sku_desc != nil {
		fields = append(fields, cart.FieldGoodsSkuDesc)
	}
	if m.num != nil {
		fields = append(fields, cart.FieldNum)
	}
	if m.created_at != nil {
		fields = append(fields, cart.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, cart.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CartMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case cart.FieldUserID:
		return m.UserID()
	case cart.FieldGoodsID:
		return m.GoodsID()
	case cart.FieldGoodsSkuID:
		return m.GoodsSkuID()
	case cart.FieldGoodsSkuDesc:
		return m.GoodsSkuDesc()
	case cart.FieldNum:
		return m.Num()
	case cart.FieldCreatedAt:
		return m.CreatedAt()
	case cart.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CartMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case cart.FieldUserID:
		return m.OldUserID(ctx)
	case cart.FieldGoodsID:
		return m.OldGoodsID(ctx)
	case cart.FieldGoodsSkuID:
		return m.OldGoodsSkuID(ctx)
	case cart.FieldGoodsSkuDesc:
		return m.OldGoodsSkuDesc(ctx)
	case cart.FieldNum:
		return m.OldNum(ctx)
	case cart.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case cart.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Cart field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CartMutation) SetField(name string, value ent.Value) error {
	switch name {
	case cart.FieldUserID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case cart.FieldGoodsID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGoodsID(v)
		return nil
	case cart.FieldGoodsSkuID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGoodsSkuID(v)
		return nil
	case cart.FieldGoodsSkuDesc:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGoodsSkuDesc(v)
		return nil
	case cart.FieldNum:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNum(v)
		return nil
	case cart.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case cart.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Cart field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CartMutation) AddedFields() []string {
	var fields []string
	if m.adduser_id != nil {
		fields = append(fields, cart.FieldUserID)
	}
	if m.addgoods_id != nil {
		fields = append(fields, cart.FieldGoodsID)
	}
	if m.addgoods_sku_id != nil {
		fields = append(fields, cart.FieldGoodsSkuID)
	}
	if m.addnum != nil {
		fields = append(fields, cart.FieldNum)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CartMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case cart.FieldUserID:
		return m.AddedUserID()
	case cart.FieldGoodsID:
		return m.AddedGoodsID()
	case cart.FieldGoodsSkuID:
		return m.AddedGoodsSkuID()
	case cart.FieldNum:
		return m.AddedNum()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CartMutation) AddField(name string, value ent.Value) error {
	switch name {
	case cart.FieldUserID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserID(v)
		return nil
	case cart.FieldGoodsID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGoodsID(v)
		return nil
	case cart.FieldGoodsSkuID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGoodsSkuID(v)
		return nil
	case cart.FieldNum:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddNum(v)
		return nil
	}
	return fmt.Errorf("unknown Cart numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CartMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CartMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CartMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Cart nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CartMutation) ResetField(name string) error {
	switch name {
	case cart.FieldUserID:
		m.ResetUserID()
		return nil
	case cart.FieldGoodsID:
		m.ResetGoodsID()
		return nil
	case cart.FieldGoodsSkuID:
		m.ResetGoodsSkuID()
		return nil
	case cart.FieldGoodsSkuDesc:
		m.ResetGoodsSkuDesc()
		return nil
	case cart.FieldNum:
		m.ResetNum()
		return nil
	case cart.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case cart.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Cart field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CartMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CartMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CartMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CartMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CartMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CartMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CartMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Cart unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CartMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Cart edge %s", name)
}
