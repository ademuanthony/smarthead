package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/datatable"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/student"
	"remoteschool/smarthead/internal/subscription"

	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

// Deposits represents the Deposits API method handler set.
type Subscriptions struct {
	Repo        *subscription.Repository
	StudentRepo *student.Repository
	Redis       *redis.Client
	Renderer    web.Renderer
}

func urlSubscriptionsIndex() string {
	return fmt.Sprintf("/admin/subscriptions")
}

func urlSubscriptionsView(subjectID string) string {
	return fmt.Sprintf("/admin/subscriptions/%s", subjectID)
}

func urlSubscriptionsUpdate(subjectID string) string {
	return fmt.Sprintf("/admin/subscriptions/%s/update", subjectID)
}

// Index handles listing all the subscriptions.
func (h *Subscriptions) Index(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}
  
	fields := []datatable.DisplayField{
		{Field: "id", Title: "", Visible: true, Searchable: true, },
		{Field: "student_id", Title: "Student", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Name"},
		{Field: "subject_id", Title: "Subject", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Username"},
		{Field: "period_id", Title: "Period", Visible: true, Searchable: true, Orderable: true, Filterable: true},
		{Field: "class_id", Title: "Class", Visible: true, Searchable: true, Orderable: true, Filterable: true},
		{Field: "start_date", Title: "Start Date", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Class"},
		{Field: "end_date", Title: "End Date", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Phone"},
		{Field: "created_at", Title: "Creation Date", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Email"},
	}

	mapFunc := func(q *subscription.Response, cols []datatable.DisplayField) (resp []datatable.ColumnValue, err error) {
		for i := 0; i < len(cols); i++ {
			col := cols[i]
			var v datatable.ColumnValue
			switch col.Field {
			case "id":
				v.Value = q.ID
				v.Formatted = fmt.Sprintf("<input type='checkbox' value='%s' class='form-control' data-target='subscription.selected'/>", q.ID)
			case "student_id":
				v.Value = q.StudentID
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlStudentsView(q.StudentID), q.Student)
			case "subject_id":
				v.Value = q.SubjectID
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlSubjectsView(q.SubjectID), q.Subject)
			case "period_id":
				v.Value = q.PeriodID
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlSubjectsView(q.PeriodID), q.Period)
			case "class_id":
				v.Value = q.ClassID
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlClassesView(q.ClassID), q.Class)
			case "start_date":
				v.Value = q.StartDate.LocalDate
				v.Formatted = v.Value
			case "end_date":
				v.Value = q.EndDate.LocalDate
				v.Formatted = v.Value
			case "created_at":
				v.Value = q.CreatedAt.LocalDate
				v.Formatted = v.Value
			default:
				return resp, errors.Errorf("Failed to map value for %s.", col.Field)
			}
			resp = append(resp, v)
		}

		return resp, nil
	}

	loadFunc := func(ctx context.Context, sorting string, fields []datatable.DisplayField) (resp [][]datatable.ColumnValue, err error) {
		if sorting == "" {
			sorting = "created_at desc"
		}
		res, err := h.Repo.Find(ctx, claims, subscription.FindRequest{
			Order: strings.Split(sorting, ","),
		})
		if err != nil {
			return resp, err
		}

		for _, a := range res {
			l, err := mapFunc(a.Response(ctx), fields)
			if err != nil {
				return resp, errors.Wrapf(err, "Failed to map student for display.")
			}

			resp = append(resp, l)
		}

		return resp, nil
	}

	dt, err := datatable.New(ctx, w, r, h.Redis, fields, loadFunc)
	if err != nil {
		return err
	}

	if dt.HasCache() {
		return nil
	}

	if ok, err := dt.Render(); ok {
		if err != nil {
			return err
		}
		return nil
	}

	data := map[string]interface{}{
		"datatable":             dt.Response(),
		"urlSubscriptionsIndex": urlSubscriptionsIndex(),
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subscriptions-index.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// My handles the listing all the subscriptions for the current account.
func (h *Subscriptions) My(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []datatable.DisplayField{
		{Field: "id", Title: "ID", Visible: false, Searchable: true, Orderable: true, Filterable: false},
		{Field: "student_id", Title: "Student", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Name"},
		{Field: "subject_id", Title: "Subject", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Username"},
		{Field: "period_id", Title: "Period", Visible: true, Searchable: true, Orderable: true},
		{Field: "start_date", Title: "Start Date", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Class"},
		{Field: "end_date", Title: "End Date", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Phone"},
		{Field: "created_at", Title: "Creation Date", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Email"},
	}

	mapFunc := func(q *subscription.Response, cols []datatable.DisplayField) (resp []datatable.ColumnValue, err error) {
		for i := 0; i < len(cols); i++ {
			col := cols[i]
			var v datatable.ColumnValue
			switch col.Field {
			case "id":
				v.Value = fmt.Sprintf("%s", q.ID)
			case "student_id":
				v.Value = q.StudentID
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlStudentsView(q.StudentID), q.Student)
			case "subject_id":
				v.Value = q.SubjectID
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlSubjectsView(q.SubjectID), q.Subject)
			case "period_id":
				v.Value = q.PeriodID
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlSubjectsView(q.PeriodID), q.Period)
			case "start_date":
				v.Value = q.StartDate.LocalDate
				v.Formatted = v.Value
			case "end_date":
				v.Value = q.EndDate.LocalDate
				v.Formatted = v.Value
			case "created_at":
				v.Value = q.CreatedAt.LocalDate
				v.Formatted = v.Value
			default:
				return resp, errors.Errorf("Failed to map value for %s.", col.Field)
			}
			resp = append(resp, v)
		}

		return resp, nil
	}

	currentStudent, err := h.StudentRepo.CurrentStudent(ctx, claims)
	if err != nil {
		return fmt.Errorf("you must be a student in order to access this page, %s", err.Error())
	}

	loadFunc := func(ctx context.Context, sorting string, fields []datatable.DisplayField) (resp [][]datatable.ColumnValue, err error) {
		res, err := h.Repo.Find(ctx, claims, subscription.FindRequest{
			Where: "student_id = ?",
			Order: strings.Split(sorting, ","),
			Args:  []interface{}{currentStudent.ID},
		})
		if err != nil {
			return resp, err
		}

		for _, a := range res {
			l, err := mapFunc(a.Response(ctx), fields)
			if err != nil {
				return resp, errors.Wrapf(err, "Failed to map student for display.")
			}

			resp = append(resp, l)
		}

		return resp, nil
	}

	dt, err := datatable.New(ctx, w, r, h.Redis, fields, loadFunc)
	if err != nil {
		return err
	}

	if dt.HasCache() {
		return nil
	}

	if ok, err := dt.Render(); ok {
		if err != nil {
			return err
		}
		return nil
	}

	data := map[string]interface{}{
		"datatable":             dt.Response(),
		"urlSubscriptionsIndex": urlSubscriptionsIndex(),
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subscriptions-index.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// View handles displaying a subscriptions.
func (h *Subscriptions) View(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	studentID := params["subscription_id"]

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})

	sub, err := h.Repo.ReadByID(ctx, claims, studentID)
	if err != nil {
		return err
	}
	data["student"] = sub.Response(ctx)
	data["urlSubscriptionsIndex"] = urlSubscriptionsIndex()
	data["urlSubscriptionsView"] = urlSubscriptionsView(studentID)

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subscriptions-view.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}
