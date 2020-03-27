package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"remoteschool/smarthead/internal/deposit"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/datatable"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/student"

	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

// Deposits represents the Deposits API method handler set.
type Deposits struct {
	Repo        *deposit.Repository
	StudentRepo *student.Repository
	Redis       *redis.Client
	Renderer    web.Renderer
}

func urlDepositsIndex() string {
	return fmt.Sprintf("/admin/deposits")
}

func urlDepositsView(subjectID string) string {
	return fmt.Sprintf("/admin/deposits/%s", subjectID)
}

func urlDepositsUpdate(subjectID string) string {
	return fmt.Sprintf("/admin/deposits/%s/update", subjectID)
}

// Index handles listing all the deposits.
func (h *Deposits) Index(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []datatable.DisplayField{
		{Field: "id", Title: "ID", Visible: false, Searchable: true, Orderable: true, Filterable: false},
		{Field: "student_id", Title: "Student", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Name"},
		{Field: "amount", Title: "Amount", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Amount"},
		{Field: "status", Title: "Status", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Status"},
		{Field: "created_at", Title: "Creation Date", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Email"},
	}

	mapFunc := func(q *deposit.Response, cols []datatable.DisplayField) (resp []datatable.ColumnValue, err error) {
		for i := 0; i < len(cols); i++ {
			col := cols[i]
			var v datatable.ColumnValue
			switch col.Field {
			case "id":
				v.Value = fmt.Sprintf("%s", q.ID)
			case "student_id":
				v.Value = q.StudentID
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlStudentsView(q.StudentID), q.Student)
			case "amount":
				v.Value = fmt.Sprintf("%d", q.Amount/100)
				v.Formatted = v.Value
			case "status":
				v.Value = q.Status
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
		res, err := h.Repo.Find(ctx, claims, deposit.FindRequest{
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
		"urlDepositsIndex": urlDepositsIndex(),
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-deposits-index.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// ApiCreate initialize a new deposit
func (h *Deposits) Initiate(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	
	ctxValues, err := webcontext.ContextValues(ctx)
	if err != nil {
		return web.RespondJsonError(ctx, w, err)
	}
	
	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return web.RespondJsonError(ctx, w, err)
	}

	currentStudent, err := h.StudentRepo.CurrentStudent(ctx, claims)
	if err != nil {
		return web.RespondJsonError(ctx, w, err)
	}

	depositReq := deposit.CreateRequest{
		Channel: "Paystack",
		Amount: 1200000, 
		Status: deposit.StatusPending,
		StudentID: currentStudent.ID,
	}
	depo, err := h.Repo.Create(ctx, claims, depositReq, ctxValues.Now)
	if err != nil {
		return web.RespondJsonError(ctx, w, err)
	}

	depo.Student = currentStudent

	return web.RespondJson(ctx, w, depo, http.StatusOK)
}

// UpdateStatus updates the status of the payment with the specified ID
func (h *Deposits) UpdateStatus(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	ctxValues, err := webcontext.ContextValues(ctx)
	if err != nil {
		return web.RespondJsonError(ctx, w, err)
	}

	depositID := params["deposit_id"]
	err = h.Repo.UpdateStatus(ctx, depositID, ctxValues.Now)
	if err != nil {
		return web.RespondJsonError(ctx, w, err)
	}
	return web.RespondJson(ctx, w, "", http.StatusOK)
}

// View handles displaying a deposits.
func (h *Deposits) View(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

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
	data["urlDepositsIndex"] = urlDepositsIndex()
	data["urlDepositsView"] = urlDepositsView(studentID)

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-deposits-view.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}
