package student

import (
	"context"
	"time"

	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/platform/web"
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

// Branch represents a workflow.
type Student struct {
	ID             string    `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name           string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Username       string    `boil:"username" json:"username" toml:"username" yaml:"username"`
	Age            int       `boil:"age" json:"age" toml:"age" yaml:"age"`
	AccountBalance int       `boil:"account_balance" json:"account_balance" toml:"account_balance" yaml:"account_balance"`
	CurrentClass   string    `boil:"current_class" json:"current_class" toml:"current_class" yaml:"current_class"`
	ClassID        string    `json:"class_id"`
	ParentPhone    string    `boil:"parent_phone" json:"parent_phone" toml:"parent_phone" yaml:"parent_phone"`
	ParentEmail    string    `boil:"parent_email" json:"parent_email" toml:"parent_email" yaml:"parent_email"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	Class *class.Class `json:"_class"`
}

func FromModel(rec *models.Student) *Student {
	b := &Student{
		ID:             rec.ID,
		Name:           rec.Name,
		AccountBalance: rec.AccountBalance,
		Age:            rec.Age,
		CreatedAt:      rec.CreatedAt,
		ClassID:        rec.ClassID.String,
		ParentEmail:    rec.ParentEmail,
		ParentPhone:    rec.ParentPhone,
		UpdatedAt:      rec.UpdatedAt,
		Username:       rec.Username,
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
	ID             string           `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name           string           `boil:"name" json:"name" toml:"name" yaml:"name"`
	Username       string           `boil:"username" json:"username" toml:"username" yaml:"username"`
	Age            int              `boil:"age" json:"age" toml:"age" yaml:"age"`
	AccountBalance int              `boil:"account_balance" json:"account_balance" toml:"account_balance" yaml:"account_balance"`
	ClassID        string           `json:"class_id"`
	CurrentClass   string           `boil:"current_class" json:"current_class" toml:"current_class" yaml:"current_class"`
	ParentPhone    string           `boil:"parent_phone" json:"parent_phone" toml:"parent_phone" yaml:"parent_phone"`
	ParentEmail    string           `boil:"parent_email" json:"parent_email" toml:"parent_email" yaml:"parent_email"`
	CreatedAt      web.TimeResponse `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      web.TimeResponse `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}

// Response transforms Branch to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Student) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:             m.ID,
		Name:           m.Name,
		Username:       m.Username,
		Age:            m.Age,
		AccountBalance: m.AccountBalance,
		ClassID:        m.ClassID,
		ParentEmail:    m.ParentEmail,
		ParentPhone:    m.ParentPhone,
		CreatedAt:      web.NewTimeResponse(ctx, m.CreatedAt),
		UpdatedAt:      web.NewTimeResponse(ctx, m.UpdatedAt),
	}

	if m.Class != nil {
		r.CurrentClass = m.Class.Name
	}

	return r
}

// Students a list of Students.
type Students []*Student

// Response transforms a list of Students to a list of Responses.
func (m *Students) Response(ctx context.Context) []*Response {
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
	Name           string `json:"name" validate:"required"  example:"Rocket Launch"`
	Username       string `json:"username" json:"username" validate:"required" toml:"username" yaml:"username"`
	Age            int    `json:"age" toml:"age" yaml:"age"`
	AccountBalance int    `json:"account_balance" toml:"account_balance" yaml:"account_balance"`
	ClassID        string `json:"class_id" toml:"class_id" yaml:"class_id"`
	ParentPhone    string `json:"parent_phone" validate:"required" toml:"parent_phone" yaml:"parent_phone"`
	ParentEmail    string `json:"parent_email" validate:"required" toml:"parent_email" yaml:"parent_email"`
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
	ID             string  `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name           *string `json:"name"  example:"Rocket Launch"`
	Username       *string `json:"username"`
	Age            *int    `json:"age"`
	AccountBalance *int    `json:"account_balance"`
	ClassID        *string `json:"class_id"`
	ParentPhone    *string `json:"parent_phone"`
	ParentEmail    *string `json:"parent_email"`
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
