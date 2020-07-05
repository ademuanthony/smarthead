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

// Account is an object representing the database table.
type Account struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name          string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Address1      string      `boil:"address1" json:"address1" toml:"address1" yaml:"address1"`
	Address2      string      `boil:"address2" json:"address2" toml:"address2" yaml:"address2"`
	City          string      `boil:"city" json:"city" toml:"city" yaml:"city"`
	Region        string      `boil:"region" json:"region" toml:"region" yaml:"region"`
	Country       string      `boil:"country" json:"country" toml:"country" yaml:"country"`
	Zipcode       string      `boil:"zipcode" json:"zipcode" toml:"zipcode" yaml:"zipcode"`
	Status        string      `boil:"status" json:"status" toml:"status" yaml:"status"`
	Timezone      string      `boil:"timezone" json:"timezone" toml:"timezone" yaml:"timezone"`
	SignupUserID  null.String `boil:"signup_user_id" json:"signup_user_id,omitempty" toml:"signup_user_id" yaml:"signup_user_id,omitempty"`
	BillingUserID null.String `boil:"billing_user_id" json:"billing_user_id,omitempty" toml:"billing_user_id" yaml:"billing_user_id,omitempty"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ArchivedAt    null.Time   `boil:"archived_at" json:"archived_at,omitempty" toml:"archived_at" yaml:"archived_at,omitempty"`

	R *accountR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accountL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccountColumns = struct {
	ID            string
	Name          string
	Address1      string
	Address2      string
	City          string
	Region        string
	Country       string
	Zipcode       string
	Status        string
	Timezone      string
	SignupUserID  string
	BillingUserID string
	CreatedAt     string
	UpdatedAt     string
	ArchivedAt    string
}{
	ID:            "id",
	Name:          "name",
	Address1:      "address1",
	Address2:      "address2",
	City:          "city",
	Region:        "region",
	Country:       "country",
	Zipcode:       "zipcode",
	Status:        "status",
	Timezone:      "timezone",
	SignupUserID:  "signup_user_id",
	BillingUserID: "billing_user_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	ArchivedAt:    "archived_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AccountWhere = struct {
	ID            whereHelperstring
	Name          whereHelperstring
	Address1      whereHelperstring
	Address2      whereHelperstring
	City          whereHelperstring
	Region        whereHelperstring
	Country       whereHelperstring
	Zipcode       whereHelperstring
	Status        whereHelperstring
	Timezone      whereHelperstring
	SignupUserID  whereHelpernull_String
	BillingUserID whereHelpernull_String
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpernull_Time
	ArchivedAt    whereHelpernull_Time
}{
	ID:            whereHelperstring{field: "\"accounts\".\"id\""},
	Name:          whereHelperstring{field: "\"accounts\".\"name\""},
	Address1:      whereHelperstring{field: "\"accounts\".\"address1\""},
	Address2:      whereHelperstring{field: "\"accounts\".\"address2\""},
	City:          whereHelperstring{field: "\"accounts\".\"city\""},
	Region:        whereHelperstring{field: "\"accounts\".\"region\""},
	Country:       whereHelperstring{field: "\"accounts\".\"country\""},
	Zipcode:       whereHelperstring{field: "\"accounts\".\"zipcode\""},
	Status:        whereHelperstring{field: "\"accounts\".\"status\""},
	Timezone:      whereHelperstring{field: "\"accounts\".\"timezone\""},
	SignupUserID:  whereHelpernull_String{field: "\"accounts\".\"signup_user_id\""},
	BillingUserID: whereHelpernull_String{field: "\"accounts\".\"billing_user_id\""},
	CreatedAt:     whereHelpertime_Time{field: "\"accounts\".\"created_at\""},
	UpdatedAt:     whereHelpernull_Time{field: "\"accounts\".\"updated_at\""},
	ArchivedAt:    whereHelpernull_Time{field: "\"accounts\".\"archived_at\""},
}

// AccountRels is where relationship names are stored.
var AccountRels = struct {
	BillingUser   string
	SignupUser    string
	UsersAccounts string
}{
	BillingUser:   "BillingUser",
	SignupUser:    "SignupUser",
	UsersAccounts: "UsersAccounts",
}

// accountR is where relationships are stored.
type accountR struct {
	BillingUser   *User
	SignupUser    *User
	UsersAccounts UsersAccountSlice
}

// NewStruct creates a new relationship struct
func (*accountR) NewStruct() *accountR {
	return &accountR{}
}

// accountL is where Load methods for each relationship are stored.
type accountL struct{}

var (
	accountAllColumns            = []string{"id", "name", "address1", "address2", "city", "region", "country", "zipcode", "status", "timezone", "signup_user_id", "billing_user_id", "created_at", "updated_at", "archived_at"}
	accountColumnsWithoutDefault = []string{"id", "name", "created_at", "updated_at", "archived_at"}
	accountColumnsWithDefault    = []string{"address1", "address2", "city", "region", "country", "zipcode", "status", "timezone", "signup_user_id", "billing_user_id"}
	accountPrimaryKeyColumns     = []string{"id"}
)

type (
	// AccountSlice is an alias for a slice of pointers to Account.
	// This should generally be used opposed to []Account.
	AccountSlice []*Account

	accountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accountType                 = reflect.TypeOf(&Account{})
	accountMapping              = queries.MakeStructMapping(accountType)
	accountPrimaryKeyMapping, _ = queries.BindMapping(accountType, accountMapping, accountPrimaryKeyColumns)
	accountInsertCacheMut       sync.RWMutex
	accountInsertCache          = make(map[string]insertCache)
	accountUpdateCacheMut       sync.RWMutex
	accountUpdateCache          = make(map[string]updateCache)
	accountUpsertCacheMut       sync.RWMutex
	accountUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single account record from the query.
func (q accountQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Account, error) {
	o := &Account{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for accounts")
	}

	return o, nil
}

// All returns all Account records from the query.
func (q accountQuery) All(ctx context.Context, exec boil.ContextExecutor) (AccountSlice, error) {
	var o []*Account

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Account slice")
	}

	return o, nil
}

// Count returns the count of all Account records in the query.
func (q accountQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count accounts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q accountQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if accounts exists")
	}

	return count > 0, nil
}

// BillingUser pointed to by the foreign key.
func (o *Account) BillingUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BillingUserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// SignupUser pointed to by the foreign key.
func (o *Account) SignupUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SignupUserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// UsersAccounts retrieves all the users_account's UsersAccounts with an executor.
func (o *Account) UsersAccounts(mods ...qm.QueryMod) usersAccountQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"users_accounts\".\"account_id\"=?", o.ID),
	)

	query := UsersAccounts(queryMods...)
	queries.SetFrom(query.Query, "\"users_accounts\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"users_accounts\".*"})
	}

	return query
}

// LoadBillingUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (accountL) LoadBillingUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAccount interface{}, mods queries.Applicator) error {
	var slice []*Account
	var object *Account

	if singular {
		object = maybeAccount.(*Account)
	} else {
		slice = *maybeAccount.(*[]*Account)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &accountR{}
		}
		if !queries.IsNil(object.BillingUserID) {
			args = append(args, object.BillingUserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &accountR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.BillingUserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.BillingUserID) {
				args = append(args, obj.BillingUserID)
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
		object.R.BillingUser = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.BillingUserAccounts = append(foreign.R.BillingUserAccounts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.BillingUserID, foreign.ID) {
				local.R.BillingUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.BillingUserAccounts = append(foreign.R.BillingUserAccounts, local)
				break
			}
		}
	}

	return nil
}

// LoadSignupUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (accountL) LoadSignupUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAccount interface{}, mods queries.Applicator) error {
	var slice []*Account
	var object *Account

	if singular {
		object = maybeAccount.(*Account)
	} else {
		slice = *maybeAccount.(*[]*Account)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &accountR{}
		}
		if !queries.IsNil(object.SignupUserID) {
			args = append(args, object.SignupUserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &accountR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.SignupUserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.SignupUserID) {
				args = append(args, obj.SignupUserID)
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
		object.R.SignupUser = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.SignupUserAccounts = append(foreign.R.SignupUserAccounts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SignupUserID, foreign.ID) {
				local.R.SignupUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.SignupUserAccounts = append(foreign.R.SignupUserAccounts, local)
				break
			}
		}
	}

	return nil
}

// LoadUsersAccounts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (accountL) LoadUsersAccounts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAccount interface{}, mods queries.Applicator) error {
	var slice []*Account
	var object *Account

	if singular {
		object = maybeAccount.(*Account)
	} else {
		slice = *maybeAccount.(*[]*Account)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &accountR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &accountR{}
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

	query := NewQuery(qm.From(`users_accounts`), qm.WhereIn(`users_accounts.account_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users_accounts")
	}

	var resultSlice []*UsersAccount
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice users_accounts")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on users_accounts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users_accounts")
	}

	if singular {
		object.R.UsersAccounts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &usersAccountR{}
			}
			foreign.R.Account = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.AccountID {
				local.R.UsersAccounts = append(local.R.UsersAccounts, foreign)
				if foreign.R == nil {
					foreign.R = &usersAccountR{}
				}
				foreign.R.Account = local
				break
			}
		}
	}

	return nil
}

