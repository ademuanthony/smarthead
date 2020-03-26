package period

import (
	"context"
	"time"

	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/postgres/models"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var (
	// ErrNotFound abstracts the postgres not found error.
	ErrNotFound = errors.New("Entity not found")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("Attempted action is not allowed")
)

// Find gets all the period from the database based on the request params.
func (repo *Repository) Find(ctx context.Context, _ auth.Claims, req FindRequest) (Periods, error) {
	var queries []QueryMod

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

	periodSlice, err := models.Periods(queries...).All(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	var result Periods
	for _, rec := range periodSlice {
		result = append(result, FromModel(rec))
	}

	return result, nil
}

// ReadByID gets the specified period by ID from the database.
func (repo *Repository) ReadByID(ctx context.Context, claims auth.Claims, id string) (*Period, error) {
	periodModel, err := models.FindPeriod(ctx, repo.DbConn, id)
	if err != nil {
		return nil, err
	}

	return FromModel(periodModel), nil
}

// Create inserts a new period into the database.
func (repo *Repository) Create(ctx context.Context, claims auth.Claims, req CreateRequest, now time.Time) (*Period, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.period.Create")
	defer span.Finish()
	if claims.Audience == "" {
		return nil, errors.WithStack(ErrForbidden)
	}

	// Admin users can update branch they have access to.
	if !claims.HasRole(auth.RoleAdmin) {
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

	m := models.Period{
		ID:        uuid.NewRandom().String(),
		StartHour: req.StartHour,
		StartMinute: req.StartMinute,
		EndHour: req.EndHour,
		EndMinute: req.EndMinute,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := m.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return nil, errors.WithMessage(err, "Insert period failed")
	}

	return FromModel(&m), nil
}

// Update replaces an period in the database.
func (repo *Repository) Update(ctx context.Context, claims auth.Claims, req UpdateRequest, now time.Time) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.period.Update")
	defer span.Finish()

	if claims.Audience == "" {
		return errors.WithStack(ErrForbidden)
	}
	// Admin users can update branches they have access to.
	if !claims.HasRole(auth.RoleAdmin) {
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
	if req.StartHour != nil {
		cols[models.PeriodColumns.StartHour] = *req.StartHour
	}

	if req.StartMinute != nil {
		cols[models.PeriodColumns.StartMinute] = *req.StartMinute
	}

	if req.EndHour != nil {
		cols[models.PeriodColumns.EndHour] = *req.EndHour
	}

	if req.EndMinute != nil {
		cols[models.PeriodColumns.EndMinute] = *req.EndMinute
	}

	cols[models.PeriodColumns.UpdatedAt] = now

	_,err = models.Periods(models.PeriodWhere.ID.EQ(req.ID)).UpdateAll(ctx, repo.DbConn, cols)

	return nil
}

// Delete removes an period from the database.
func (repo *Repository) Delete(ctx context.Context, claims auth.Claims, req DeleteRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.period.Delete")
	defer span.Finish()

	// Validate the request.
	v := webcontext.Validator()
	err := v.Struct(req)
	if err != nil {
		return err
	}

	// Ensure the claims can modify the project specified in the request.
	if claims.Audience == "" {
		return errors.WithStack(ErrForbidden)
	}
	// Admin users can update Categories they have access to.
	if !claims.HasRole(auth.RoleAdmin) {
		return errors.WithStack(ErrForbidden)
	}

	_, err = models.Periods(models.PeriodWhere.ID.EQ(req.ID)).DeleteAll(ctx, repo.DbConn)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
