package handlers

import (
	"context"
	"database/sql"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"remoteschool/smarthead/internal/account"
	"remoteschool/smarthead/internal/class"
	"remoteschool/smarthead/internal/deposit"
	"remoteschool/smarthead/internal/geonames"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/notify"
	"remoteschool/smarthead/internal/platform/web"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/platform/web/weberror"
	"remoteschool/smarthead/internal/postgres/models"
	"remoteschool/smarthead/internal/signup"
	"remoteschool/smarthead/internal/student"
	"remoteschool/smarthead/internal/subclass"
	"remoteschool/smarthead/internal/subject"
	"remoteschool/smarthead/internal/subscription"
	"remoteschool/smarthead/internal/user_auth"

	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Signup represents the Signup API method handler set.
type Signup struct {
	AccountRepo      *account.Repository
	SignupRepo       *signup.Repository
	AuthRepo         *user_auth.Repository
	GeoRepo          *geonames.Repository
	StudentRepo      *student.Repository
	ClassRepo        *class.Repository
	SubclassRepo	 *subclass.Repository
	SubscriptionRepo *subscription.Repository
	SubjectRepo      *subject.Repository
	DepositRepo      *deposit.Repository
	MasterDB         *sqlx.DB
	Renderer         web.Renderer
	EmailNotifier	 notify.Email
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
					status := account.AccountStatus_Active
					acc, err := h.AccountRepo.Create(ctx, claims, account.AccountCreateRequest{
						Name:   "Main Account",
						Status: &status,
					}, ctxValues.Now)
					if err != nil {
						return false, err
					}
					id = &acc.ID
				}
				req.Account.ID = *id
			}

			// Execute the account / user signup.
			signupResult, err := h.SignupRepo.Signup(ctx, claims, *req, ctxValues.Now)
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

			claims.Audience = signupResult.User.ID

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
				trailDeposit, err := h.DepositRepo.TrailDeposit(ctx)
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
					DepositID: trailDeposit.ID,
				}

				_, err = h.SubscriptionRepo.Create(ctx, claims, subReq, ctxValues.Now)

				if err != nil {
					return false, weberror.NewErrorMessage(ctx, err, 400, "Unable to create free trial for your new account. Please contact the admin")
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
					DepositID: trailDeposit.ID,
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

func (h *Signup) GetStarted(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v, err := webcontext.ContextValues(ctx)
	if err != nil {
		return err
	}

	// Claims are optional as authentication is not required ATM for this method.
	claims, _ := auth.ClaimsFromContext(ctx)

	var req struct {
		ClassID string `json:"class_id"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	}

	if err := web.Decode(ctx, r, &req); err != nil {
		if _, ok := errors.Cause(err).(*weberror.Error); !ok {
			err = weberror.NewError(ctx, err, http.StatusBadRequest)
		}
		return web.RespondJsonError(ctx, w, err)
	}

	subclassA, err := h.SubclassRepo.NextSubclass(ctx, req.ClassID)
	if err != nil {
		return web.RespondJsonError(ctx, w, err)
	}

	pass := randomPassword()
	names := strings.Split(req.Name, " ")
	firstName := names[0]
	lastName := names[len(names) - 1]
	regReq := signup.SignupRequest{
		ClassID: req.ClassID,
		User: signup.SignupUser{
			Email:           req.Email,
			Phone:           req.Phone,
			FirstName:       firstName,
			LastName: 		 lastName,
			Password:        pass,
			PasswordConfirm: pass,
		},
	}
	id, err := h.AccountRepo.First(ctx)
	if err != nil {
		status := account.AccountStatus_Active
		acc, err := h.AccountRepo.Create(ctx, claims, account.AccountCreateRequest{
			Name:   "Main Account",
			Status: &status,
		}, v.Now)
		if err != nil {
			return err
		}
		id = &acc.ID
	}
	regReq.Account.ID = *id
	res, err := h.SignupRepo.Signup(ctx, claims, regReq, v.Now)
	if err != nil {
		switch errors.Cause(err) {
		case account.ErrForbidden:
			return web.RespondJsonError(ctx, w, weberror.NewError(ctx, err, http.StatusForbidden))
		default:
			_, ok := err.(validator.ValidationErrors)
			if ok {
				return web.RespondJsonError(ctx, w, weberror.NewError(ctx, err, http.StatusBadRequest))
			}

			return errors.Wrapf(err, "Signup: %+v", &req)
		}
	}

	claims.Audience = res.User.ID

	// create the student account
	s, err := h.StudentRepo.Create(ctx, student.CreateRequest{
		Name:        req.Name,
		ParentEmail: req.Email,
		ParentPhone: req.Phone,
		Username:    req.Email,
		ClassID:     req.ClassID,
		SubclassID:	 subclassA.ID,
	}, v.Now)
	if err != nil {
		return err
	}

	// create the one week trail lesson
	startDate := subscription.NextMonday(v.Now)
	startDate = time.Now()
	endDate := startDate.Add(7 * 24 * time.Hour)

	period, err := h.SubscriptionRepo.TrailPeriodID(ctx)
	if err != nil {
		return err
	}

	maths, err := h.SubjectRepo.MathsID(ctx)
	if err != nil {
		return err
	}

	trailDeposit, err := h.DepositRepo.TrailDeposit(ctx)
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error(){
			return err
		}
		dept := models.Deposit{
			ID: uuid.NewRandom().String(),
			Amount: 0,
			Channel: "trail",
			ClassID: req.ClassID,
			CreatedAt: time.Now(),
			StudentID: s.ID,
			SubjectID: maths.ID,
			Status: "paid",
		}
		err = h.DepositRepo.Insert(ctx, dept)
		if err != nil {
			return err
		}
		trailDeposit = &deposit.Deposit{
			ID: dept.ID,
		}
	}

	// Maths trail
	subReq := subscription.CreateRequest{
		StudentID: s.ID,
		StartDate: startDate.Unix(),
		EndDate:   endDate.Unix(),
		PeriodID:  period.ID,
		SubjectID: maths.ID,
		ClassID:   s.ClassID,
		DepositID: trailDeposit.ID,
	}

	_, err = h.SubscriptionRepo.Create(ctx, claims, subReq, v.Now)

	if err != nil {
		return errors.New("Unable to create free trial for your new account. Please contact the admin")
	}

	data := map[string]interface{}{
		"Name": req.Name,
		"Email": req.Email,
		"Password": pass,
		"Lesson1Date": "Monday, 8:00 AM",
		"Lesson2Date": "Wednesday, 8:00 AM",
		"Subject1": maths.Name,
		"Subject2": maths.Name,
	}
	err = h.EmailNotifier.Send(ctx, req.Email, "Welcome to Remote School", "welcome_email", data)
	if err != nil {
		return err
	}
	return web.RespondJson(ctx, w, res.Response(ctx), http.StatusCreated)
}

func (h *Signup) ThankYou(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	return h.Renderer.Render(ctx, w, r, tmplLayoutSite, "dtox-thank-you.html", web.MIMETextHTMLCharsetUTF8, http.StatusOK, nil)
}

func randomPassword () string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (h *Signup) Ping(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	ctxValues, err := webcontext.ContextValues(ctx)
	if err != nil {
		return err
	}
	claims, _ := auth.ClaimsFromContext(ctx)
	err = h.SignupRepo.CreateDefaultAdmin(ctx, claims, ctxValues.Now)
	if err != nil {
		return err
	}
	return web.Redirect(ctx, w, r, "/user/login", http.StatusCreated)
}