// SetBillingUser of the account to the related item.
// Sets o.R.BillingUser to related.
// Adds o to related.R.BillingUserAccounts.
func (o *Account) SetBillingUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"accounts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"billing_user_id"}),
		strmangle.WhereClause("\"", "\"", 2, accountPrimaryKeyColumns),
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

	queries.Assign(&o.BillingUserID, related.ID)
	if o.R == nil {
		o.R = &accountR{
			BillingUser: related,
		}
	} else {
		o.R.BillingUser = related
	}

	if related.R == nil {
		related.R = &userR{
			BillingUserAccounts: AccountSlice{o},
		}
	} else {
		related.R.BillingUserAccounts = append(related.R.BillingUserAccounts, o)
	}

	return nil
}

// RemoveBillingUser relationship.
// Sets o.R.BillingUser to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Account) RemoveBillingUser(ctx context.Context, exec boil.ContextExecutor, related *User) error {
	var err error

	queries.SetScanner(&o.BillingUserID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("billing_user_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.BillingUser = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.BillingUserAccounts {
		if queries.Equal(o.BillingUserID, ri.BillingUserID) {
			continue
		}

		ln := len(related.R.BillingUserAccounts)
		if ln > 1 && i < ln-1 {
			related.R.BillingUserAccounts[i] = related.R.BillingUserAccounts[ln-1]
		}
		related.R.BillingUserAccounts = related.R.BillingUserAccounts[:ln-1]
		break
	}
	return nil
}

