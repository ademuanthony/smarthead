package timetable

import (
	"context"
	"time"

	"remoteschool/smarthead/internal/period"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/postgres/models"
	"remoteschool/smarthead/internal/subclass"
	"remoteschool/smarthead/internal/subject"

	"github.com/jmoiron/sqlx"
)

// Repository defines the required dependencies for Timtable.
type Repository struct {
	DbConn *sqlx.DB
}

// NewRepository creates a new Repository that defines dependencies for Timtable.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DbConn: db,
	}
}

// Timtable represents a timetable .
type Timetable struct {
	ID         string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	SubclassID string `boil:"subclass_id" json:"subclass_id" toml:"subclass_id" yaml:"subclass_id"`
	SubjectID  string `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	TeacherID  string `boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	PeriodID   string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	Day        int    `boil:"day" json:"day" toml:"day" yaml:"day"`

	Subclass *subclass.Subclass
	Subject  *subject.Subject
	Teacher  *models.User
	Period   *period.Period
}


func (m *Timetable) NextLessonDate(ctx context.Context, now time.Time) web.TimeResponse {
	daysLeft := m.Day - now.Day()
	if daysLeft < 0 {
		daysLeft = (7 - now.Day()) + m.Day
	}
	startDate := time.Date(now.Year(), now.Month(), daysLeft, 
				m.Period.StartHour - 1, m.Period.StartMinute, 0, 0, time.UTC)
	return web.NewTimeResponse(ctx, startDate)
}

func FromModel(rec *models.Timetable) *Timetable {
	b := &Timetable{
		ID: rec.ID,
		Day: rec.Day,
		PeriodID: rec.PeriodID,
		SubclassID: rec.SubclassID,
		SubjectID: rec.SubjectID,
		TeacherID: rec.TeacherID,
	}

	if rec.R != nil {
		if rec.R.Period != nil {
			b.Period = period.FromModel(rec.R.Period)
		}
		if rec.R.Subclass != nil {
			b.Subclass = subclass.FromModel(rec.R.Subclass)
		}
		if rec.R.Subject != nil {
			b.Subject = subject.FromModel(rec.R.Subject)
		}
		if rec.R.Teacher != nil {
			b.Teacher = rec.R.Teacher
		}
	}

	return b
}

// Response represents a workflow that is returned for display.
type Response struct {
	ID         string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	SubclassID string `boil:"subclass_id" json:"subclass_id" toml:"subclass_id" yaml:"subclass_id"`
	SubjectID  string `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	TeacherID  string `boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	PeriodID   string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	Day        time.Weekday    `boil:"day" json:"day" toml:"day" yaml:"day"`

	Subclass string `json:"subclass"`
	Subject  string `json:"subject"`
	Teacher  string `json:"teacher"`
	Period   string `json:"period"`
}

// Response transforms Subject to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Timetable) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:         m.ID,
		SubjectID:  m.SubjectID,
		SubclassID: m.SubclassID,
		TeacherID:  m.TeacherID,
		PeriodID:   m.PeriodID,
		Day:        time.Weekday(m.Day),
	}

	if m.Subclass != nil {
		r.Subclass = m.Subclass.Name
	}
	if m.Subject != nil {
		r.Subject = m.Subject.Name
	}
	if m.Teacher != nil {
		r.Teacher = m.Teacher.FirstName + " " + m.Teacher.LastName
	}
	if m.Period != nil {
		r.Period = m.Period.String()
	}

	return r
}

// Timetables a list of Timetables.
type Timetables []*Timetable

// Response transforms a list of Branches to a list of Responses.
func (m *Timetables) Response(ctx context.Context) []*Response {
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
	SubclassID string `validate:"required" boil:"subclass_id" json:"subclass_id" toml:"subclass_id" yaml:"subclass_id"`
	SubjectID  string `validate:"required" boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	TeacherID  string `validate:"required" boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	PeriodID   string `validate:"required" boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	Day        int    `validate:"required" boil:"day" json:"day" toml:"day" yaml:"day"`
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
	ID         string  `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	SubclassID *string `validate:"required" boil:"subclass_id" json:"subclass_id" toml:"subclass_id" yaml:"subclass_id"`
	SubjectID  *string `validate:"required" boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	TeacherID  *string `validate:"required" boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	PeriodID   *string `validate:"required" boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	Day        *int    `validate:"required" boil:"day" json:"day" toml:"day" yaml:"day"`
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
	IncludeSubclass bool          `json:"include-subclass" example:"false"`
	IncludeSubject  bool          `json:"include-subject" example:"false"`
	IncludePeriod   bool          `json:"include-period" example:"false"`
	IncludeTeacher  bool          `json:"include-teacher" example:"false"`
}
