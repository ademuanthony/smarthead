package deposit

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"remoteschool/smarthead/internal/paystack"
	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/postgres/models"
	"remoteschool/smarthead/internal/subscription"

	"github.com/opentracing/opentracing-go/log"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var (
	// ErrNotFound abstracts the postgres not found error.
	ErrNotFound = errors.New("Entity not found")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("Attempted action is not allowed")
)

// Find gets all the subscription from the database based on the request params.
func (repo *Repository) Find(ctx context.Context, _ auth.Claims, req FindRequest) (Deposits, error) {
	var queries []QueryMod

	queries = append(queries, qm.Load(models.DepositRels.Student), qm.Load(models.DepositRels.Class))

	if req.Where != "" {
		queries = append(queries, Where(req.Where, req.Args...))
	}

	if len(req.Order) > 0 {
		for _, s := range req.Order {
			queries = append(queries, OrderBy(s))
		}
	}

	if req.Limit != nil {
		queries = append(queries, Limit(int(*req.Limit)))
	}

	if req.Offset != nil {
		queries = append(queries, Offset(int(*req.Offset)))
	}

	depositSlice, err := models.Deposits(queries...).All(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	var result Deposits
	for _, rec := range depositSlice {
		result = append(result, FromModel(rec))
	}

	return result, nil
}

// ReadByID gets the specified deposit by ID from the database.
func (repo *Repository) ReadByID(ctx context.Context, claims auth.Claims, id string) (*Deposit, error) {
	depositModel, err := models.FindDeposit(ctx, repo.DbConn, id)
	if err != nil {
		return nil, err
	}

	return FromModel(depositModel), nil
}

func (repo *Repository) TrailDeposit(ctx context.Context) (*Deposit, error) {
	m, err := models.Deposits(qm.Limit(1)).One(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}
	return FromModel(m), nil
}

func subscriptionAmount(reqCount int) int {
	var amount int
	if reqCount >= 5 {
		fives := (reqCount - reqCount%5) / 5
		amount = 1500000 * fives
		reqCount -= fives * 5
	}
	if reqCount >= 3 {
		threes := (reqCount - reqCount%3) / 3
		amount += 1200000 * threes
		reqCount -= threes * 3
	}
	amount += reqCount * 500000

	return amount
}

// Create inserts a new subscription into the database.
func (repo *Repository) Create(ctx context.Context, claims auth.Claims, req CreateRequest, now time.Time) (*Deposit, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.deposit.Create")
	defer span.Finish()
	if claims.Audience == "" {
		return nil, errors.WithStack(ErrForbidden)
	}

	// Validate the request.
	v := webcontext.Validator()
	err := v.StructCtx(ctx, req)
	if err != nil {
		return nil, err
	}

	var amount = subscriptionAmount(req.Count)

	// If now empty set it to the current time.
	if now.IsZero() {
		now = time.Now()
	}

	// Always store the time as UTC.
	now = now.UTC()
	// Postgres truncates times to milliseconds when storing. We and do the same
	// here so the value we return is consistent with what we store.
	now = now.Truncate(time.Millisecond)
	m := models.Deposit{
		ID:         uuid.NewRandom().String(),
		StudentID:  req.StudentID,
		SubjectID:  req.SubjectID,
		DaysOfWeek: req.DaysOfWeek,
		PeriodID:   null.StringFrom(req.PeriodID),
		ClassID:    req.ClassID,
		CreatedAt:  now,
		Amount:     amount,
		Channel:    req.Channel,
		Status:     req.Status,
		Ref:        req.Ref,
		PaymentRef: req.PaymentRef,
	}

	if err := m.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return nil, errors.WithMessage(err, "Insert deposit failed")
	}

	return &Deposit{
		ID:        m.ID,
		StudentID: m.StudentID,
		CreatedAt: now,
		Amount:    m.Amount,
		Channel:   m.Channel,
		Status:    m.Status,
		Ref:       m.Ref,
	}, nil
}

func (repo *Repository) Insert(ctx context.Context, dept models.Deposit) error {
	return dept.Insert(ctx, repo.DbConn, boil.Infer())
}

// Update replaces an subject in the database.
func (repo *Repository) Update(ctx context.Context, claims auth.Claims, req UpdateRequest, now time.Time) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.deposit.Update")
	defer span.Finish()

	if claims.Audience == "" {
		return errors.WithStack(ErrForbidden)
	}

	// Validate the request.
	v := webcontext.Validator()
	err := v.StructCtx(ctx, req)
	if err != nil {
		return err
	}

	// If now empty set it to the current time.
	if now.IsZero() {
		now = time.Now()
	}

	// Always store the time as UTC.
	now = now.UTC()

	cols := models.M{}
	cols[models.DepositColumns.UpdatedAt] = now

	if req.Ref != nil {
		cols[models.DepositColumns.Ref] = *req.Ref
	}

	if req.Status != nil {
		cols[models.DepositColumns.Status] = *req.Status
	}

	if len(cols) == 0 {
		return nil
	}

	_, err = models.Deposits(models.DepositWhere.ID.EQ(req.ID)).UpdateAll(ctx, repo.DbConn, cols)

	return nil
}

