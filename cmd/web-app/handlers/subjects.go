package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"remoteschool/smarthead/internal/subject"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/datatable"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/platform/web/weberror"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

// Subjects represents the Subjects API method handler set.
type Subjects struct {
	Repo *subject.Repository
	Redis         *redis.Client
	Renderer      web.Renderer
}
 
func urlSubjectsIndex() string {
	return fmt.Sprintf("/admin/subjects")
}

func urlSubjectsCreate() string {
	return fmt.Sprintf("/admin/subjects/create")
}

func urlSubjectsView(subjectID string) string {
	return fmt.Sprintf("/admin/subjects/%s", subjectID) 
}
 
func urlSubjectsUpdate(subjectID string) string {
	return fmt.Sprintf("/admin/subjects/%s/update", subjectID)
}

// Index handles listing all the subjects for the current account.
func (h *Subjects) Index(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []datatable.DisplayField{
		{Field: "id", Title: "ID", Visible: false, Searchable: true, Orderable: true, Filterable: false},
		{Field: "name", Title: "Subject", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Name"},
		{Field: "school_order", Title: "School(s)", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Name"},
	}

	mapFunc := func(q *subject.Response, cols []datatable.DisplayField) (resp []datatable.ColumnValue, err error) {
		for i := 0; i < len(cols); i++ {
			col := cols[i]
			var v datatable.ColumnValue
			switch col.Field {
			case "id":
				v.Value = fmt.Sprintf("%s", q.ID)
			case "name":
				v.Value = q.Name
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlSubjectsView(q.ID), v.Value)
			case "school_order":
				v.Value = fmt.Sprintf("%v", q.SchoolOrders)
				v.Formatted = v.Value
			default:
				return resp, errors.Errorf("Failed to map value for %s.", col.Field)
			}
			resp = append(resp, v)
		}

		return resp, nil
	}

	loadFunc := func(ctx context.Context, sorting string, fields []datatable.DisplayField) (resp [][]datatable.ColumnValue, err error) {
		res, err := h.Repo.Find(ctx, claims, subject.FindRequest{
			Order: strings.Split(sorting, ","),
		})
		if err != nil {
			return resp, err
		}

		for _, a := range res {
			l, err := mapFunc(a.Response(ctx), fields)
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
		"urlSubjectsCreate": urlSubjectsCreate(),
		"urlSubjectsIndex": urlSubjectsIndex(),
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subjects-index.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// Create handles creating a new subject.
func (h *Subjects) Create(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	//
	req := new(subject.CreateRequest)
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

			sub, err := h.Repo.Create(ctx, claims, *req)
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
				"Subject Created",
				"Subject successfully created.")

			return true, web.Redirect(ctx, w, r, urlSubjectsView(sub.ID), http.StatusFound)
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
	data["urlSubjectsIndex"] = urlSubjectsIndex() 

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(subject.CreateRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subjects-create.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// View handles displaying a subjects.
func (h *Subjects) View(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	subjectID := params["subject_id"]

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
				err = h.Repo.Delete(ctx, claims, subject.DeleteRequest{
					ID: subjectID,
				})
				if err != nil {
					return false, err
				}

				webcontext.SessionFlashSuccess(ctx,
					"Subject Archive",
					"Subject successfully archive.")

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

	sub, err := h.Repo.ReadByID(ctx, claims, subjectID)
	if err != nil {
		return err
	}
	data["subject"] = sub.Response(ctx)
	data["urlSubjectsIndex"] = urlSubjectsIndex()
	data["urlSubjectsView"] = urlSubjectsView(subjectID)
	data["urlSubjectsUpdate"] = urlSubjectsUpdate(subjectID)

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subjects-view.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// Update handles updating a subject.
func (h *Subjects) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	subjectID := params["subject_id"]

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	} 

	req := new(subject.UpdateRequest)
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
			req.ID = subjectID
			err = h.Repo.Update(ctx, claims, *req)
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
				"Subject Updated",
				"Subject successfully updated.")

			return true, web.Redirect(ctx, w, r, urlSubjectsView(req.ID), http.StatusFound)
		}

		return false, nil
	}

	end, err := f()
	if err != nil {
		return web.RenderError(ctx, w, r, err, h.Renderer, TmplLayoutBase, TmplContentErrorGeneric, web.MIMETextHTMLCharsetUTF8)
	} else if end {
		return nil
	}

	sub, err := h.Repo.ReadByID(ctx, claims, subjectID)
	if err != nil {
		return err
	}
	data["subject"] = sub.Response(ctx)

	data["urlSubjectsIndex"] = urlSubjectsIndex()
	data["urlSubjectsView"] = urlSubjectsView(subjectID)

	if req.ID == "" {
		req.Name = &sub.Name
		var schoolOrders []string
		for _, s := range sub.SchoolOrders {
			schoolOrders = append(schoolOrders, fmt.Sprintf("%d", s))
		}
		s := strings.Join(schoolOrders, ", ")
		req.SchoolOrder = &s
	}
	data["form"] = req

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(subject.UpdateRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subjects-update.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}
