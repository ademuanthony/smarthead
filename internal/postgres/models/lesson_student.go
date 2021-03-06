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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// LessonStudent is an object representing the database table.
type LessonStudent struct {
	ID             string `boil:"id" json:"id" toml:"id" yaml:"id"`
	LessonID       string `boil:"lesson_id" json:"lesson_id" toml:"lesson_id" yaml:"lesson_id"`
	StudentID      string `boil:"student_id" json:"student_id" toml:"student_id" yaml:"student_id"`
	AttendanceDate int64  `boil:"attendance_date" json:"attendance_date" toml:"attendance_date" yaml:"attendance_date"`

	R *lessonStudentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lessonStudentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LessonStudentColumns = struct {
	ID             string
	LessonID       string
	StudentID      string
	AttendanceDate string
}{
	ID:             "id",
	LessonID:       "lesson_id",
	StudentID:      "student_id",
	AttendanceDate: "attendance_date",
}

// Generated where

var LessonStudentWhere = struct {
	ID             whereHelperstring
	LessonID       whereHelperstring
	StudentID      whereHelperstring
	AttendanceDate whereHelperint64
}{
	ID:             whereHelperstring{field: "\"lesson_student\".\"id\""},
	LessonID:       whereHelperstring{field: "\"lesson_student\".\"lesson_id\""},
	StudentID:      whereHelperstring{field: "\"lesson_student\".\"student_id\""},
	AttendanceDate: whereHelperint64{field: "\"lesson_student\".\"attendance_date\""},
}

// LessonStudentRels is where relationship names are stored.
var LessonStudentRels = struct {
	Lesson  string
	Student string
}{
	Lesson:  "Lesson",
	Student: "Student",
}

// lessonStudentR is where relationships are stored.
type lessonStudentR struct {
	Lesson  *Lesson
	Student *Student
}

// NewStruct creates a new relationship struct
func (*lessonStudentR) NewStruct() *lessonStudentR {
	return &lessonStudentR{}
}

// lessonStudentL is where Load methods for each relationship are stored.
type lessonStudentL struct{}

var (
	lessonStudentAllColumns            = []string{"id", "lesson_id", "student_id", "attendance_date"}
	lessonStudentColumnsWithoutDefault = []string{"id", "lesson_id", "student_id"}
	lessonStudentColumnsWithDefault    = []string{"attendance_date"}
	lessonStudentPrimaryKeyColumns     = []string{"id"}
)

type (
	// LessonStudentSlice is an alias for a slice of pointers to LessonStudent.
	// This should generally be used opposed to []LessonStudent.
	LessonStudentSlice []*LessonStudent

	lessonStudentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lessonStudentType                 = reflect.TypeOf(&LessonStudent{})
	lessonStudentMapping              = queries.MakeStructMapping(lessonStudentType)
	lessonStudentPrimaryKeyMapping, _ = queries.BindMapping(lessonStudentType, lessonStudentMapping, lessonStudentPrimaryKeyColumns)
	lessonStudentInsertCacheMut       sync.RWMutex
	lessonStudentInsertCache          = make(map[string]insertCache)
	lessonStudentUpdateCacheMut       sync.RWMutex
	lessonStudentUpdateCache          = make(map[string]updateCache)
	lessonStudentUpsertCacheMut       sync.RWMutex
	lessonStudentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single lessonStudent record from the query.
func (q lessonStudentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LessonStudent, error) {
	o := &LessonStudent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for lesson_student")
	}

	return o, nil
}

// All returns all LessonStudent records from the query.
func (q lessonStudentQuery) All(ctx context.Context, exec boil.ContextExecutor) (LessonStudentSlice, error) {
	var o []*LessonStudent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LessonStudent slice")
	}

	return o, nil
}

// Count returns the count of all LessonStudent records in the query.
func (q lessonStudentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count lesson_student rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lessonStudentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if lesson_student exists")
	}

	return count > 0, nil
}

// Lesson pointed to by the foreign key.
func (o *LessonStudent) Lesson(mods ...qm.QueryMod) lessonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LessonID),
	}

	queryMods = append(queryMods, mods...)

	query := Lessons(queryMods...)
	queries.SetFrom(query.Query, "\"lesson\"")

	return query
}

// Student pointed to by the foreign key.
func (o *LessonStudent) Student(mods ...qm.QueryMod) studentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.StudentID),
	}

	queryMods = append(queryMods, mods...)

	query := Students(queryMods...)
	queries.SetFrom(query.Query, "\"student\"")

	return query
}

// LoadLesson allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lessonStudentL) LoadLesson(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLessonStudent interface{}, mods queries.Applicator) error {
	var slice []*LessonStudent
	var object *LessonStudent

	if singular {
		object = maybeLessonStudent.(*LessonStudent)
	} else {
		slice = *maybeLessonStudent.(*[]*LessonStudent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lessonStudentR{}
		}
		args = append(args, object.LessonID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lessonStudentR{}
			}

			for _, a := range args {
				if a == obj.LessonID {
					continue Outer
				}
			}

			args = append(args, obj.LessonID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`lesson`), qm.WhereIn(`lesson.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Lesson")
	}

	var resultSlice []*Lesson
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Lesson")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for lesson")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for lesson")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Lesson = foreign
		if foreign.R == nil {
			foreign.R = &lessonR{}
		}
		foreign.R.LessonStudents = append(foreign.R.LessonStudents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LessonID == foreign.ID {
				local.R.Lesson = foreign
				if foreign.R == nil {
					foreign.R = &lessonR{}
				}
				foreign.R.LessonStudents = append(foreign.R.LessonStudents, local)
				break
			}
		}
	}

	return nil
}

// LoadStudent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lessonStudentL) LoadStudent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLessonStudent interface{}, mods queries.Applicator) error {
	var slice []*LessonStudent
	var object *LessonStudent

	if singular {
		object = maybeLessonStudent.(*LessonStudent)
	} else {
		slice = *maybeLessonStudent.(*[]*LessonStudent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lessonStudentR{}
		}
		args = append(args, object.StudentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lessonStudentR{}
			}

			for _, a := range args {
				if a == obj.StudentID {
					continue Outer
				}
			}

			args = append(args, obj.StudentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`student`), qm.WhereIn(`student.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Student")
	}

	var resultSlice []*Student
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Student")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for student")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for student")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Student = foreign
		if foreign.R == nil {
			foreign.R = &studentR{}
		}
		foreign.R.LessonStudents = append(foreign.R.LessonStudents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StudentID == foreign.ID {
				local.R.Student = foreign
				if foreign.R == nil {
					foreign.R = &studentR{}
				}
				foreign.R.LessonStudents = append(foreign.R.LessonStudents, local)
				break
			}
		}
	}

	return nil
}

