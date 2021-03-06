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
	"github.com/volatiletech/strmangle"
)

// CMFLiveKick is an object representing the database table.
type CMFLiveKick struct {
	ID       uint64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID      uint   `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Liveuid  uint   `boil:"liveuid" json:"liveuid" toml:"liveuid" yaml:"liveuid"`
	Addtime  int    `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Actionid int    `boil:"actionid" json:"actionid" toml:"actionid" yaml:"actionid"`

	R *cmfLiveKickR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfLiveKickL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFLiveKickColumns = struct {
	ID       string
	UID      string
	Liveuid  string
	Addtime  string
	Actionid string
}{
	ID:       "id",
	UID:      "uid",
	Liveuid:  "liveuid",
	Addtime:  "addtime",
	Actionid: "actionid",
}

// Generated where

var CMFLiveKickWhere = struct {
	ID       whereHelperuint64
	UID      whereHelperuint
	Liveuid  whereHelperuint
	Addtime  whereHelperint
	Actionid whereHelperint
}{
	ID:       whereHelperuint64{field: "`cmf_live_kick`.`id`"},
	UID:      whereHelperuint{field: "`cmf_live_kick`.`uid`"},
	Liveuid:  whereHelperuint{field: "`cmf_live_kick`.`liveuid`"},
	Addtime:  whereHelperint{field: "`cmf_live_kick`.`addtime`"},
	Actionid: whereHelperint{field: "`cmf_live_kick`.`actionid`"},
}

// CMFLiveKickRels is where relationship names are stored.
var CMFLiveKickRels = struct {
}{}

// cmfLiveKickR is where relationships are stored.
type cmfLiveKickR struct {
}

// NewStruct creates a new relationship struct
func (*cmfLiveKickR) NewStruct() *cmfLiveKickR {
	return &cmfLiveKickR{}
}

// cmfLiveKickL is where Load methods for each relationship are stored.
type cmfLiveKickL struct{}

var (
	cmfLiveKickAllColumns            = []string{"id", "uid", "liveuid", "addtime", "actionid"}
	cmfLiveKickColumnsWithoutDefault = []string{}
	cmfLiveKickColumnsWithDefault    = []string{"id", "uid", "liveuid", "addtime", "actionid"}
	cmfLiveKickPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFLiveKickSlice is an alias for a slice of pointers to CMFLiveKick.
	// This should generally be used opposed to []CMFLiveKick.
	CMFLiveKickSlice []*CMFLiveKick
	// CMFLiveKickHook is the signature for custom CMFLiveKick hook methods
	CMFLiveKickHook func(context.Context, boil.ContextExecutor, *CMFLiveKick) error

	cmfLiveKickQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfLiveKickType                 = reflect.TypeOf(&CMFLiveKick{})
	cmfLiveKickMapping              = queries.MakeStructMapping(cmfLiveKickType)
	cmfLiveKickPrimaryKeyMapping, _ = queries.BindMapping(cmfLiveKickType, cmfLiveKickMapping, cmfLiveKickPrimaryKeyColumns)
	cmfLiveKickInsertCacheMut       sync.RWMutex
	cmfLiveKickInsertCache          = make(map[string]insertCache)
	cmfLiveKickUpdateCacheMut       sync.RWMutex
	cmfLiveKickUpdateCache          = make(map[string]updateCache)
	cmfLiveKickUpsertCacheMut       sync.RWMutex
	cmfLiveKickUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfLiveKickBeforeInsertHooks []CMFLiveKickHook
var cmfLiveKickBeforeUpdateHooks []CMFLiveKickHook
var cmfLiveKickBeforeDeleteHooks []CMFLiveKickHook
var cmfLiveKickBeforeUpsertHooks []CMFLiveKickHook

var cmfLiveKickAfterInsertHooks []CMFLiveKickHook
var cmfLiveKickAfterSelectHooks []CMFLiveKickHook
var cmfLiveKickAfterUpdateHooks []CMFLiveKickHook
var cmfLiveKickAfterDeleteHooks []CMFLiveKickHook
var cmfLiveKickAfterUpsertHooks []CMFLiveKickHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFLiveKick) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFLiveKick) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFLiveKick) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFLiveKick) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFLiveKick) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFLiveKick) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFLiveKick) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFLiveKick) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFLiveKick) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveKickAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFLiveKickHook registers your hook function for all future operations.
func AddCMFLiveKickHook(hookPoint boil.HookPoint, cmfLiveKickHook CMFLiveKickHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfLiveKickBeforeInsertHooks = append(cmfLiveKickBeforeInsertHooks, cmfLiveKickHook)
	case boil.BeforeUpdateHook:
		cmfLiveKickBeforeUpdateHooks = append(cmfLiveKickBeforeUpdateHooks, cmfLiveKickHook)
	case boil.BeforeDeleteHook:
		cmfLiveKickBeforeDeleteHooks = append(cmfLiveKickBeforeDeleteHooks, cmfLiveKickHook)
	case boil.BeforeUpsertHook:
		cmfLiveKickBeforeUpsertHooks = append(cmfLiveKickBeforeUpsertHooks, cmfLiveKickHook)
	case boil.AfterInsertHook:
		cmfLiveKickAfterInsertHooks = append(cmfLiveKickAfterInsertHooks, cmfLiveKickHook)
	case boil.AfterSelectHook:
		cmfLiveKickAfterSelectHooks = append(cmfLiveKickAfterSelectHooks, cmfLiveKickHook)
	case boil.AfterUpdateHook:
		cmfLiveKickAfterUpdateHooks = append(cmfLiveKickAfterUpdateHooks, cmfLiveKickHook)
	case boil.AfterDeleteHook:
		cmfLiveKickAfterDeleteHooks = append(cmfLiveKickAfterDeleteHooks, cmfLiveKickHook)
	case boil.AfterUpsertHook:
		cmfLiveKickAfterUpsertHooks = append(cmfLiveKickAfterUpsertHooks, cmfLiveKickHook)
	}
}

