// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/ent/namespace"
	"github.com/direktiv/direktiv/ent/workflow"
	"github.com/google/uuid"
)

// Workflow is the model entity for the Workflow schema.
type Workflow struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Revision holds the value of the "revision" field.
	Revision int `json:"revision,omitempty"`
	// Workflow holds the value of the "workflow" field.
	Workflow []byte `json:"workflow,omitempty"`
	// LogToEvents holds the value of the "logToEvents" field.
	LogToEvents string `json:"logToEvents,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowQuery when eager-loading is set.
	Edges               WorkflowEdges `json:"edges"`
	namespace_workflows *string
}

// WorkflowEdges holds the relations/edges for other nodes in the graph.
type WorkflowEdges struct {
	// Namespace holds the value of the namespace edge.
	Namespace *Namespace `json:"namespace,omitempty"`
	// Instances holds the value of the instances edge.
	Instances []*WorkflowInstance `json:"instances,omitempty"`
	// Wfevents holds the value of the wfevents edge.
	Wfevents []*WorkflowEvents `json:"wfevents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowEdges) NamespaceOrErr() (*Namespace, error) {
	if e.loadedTypes[0] {
		if e.Namespace == nil {
			// The edge namespace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: namespace.Label}
		}
		return e.Namespace, nil
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// InstancesOrErr returns the Instances value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowEdges) InstancesOrErr() ([]*WorkflowInstance, error) {
	if e.loadedTypes[1] {
		return e.Instances, nil
	}
	return nil, &NotLoadedError{edge: "instances"}
}

// WfeventsOrErr returns the Wfevents value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowEdges) WfeventsOrErr() ([]*WorkflowEvents, error) {
	if e.loadedTypes[2] {
		return e.Wfevents, nil
	}
	return nil, &NotLoadedError{edge: "wfevents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workflow) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflow.FieldWorkflow:
			values[i] = new([]byte)
		case workflow.FieldActive:
			values[i] = new(sql.NullBool)
		case workflow.FieldRevision:
			values[i] = new(sql.NullInt64)
		case workflow.FieldName, workflow.FieldDescription, workflow.FieldLogToEvents:
			values[i] = new(sql.NullString)
		case workflow.FieldCreated:
			values[i] = new(sql.NullTime)
		case workflow.FieldID:
			values[i] = new(uuid.UUID)
		case workflow.ForeignKeys[0]: // namespace_workflows
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Workflow", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workflow fields.
func (w *Workflow) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflow.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				w.ID = *value
			}
		case workflow.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		case workflow.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				w.Created = value.Time
			}
		case workflow.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				w.Description = value.String
			}
		case workflow.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				w.Active = value.Bool
			}
		case workflow.FieldRevision:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revision", values[i])
			} else if value.Valid {
				w.Revision = int(value.Int64)
			}
		case workflow.FieldWorkflow:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field workflow", values[i])
			} else if value != nil {
				w.Workflow = *value
			}
		case workflow.FieldLogToEvents:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logToEvents", values[i])
			} else if value.Valid {
				w.LogToEvents = value.String
			}
		case workflow.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace_workflows", values[i])
			} else if value.Valid {
				w.namespace_workflows = new(string)
				*w.namespace_workflows = value.String
			}
		}
	}
	return nil
}

// QueryNamespace queries the "namespace" edge of the Workflow entity.
func (w *Workflow) QueryNamespace() *NamespaceQuery {
	return (&WorkflowClient{config: w.config}).QueryNamespace(w)
}

// QueryInstances queries the "instances" edge of the Workflow entity.
func (w *Workflow) QueryInstances() *WorkflowInstanceQuery {
	return (&WorkflowClient{config: w.config}).QueryInstances(w)
}

// QueryWfevents queries the "wfevents" edge of the Workflow entity.
func (w *Workflow) QueryWfevents() *WorkflowEventsQuery {
	return (&WorkflowClient{config: w.config}).QueryWfevents(w)
}

// Update returns a builder for updating this Workflow.
// Note that you need to call Workflow.Unwrap() before calling this method if this Workflow
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workflow) Update() *WorkflowUpdateOne {
	return (&WorkflowClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Workflow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Workflow) Unwrap() *Workflow {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Workflow is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workflow) String() string {
	var builder strings.Builder
	builder.WriteString("Workflow(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", name=")
	builder.WriteString(w.Name)
	builder.WriteString(", created=")
	builder.WriteString(w.Created.Format(time.ANSIC))
	builder.WriteString(", description=")
	builder.WriteString(w.Description)
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", w.Active))
	builder.WriteString(", revision=")
	builder.WriteString(fmt.Sprintf("%v", w.Revision))
	builder.WriteString(", workflow=")
	builder.WriteString(fmt.Sprintf("%v", w.Workflow))
	builder.WriteString(", logToEvents=")
	builder.WriteString(w.LogToEvents)
	builder.WriteByte(')')
	return builder.String()
}

// Workflows is a parsable slice of Workflow.
type Workflows []*Workflow

func (w Workflows) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
