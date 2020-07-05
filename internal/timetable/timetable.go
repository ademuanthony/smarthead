package timetable

import (
	"context"
	"database/sql"
	"fmt"
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

// Find gets all the timtables from the database based on the request params.
func (repo *Repository) Find(ctx context.Context, req FindRequest) (Timetables, error) {
	var queries []QueryMod

	if req.IncludeSubclass {
		queries = append(queries, qm.Load(models.TimetableRels.Subclass))
	}

	if req.IncludeSubject {
		queries = append(queries, qm.Load(models.TimetableRels.Subject))
	}

	if req.IncludePeriod {
		queries = append(queries, qm.Load(models.TimetableRels.Period))
	}

	if req.IncludeTeacher {
		queries = append(queries, qm.Load(models.TimetableRels.Teacher))
	}

	if req.Where != "" {
		queries = append(queries, Where(req.Where, req.Args...))
	}

	req.Order = []string{"day"}

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

	timetableSlice, err := models.Timetables(queries...).All(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	var result Timetables
	for _, rec := range timetableSlice {
		result = append(result, FromModel(rec))
	}

	return result, nil
}

func (repo *Repository) StudentsTimetables(ctx context.Context, studentID string, now time.Time) (Timetables, error) {
	student, err := models.FindStudent(ctx, repo.DbConn, studentID)
	if err != nil {
		return nil, fmt.Errorf("Cannot fetch student, %s", err)
	}

	if !student.SubclassID.Valid {
		return Timetables{}, nil
	}

	thirtyDaysAgo := now.Add(-30 * 24 * time.Hour).Unix()
	if student.LastPaymentDate < thirtyDaysAgo {
		return nil, errors.New("payment required")
	}

	// subscriptions, err := models.Subscriptions(
	// 	models.SubscriptionWhere.StudentID.EQ(student.ID),
	// 	models.SubscriptionWhere.StartDate.LTE(time.Now().UTC().Unix()),
	// 	models.SubscriptionWhere.EndDate.GTE(time.Now().UTC().Unix()),
	// ).All(ctx, repo.DbConn)
	// if err != nil {
	// 	if err.Error() == sql.ErrNoRows.Error() {
	// 		return Timetables{}, nil
	// 	}
	// 	return nil, err
	// }

	// var subjectIDs string
	// for _, s := range subscriptions {
	// 	subjectIDs += s.SubjectID + "|"
	// }

	timetables, err := models.Timetables(
		models.TimetableWhere.SubclassID.EQ(student.SubclassID.String),
		qm.Load(models.TimetableRels.Period),
		qm.Load(models.TimetableRels.Subclass),
		qm.Load(models.TimetableRels.Subject),
		qm.Load(models.TimetableRels.Teacher),
	).All(ctx, repo.DbConn)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return Timetables{}, nil
		}
		return nil, err
	}

	var result Timetables
	for _, t := range timetables {
		result = append(result, FromModel(t))
		// if strings.Contains(subjectIDs, t.SubjectID) {
		// 	result = append(result, FromModel(t))
		// }
	}

	return result, nil
}

func (repo *Repository) TeachersTimetables(ctx context.Context, teacherID string) (Timetables, error) {
	
	timetables, err := models.Timetables(
		models.TimetableWhere.TeacherID.EQ(teacherID),
		qm.Load(models.TimetableRels.Period),
		qm.Load(models.TimetableRels.Subclass),
		qm.Load(models.TimetableRels.Subject),
		qm.Load(models.TimetableRels.Teacher),
	).All(ctx, repo.DbConn)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return Timetables{}, nil
		}
		return nil, err
	}

	var result Timetables
	for _, t := range timetables {
		result = append(result, FromModel(t))
	}

	return result, nil
}

// ReadByID gets the specified class by ID from the database.
func (repo *Repository) ReadByID(ctx context.Context, claims auth.Claims, id string) (*Timetable, error) {
	classModel, err := models.Timetables(
		models.TimetableWhere.ID.EQ(id),
		qm.Load(models.TimetableRels.Subclass),
	).One(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}

	return FromModel(classModel), nil
}

// Create inserts a new class into the database.
func (repo *Repository) Create(ctx context.Context, claims auth.Claims, req CreateRequest) (*Timetable, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.timetable.Create")
	defer span.Finish()
	if claims.Audience == "" {
		return nil, errors.WithStack(ErrForbidden)
	}

	// Admin users can update branch they have access to.
	if !claims.HasRole(auth.RoleAdmin) {
		return nil, errors.WithStack(ErrForbidden)
	}

	exists, err := models.Timetables(
		models.TimetableWhere.SubclassID.EQ(req.SubclassID),
		models.TimetableWhere.PeriodID.EQ(req.PeriodID),
		models.TimetableWhere.Day.EQ(req.Day),
	).Exists(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("The selected period is already occupied for this class. Please use the update action")
	}

	exists, err = models.Timetables(
		models.TimetableWhere.TeacherID.EQ(req.TeacherID),
		models.TimetableWhere.PeriodID.EQ(req.PeriodID),
		models.TimetableWhere.Day.EQ(req.Day),
	).Exists(ctx, repo.DbConn)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("The selected teacher is already assigned a class for this period.")
	}

	// Validate the request.
	v := webcontext.Validator()
	err = v.StructCtx(ctx, req)
	if err != nil {
		return nil, err
	}

	m := models.Timetable{
		ID:         uuid.NewRandom().String(),
		Day:        req.Day,
		PeriodID:   req.PeriodID,
		SubclassID: req.SubclassID,
		SubjectID:  req.SubjectID,
		TeacherID:  req.TeacherID,
	}

	if err := m.Insert(ctx, repo.DbConn, boil.Infer()); err != nil {
		return nil, errors.WithMessage(err, "Insert timetable failed")
	}

	return &Timetable{
		ID:         m.ID,
		Day:        req.Day,
		PeriodID:   req.PeriodID,
		SubclassID: req.SubclassID,
		SubjectID:  req.SubjectID,
		TeacherID:  req.TeacherID,
	}, nil
}

// Delete removes an class from the database.
func (repo *Repository) Delete(ctx context.Context, claims auth.Claims, req DeleteRequest) error {
	span, ctx := tracer.StartSpanFromContext(ctx, "internal.timetable.Delete")
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

	_, err = models.Timetables(models.TimetableWhere.ID.EQ(req.ID)).DeleteAll(ctx, repo.DbConn)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
