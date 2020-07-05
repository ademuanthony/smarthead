package subject

import (
	"context"
	"strconv"
	"strings"

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

// Find gets all the subjects from the database based on the request params.
func (repo *Repository) Find(ctx context.Context, _ auth.Claims, req FindRequest) (Subjects, error) {
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

	subjectSlice, err := models.Subjects(queries...).All(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	var result Subjects
	for _, rec := range subjectSlice {
		result = append(result, FromModel(rec))
	}

	return result, nil
}

// ReadByID gets the specified subject by ID from the database.
func (repo *Repository) ReadByID(ctx context.Context, claims auth.Claims, id string) (*Subject, error) {
	subjectModel, err := models.FindSubject(ctx, repo.DbConn, id)
	if err != nil {
		return nil, err
	}

	return FromModel(subjectModel), nil
}

func (repo *Repository) EnglishID(ctx context.Context) (*Subject, error) {
	m, err := models.Subjects(
		models.SubjectWhere.Name.EQ("English Language"),
	).One(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	return FromModel(m), nil
}

func (repo *Repository) MathsID(ctx context.Context) (*Subject, error) {
	m, err := models.Subjects(
		models.SubjectWhere.Name.EQ("Mathematics"),
	).One(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	return FromModel(m), nil
}

// Create inserts a new subject into the database.
func (repo *Repository) Create(ctx context.Context, claims auth.Claims, req CreateRequest) (*Subject, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.subject.Create")
	defer span.Finish()
	if claims.Audience == "" {
		return nil, errors.WithStack(ErrForbidden)
	}

	// Admin users can update branch they have access to.
	if !claims.HasRole(auth.RoleAdmin) {
		return nil, errors.WithStack(ErrForbidden)
	}

	exists, err := models.Subjects(models.SubjectWhere.Name.EQ(req.Name)).Exists(ctx, repo.DbConn)
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

	schoolOrdersStrs := strings.Split(req.SchoolOrder, ",")
	var schoolOrders = make([]int64, len(schoolOrdersStrs))
	for i, s := range schoolOrdersStrs {
		s = strings.Trim(s, " ")
		sInt, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("Invalid character for school order")
		}
		schoolOrders[i] = int64(sInt)
	}
	m := models.Subject{
		ID:          uuid.NewRandom().String(),
		Name:        req.Name,
		SchoolOrder: schoolOrders,
	}

	if err := m.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return nil, errors.WithMessage(err, "Insert subject failed")
	}

	return &Subject{
		ID:           m.ID,
		Name:         m.Name,
		SchoolOrders: m.SchoolOrder,
	}, nil
}

// Update replaces an subject in the database.
func (repo *Repository) Update(ctx context.Context, claims auth.Claims, req UpdateRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.subject.Update")
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
		exists, err := models.Subjects(models.SubjectWhere.Name.EQ(*req.Name), models.SubjectWhere.ID.NEQ(req.ID)).Exists(ctx, repo.DbConn)
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
		cols[models.SubjectColumns.Name] = *req.Name
	}

	if req.SchoolOrder != nil {
		schoolOrdersStrs := strings.Split(*req.SchoolOrder, ",")
		var schoolOrders = make([]int64, len(schoolOrdersStrs))
		for i, s := range schoolOrdersStrs {
			s = strings.Trim(s, " ")
			sInt, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("Invalid character for school order")
			}
			schoolOrders[i] = int64(sInt)
		}
		cols[models.SubjectColumns.SchoolOrder] = schoolOrders
	}

	if len(cols) == 0 {
		return nil
	}

	_, err = models.Subjects(models.SubjectWhere.ID.EQ(req.ID)).UpdateAll(ctx, repo.DbConn, cols)

	return nil
}

// Delete removes an subject from the database.
func (repo *Repository) Delete(ctx context.Context, claims auth.Claims, req DeleteRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.subject.Delete")
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

	_, err = models.Subjects(models.SubjectWhere.ID.EQ(req.ID)).DeleteAll(ctx, repo.DbConn)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
