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

// Timetable is an object representing the database table.
type Timetable struct {
	ID         string `boil:"id" json:"id" toml:"id" yaml:"id"`
	SubclassID string `boil:"subclass_id" json:"subclass_id" toml:"subclass_id" yaml:"subclass_id"`
	SubjectID  string `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	TeacherID  string `boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	PeriodID   string `boil:"period_id" json:"period_id" toml:"period_id" yaml:"period_id"`
	Day        int    `boil:"day" json:"day" toml:"day" yaml:"day"`

	R *timetableR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L timetableL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TimetableColumns = struct {
	ID         string
	SubclassID string
	SubjectID  string
	TeacherID  string
	PeriodID   string
	Day        string
}{
	ID:         "id",
	SubclassID: "subclass_id",
	SubjectID:  "subject_id",
	TeacherID:  "teacher_id",
	PeriodID:   "period_id",
	Day:        "day",
}

// Generated where

var TimetableWhere = struct {
	ID         whereHelperstring
	SubclassID whereHelperstring
	SubjectID  whereHelperstring
	TeacherID  whereHelperstring
	PeriodID   whereHelperstring
	Day        whereHelperint
}{
	ID:         whereHelperstring{field: "\"timetable\".\"id\""},
	SubclassID: whereHelperstring{field: "\"timetable\".\"subclass_id\""},
	SubjectID:  whereHelperstring{field: "\"timetable\".\"subject_id\""},
	TeacherID:  whereHelperstring{field: "\"timetable\".\"teacher_id\""},
	PeriodID:   whereHelperstring{field: "\"timetable\".\"period_id\""},
	Day:        whereHelperint{field: "\"timetable\".\"day\""},
}

// TimetableRels is where relationship names are stored.
var TimetableRels = struct {
	Period   string
	Subclass string
	Subject  string
	Teacher  string
	Lessons  string
}{
	Period:   "Period",
	Subclass: "Subclass",
	Subject:  "Subject",
	Teacher:  "Teacher",
	Lessons:  "Lessons",
}

// timetableR is where relationships are stored.
type timetableR struct {
	Period   *Period
	Subclass *Subclass
	Subject  *Subject
	Teacher  *User
	Lessons  LessonSlice
}

// NewStruct creates a new relationship struct
func (*timetableR) NewStruct() *timetableR {
	return &timetableR{}
}

// timetableL is where Load methods for each relationship are stored.
type timetableL struct{}

var (
	timetableAllColumns            = []string{"id", "subclass_id", "subject_id", "teacher_id", "period_id", "day"}
	timetableColumnsWithoutDefault = []string{"id", "subclass_id", "subject_id", "teacher_id", "period_id", "day"}
	timetableColumnsWithDefault    = []string{}
	timetablePrimaryKeyColumns     = []string{"id"}
)

type (
	// TimetableSlice is an alias for a slice of pointers to Timetable.
	// This should generally be used opposed to []Timetable.
	TimetableSlice []*Timetable

	timetableQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	timetableType                 = reflect.TypeOf(&Timetable{})
	timetableMapping              = queries.MakeStructMapping(timetableType)
	timetablePrimaryKeyMapping, _ = queries.BindMapping(timetableType, timetableMapping, timetablePrimaryKeyColumns)
	timetableInsertCacheMut       sync.RWMutex
	timetableInsertCache          = make(map[string]insertCache)
	timetableUpdateCacheMut       sync.RWMutex
	timetableUpdateCache          = make(map[string]updateCache)
	timetableUpsertCacheMut       sync.RWMutex
	timetableUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single timetable record from the query.
func (q timetableQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Timetable, error) {
	o := &Timetable{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for timetable")
	}

	return o, nil
}

// All returns all Timetable records from the query.
func (q timetableQuery) All(ctx context.Context, exec boil.ContextExecutor) (TimetableSlice, error) {
	var o []*Timetable

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Timetable slice")
	}

	return o, nil
}

// Count returns the count of all Timetable records in the query.
func (q timetableQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count timetable rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q timetableQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if timetable exists")
	}

	return count > 0, nil
}