// SetLesson of the lessonStudent to the related item.
// Sets o.R.Lesson to related.
// Adds o to related.R.LessonStudents.
func (o *LessonStudent) SetLesson(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Lesson) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lesson_student\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"lesson_id"}),
		strmangle.WhereClause("\"", "\"", 2, lessonStudentPrimaryKeyColumns),
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

	o.LessonID = related.ID
	if o.R == nil {
		o.R = &lessonStudentR{
			Lesson: related,
		}
	} else {
		o.R.Lesson = related
	}

	if related.R == nil {
		related.R = &lessonR{
			LessonStudents: LessonStudentSlice{o},
		}
	} else {
		related.R.LessonStudents = append(related.R.LessonStudents, o)
	}

	return nil
}

// SetStudent of the lessonStudent to the related item.
// Sets o.R.Student to related.
// Adds o to related.R.LessonStudents.
func (o *LessonStudent) SetStudent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Student) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lesson_student\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"student_id"}),
		strmangle.WhereClause("\"", "\"", 2, lessonStudentPrimaryKeyColumns),
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

	o.StudentID = related.ID
	if o.R == nil {
		o.R = &lessonStudentR{
			Student: related,
		}
	} else {
		o.R.Student = related
	}

	if related.R == nil {
		related.R = &studentR{
			LessonStudents: LessonStudentSlice{o},
		}
	} else {
		related.R.LessonStudents = append(related.R.LessonStudents, o)
	}

	return nil
}

// LessonStudents retrieves all the records using an executor.
func LessonStudents(mods ...qm.QueryMod) lessonStudentQuery {
	mods = append(mods, qm.From("\"lesson_student\""))
	return lessonStudentQuery{NewQuery(mods...)}
}

// FindLessonStudent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLessonStudent(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*LessonStudent, error) {
	lessonStudentObj := &LessonStudent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"lesson_student\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, lessonStudentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from lesson_student")
	}

	return lessonStudentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LessonStudent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no lesson_student provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(lessonStudentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lessonStudentInsertCacheMut.RLock()
	cache, cached := lessonStudentInsertCache[key]
	lessonStudentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lessonStudentAllColumns,
			lessonStudentColumnsWithDefault,
			lessonStudentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lessonStudentType, lessonStudentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lessonStudentType, lessonStudentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"lesson_student\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"lesson_student\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into lesson_student")
	}

	if !cached {
		lessonStudentInsertCacheMut.Lock()
		lessonStudentInsertCache[key] = cache
		lessonStudentInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the LessonStudent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LessonStudent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	lessonStudentUpdateCacheMut.RLock()
	cache, cached := lessonStudentUpdateCache[key]
	lessonStudentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lessonStudentAllColumns,
			lessonStudentPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update lesson_student, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"lesson_student\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, lessonStudentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lessonStudentType, lessonStudentMapping, append(wl, lessonStudentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update lesson_student row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for lesson_student")
	}

	if !cached {
		lessonStudentUpdateCacheMut.Lock()
		lessonStudentUpdateCache[key] = cache
		lessonStudentUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q lessonStudentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for lesson_student")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for lesson_student")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LessonStudentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lessonStudentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"lesson_student\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, lessonStudentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in lessonStudent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all lessonStudent")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LessonStudent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no lesson_student provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(lessonStudentColumnsWithDefault, o)

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

	lessonStudentUpsertCacheMut.RLock()
	cache, cached := lessonStudentUpsertCache[key]
	lessonStudentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lessonStudentAllColumns,
			lessonStudentColumnsWithDefault,
			lessonStudentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			lessonStudentAllColumns,
			lessonStudentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert lesson_student, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(lessonStudentPrimaryKeyColumns))
			copy(conflict, lessonStudentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"lesson_student\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(lessonStudentType, lessonStudentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lessonStudentType, lessonStudentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert lesson_student")
	}

	if !cached {
		lessonStudentUpsertCacheMut.Lock()
		lessonStudentUpsertCache[key] = cache
		lessonStudentUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single LessonStudent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LessonStudent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no LessonStudent provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lessonStudentPrimaryKeyMapping)
	sql := "DELETE FROM \"lesson_student\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from lesson_student")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for lesson_student")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lessonStudentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no lessonStudentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from lesson_student")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for lesson_student")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LessonStudentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lessonStudentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"lesson_student\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lessonStudentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from lessonStudent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for lesson_student")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LessonStudent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLessonStudent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LessonStudentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LessonStudentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lessonStudentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"lesson_student\".* FROM \"lesson_student\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lessonStudentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LessonStudentSlice")
	}

	*o = slice

	return nil
}

// LessonStudentExists checks if the LessonStudent row exists.
func LessonStudentExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"lesson_student\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if lesson_student exists")
	}

	return exists, nil
}
