package subclass

import (
	"context"

	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/postgres/models"

	"github.com/jmoiron/sqlx"
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

// Branch represents a Subclass.
type Subclass struct {
	ID          string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ClassID     int    `json:"class_id"`
	SchoolOrder int    `json:"school_order"`

	Class *class.Class `json:"_class"`
}

func FromModel(rec *models.Subclass) *Subclass {
	b := &Subclass{
		ID:          rec.ID,
		Name:        rec.Name,
		ClassID:     rec.ClassID,
		SchoolOrder: rec.SchoolOrder,
	}

	if rec.R != nil {
		if rec.R.Class != nil {
			b.Class = class.FromModel(rec.R.Class)
		}
	}

	return b
}

// Response represents a workflow that is returned for display.
type Response struct {
	ID          string `json:"id" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name        string `json:"name" example:"Rocket Launch"`
	ClassID     string    `json:"class_id"`
	SchoolOrder int    `json:"school_order"`
	Class       string `json:"_class"`
}

// Response transforms Subclass to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Subclass) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:          m.ID,
		Name:        m.Name,
		ClassID:     m.ClassID,
		SchoolOrder: m.SchoolOrder,
	}

	if m.Class != nil {
		r.Class = m.Class.Name
	}

	return r
}

// Classes a list of Classes.
type Classes []*Subclass

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
	ClassID     string `json:"class_id" validate:"required"`
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
	ClassID     *string `json:"class_id" validate:"required"`
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
