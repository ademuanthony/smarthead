package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/datatable"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/platform/web/weberror"
	"remoteschool/smarthead/internal/subclass"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

// Sublasses represents the Sublasses API method handler set.
type Sublasses struct {
	Repo     *subclass.Repository
	Redis    *redis.Client
	Renderer web.Renderer
}

func urlSubclassesIndex() string {
	return fmt.Sprintf("/admin/subclasses")
}

func urlSubclassesCreate() string {
	return fmt.Sprintf("/admin/subclasses/create")
}

func urlSubclassesView(classID string) string {
	return fmt.Sprintf("/admin/subclasses/%s", classID)
}

func urlSubclassesUpdate(classID string) string {
	return fmt.Sprintf("/admin/subclasses/%s/update", classID)
}

// Index handles listing all the classes for the current account.
func (h *Subclasses) Index(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	fields := []datatable.DisplayField{
		{Field: "id", Title: "ID", Visible: false, Searchable: true, Orderable: true, Filterable: false},
		{Field: "subclass", Title: "Subclass", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Name"},
		{Field: "class", Title: "Class", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Class"},
		{Field: "school_order", Title: "School Order", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter School"},
	}

	mapFunc := func(q *subclass.Response, cols []datatable.DisplayField) (resp []datatable.ColumnValue, err error) {
		for i := 0; i < len(cols); i++ {
			col := cols[i]
			var v datatable.ColumnValue
			switch col.Field {
			case "id":
				v.Value = fmt.Sprintf("%s", q.ID)
			case "name":
				v.Value = q.Name
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlSubclassesView(q.ID), v.Value)
			case "class":
				v.Value = q.Class
				v.Formatted = fmt.Sprintf("<a href='%s'>%s</a>", urlClassesView(q.ClassID), v.Value)
			case "school_order":
				v.Value = fmt.Sprint(q.SchoolOrder)
				v.Formatted = v.Value
			default:
				return resp, errors.Errorf("Failed to map value for %s.", col.Field)
			}
			resp = append(resp, v)
		}

		return resp, nil
	}

	loadFunc := func(ctx context.Context, sorting string, fields []datatable.DisplayField) (resp [][]datatable.ColumnValue, err error) {
		res, err := h.Repo.Find(ctx, class.FindRequest{
			Order: strings.Split(sorting, ","),
		})
		if err != nil {
			return resp, err
		}

		for _, a := range res {
			l, err := mapFunc(a, fields)
			if err != nil {
				return resp, errors.Wrapf(err, "Failed to map class for display.")
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
		"datatable":        dt.Response(),
		"urlSubclassesCreate": urlSubclassesCreate(),
		"urlSubclassesIndex":  urlSubclassesIndex(),
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subclasses-index.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// Create handles creating a new class.
func (h *Subclasses) Create(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err 
	}

	//
	req := new(class.CreateRequest)
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
				"Class Created",
				"Class successfully created.")

			return true, web.Redirect(ctx, w, r, urlClassesView(sub.ID), http.StatusFound)
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
	data["urlClassesIndex"] = urlClassesIndex()
	spew.Dump(req)

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(class.CreateRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-classes-create.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// View handles displaying a classes.
func (h *Classes) View(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	classID := params["class_id"]

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
				err = h.Repo.Delete(ctx, claims, class.DeleteRequest{
					ID: classID,
				})
				if err != nil {
					return false, err
				}

				webcontext.SessionFlashSuccess(ctx,
					"Class Archive",
					"Class successfully archive.")

				return true, web.Redirect(ctx, w, r, urlClassesIndex(), http.StatusFound)
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

	sub, err := h.Repo.ReadByID(ctx, claims, classID)
	if err != nil {
		return err
	}
	data["class"] = sub.Response(ctx)
	data["urlClassesIndex"] = urlClassesIndex()
	data["urlClassesView"] = urlClassesView(classID)
	data["urlClassesUpdate"] = urlClassesUpdate(classID)

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-classes-view.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// Update handles updating a class.
func (h *Classes) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	classID := params["class_id"]

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	//
	req := new(class.UpdateRequest)
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
			req.ID = classID

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

			// Display a success message to the class.
			webcontext.SessionFlashSuccess(ctx,
				"Class Updated",
				"Class successfully updated.")

			return true, web.Redirect(ctx, w, r, urlClassesView(req.ID), http.StatusFound)
		}

		return false, nil
	}

	end, err := f()
	if err != nil {
		return web.RenderError(ctx, w, r, err, h.Renderer, TmplLayoutBase, TmplContentErrorGeneric, web.MIMETextHTMLCharsetUTF8)
	} else if end {
		return nil
	}

	sub, err := h.Repo.ReadByID(ctx, claims, classID)
	if err != nil {
		return err
	}
	data["class"] = sub.Response(ctx)

	data["urlClassesIndex"] = urlClassesIndex()
	data["urlClassesView"] = urlClassesView(classID)

	if req.ID == "" {
		req.Name = &sub.Name
	}
	data["form"] = req

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(class.UpdateRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-classes-update.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}
