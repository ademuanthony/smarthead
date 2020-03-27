package deposit

import (
	"context"
	"time"

	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/postgres/models"
	"remoteschool/smarthead/internal/student"

	"github.com/jmoiron/sqlx"
)

// Repository defines the required dependencies for Branch.
type Repository struct {
	DbConn *sqlx.DB
}

const (
	StatusPending = "pending"
	StatusPaid = "paid"
	StatusSubscribed = "subscribed"
)

// NewRepository creates a new Repository that defines dependencies for Branch.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DbConn: db,
	}
}

// Branch represents a workflow.
type Deposit struct {
	ID        string    `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StudentID string    `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	Amount    int       `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Ref       string    `boil:"ref" json:"ref" toml:"ref" yaml:"ref"`
	Status    string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	Channel   string    `boil:"channel" json:"channel" toml:"channel" yaml:"channel"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	Student *student.Student `json:"student"`
}

func FromModel(rec *models.Deposit) *Deposit {
	b := &Deposit{
		ID:        rec.ID,
		StudentID: rec.StudentID,
		Amount:    rec.Amount,
		Ref:       rec.Ref,
		Status:    rec.Status,
		Channel:   rec.Channel,
		CreatedAt: rec.CreatedAt,
	}
	if rec.R != nil {
		if rec.R.Student != nil {
			b.Student = student.FromModel(rec.R.Student)
		}
	}
	return b
}

// Response represents a workflow that is returned for display.
type Response struct {
	ID        string           `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StudentID string           `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	Amount    int              `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Ref       string           `boil:"ref" json:"ref" toml:"ref" yaml:"ref"`
	Status    string           `boil:"status" json:"status" toml:"status" yaml:"status"`
	Channel   string           `boil:"channel" json:"channel" toml:"channel" yaml:"channel"`
	CreatedAt web.TimeResponse `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	Student string `json:"student"`
}

// Response transforms Branch to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Deposit) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:        m.ID,
		CreatedAt: web.NewTimeResponse(ctx, m.CreatedAt),
		StudentID: m.StudentID,
		Amount:    m.Amount,
		Ref:       m.Ref,
		Status:    m.Status,
		Channel:   m.Channel,
	}

	if m.Student != nil {
		r.Student = m.Student.Name
	}

	return r
}

// Deposits a list of Deposits.
type Deposits []*Deposit

// Response transforms a list of Deposits to a list of Responses.
func (m *Deposits) Response(ctx context.Context) []*Response {
	var l []*Response
	if m != nil && len(*m) > 0 {
		for _, n := range *m {
			l = append(l, n.Response(ctx))
		}
	}

	return l
}

// CreateRequest contains information needed to create a new Deposit.
type CreateRequest struct {
	StudentID string `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	Amount    int    `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Ref       string `boil:"ref" json:"ref" toml:"ref" yaml:"ref"`
	Status    string `boil:"status" json:"status" toml:"status" yaml:"status"`
	Channel   string `boil:"channel" json:"channel" toml:"channel" yaml:"channel"`
}

// UpdateRequest contains information needed to update a Deposit.
type UpdateRequest struct {
	ID     string `boil:"id" json:"id" validate:"required" yaml:"id"`
	Ref    *string `boil:"ref" json:"ref" toml:"ref" yaml:"ref"`
	Status *string `boil:"status" json:"status" toml:"status" yaml:"status"`
}

// ReadRequest defines the information needed to read a checklist.
type ReadRequest struct {
	ID              string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	IncludeArchived bool   `json:"include-archived" example:"false"`
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
