package class

import (
	"context"

	"github.com/jmoiron/sqlx"
	"remoteschool/smarthead/internal/postgres/models"
)

// Repository defines the required dependencies for Branch.
type Repository struct {
	DbConn *sqlx.DB
}

// NewRepository creates a new Repository that defines dependencies for Branch.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DbConn: db,
	}
}

// Branch represents a workflow.
type Class struct {
	ID          string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	SchoolOrder int    `json:"school_order"`
}

func FromModel(rec *models.Class) *Class {
	b := &Class{
		ID:          rec.ID,
		Name:        rec.Name,
		SchoolOrder: rec.SchoolOrder,
	}

	return b
}

// Response represents a workflow that is returned for display.
type Response struct {
	ID          string `json:"id" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name        string `json:"name" example:"Rocket Launch"`
	SchoolOrder int    `json:"school_order"`
}

// Response transforms Subject to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Class) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:          m.ID,
		Name:        m.Name,
		SchoolOrder: m.SchoolOrder,
	}

	return r
}

// Classes a list of Classes.
type Classes []*Class

// Response transforms a list of Branches to a list of Responses.
func (m *Classes) Response(ctx context.Context) []*Response {
	var l []*Response
	if m != nil && len(*m) > 0 {
		for _, n := range *m {
			l = append(l, n.Response(ctx))
		}
	}

	return l
}

// CreateRequest contains information needed to create a new Branch.
type CreateRequest struct {
	Name        string `json:"name" validate:"required"  example:"Rocket Launch"`
	SchoolOrder int    `json:"school_order"`
}

// ReadRequest defines the information needed to read a checklist.
type ReadRequest struct {
	ID              string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	IncludeArchived bool   `json:"include-archived" example:"false"`
}

// UpdateRequest defines what information may be provided to modify an existing
// Branch. All fields are optional so clients can send just the fields they want
// changed. It uses pointer fields so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank.
type UpdateRequest struct {
	ID          string  `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name        *string `json:"name,omitempty" validate:"omitempty,unique" example:"Rocket Launch to Moon"`
	SchoolOrder *int    `json:"school_order"`
}

// ArchiveRequest defines the information needed to archive a checklist. This will archive (soft-delete) the
// existing database entry.
type ArchiveRequest struct {
	ID string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
}

// DeleteRequest defines the information needed to delete a branch.
type DeleteRequest struct {
	ID string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
}

// FindRequest defines the possible options to search for branches. By default
// archived checklist will be excluded from response.
type FindRequest struct {
	Where           string        `json:"where" example:"name = ? and status = ?"`
	Args            []interface{} `json:"args" swaggertype:"array,string" example:"Moon Launch,active"`
	Order           []string      `json:"order" example:"created_at desc"`
	Limit           *uint         `json:"limit" example:"10"`
	Offset          *uint         `json:"offset" example:"20"`
	IncludeArchived bool          `json:"include-archived" example:"false"`
}
