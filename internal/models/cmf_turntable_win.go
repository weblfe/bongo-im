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

// CMFTurntableWin is an object representing the database table.
type CMFTurntableWin struct {
	ID      uint64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	Logid   int64  `boil:"logid" json:"logid" toml:"logid" yaml:"logid"`
	UID     int64  `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Type    bool   `boil:"type" json:"type" toml:"type" yaml:"type"`
	TypeVal string `boil:"type_val" json:"type_val" toml:"type_val" yaml:"type_val"`
	Thumb   string `boil:"thumb" json:"thumb" toml:"thumb" yaml:"thumb"`
	Nums    int    `boil:"nums" json:"nums" toml:"nums" yaml:"nums"`
	Addtime int64  `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Status  bool   `boil:"status" json:"status" toml:"status" yaml:"status"`
	Uptime  int64  `boil:"uptime" json:"uptime" toml:"uptime" yaml:"uptime"`

	R *cmfTurntableWinR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfTurntableWinL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFTurntableWinColumns = struct {
	ID      string
	Logid   string
	UID     string
	Type    string
	TypeVal string
	Thumb   string
	Nums    string
	Addtime string
	Status  string
	Uptime  string
}{
	ID:      "id",
	Logid:   "logid",
	UID:     "uid",
	Type:    "type",
	TypeVal: "type_val",
	Thumb:   "thumb",
	Nums:    "nums",
	Addtime: "addtime",
	Status:  "status",
	Uptime:  "uptime",
}

// Generated where

var CMFTurntableWinWhere = struct {
	ID      whereHelperuint64
	Logid   whereHelperint64
	UID     whereHelperint64
	Type    whereHelperbool
	TypeVal whereHelperstring
	Thumb   whereHelperstring
	Nums    whereHelperint
	Addtime whereHelperint64
	Status  whereHelperbool
	Uptime  whereHelperint64
}{
	ID:      whereHelperuint64{field: "`cmf_turntable_win`.`id`"},
	Logid:   whereHelperint64{field: "`cmf_turntable_win`.`logid`"},
	UID:     whereHelperint64{field: "`cmf_turntable_win`.`uid`"},
	Type:    whereHelperbool{field: "`cmf_turntable_win`.`type`"},
	TypeVal: whereHelperstring{field: "`cmf_turntable_win`.`type_val`"},
	Thumb:   whereHelperstring{field: "`cmf_turntable_win`.`thumb`"},
	Nums:    whereHelperint{field: "`cmf_turntable_win`.`nums`"},
	Addtime: whereHelperint64{field: "`cmf_turntable_win`.`addtime`"},
	Status:  whereHelperbool{field: "`cmf_turntable_win`.`status`"},
	Uptime:  whereHelperint64{field: "`cmf_turntable_win`.`uptime`"},
}

// CMFTurntableWinRels is where relationship names are stored.
var CMFTurntableWinRels = struct {
}{}

// cmfTurntableWinR is where relationships are stored.
type cmfTurntableWinR struct {
}

// NewStruct creates a new relationship struct
func (*cmfTurntableWinR) NewStruct() *cmfTurntableWinR {
	return &cmfTurntableWinR{}
}

// cmfTurntableWinL is where Load methods for each relationship are stored.
type cmfTurntableWinL struct{}

var (
	cmfTurntableWinAllColumns            = []string{"id", "logid", "uid", "type", "type_val", "thumb", "nums", "addtime", "status", "uptime"}
	cmfTurntableWinColumnsWithoutDefault = []string{"thumb"}
	cmfTurntableWinColumnsWithDefault    = []string{"id", "logid", "uid", "type", "type_val", "nums", "addtime", "status", "uptime"}
	cmfTurntableWinPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFTurntableWinSlice is an alias for a slice of pointers to CMFTurntableWin.
	// This should generally be used opposed to []CMFTurntableWin.
	CMFTurntableWinSlice []*CMFTurntableWin
	// CMFTurntableWinHook is the signature for custom CMFTurntableWin hook methods
	CMFTurntableWinHook func(context.Context, boil.ContextExecutor, *CMFTurntableWin) error

	cmfTurntableWinQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfTurntableWinType                 = reflect.TypeOf(&CMFTurntableWin{})
	cmfTurntableWinMapping              = queries.MakeStructMapping(cmfTurntableWinType)
	cmfTurntableWinPrimaryKeyMapping, _ = queries.BindMapping(cmfTurntableWinType, cmfTurntableWinMapping, cmfTurntableWinPrimaryKeyColumns)
	cmfTurntableWinInsertCacheMut       sync.RWMutex
	cmfTurntableWinInsertCache          = make(map[string]insertCache)
	cmfTurntableWinUpdateCacheMut       sync.RWMutex
	cmfTurntableWinUpdateCache          = make(map[string]updateCache)
	cmfTurntableWinUpsertCacheMut       sync.RWMutex
	cmfTurntableWinUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfTurntableWinBeforeInsertHooks []CMFTurntableWinHook
var cmfTurntableWinBeforeUpdateHooks []CMFTurntableWinHook
var cmfTurntableWinBeforeDeleteHooks []CMFTurntableWinHook
var cmfTurntableWinBeforeUpsertHooks []CMFTurntableWinHook

var cmfTurntableWinAfterInsertHooks []CMFTurntableWinHook
var cmfTurntableWinAfterSelectHooks []CMFTurntableWinHook
var cmfTurntableWinAfterUpdateHooks []CMFTurntableWinHook
var cmfTurntableWinAfterDeleteHooks []CMFTurntableWinHook
var cmfTurntableWinAfterUpsertHooks []CMFTurntableWinHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFTurntableWin) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFTurntableWin) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFTurntableWin) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFTurntableWin) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFTurntableWin) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFTurntableWin) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFTurntableWin) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFTurntableWin) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFTurntableWin) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfTurntableWinAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFTurntableWinHook registers your hook function for all future operations.
func AddCMFTurntableWinHook(hookPoint boil.HookPoint, cmfTurntableWinHook CMFTurntableWinHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfTurntableWinBeforeInsertHooks = append(cmfTurntableWinBeforeInsertHooks, cmfTurntableWinHook)
	case boil.BeforeUpdateHook:
		cmfTurntableWinBeforeUpdateHooks = append(cmfTurntableWinBeforeUpdateHooks, cmfTurntableWinHook)
	case boil.BeforeDeleteHook:
		cmfTurntableWinBeforeDeleteHooks = append(cmfTurntableWinBeforeDeleteHooks, cmfTurntableWinHook)
	case boil.BeforeUpsertHook:
		cmfTurntableWinBeforeUpsertHooks = append(cmfTurntableWinBeforeUpsertHooks, cmfTurntableWinHook)
	case boil.AfterInsertHook:
		cmfTurntableWinAfterInsertHooks = append(cmfTurntableWinAfterInsertHooks, cmfTurntableWinHook)
	case boil.AfterSelectHook:
		cmfTurntableWinAfterSelectHooks = append(cmfTurntableWinAfterSelectHooks, cmfTurntableWinHook)
	case boil.AfterUpdateHook:
		cmfTurntableWinAfterUpdateHooks = append(cmfTurntableWinAfterUpdateHooks, cmfTurntableWinHook)
	case boil.AfterDeleteHook:
		cmfTurntableWinAfterDeleteHooks = append(cmfTurntableWinAfterDeleteHooks, cmfTurntableWinHook)
	case boil.AfterUpsertHook:
		cmfTurntableWinAfterUpsertHooks = append(cmfTurntableWinAfterUpsertHooks, cmfTurntableWinHook)
	}
}

