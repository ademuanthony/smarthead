package subscription

import (
	"context"
	"time"

	"remoteschool/smarthead/internal/period"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/postgres/models"
	"remoteschool/smarthead/internal/student"
	"remoteschool/smarthead/internal/subject"

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
type Subscription struct {
	ID         string `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StudentID  string `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	SubjectID  string `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	PeriodID   string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	DaysOfWeek int    `boil:"days_of_week" json:"days_of_week" toml:"days_of_week" yaml:"days_of_week"`
	StartDate  int64  `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	EndDate    int64  `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`
	CreatedAt  int64  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Status	   string `json:"status"`

	Subject *subject.Subject `json:"subject"`
	Student *student.Student `json:"student"`
	Period  *period.Period   `json:"period"`
}

func FromModel(rec *models.Subscription) *Subscription {
	b := &Subscription{
		ID:         rec.ID,
		CreatedAt:  rec.CreatedAt,
		DaysOfWeek: rec.DaysOfWeek,
		EndDate:    rec.EndDate,
		PeriodID:   rec.PeriodID,
		StartDate:  rec.StartDate,
		StudentID:  rec.StudentID,
		SubjectID:  rec.SubjectID,
	}
	if rec.R != nil {
		if rec.R.Period != nil {
			b.Period = period.FromModel(rec.R.Period)
		}

		if rec.R.Subject != nil {
			b.Subject = subject.FromModel(rec.R.Subject)
		}

		if rec.R.Student != nil {
			b.Student = student.FromModel(rec.R.Student)
		}
	}
	return b
}

// Response represents a workflow that is returned for display.
type Response struct {
	ID         string           `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StudentID  string           `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	SubjectID  string           `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	PeriodID   string           `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	DaysOfWeek int              `boil:"days_of_week" json:"days_of_week" toml:"days_of_week" yaml:"days_of_week"`
	StartDate  web.TimeResponse `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	EndDate    web.TimeResponse `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`
	CreatedAt  web.TimeResponse `json:"created_at"` // CreatedAt contains multiple format options for display.
	Status	   string 			`json:"status"`

	Subject string `json:"subject"`
	Student string `json:"student"`
	Period  string `json:"period"`
}

// Response transforms Branch to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Subscription) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:         m.ID,
		CreatedAt:  web.NewTimeResponse(ctx, time.Unix(m.CreatedAt, 0)),
		DaysOfWeek: m.DaysOfWeek,
		EndDate:    web.NewTimeResponse(ctx, time.Unix(m.EndDate, 0)),
		PeriodID:   m.PeriodID,
		StartDate:  web.NewTimeResponse(ctx, time.Unix(m.StartDate, 0)),
		StudentID:  m.StudentID,
		SubjectID:  m.SubjectID,
	}

	if m.Student != nil {
		r.Student = m.Student.Name
	}

	if m.Subject != nil {
		r.Subject = m.Subject.Name
	}

	if m.Period != nil {
		r.Period = m.Period.String()
	}

	return r
}

// Subscriptions a list of Subscriptions.
type Subscriptions []*Subscription

// Response transforms a list of Subscriptions to a list of Responses.
func (m *Subscriptions) Response(ctx context.Context) []*Response {
	var l []*Response
	if m != nil && len(*m) > 0 {
		for _, n := range *m {
			l = append(l, n.Response(ctx))
		}
	}

	return l
}

// CreateRequest contains information needed to create a new Subscription.
type CreateRequest struct {
	StudentID  string `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	SubjectID  string `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	PeriodID   string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	DaysOfWeek int    `boil:"days_of_week" json:"days_of_week" toml:"days_of_week" yaml:"days_of_week"`
	StartDate  int64  `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	EndDate    int64  `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`
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
