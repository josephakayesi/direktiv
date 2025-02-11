// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/ent/workflow"
	"github.com/direktiv/direktiv/ent/workflowinstance"
	"github.com/google/uuid"
)

// WorkflowInstance is the model entity for the WorkflowInstance schema.
type WorkflowInstance struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// InstanceID holds the value of the "instanceID" field.
	InstanceID string `json:"instanceID,omitempty"`
	// InvokedBy holds the value of the "invokedBy" field.
	InvokedBy string `json:"invokedBy,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Revision holds the value of the "revision" field.
	Revision int `json:"revision,omitempty"`
	// BeginTime holds the value of the "beginTime" field.
	BeginTime time.Time `json:"beginTime,omitempty"`
	// EndTime holds the value of the "endTime" field.
	EndTime time.Time `json:"endTime,omitempty"`
	// Flow holds the value of the "flow" field.
	Flow []string `json:"flow,omitempty"`
	// Input holds the value of the "input" field.
	Input string `json:"input,omitempty"`
	// Output holds the value of the "output" field.
	Output string `json:"output,omitempty"`
	// StateData holds the value of the "stateData" field.
	StateData string `json:"stateData,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
	// Deadline holds the value of the "deadline" field.
	Deadline time.Time `json:"deadline,omitempty"`
	// Attempts holds the value of the "attempts" field.
	Attempts int `json:"attempts,omitempty"`
	// ErrorCode holds the value of the "errorCode" field.
	ErrorCode string `json:"errorCode,omitempty"`
	// ErrorMessage holds the value of the "errorMessage" field.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// StateBeginTime holds the value of the "stateBeginTime" field.
	StateBeginTime time.Time `json:"stateBeginTime,omitempty"`
	// Controller holds the value of the "controller" field.
	Controller string `json:"controller,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowInstanceQuery when eager-loading is set.
	Edges              WorkflowInstanceEdges `json:"edges"`
	workflow_instances *uuid.UUID
}

// WorkflowInstanceEdges holds the relations/edges for other nodes in the graph.
type WorkflowInstanceEdges struct {
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// Instance holds the value of the instance edge.
	Instance []*WorkflowEvents `json:"instance,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowInstanceEdges) WorkflowOrErr() (*Workflow, error) {
	if e.loadedTypes[0] {
		if e.Workflow == nil {
			// The edge workflow was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workflow.Label}
		}
		return e.Workflow, nil
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// InstanceOrErr returns the Instance value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowInstanceEdges) InstanceOrErr() ([]*WorkflowEvents, error) {
	if e.loadedTypes[1] {
		return e.Instance, nil
	}
	return nil, &NotLoadedError{edge: "instance"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkflowInstance) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflowinstance.FieldFlow:
			values[i] = new([]byte)
		case workflowinstance.FieldID, workflowinstance.FieldRevision, workflowinstance.FieldAttempts:
			values[i] = new(sql.NullInt64)
		case workflowinstance.FieldInstanceID, workflowinstance.FieldInvokedBy, workflowinstance.FieldStatus, workflowinstance.FieldInput, workflowinstance.FieldOutput, workflowinstance.FieldStateData, workflowinstance.FieldMemory, workflowinstance.FieldErrorCode, workflowinstance.FieldErrorMessage, workflowinstance.FieldController:
			values[i] = new(sql.NullString)
		case workflowinstance.FieldBeginTime, workflowinstance.FieldEndTime, workflowinstance.FieldDeadline, workflowinstance.FieldStateBeginTime:
			values[i] = new(sql.NullTime)
		case workflowinstance.ForeignKeys[0]: // workflow_instances
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkflowInstance", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkflowInstance fields.
func (wi *WorkflowInstance) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflowinstance.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wi.ID = int(value.Int64)
		case workflowinstance.FieldInstanceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instanceID", values[i])
			} else if value.Valid {
				wi.InstanceID = value.String
			}
		case workflowinstance.FieldInvokedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invokedBy", values[i])
			} else if value.Valid {
				wi.InvokedBy = value.String
			}
		case workflowinstance.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				wi.Status = value.String
			}
		case workflowinstance.FieldRevision:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revision", values[i])
			} else if value.Valid {
				wi.Revision = int(value.Int64)
			}
		case workflowinstance.FieldBeginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field beginTime", values[i])
			} else if value.Valid {
				wi.BeginTime = value.Time
			}
		case workflowinstance.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field endTime", values[i])
			} else if value.Valid {
				wi.EndTime = value.Time
			}
		case workflowinstance.FieldFlow:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field flow", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wi.Flow); err != nil {
					return fmt.Errorf("unmarshal field flow: %w", err)
				}
			}
		case workflowinstance.FieldInput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field input", values[i])
			} else if value.Valid {
				wi.Input = value.String
			}
		case workflowinstance.FieldOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[i])
			} else if value.Valid {
				wi.Output = value.String
			}
		case workflowinstance.FieldStateData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stateData", values[i])
			} else if value.Valid {
				wi.StateData = value.String
			}
		case workflowinstance.FieldMemory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				wi.Memory = value.String
			}
		case workflowinstance.FieldDeadline:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deadline", values[i])
			} else if value.Valid {
				wi.Deadline = value.Time
			}
		case workflowinstance.FieldAttempts:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attempts", values[i])
			} else if value.Valid {
				wi.Attempts = int(value.Int64)
			}
		case workflowinstance.FieldErrorCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field errorCode", values[i])
			} else if value.Valid {
				wi.ErrorCode = value.String
			}
		case workflowinstance.FieldErrorMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field errorMessage", values[i])
			} else if value.Valid {
				wi.ErrorMessage = value.String
			}
		case workflowinstance.FieldStateBeginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field stateBeginTime", values[i])
			} else if value.Valid {
				wi.StateBeginTime = value.Time
			}
		case workflowinstance.FieldController:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field controller", values[i])
			} else if value.Valid {
				wi.Controller = value.String
			}
		case workflowinstance.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_instances", values[i])
			} else if value.Valid {
				wi.workflow_instances = new(uuid.UUID)
				*wi.workflow_instances = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryWorkflow queries the "workflow" edge of the WorkflowInstance entity.
func (wi *WorkflowInstance) QueryWorkflow() *WorkflowQuery {
	return (&WorkflowInstanceClient{config: wi.config}).QueryWorkflow(wi)
}

// QueryInstance queries the "instance" edge of the WorkflowInstance entity.
func (wi *WorkflowInstance) QueryInstance() *WorkflowEventsQuery {
	return (&WorkflowInstanceClient{config: wi.config}).QueryInstance(wi)
}

// Update returns a builder for updating this WorkflowInstance.
// Note that you need to call WorkflowInstance.Unwrap() before calling this method if this WorkflowInstance
// was returned from a transaction, and the transaction was committed or rolled back.
func (wi *WorkflowInstance) Update() *WorkflowInstanceUpdateOne {
	return (&WorkflowInstanceClient{config: wi.config}).UpdateOne(wi)
}

// Unwrap unwraps the WorkflowInstance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wi *WorkflowInstance) Unwrap() *WorkflowInstance {
	tx, ok := wi.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkflowInstance is not a transactional entity")
	}
	wi.config.driver = tx.drv
	return wi
}

