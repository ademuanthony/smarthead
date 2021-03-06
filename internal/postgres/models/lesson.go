// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Lesson is an object representing the database table.
type Lesson struct {
	ID                    string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	TimetableID           string      `boil:"timetable_id" json:"timetable_id" toml:"timetable_id" yaml:"timetable_id"`
	Date                  int64       `boil:"date" json:"date" toml:"date" yaml:"date"`
	StartDate             int64       `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	EndDate               int64       `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`
	TeacherID             null.String `boil:"teacher_id" json:"teacher_id,omitempty" toml:"teacher_id" yaml:"teacher_id,omitempty"`
	TeacherAttendanceDate int64       `boil:"teacher_attendance_date" json:"teacher_attendance_date" toml:"teacher_attendance_date" yaml:"teacher_attendance_date"`

	R *lessonR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lessonL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LessonColumns = struct {
	ID                    string
	TimetableID           string
	Date                  string
	StartDate             string
	EndDate               string
	TeacherID             string
	TeacherAttendanceDate string
}{
	ID:                    "id",
	TimetableID:           "timetable_id",
	Date:                  "date",
	StartDate:             "start_date",
	EndDate:               "end_date",
	TeacherID:             "teacher_id",
	TeacherAttendanceDate: "teacher_attendance_date",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

var LessonWhere = struct {
	ID                    whereHelperstring
	TimetableID           whereHelperstring
	Date                  whereHelperint64
	StartDate             whereHelperint64
	EndDate               whereHelperint64
	TeacherID             whereHelpernull_String
	TeacherAttendanceDate whereHelperint64
}{
	ID:                    whereHelperstring{field: "\"lesson\".\"id\""},
	TimetableID:           whereHelperstring{field: "\"lesson\".\"timetable_id\""},
	Date:                  whereHelperint64{field: "\"lesson\".\"date\""},
	StartDate:             whereHelperint64{field: "\"lesson\".\"start_date\""},
	EndDate:               whereHelperint64{field: "\"lesson\".\"end_date\""},
	TeacherID:             whereHelpernull_String{field: "\"lesson\".\"teacher_id\""},
	TeacherAttendanceDate: whereHelperint64{field: "\"lesson\".\"teacher_attendance_date\""},
}

// LessonRels is where relationship names are stored.
var LessonRels = struct {
	Teacher        string
	Timetable      string
	LessonStudents string
}{
	Teacher:        "Teacher",
	Timetable:      "Timetable",
	LessonStudents: "LessonStudents",
}

// lessonR is where relationships are stored.
type lessonR struct {
	Teacher        *User
	Timetable      *Timetable
	LessonStudents LessonStudentSlice
}

// NewStruct creates a new relationship struct
func (*lessonR) NewStruct() *lessonR {
	return &lessonR{}
}

// lessonL is where Load methods for each relationship are stored.
type lessonL struct{}

var (
	lessonAllColumns            = []string{"id", "timetable_id", "date", "start_date", "end_date", "teacher_id", "teacher_attendance_date"}
	lessonColumnsWithoutDefault = []string{"id", "timetable_id", "date", "start_date", "end_date"}
	lessonColumnsWithDefault    = []string{"teacher_id", "teacher_attendance_date"}
	lessonPrimaryKeyColumns     = []string{"id"}
)

type (
	// LessonSlice is an alias for a slice of pointers to Lesson.
	// This should generally be used opposed to []Lesson.
	LessonSlice []*Lesson

	lessonQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lessonType                 = reflect.TypeOf(&Lesson{})
	lessonMapping              = queries.MakeStructMapping(lessonType)
	lessonPrimaryKeyMapping, _ = queries.BindMapping(lessonType, lessonMapping, lessonPrimaryKeyColumns)
	lessonInsertCacheMut       sync.RWMutex
	lessonInsertCache          = make(map[string]insertCache)
	lessonUpdateCacheMut       sync.RWMutex
	lessonUpdateCache          = make(map[string]updateCache)
	lessonUpsertCacheMut       sync.RWMutex
	lessonUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single lesson record from the query.
func (q lessonQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Lesson, error) {
	o := &Lesson{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for lesson")
	}

	return o, nil
}

// All returns all Lesson records from the query.
func (q lessonQuery) All(ctx context.Context, exec boil.ContextExecutor) (LessonSlice, error) {
	var o []*Lesson

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Lesson slice")
	}

	return o, nil
}

// Count returns the count of all Lesson records in the query.
func (q lessonQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count lesson rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lessonQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if lesson exists")
	}

	return count > 0, nil
}

// Teacher pointed to by the foreign key.
func (o *Lesson) Teacher(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TeacherID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Timetable pointed to by the foreign key.
func (o *Lesson) Timetable(mods ...qm.QueryMod) timetableQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TimetableID),
	}

	queryMods = append(queryMods, mods...)

	query := Timetables(queryMods...)
	queries.SetFrom(query.Query, "\"timetable\"")

	return query
}

// LessonStudents retrieves all the lesson_student's LessonStudents with an executor.
func (o *Lesson) LessonStudents(mods ...qm.QueryMod) lessonStudentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"lesson_student\".\"lesson_id\"=?", o.ID),
	)

	query := LessonStudents(queryMods...)
	queries.SetFrom(query.Query, "\"lesson_student\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"lesson_student\".*"})
	}

	return query
}

// LoadTeacher allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lessonL) LoadTeacher(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLesson interface{}, mods queries.Applicator) error {
	var slice []*Lesson
	var object *Lesson

	if singular {
		object = maybeLesson.(*Lesson)
	} else {
		slice = *maybeLesson.(*[]*Lesson)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lessonR{}
		}
		if !queries.IsNil(object.TeacherID) {
			args = append(args, object.TeacherID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lessonR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.TeacherID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.TeacherID) {
				args = append(args, obj.TeacherID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`users.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Teacher = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.TeacherLessons = append(foreign.R.TeacherLessons, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.TeacherID, foreign.ID) {
				local.R.Teacher = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.TeacherLessons = append(foreign.R.TeacherLessons, local)
				break
			}
		}
	}

	return nil
}

