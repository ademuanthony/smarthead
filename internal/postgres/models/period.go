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

// Period is an object representing the database table.
type Period struct {
	ID          string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	StartHour   int       `boil:"start_hour" json:"start_hour" toml:"start_hour" yaml:"start_hour"`
	StartMinute int       `boil:"start_minute" json:"start_minute" toml:"start_minute" yaml:"start_minute"`
	EndHour     int       `boil:"end_hour" json:"end_hour" toml:"end_hour" yaml:"end_hour"`
	EndMinute   int       `boil:"end_minute" json:"end_minute" toml:"end_minute" yaml:"end_minute"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *periodR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L periodL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PeriodColumns = struct {
	ID          string
	StartHour   string
	StartMinute string
	EndHour     string
	EndMinute   string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	StartHour:   "start_hour",
	StartMinute: "start_minute",
	EndHour:     "end_hour",
	EndMinute:   "end_minute",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// Generated where

var PeriodWhere = struct {
	ID          whereHelperstring
	StartHour   whereHelperint
	StartMinute whereHelperint
	EndHour     whereHelperint
	EndMinute   whereHelperint
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperstring{field: "\"period\".\"id\""},
	StartHour:   whereHelperint{field: "\"period\".\"start_hour\""},
	StartMinute: whereHelperint{field: "\"period\".\"start_minute\""},
	EndHour:     whereHelperint{field: "\"period\".\"end_hour\""},
	EndMinute:   whereHelperint{field: "\"period\".\"end_minute\""},
	CreatedAt:   whereHelpertime_Time{field: "\"period\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"period\".\"updated_at\""},
}

// PeriodRels is where relationship names are stored.
var PeriodRels = struct {
	Deposits      string
	Students      string
	Subscriptions string
}{
	Deposits:      "Deposits",
	Students:      "Students",
	Subscriptions: "Subscriptions",
}

// periodR is where relationships are stored.
type periodR struct {
	Deposits      DepositSlice
	Students      StudentSlice
	Subscriptions SubscriptionSlice
}

// NewStruct creates a new relationship struct
func (*periodR) NewStruct() *periodR {
	return &periodR{}
}

// periodL is where Load methods for each relationship are stored.
type periodL struct{}

var (
	periodAllColumns            = []string{"id", "start_hour", "start_minute", "end_hour", "end_minute", "created_at", "updated_at"}
	periodColumnsWithoutDefault = []string{"id", "start_hour", "start_minute", "end_hour", "end_minute", "created_at", "updated_at"}
	periodColumnsWithDefault    = []string{}
	periodPrimaryKeyColumns     = []string{"id"}
)

type (
	// PeriodSlice is an alias for a slice of pointers to Period.
	// This should generally be used opposed to []Period.
	PeriodSlice []*Period

	periodQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	periodType                 = reflect.TypeOf(&Period{})
	periodMapping              = queries.MakeStructMapping(periodType)
	periodPrimaryKeyMapping, _ = queries.BindMapping(periodType, periodMapping, periodPrimaryKeyColumns)
	periodInsertCacheMut       sync.RWMutex
	periodInsertCache          = make(map[string]insertCache)
	periodUpdateCacheMut       sync.RWMutex
	periodUpdateCache          = make(map[string]updateCache)
	periodUpsertCacheMut       sync.RWMutex
	periodUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single period record from the query.
func (q periodQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Period, error) {
	o := &Period{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for period")
	}

	return o, nil
}

// All returns all Period records from the query.
func (q periodQuery) All(ctx context.Context, exec boil.ContextExecutor) (PeriodSlice, error) {
	var o []*Period

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Period slice")
	}

	return o, nil
}

// Count returns the count of all Period records in the query.
func (q periodQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count period rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q periodQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if period exists")
	}

	return count > 0, nil
}

// Deposits retrieves all the deposit's Deposits with an executor.
func (o *Period) Deposits(mods ...qm.QueryMod) depositQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"deposits\".\"period_id\"=?", o.ID),
	)

	query := Deposits(queryMods...)
	queries.SetFrom(query.Query, "\"deposits\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"deposits\".*"})
	}

	return query
}

// Students retrieves all the student's Students with an executor.
func (o *Period) Students(mods ...qm.QueryMod) studentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"students_periods\" on \"student\".\"id\" = \"students_periods\".\"student_id\""),
		qm.Where("\"students_periods\".\"period_id\"=?", o.ID),
	)

	query := Students(queryMods...)
	queries.SetFrom(query.Query, "\"student\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"student\".*"})
	}

	return query
}

// Subscriptions retrieves all the subscription's Subscriptions with an executor.
func (o *Period) Subscriptions(mods ...qm.QueryMod) subscriptionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"subscription\".\"period_id\"=?", o.ID),
	)

	query := Subscriptions(queryMods...)
	queries.SetFrom(query.Query, "\"subscription\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"subscription\".*"})
	}

	return query
}

// LoadDeposits allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (periodL) LoadDeposits(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeriod interface{}, mods queries.Applicator) error {
	var slice []*Period
	var object *Period

	if singular {
		object = maybePeriod.(*Period)
	} else {
		slice = *maybePeriod.(*[]*Period)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &periodR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &periodR{}
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

	query := NewQuery(qm.From(`deposits`), qm.WhereIn(`deposits.period_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load deposits")
	}

	var resultSlice []*Deposit
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice deposits")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on deposits")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for deposits")
	}

	if singular {
		object.R.Deposits = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &depositR{}
			}
			foreign.R.Period = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PeriodID {
				local.R.Deposits = append(local.R.Deposits, foreign)
				if foreign.R == nil {
					foreign.R = &depositR{}
				}
				foreign.R.Period = local
				break
			}
		}
	}

	return nil
}

