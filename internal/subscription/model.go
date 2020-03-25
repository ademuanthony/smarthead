package branch

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"remoteschool/smarthead/internal/platform/web"
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
type Subscription struct {
	ID         string  `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StudentID  string `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	SubjectID  string `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	PeriodID   string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	DaysOfWeek int    `boil:"days_of_week" json:"days_of_week" toml:"days_of_week" yaml:"days_of_week"`
	StartDate  int64  `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	EndDate    int64  `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`
	CreatedAt  int64  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	
	
}

func FromModel(rec *models.Subscription) *Subscription {
	b := &Subscription{
		ID:         rec.ID,
		Name:       rec.Name,
		CreatedAt:  time.Unix(rec.CreatedAt, 0),
		UpdatedAt:  time.Unix(rec.UpdatedAt, 0),
	}
	if rec.ArchivedAt.Valid {
		archivedAt := time.Unix(rec.ArchivedAt.Int64, 0)
		b.ArchivedAt = &archivedAt
	}

	return b
}

// Response represents a workflow that is returned for display.
type Response struct {
	ID         string            `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name       string            `json:"name"  validate:"required" example:"Rocket Launch"`
	CreatedAt  web.TimeResponse  `json:"created_at"`            // CreatedAt contains multiple format options for display.
	UpdatedAt  web.TimeResponse  `json:"updated_at"`            // UpdatedAt contains multiple format options for display.
	ArchivedAt *web.TimeResponse `json:"archived_at,omitempty"` // ArchivedAt contains multiple format options for display.
}

// Response transforms Branch to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Subscription) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: web.NewTimeResponse(ctx, m.CreatedAt),
		UpdatedAt: web.NewTimeResponse(ctx, m.UpdatedAt),
	}

	if m.ArchivedAt != nil && !m.ArchivedAt.IsZero() {
		at := web.NewTimeResponse(ctx, *m.ArchivedAt)
		r.ArchivedAt = &at
	}

	return r
}

// Branches a list of Branches.
type Branches []*Subscription

// Response transforms a list of Branches to a list of Responses.
func (m *Branches) Response(ctx context.Context) []*Response {
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
	Name      string           `json:"name" validate:"required"  example:"Rocket Launch"`
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
	ID     string           `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name   *string          `json:"name,omitempty" validate:"omitempty,unique" example:"Rocket Launch to Moon"`
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