// Period pointed to by the foreign key.
func (o *Timetable) Period(mods ...qm.QueryMod) periodQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PeriodID),
	}

	queryMods = append(queryMods, mods...)

	query := Periods(queryMods...)
	queries.SetFrom(query.Query, "\"period\"")

	return query
}

// Subclass pointed to by the foreign key.
func (o *Timetable) Subclass(mods ...qm.QueryMod) subclassQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SubclassID),
	}

	queryMods = append(queryMods, mods...)

	query := Subclasses(queryMods...)
	queries.SetFrom(query.Query, "\"subclass\"")

	return query
}

// Subject pointed to by the foreign key.
func (o *Timetable) Subject(mods ...qm.QueryMod) subjectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SubjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Subjects(queryMods...)
	queries.SetFrom(query.Query, "\"subject\"")

	return query
}

// Teacher pointed to by the foreign key.
func (o *Timetable) Teacher(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TeacherID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Lessons retrieves all the lesson's Lessons with an executor.
func (o *Timetable) Lessons(mods ...qm.QueryMod) lessonQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"lesson\".\"timetable_id\"=?", o.ID),
	)

	query := Lessons(queryMods...)
	queries.SetFrom(query.Query, "\"lesson\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"lesson\".*"})
	}

	return query
}

// LoadPeriod allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (timetableL) LoadPeriod(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTimetable interface{}, mods queries.Applicator) error {
	var slice []*Timetable
	var object *Timetable

	if singular {
		object = maybeTimetable.(*Timetable)
	} else {
		slice = *maybeTimetable.(*[]*Timetable)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &timetableR{}
		}
		args = append(args, object.PeriodID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &timetableR{}
			}

			for _, a := range args {
				if a == obj.PeriodID {
					continue Outer
				}
			}

			args = append(args, obj.PeriodID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`period`), qm.WhereIn(`period.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Period")
	}

	var resultSlice []*Period
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Period")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for period")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for period")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Period = foreign
		if foreign.R == nil {
			foreign.R = &periodR{}
		}
		foreign.R.Timetables = append(foreign.R.Timetables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PeriodID == foreign.ID {
				local.R.Period = foreign
				if foreign.R == nil {
					foreign.R = &periodR{}
				}
				foreign.R.Timetables = append(foreign.R.Timetables, local)
				break
			}
		}
	}

	return nil
}

// LoadSubclass allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (timetableL) LoadSubclass(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTimetable interface{}, mods queries.Applicator) error {
	var slice []*Timetable
	var object *Timetable

	if singular {
		object = maybeTimetable.(*Timetable)
	} else {
		slice = *maybeTimetable.(*[]*Timetable)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &timetableR{}
		}
		args = append(args, object.SubclassID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &timetableR{}
			}

			for _, a := range args {
				if a == obj.SubclassID {
					continue Outer
				}
			}

			args = append(args, obj.SubclassID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`subclass`), qm.WhereIn(`subclass.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Subclass")
	}

	var resultSlice []*Subclass
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Subclass")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for subclass")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subclass")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Subclass = foreign
		if foreign.R == nil {
			foreign.R = &subclassR{}
		}
		foreign.R.Timetables = append(foreign.R.Timetables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SubclassID == foreign.ID {
				local.R.Subclass = foreign
				if foreign.R == nil {
					foreign.R = &subclassR{}
				}
				foreign.R.Timetables = append(foreign.R.Timetables, local)
				break
			}
		}
	}

	return nil
}

