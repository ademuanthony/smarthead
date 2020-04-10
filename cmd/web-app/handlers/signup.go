package handlers

import (
	"context"
	"net/http"
	"time"

	"remoteschool/smarthead/internal/account"
	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/geonames"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/platform/web/weberror"
	"remoteschool/smarthead/internal/signup"
	"remoteschool/smarthead/internal/student"
	"remoteschool/smarthead/internal/subject"
	"remoteschool/smarthead/internal/subscription"
	"remoteschool/smarthead/internal/user_auth"

	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Signup represents the Signup API method handler set.
type Signup struct {
	AccountRepo      *account.Repository
	SignupRepo       *signup.Repository
	AuthRepo         *user_auth.Repository
	GeoRepo          *geonames.Repository
	StudentRepo      *student.Repository
	ClassRepo        *class.Repository
	SubscriptionRepo *subscription.Repository
	SubjectRepo      *subject.Repository
	MasterDB         *sqlx.DB
	Renderer         web.Renderer
}

// Step1 handles collecting the first detailed needed to create a new account.
func (h *Signup) Step1(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {

	ctxValues, err := webcontext.ContextValues(ctx)
	if err != nil {
		return err
	}

	//
	req := new(signup.SignupRequest)
	data := make(map[string]interface{})
	f := func() (bool, error) {
		claims, _ := auth.ClaimsFromContext(ctx)

		if r.Method == http.MethodPost {

			err := r.ParseForm()
			if err != nil {
				return false, err
			}

			decoder := schema.NewDecoder()
			if err := decoder.Decode(req, r.PostForm); err != nil {
				return false, err
			}

			var isFirst bool
			if req.Account.Name == "" {
				id, err := h.AccountRepo.First(ctx)
				if err != nil {
					isFirst = true
					acc, err := h.AccountRepo.Create(ctx, claims, account.AccountCreateRequest{
						Name: "Main Account",
					}, ctxValues.Now)
					if err != nil {
						return false, err
					}
					id = &acc.ID
				}
				req.Account.ID = *id
			}

			// Execute the account / user signup.
			_, err = h.SignupRepo.Signup(ctx, claims, *req, ctxValues.Now)
			if err != nil {
				switch errors.Cause(err) {
				case account.ErrForbidden:
					return false, web.RespondError(ctx, w, weberror.NewError(ctx, err, http.StatusForbidden))
				default:
					if verr, ok := weberror.NewValidationError(ctx, err); ok {
						data["validationErrors"] = verr.(*weberror.Error)
						return false, nil
					} else {
						return false, err
					}
				}
			}

			if !isFirst {
				// create the student account
				s, err := h.StudentRepo.Create(ctx, student.CreateRequest{
					Name:        req.User.FirstName + " " + req.User.LastName,
					ParentEmail: req.User.Email,
					ParentPhone: req.User.Phone,
					Username:    req.User.Email,
					ClassID:     req.ClassID,
				}, ctxValues.Now)
				if err != nil {
					return false, err
				}

				// create the one week trail lesson
				startDate := subscription.NextMonday(ctxValues.Now)
				endDate := startDate.Add(7 * 24 * time.Hour)

				period, err := h.SubscriptionRepo.TrailPeriodID(ctx)
				if err != nil {
					return false, err
				}
				eng, err := h.SubjectRepo.EnglishID(ctx)
				if err != nil {
					return false, err
				}

				// Eng trail
				subReq := subscription.CreateRequest{ 
					StudentID: s.ID,
					StartDate: startDate.Unix(),
					EndDate:   endDate.Unix(),
					PeriodID:  period.ID,
					SubjectID: eng.ID,
					ClassID:   s.ClassID,
				}

				_, err = h.SubscriptionRepo.Create(ctx, claims, subReq, ctxValues.Now)

				if err != nil {
					return false, errors.New("Unable to create free trial for your new account. Please contact the admin")
				}

				maths, err := h.SubjectRepo.MathsID(ctx)
				if err != nil {
					return false, err
				}
				// Maths trail
				subReq = subscription.CreateRequest{
					StudentID: s.ID,
					StartDate: startDate.Unix(),
					EndDate:   endDate.Unix(),
					PeriodID:  period.ID,
					SubjectID: maths.ID,
					ClassID:   s.ClassID,
				}

				_, err = h.SubscriptionRepo.Create(ctx, claims, subReq, ctxValues.Now)

				if err != nil {
					return false, errors.New("Unable to create free trial for your new account. Please contact the admin")
				}
			}

			// Authenticate the new user.
			token, err := h.AuthRepo.Authenticate(ctx, user_auth.AuthenticateRequest{
				Email:    req.User.Email,
				Password: req.User.Password,
			}, time.Hour, ctxValues.Now)
			if err != nil {
				return false, err
			}

			// Add the token to the users session.
			err = handleSessionToken(ctx, w, r, token)
			if err != nil {
				return false, err
			}

			// Display a welcome message to the user.
			webcontext.SessionFlashSuccess(ctx,
				"Thank you for Joining",
				"You workflow will be a breeze starting today.")

			// Redirect the user to the dashboard.
			return true, web.Redirect(ctx, w, r, "/?s=new", http.StatusFound)
		}

		return false, nil
	}

	end, err := f()
	if err != nil {
		return web.RenderError(ctx, w, r, err, h.Renderer, TmplLayoutBase, TmplContentErrorGeneric, web.MIMETextHTMLCharsetUTF8)
	} else if end {
		err = webcontext.ContextSession(ctx).Save(r, w)
		if err != nil {
			return err
		}
		return nil
	}

	data["geonameCountries"] = geonames.ValidGeonameCountries(ctx)

	data["countries"], err = h.GeoRepo.FindCountries(ctx, "name", "")
	if err != nil {
		return err
	}

	data["classes"], err = h.ClassRepo.Find(ctx, class.FindRequest{
		Order: []string{"school_order", "name"},
	})
	if err != nil {
		return err
	}

	data["form"] = req

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(signup.SignupRequest{})); ok {
		data["validationDefaults"] = verr.(*weberror.Error)
	}

	return h.Renderer.Render(ctx, w, r, TmplLayoutBase, "signup-step1.gohtml", web.MIMETextHTMLCharsetUTF8, http.StatusOK, data)
}
