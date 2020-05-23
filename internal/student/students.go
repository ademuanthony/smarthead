package student

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"remoteschool/smarthead/internal/platform/auth"
	"remoteschool/smarthead/internal/platform/web/webcontext"
	"remoteschool/smarthead/internal/postgres/models"

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

// Find gets all the students from the database based on the request params.
func (repo *Repository) Find(ctx context.Context, _ auth.Claims, req FindRequest) (Students, error) {
	var queries []QueryMod

	if req.IncludeClass {
		queries = append(queries, qm.Load(models.StudentRels.Class))
	}
	if req.IncludeSubclass {
		queries = append(queries, qm.Load(models.StudentRels.Subclass))
	}
	
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

	studentSlice, err := models.Students(queries...).All(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	var result Students
	for _, rec := range studentSlice {
		result = append(result, FromModel(rec))
	}

	return result, nil
}

// ReadByID gets the specified student by ID from the database.
func (repo *Repository) ReadByID(ctx context.Context, claims auth.Claims, id string) (*Student, error) {
	studentModel, err := models.FindStudent(ctx, repo.DbConn, id)
	if err != nil {
		return nil, err
	}

	return FromModel(studentModel), nil
}

// CurrentStudent gets the currently logged in student from the database
func (repo *Repository) CurrentStudent(ctx context.Context, claims auth.Claims) (*Student, error) {
	user, err := models.FindUser(ctx, repo.DbConn, claims.Subject)
	if err != nil {
		return nil, err
	}
	studentModel, err := models.Students(
		qm.Load(models.StudentRels.Subclass),
		models.StudentWhere.Username.EQ(user.Email),
	).One(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	return FromModel(studentModel), nil
}

// Create inserts a new student into the database.
func (repo *Repository) Create(ctx context.Context, req CreateRequest, now time.Time) (*Student, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.student.Create")
	defer span.Finish()

	// Validate the request.
	v := webcontext.Validator()
	err := v.StructCtx(ctx, req)
	if err != nil {
		return nil, err
	}

	now = now.UTC().Truncate(time.Millisecond)
	regNo, err := repo.generateRegNo(ctx)

	m := models.Student{
		ID:          uuid.NewRandom().String(),
		Name:        req.Name,
		Age:         req.Age,
		ClassID:     null.StringFrom(req.ClassID),
		SubclassID: null.StringFrom(req.SubclassID),
		ParentEmail: req.ParentEmail,
		ParentPhone: req.ParentPhone,
		RegNo: 		 regNo,
		Username:    req.Username,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := m.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return nil, errors.WithMessage(err, "Insert subject failed")
	}

	return &Student{
		ID:          m.ID,
		Name:        m.Name,
		Age:         m.Age,
		ClassID:     m.ClassID.String,
		SubclassID:  &req.SubclassID,
		ParentEmail: m.ParentEmail,
		ParentPhone: m.ParentPhone,
		Username:    m.Username,
		RegNo: 		 regNo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (repo *Repository) generateRegNo(ctx context.Context) (string, error) {
	var regNo string
	regNoExists := func (regNo string) (bool) {
		e, _ := models.Students(models.StudentWhere.RegNo.EQ(regNo)).Exists(ctx, repo.DbConn)
		return e
	}
	for regNo == "" || regNoExists(regNo) {
		regNo = "RS"
		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 6; i++ {
			regNo += strconv.Itoa(rand.Intn(10))
		}
	}
	return regNo, nil
}

// Update replaces an student in the database.
func (repo *Repository) Update(ctx context.Context, claims auth.Claims, req UpdateRequest, now time.Time) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.student.Update")
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

	now = now.UTC().Truncate(time.Millisecond)

	cols := models.M{}
	cols[models.StudentColumns.UpdatedAt] = now
	if req.Name != nil {
		cols[models.StudentColumns.Name] = *req.Name
	}
	if req.Age != nil {
		cols[models.StudentColumns.Age] = *req.Age
	}
	if req.ClassID != nil {
		cols[models.StudentColumns.ClassID] = *req.ClassID
	}
	if req.SubclassID != nil {
		cols[models.StudentColumns.SubclassID] = *req.SubclassID
	}
	if req.ParentPhone != nil {
		cols[models.StudentColumns.ParentPhone] = *req.ParentPhone
	}
	if req.ParentEmail != nil {
		cols[models.StudentColumns.ParentEmail] = *req.ParentEmail
	}

	if len(cols) == 0 {
		return nil
	}

	_, err = models.Students(models.StudentWhere.ID.EQ(req.ID)).UpdateAll(ctx, repo.DbConn, cols)

	return nil
}

// Delete removes an student from the database.
func (repo *Repository) Delete(ctx context.Context, claims auth.Claims, req DeleteRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.student.Delete")
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

	_, err = models.Students(models.StudentWhere.ID.EQ(req.ID)).DeleteAll(ctx, repo.DbConn)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