// SetSignupUser of the account to the related item.
// Sets o.R.SignupUser to related.
// Adds o to related.R.SignupUserAccounts.
func (o *Account) SetSignupUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"accounts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"signup_user_id"}),
		strmangle.WhereClause("\"", "\"", 2, accountPrimaryKeyColumns),
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

	queries.Assign(&o.SignupUserID, related.ID)
	if o.R == nil {
		o.R = &accountR{
			SignupUser: related,
		}
	} else {
		o.R.SignupUser = related
	}

	if related.R == nil {
		related.R = &userR{
			SignupUserAccounts: AccountSlice{o},
		}
	} else {
		related.R.SignupUserAccounts = append(related.R.SignupUserAccounts, o)
	}

	return nil
}

// RemoveSignupUser relationship.
// Sets o.R.SignupUser to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Account) RemoveSignupUser(ctx context.Context, exec boil.ContextExecutor, related *User) error {
	var err error

	queries.SetScanner(&o.SignupUserID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("signup_user_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.SignupUser = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.SignupUserAccounts {
		if queries.Equal(o.SignupUserID, ri.SignupUserID) {
			continue
		}

		ln := len(related.R.SignupUserAccounts)
		if ln > 1 && i < ln-1 {
			related.R.SignupUserAccounts[i] = related.R.SignupUserAccounts[ln-1]
		}
		related.R.SignupUserAccounts = related.R.SignupUserAccounts[:ln-1]
		break
	}
	return nil
}

// AddUsersAccounts adds the given related objects to the existing relationships
// of the account, optionally inserting them as new records.
// Appends related to o.R.UsersAccounts.
// Sets related.R.Account appropriately.
func (o *Account) AddUsersAccounts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UsersAccount) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.AccountID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"users_accounts\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"account_id"}),
				strmangle.WhereClause("\"", "\"", 2, usersAccountPrimaryKeyColumns),
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

			rel.AccountID = o.ID
		}
	}

	if o.R == nil {
		o.R = &accountR{
			UsersAccounts: related,
		}
	} else {
		o.R.UsersAccounts = append(o.R.UsersAccounts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &usersAccountR{
				Account: o,
			}
		} else {
			rel.R.Account = o
		}
	}
	return nil
}

// Accounts retrieves all the records using an executor.
func Accounts(mods ...qm.QueryMod) accountQuery {
	mods = append(mods, qm.From("\"accounts\""))
	return accountQuery{NewQuery(mods...)}
}

// FindAccount retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccount(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Account, error) {
	accountObj := &Account{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"accounts\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, accountObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from accounts")
	}

	return accountObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Account) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no accounts provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(accountColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accountInsertCacheMut.RLock()
	cache, cached := accountInsertCache[key]
	accountInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accountAllColumns,
			accountColumnsWithDefault,
			accountColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accountType, accountMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"accounts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"accounts\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into accounts")
	}

	if !cached {
		accountInsertCacheMut.Lock()
		accountInsertCache[key] = cache
		accountInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Account.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Account) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	accountUpdateCacheMut.RLock()
	cache, cached := accountUpdateCache[key]
	accountUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accountAllColumns,
			accountPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update accounts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"accounts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, accountPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, append(wl, accountPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update accounts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for accounts")
	}

	if !cached {
		accountUpdateCacheMut.Lock()
		accountUpdateCache[key] = cache
		accountUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q accountQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for accounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for accounts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccountSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"accounts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, accountPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in account slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all account")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Account) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no accounts provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(accountColumnsWithDefault, o)

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

	accountUpsertCacheMut.RLock()
	cache, cached := accountUpsertCache[key]
	accountUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			accountAllColumns,
			accountColumnsWithDefault,
			accountColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			accountAllColumns,
			accountPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert accounts, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(accountPrimaryKeyColumns))
			copy(conflict, accountPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"accounts\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accountType, accountMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert accounts")
	}

	if !cached {
		accountUpsertCacheMut.Lock()
		accountUpsertCache[key] = cache
		accountUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Account record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Account) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Account provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accountPrimaryKeyMapping)
	sql := "DELETE FROM \"accounts\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from accounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for accounts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q accountQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no accountQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from accounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for accounts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccountSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"accounts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accountPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from account slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for accounts")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Account) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAccount(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccountSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"accounts\".* FROM \"accounts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accountPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AccountSlice")
	}

	*o = slice

	return nil
}

// AccountExists checks if the Account row exists.
func AccountExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"accounts\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if accounts exists")
	}

	return exists, nil
}