// ProccessPaystackSubscriptionEvent
func (repo *Repository) ProccessPaystackSubscriptionEvent(ctx context.Context, req PaystackSubscriptionEvent, now time.Time) error {
	switch req.Event {
	default:
		return errors.New("unsupported event")
	case "charge.success":
		client := paystack.NewClient(repo.PaystackSecret, http.DefaultClient)
		payment, err := client.Transaction.Verify(req.Data.Reference)
		if err != nil {
			// TODO: inform the admin of verification problem
			return err
		}
		switch req.Data.Plan {
		default:
			return errors.New("unsupported plan")
		case repo.PrimaryPlan:
			if payment.Amount < repo.PrimaryAmount {
				return errors.New("partial payment received")
			}
		case repo.SecondaryPlan:
			if payment.Amount < repo.SecondaryAmount {
				return errors.New("partial payment received")
			}
		}
		student, err := models.Students(
			models.StudentWhere.RegNo.EQ(req.Data.Customer.Metadata.RegistrationNumber),
			qm.Load(models.StudentRels.Subclass),
		).One(ctx, repo.DbConn)
		if err != nil {
			// send registration instruction to user
			return errors.New("student not found")
		}
		cols := models.M{}
		cols[models.StudentColumns.LastPaymentDate] = now
		var subclassID string
		if !student.SubclassID.Valid {
			subclassID = repo.getAvailableSubclass(ctx, student.ClassID.String)
		} else if strings.Contains(strings.ToLower(student.R.Subclass.Name), "free") {
			subclassID = repo.getAvailableSubclass(ctx, student.ClassID.String)
		}
		if subclassID != "" {
			cols[models.StudentColumns.SubclassID] = subclassID
		}
		if _, err = models.Students(models.StudentWhere.ID.EQ(student.ID)).UpdateAll(ctx, repo.DbConn, cols); err != nil {
			return err
		}
		if err = repo.NotifyEmail.Send(ctx, student.ParentEmail, "Payment Received", "payment_received", map[string]interface{}{
			"Name": student.Name, "Amount": payment.Amount,
		}); err != nil {
			log.Error(err)
		}
		return nil
	}
}

func (repo *Repository) getAvailableSubclass(ctx context.Context, classID string) string {
	subclasses, err := models.Subclasses(
		models.SubclassWhere.ClassID.EQ(classID),
		qm.OrderBy(models.SubclassColumns.Name),
	).All(ctx, repo.DbConn)

	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return ""
		}
	} else {
		for _, s := range subclasses {
			count, err := models.Students(models.StudentWhere.SubclassID.EQ(null.StringFrom(s.ID))).Count(ctx, repo.DbConn)
			if err != nil {
				log.Error(err)
			}
			if count < 10 {
				return s.ID
			}
		}
	}

	class, err := models.FindClass(ctx, repo.DbConn, classID)
	if err != nil {
		return ""
	}
	subclass := models.Subclass{
		ID:          uuid.NewRandom().String(),
		ClassID:     classID,
		Name:        fmt.Sprintf("%sA", class.Name),
		SchoolOrder: class.SchoolOrder,
	}
	if err = subclass.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return ""
	}
	return subclass.ID

	return ""
}

// UpdateStatus updates the status of the the supplied deposit by quering the channel
func (repo *Repository) UpdateStatus(ctx context.Context, req UpdateStatusRequest, claims auth.Claims, now time.Time) ([]*subscription.Subscription, error) {
	depositModel, err := models.FindDeposit(ctx, repo.DbConn, req.ID)
	if err != nil {
		return nil, err
	}
	if depositModel.Status == StatusPaid {
		return nil, errors.New("duplicate payment verification request")
	}
	client := paystack.NewClient(repo.PaystackSecret, http.DefaultClient)
	payment, err := client.Transaction.Verify(req.ID)
	if err != nil {
		// TODO: inform the admin of verification problem
		return nil, err
	}
	if int(payment.Amount)*100 < depositModel.Amount {
		return nil, errors.Errorf("partial payment received. Expected %d, got %f", depositModel.Amount/100, payment.Amount)
	}

	tx, err := repo.DbConn.Begin()
	if err != nil {
		return nil, err
	}

	depositModel.Status = StatusPaid
	_, err = depositModel.Update(ctx, tx, boil.Infer())
	if err != nil {
		_ = tx.Rollback()
		//Todo: log fatal error for admin to resolve
		return nil, errors.New("payment received but unable to update status. contact admin for help")
	}

	amount := subscriptionAmount(len(req.Items))
	if amount > depositModel.Amount {
		_ = tx.Rollback()
		return nil, errors.New("Wrong amount received. Please contact the admin for help")
	}

	period, err := repo.SubscriptionRepo.TrailPeriodID(ctx)
	if err != nil {
		return nil, err
	}

	startDate := subscription.NextMonday(now)
	endDate := startDate.Add(30 * 24 * time.Hour)

	var subs []*subscription.Subscription

	for _, item := range req.Items {
		if item.PeriodID == "" {
			item.PeriodID = period.ID
		}
		subReq := subscription.CreateRequest{
			StudentID:  depositModel.StudentID,
			StartDate:  startDate.Unix(),
			EndDate:    endDate.Unix(),
			DaysOfWeek: depositModel.DaysOfWeek,
			PeriodID:   item.PeriodID,
			SubjectID:  item.SubjectID,
			ClassID:    item.ClassID,
			DepositID:  depositModel.ID,
		}

		sub, err := repo.SubscriptionRepo.CreateTx(ctx, tx, claims, subReq, now)

		if err != nil {
			_ = tx.Rollback()
			//TODO: log critical error and inform admin
			return nil, errors.New("payment received but unable to create subscription. Please contact the admin")
		}
		subs = append(subs, sub)
	}

	depositModel.Status = StatusSubscribed
	_, err = depositModel.Update(ctx, tx, boil.Infer())
	if err != nil {
		_ = tx.Rollback()
		//TODO: log fatal error for admin to resolve
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		// TODO: log critical error
		_ = tx.Rollback()
		return nil, err
	}

	return subs, nil
}