// LoadSubject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (timetableL) LoadSubject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTimetable interface{}, mods queries.Applicator) error {
	var slice []*Timetable
	var object *Timetable

	if singular {
		object = maybeTimetable.(*Timetable)
	} else {
		slice = *maybeTimetable.(*[]*Timetable)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &timetableR{}
		}
		args = append(args, object.SubjectID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &timetableR{}
			}

			for _, a := range args {
				if a == obj.SubjectID {
					continue Outer
				}
			}

			args = append(args, obj.SubjectID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`subject`), qm.WhereIn(`subject.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Subject")
	}

	var resultSlice []*Subject
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Subject")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for subject")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subject")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Subject = foreign
		if foreign.R == nil {
			foreign.R = &subjectR{}
		}
		foreign.R.Timetables = append(foreign.R.Timetables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SubjectID == foreign.ID {
				local.R.Subject = foreign
				if foreign.R == nil {
					foreign.R = &subjectR{}
				}
				foreign.R.Timetables = append(foreign.R.Timetables, local)
				break
			}
		}
	}

	return nil
}

// LoadTeacher allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (timetableL) LoadTeacher(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTimetable interface{}, mods queries.Applicator) error {
	var slice []*Timetable
	var object *Timetable

	if singular {
		object = maybeTimetable.(*Timetable)
	} else {
		slice = *maybeTimetable.(*[]*Timetable)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &timetableR{}
		}
		args = append(args, object.TeacherID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &timetableR{}
			}

			for _, a := range args {
				if a == obj.TeacherID {
					continue Outer
				}
			}

			args = append(args, obj.TeacherID)

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
		foreign.R.TeacherTimetables = append(foreign.R.TeacherTimetables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TeacherID == foreign.ID {
				local.R.Teacher = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.TeacherTimetables = append(foreign.R.TeacherTimetables, local)
				break
			}
		}
	}

	return nil
}

// LoadLessons allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (timetableL) LoadLessons(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTimetable interface{}, mods queries.Applicator) error {
	var slice []*Timetable
	var object *Timetable

	if singular {
		object = maybeTimetable.(*Timetable)
	} else {
		slice = *maybeTimetable.(*[]*Timetable)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &timetableR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &timetableR{}
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

	query := NewQuery(qm.From(`lesson`), qm.WhereIn(`lesson.timetable_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load lesson")
	}

	var resultSlice []*Lesson
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice lesson")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on lesson")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for lesson")
	}

	if singular {
		object.R.Lessons = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &lessonR{}
			}
			foreign.R.Timetable = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TimetableID {
				local.R.Lessons = append(local.R.Lessons, foreign)
				if foreign.R == nil {
					foreign.R = &lessonR{}
				}
				foreign.R.Timetable = local
				break
			}
		}
	}

	return nil
}

// SetPeriod of the timetable to the related item.
// Sets o.R.Period to related.
// Adds o to related.R.Timetables.
func (o *Timetable) SetPeriod(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Period) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"period_id"}),
		strmangle.WhereClause("\"", "\"", 2, timetablePrimaryKeyColumns),
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

	o.PeriodID = related.ID
	if o.R == nil {
		o.R = &timetableR{
			Period: related,
		}
	} else {
		o.R.Period = related
	}

	if related.R == nil {
		related.R = &periodR{
			Timetables: TimetableSlice{o},
		}
	} else {
		related.R.Timetables = append(related.R.Timetables, o)
	}

	return nil
}

// SetSubclass of the timetable to the related item.
// Sets o.R.Subclass to related.
// Adds o to related.R.Timetables.
func (o *Timetable) SetSubclass(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Subclass) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"subclass_id"}),
		strmangle.WhereClause("\"", "\"", 2, timetablePrimaryKeyColumns),
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

	o.SubclassID = related.ID
	if o.R == nil {
		o.R = &timetableR{
			Subclass: related,
		}
	} else {
		o.R.Subclass = related
	}

	if related.R == nil {
		related.R = &subclassR{
			Timetables: TimetableSlice{o},
		}
	} else {
		related.R.Timetables = append(related.R.Timetables, o)
	}

	return nil
}

// SetSubject of the timetable to the related item.
// Sets o.R.Subject to related.
// Adds o to related.R.Timetables.
func (o *Timetable) SetSubject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Subject) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
		strmangle.WhereClause("\"", "\"", 2, timetablePrimaryKeyColumns),
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

	o.SubjectID = related.ID
	if o.R == nil {
		o.R = &timetableR{
			Subject: related,
		}
	} else {
		o.R.Subject = related
	}

	if related.R == nil {
		related.R = &subjectR{
			Timetables: TimetableSlice{o},
		}
	} else {
		related.R.Timetables = append(related.R.Timetables, o)
	}

	return nil
}

