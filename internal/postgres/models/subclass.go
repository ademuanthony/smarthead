// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Subclass is an object representing the database table.
type Subclass struct {
	ID          string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ClassID     string `boil:"class_id" json:"class_id" toml:"class_id" yaml:"class_id"`
	SchoolOrder int    `boil:"school_order" json:"school_order" toml:"school_order" yaml:"school_order"`
	Link        string `boil:"link" json:"link" toml:"link" yaml:"link"`

	R *subclassR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L subclassL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SubclassColumns = struct {
	ID          string
	Name        string
	ClassID     string
	SchoolOrder string
	Link        string
}{
	ID:          "id",
	Name:        "name",
	ClassID:     "class_id",
	SchoolOrder: "school_order",
	Link:        "link",
}

// Generated where

var SubclassWhere = struct {
	ID          whereHelperstring
	Name        whereHelperstring
	ClassID     whereHelperstring
	SchoolOrder whereHelperint
	Link        whereHelperstring
}{
	ID:          whereHelperstring{field: "\"subclass\".\"id\""},
	Name:        whereHelperstring{field: "\"subclass\".\"name\""},
	ClassID:     whereHelperstring{field: "\"subclass\".\"class_id\""},
	SchoolOrder: whereHelperint{field: "\"subclass\".\"school_order\""},
	Link:        whereHelperstring{field: "\"subclass\".\"link\""},
}

// SubclassRels is where relationship names are stored.
var SubclassRels = struct {
	Class      string
	Students   string
	Timetables string
}{
	Class:      "Class",
	Students:   "Students",
	Timetables: "Timetables",
}

// subclassR is where relationships are stored.
type subclassR struct {
	Class      *Class
	Students   StudentSlice
	Timetables TimetableSlice
}

// NewStruct creates a new relationship struct
func (*subclassR) NewStruct() *subclassR {
	return &subclassR{}
}

// subclassL is where Load methods for each relationship are stored.
type subclassL struct{}

var (
	subclassAllColumns            = []string{"id", "name", "class_id", "school_order", "link"}
	subclassColumnsWithoutDefault = []string{"id", "name", "class_id", "school_order"}
	subclassColumnsWithDefault    = []string{"link"}
	subclassPrimaryKeyColumns     = []string{"id"}
)

type (
	// SubclassSlice is an alias for a slice of pointers to Subclass.
	// This should generally be used opposed to []Subclass.
	SubclassSlice []*Subclass

	subclassQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	subclassType                 = reflect.TypeOf(&Subclass{})
	subclassMapping              = queries.MakeStructMapping(subclassType)
	subclassPrimaryKeyMapping, _ = queries.BindMapping(subclassType, subclassMapping, subclassPrimaryKeyColumns)
	subclassInsertCacheMut       sync.RWMutex
	subclassInsertCache          = make(map[string]insertCache)
	subclassUpdateCacheMut       sync.RWMutex
	subclassUpdateCache          = make(map[string]updateCache)
	subclassUpsertCacheMut       sync.RWMutex
	subclassUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single subclass record from the query.
func (q subclassQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Subclass, error) {
	o := &Subclass{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for subclass")
	}

	return o, nil
}

// All returns all Subclass records from the query.
func (q subclassQuery) All(ctx context.Context, exec boil.ContextExecutor) (SubclassSlice, error) {
	var o []*Subclass

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Subclass slice")
	}

	return o, nil
}

// Count returns the count of all Subclass records in the query.
func (q subclassQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count subclass rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q subclassQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if subclass exists")
	}

	return count > 0, nil
}

// Class pointed to by the foreign key.
func (o *Subclass) Class(mods ...qm.QueryMod) classQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ClassID),
	}

	queryMods = append(queryMods, mods...)

	query := Classes(queryMods...)
	queries.SetFrom(query.Query, "\"classes\"")

	return query
}

// Students retrieves all the student's Students with an executor.
func (o *Subclass) Students(mods ...qm.QueryMod) studentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"student\".\"subclass_id\"=?", o.ID),
	)

	query := Students(queryMods...)
	queries.SetFrom(query.Query, "\"student\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"student\".*"})
	}

	return query
}

// Timetables retrieves all the timetable's Timetables with an executor.
func (o *Subclass) Timetables(mods ...qm.QueryMod) timetableQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"timetable\".\"subclass_id\"=?", o.ID),
	)

	query := Timetables(queryMods...)
	queries.SetFrom(query.Query, "\"timetable\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"timetable\".*"})
	}

	return query
}