// One returns a single cmfLiveKick record from the query.
func (q cmfLiveKickQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFLiveKick, error) {
	o := &CMFLiveKick{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_live_kick")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFLiveKick records from the query.
func (q cmfLiveKickQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFLiveKickSlice, error) {
	var o []*CMFLiveKick

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFLiveKick slice")
	}

	if len(cmfLiveKickAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFLiveKick records in the query.
func (q cmfLiveKickQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_live_kick rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfLiveKickQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_live_kick exists")
	}

	return count > 0, nil
}

// CMFLiveKicks retrieves all the records using an executor.
func CMFLiveKicks(mods ...qm.QueryMod) cmfLiveKickQuery {
	mods = append(mods, qm.From("`cmf_live_kick`"))
	return cmfLiveKickQuery{NewQuery(mods...)}
}

// FindCMFLiveKick retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFLiveKick(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*CMFLiveKick, error) {
	cmfLiveKickObj := &CMFLiveKick{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_live_kick` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfLiveKickObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_live_kick")
	}

	return cmfLiveKickObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFLiveKick) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_live_kick provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfLiveKickColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfLiveKickInsertCacheMut.RLock()
	cache, cached := cmfLiveKickInsertCache[key]
	cmfLiveKickInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfLiveKickAllColumns,
			cmfLiveKickColumnsWithDefault,
			cmfLiveKickColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfLiveKickType, cmfLiveKickMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfLiveKickType, cmfLiveKickMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_live_kick` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_live_kick` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_live_kick` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfLiveKickPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_live_kick")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfLiveKickMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_live_kick")
	}

CacheNoHooks:
	if !cached {
		cmfLiveKickInsertCacheMut.Lock()
		cmfLiveKickInsertCache[key] = cache
		cmfLiveKickInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFLiveKick.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFLiveKick) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfLiveKickUpdateCacheMut.RLock()
	cache, cached := cmfLiveKickUpdateCache[key]
	cmfLiveKickUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfLiveKickAllColumns,
			cmfLiveKickPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_live_kick, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_live_kick` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfLiveKickPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfLiveKickType, cmfLiveKickMapping, append(wl, cmfLiveKickPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_live_kick row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_live_kick")
	}

	if !cached {
		cmfLiveKickUpdateCacheMut.Lock()
		cmfLiveKickUpdateCache[key] = cache
		cmfLiveKickUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfLiveKickQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_live_kick")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_live_kick")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFLiveKickSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLiveKickPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_live_kick` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLiveKickPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfLiveKick slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfLiveKick")
	}
	return rowsAff, nil
}

var mySQLCMFLiveKickUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFLiveKick) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_live_kick provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfLiveKickColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFLiveKickUniqueColumns, o)

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

	cmfLiveKickUpsertCacheMut.RLock()
	cache, cached := cmfLiveKickUpsertCache[key]
	cmfLiveKickUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfLiveKickAllColumns,
			cmfLiveKickColumnsWithDefault,
			cmfLiveKickColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfLiveKickAllColumns,
			cmfLiveKickPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_live_kick, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_live_kick`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_live_kick` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfLiveKickType, cmfLiveKickMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfLiveKickType, cmfLiveKickMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_live_kick")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfLiveKickMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfLiveKickType, cmfLiveKickMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_live_kick")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_live_kick")
	}

CacheNoHooks:
	if !cached {
		cmfLiveKickUpsertCacheMut.Lock()
		cmfLiveKickUpsertCache[key] = cache
		cmfLiveKickUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFLiveKick record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFLiveKick) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFLiveKick provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfLiveKickPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_live_kick` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_live_kick")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_live_kick")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfLiveKickQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfLiveKickQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_live_kick")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_live_kick")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFLiveKickSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfLiveKickBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLiveKickPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_live_kick` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLiveKickPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfLiveKick slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_live_kick")
	}

	if len(cmfLiveKickAfterDeleteHooks) != 0 {
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
func (o *CMFLiveKick) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFLiveKick(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFLiveKickSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFLiveKickSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLiveKickPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_live_kick`.* FROM `cmf_live_kick` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLiveKickPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFLiveKickSlice")
	}

	*o = slice

	return nil
}

// CMFLiveKickExists checks if the CMFLiveKick row exists.
func CMFLiveKickExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_live_kick` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_live_kick exists")
	}

	return exists, nil
}