// String implements the fmt.Stringer.
func (wi *WorkflowInstance) String() string {
	var builder strings.Builder
	builder.WriteString("WorkflowInstance(")
	builder.WriteString(fmt.Sprintf("id=%v", wi.ID))
	builder.WriteString(", instanceID=")
	builder.WriteString(wi.InstanceID)
	builder.WriteString(", invokedBy=")
	builder.WriteString(wi.InvokedBy)
	builder.WriteString(", status=")
	builder.WriteString(wi.Status)
	builder.WriteString(", revision=")
	builder.WriteString(fmt.Sprintf("%v", wi.Revision))
	builder.WriteString(", beginTime=")
	builder.WriteString(wi.BeginTime.Format(time.ANSIC))
	builder.WriteString(", endTime=")
	builder.WriteString(wi.EndTime.Format(time.ANSIC))
	builder.WriteString(", flow=")
	builder.WriteString(fmt.Sprintf("%v", wi.Flow))
	builder.WriteString(", input=")
	builder.WriteString(wi.Input)
	builder.WriteString(", output=")
	builder.WriteString(wi.Output)
	builder.WriteString(", stateData=")
	builder.WriteString(wi.StateData)
	builder.WriteString(", memory=")
	builder.WriteString(wi.Memory)
	builder.WriteString(", deadline=")
	builder.WriteString(wi.Deadline.Format(time.ANSIC))
	builder.WriteString(", attempts=")
	builder.WriteString(fmt.Sprintf("%v", wi.Attempts))
	builder.WriteString(", errorCode=")
	builder.WriteString(wi.ErrorCode)
	builder.WriteString(", errorMessage=")
	builder.WriteString(wi.ErrorMessage)
	builder.WriteString(", stateBeginTime=")
	builder.WriteString(wi.StateBeginTime.Format(time.ANSIC))
	builder.WriteString(", controller=")
	builder.WriteString(wi.Controller)
	builder.WriteByte(')')
	return builder.String()
}

// WorkflowInstances is a parsable slice of WorkflowInstance.
type WorkflowInstances []*WorkflowInstance

func (wi WorkflowInstances) config(cfg config) {
	for _i := range wi {
		wi[_i].config = cfg
	}
}
