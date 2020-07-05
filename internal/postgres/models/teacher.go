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

// Teacher is an object representing the database table.
type Teacher struct {
	ID                string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name              string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CV                string    `boil:"cv" json:"cv" toml:"cv" yaml:"cv"`
	YearsOfExperience int       `boil:"years_of_experience" json:"years_of_experience" toml:"years_of_experience" yaml:"years_of_experience"`
	CreatedAt         time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt         time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *teacherR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L teacherL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TeacherColumns = struct {
	ID                string
	Name              string
	CV                string
	YearsOfExperience string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	Name:              "name",
	CV:                "cv",
	YearsOfExperience: "years_of_experience",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// Generated where

var TeacherWhere = struct {
	ID                whereHelperstring
	Name              whereHelperstring
	CV                whereHelperstring
	YearsOfExperience whereHelperint
	CreatedAt         whereHelpertime_Time
	UpdatedAt         whereHelpertime_Time
}{
	ID:                whereHelperstring{field: "\"teacher\".\"id\""},
	Name:              whereHelperstring{field: "\"teacher\".\"name\""},
	CV:                whereHelperstring{field: "\"teacher\".\"cv\""},
	YearsOfExperience: whereHelperint{field: "\"teacher\".\"years_of_experience\""},
	CreatedAt:         whereHelpertime_Time{field: "\"teacher\".\"created_at\""},
	UpdatedAt:         whereHelpertime_Time{field: "\"teacher\".\"updated_at\""},
}

// TeacherRels is where relationship names are stored.
var TeacherRels = struct {
	Subjects string
}{
	Subjects: "Subjects",
}

// teacherR is where relationships are stored.
type teacherR struct {
	Subjects SubjectSlice
}

// NewStruct creates a new relationship struct
func (*teacherR) NewStruct() *teacherR {
	return &teacherR{}
}

// teacherL is where Load methods for each relationship are stored.
type teacherL struct{}

var (
	teacherAllColumns            = []string{"id", "name", "cv", "years_of_experience", "created_at", "updated_at"}
	teacherColumnsWithoutDefault = []string{"id", "name", "cv", "years_of_experience", "created_at", "updated_at"}
	teacherColumnsWithDefault    = []string{}
	teacherPrimaryKeyColumns     = []string{"id"}
)

type (
	// TeacherSlice is an alias for a slice of pointers to Teacher.
	// This should generally be used opposed to []Teacher.
	TeacherSlice []*Teacher

	teacherQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	teacherType                 = reflect.TypeOf(&Teacher{})
	teacherMapping              = queries.MakeStructMapping(teacherType)
	teacherPrimaryKeyMapping, _ = queries.BindMapping(teacherType, teacherMapping, teacherPrimaryKeyColumns)
	teacherInsertCacheMut       sync.RWMutex
	teacherInsertCache          = make(map[string]insertCache)
	teacherUpdateCacheMut       sync.RWMutex
	teacherUpdateCache          = make(map[string]updateCache)
	teacherUpsertCacheMut       sync.RWMutex
	teacherUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single teacher record from the query.
func (q teacherQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Teacher, error) {
	o := &Teacher{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for teacher")
	}

	return o, nil
}

// All returns all Teacher records from the query.
func (q teacherQuery) All(ctx context.Context, exec boil.ContextExecutor) (TeacherSlice, error) {
	var o []*Teacher

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Teacher slice")
	}

	return o, nil
}

// Count returns the count of all Teacher records in the query.
func (q teacherQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count teacher rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q teacherQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if teacher exists")
	}

	return count > 0, nil
}

// Subjects retrieves all the subject's Subjects with an executor.
func (o *Teacher) Subjects(mods ...qm.QueryMod) subjectQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"specialization\" on \"subject\".\"id\" = \"specialization\".\"subject_id\""),
		qm.Where("\"specialization\".\"teacher_id\"=?", o.ID),
	)

	query := Subjects(queryMods...)
	queries.SetFrom(query.Query, "\"subject\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"subject\".*"})
	}

	return query
}

// LoadSubjects allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (teacherL) LoadSubjects(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTeacher interface{}, mods queries.Applicator) error {
	var slice []*Teacher
	var object *Teacher

	if singular {
		object = maybeTeacher.(*Teacher)
	} else {
		slice = *maybeTeacher.(*[]*Teacher)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &teacherR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &teacherR{}
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

	query := NewQuery(
		qm.Select("\"subject\".*, \"a\".\"teacher_id\""),
		qm.From("\"subject\""),
		qm.InnerJoin("\"specialization\" as \"a\" on \"subject\".\"id\" = \"a\".\"subject_id\""),
		qm.WhereIn("\"a\".\"teacher_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load subject")
	}

	var resultSlice []*Subject

	var localJoinCols []string
	for results.Next() {
		one := new(Subject)
		var localJoinCol string

		err = results.Scan(&one.ID, &one.Name, &one.CreatedAt, &one.UpdatedAt, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for subject")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice subject")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on subject")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subject")
	}

	if singular {
		object.R.Subjects = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &subjectR{}
			}
			foreign.R.Teachers = append(foreign.R.Teachers, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Subjects = append(local.R.Subjects, foreign)
				if foreign.R == nil {
					foreign.R = &subjectR{}
				}
				foreign.R.Teachers = append(foreign.R.Teachers, local)
				break
			}
		}
	}

	return nil
}

