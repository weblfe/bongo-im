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

// CMFPlatformPolicy is an object representing the database table.
type CMFPlatformPolicy struct {
	ID         int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name       string            `boil:"name" json:"name" toml:"name" yaml:"name"`
	Key        string            `boil:"key" json:"key" toml:"key" yaml:"key"`
	Type       int16             `boil:"type" json:"type" toml:"type" yaml:"type"`
	State      bool              `boil:"state" json:"state" toml:"state" yaml:"state"`
	Proportion types.Decimal     `boil:"proportion" json:"proportion" toml:"proportion" yaml:"proportion"`
	IncomeMin  types.NullDecimal `boil:"income_min" json:"income_min,omitempty" toml:"income_min" yaml:"income_min,omitempty"`
	IncomeMax  types.NullDecimal `boil:"income_max" json:"income_max,omitempty" toml:"income_max" yaml:"income_max,omitempty"`
	Comment    null.String       `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`
	ModifyUID  int64             `boil:"modify_uid" json:"modify_uid" toml:"modify_uid" yaml:"modify_uid"`
	CreatedAt  time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time         `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *cmfPlatformPolicyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfPlatformPolicyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFPlatformPolicyColumns = struct {
	ID         string
	Name       string
	Key        string
	Type       string
	State      string
	Proportion string
	IncomeMin  string
	IncomeMax  string
	Comment    string
	ModifyUID  string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	Name:       "name",
	Key:        "key",
	Type:       "type",
	State:      "state",
	Proportion: "proportion",
	IncomeMin:  "income_min",
	IncomeMax:  "income_max",
	Comment:    "comment",
	ModifyUID:  "modify_uid",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// Generated where

var CMFPlatformPolicyWhere = struct {
	ID         whereHelperint
	Name       whereHelperstring
	Key        whereHelperstring
	Type       whereHelperint16
	State      whereHelperbool
	Proportion whereHelpertypes_Decimal
	IncomeMin  whereHelpertypes_NullDecimal
	IncomeMax  whereHelpertypes_NullDecimal
	Comment    whereHelpernull_String
	ModifyUID  whereHelperint64
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperint{field: "`cmf_platform_policy`.`id`"},
	Name:       whereHelperstring{field: "`cmf_platform_policy`.`name`"},
	Key:        whereHelperstring{field: "`cmf_platform_policy`.`key`"},
	Type:       whereHelperint16{field: "`cmf_platform_policy`.`type`"},
	State:      whereHelperbool{field: "`cmf_platform_policy`.`state`"},
	Proportion: whereHelpertypes_Decimal{field: "`cmf_platform_policy`.`proportion`"},
	IncomeMin:  whereHelpertypes_NullDecimal{field: "`cmf_platform_policy`.`income_min`"},
	IncomeMax:  whereHelpertypes_NullDecimal{field: "`cmf_platform_policy`.`income_max`"},
	Comment:    whereHelpernull_String{field: "`cmf_platform_policy`.`comment`"},
	ModifyUID:  whereHelperint64{field: "`cmf_platform_policy`.`modify_uid`"},
	CreatedAt:  whereHelpertime_Time{field: "`cmf_platform_policy`.`created_at`"},
	UpdatedAt:  whereHelpertime_Time{field: "`cmf_platform_policy`.`updated_at`"},
}

// CMFPlatformPolicyRels is where relationship names are stored.
var CMFPlatformPolicyRels = struct {
}{}

// cmfPlatformPolicyR is where relationships are stored.
type cmfPlatformPolicyR struct {
}

// NewStruct creates a new relationship struct
func (*cmfPlatformPolicyR) NewStruct() *cmfPlatformPolicyR {
	return &cmfPlatformPolicyR{}
}

// cmfPlatformPolicyL is where Load methods for each relationship are stored.
type cmfPlatformPolicyL struct{}

var (
	cmfPlatformPolicyAllColumns            = []string{"id", "name", "key", "type", "state", "proportion", "income_min", "income_max", "comment", "modify_uid", "created_at", "updated_at"}
	cmfPlatformPolicyColumnsWithoutDefault = []string{"name", "key", "income_min", "income_max", "comment", "created_at", "updated_at"}
	cmfPlatformPolicyColumnsWithDefault    = []string{"id", "type", "state", "proportion", "modify_uid"}
	cmfPlatformPolicyPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFPlatformPolicySlice is an alias for a slice of pointers to CMFPlatformPolicy.
	// This should generally be used opposed to []CMFPlatformPolicy.
	CMFPlatformPolicySlice []*CMFPlatformPolicy
	// CMFPlatformPolicyHook is the signature for custom CMFPlatformPolicy hook methods
	CMFPlatformPolicyHook func(context.Context, boil.ContextExecutor, *CMFPlatformPolicy) error

	cmfPlatformPolicyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfPlatformPolicyType                 = reflect.TypeOf(&CMFPlatformPolicy{})
	cmfPlatformPolicyMapping              = queries.MakeStructMapping(cmfPlatformPolicyType)
	cmfPlatformPolicyPrimaryKeyMapping, _ = queries.BindMapping(cmfPlatformPolicyType, cmfPlatformPolicyMapping, cmfPlatformPolicyPrimaryKeyColumns)
	cmfPlatformPolicyInsertCacheMut       sync.RWMutex
	cmfPlatformPolicyInsertCache          = make(map[string]insertCache)
	cmfPlatformPolicyUpdateCacheMut       sync.RWMutex
	cmfPlatformPolicyUpdateCache          = make(map[string]updateCache)
	cmfPlatformPolicyUpsertCacheMut       sync.RWMutex
	cmfPlatformPolicyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfPlatformPolicyBeforeInsertHooks []CMFPlatformPolicyHook
var cmfPlatformPolicyBeforeUpdateHooks []CMFPlatformPolicyHook
var cmfPlatformPolicyBeforeDeleteHooks []CMFPlatformPolicyHook
var cmfPlatformPolicyBeforeUpsertHooks []CMFPlatformPolicyHook

var cmfPlatformPolicyAfterInsertHooks []CMFPlatformPolicyHook
var cmfPlatformPolicyAfterSelectHooks []CMFPlatformPolicyHook
var cmfPlatformPolicyAfterUpdateHooks []CMFPlatformPolicyHook
var cmfPlatformPolicyAfterDeleteHooks []CMFPlatformPolicyHook
var cmfPlatformPolicyAfterUpsertHooks []CMFPlatformPolicyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFPlatformPolicy) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFPlatformPolicy) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFPlatformPolicy) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFPlatformPolicy) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFPlatformPolicy) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFPlatformPolicy) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFPlatformPolicy) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFPlatformPolicy) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFPlatformPolicy) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPlatformPolicyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFPlatformPolicyHook registers your hook function for all future operations.
func AddCMFPlatformPolicyHook(hookPoint boil.HookPoint, cmfPlatformPolicyHook CMFPlatformPolicyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfPlatformPolicyBeforeInsertHooks = append(cmfPlatformPolicyBeforeInsertHooks, cmfPlatformPolicyHook)
	case boil.BeforeUpdateHook:
		cmfPlatformPolicyBeforeUpdateHooks = append(cmfPlatformPolicyBeforeUpdateHooks, cmfPlatformPolicyHook)
	case boil.BeforeDeleteHook:
		cmfPlatformPolicyBeforeDeleteHooks = append(cmfPlatformPolicyBeforeDeleteHooks, cmfPlatformPolicyHook)
	case boil.BeforeUpsertHook:
		cmfPlatformPolicyBeforeUpsertHooks = append(cmfPlatformPolicyBeforeUpsertHooks, cmfPlatformPolicyHook)
	case boil.AfterInsertHook:
		cmfPlatformPolicyAfterInsertHooks = append(cmfPlatformPolicyAfterInsertHooks, cmfPlatformPolicyHook)
	case boil.AfterSelectHook:
		cmfPlatformPolicyAfterSelectHooks = append(cmfPlatformPolicyAfterSelectHooks, cmfPlatformPolicyHook)
	case boil.AfterUpdateHook:
		cmfPlatformPolicyAfterUpdateHooks = append(cmfPlatformPolicyAfterUpdateHooks, cmfPlatformPolicyHook)
	case boil.AfterDeleteHook:
		cmfPlatformPolicyAfterDeleteHooks = append(cmfPlatformPolicyAfterDeleteHooks, cmfPlatformPolicyHook)
	case boil.AfterUpsertHook:
		cmfPlatformPolicyAfterUpsertHooks = append(cmfPlatformPolicyAfterUpsertHooks, cmfPlatformPolicyHook)
	}
}

// One returns a single cmfPlatformPolicy record from the query.
func (q cmfPlatformPolicyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFPlatformPolicy, error) {
	o := &CMFPlatformPolicy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_platform_policy")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFPlatformPolicy records from the query.
func (q cmfPlatformPolicyQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFPlatformPolicySlice, error) {
	var o []*CMFPlatformPolicy

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFPlatformPolicy slice")
	}

	if len(cmfPlatformPolicyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFPlatformPolicy records in the query.
func (q cmfPlatformPolicyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_platform_policy rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfPlatformPolicyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_platform_policy exists")
	}

	return count > 0, nil
}

// CMFPlatformPolicies retrieves all the records using an executor.
func CMFPlatformPolicies(mods ...qm.QueryMod) cmfPlatformPolicyQuery {
	mods = append(mods, qm.From("`cmf_platform_policy`"))
	return cmfPlatformPolicyQuery{NewQuery(mods...)}
}

// FindCMFPlatformPolicy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFPlatformPolicy(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFPlatformPolicy, error) {
	cmfPlatformPolicyObj := &CMFPlatformPolicy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_platform_policy` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfPlatformPolicyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_platform_policy")
	}

	return cmfPlatformPolicyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFPlatformPolicy) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_platform_policy provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(cmfPlatformPolicyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfPlatformPolicyInsertCacheMut.RLock()
	cache, cached := cmfPlatformPolicyInsertCache[key]
	cmfPlatformPolicyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfPlatformPolicyAllColumns,
			cmfPlatformPolicyColumnsWithDefault,
			cmfPlatformPolicyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfPlatformPolicyType, cmfPlatformPolicyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfPlatformPolicyType, cmfPlatformPolicyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_platform_policy` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_platform_policy` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_platform_policy` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfPlatformPolicyPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_platform_policy")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPlatformPolicyMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_platform_policy")
	}

CacheNoHooks:
	if !cached {
		cmfPlatformPolicyInsertCacheMut.Lock()
		cmfPlatformPolicyInsertCache[key] = cache
		cmfPlatformPolicyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFPlatformPolicy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFPlatformPolicy) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfPlatformPolicyUpdateCacheMut.RLock()
	cache, cached := cmfPlatformPolicyUpdateCache[key]
	cmfPlatformPolicyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfPlatformPolicyAllColumns,
			cmfPlatformPolicyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_platform_policy, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_platform_policy` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfPlatformPolicyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfPlatformPolicyType, cmfPlatformPolicyMapping, append(wl, cmfPlatformPolicyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_platform_policy row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_platform_policy")
	}

	if !cached {
		cmfPlatformPolicyUpdateCacheMut.Lock()
		cmfPlatformPolicyUpdateCache[key] = cache
		cmfPlatformPolicyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfPlatformPolicyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_platform_policy")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_platform_policy")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFPlatformPolicySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPlatformPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_platform_policy` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPlatformPolicyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfPlatformPolicy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfPlatformPolicy")
	}
	return rowsAff, nil
}

var mySQLCMFPlatformPolicyUniqueColumns = []string{
	"id",
	"key",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFPlatformPolicy) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_platform_policy provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(cmfPlatformPolicyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFPlatformPolicyUniqueColumns, o)

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

	cmfPlatformPolicyUpsertCacheMut.RLock()
	cache, cached := cmfPlatformPolicyUpsertCache[key]
	cmfPlatformPolicyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfPlatformPolicyAllColumns,
			cmfPlatformPolicyColumnsWithDefault,
			cmfPlatformPolicyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfPlatformPolicyAllColumns,
			cmfPlatformPolicyPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_platform_policy, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_platform_policy`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_platform_policy` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfPlatformPolicyType, cmfPlatformPolicyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfPlatformPolicyType, cmfPlatformPolicyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_platform_policy")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPlatformPolicyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfPlatformPolicyType, cmfPlatformPolicyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_platform_policy")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_platform_policy")
	}

CacheNoHooks:
	if !cached {
		cmfPlatformPolicyUpsertCacheMut.Lock()
		cmfPlatformPolicyUpsertCache[key] = cache
		cmfPlatformPolicyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFPlatformPolicy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFPlatformPolicy) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFPlatformPolicy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfPlatformPolicyPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_platform_policy` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_platform_policy")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_platform_policy")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfPlatformPolicyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfPlatformPolicyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_platform_policy")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_platform_policy")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFPlatformPolicySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfPlatformPolicyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPlatformPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_platform_policy` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPlatformPolicyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfPlatformPolicy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_platform_policy")
	}

	if len(cmfPlatformPolicyAfterDeleteHooks) != 0 {
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
func (o *CMFPlatformPolicy) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFPlatformPolicy(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFPlatformPolicySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFPlatformPolicySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPlatformPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_platform_policy`.* FROM `cmf_platform_policy` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPlatformPolicyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFPlatformPolicySlice")
	}

	*o = slice

	return nil
}

// CMFPlatformPolicyExists checks if the CMFPlatformPolicy row exists.
func CMFPlatformPolicyExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_platform_policy` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_platform_policy exists")
	}

	return exists, nil
}