// AddManualDeposit
func (repo *Repository) AddManualDeposit(ctx context.Context, req AddManualDepositRequest, claims auth.Claims, now time.Time) error {
	if !claims.HasRole(auth.RoleAdmin) {
		return ErrForbidden
	}

	v := webcontext.Validator()
	err := v.StructCtx(ctx, req)
	if err != nil {
		return err
	}

	tx, err := repo.DbConn.Begin()
	if err != nil {
		return errors.Wrap(err, "starting transaction")
	}

	student, err := models.Students(models.StudentWhere.RegNo.EQ(req.StudentRegNo)).One(ctx, repo.DbConn)
	if err != nil {
		return errors.New("invalid Reg number")
	}

	if !student.ClassID.Valid {
		return errors.New("Please set the student class to continue")
	}

	if len(req.Items) == 0 {
		return errors.New("please add at least one subject")
	}

	// If now empty set it to the current time.
	if now.IsZero() {
		now = time.Now()
	}

	// Always store the time as UTC.
	now = now.UTC()
	// Postgres truncates times to milliseconds when storing. We and do the same
	// here so the value we return is consistent with what we store.
	now = now.Truncate(time.Millisecond)

	period, err := repo.SubscriptionRepo.TrailPeriodID(ctx)
	if err != nil {
		return errors.New("could not get a trail period")
	}

	amount := subscriptionAmount(len(req.Items))

	m := models.Deposit{
		ID:         uuid.NewRandom().String(),
		StudentID:  student.ID,
		SubjectID:  req.Items[0].SubjectID,
		DaysOfWeek: 1,
		PeriodID:   null.StringFrom(period.ID),
		ClassID:    student.ClassID.String,
		CreatedAt:  now,
		Amount:     amount,
		Channel:    "Manual",
		Status:     StatusSubscribed,
		Ref:        "manual",
		PaymentRef: "manual",
	}

	if err := m.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return errors.WithMessage(err, "Insert deposit failed")
	}

	startDate, err := time.Parse("01/02/2006", req.StartDate)
	if err != nil {
		startDate = now
	}

	endDate, err := time.Parse("01/02/2006", req.StartDate)
	if err != nil {
		endDate = now.Add(30 * time.Hour * 24)
	}

	for _, item := range req.Items {
		if item.PeriodID == "" {
			item.PeriodID = period.ID
		}
		subReq := subscription.CreateRequest{
			StudentID:  student.ID,
			StartDate:  startDate.Unix(),
			EndDate:    endDate.Unix(),
			DaysOfWeek: 1,
			PeriodID:   period.ID,
			SubjectID:  item.SubjectID,
			ClassID:    student.ClassID.String,
			DepositID:  m.ID,
		}

		_, err := repo.SubscriptionRepo.CreateTx(ctx, tx, claims, subReq, now)

		if err != nil {
			_ = tx.Rollback()
			//TODO: log critical error and inform admin
			log.Error(err)
			return errors.Wrap(err, "Error creating subscription")
		}
	}

	if err = tx.Commit(); err != nil {
		// TODO: log critical error
		_ = tx.Rollback()
		return err
	}
	return nil
}

// Delete removes a deposit from the database.
func (repo *Repository) Delete(ctx context.Context, claims auth.Claims, req DeleteRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.deposit.Delete")
	defer span.Finish()

	// Validate the request.
	v := webcontext.Validator()
	err := v.Struct(req)
	if err != nil {
		return err
	}

	if claims.Audience == "" {
		return errors.WithStack(ErrForbidden)
	}
	// Admin users can update Categories they have access to.
	if !claims.HasRole(auth.RoleAdmin) {
		return errors.WithStack(ErrForbidden)
	}

	_, err = models.Deposits(models.DepositWhere.ID.EQ(req.ID)).DeleteAll(ctx, repo.DbConn)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