// LoadTimetable allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lessonL) LoadTimetable(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLesson interface{}, mods queries.Applicator) error {
	var slice []*Lesson
	var object *Lesson

	if singular {
		object = maybeLesson.(*Lesson)
	} else {
		slice = *maybeLesson.(*[]*Lesson)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lessonR{}
		}
		args = append(args, object.TimetableID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lessonR{}
			}

			for _, a := range args {
				if a == obj.TimetableID {
					continue Outer
				}
			}

			args = append(args, obj.TimetableID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`timetable`), qm.WhereIn(`timetable.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Timetable")
	}

	var resultSlice []*Timetable
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Timetable")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for timetable")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for timetable")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Timetable = foreign
		if foreign.R == nil {
			foreign.R = &timetableR{}
		}
		foreign.R.Lessons = append(foreign.R.Lessons, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TimetableID == foreign.ID {
				local.R.Timetable = foreign
				if foreign.R == nil {
					foreign.R = &timetableR{}
				}
				foreign.R.Lessons = append(foreign.R.Lessons, local)
				break
			}
		}
	}

	return nil
}

// LoadLessonStudents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (lessonL) LoadLessonStudents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLesson interface{}, mods queries.Applicator) error {
	var slice []*Lesson
	var object *Lesson

	if singular {
		object = maybeLesson.(*Lesson)
	} else {
		slice = *maybeLesson.(*[]*Lesson)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lessonR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lessonR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`lesson_student`), qm.WhereIn(`lesson_student.lesson_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load lesson_student")
	}

	var resultSlice []*LessonStudent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice lesson_student")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on lesson_student")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for lesson_student")
	}

	if singular {
		object.R.LessonStudents = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &lessonStudentR{}
			}
			foreign.R.Lesson = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.LessonID {
				local.R.LessonStudents = append(local.R.LessonStudents, foreign)
				if foreign.R == nil {
					foreign.R = &lessonStudentR{}
				}
				foreign.R.Lesson = local
				break
			}
		}
	}

	return nil
}

// SetTeacher of the lesson to the related item.
// Sets o.R.Teacher to related.
// Adds o to related.R.TeacherLessons.
func (o *Lesson) SetTeacher(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lesson\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"teacher_id"}),
		strmangle.WhereClause("\"", "\"", 2, lessonPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.TeacherID, related.ID)
	if o.R == nil {
		o.R = &lessonR{
			Teacher: related,
		}
	} else {
		o.R.Teacher = related
	}

	if related.R == nil {
		related.R = &userR{
			TeacherLessons: LessonSlice{o},
		}
	} else {
		related.R.TeacherLessons = append(related.R.TeacherLessons, o)
	}

	return nil
}