// One returns a single cmfTurntableWin record from the query.
func (q cmfTurntableWinQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFTurntableWin, error) {
	o := &CMFTurntableWin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_turntable_win")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFTurntableWin records from the query.
func (q cmfTurntableWinQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFTurntableWinSlice, error) {
	var o []*CMFTurntableWin

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFTurntableWin slice")
	}

	if len(cmfTurntableWinAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFTurntableWin records in the query.
func (q cmfTurntableWinQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_turntable_win rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfTurntableWinQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_turntable_win exists")
	}

	return count > 0, nil
}

// CMFTurntableWins retrieves all the records using an executor.
func CMFTurntableWins(mods ...qm.QueryMod) cmfTurntableWinQuery {
	mods = append(mods, qm.From("`cmf_turntable_win`"))
	return cmfTurntableWinQuery{NewQuery(mods...)}
}

// FindCMFTurntableWin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFTurntableWin(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*CMFTurntableWin, error) {
	cmfTurntableWinObj := &CMFTurntableWin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_turntable_win` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfTurntableWinObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_turntable_win")
	}

	return cmfTurntableWinObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFTurntableWin) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_turntable_win provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfTurntableWinColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfTurntableWinInsertCacheMut.RLock()
	cache, cached := cmfTurntableWinInsertCache[key]
	cmfTurntableWinInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfTurntableWinAllColumns,
			cmfTurntableWinColumnsWithDefault,
			cmfTurntableWinColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfTurntableWinType, cmfTurntableWinMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfTurntableWinType, cmfTurntableWinMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_turntable_win` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_turntable_win` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_turntable_win` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfTurntableWinPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_turntable_win")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfTurntableWinMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_turntable_win")
	}

