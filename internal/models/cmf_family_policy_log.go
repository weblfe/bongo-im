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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// CMFFamilyPolicyLog is an object representing the database table.
type CMFFamilyPolicyLog struct {
	ID             int           `boil:"id" json:"id" toml:"id" yaml:"id"`
	FamilyID       int           `boil:"family_id" json:"family_id" toml:"family_id" yaml:"family_id"`
	Policy         int16         `boil:"policy" json:"policy" toml:"policy" yaml:"policy"`
	Proportion     types.Decimal `boil:"proportion" json:"proportion" toml:"proportion" yaml:"proportion"`
	PrevPolicy     int16         `boil:"prev_policy" json:"prev_policy" toml:"prev_policy" yaml:"prev_policy"`
	PrevProportion types.Decimal `boil:"prev_proportion" json:"prev_proportion" toml:"prev_proportion" yaml:"prev_proportion"`
	CreatorUID     int           `boil:"creator_uid" json:"creator_uid" toml:"creator_uid" yaml:"creator_uid"`
	EffectedAt     time.Time     `boil:"effected_at" json:"effected_at" toml:"effected_at" yaml:"effected_at"`
	CreatedAt      time.Time     `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *cmfFamilyPolicyLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfFamilyPolicyLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFFamilyPolicyLogColumns = struct {
	ID             string
	FamilyID       string
	Policy         string
	Proportion     string
	PrevPolicy     string
	PrevProportion string
	CreatorUID     string
	EffectedAt     string
	CreatedAt      string
}{
	ID:             "id",
	FamilyID:       "family_id",
	Policy:         "policy",
	Proportion:     "proportion",
	PrevPolicy:     "prev_policy",
	PrevProportion: "prev_proportion",
	CreatorUID:     "creator_uid",
	EffectedAt:     "effected_at",
	CreatedAt:      "created_at",
}

// Generated where

var CMFFamilyPolicyLogWhere = struct {
	ID             whereHelperint
	FamilyID       whereHelperint
	Policy         whereHelperint16
	Proportion     whereHelpertypes_Decimal
	PrevPolicy     whereHelperint16
	PrevProportion whereHelpertypes_Decimal
	CreatorUID     whereHelperint
	EffectedAt     whereHelpertime_Time
	CreatedAt      whereHelpertime_Time
}{
	ID:             whereHelperint{field: "`cmf_family_policy_log`.`id`"},
	FamilyID:       whereHelperint{field: "`cmf_family_policy_log`.`family_id`"},
	Policy:         whereHelperint16{field: "`cmf_family_policy_log`.`policy`"},
	Proportion:     whereHelpertypes_Decimal{field: "`cmf_family_policy_log`.`proportion`"},
	PrevPolicy:     whereHelperint16{field: "`cmf_family_policy_log`.`prev_policy`"},
	PrevProportion: whereHelpertypes_Decimal{field: "`cmf_family_policy_log`.`prev_proportion`"},
	CreatorUID:     whereHelperint{field: "`cmf_family_policy_log`.`creator_uid`"},
	EffectedAt:     whereHelpertime_Time{field: "`cmf_family_policy_log`.`effected_at`"},
	CreatedAt:      whereHelpertime_Time{field: "`cmf_family_policy_log`.`created_at`"},
}

// CMFFamilyPolicyLogRels is where relationship names are stored.
var CMFFamilyPolicyLogRels = struct {
}{}

// cmfFamilyPolicyLogR is where relationships are stored.
type cmfFamilyPolicyLogR struct {
}

// NewStruct creates a new relationship struct
func (*cmfFamilyPolicyLogR) NewStruct() *cmfFamilyPolicyLogR {
	return &cmfFamilyPolicyLogR{}
}

// cmfFamilyPolicyLogL is where Load methods for each relationship are stored.
type cmfFamilyPolicyLogL struct{}

var (
	cmfFamilyPolicyLogAllColumns            = []string{"id", "family_id", "policy", "proportion", "prev_policy", "prev_proportion", "creator_uid", "effected_at", "created_at"}
	cmfFamilyPolicyLogColumnsWithoutDefault = []string{"family_id", "policy", "prev_policy", "creator_uid", "effected_at", "created_at"}
	cmfFamilyPolicyLogColumnsWithDefault    = []string{"id", "proportion", "prev_proportion"}
	cmfFamilyPolicyLogPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFFamilyPolicyLogSlice is an alias for a slice of pointers to CMFFamilyPolicyLog.
	// This should generally be used opposed to []CMFFamilyPolicyLog.
	CMFFamilyPolicyLogSlice []*CMFFamilyPolicyLog
	// CMFFamilyPolicyLogHook is the signature for custom CMFFamilyPolicyLog hook methods
	CMFFamilyPolicyLogHook func(context.Context, boil.ContextExecutor, *CMFFamilyPolicyLog) error

	cmfFamilyPolicyLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfFamilyPolicyLogType                 = reflect.TypeOf(&CMFFamilyPolicyLog{})
	cmfFamilyPolicyLogMapping              = queries.MakeStructMapping(cmfFamilyPolicyLogType)
	cmfFamilyPolicyLogPrimaryKeyMapping, _ = queries.BindMapping(cmfFamilyPolicyLogType, cmfFamilyPolicyLogMapping, cmfFamilyPolicyLogPrimaryKeyColumns)
	cmfFamilyPolicyLogInsertCacheMut       sync.RWMutex
	cmfFamilyPolicyLogInsertCache          = make(map[string]insertCache)
	cmfFamilyPolicyLogUpdateCacheMut       sync.RWMutex
	cmfFamilyPolicyLogUpdateCache          = make(map[string]updateCache)
	cmfFamilyPolicyLogUpsertCacheMut       sync.RWMutex
	cmfFamilyPolicyLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfFamilyPolicyLogBeforeInsertHooks []CMFFamilyPolicyLogHook
var cmfFamilyPolicyLogBeforeUpdateHooks []CMFFamilyPolicyLogHook
var cmfFamilyPolicyLogBeforeDeleteHooks []CMFFamilyPolicyLogHook
var cmfFamilyPolicyLogBeforeUpsertHooks []CMFFamilyPolicyLogHook

var cmfFamilyPolicyLogAfterInsertHooks []CMFFamilyPolicyLogHook
var cmfFamilyPolicyLogAfterSelectHooks []CMFFamilyPolicyLogHook
var cmfFamilyPolicyLogAfterUpdateHooks []CMFFamilyPolicyLogHook
var cmfFamilyPolicyLogAfterDeleteHooks []CMFFamilyPolicyLogHook
var cmfFamilyPolicyLogAfterUpsertHooks []CMFFamilyPolicyLogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFFamilyPolicyLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFFamilyPolicyLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFFamilyPolicyLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFFamilyPolicyLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFFamilyPolicyLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFFamilyPolicyLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFFamilyPolicyLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFFamilyPolicyLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFFamilyPolicyLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfFamilyPolicyLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFFamilyPolicyLogHook registers your hook function for all future operations.
func AddCMFFamilyPolicyLogHook(hookPoint boil.HookPoint, cmfFamilyPolicyLogHook CMFFamilyPolicyLogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfFamilyPolicyLogBeforeInsertHooks = append(cmfFamilyPolicyLogBeforeInsertHooks, cmfFamilyPolicyLogHook)
	case boil.BeforeUpdateHook:
		cmfFamilyPolicyLogBeforeUpdateHooks = append(cmfFamilyPolicyLogBeforeUpdateHooks, cmfFamilyPolicyLogHook)
	case boil.BeforeDeleteHook:
		cmfFamilyPolicyLogBeforeDeleteHooks = append(cmfFamilyPolicyLogBeforeDeleteHooks, cmfFamilyPolicyLogHook)
	case boil.BeforeUpsertHook:
		cmfFamilyPolicyLogBeforeUpsertHooks = append(cmfFamilyPolicyLogBeforeUpsertHooks, cmfFamilyPolicyLogHook)
	case boil.AfterInsertHook:
		cmfFamilyPolicyLogAfterInsertHooks = append(cmfFamilyPolicyLogAfterInsertHooks, cmfFamilyPolicyLogHook)
	case boil.AfterSelectHook:
		cmfFamilyPolicyLogAfterSelectHooks = append(cmfFamilyPolicyLogAfterSelectHooks, cmfFamilyPolicyLogHook)
	case boil.AfterUpdateHook:
		cmfFamilyPolicyLogAfterUpdateHooks = append(cmfFamilyPolicyLogAfterUpdateHooks, cmfFamilyPolicyLogHook)
	case boil.AfterDeleteHook:
		cmfFamilyPolicyLogAfterDeleteHooks = append(cmfFamilyPolicyLogAfterDeleteHooks, cmfFamilyPolicyLogHook)
	case boil.AfterUpsertHook:
		cmfFamilyPolicyLogAfterUpsertHooks = append(cmfFamilyPolicyLogAfterUpsertHooks, cmfFamilyPolicyLogHook)
	}
}

// One returns a single cmfFamilyPolicyLog record from the query.
func (q cmfFamilyPolicyLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFFamilyPolicyLog, error) {
	o := &CMFFamilyPolicyLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_family_policy_log")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFFamilyPolicyLog records from the query.
func (q cmfFamilyPolicyLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFFamilyPolicyLogSlice, error) {
	var o []*CMFFamilyPolicyLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFFamilyPolicyLog slice")
	}

	if len(cmfFamilyPolicyLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFFamilyPolicyLog records in the query.
func (q cmfFamilyPolicyLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_family_policy_log rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfFamilyPolicyLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_family_policy_log exists")
	}

	return count > 0, nil
}

// CMFFamilyPolicyLogs retrieves all the records using an executor.
func CMFFamilyPolicyLogs(mods ...qm.QueryMod) cmfFamilyPolicyLogQuery {
	mods = append(mods, qm.From("`cmf_family_policy_log`"))
	return cmfFamilyPolicyLogQuery{NewQuery(mods...)}
}

// FindCMFFamilyPolicyLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFFamilyPolicyLog(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFFamilyPolicyLog, error) {
	cmfFamilyPolicyLogObj := &CMFFamilyPolicyLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_family_policy_log` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfFamilyPolicyLogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_family_policy_log")
	}

	return cmfFamilyPolicyLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFFamilyPolicyLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_family_policy_log provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfFamilyPolicyLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfFamilyPolicyLogInsertCacheMut.RLock()
	cache, cached := cmfFamilyPolicyLogInsertCache[key]
	cmfFamilyPolicyLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfFamilyPolicyLogAllColumns,
			cmfFamilyPolicyLogColumnsWithDefault,
			cmfFamilyPolicyLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfFamilyPolicyLogType, cmfFamilyPolicyLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfFamilyPolicyLogType, cmfFamilyPolicyLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_family_policy_log` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_family_policy_log` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_family_policy_log` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfFamilyPolicyLogPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_family_policy_log")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfFamilyPolicyLogMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_family_policy_log")
	}

CacheNoHooks:
	if !cached {
		cmfFamilyPolicyLogInsertCacheMut.Lock()
		cmfFamilyPolicyLogInsertCache[key] = cache
		cmfFamilyPolicyLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFFamilyPolicyLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFFamilyPolicyLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfFamilyPolicyLogUpdateCacheMut.RLock()
	cache, cached := cmfFamilyPolicyLogUpdateCache[key]
	cmfFamilyPolicyLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfFamilyPolicyLogAllColumns,
			cmfFamilyPolicyLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_family_policy_log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_family_policy_log` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfFamilyPolicyLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfFamilyPolicyLogType, cmfFamilyPolicyLogMapping, append(wl, cmfFamilyPolicyLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_family_policy_log row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_family_policy_log")
	}

	if !cached {
		cmfFamilyPolicyLogUpdateCacheMut.Lock()
		cmfFamilyPolicyLogUpdateCache[key] = cache
		cmfFamilyPolicyLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfFamilyPolicyLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_family_policy_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_family_policy_log")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFFamilyPolicyLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamilyPolicyLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_family_policy_log` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamilyPolicyLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfFamilyPolicyLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfFamilyPolicyLog")
	}
	return rowsAff, nil
}

var mySQLCMFFamilyPolicyLogUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFFamilyPolicyLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_family_policy_log provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfFamilyPolicyLogColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFFamilyPolicyLogUniqueColumns, o)

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

	cmfFamilyPolicyLogUpsertCacheMut.RLock()
	cache, cached := cmfFamilyPolicyLogUpsertCache[key]
	cmfFamilyPolicyLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfFamilyPolicyLogAllColumns,
			cmfFamilyPolicyLogColumnsWithDefault,
			cmfFamilyPolicyLogColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfFamilyPolicyLogAllColumns,
			cmfFamilyPolicyLogPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_family_policy_log, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_family_policy_log`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_family_policy_log` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfFamilyPolicyLogType, cmfFamilyPolicyLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfFamilyPolicyLogType, cmfFamilyPolicyLogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_family_policy_log")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfFamilyPolicyLogMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfFamilyPolicyLogType, cmfFamilyPolicyLogMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_family_policy_log")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_family_policy_log")
	}

CacheNoHooks:
	if !cached {
		cmfFamilyPolicyLogUpsertCacheMut.Lock()
		cmfFamilyPolicyLogUpsertCache[key] = cache
		cmfFamilyPolicyLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFFamilyPolicyLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFFamilyPolicyLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFFamilyPolicyLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfFamilyPolicyLogPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_family_policy_log` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_family_policy_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_family_policy_log")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfFamilyPolicyLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfFamilyPolicyLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_family_policy_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_family_policy_log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFFamilyPolicyLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfFamilyPolicyLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamilyPolicyLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_family_policy_log` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamilyPolicyLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfFamilyPolicyLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_family_policy_log")
	}

	if len(cmfFamilyPolicyLogAfterDeleteHooks) != 0 {
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
func (o *CMFFamilyPolicyLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFFamilyPolicyLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFFamilyPolicyLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFFamilyPolicyLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfFamilyPolicyLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_family_policy_log`.* FROM `cmf_family_policy_log` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfFamilyPolicyLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFFamilyPolicyLogSlice")
	}

	*o = slice

	return nil
}

// CMFFamilyPolicyLogExists checks if the CMFFamilyPolicyLog row exists.
func CMFFamilyPolicyLogExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_family_policy_log` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_family_policy_log exists")
	}

	return exists, nil
}
