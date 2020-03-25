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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// User is an object representing the database table.
type User struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Email         string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	FirstName     string      `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	PasswordHash  string      `boil:"password_hash" json:"password_hash" toml:"password_hash" yaml:"password_hash"`
	PasswordSalt  string      `boil:"password_salt" json:"password_salt" toml:"password_salt" yaml:"password_salt"`
	PasswordReset null.String `boil:"password_reset" json:"password_reset,omitempty" toml:"password_reset" yaml:"password_reset,omitempty"`
	Timezone      null.String `boil:"timezone" json:"timezone,omitempty" toml:"timezone" yaml:"timezone,omitempty"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ArchivedAt    null.Time   `boil:"archived_at" json:"archived_at,omitempty" toml:"archived_at" yaml:"archived_at,omitempty"`
	LastName      string      `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	Phone         string      `boil:"phone" json:"phone" toml:"phone" yaml:"phone"`

	R *userR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserColumns = struct {
	ID            string
	Email         string
	FirstName     string
	PasswordHash  string
	PasswordSalt  string
	PasswordReset string
	Timezone      string
	CreatedAt     string
	UpdatedAt     string
	ArchivedAt    string
	LastName      string
	Phone         string
}{
	ID:            "id",
	Email:         "email",
	FirstName:     "first_name",
	PasswordHash:  "password_hash",
	PasswordSalt:  "password_salt",
	PasswordReset: "password_reset",
	Timezone:      "timezone",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	ArchivedAt:    "archived_at",
	LastName:      "last_name",
	Phone:         "phone",
}

// Generated where

var UserWhere = struct {
	ID            whereHelperstring
	Email         whereHelperstring
	FirstName     whereHelperstring
	PasswordHash  whereHelperstring
	PasswordSalt  whereHelperstring
	PasswordReset whereHelpernull_String
	Timezone      whereHelpernull_String
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpernull_Time
	ArchivedAt    whereHelpernull_Time
	LastName      whereHelperstring
	Phone         whereHelperstring
}{
	ID:            whereHelperstring{field: "\"users\".\"id\""},
	Email:         whereHelperstring{field: "\"users\".\"email\""},
	FirstName:     whereHelperstring{field: "\"users\".\"first_name\""},
	PasswordHash:  whereHelperstring{field: "\"users\".\"password_hash\""},
	PasswordSalt:  whereHelperstring{field: "\"users\".\"password_salt\""},
	PasswordReset: whereHelpernull_String{field: "\"users\".\"password_reset\""},
	Timezone:      whereHelpernull_String{field: "\"users\".\"timezone\""},
	CreatedAt:     whereHelpertime_Time{field: "\"users\".\"created_at\""},
	UpdatedAt:     whereHelpernull_Time{field: "\"users\".\"updated_at\""},
	ArchivedAt:    whereHelpernull_Time{field: "\"users\".\"archived_at\""},
	LastName:      whereHelperstring{field: "\"users\".\"last_name\""},
	Phone:         whereHelperstring{field: "\"users\".\"phone\""},
}

// UserRels is where relationship names are stored.
var UserRels = struct {
	BillingUserAccounts string
	SignupUserAccounts  string
}{
	BillingUserAccounts: "BillingUserAccounts",
	SignupUserAccounts:  "SignupUserAccounts",
}

// userR is where relationships are stored.
type userR struct {
	BillingUserAccounts AccountSlice
	SignupUserAccounts  AccountSlice
}

// NewStruct creates a new relationship struct
func (*userR) NewStruct() *userR {
	return &userR{}
}

// userL is where Load methods for each relationship are stored.
type userL struct{}

var (
	userAllColumns            = []string{"id", "email", "first_name", "password_hash", "password_salt", "password_reset", "timezone", "created_at", "updated_at", "archived_at", "last_name", "phone"}
	userColumnsWithoutDefault = []string{"id", "email", "password_hash", "password_salt", "timezone", "created_at", "updated_at", "archived_at"}
	userColumnsWithDefault    = []string{"first_name", "password_reset", "last_name", "phone"}
	userPrimaryKeyColumns     = []string{"id"}
)

type (
	// UserSlice is an alias for a slice of pointers to User.
	// This should generally be used opposed to []User.
	UserSlice []*User

	userQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userType                 = reflect.TypeOf(&User{})
	userMapping              = queries.MakeStructMapping(userType)
	userPrimaryKeyMapping, _ = queries.BindMapping(userType, userMapping, userPrimaryKeyColumns)
	userInsertCacheMut       sync.RWMutex
	userInsertCache          = make(map[string]insertCache)
	userUpdateCacheMut       sync.RWMutex
	userUpdateCache          = make(map[string]updateCache)
	userUpsertCacheMut       sync.RWMutex
	userUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single user record from the query.
func (q userQuery) One(ctx context.Context, exec boil.ContextExecutor) (*User, error) {
	o := &User{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for users")
	}

	return o, nil
}

// All returns all User records from the query.
func (q userQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserSlice, error) {
	var o []*User

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to User slice")
	}

	return o, nil
}

// Count returns the count of all User records in the query.
func (q userQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if users exists")
	}

	return count > 0, nil
}

// BillingUserAccounts retrieves all the account's Accounts with an executor via billing_user_id column.
func (o *User) BillingUserAccounts(mods ...qm.QueryMod) accountQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"accounts\".\"billing_user_id\"=?", o.ID),
	)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "\"accounts\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"accounts\".*"})
	}

	return query
}

// SignupUserAccounts retrieves all the account's Accounts with an executor via signup_user_id column.
func (o *User) SignupUserAccounts(mods ...qm.QueryMod) accountQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"accounts\".\"signup_user_id\"=?", o.ID),
	)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "\"accounts\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"accounts\".*"})
	}

	return query
}

// LoadBillingUserAccounts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (userL) LoadBillingUserAccounts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUser interface{}, mods queries.Applicator) error {
	var slice []*User
	var object *User

	if singular {
		object = maybeUser.(*User)
	} else {
		slice = *maybeUser.(*[]*User)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userR{}
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

	query := NewQuery(qm.From(`accounts`), qm.WhereIn(`accounts.billing_user_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load accounts")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice accounts")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on accounts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for accounts")
	}

	if singular {
		object.R.BillingUserAccounts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &accountR{}
			}
			foreign.R.BillingUser = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.BillingUserID) {
				local.R.BillingUserAccounts = append(local.R.BillingUserAccounts, foreign)
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.BillingUser = local
				break
			}
		}
	}

	return nil
}

