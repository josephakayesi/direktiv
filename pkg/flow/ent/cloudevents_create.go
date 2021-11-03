// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudevents"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// CloudEventsCreate is the builder for creating a CloudEvents entity.
type CloudEventsCreate struct {
	config
	mutation *CloudEventsMutation
	hooks    []Hook
}

// SetEventId sets the "eventId" field.
func (cec *CloudEventsCreate) SetEventId(s string) *CloudEventsCreate {
	cec.mutation.SetEventId(s)
	return cec
}

// SetEvent sets the "event" field.
func (cec *CloudEventsCreate) SetEvent(e event.Event) *CloudEventsCreate {
	cec.mutation.SetEvent(e)
	return cec
}

// SetFire sets the "fire" field.
func (cec *CloudEventsCreate) SetFire(t time.Time) *CloudEventsCreate {
	cec.mutation.SetFire(t)
	return cec
}

// SetNillableFire sets the "fire" field if the given value is not nil.
func (cec *CloudEventsCreate) SetNillableFire(t *time.Time) *CloudEventsCreate {
	if t != nil {
		cec.SetFire(*t)
	}
	return cec
}

// SetCreated sets the "created" field.
func (cec *CloudEventsCreate) SetCreated(t time.Time) *CloudEventsCreate {
	cec.mutation.SetCreated(t)
	return cec
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (cec *CloudEventsCreate) SetNillableCreated(t *time.Time) *CloudEventsCreate {
	if t != nil {
		cec.SetCreated(*t)
	}
	return cec
}

// SetProcessed sets the "processed" field.
func (cec *CloudEventsCreate) SetProcessed(b bool) *CloudEventsCreate {
	cec.mutation.SetProcessed(b)
	return cec
}

// SetID sets the "id" field.
func (cec *CloudEventsCreate) SetID(u uuid.UUID) *CloudEventsCreate {
	cec.mutation.SetID(u)
	return cec
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (cec *CloudEventsCreate) SetNamespaceID(id uuid.UUID) *CloudEventsCreate {
	cec.mutation.SetNamespaceID(id)
	return cec
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (cec *CloudEventsCreate) SetNamespace(n *Namespace) *CloudEventsCreate {
	return cec.SetNamespaceID(n.ID)
}

// Mutation returns the CloudEventsMutation object of the builder.
func (cec *CloudEventsCreate) Mutation() *CloudEventsMutation {
	return cec.mutation
}

// Save creates the CloudEvents in the database.
func (cec *CloudEventsCreate) Save(ctx context.Context) (*CloudEvents, error) {
	var (
		err  error
		node *CloudEvents
	)
	cec.defaults()
	if len(cec.hooks) == 0 {
		if err = cec.check(); err != nil {
			return nil, err
		}
		node, err = cec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CloudEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cec.check(); err != nil {
				return nil, err
			}
			cec.mutation = mutation
			if node, err = cec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cec.hooks) - 1; i >= 0; i-- {
			if cec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cec *CloudEventsCreate) SaveX(ctx context.Context) *CloudEvents {
	v, err := cec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cec *CloudEventsCreate) Exec(ctx context.Context) error {
	_, err := cec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cec *CloudEventsCreate) ExecX(ctx context.Context) {
	if err := cec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cec *CloudEventsCreate) defaults() {
	if _, ok := cec.mutation.Fire(); !ok {
		v := cloudevents.DefaultFire()
		cec.mutation.SetFire(v)
	}
	if _, ok := cec.mutation.Created(); !ok {
		v := cloudevents.DefaultCreated()
		cec.mutation.SetCreated(v)
	}
	if _, ok := cec.mutation.ID(); !ok {
		v := cloudevents.DefaultID()
		cec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cec *CloudEventsCreate) check() error {
	if _, ok := cec.mutation.EventId(); !ok {
		return &ValidationError{Name: "eventId", err: errors.New(`ent: missing required field "eventId"`)}
	}
	if v, ok := cec.mutation.EventId(); ok {
		if err := cloudevents.EventIdValidator(v); err != nil {
			return &ValidationError{Name: "eventId", err: fmt.Errorf(`ent: validator failed for field "eventId": %w`, err)}
		}
	}
	if _, ok := cec.mutation.Event(); !ok {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required field "event"`)}
	}
	if _, ok := cec.mutation.Fire(); !ok {
		return &ValidationError{Name: "fire", err: errors.New(`ent: missing required field "fire"`)}
	}
	if _, ok := cec.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "created"`)}
	}
	if _, ok := cec.mutation.Processed(); !ok {
		return &ValidationError{Name: "processed", err: errors.New(`ent: missing required field "processed"`)}
	}
	if _, ok := cec.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New("ent: missing required edge \"namespace\"")}
	}
	return nil
}

func (cec *CloudEventsCreate) sqlSave(ctx context.Context) (*CloudEvents, error) {
	_node, _spec := cec.createSpec()
	if err := sqlgraph.CreateNode(ctx, cec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (cec *CloudEventsCreate) createSpec() (*CloudEvents, *sqlgraph.CreateSpec) {
	var (
		_node = &CloudEvents{config: cec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cloudevents.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cloudevents.FieldID,
			},
		}
	)
	if id, ok := cec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cec.mutation.EventId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cloudevents.FieldEventId,
		})
		_node.EventId = value
	}
	if value, ok := cec.mutation.Event(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: cloudevents.FieldEvent,
		})
		_node.Event = value
	}
	if value, ok := cec.mutation.Fire(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cloudevents.FieldFire,
		})
		_node.Fire = value
	}
	if value, ok := cec.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cloudevents.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := cec.mutation.Processed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cloudevents.FieldProcessed,
		})
		_node.Processed = value
	}
	if nodes := cec.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cloudevents.NamespaceTable,
			Columns: []string{cloudevents.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_cloudevents = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CloudEventsCreateBulk is the builder for creating many CloudEvents entities in bulk.
type CloudEventsCreateBulk struct {
	config
	builders []*CloudEventsCreate
}

// Save creates the CloudEvents entities in the database.
func (cecb *CloudEventsCreateBulk) Save(ctx context.Context) ([]*CloudEvents, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cecb.builders))
	nodes := make([]*CloudEvents, len(cecb.builders))
	mutators := make([]Mutator, len(cecb.builders))
	for i := range cecb.builders {
		func(i int, root context.Context) {
			builder := cecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CloudEventsMutation)
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
					_, err = mutators[i+1].Mutate(root, cecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, cecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cecb *CloudEventsCreateBulk) SaveX(ctx context.Context) []*CloudEvents {
	v, err := cecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cecb *CloudEventsCreateBulk) Exec(ctx context.Context) error {
	_, err := cecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cecb *CloudEventsCreateBulk) ExecX(ctx context.Context) {
	if err := cecb.Exec(ctx); err != nil {
		panic(err)
	}
}