// RemoveTeacher relationship.
// Sets o.R.Teacher to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Lesson) RemoveTeacher(ctx context.Context, exec boil.ContextExecutor, related *User) error {
	var err error

	queries.SetScanner(&o.TeacherID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("teacher_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Teacher = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.TeacherLessons {
		if queries.Equal(o.TeacherID, ri.TeacherID) {
			continue
		}

		ln := len(related.R.TeacherLessons)
		if ln > 1 && i < ln-1 {
			related.R.TeacherLessons[i] = related.R.TeacherLessons[ln-1]
		}
		related.R.TeacherLessons = related.R.TeacherLessons[:ln-1]
		break
	}
	return nil
}

// SetTimetable of the lesson to the related item.
// Sets o.R.Timetable to related.
// Adds o to related.R.Lessons.
func (o *Lesson) SetTimetable(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Timetable) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lesson\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"timetable_id"}),
		strmangle.WhereClause("\"", "\"", 2, lessonPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TimetableID = related.ID
	if o.R == nil {
		o.R = &lessonR{
			Timetable: related,
		}
	} else {
		o.R.Timetable = related
	}

	if related.R == nil {
		related.R = &timetableR{
			Lessons: LessonSlice{o},
		}
	} else {
		related.R.Lessons = append(related.R.Lessons, o)
	}

	return nil
}

// AddLessonStudents adds the given related objects to the existing relationships
// of the lesson, optionally inserting them as new records.
// Appends related to o.R.LessonStudents.
// Sets related.R.Lesson appropriately.
func (o *Lesson) AddLessonStudents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*LessonStudent) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.LessonID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"lesson_student\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"lesson_id"}),
				strmangle.WhereClause("\"", "\"", 2, lessonStudentPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.LessonID = o.ID
		}
	}

	if o.R == nil {
		o.R = &lessonR{
			LessonStudents: related,
		}
	} else {
		o.R.LessonStudents = append(o.R.LessonStudents, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &lessonStudentR{
				Lesson: o,
			}
		} else {
			rel.R.Lesson = o
		}
	}
	return nil
}

// Lessons retrieves all the records using an executor.
func Lessons(mods ...qm.QueryMod) lessonQuery {
	mods = append(mods, qm.From("\"lesson\""))
	return lessonQuery{NewQuery(mods...)}
}

// FindLesson retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLesson(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Lesson, error) {
	lessonObj := &Lesson{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"lesson\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, lessonObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from lesson")
	}

	return lessonObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Lesson) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no lesson provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(lessonColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lessonInsertCacheMut.RLock()
	cache, cached := lessonInsertCache[key]
	lessonInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lessonAllColumns,
			lessonColumnsWithDefault,
			lessonColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lessonType, lessonMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lessonType, lessonMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"lesson\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"lesson\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into lesson")
	}

	if !cached {
		lessonInsertCacheMut.Lock()
		lessonInsertCache[key] = cache
		lessonInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Lesson.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Lesson) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	lessonUpdateCacheMut.RLock()
	cache, cached := lessonUpdateCache[key]
	lessonUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lessonAllColumns,
			lessonPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update lesson, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"lesson\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, lessonPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lessonType, lessonMapping, append(wl, lessonPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update lesson row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for lesson")
	}

	if !cached {
		lessonUpdateCacheMut.Lock()
		lessonUpdateCache[key] = cache
		lessonUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q lessonQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for lesson")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for lesson")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LessonSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lessonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"lesson\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, lessonPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in lesson slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all lesson")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Lesson) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no lesson provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(lessonColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	lessonUpsertCacheMut.RLock()
	cache, cached := lessonUpsertCache[key]
	lessonUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lessonAllColumns,
			lessonColumnsWithDefault,
			lessonColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			lessonAllColumns,
			lessonPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert lesson, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(lessonPrimaryKeyColumns))
			copy(conflict, lessonPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"lesson\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(lessonType, lessonMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lessonType, lessonMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert lesson")
	}

	if !cached {
		lessonUpsertCacheMut.Lock()
		lessonUpsertCache[key] = cache
		lessonUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Lesson record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Lesson) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Lesson provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lessonPrimaryKeyMapping)
	sql := "DELETE FROM \"lesson\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from lesson")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for lesson")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lessonQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no lessonQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from lesson")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for lesson")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LessonSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lessonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"lesson\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lessonPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from lesson slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for lesson")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Lesson) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLesson(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LessonSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LessonSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lessonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"lesson\".* FROM \"lesson\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lessonPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LessonSlice")
	}

	*o = slice

	return nil
}

// LessonExists checks if the Lesson row exists.
func LessonExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"lesson\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if lesson exists")
	}

	return exists, nil
}