// LoadStudents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (periodL) LoadStudents(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeriod interface{}, mods queries.Applicator) error {
	var slice []*Period
	var object *Period

	if singular {
		object = maybePeriod.(*Period)
	} else {
		slice = *maybePeriod.(*[]*Period)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &periodR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &periodR{}
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
		qm.Select("\"student\".*, \"a\".\"period_id\""),
		qm.From("\"student\""),
		qm.InnerJoin("\"students_periods\" as \"a\" on \"student\".\"id\" = \"a\".\"student_id\""),
		qm.WhereIn("\"a\".\"period_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load student")
	}

	var resultSlice []*Student

	var localJoinCols []string
	for results.Next() {
		one := new(Student)
		var localJoinCol string

		err = results.Scan(&one.ID, &one.Name, &one.Username, &one.Age, &one.AccountBalance, &one.CurrentClass, &one.ParentPhone, &one.ParentEmail, &one.CreatedAt, &one.UpdatedAt, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for student")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice student")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
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
			foreign.R.Periods = append(foreign.R.Periods, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Students = append(local.R.Students, foreign)
				if foreign.R == nil {
					foreign.R = &studentR{}
				}
				foreign.R.Periods = append(foreign.R.Periods, local)
				break
			}
		}
	}

	return nil
}

