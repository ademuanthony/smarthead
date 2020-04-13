package deposit

import (
	"context"
	"time"

	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/postgres/models"
	"remoteschool/smarthead/internal/student"
	"remoteschool/smarthead/internal/subscription"

	"github.com/jmoiron/sqlx"
)

// Repository defines the required dependencies for Branch.
type Repository struct {
	DbConn           *sqlx.DB
	SubscriptionRepo *subscription.Repository
	PaystackSecret   string
}

const (
	StatusPending    = "pending"
	StatusPaid       = "paid"
	StatusSubscribed = "subscribed"
)

// NewRepository creates a new Repository that defines dependencies for Branch.
func NewRepository(db *sqlx.DB, subscriptionRepo *subscription.Repository, paystackSecret string) *Repository {
	return &Repository{
		DbConn:           db,
		SubscriptionRepo: subscriptionRepo,
		PaystackSecret:   paystackSecret,
	}
}

// Branch represents a workflow.
type Deposit struct {
	ID         string    `json:"id" validate:"required,uuid" example:"985f1746-1d9f-459f-a2d9-fc53ece5ae86"`
	StudentID  string    `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	SubjectID  string    `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	PeriodID   string    `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	ClassID    string    `json:"class_id"`
	DaysOfWeek int       `boil:"days_of_week" json:"days_of_week" toml:"days_of_week" yaml:"days_of_week"`
	Amount     int       `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	PaymentRef string 	 `boil:"payment_ref" json:"payment_ref" toml:"payment_ref" yaml:"payment_ref"`
	Ref        string    `boil:"ref" json:"ref" toml:"ref" yaml:"ref"`
	Status     string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	Channel    string    `boil:"channel" json:"channel" toml:"channel" yaml:"channel"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	Student *student.Student `json:"student"`
	Class   *class.Class     `json:"_class"`
}

func FromModel(rec *models.Deposit) *Deposit {
	b := &Deposit{
		ID:         rec.ID,
		StudentID:  rec.StudentID,
		SubjectID:  rec.SubjectID,
		PeriodID:   rec.PeriodID.String,
		ClassID:    rec.ClassID,
		DaysOfWeek: rec.DaysOfWeek,
		Amount:     rec.Amount,
		Ref:        rec.Ref,
		Status:     rec.Status,
		Channel:    rec.Channel,
		CreatedAt:  rec.CreatedAt,
	}
	if rec.R != nil {
		if rec.R.Student != nil {
			b.Student = student.FromModel(rec.R.Student)
		}

		if rec.R.Class != nil {
			b.Class = class.FromModel(rec.R.Class)
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
	ClassID    string           `json:"class_id"`
	DaysOfWeek int              `boil:"days_of_week" json:"days_of_week" toml:"days_of_week" yaml:"days_of_week"`
	Amount     int              `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	PaymentRef string 	 		`boil:"payment_ref" json:"payment_ref" toml:"payment_ref" yaml:"payment_ref"`
	Ref        string           `boil:"ref" json:"ref" toml:"ref" yaml:"ref"`
	Status     string           `boil:"status" json:"status" toml:"status" yaml:"status"`
	Channel    string           `boil:"channel" json:"channel" toml:"channel" yaml:"channel"`
	CreatedAt  web.TimeResponse `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	Student string `json:"student"`
	Class   string `json:"_class"`
}

// Response transforms Branch to the Response that is used for display.
// Additional filtering by context values or translations could be applied.
func (m *Deposit) Response(ctx context.Context) *Response {
	if m == nil {
		return nil
	}

	r := &Response{
		ID:         m.ID,
		CreatedAt:  web.NewTimeResponse(ctx, m.CreatedAt),
		StudentID:  m.StudentID,
		SubjectID:  m.SubjectID,
		PeriodID:   m.PeriodID,
		DaysOfWeek: m.DaysOfWeek,
		Amount:     m.Amount,
		Ref:        m.Ref,
		Status:     m.Status,
		Channel:    m.Channel,
	}

	if m.Student != nil {
		r.Student = m.Student.Name
	}

	if m.Class != nil {
		r.Class = m.Class.Name
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
	Count 	   int    `json:"count"`
	StudentID  string `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	SubjectID  string `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	PeriodID   string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	ClassID    string `json:"class_id"`
	DaysOfWeek int    `boil:"days_of_week" json:"days_of_week" toml:"days_of_week" yaml:"days_of_week"`
	Amount     int    `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Ref        string `boil:"ref" json:"ref" toml:"ref" yaml:"ref"`
	PaymentRef string `boil:"payment_ref" json:"payment_ref" toml:"payment_ref" yaml:"payment_ref"`
	Status     string `boil:"status" json:"status" toml:"status" yaml:"status"`
	Channel    string `boil:"channel" json:"channel" toml:"channel" yaml:"channel"`
}

// UpdateRequest contains information needed to update a Deposit.
type UpdateRequest struct {
	ID     string  `boil:"id" json:"id" validate:"required" yaml:"id"`
	Ref    *string `boil:"ref" json:"ref" toml:"ref" yaml:"ref"`
	Status *string `boil:"status" json:"status" toml:"status" yaml:"status"`
}

type SubscriptionItem struct {
	SubjectID  string `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	PeriodID   string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	ClassID    string `json:"class_id"`
}

type UpdateStatusRequest struct {
	ID     string  			  `json:"id" validate:"required" yaml:"id"`
	Items  []SubscriptionItem `json:"items"`
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
