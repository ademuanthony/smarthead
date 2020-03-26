package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"remoteschool/smarthead/internal/period"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/datatable"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/platform/web/weberror"
	"remoteschool/smarthead/internal/subject"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

// Periods represents the Periods API method handler set.
type Periods struct {
	Repo 	 *period.Repository
	Redis    *redis.Client
	Renderer web.Renderer
}

func urlPeriodsIndex() string {
	return fmt.Sprintf("/admin/periods")
}

func urlPeriodsCreate() string {
	return fmt.Sprintf("/admin/periods/create")
}

func urlPeriodsView(subjectID string) string {
	return fmt.Sprintf("/admin/periods/%s", subjectID)
}

func urlPeriodsUpdate(subjectID string) string {
	return fmt.Sprintf("/admin/periods/%s/update", subjectID)
}

// Index handles listing all the periods.
func (h *Periods) Index(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	} 

	fields := []datatable.DisplayField{
		{Field: "id", Title: "ID", Visible: false, Searchable: true, Orderable: true, Filterable: false},
		{Field: "period", Title: "Period", Visible: true, Orderable: true, },
		{Field: "start_hour", Title: "Start Hour", Visible: true, Orderable: true, },
		{Field: "start_minute", Title: "Start Minute", Visible: true, Orderable: true, },
		{Field: "end_hour", Title: "End Hour", Visible: true, Orderable: true, },
		{Field: "end_minute", Title: "End Minute", Visible: true, Orderable: true, },
	}

	mapFunc := func(q *period.Period, cols []datatable.DisplayField) (resp []datatable.ColumnValue, err error) {
		for i := 0; i < len(cols); i++ {
			col := cols[i]
			var v datatable.ColumnValue
			switch col.Field {
			case "id":
				v.Value = fmt.Sprintf("%s", q.ID)
			case "period":
				v.Value = q.String()
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlPeriodsView(q.ID), v.Value)
			case "start_hour":
				v.Value = fmt.Sprintf("%d", q.StartHour)
				v.Formatted = v.Value
			case "start_minute":
				v.Value = fmt.Sprintf("%20d", q.StartMinute)
				v.Formatted = v.Value
			case "end_hour":
				v.Value = fmt.Sprintf("%d", q.EndHour)
				v.Formatted = v.Value
			case "end_minute":
				v.Value = fmt.Sprintf("%20d", q.EndMinute)
				v.Formatted = v.Value
			default:
				return resp, errors.Errorf("Failed to map value for %s.", col.Field)
			}
			resp = append(resp, v)
		}

		return resp, nil
	}

	loadFunc := func(ctx context.Context, sorting string, fields []datatable.DisplayField) (resp [][]datatable.ColumnValue, err error) {
		res, err := h.Repo.Find(ctx, claims, period.FindRequest{
			Order: strings.Split(sorting, ","),
		})
		if err != nil {
			return resp, err
		}

		for _, a := range res {
			l, err := mapFunc(a, fields)
			if err != nil {
				return resp, errors.Wrapf(err, "Failed to map checklist for display.")
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
		"datatable":           dt.Response(),
		"urlPeriodsCreate": urlPeriodsCreate(),
		"urlPeriodsIndex": urlPeriodsIndex(),
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-periods-index.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// Create handles creating a new period.
func (h *Periods) Create(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	ctxValues, err := webcontext.ContextValues(ctx)
	if err != nil {
		return err
	}

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	//
	req := new(period.CreateRequest)
	data := make(map[string]interface{})
	f := func() (bool, error) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				return false, err
			}

			decoder := schema.NewDecoder()
			decoder.IgnoreUnknownKeys(true)

			if err := decoder.Decode(req, r.PostForm); err != nil {
				return false, err
			}

			sub, err := h.Repo.Create(ctx, claims, *req, ctxValues.Now)
			if err != nil {
				switch errors.Cause(err) {
				default:
					if verr, ok := weberror.NewValidationError(ctx, err); ok {
						data["validationErrors"] = verr.(*weberror.Error)
						return false, nil
					} else {
						return false, err
					}
				}
			}

			// Display a success message to the checklist.
			webcontext.SessionFlashSuccess(ctx,
				"Period Created",
				"Period successfully created.")

			return true, web.Redirect(ctx, w, r, urlPeriodsView(sub.ID), http.StatusFound)
		}

		return false, nil
	}

	end, err := f()
	if err != nil {
		return web.RenderError(ctx, w, r, err, h.Renderer, TmplLayoutBase, TmplContentErrorGeneric, web.MIMETextHTMLCharsetUTF8)
	} else if end {
		return nil
	}

	data["form"] = req
	data["urlPeriodsIndex"] = urlPeriodsIndex() 

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(period.CreateRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-periods-create.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// View handles displaying a period.
func (h *Periods) View(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	periodID := params["period_id"]

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	f := func() (bool, error) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				return false, err
			}

			switch r.PostForm.Get("action") {
			case "archive":
				err = h.Repo.Delete(ctx, claims, period.DeleteRequest{
					ID: periodID,
				})
				if err != nil {
					return false, err
				}

				webcontext.SessionFlashSuccess(ctx,
					"Period Archive",
					"Period successfully archive.")

				return true, web.Redirect(ctx, w, r, urlSubjectsIndex(), http.StatusFound)
			}
		}

		return false, nil
	}

	end, err := f()
	if err != nil {
		return web.RenderError(ctx, w, r, err, h.Renderer, TmplLayoutBase, TmplContentErrorGeneric, web.MIMETextHTMLCharsetUTF8)
	} else if end {
		return nil
	}

	sub, err := h.Repo.ReadByID(ctx, claims, periodID)
	if err != nil {
		return err
	}
	data["period"] = sub.Response(ctx)
	data["urlPeriodsIndex"] = urlPeriodsIndex()
	data["urlPeriodsView"] = urlPeriodsView(periodID)
	data["urlPeriodsUpdate"] = urlPeriodsUpdate(periodID)

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-periods-view.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// Update handles updating a period.
func (h *Periods) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	ctxValues, err := webcontext.ContextValues(ctx)
	if err != nil {
		return err
	}

	periodID := params["period_id"]

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	//
	req := new(period.UpdateRequest)
	data := make(map[string]interface{})
	f := func() (bool, error) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				return false, err
			}

			decoder := schema.NewDecoder()
			decoder.IgnoreUnknownKeys(true)

			if err := decoder.Decode(req, r.PostForm); err != nil {
				return false, err
			}
			req.ID = periodID

			err = h.Repo.Update(ctx, claims, *req, ctxValues.Now)
			if err != nil {
				switch errors.Cause(err) {
				default:
					if verr, ok := weberror.NewValidationError(ctx, err); ok {
						data["validationErrors"] = verr.(*weberror.Error)
						return false, nil
					} else {
						return false, err
					}
				}
			}

			// Display a success message to the checklist.
			webcontext.SessionFlashSuccess(ctx,
				"Period Updated",
				"Period successfully updated.")

			return true, web.Redirect(ctx, w, r, urlPeriodsView(req.ID), http.StatusFound)
		}

		return false, nil
	}

	end, err := f()
	if err != nil {
		return web.RenderError(ctx, w, r, err, h.Renderer, TmplLayoutBase, TmplContentErrorGeneric, web.MIMETextHTMLCharsetUTF8)
	} else if end {
		return nil
	}

	prd, err := h.Repo.ReadByID(ctx, claims, periodID)
	if err != nil {
		return err
	}
	data["period"] = prd.Response(ctx)

	data["urlPeriodsIndex"] = urlPeriodsIndex()
	data["urlPeriodsView"] = urlPeriodsView(periodID)

	if req.ID == "" {
		req.StartHour = &prd.StartHour
		req.StartMinute = &prd.StartMinute
		req.EndHour = &prd.EndHour
		req.EndMinute = &prd.EndMinute
	}
	data["form"] = req

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(subject.UpdateRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-periods-update.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}