// LoadClass allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (subclassL) LoadClass(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSubclass interface{}, mods queries.Applicator) error {
	var slice []*Subclass
	var object *Subclass

	if singular {
		object = maybeSubclass.(*Subclass)
	} else {
		slice = *maybeSubclass.(*[]*Subclass)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &subclassR{}
		}
		args = append(args, object.ClassID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &subclassR{}
			}

			for _, a := range args {
				if a == obj.ClassID {
					continue Outer
				}
			}

			args = append(args, obj.ClassID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`classes`), qm.WhereIn(`classes.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Class")
	}

	var resultSlice []*Class
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Class")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for classes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for classes")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Class = foreign
		if foreign.R == nil {
			foreign.R = &classR{}
		}
		foreign.R.Subclasses = append(foreign.R.Subclasses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ClassID == foreign.ID {
				local.R.Class = foreign
				if foreign.R == nil {
					foreign.R = &classR{}
				}
				foreign.R.Subclasses = append(foreign.R.Subclasses, local)
				break
			}
		}
	}

	return nil
}

// LoadStudents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (subclassL) LoadStudents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSubclass interface{}, mods queries.Applicator) error {
	var slice []*Subclass
	var object *Subclass

	if singular {
		object = maybeSubclass.(*Subclass)
	} else {
		slice = *maybeSubclass.(*[]*Subclass)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &subclassR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &subclassR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`student`), qm.WhereIn(`student.subclass_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load student")
	}

	var resultSlice []*Student
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice student")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on student")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for student")
	}

	if singular {
		object.R.Students = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &studentR{}
			}
			foreign.R.Subclass = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.SubclassID) {
				local.R.Students = append(local.R.Students, foreign)
				if foreign.R == nil {
					foreign.R = &studentR{}
				}
				foreign.R.Subclass = local
				break
			}
		}
	}

	return nil
}