// SetTeacher of the timetable to the related item.
// Sets o.R.Teacher to related.
// Adds o to related.R.TeacherTimetables.
func (o *Timetable) SetTeacher(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"teacher_id"}),
		strmangle.WhereClause("\"", "\"", 2, timetablePrimaryKeyColumns),
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

	o.TeacherID = related.ID
	if o.R == nil {
		o.R = &timetableR{
			Teacher: related,
		}
	} else {
		o.R.Teacher = related
	}

	if related.R == nil {
		related.R = &userR{
			TeacherTimetables: TimetableSlice{o},
		}
	} else {
		related.R.TeacherTimetables = append(related.R.TeacherTimetables, o)
	}

	return nil
}

// AddLessons adds the given related objects to the existing relationships
// of the timetable, optionally inserting them as new records.
// Appends related to o.R.Lessons.
// Sets related.R.Timetable appropriately.
func (o *Timetable) AddLessons(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Lesson) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TimetableID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"lesson\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"timetable_id"}),
				strmangle.WhereClause("\"", "\"", 2, lessonPrimaryKeyColumns),
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

			rel.TimetableID = o.ID
		}
	}

	if o.R == nil {
		o.R = &timetableR{
			Lessons: related,
		}
	} else {
		o.R.Lessons = append(o.R.Lessons, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &lessonR{
				Timetable: o,
			}
		} else {
			rel.R.Timetable = o
		}
	}
	return nil
}

// Timetables retrieves all the records using an executor.
func Timetables(mods ...qm.QueryMod) timetableQuery {
	mods = append(mods, qm.From("\"timetable\""))
	return timetableQuery{NewQuery(mods...)}
}

// FindTimetable retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTimetable(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Timetable, error) {
	timetableObj := &Timetable{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"timetable\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, timetableObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from timetable")
	}

	return timetableObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Timetable) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no timetable provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(timetableColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	timetableInsertCacheMut.RLock()
	cache, cached := timetableInsertCache[key]
	timetableInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			timetableAllColumns,
			timetableColumnsWithDefault,
			timetableColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(timetableType, timetableMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(timetableType, timetableMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"timetable\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"timetable\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into timetable")
	}

	if !cached {
		timetableInsertCacheMut.Lock()
		timetableInsertCache[key] = cache
		timetableInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Timetable.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Timetable) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	timetableUpdateCacheMut.RLock()
	cache, cached := timetableUpdateCache[key]
	timetableUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			timetableAllColumns,
			timetablePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update timetable, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"timetable\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, timetablePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(timetableType, timetableMapping, append(wl, timetablePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update timetable row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for timetable")
	}

	if !cached {
		timetableUpdateCacheMut.Lock()
		timetableUpdateCache[key] = cache
		timetableUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q timetableQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for timetable")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TimetableSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), timetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, timetablePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in timetable slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all timetable")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Timetable) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no timetable provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(timetableColumnsWithDefault, o)

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

	timetableUpsertCacheMut.RLock()
	cache, cached := timetableUpsertCache[key]
	timetableUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			timetableAllColumns,
			timetableColumnsWithDefault,
			timetableColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			timetableAllColumns,
			timetablePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert timetable, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(timetablePrimaryKeyColumns))
			copy(conflict, timetablePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"timetable\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(timetableType, timetableMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(timetableType, timetableMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert timetable")
	}

	if !cached {
		timetableUpsertCacheMut.Lock()
		timetableUpsertCache[key] = cache
		timetableUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Timetable record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Timetable) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Timetable provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), timetablePrimaryKeyMapping)
	sql := "DELETE FROM \"timetable\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for timetable")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q timetableQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no timetableQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for timetable")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TimetableSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), timetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"timetable\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, timetablePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from timetable slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for timetable")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Timetable) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTimetable(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TimetableSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TimetableSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), timetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"timetable\".* FROM \"timetable\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, timetablePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TimetableSlice")
	}

	*o = slice

	return nil
}

// TimetableExists checks if the Timetable row exists.
func TimetableExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"timetable\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if timetable exists")
	}

	return exists, nil
}