// LoadSignupUserAccounts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (userL) LoadSignupUserAccounts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUser interface{}, mods queries.Applicator) error {
	var slice []*User
	var object *User

	if singular {
		object = maybeUser.(*User)
	} else {
		slice = *maybeUser.(*[]*User)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userR{}
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

	query := NewQuery(qm.From(`accounts`), qm.WhereIn(`accounts.signup_user_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load accounts")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice accounts")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on accounts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for accounts")
	}

	if singular {
		object.R.SignupUserAccounts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &accountR{}
			}
			foreign.R.SignupUser = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.SignupUserID) {
				local.R.SignupUserAccounts = append(local.R.SignupUserAccounts, foreign)
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.SignupUser = local
				break
			}
		}
	}

	return nil
}

// AddBillingUserAccounts adds the given related objects to the existing relationships
// of the user, optionally inserting them as new records.
// Appends related to o.R.BillingUserAccounts.
// Sets related.R.BillingUser appropriately.
func (o *User) AddBillingUserAccounts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Account) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.BillingUserID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"accounts\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"billing_user_id"}),
				strmangle.WhereClause("\"", "\"", 2, accountPrimaryKeyColumns),
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

			queries.Assign(&rel.BillingUserID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &userR{
			BillingUserAccounts: related,
		}
	} else {
		o.R.BillingUserAccounts = append(o.R.BillingUserAccounts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &accountR{
				BillingUser: o,
			}
		} else {
			rel.R.BillingUser = o
		}
	}
	return nil
}

// SetBillingUserAccounts removes all previously related items of the
// user replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.BillingUser's BillingUserAccounts accordingly.
// Replaces o.R.BillingUserAccounts with related.
// Sets related.R.BillingUser's BillingUserAccounts accordingly.
func (o *User) SetBillingUserAccounts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Account) error {
	query := "update \"accounts\" set \"billing_user_id\" = null where \"billing_user_id\" = $1"
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
		for _, rel := range o.R.BillingUserAccounts {
			queries.SetScanner(&rel.BillingUserID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.BillingUser = nil
		}

		o.R.BillingUserAccounts = nil
	}
	return o.AddBillingUserAccounts(ctx, exec, insert, related...)
}

// RemoveBillingUserAccounts relationships from objects passed in.
// Removes related items from R.BillingUserAccounts (uses pointer comparison, removal does not keep order)
// Sets related.R.BillingUser.
func (o *User) RemoveBillingUserAccounts(ctx context.Context, exec boil.ContextExecutor, related ...*Account) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.BillingUserID, nil)
		if rel.R != nil {
			rel.R.BillingUser = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("billing_user_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.BillingUserAccounts {
			if rel != ri {
				continue
			}

			ln := len(o.R.BillingUserAccounts)
			if ln > 1 && i < ln-1 {
				o.R.BillingUserAccounts[i] = o.R.BillingUserAccounts[ln-1]
			}
			o.R.BillingUserAccounts = o.R.BillingUserAccounts[:ln-1]
			break
		}
	}

	return nil
}

// AddSignupUserAccounts adds the given related objects to the existing relationships
// of the user, optionally inserting them as new records.
// Appends related to o.R.SignupUserAccounts.
// Sets related.R.SignupUser appropriately.
func (o *User) AddSignupUserAccounts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Account) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.SignupUserID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"accounts\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"signup_user_id"}),
				strmangle.WhereClause("\"", "\"", 2, accountPrimaryKeyColumns),
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

			queries.Assign(&rel.SignupUserID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &userR{
			SignupUserAccounts: related,
		}
	} else {
		o.R.SignupUserAccounts = append(o.R.SignupUserAccounts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &accountR{
				SignupUser: o,
			}
		} else {
			rel.R.SignupUser = o
		}
	}
	return nil
}

// SetSignupUserAccounts removes all previously related items of the
// user replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.SignupUser's SignupUserAccounts accordingly.
// Replaces o.R.SignupUserAccounts with related.
// Sets related.R.SignupUser's SignupUserAccounts accordingly.
func (o *User) SetSignupUserAccounts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Account) error {
	query := "update \"accounts\" set \"signup_user_id\" = null where \"signup_user_id\" = $1"
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
		for _, rel := range o.R.SignupUserAccounts {
			queries.SetScanner(&rel.SignupUserID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.SignupUser = nil
		}

		o.R.SignupUserAccounts = nil
	}
	return o.AddSignupUserAccounts(ctx, exec, insert, related...)
}

// RemoveSignupUserAccounts relationships from objects passed in.
// Removes related items from R.SignupUserAccounts (uses pointer comparison, removal does not keep order)
// Sets related.R.SignupUser.
func (o *User) RemoveSignupUserAccounts(ctx context.Context, exec boil.ContextExecutor, related ...*Account) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.SignupUserID, nil)
		if rel.R != nil {
			rel.R.SignupUser = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("signup_user_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.SignupUserAccounts {
			if rel != ri {
				continue
			}

			ln := len(o.R.SignupUserAccounts)
			if ln > 1 && i < ln-1 {
				o.R.SignupUserAccounts[i] = o.R.SignupUserAccounts[ln-1]
			}
			o.R.SignupUserAccounts = o.R.SignupUserAccounts[:ln-1]
			break
		}
	}

	return nil
}

// Users retrieves all the records using an executor.
func Users(mods ...qm.QueryMod) userQuery {
	mods = append(mods, qm.From("\"users\""))
	return userQuery{NewQuery(mods...)}
}

// FindUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUser(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*User, error) {
	userObj := &User{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from users")
	}

	return userObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *User) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no users provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(userColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userInsertCacheMut.RLock()
	cache, cached := userInsertCache[key]
	userInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userAllColumns,
			userColumnsWithDefault,
			userColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userType, userMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userType, userMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into users")
	}

	if !cached {
		userInsertCacheMut.Lock()
		userInsertCache[key] = cache
		userInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the User.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *User) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	userUpdateCacheMut.RLock()
	cache, cached := userUpdateCache[key]
	userUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userAllColumns,
			userPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userType, userMapping, append(wl, userPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for users")
	}

	if !cached {
		userUpdateCacheMut.Lock()
		userUpdateCache[key] = cache
		userUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q userQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in user slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all user")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *User) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no users provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(userColumnsWithDefault, o)

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

	userUpsertCacheMut.RLock()
	cache, cached := userUpsertCache[key]
	userUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userAllColumns,
			userColumnsWithDefault,
			userColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userAllColumns,
			userPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userPrimaryKeyColumns))
			copy(conflict, userPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"users\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userType, userMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userType, userMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert users")
	}

	if !cached {
		userUpsertCacheMut.Lock()
		userUpsertCache[key] = cache
		userUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single User record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *User) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no User provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userPrimaryKeyMapping)
	sql := "DELETE FROM \"users\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for users")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *User) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"users\".* FROM \"users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserSlice")
	}

	*o = slice

	return nil
}

// UserExists checks if the User row exists.
func UserExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"users\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if users exists")
	}

	return exists, nil
}