CacheNoHooks:
	if !cached {
		cmfTurntableWinInsertCacheMut.Lock()
		cmfTurntableWinInsertCache[key] = cache
		cmfTurntableWinInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFTurntableWin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFTurntableWin) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfTurntableWinUpdateCacheMut.RLock()
	cache, cached := cmfTurntableWinUpdateCache[key]
	cmfTurntableWinUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfTurntableWinAllColumns,
			cmfTurntableWinPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_turntable_win, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_turntable_win` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfTurntableWinPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfTurntableWinType, cmfTurntableWinMapping, append(wl, cmfTurntableWinPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_turntable_win row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_turntable_win")
	}

	if !cached {
		cmfTurntableWinUpdateCacheMut.Lock()
		cmfTurntableWinUpdateCache[key] = cache
		cmfTurntableWinUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfTurntableWinQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_turntable_win")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_turntable_win")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFTurntableWinSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfTurntableWinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_turntable_win` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfTurntableWinPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfTurntableWin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfTurntableWin")
	}
	return rowsAff, nil
}

var mySQLCMFTurntableWinUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFTurntableWin) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_turntable_win provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfTurntableWinColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFTurntableWinUniqueColumns, o)

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

	cmfTurntableWinUpsertCacheMut.RLock()
	cache, cached := cmfTurntableWinUpsertCache[key]
	cmfTurntableWinUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfTurntableWinAllColumns,
			cmfTurntableWinColumnsWithDefault,
			cmfTurntableWinColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfTurntableWinAllColumns,
			cmfTurntableWinPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_turntable_win, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_turntable_win`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_turntable_win` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfTurntableWinType, cmfTurntableWinMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfTurntableWinType, cmfTurntableWinMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_turntable_win")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfTurntableWinMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfTurntableWinType, cmfTurntableWinMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_turntable_win")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_turntable_win")
	}

CacheNoHooks:
	if !cached {
		cmfTurntableWinUpsertCacheMut.Lock()
		cmfTurntableWinUpsertCache[key] = cache
		cmfTurntableWinUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFTurntableWin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFTurntableWin) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFTurntableWin provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfTurntableWinPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_turntable_win` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_turntable_win")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_turntable_win")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfTurntableWinQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfTurntableWinQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_turntable_win")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_turntable_win")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFTurntableWinSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfTurntableWinBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfTurntableWinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_turntable_win` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfTurntableWinPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfTurntableWin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_turntable_win")
	}

	if len(cmfTurntableWinAfterDeleteHooks) != 0 {
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
func (o *CMFTurntableWin) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFTurntableWin(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFTurntableWinSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFTurntableWinSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfTurntableWinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_turntable_win`.* FROM `cmf_turntable_win` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfTurntableWinPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFTurntableWinSlice")
	}

	*o = slice

	return nil
}

// CMFTurntableWinExists checks if the CMFTurntableWin row exists.
func CMFTurntableWinExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_turntable_win` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_turntable_win exists")
	}

	return exists, nil
}
