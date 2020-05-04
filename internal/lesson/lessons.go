package lesson

import (
	"context"

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

// Find gets all the classes from the database based on the request params.
func (repo *Repository) Find(ctx context.Context, req FindRequest) (Lessons, error) {
	var queries []QueryMod

	if req.Where != "" {
		queries = append(queries, Where(req.Where, req.Args...))
	}

	req.Order = []string{"school_order", "name"}

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

	classSlice, err := models.Classes(queries...).All(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	var result Lessons
	for _, rec := range classSlice {
		result = append(result, FromModel(rec))
	}

	return result, nil
}

// ReadByID gets the specified class by ID from the database.
func (repo *Repository) ReadByID(ctx context.Context, claims auth.Claims, id string) (*Lesson, error) {
	classModel, err := models.FindClass(ctx, repo.DbConn, id)
	if err != nil {
		return nil, err
	}

	return FromModel(classModel), nil
}

// Create inserts a new class into the database.
func (repo *Repository) Create(ctx context.Context, claims auth.Claims, req CreateRequest) (*Lesson, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.class.Create")
	defer span.Finish()
	if claims.Audience == "" {
		return nil, errors.WithStack(ErrForbidden)
	}

	// Admin users can update branch they have access to.
	if !claims.HasRole(auth.RoleAdmin) {
		return nil, errors.WithStack(ErrForbidden)
	}

	exists, err := models.Classes(models.ClassWhere.Name.EQ(req.Name)).Exists(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}
	ctx = webcontext.ContextAddUniqueValue(ctx, req, "Name", !exists)

	// Validate the request.
	v := webcontext.Validator()
	err = v.StructCtx(ctx, req)
	if err != nil {
		return nil, err
	}

	m := models.Class {
		ID:        uuid.NewRandom().String(),
		Name:      req.Name,
		SchoolOrder: req.SchoolOrder,
	}

	if err := m.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return nil, errors.WithMessage(err, "Insert class failed")
	}

	return &Lesson{
		ID:         m.ID,
		Name:       m.Name,
		SchoolOrder: m.SchoolOrder,
	}, nil
}

// Update replaces an class in the database.
func (repo *Repository) Update(ctx context.Context, claims auth.Claims, req UpdateRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.class.Update")
	defer span.Finish()

	if claims.Audience == "" {
		return errors.WithStack(ErrForbidden)
	}
	// Admin users can update branches they have access to.
	if !claims.HasRole(auth.RoleAdmin) {
		return errors.WithStack(ErrForbidden)
	}

	var unique = true
	if req.Name != nil {
		exists, err := models.Classes(models.ClassWhere.Name.EQ(*req.Name), models.ClassWhere.ID.NEQ(req.ID)).Exists(ctx, repo.DbConn)
		if err != nil {
			return err
		}
		unique = !exists
	}

	ctx = webcontext.ContextAddUniqueValue(ctx, req, "Name", unique)

	// Validate the request.
	v := webcontext.Validator()
	err := v.StructCtx(ctx, req)
	if err != nil {
		return err
	}

	cols := models.M{}
	if req.Name != nil {
		cols[models.ClassColumns.Name] = *req.Name
	}

	if req.SchoolOrder != nil {
		cols[models.ClassColumns.SchoolOrder] = *req.SchoolOrder
	}

	if len(cols) == 0 {
		return nil
	}

	_,err = models.Classes(models.ClassWhere.ID.EQ(req.ID)).UpdateAll(ctx, repo.DbConn, cols)

	return nil
}

// Delete removes an class from the database.
func (repo *Repository) Delete(ctx context.Context, claims auth.Claims, req DeleteRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.class.Delete")
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

	_, err = models.Classes(models.ClassWhere.ID.EQ(req.ID)).DeleteAll(ctx, repo.DbConn)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