// LoadTimetables allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (subclassL) LoadTimetables(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSubclass interface{}, mods queries.Applicator) error {
	var slice []*Subclass
	var object *Subclass

	if singular {
		object = maybeSubclass.(*Subclass)
	} else {
		slice = *maybeSubclass.(*[]*Subclass)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &subclassR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &subclassR{}
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

	query := NewQuery(qm.From(`timetable`), qm.WhereIn(`timetable.subclass_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load timetable")
	}

	var resultSlice []*Timetable
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice timetable")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on timetable")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for timetable")
	}

	if singular {
		object.R.Timetables = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &timetableR{}
			}
			foreign.R.Subclass = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.SubclassID {
				local.R.Timetables = append(local.R.Timetables, foreign)
				if foreign.R == nil {
					foreign.R = &timetableR{}
				}
				foreign.R.Subclass = local
				break
			}
		}
	}

	return nil
}

// SetClass of the subclass to the related item.
// Sets o.R.Class to related.
// Adds o to related.R.Subclasses.
func (o *Subclass) SetClass(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Class) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"subclass\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"class_id"}),
		strmangle.WhereClause("\"", "\"", 2, subclassPrimaryKeyColumns),
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

	o.ClassID = related.ID
	if o.R == nil {
		o.R = &subclassR{
			Class: related,
		}
	} else {
		o.R.Class = related
	}

	if related.R == nil {
		related.R = &classR{
			Subclasses: SubclassSlice{o},
		}
	} else {
		related.R.Subclasses = append(related.R.Subclasses, o)
	}

	return nil
}

// AddStudents adds the given related objects to the existing relationships
// of the subclass, optionally inserting them as new records.
// Appends related to o.R.Students.
// Sets related.R.Subclass appropriately.
func (o *Subclass) AddStudents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Student) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.SubclassID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"student\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"subclass_id"}),
				strmangle.WhereClause("\"", "\"", 2, studentPrimaryKeyColumns),
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

			queries.Assign(&rel.SubclassID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &subclassR{
			Students: related,
		}
	} else {
		o.R.Students = append(o.R.Students, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &studentR{
				Subclass: o,
			}
		} else {
			rel.R.Subclass = o
		}
	}
	return nil
}

// SetStudents removes all previously related items of the
// subclass replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Subclass's Students accordingly.
// Replaces o.R.Students with related.
// Sets related.R.Subclass's Students accordingly.
func (o *Subclass) SetStudents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Student) error {
	query := "update \"student\" set \"subclass_id\" = null where \"subclass_id\" = $1"
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

	if o.R != nil {
		for _, rel := range o.R.Students {
			queries.SetScanner(&rel.SubclassID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Subclass = nil
		}

		o.R.Students = nil
	}
	return o.AddStudents(ctx, exec, insert, related...)
}

// RemoveStudents relationships from objects passed in.
// Removes related items from R.Students (uses pointer comparison, removal does not keep order)
// Sets related.R.Subclass.
func (o *Subclass) RemoveStudents(ctx context.Context, exec boil.ContextExecutor, related ...*Student) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.SubclassID, nil)
		if rel.R != nil {
			rel.R.Subclass = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("subclass_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Students {
			if rel != ri {
				continue
			}

			ln := len(o.R.Students)
			if ln > 1 && i < ln-1 {
				o.R.Students[i] = o.R.Students[ln-1]
			}
			o.R.Students = o.R.Students[:ln-1]
			break
		}
	}

	return nil
}

// AddTimetables adds the given related objects to the existing relationships
// of the subclass, optionally inserting them as new records.
// Appends related to o.R.Timetables.
// Sets related.R.Subclass appropriately.
func (o *Subclass) AddTimetables(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Timetable) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SubclassID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"timetable\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"subclass_id"}),
				strmangle.WhereClause("\"", "\"", 2, timetablePrimaryKeyColumns),
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

			rel.SubclassID = o.ID
		}
	}

	if o.R == nil {
		o.R = &subclassR{
			Timetables: related,
		}
	} else {
		o.R.Timetables = append(o.R.Timetables, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &timetableR{
				Subclass: o,
			}
		} else {
			rel.R.Subclass = o
		}
	}
	return nil
}

// Subclasses retrieves all the records using an executor.
func Subclasses(mods ...qm.QueryMod) subclassQuery {
	mods = append(mods, qm.From("\"subclass\""))
	return subclassQuery{NewQuery(mods...)}
}

// FindSubclass retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSubclass(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Subclass, error) {
	subclassObj := &Subclass{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"subclass\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, subclassObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from subclass")
	}

	return subclassObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Subclass) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no subclass provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(subclassColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	subclassInsertCacheMut.RLock()
	cache, cached := subclassInsertCache[key]
	subclassInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			subclassAllColumns,
			subclassColumnsWithDefault,
			subclassColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(subclassType, subclassMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(subclassType, subclassMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"subclass\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"subclass\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into subclass")
	}

	if !cached {
		subclassInsertCacheMut.Lock()
		subclassInsertCache[key] = cache
		subclassInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Subclass.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Subclass) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	subclassUpdateCacheMut.RLock()
	cache, cached := subclassUpdateCache[key]
	subclassUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			subclassAllColumns,
			subclassPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update subclass, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"subclass\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, subclassPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(subclassType, subclassMapping, append(wl, subclassPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update subclass row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for subclass")
	}

	if !cached {
		subclassUpdateCacheMut.Lock()
		subclassUpdateCache[key] = cache
		subclassUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q subclassQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for subclass")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for subclass")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SubclassSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subclassPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"subclass\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, subclassPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in subclass slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all subclass")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Subclass) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no subclass provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(subclassColumnsWithDefault, o)

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

	subclassUpsertCacheMut.RLock()
	cache, cached := subclassUpsertCache[key]
	subclassUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			subclassAllColumns,
			subclassColumnsWithDefault,
			subclassColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			subclassAllColumns,
			subclassPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert subclass, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(subclassPrimaryKeyColumns))
			copy(conflict, subclassPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"subclass\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(subclassType, subclassMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(subclassType, subclassMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert subclass")
	}

	if !cached {
		subclassUpsertCacheMut.Lock()
		subclassUpsertCache[key] = cache
		subclassUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Subclass record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Subclass) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Subclass provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), subclassPrimaryKeyMapping)
	sql := "DELETE FROM \"subclass\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from subclass")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for subclass")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q subclassQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no subclassQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from subclass")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for subclass")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SubclassSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subclassPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"subclass\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, subclassPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from subclass slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for subclass")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Subclass) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSubclass(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SubclassSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SubclassSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subclassPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"subclass\".* FROM \"subclass\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, subclassPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SubclassSlice")
	}

	*o = slice

	return nil
}

// SubclassExists checks if the Subclass row exists.
func SubclassExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"subclass\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if subclass exists")
	}

	return exists, nil
}
