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

// CMFFamiliesPolicy is an object representing the database table.
type CMFFamiliesPolicy struct {
	ID         int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	Type       int16             `boil:"type" json:"type" toml:"type" yaml:"type"`
	FamilyID   int64             `boil:"family_id" json:"family_id" toml:"family_id" yaml:"family_id"`
	State      bool              `boil:"state" json:"state" toml:"state" yaml:"state"`
	Proportion types.Decimal     `boil:"proportion" json:"proportion" toml:"proportion" yaml:"proportion"`
	IncomeMin  types.NullDecimal `boil:"income_min" json:"income_min,omitempty" toml:"income_min" yaml:"income_min,omitempty"`
	IncomeMax  types.NullDecimal `boil:"income_max" json:"income_max,omitempty" toml:"income_max" yaml:"income_max,omitempty"`
	Unit       int16             `boil:"unit" json:"unit" toml:"unit" yaml:"unit"`
	Comment    null.String       `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`
	ModifyUID  int               `boil:"modify_uid" json:"modify_uid" toml:"modify_uid" yaml:"modify_uid"`
	CreatorUID int               `boil:"creator_uid" json:"creator_uid" toml:"creator_uid" yaml:"creator_uid"`
	CreatedAt  time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time         `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *cmfFamiliesPolicyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfFamiliesPolicyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFFamiliesPolicyColumns = struct {
	ID         string
	Type       string
	FamilyID   string
	State      string
	Proportion string
	IncomeMin  string
	IncomeMax  string
	Unit       string
	Comment    string
	ModifyUID  string
	CreatorUID string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	Type:       "type",
	FamilyID:   "family_id",
	State:      "state",
	Proportion: "proportion",
	IncomeMin:  "income_min",
	IncomeMax:  "income_max",
	Unit:       "unit",
	Comment:    "comment",
	ModifyUID:  "modify_uid",
	CreatorUID: "creator_uid",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// Generated where

var CMFFamiliesPolicyWhere = struct {
	ID         whereHelperint
	Type       whereHelperint16
	FamilyID   whereHelperint64
	State      whereHelperbool
	Proportion whereHelpertypes_Decimal
	IncomeMin  whereHelpertypes_NullDecimal
	IncomeMax  whereHelpertypes_NullDecimal
	Unit       whereHelperint16
	Comment    whereHelpernull_String
	ModifyUID  whereHelperint
	CreatorUID whereHelperint
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperint{field: "`cmf_families_policies`.`id`"},
	Type:       whereHelperint16{field: "`cmf_families_policies`.`type`"},
	FamilyID:   whereHelperint64{field: "`cmf_families_policies`.`family_id`"},
	State:      whereHelperbool{field: "`cmf_families_policies`.`state`"},
	Proportion: whereHelpertypes_Decimal{field: "`cmf_families_policies`.`proportion`"},
	IncomeMin:  whereHelpertypes_NullDecimal{field: "`cmf_families_policies`.`income_min`"},
	IncomeMax:  whereHelpertypes_NullDecimal{field: "`cmf_families_policies`.`income_max`"},
	Unit:       whereHelperint16{field: "`cmf_families_policies`.`unit`"},
	Comment:    whereHelpernull_String{field: "`cmf_families_policies`.`comment`"},
	ModifyUID:  whereHelperint{field: "`cmf_families_policies`.`modify_uid`"},
	CreatorUID: whereHelperint{field: "`cmf_families_policies`.`creator_uid`"},
	CreatedAt:  whereHelpertime_Time{field: "`cmf_families_policies`.`created_at`"},
	UpdatedAt:  whereHelpertime_Time{field: "`cmf_families_policies`.`updated_at`"},
}

// CMFFamiliesPolicyRels is where relationship names are stored.
var CMFFamiliesPolicyRels = struct {
}{}

// cmfFamiliesPolicyR is where relationships are stored.
type cmfFamiliesPolicyR struct {
}

// NewStruct creates a new relationship struct
func (*cmfFamiliesPolicyR) NewStruct() *cmfFamiliesPolicyR {
	return &cmfFamiliesPolicyR{}
}

// cmfFamiliesPolicyL is where Load methods for each relationship are stored.
type cmfFamiliesPolicyL struct{}

var (
	cmfFamiliesPolicyAllColumns            = []string{"id", "type", "family_id", "state", "proportion", "income_min", "income_max", "unit", "comment", "modify_uid", "creator_uid", "created_at", "updated_at"}
	cmfFamiliesPolicyColumnsWithoutDefault = []string{"family_id", "income_min", "income_max", "comment", "created_at", "updated_at"}
	cmfFamiliesPolicyColumnsWithDefault    = []string{"id", "type", "state", "proportion", "unit", "modify_uid", "creator_uid"}
	cmfFamiliesPolicyPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFFamiliesPolicySlice is an alias for a slice of pointers to CMFFamiliesPolicy.
	// This should generally be used opposed to []CMFFamiliesPolicy.
	CMFFamiliesPolicySlice []*CMFFamiliesPolicy
	// CMFFamiliesPolicyHook is the signature for custom CMFFamiliesPolicy hook methods
	CMFFamiliesPolicyHook func(context.Context, boil.ContextExecutor, *CMFFamiliesPolicy) error

	cmfFamiliesPolicyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfFamiliesPolicyType                 = reflect.TypeOf(&CMFFamiliesPolicy{})
	cmfFamiliesPolicyMapping              = queries.MakeStructMapping(cmfFamiliesPolicyType)
	cmfFamiliesPolicyPrimaryKeyMapping, _ = queries.BindMapping(cmfFamiliesPolicyType, cmfFamiliesPolicyMapping, cmfFamiliesPolicyPrimaryKeyColumns)
	cmfFamiliesPolicyInsertCacheMut       sync.RWMutex
	cmfFamiliesPolicyInsertCache          = make(map[string]insertCache)
	cmfFamiliesPolicyUpdateCacheMut       sync.RWMutex
	cmfFamiliesPolicyUpdateCache          = make(map[string]updateCache)
	cmfFamiliesPolicyUpsertCacheMut       sync.RWMutex
	cmfFamiliesPolicyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfFamiliesPolicyBeforeInsertHooks []CMFFamiliesPolicyHook
var cmfFamiliesPolicyBeforeUpdateHooks []CMFFamiliesPolicyHook
var cmfFamiliesPolicyBeforeDeleteHooks []CMFFamiliesPolicyHook
var cmfFamiliesPolicyBeforeUpsertHooks []CMFFamiliesPolicyHook

var cmfFamiliesPolicyAfterInsertHooks []CMFFamiliesPolicyHook
var cmfFamiliesPolicyAfterSelectHooks []CMFFamiliesPolicyHook
var cmfFamiliesPolicyAfterUpdateHooks []CMFFamiliesPolicyHook
var cmfFamiliesPolicyAfterDeleteHooks []CMFFamiliesPolicyHook
var cmfFamiliesPolicyAfterUpsertHooks []CMFFamiliesPolicyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFFamiliesPolicy) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFFamiliesPolicy) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFFamiliesPolicy) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFFamiliesPolicy) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFFamiliesPolicy) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFFamiliesPolicy) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFFamiliesPolicy) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFFamiliesPolicy) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFFamiliesPolicy) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamiliesPolicyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFFamiliesPolicyHook registers your hook function for all future operations.
func AddCMFFamiliesPolicyHook(hookPoint boil.HookPoint, cmfFamiliesPolicyHook CMFFamiliesPolicyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfFamiliesPolicyBeforeInsertHooks = append(cmfFamiliesPolicyBeforeInsertHooks, cmfFamiliesPolicyHook)
	case boil.BeforeUpdateHook:
		cmfFamiliesPolicyBeforeUpdateHooks = append(cmfFamiliesPolicyBeforeUpdateHooks, cmfFamiliesPolicyHook)
	case boil.BeforeDeleteHook:
		cmfFamiliesPolicyBeforeDeleteHooks = append(cmfFamiliesPolicyBeforeDeleteHooks, cmfFamiliesPolicyHook)
	case boil.BeforeUpsertHook:
		cmfFamiliesPolicyBeforeUpsertHooks = append(cmfFamiliesPolicyBeforeUpsertHooks, cmfFamiliesPolicyHook)
	case boil.AfterInsertHook:
		cmfFamiliesPolicyAfterInsertHooks = append(cmfFamiliesPolicyAfterInsertHooks, cmfFamiliesPolicyHook)
	case boil.AfterSelectHook:
		cmfFamiliesPolicyAfterSelectHooks = append(cmfFamiliesPolicyAfterSelectHooks, cmfFamiliesPolicyHook)
	case boil.AfterUpdateHook:
		cmfFamiliesPolicyAfterUpdateHooks = append(cmfFamiliesPolicyAfterUpdateHooks, cmfFamiliesPolicyHook)
	case boil.AfterDeleteHook:
		cmfFamiliesPolicyAfterDeleteHooks = append(cmfFamiliesPolicyAfterDeleteHooks, cmfFamiliesPolicyHook)
	case boil.AfterUpsertHook:
		cmfFamiliesPolicyAfterUpsertHooks = append(cmfFamiliesPolicyAfterUpsertHooks, cmfFamiliesPolicyHook)
	}
}

// One returns a single cmfFamiliesPolicy record from the query.
func (q cmfFamiliesPolicyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFFamiliesPolicy, error) {
	o := &CMFFamiliesPolicy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_families_policies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFFamiliesPolicy records from the query.
func (q cmfFamiliesPolicyQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFFamiliesPolicySlice, error) {
	var o []*CMFFamiliesPolicy

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFFamiliesPolicy slice")
	}

	if len(cmfFamiliesPolicyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFFamiliesPolicy records in the query.
func (q cmfFamiliesPolicyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_families_policies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfFamiliesPolicyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_families_policies exists")
	}

	return count > 0, nil
}

// CMFFamiliesPolicies retrieves all the records using an executor.
func CMFFamiliesPolicies(mods ...qm.QueryMod) cmfFamiliesPolicyQuery {
	mods = append(mods, qm.From("`cmf_families_policies`"))
	return cmfFamiliesPolicyQuery{NewQuery(mods...)}
}

// FindCMFFamiliesPolicy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFFamiliesPolicy(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFFamiliesPolicy, error) {
	cmfFamiliesPolicyObj := &CMFFamiliesPolicy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_families_policies` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfFamiliesPolicyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_families_policies")
	}

	return cmfFamiliesPolicyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFFamiliesPolicy) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_families_policies provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(cmfFamiliesPolicyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfFamiliesPolicyInsertCacheMut.RLock()
	cache, cached := cmfFamiliesPolicyInsertCache[key]
	cmfFamiliesPolicyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfFamiliesPolicyAllColumns,
			cmfFamiliesPolicyColumnsWithDefault,
			cmfFamiliesPolicyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfFamiliesPolicyType, cmfFamiliesPolicyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfFamiliesPolicyType, cmfFamiliesPolicyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_families_policies` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_families_policies` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_families_policies` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfFamiliesPolicyPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_families_policies")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfFamiliesPolicyMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_families_policies")
	}

CacheNoHooks:
	if !cached {
		cmfFamiliesPolicyInsertCacheMut.Lock()
		cmfFamiliesPolicyInsertCache[key] = cache
		cmfFamiliesPolicyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFFamiliesPolicy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFFamiliesPolicy) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfFamiliesPolicyUpdateCacheMut.RLock()
	cache, cached := cmfFamiliesPolicyUpdateCache[key]
	cmfFamiliesPolicyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfFamiliesPolicyAllColumns,
			cmfFamiliesPolicyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_families_policies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_families_policies` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfFamiliesPolicyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfFamiliesPolicyType, cmfFamiliesPolicyMapping, append(wl, cmfFamiliesPolicyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_families_policies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_families_policies")
	}

	if !cached {
		cmfFamiliesPolicyUpdateCacheMut.Lock()
		cmfFamiliesPolicyUpdateCache[key] = cache
		cmfFamiliesPolicyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfFamiliesPolicyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_families_policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_families_policies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFFamiliesPolicySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamiliesPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_families_policies` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamiliesPolicyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfFamiliesPolicy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfFamiliesPolicy")
	}
	return rowsAff, nil
}

var mySQLCMFFamiliesPolicyUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFFamiliesPolicy) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_families_policies provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(cmfFamiliesPolicyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFFamiliesPolicyUniqueColumns, o)

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

	cmfFamiliesPolicyUpsertCacheMut.RLock()
	cache, cached := cmfFamiliesPolicyUpsertCache[key]
	cmfFamiliesPolicyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfFamiliesPolicyAllColumns,
			cmfFamiliesPolicyColumnsWithDefault,
			cmfFamiliesPolicyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfFamiliesPolicyAllColumns,
			cmfFamiliesPolicyPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_families_policies, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_families_policies`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_families_policies` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfFamiliesPolicyType, cmfFamiliesPolicyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfFamiliesPolicyType, cmfFamiliesPolicyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_families_policies")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfFamiliesPolicyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfFamiliesPolicyType, cmfFamiliesPolicyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_families_policies")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_families_policies")
	}

CacheNoHooks:
	if !cached {
		cmfFamiliesPolicyUpsertCacheMut.Lock()
		cmfFamiliesPolicyUpsertCache[key] = cache
		cmfFamiliesPolicyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFFamiliesPolicy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFFamiliesPolicy) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFFamiliesPolicy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfFamiliesPolicyPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_families_policies` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_families_policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_families_policies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfFamiliesPolicyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfFamiliesPolicyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_families_policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_families_policies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFFamiliesPolicySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfFamiliesPolicyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamiliesPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_families_policies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamiliesPolicyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfFamiliesPolicy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_families_policies")
	}

	if len(cmfFamiliesPolicyAfterDeleteHooks) != 0 {
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
func (o *CMFFamiliesPolicy) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFFamiliesPolicy(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFFamiliesPolicySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFFamiliesPolicySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamiliesPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_families_policies`.* FROM `cmf_families_policies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamiliesPolicyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFFamiliesPolicySlice")
	}

	*o = slice

	return nil
}

// CMFFamiliesPolicyExists checks if the CMFFamiliesPolicy row exists.
func CMFFamiliesPolicyExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_families_policies` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_families_policies exists")
	}

	return exists, nil
}
