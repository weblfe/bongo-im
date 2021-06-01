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
	"github.com/volatiletech/strmangle"
)

// CMFUserExperienceLog is an object representing the database table.
type CMFUserExperienceLog struct {
	ID         int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID     int64       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ExpNum     int         `boil:"exp_num" json:"exp_num" toml:"exp_num" yaml:"exp_num"`
	Type       string      `boil:"type" json:"type" toml:"type" yaml:"type"`
	Option     int16       `boil:"option" json:"option" toml:"option" yaml:"option"`
	Source     string      `boil:"source" json:"source" toml:"source" yaml:"source"`
	Extras     null.String `boil:"extras" json:"extras,omitempty" toml:"extras" yaml:"extras,omitempty"`
	CreateTime time.Time   `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`

	R *cmfUserExperienceLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfUserExperienceLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFUserExperienceLogColumns = struct {
	ID         string
	UserID     string
	ExpNum     string
	Type       string
	Option     string
	Source     string
	Extras     string
	CreateTime string
}{
	ID:         "id",
	UserID:     "user_id",
	ExpNum:     "exp_num",
	Type:       "type",
	Option:     "option",
	Source:     "source",
	Extras:     "extras",
	CreateTime: "create_time",
}

// Generated where

var CMFUserExperienceLogWhere = struct {
	ID         whereHelperint64
	UserID     whereHelperint64
	ExpNum     whereHelperint
	Type       whereHelperstring
	Option     whereHelperint16
	Source     whereHelperstring
	Extras     whereHelpernull_String
	CreateTime whereHelpertime_Time
}{
	ID:         whereHelperint64{field: "`cmf_user_experience_log`.`id`"},
	UserID:     whereHelperint64{field: "`cmf_user_experience_log`.`user_id`"},
	ExpNum:     whereHelperint{field: "`cmf_user_experience_log`.`exp_num`"},
	Type:       whereHelperstring{field: "`cmf_user_experience_log`.`type`"},
	Option:     whereHelperint16{field: "`cmf_user_experience_log`.`option`"},
	Source:     whereHelperstring{field: "`cmf_user_experience_log`.`source`"},
	Extras:     whereHelpernull_String{field: "`cmf_user_experience_log`.`extras`"},
	CreateTime: whereHelpertime_Time{field: "`cmf_user_experience_log`.`create_time`"},
}

// CMFUserExperienceLogRels is where relationship names are stored.
var CMFUserExperienceLogRels = struct {
}{}

// cmfUserExperienceLogR is where relationships are stored.
type cmfUserExperienceLogR struct {
}

// NewStruct creates a new relationship struct
func (*cmfUserExperienceLogR) NewStruct() *cmfUserExperienceLogR {
	return &cmfUserExperienceLogR{}
}

// cmfUserExperienceLogL is where Load methods for each relationship are stored.
type cmfUserExperienceLogL struct{}

var (
	cmfUserExperienceLogAllColumns            = []string{"id", "user_id", "exp_num", "type", "option", "source", "extras", "create_time"}
	cmfUserExperienceLogColumnsWithoutDefault = []string{"user_id", "exp_num", "type", "option", "source", "extras", "create_time"}
	cmfUserExperienceLogColumnsWithDefault    = []string{"id"}
	cmfUserExperienceLogPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFUserExperienceLogSlice is an alias for a slice of pointers to CMFUserExperienceLog.
	// This should generally be used opposed to []CMFUserExperienceLog.
	CMFUserExperienceLogSlice []*CMFUserExperienceLog
	// CMFUserExperienceLogHook is the signature for custom CMFUserExperienceLog hook methods
	CMFUserExperienceLogHook func(context.Context, boil.ContextExecutor, *CMFUserExperienceLog) error

	cmfUserExperienceLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfUserExperienceLogType                 = reflect.TypeOf(&CMFUserExperienceLog{})
	cmfUserExperienceLogMapping              = queries.MakeStructMapping(cmfUserExperienceLogType)
	cmfUserExperienceLogPrimaryKeyMapping, _ = queries.BindMapping(cmfUserExperienceLogType, cmfUserExperienceLogMapping, cmfUserExperienceLogPrimaryKeyColumns)
	cmfUserExperienceLogInsertCacheMut       sync.RWMutex
	cmfUserExperienceLogInsertCache          = make(map[string]insertCache)
	cmfUserExperienceLogUpdateCacheMut       sync.RWMutex
	cmfUserExperienceLogUpdateCache          = make(map[string]updateCache)
	cmfUserExperienceLogUpsertCacheMut       sync.RWMutex
	cmfUserExperienceLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfUserExperienceLogBeforeInsertHooks []CMFUserExperienceLogHook
var cmfUserExperienceLogBeforeUpdateHooks []CMFUserExperienceLogHook
var cmfUserExperienceLogBeforeDeleteHooks []CMFUserExperienceLogHook
var cmfUserExperienceLogBeforeUpsertHooks []CMFUserExperienceLogHook

var cmfUserExperienceLogAfterInsertHooks []CMFUserExperienceLogHook
var cmfUserExperienceLogAfterSelectHooks []CMFUserExperienceLogHook
var cmfUserExperienceLogAfterUpdateHooks []CMFUserExperienceLogHook
var cmfUserExperienceLogAfterDeleteHooks []CMFUserExperienceLogHook
var cmfUserExperienceLogAfterUpsertHooks []CMFUserExperienceLogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFUserExperienceLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFUserExperienceLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFUserExperienceLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFUserExperienceLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFUserExperienceLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFUserExperienceLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFUserExperienceLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFUserExperienceLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFUserExperienceLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserExperienceLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFUserExperienceLogHook registers your hook function for all future operations.
func AddCMFUserExperienceLogHook(hookPoint boil.HookPoint, cmfUserExperienceLogHook CMFUserExperienceLogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfUserExperienceLogBeforeInsertHooks = append(cmfUserExperienceLogBeforeInsertHooks, cmfUserExperienceLogHook)
	case boil.BeforeUpdateHook:
		cmfUserExperienceLogBeforeUpdateHooks = append(cmfUserExperienceLogBeforeUpdateHooks, cmfUserExperienceLogHook)
	case boil.BeforeDeleteHook:
		cmfUserExperienceLogBeforeDeleteHooks = append(cmfUserExperienceLogBeforeDeleteHooks, cmfUserExperienceLogHook)
	case boil.BeforeUpsertHook:
		cmfUserExperienceLogBeforeUpsertHooks = append(cmfUserExperienceLogBeforeUpsertHooks, cmfUserExperienceLogHook)
	case boil.AfterInsertHook:
		cmfUserExperienceLogAfterInsertHooks = append(cmfUserExperienceLogAfterInsertHooks, cmfUserExperienceLogHook)
	case boil.AfterSelectHook:
		cmfUserExperienceLogAfterSelectHooks = append(cmfUserExperienceLogAfterSelectHooks, cmfUserExperienceLogHook)
	case boil.AfterUpdateHook:
		cmfUserExperienceLogAfterUpdateHooks = append(cmfUserExperienceLogAfterUpdateHooks, cmfUserExperienceLogHook)
	case boil.AfterDeleteHook:
		cmfUserExperienceLogAfterDeleteHooks = append(cmfUserExperienceLogAfterDeleteHooks, cmfUserExperienceLogHook)
	case boil.AfterUpsertHook:
		cmfUserExperienceLogAfterUpsertHooks = append(cmfUserExperienceLogAfterUpsertHooks, cmfUserExperienceLogHook)
	}
}

// One returns a single cmfUserExperienceLog record from the query.
func (q cmfUserExperienceLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFUserExperienceLog, error) {
	o := &CMFUserExperienceLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_user_experience_log")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFUserExperienceLog records from the query.
func (q cmfUserExperienceLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFUserExperienceLogSlice, error) {
	var o []*CMFUserExperienceLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFUserExperienceLog slice")
	}

	if len(cmfUserExperienceLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFUserExperienceLog records in the query.
func (q cmfUserExperienceLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_user_experience_log rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfUserExperienceLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_user_experience_log exists")
	}

	return count > 0, nil
}

// CMFUserExperienceLogs retrieves all the records using an executor.
func CMFUserExperienceLogs(mods ...qm.QueryMod) cmfUserExperienceLogQuery {
	mods = append(mods, qm.From("`cmf_user_experience_log`"))
	return cmfUserExperienceLogQuery{NewQuery(mods...)}
}

// FindCMFUserExperienceLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFUserExperienceLog(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CMFUserExperienceLog, error) {
	cmfUserExperienceLogObj := &CMFUserExperienceLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_user_experience_log` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfUserExperienceLogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_user_experience_log")
	}

	return cmfUserExperienceLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFUserExperienceLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_experience_log provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserExperienceLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfUserExperienceLogInsertCacheMut.RLock()
	cache, cached := cmfUserExperienceLogInsertCache[key]
	cmfUserExperienceLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfUserExperienceLogAllColumns,
			cmfUserExperienceLogColumnsWithDefault,
			cmfUserExperienceLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserExperienceLogType, cmfUserExperienceLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfUserExperienceLogType, cmfUserExperienceLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_user_experience_log` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_user_experience_log` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_user_experience_log` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfUserExperienceLogPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_user_experience_log")
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

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfUserExperienceLogMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_experience_log")
	}

CacheNoHooks:
	if !cached {
		cmfUserExperienceLogInsertCacheMut.Lock()
		cmfUserExperienceLogInsertCache[key] = cache
		cmfUserExperienceLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFUserExperienceLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFUserExperienceLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfUserExperienceLogUpdateCacheMut.RLock()
	cache, cached := cmfUserExperienceLogUpdateCache[key]
	cmfUserExperienceLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfUserExperienceLogAllColumns,
			cmfUserExperienceLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_user_experience_log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_user_experience_log` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfUserExperienceLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfUserExperienceLogType, cmfUserExperienceLogMapping, append(wl, cmfUserExperienceLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_user_experience_log row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_user_experience_log")
	}

	if !cached {
		cmfUserExperienceLogUpdateCacheMut.Lock()
		cmfUserExperienceLogUpdateCache[key] = cache
		cmfUserExperienceLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfUserExperienceLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_user_experience_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_user_experience_log")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFUserExperienceLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserExperienceLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_user_experience_log` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserExperienceLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfUserExperienceLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfUserExperienceLog")
	}
	return rowsAff, nil
}

var mySQLCMFUserExperienceLogUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFUserExperienceLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_experience_log provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserExperienceLogColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFUserExperienceLogUniqueColumns, o)

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

	cmfUserExperienceLogUpsertCacheMut.RLock()
	cache, cached := cmfUserExperienceLogUpsertCache[key]
	cmfUserExperienceLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfUserExperienceLogAllColumns,
			cmfUserExperienceLogColumnsWithDefault,
			cmfUserExperienceLogColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfUserExperienceLogAllColumns,
			cmfUserExperienceLogPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_user_experience_log, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_user_experience_log`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_user_experience_log` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserExperienceLogType, cmfUserExperienceLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfUserExperienceLogType, cmfUserExperienceLogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_user_experience_log")
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

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfUserExperienceLogMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfUserExperienceLogType, cmfUserExperienceLogMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_user_experience_log")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_experience_log")
	}

CacheNoHooks:
	if !cached {
		cmfUserExperienceLogUpsertCacheMut.Lock()
		cmfUserExperienceLogUpsertCache[key] = cache
		cmfUserExperienceLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFUserExperienceLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFUserExperienceLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFUserExperienceLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfUserExperienceLogPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_user_experience_log` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_user_experience_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_user_experience_log")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfUserExperienceLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfUserExperienceLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_user_experience_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_experience_log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFUserExperienceLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfUserExperienceLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserExperienceLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_user_experience_log` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserExperienceLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfUserExperienceLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_experience_log")
	}

	if len(cmfUserExperienceLogAfterDeleteHooks) != 0 {
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
func (o *CMFUserExperienceLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFUserExperienceLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFUserExperienceLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFUserExperienceLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserExperienceLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_user_experience_log`.* FROM `cmf_user_experience_log` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserExperienceLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFUserExperienceLogSlice")
	}

	*o = slice

	return nil
}

// CMFUserExperienceLogExists checks if the CMFUserExperienceLog row exists.
func CMFUserExperienceLogExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_user_experience_log` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_user_experience_log exists")
	}

	return exists, nil
}