// AddSubjects adds the given related objects to the existing relationships
// of the teacher, optionally inserting them as new records.
// Appends related to o.R.Subjects.
// Sets related.R.Teachers appropriately.
func (o *Teacher) AddSubjects(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Subject) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"specialization\" (\"teacher_id\", \"subject_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &teacherR{
			Subjects: related,
		}
	} else {
		o.R.Subjects = append(o.R.Subjects, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &subjectR{
				Teachers: TeacherSlice{o},
			}
		} else {
			rel.R.Teachers = append(rel.R.Teachers, o)
		}
	}
	return nil
}

// SetSubjects removes all previously related items of the
// teacher replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Teachers's Subjects accordingly.
// Replaces o.R.Subjects with related.
// Sets related.R.Teachers's Subjects accordingly.
func (o *Teacher) SetSubjects(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Subject) error {
	query := "delete from \"specialization\" where \"teacher_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeSubjectsFromTeachersSlice(o, related)
	if o.R != nil {
		o.R.Subjects = nil
	}
	return o.AddSubjects(ctx, exec, insert, related...)
}

// RemoveSubjects relationships from objects passed in.
// Removes related items from R.Subjects (uses pointer comparison, removal does not keep order)
// Sets related.R.Teachers.
func (o *Teacher) RemoveSubjects(ctx context.Context, exec boil.ContextExecutor, related ...*Subject) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"specialization\" where \"teacher_id\" = $1 and \"subject_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeSubjectsFromTeachersSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Subjects {
			if rel != ri {
				continue
			}

			ln := len(o.R.Subjects)
			if ln > 1 && i < ln-1 {
				o.R.Subjects[i] = o.R.Subjects[ln-1]
			}
			o.R.Subjects = o.R.Subjects[:ln-1]
			break
		}
	}

	return nil
}

func removeSubjectsFromTeachersSlice(o *Teacher, related []*Subject) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Teachers {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Teachers)
			if ln > 1 && i < ln-1 {
				rel.R.Teachers[i] = rel.R.Teachers[ln-1]
			}
			rel.R.Teachers = rel.R.Teachers[:ln-1]
			break
		}
	}
}

// Teachers retrieves all the records using an executor.
func Teachers(mods ...qm.QueryMod) teacherQuery {
	mods = append(mods, qm.From("\"teacher\""))
	return teacherQuery{NewQuery(mods...)}
}

// FindTeacher retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTeacher(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Teacher, error) {
	teacherObj := &Teacher{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"teacher\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, teacherObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from teacher")
	}

	return teacherObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Teacher) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no teacher provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(teacherColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	teacherInsertCacheMut.RLock()
	cache, cached := teacherInsertCache[key]
	teacherInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			teacherAllColumns,
			teacherColumnsWithDefault,
			teacherColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(teacherType, teacherMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(teacherType, teacherMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"teacher\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"teacher\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into teacher")
	}

	if !cached {
		teacherInsertCacheMut.Lock()
		teacherInsertCache[key] = cache
		teacherInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Teacher.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Teacher) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	teacherUpdateCacheMut.RLock()
	cache, cached := teacherUpdateCache[key]
	teacherUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			teacherAllColumns,
			teacherPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update teacher, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"teacher\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, teacherPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(teacherType, teacherMapping, append(wl, teacherPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update teacher row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for teacher")
	}

	if !cached {
		teacherUpdateCacheMut.Lock()
		teacherUpdateCache[key] = cache
		teacherUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q teacherQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for teacher")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TeacherSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"teacher\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, teacherPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in teacher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all teacher")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Teacher) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no teacher provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(teacherColumnsWithDefault, o)

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

	teacherUpsertCacheMut.RLock()
	cache, cached := teacherUpsertCache[key]
	teacherUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			teacherAllColumns,
			teacherColumnsWithDefault,
			teacherColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			teacherAllColumns,
			teacherPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert teacher, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(teacherPrimaryKeyColumns))
			copy(conflict, teacherPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"teacher\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(teacherType, teacherMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(teacherType, teacherMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert teacher")
	}

	if !cached {
		teacherUpsertCacheMut.Lock()
		teacherUpsertCache[key] = cache
		teacherUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Teacher record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Teacher) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Teacher provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), teacherPrimaryKeyMapping)
	sql := "DELETE FROM \"teacher\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for teacher")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q teacherQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no teacherQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for teacher")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TeacherSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"teacher\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teacherPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from teacher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for teacher")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Teacher) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTeacher(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TeacherSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TeacherSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"teacher\".* FROM \"teacher\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teacherPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TeacherSlice")
	}

	*o = slice

	return nil
}

// TeacherExists checks if the Teacher row exists.
func TeacherExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"teacher\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if teacher exists")
	}

	return exists, nil
}
