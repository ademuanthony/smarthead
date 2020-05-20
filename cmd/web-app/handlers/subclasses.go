package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/datatable"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/platform/web/weberror"
	"remoteschool/smarthead/internal/student"
	"remoteschool/smarthead/internal/subclass"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"
)

// Subclasses represents the Subclasses API method handler set.
type Subclasses struct {
	Repo      *subclass.Repository
	ClassRepo *class.Repository
	StudentRepo *student.Repository
	Redis     *redis.Client
	Renderer  web.Renderer
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
		{Field: "class", Title: "Class", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Class"},
		{Field: "name", Title: "Subclass", Visible: true, Searchable: true, Orderable: true, Filterable: true, FilterPlaceholder: "filter Name"},
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
		result, err := h.Repo.Find(ctx, subclass.FindRequest{
			Order: strings.Split(sorting, ","),
			IncludeClass: true,
		})
		if err != nil {
			return resp, err
		}

		res := result.Response(ctx)
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
		"datatable":           dt.Response(),
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
	req := new(subclass.CreateRequest)
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

			return true, web.Redirect(ctx, w, r, urlSubclassesView(sub.ID), http.StatusFound)
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
	data["urlSubclassesIndex"] = urlSubclassesIndex()
	classes, err := h.ClassRepo.Find(ctx, class.FindRequest{
		Order: []string{"school_order", "name"},
	})
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return err
		}
	}
	data["classes"] = classes

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(class.CreateRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subclasses-create.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// View handles displaying a classes.
func (h *Subclasses) View(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	classID := params["subclass_id"]

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	ctxValue, err := webcontext.ContextValues(ctx)
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
				err = h.Repo.Delete(ctx, claims, subclass.DeleteRequest{
					ID: classID,
				})
				if err != nil { 
					return false, err
				}

				webcontext.SessionFlashSuccess(ctx,
					"Class Archive",
					"Class successfully archive.")

				return true, web.Redirect(ctx, w, r, urlClassesIndex(), http.StatusFound)
			case "add-student":
				r.ParseForm()
				regNo := r.FormValue("RegNo")
				if regNo == "" {
					webcontext.SessionFlashSuccess(ctx,
						"Invalid Request",
						"Student ID cannot be empty.")
						return false, nil
				}
				stud, err := h.StudentRepo.Find(ctx, claims, student.FindRequest{
					Where: "reg_no = $1",
					Args: []interface{}{regNo},
				})
				if err != nil || len(stud) == 0 {
					webcontext.SessionFlashSuccess(ctx,
						"Invalid Request",
						"Student not found")
						return false, nil
				}
				stuErr := h.StudentRepo.Update(ctx, claims, student.UpdateRequest{
					ID: stud[0].ID,
					SubclassID: &classID,
				}, ctxValue.Now)
				if stuErr != nil {
					return false, stuErr
				}
				webcontext.SessionFlashSuccess(ctx,
					"Student Added",
					"Student successfully added.")
				return false, nil
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
	cla := sub.Response(ctx)
	data["class"] = cla
	data["urlSubclassesIndex"] = urlSubclassesIndex()
	data["urlSubclassesView"] = urlSubclassesView(classID)
	data["urlSubclassesUpdate"] = urlSubclassesUpdate(classID)

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subclasses-view.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}

// Update handles updating a class.
func (h *Subclasses) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	classID := params["subclass_id"]

	claims, err := auth.ClaimsFromContext(ctx)
	if err != nil {
		return err
	}

	//
	req := new(subclass.UpdateRequest)
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

			return true, web.Redirect(ctx, w, r, urlSubclassesView(req.ID), http.StatusFound)
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

	data["urlSubclassesIndex"] = urlSubclassesIndex()
	data["urlSubclassesView"] = urlSubclassesView(classID)

	if req.ID == "" {
		req.Name = &sub.Name
		req.ClassID = &sub.ClassID
		req.SchoolOrder = &sub.SchoolOrder
	}
	data["form"] = req

	classes, err := h.ClassRepo.Find(ctx, class.FindRequest{
		Order: []string{"school_order", "name"},
	})
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return err
		}
	}
	data["classes"] = classes

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(class.UpdateRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "admin-subclasses-update.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}
