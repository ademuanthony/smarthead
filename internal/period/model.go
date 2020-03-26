package period

import (
	"context"
	"fmt"
	"time"

	"remoteschool/smarthead/internal/postgres/models"

	"github.com/jmoiron/sqlx"
)

// Repository defines the required dependencies for Period.
type Repository struct {
	DbConn *sqlx.DB
}

// NewRepository creates a new Repository that defines dependencies for Period.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DbConn: db,
	}
}

// Period represents a workflow.
type Period struct {
	ID         string     `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StartHour   int       `boil:"start_hour" json:"start_hour" toml:"start_hour" yaml:"start_hour"`
	StartMinute int       `boil:"start_minute" json:"start_minute" toml:"start_minute" yaml:"start_minute"`
	EndHour     int       `boil:"end_hour" json:"end_hour" toml:"end_hour" yaml:"end_hour"`
	EndMinute   int       `boil:"end_minute" json:"end_minute" toml:"end_minute" yaml:"end_minute"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}

func (p Period) String() string {
	return fmt.Sprintf("%d:%02d - %d:%02d", p.StartHour, p.StartMinute, p.EndHour, p.EndMinute)
}

func FromModel(rec *models.Period) *Period {
	b := &Period{
		ID:         rec.ID,
		StartHour: rec.StartHour,
		StartMinute: rec.StartMinute,
		EndHour: rec.EndHour,
		EndMinute: rec.EndMinute,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}

	return b
}

// Response represents a workflow that is returned for display.
type Response struct {
	ID         string            `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StartHour   int       `boil:"start_hour" json:"start_hour" toml:"start_hour" yaml:"start_hour"`
	StartMinute int       `boil:"start_minute" json:"start_minute" toml:"start_minute" yaml:"start_minute"`
	EndHour     int       `boil:"end_hour" json:"end_hour" toml:"end_hour" yaml:"end_hour"`
	EndMinute   int       `boil:"end_minute" json:"end_minute" toml:"end_minute" yaml:"end_minute"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}

func (p Response) String() string {
	return fmt.Sprintf("%d:%02d - %d:%02d", p.StartHour, p.StartMinute, p.EndHour, p.EndMinute)
}

// Response transforms Branch to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Period) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:        m.ID,
		StartHour: m.StartHour,
		StartMinute: m.StartMinute,
		EndHour:  m.EndHour,
		EndMinute: m.EndMinute,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return r
}

// Periods a list of Periods.
type Periods []*Period

// Response transforms a list of Branches to a list of Responses.
func (m *Periods) Response(ctx context.Context) []*Response {
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
	StartHour   int       `json:"start_hour" toml:"start_hour" yaml:"start_hour"`
	StartMinute int       `json:"start_minute" toml:"start_minute" yaml:"start_minute"`
	EndHour     int       `json:"end_hour" toml:"end_hour" yaml:"end_hour"`
	EndMinute   int       `json:"end_minute" toml:"end_minute" yaml:"end_minute"`
}

// ReadRequest defines the information needed to read a period.
type ReadRequest struct {
	ID              string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	IncludeArchived bool   `json:"include-archived" example:"false"`
}

// UpdateRequest defines what information may be provided to modify an existing
// Branch. All fields are optional so clients can send just the fields they want
// changed. It uses pointer fields so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank.
type UpdateRequest struct {
	ID     		string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StartHour   *int   `json:"start_hour" toml:"start_hour" yaml:"start_hour"`
	StartMinute *int   `json:"start_minute" toml:"start_minute" yaml:"start_minute"`
	EndHour     *int   `json:"end_hour" toml:"end_hour" yaml:"end_hour"`
	EndMinute   *int   `json:"end_minute" toml:"end_minute" yaml:"end_minute"`
}

// DeleteRequest defines the information needed to delete a period.
type DeleteRequest struct {
	ID string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
}

// FindRequest defines the possible options to search for branches. By default
// archived period will be excluded from response.
type FindRequest struct {
	Where           string        `json:"where" example:"name = ? and status = ?"`
	Args            []interface{} `json:"args" swaggertype:"array,string" example:"Moon Launch,active"`
	Order           []string      `json:"order" example:"created_at desc"`
	Limit           *uint         `json:"limit" example:"10"`
	Offset          *uint         `json:"offset" example:"20"`
	IncludeArchived bool          `json:"include-archived" example:"false"`
}
