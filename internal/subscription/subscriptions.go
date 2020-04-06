package subscription

import (
	"context"
	"time"

	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/postgres/models"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
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
func (repo *Repository) Find(ctx context.Context, _ auth.Claims, req FindRequest) (Subscriptions, error) {
	var queries []QueryMod

	queries = append(queries, qm.Load(models.SubscriptionRels.Subject))
	queries = append(queries, qm.Load(models.SubscriptionRels.Student))
	queries = append(queries, qm.Load(models.SubscriptionRels.Period))
	queries = append(queries, qm.Load(models.SubscriptionRels.Class))

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

	subscriptionSlice, err := models.Subscriptions(queries...).All(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	var result Subscriptions
	for _, rec := range subscriptionSlice {
		result = append(result, FromModel(rec))
	}

	return result, nil
}

// ReadByID gets the specified subscription by ID from the database.
func (repo *Repository) ReadByID(ctx context.Context, claims auth.Claims, id string) (*Subscription, error) {
	subscriptionModel, err := models.FindSubscription(ctx, repo.DbConn, id)
	if err != nil {
		return nil, err
	}

	return FromModel(subscriptionModel), nil
}

func (repo *Repository) CountActiveSubscriptions(ctx context.Context, studentID string, now time.Time) (int64, error) {
	return models.Subscriptions(models.SubscriptionWhere.StudentID.EQ(studentID), 
	models.SubscriptionWhere.EndDate.GT(now.Unix())).Count(ctx, repo.DbConn)
}
// Create inserts a new subscription into the database.
func (repo *Repository) Create(ctx context.Context, claims auth.Claims, req CreateRequest, now time.Time) (*Subscription, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.subscription.Create")
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

	// If now empty set it to the current time.
	if now.IsZero() {
		now = time.Now()
	}

	// Always store the time as UTC.
	now = now.UTC()
	// Postgres truncates times to milliseconds when storing. We and do the same
	// here so the value we return is consistent with what we store.
	now = now.Truncate(time.Millisecond)
	m := models.Subscription{
		ID:         uuid.NewRandom().String(),
		DaysOfWeek: req.DaysOfWeek,
		EndDate:    req.EndDate,
		PeriodID:   req.PeriodID,
		ClassID:    req.ClassID,
		DepositID:  req.DepositID,
		StartDate:  req.StartDate,
		StudentID:  req.StudentID,
		SubjectID:  req.SubjectID,
		CreatedAt:  now.Unix(),
	}

	if err := m.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return nil, errors.WithMessage(err, "Insert subscription failed")
	}

	// TODO: get the associated subject and create lesson

	return &Subscription{
		ID: m.ID,
	}, nil
}

// Delete removes an checklist from the database.
func (repo *Repository) Delete(ctx context.Context, claims auth.Claims, req DeleteRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.subscription.Delete")
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

	_, err = models.Subscriptions(models.SubscriptionWhere.ID.EQ(req.ID)).DeleteAll(ctx, repo.DbConn)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
