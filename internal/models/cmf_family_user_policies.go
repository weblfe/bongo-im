// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// CMFFamilyUserPolicy is an object representing the database table.
type CMFFamilyUserPolicy struct {
	ID             int           `boil:"id" json:"id" toml:"id" yaml:"id"`
	FamilyID       int64         `boil:"family_id" json:"family_id" toml:"family_id" yaml:"family_id"`
	UID            int64         `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	State          bool          `boil:"state" json:"state" toml:"state" yaml:"state"`
	Proportion     types.Decimal `boil:"proportion" json:"proportion" toml:"proportion" yaml:"proportion"`
	FamilyPolicyID null.Int      `boil:"family_policy_id" json:"family_policy_id,omitempty" toml:"family_policy_id" yaml:"family_policy_id,omitempty"`
	Comment        null.String   `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`
	ModifyUID      int           `boil:"modify_uid" json:"modify_uid" toml:"modify_uid" yaml:"modify_uid"`
	CreatorUID     int           `boil:"creator_uid" json:"creator_uid" toml:"creator_uid" yaml:"creator_uid"`
	CreatedAt      time.Time     `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time     `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *cmfFamilyUserPolicyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfFamilyUserPolicyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFFamilyUserPolicyColumns = struct {
	ID             string
	FamilyID       string
	UID            string
	State          string
	Proportion     string
	FamilyPolicyID string
	Comment        string
	ModifyUID      string
	CreatorUID     string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	FamilyID:       "family_id",
	UID:            "uid",
	State:          "state",
	Proportion:     "proportion",
	FamilyPolicyID: "family_policy_id",
	Comment:        "comment",
	ModifyUID:      "modify_uid",
	CreatorUID:     "creator_uid",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// Generated where

var CMFFamilyUserPolicyWhere = struct {
	ID             whereHelperint
	FamilyID       whereHelperint64
	UID            whereHelperint64
	State          whereHelperbool
	Proportion     whereHelpertypes_Decimal
	FamilyPolicyID whereHelpernull_Int
	Comment        whereHelpernull_String
	ModifyUID      whereHelperint
	CreatorUID     whereHelperint
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
}{
	ID:             whereHelperint{field: "`cmf_family_user_policies`.`id`"},
	FamilyID:       whereHelperint64{field: "`cmf_family_user_policies`.`family_id`"},
	UID:            whereHelperint64{field: "`cmf_family_user_policies`.`uid`"},
	State:          whereHelperbool{field: "`cmf_family_user_policies`.`state`"},
	Proportion:     whereHelpertypes_Decimal{field: "`cmf_family_user_policies`.`proportion`"},
	FamilyPolicyID: whereHelpernull_Int{field: "`cmf_family_user_policies`.`family_policy_id`"},
	Comment:        whereHelpernull_String{field: "`cmf_family_user_policies`.`comment`"},
	ModifyUID:      whereHelperint{field: "`cmf_family_user_policies`.`modify_uid`"},
	CreatorUID:     whereHelperint{field: "`cmf_family_user_policies`.`creator_uid`"},
	CreatedAt:      whereHelpertime_Time{field: "`cmf_family_user_policies`.`created_at`"},
	UpdatedAt:      whereHelpertime_Time{field: "`cmf_family_user_policies`.`updated_at`"},
}

// CMFFamilyUserPolicyRels is where relationship names are stored.
var CMFFamilyUserPolicyRels = struct {
}{}

// cmfFamilyUserPolicyR is where relationships are stored.
type cmfFamilyUserPolicyR struct {
}

// NewStruct creates a new relationship struct
func (*cmfFamilyUserPolicyR) NewStruct() *cmfFamilyUserPolicyR {
	return &cmfFamilyUserPolicyR{}
}

// cmfFamilyUserPolicyL is where Load methods for each relationship are stored.
type cmfFamilyUserPolicyL struct{}

var (
	cmfFamilyUserPolicyAllColumns            = []string{"id", "family_id", "uid", "state", "proportion", "family_policy_id", "comment", "modify_uid", "creator_uid", "created_at", "updated_at"}
	cmfFamilyUserPolicyColumnsWithoutDefault = []string{"family_id", "uid", "family_policy_id", "comment", "created_at", "updated_at"}
	cmfFamilyUserPolicyColumnsWithDefault    = []string{"id", "state", "proportion", "modify_uid", "creator_uid"}
	cmfFamilyUserPolicyPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFFamilyUserPolicySlice is an alias for a slice of pointers to CMFFamilyUserPolicy.
	// This should generally be used opposed to []CMFFamilyUserPolicy.
	CMFFamilyUserPolicySlice []*CMFFamilyUserPolicy
	// CMFFamilyUserPolicyHook is the signature for custom CMFFamilyUserPolicy hook methods
	CMFFamilyUserPolicyHook func(context.Context, boil.ContextExecutor, *CMFFamilyUserPolicy) error

	cmfFamilyUserPolicyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfFamilyUserPolicyType                 = reflect.TypeOf(&CMFFamilyUserPolicy{})
	cmfFamilyUserPolicyMapping              = queries.MakeStructMapping(cmfFamilyUserPolicyType)
	cmfFamilyUserPolicyPrimaryKeyMapping, _ = queries.BindMapping(cmfFamilyUserPolicyType, cmfFamilyUserPolicyMapping, cmfFamilyUserPolicyPrimaryKeyColumns)
	cmfFamilyUserPolicyInsertCacheMut       sync.RWMutex
	cmfFamilyUserPolicyInsertCache          = make(map[string]insertCache)
	cmfFamilyUserPolicyUpdateCacheMut       sync.RWMutex
	cmfFamilyUserPolicyUpdateCache          = make(map[string]updateCache)
	cmfFamilyUserPolicyUpsertCacheMut       sync.RWMutex
	cmfFamilyUserPolicyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfFamilyUserPolicyBeforeInsertHooks []CMFFamilyUserPolicyHook
var cmfFamilyUserPolicyBeforeUpdateHooks []CMFFamilyUserPolicyHook
var cmfFamilyUserPolicyBeforeDeleteHooks []CMFFamilyUserPolicyHook
var cmfFamilyUserPolicyBeforeUpsertHooks []CMFFamilyUserPolicyHook

var cmfFamilyUserPolicyAfterInsertHooks []CMFFamilyUserPolicyHook
var cmfFamilyUserPolicyAfterSelectHooks []CMFFamilyUserPolicyHook
var cmfFamilyUserPolicyAfterUpdateHooks []CMFFamilyUserPolicyHook
var cmfFamilyUserPolicyAfterDeleteHooks []CMFFamilyUserPolicyHook
var cmfFamilyUserPolicyAfterUpsertHooks []CMFFamilyUserPolicyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFFamilyUserPolicy) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFFamilyUserPolicy) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFFamilyUserPolicy) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFFamilyUserPolicy) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFFamilyUserPolicy) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFFamilyUserPolicy) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFFamilyUserPolicy) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFFamilyUserPolicy) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFFamilyUserPolicy) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyUserPolicyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFFamilyUserPolicyHook registers your hook function for all future operations.
func AddCMFFamilyUserPolicyHook(hookPoint boil.HookPoint, cmfFamilyUserPolicyHook CMFFamilyUserPolicyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfFamilyUserPolicyBeforeInsertHooks = append(cmfFamilyUserPolicyBeforeInsertHooks, cmfFamilyUserPolicyHook)
	case boil.BeforeUpdateHook:
		cmfFamilyUserPolicyBeforeUpdateHooks = append(cmfFamilyUserPolicyBeforeUpdateHooks, cmfFamilyUserPolicyHook)
	case boil.BeforeDeleteHook:
		cmfFamilyUserPolicyBeforeDeleteHooks = append(cmfFamilyUserPolicyBeforeDeleteHooks, cmfFamilyUserPolicyHook)
	case boil.BeforeUpsertHook:
		cmfFamilyUserPolicyBeforeUpsertHooks = append(cmfFamilyUserPolicyBeforeUpsertHooks, cmfFamilyUserPolicyHook)
	case boil.AfterInsertHook:
		cmfFamilyUserPolicyAfterInsertHooks = append(cmfFamilyUserPolicyAfterInsertHooks, cmfFamilyUserPolicyHook)
	case boil.AfterSelectHook:
		cmfFamilyUserPolicyAfterSelectHooks = append(cmfFamilyUserPolicyAfterSelectHooks, cmfFamilyUserPolicyHook)
	case boil.AfterUpdateHook:
		cmfFamilyUserPolicyAfterUpdateHooks = append(cmfFamilyUserPolicyAfterUpdateHooks, cmfFamilyUserPolicyHook)
	case boil.AfterDeleteHook:
		cmfFamilyUserPolicyAfterDeleteHooks = append(cmfFamilyUserPolicyAfterDeleteHooks, cmfFamilyUserPolicyHook)
	case boil.AfterUpsertHook:
		cmfFamilyUserPolicyAfterUpsertHooks = append(cmfFamilyUserPolicyAfterUpsertHooks, cmfFamilyUserPolicyHook)
	}
}

// One returns a single cmfFamilyUserPolicy record from the query.
func (q cmfFamilyUserPolicyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFFamilyUserPolicy, error) {
	o := &CMFFamilyUserPolicy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_family_user_policies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFFamilyUserPolicy records from the query.
func (q cmfFamilyUserPolicyQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFFamilyUserPolicySlice, error) {
	var o []*CMFFamilyUserPolicy

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFFamilyUserPolicy slice")
	}

	if len(cmfFamilyUserPolicyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFFamilyUserPolicy records in the query.
func (q cmfFamilyUserPolicyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_family_user_policies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfFamilyUserPolicyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_family_user_policies exists")
	}

	return count > 0, nil
}

// CMFFamilyUserPolicies retrieves all the records using an executor.
func CMFFamilyUserPolicies(mods ...qm.QueryMod) cmfFamilyUserPolicyQuery {
	mods = append(mods, qm.From("`cmf_family_user_policies`"))
	return cmfFamilyUserPolicyQuery{NewQuery(mods...)}
}

// FindCMFFamilyUserPolicy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFFamilyUserPolicy(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFFamilyUserPolicy, error) {
	cmfFamilyUserPolicyObj := &CMFFamilyUserPolicy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_family_user_policies` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfFamilyUserPolicyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_family_user_policies")
	}

	return cmfFamilyUserPolicyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFFamilyUserPolicy) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_family_user_policies provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfFamilyUserPolicyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfFamilyUserPolicyInsertCacheMut.RLock()
	cache, cached := cmfFamilyUserPolicyInsertCache[key]
	cmfFamilyUserPolicyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfFamilyUserPolicyAllColumns,
			cmfFamilyUserPolicyColumnsWithDefault,
			cmfFamilyUserPolicyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfFamilyUserPolicyType, cmfFamilyUserPolicyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfFamilyUserPolicyType, cmfFamilyUserPolicyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_family_user_policies` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_family_user_policies` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_family_user_policies` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfFamilyUserPolicyPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into cmf_family_user_policies")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfFamilyUserPolicyMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_family_user_policies")
	}

CacheNoHooks:
	if !cached {
		cmfFamilyUserPolicyInsertCacheMut.Lock()
		cmfFamilyUserPolicyInsertCache[key] = cache
		cmfFamilyUserPolicyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFFamilyUserPolicy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFFamilyUserPolicy) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfFamilyUserPolicyUpdateCacheMut.RLock()
	cache, cached := cmfFamilyUserPolicyUpdateCache[key]
	cmfFamilyUserPolicyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfFamilyUserPolicyAllColumns,
			cmfFamilyUserPolicyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_family_user_policies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_family_user_policies` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfFamilyUserPolicyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfFamilyUserPolicyType, cmfFamilyUserPolicyMapping, append(wl, cmfFamilyUserPolicyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_family_user_policies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_family_user_policies")
	}

	if !cached {
		cmfFamilyUserPolicyUpdateCacheMut.Lock()
		cmfFamilyUserPolicyUpdateCache[key] = cache
		cmfFamilyUserPolicyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfFamilyUserPolicyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_family_user_policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_family_user_policies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFFamilyUserPolicySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamilyUserPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_family_user_policies` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamilyUserPolicyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfFamilyUserPolicy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfFamilyUserPolicy")
	}
	return rowsAff, nil
}

var mySQLCMFFamilyUserPolicyUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFFamilyUserPolicy) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_family_user_policies provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfFamilyUserPolicyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFFamilyUserPolicyUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	cmfFamilyUserPolicyUpsertCacheMut.RLock()
	cache, cached := cmfFamilyUserPolicyUpsertCache[key]
	cmfFamilyUserPolicyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfFamilyUserPolicyAllColumns,
			cmfFamilyUserPolicyColumnsWithDefault,
			cmfFamilyUserPolicyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfFamilyUserPolicyAllColumns,
			cmfFamilyUserPolicyPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_family_user_policies, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_family_user_policies`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_family_user_policies` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfFamilyUserPolicyType, cmfFamilyUserPolicyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfFamilyUserPolicyType, cmfFamilyUserPolicyMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for cmf_family_user_policies")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfFamilyUserPolicyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfFamilyUserPolicyType, cmfFamilyUserPolicyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_family_user_policies")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_family_user_policies")
	}

CacheNoHooks:
	if !cached {
		cmfFamilyUserPolicyUpsertCacheMut.Lock()
		cmfFamilyUserPolicyUpsertCache[key] = cache
		cmfFamilyUserPolicyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFFamilyUserPolicy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFFamilyUserPolicy) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFFamilyUserPolicy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfFamilyUserPolicyPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_family_user_policies` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_family_user_policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_family_user_policies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfFamilyUserPolicyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfFamilyUserPolicyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_family_user_policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_family_user_policies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFFamilyUserPolicySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfFamilyUserPolicyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamilyUserPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_family_user_policies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamilyUserPolicyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfFamilyUserPolicy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_family_user_policies")
	}

	if len(cmfFamilyUserPolicyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CMFFamilyUserPolicy) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFFamilyUserPolicy(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFFamilyUserPolicySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFFamilyUserPolicySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamilyUserPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_family_user_policies`.* FROM `cmf_family_user_policies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamilyUserPolicyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFFamilyUserPolicySlice")
	}

	*o = slice

	return nil
}

// CMFFamilyUserPolicyExists checks if the CMFFamilyUserPolicy row exists.
func CMFFamilyUserPolicyExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_family_user_policies` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_family_user_policies exists")
	}

	return exists, nil
}
