package lesson

import (
	"context"
	"fmt"

	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/period"
	"remoteschool/smarthead/internal/postgres/models"
	"remoteschool/smarthead/internal/user"

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
type Lesson struct {
	ID        string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name      string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ClassID   string `boil:"class_id" json:"class_id" toml:"class_id" yaml:"class_id"`
	TeacherID string `boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	PeriodID  string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	Date      int64  `boil:"date" json:"date" toml:"date" yaml:"date"`
	Code      string `boil:"code" json:"code" toml:"code" yaml:"code"`

	Students []string       `json:"students"`
	Class    *class.Class   `json:"_class"`
	Teacher  *models.User   `json:"teacher"`
	Period   *period.Period `json:"period"`
}

func FromModel(rec *models.Lesson) *Lesson {
	b := &Lesson{
		ID:        rec.ID,
		Name:      rec.Name,
		ClassID:   rec.ClassID,
		TeacherID: rec.TeacherID,
		PeriodID:  rec.PeriodID,
		Date:      rec.Date,
		Code:      rec.Code,
	}

	if rec.R == nil {
		if rec.R.Class != nil {
			b.Class = class.FromModel(rec.R.Class)
		}
		if rec.R.Period != nil {
			b.Period = period.FromModel(rec.R.Period)
		}
		if rec.R.Teacher != nil {
			b.Teacher = rec.R.Teacher
		}
	}

	return b
}

// Response represents a workflow that is returned for display.
type Response struct {
	ID        string `json:"id" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	Name      string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ClassID   string `boil:"class_id" json:"class_id" toml:"class_id" yaml:"class_id"`
	TeacherID string `boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	PeriodID  string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	Date      int64  `boil:"date" json:"date" toml:"date" yaml:"date"`
	Code      string `boil:"code" json:"code" toml:"code" yaml:"code"`

	Class    string   `json:"_class"`
	Teacher  string   `json:"teacher"`
	Period   string   `json:"period"`
	Students []string `json:"students"`
}

// Response transforms Subject to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Lesson) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:        m.ID,
		Name:      m.Name,
		ClassID:   m.ClassID,
		TeacherID: m.TeacherID,
		PeriodID:  m.PeriodID,
		Date:      m.Date,
		Code:      m.Code,
		Students:  m.Students,
	}

	if m.Class != nil {
		r.Class = m.Class.Name
	}
	if m.Teacher != nil {
		r.Teacher = fmt.Sprintf("%s %s", m.Teacher.LastName, m.Teacher.FirstName)
	}
	if m.Period != nil {
		r.Period = m.Period.String()
	}

	return r
}

// Lessons a list of Lessons.
type Lessons []*Lesson

// Response transforms a list of Branches to a list of Responses.
func (m *Lessons) Response(ctx context.Context) []*Response {
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
	Name      string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ClassID   string `boil:"class_id" json:"class_id" toml:"class_id" yaml:"class_id"`
	TeacherID string `boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	PeriodID  string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	Date      int64  `boil:"date" json:"date" toml:"date" yaml:"date"`
	Code      string `boil:"code" json:"code" toml:"code" yaml:"code"`

	Students []string       `json:"students"`
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