// LoadSubscriptions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (periodL) LoadSubscriptions(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeriod interface{}, mods queries.Applicator) error {
	var slice []*Period
	var object *Period

	if singular {
		object = maybePeriod.(*Period)
	} else {
		slice = *maybePeriod.(*[]*Period)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &periodR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &periodR{}
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

	query := NewQuery(qm.From(`subscription`), qm.WhereIn(`subscription.period_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load subscription")
	}

	var resultSlice []*Subscription
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice subscription")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on subscription")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subscription")
	}

	if singular {
		object.R.Subscriptions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &subscriptionR{}
			}
			foreign.R.Period = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PeriodID {
				local.R.Subscriptions = append(local.R.Subscriptions, foreign)
				if foreign.R == nil {
					foreign.R = &subscriptionR{}
				}
				foreign.R.Period = local
				break
			}
		}
	}

	return nil
}

// AddDeposits adds the given related objects to the existing relationships
// of the period, optionally inserting them as new records.
// Appends related to o.R.Deposits.
// Sets related.R.Period appropriately.
func (o *Period) AddDeposits(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Deposit) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PeriodID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"deposits\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"period_id"}),
				strmangle.WhereClause("\"", "\"", 2, depositPrimaryKeyColumns),
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

			rel.PeriodID = o.ID
		}
	}

	if o.R == nil {
		o.R = &periodR{
			Deposits: related,
		}
	} else {
		o.R.Deposits = append(o.R.Deposits, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &depositR{
				Period: o,
			}
		} else {
			rel.R.Period = o
		}
	}
	return nil
}

// AddStudents adds the given related objects to the existing relationships
// of the period, optionally inserting them as new records.
// Appends related to o.R.Students.
// Sets related.R.Periods appropriately.
func (o *Period) AddStudents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Student) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"students_periods\" (\"period_id\", \"student_id\") values ($1, $2)"
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
		o.R = &periodR{
			Students: related,
		}
	} else {
		o.R.Students = append(o.R.Students, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &studentR{
				Periods: PeriodSlice{o},
			}
		} else {
			rel.R.Periods = append(rel.R.Periods, o)
		}
	}
	return nil
}

// SetStudents removes all previously related items of the
// period replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Periods's Students accordingly.
// Replaces o.R.Students with related.
// Sets related.R.Periods's Students accordingly.
func (o *Period) SetStudents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Student) error {
	query := "delete from \"students_periods\" where \"period_id\" = $1"
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

	removeStudentsFromPeriodsSlice(o, related)
	if o.R != nil {
		o.R.Students = nil
	}
	return o.AddStudents(ctx, exec, insert, related...)
}

// RemoveStudents relationships from objects passed in.
// Removes related items from R.Students (uses pointer comparison, removal does not keep order)
// Sets related.R.Periods.
func (o *Period) RemoveStudents(ctx context.Context, exec boil.ContextExecutor, related ...*Student) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"students_periods\" where \"period_id\" = $1 and \"student_id\" in (%s)",
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
	removeStudentsFromPeriodsSlice(o, related)
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

func removeStudentsFromPeriodsSlice(o *Period, related []*Student) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Periods {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Periods)
			if ln > 1 && i < ln-1 {
				rel.R.Periods[i] = rel.R.Periods[ln-1]
			}
			rel.R.Periods = rel.R.Periods[:ln-1]
			break
		}
	}
}

// AddSubscriptions adds the given related objects to the existing relationships
// of the period, optionally inserting them as new records.
// Appends related to o.R.Subscriptions.
// Sets related.R.Period appropriately.
func (o *Period) AddSubscriptions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Subscription) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PeriodID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"subscription\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"period_id"}),
				strmangle.WhereClause("\"", "\"", 2, subscriptionPrimaryKeyColumns),
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

			rel.PeriodID = o.ID
		}
	}

	if o.R == nil {
		o.R = &periodR{
			Subscriptions: related,
		}
	} else {
		o.R.Subscriptions = append(o.R.Subscriptions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &subscriptionR{
				Period: o,
			}
		} else {
			rel.R.Period = o
		}
	}
	return nil
}

// Periods retrieves all the records using an executor.
func Periods(mods ...qm.QueryMod) periodQuery {
	mods = append(mods, qm.From("\"period\""))
	return periodQuery{NewQuery(mods...)}
}

// FindPeriod retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPeriod(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Period, error) {
	periodObj := &Period{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"period\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, periodObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from period")
	}

	return periodObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Period) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no period provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(periodColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	periodInsertCacheMut.RLock()
	cache, cached := periodInsertCache[key]
	periodInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			periodAllColumns,
			periodColumnsWithDefault,
			periodColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(periodType, periodMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(periodType, periodMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"period\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"period\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into period")
	}

	if !cached {
		periodInsertCacheMut.Lock()
		periodInsertCache[key] = cache
		periodInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Period.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Period) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	periodUpdateCacheMut.RLock()
	cache, cached := periodUpdateCache[key]
	periodUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			periodAllColumns,
			periodPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update period, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"period\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, periodPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(periodType, periodMapping, append(wl, periodPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update period row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for period")
	}

	if !cached {
		periodUpdateCacheMut.Lock()
		periodUpdateCache[key] = cache
		periodUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q periodQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for period")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for period")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PeriodSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"period\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, periodPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in period slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all period")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Period) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no period provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(periodColumnsWithDefault, o)

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

	periodUpsertCacheMut.RLock()
	cache, cached := periodUpsertCache[key]
	periodUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			periodAllColumns,
			periodColumnsWithDefault,
			periodColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			periodAllColumns,
			periodPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert period, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(periodPrimaryKeyColumns))
			copy(conflict, periodPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"period\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(periodType, periodMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(periodType, periodMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert period")
	}

	if !cached {
		periodUpsertCacheMut.Lock()
		periodUpsertCache[key] = cache
		periodUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Period record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Period) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Period provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), periodPrimaryKeyMapping)
	sql := "DELETE FROM \"period\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from period")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for period")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q periodQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no periodQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from period")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for period")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PeriodSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"period\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, periodPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from period slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for period")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Period) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPeriod(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PeriodSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PeriodSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"period\".* FROM \"period\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, periodPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PeriodSlice")
	}

	*o = slice

	return nil
}

// PeriodExists checks if the Period row exists.
func PeriodExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"period\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if period exists")
	}

	return exists, nil
}
