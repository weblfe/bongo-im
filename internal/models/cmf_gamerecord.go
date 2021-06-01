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

// CMFGamerecord is an object representing the database table.
type CMFGamerecord struct {
	ID      int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Action  null.Bool `boil:"action" json:"action,omitempty" toml:"action" yaml:"action,omitempty"`
	UID     null.Int  `boil:"uid" json:"uid,omitempty" toml:"uid" yaml:"uid,omitempty"`
	Liveuid null.Int  `boil:"liveuid" json:"liveuid,omitempty" toml:"liveuid" yaml:"liveuid,omitempty"`
	Gameid  null.Int  `boil:"gameid" json:"gameid,omitempty" toml:"gameid" yaml:"gameid,omitempty"`
	Coin1   null.Int  `boil:"coin_1" json:"coin_1,omitempty" toml:"coin_1" yaml:"coin_1,omitempty"`
	Coin2   null.Int  `boil:"coin_2" json:"coin_2,omitempty" toml:"coin_2" yaml:"coin_2,omitempty"`
	Coin3   null.Int  `boil:"coin_3" json:"coin_3,omitempty" toml:"coin_3" yaml:"coin_3,omitempty"`
	Coin4   null.Int  `boil:"coin_4" json:"coin_4,omitempty" toml:"coin_4" yaml:"coin_4,omitempty"`
	Coin5   null.Int  `boil:"coin_5" json:"coin_5,omitempty" toml:"coin_5" yaml:"coin_5,omitempty"`
	Coin6   null.Int  `boil:"coin_6" json:"coin_6,omitempty" toml:"coin_6" yaml:"coin_6,omitempty"`
	Status  null.Bool `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	Addtime null.Int  `boil:"addtime" json:"addtime,omitempty" toml:"addtime" yaml:"addtime,omitempty"`

	R *cmfGamerecordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfGamerecordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFGamerecordColumns = struct {
	ID      string
	Action  string
	UID     string
	Liveuid string
	Gameid  string
	Coin1   string
	Coin2   string
	Coin3   string
	Coin4   string
	Coin5   string
	Coin6   string
	Status  string
	Addtime string
}{
	ID:      "id",
	Action:  "action",
	UID:     "uid",
	Liveuid: "liveuid",
	Gameid:  "gameid",
	Coin1:   "coin_1",
	Coin2:   "coin_2",
	Coin3:   "coin_3",
	Coin4:   "coin_4",
	Coin5:   "coin_5",
	Coin6:   "coin_6",
	Status:  "status",
	Addtime: "addtime",
}

// Generated where

var CMFGamerecordWhere = struct {
	ID      whereHelperint
	Action  whereHelpernull_Bool
	UID     whereHelpernull_Int
	Liveuid whereHelpernull_Int
	Gameid  whereHelpernull_Int
	Coin1   whereHelpernull_Int
	Coin2   whereHelpernull_Int
	Coin3   whereHelpernull_Int
	Coin4   whereHelpernull_Int
	Coin5   whereHelpernull_Int
	Coin6   whereHelpernull_Int
	Status  whereHelpernull_Bool
	Addtime whereHelpernull_Int
}{
	ID:      whereHelperint{field: "`cmf_gamerecord`.`id`"},
	Action:  whereHelpernull_Bool{field: "`cmf_gamerecord`.`action`"},
	UID:     whereHelpernull_Int{field: "`cmf_gamerecord`.`uid`"},
	Liveuid: whereHelpernull_Int{field: "`cmf_gamerecord`.`liveuid`"},
	Gameid:  whereHelpernull_Int{field: "`cmf_gamerecord`.`gameid`"},
	Coin1:   whereHelpernull_Int{field: "`cmf_gamerecord`.`coin_1`"},
	Coin2:   whereHelpernull_Int{field: "`cmf_gamerecord`.`coin_2`"},
	Coin3:   whereHelpernull_Int{field: "`cmf_gamerecord`.`coin_3`"},
	Coin4:   whereHelpernull_Int{field: "`cmf_gamerecord`.`coin_4`"},
	Coin5:   whereHelpernull_Int{field: "`cmf_gamerecord`.`coin_5`"},
	Coin6:   whereHelpernull_Int{field: "`cmf_gamerecord`.`coin_6`"},
	Status:  whereHelpernull_Bool{field: "`cmf_gamerecord`.`status`"},
	Addtime: whereHelpernull_Int{field: "`cmf_gamerecord`.`addtime`"},
}

// CMFGamerecordRels is where relationship names are stored.
var CMFGamerecordRels = struct {
}{}

// cmfGamerecordR is where relationships are stored.
type cmfGamerecordR struct {
}

// NewStruct creates a new relationship struct
func (*cmfGamerecordR) NewStruct() *cmfGamerecordR {
	return &cmfGamerecordR{}
}

// cmfGamerecordL is where Load methods for each relationship are stored.
type cmfGamerecordL struct{}

var (
	cmfGamerecordAllColumns            = []string{"id", "action", "uid", "liveuid", "gameid", "coin_1", "coin_2", "coin_3", "coin_4", "coin_5", "coin_6", "status", "addtime"}
	cmfGamerecordColumnsWithoutDefault = []string{}
	cmfGamerecordColumnsWithDefault    = []string{"id", "action", "uid", "liveuid", "gameid", "coin_1", "coin_2", "coin_3", "coin_4", "coin_5", "coin_6", "status", "addtime"}
	cmfGamerecordPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFGamerecordSlice is an alias for a slice of pointers to CMFGamerecord.
	// This should generally be used opposed to []CMFGamerecord.
	CMFGamerecordSlice []*CMFGamerecord
	// CMFGamerecordHook is the signature for custom CMFGamerecord hook methods
	CMFGamerecordHook func(context.Context, boil.ContextExecutor, *CMFGamerecord) error

	cmfGamerecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfGamerecordType                 = reflect.TypeOf(&CMFGamerecord{})
	cmfGamerecordMapping              = queries.MakeStructMapping(cmfGamerecordType)
	cmfGamerecordPrimaryKeyMapping, _ = queries.BindMapping(cmfGamerecordType, cmfGamerecordMapping, cmfGamerecordPrimaryKeyColumns)
	cmfGamerecordInsertCacheMut       sync.RWMutex
	cmfGamerecordInsertCache          = make(map[string]insertCache)
	cmfGamerecordUpdateCacheMut       sync.RWMutex
	cmfGamerecordUpdateCache          = make(map[string]updateCache)
	cmfGamerecordUpsertCacheMut       sync.RWMutex
	cmfGamerecordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfGamerecordBeforeInsertHooks []CMFGamerecordHook
var cmfGamerecordBeforeUpdateHooks []CMFGamerecordHook
var cmfGamerecordBeforeDeleteHooks []CMFGamerecordHook
var cmfGamerecordBeforeUpsertHooks []CMFGamerecordHook

var cmfGamerecordAfterInsertHooks []CMFGamerecordHook
var cmfGamerecordAfterSelectHooks []CMFGamerecordHook
var cmfGamerecordAfterUpdateHooks []CMFGamerecordHook
var cmfGamerecordAfterDeleteHooks []CMFGamerecordHook
var cmfGamerecordAfterUpsertHooks []CMFGamerecordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFGamerecord) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFGamerecord) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFGamerecord) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFGamerecord) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFGamerecord) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFGamerecord) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFGamerecord) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFGamerecord) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFGamerecord) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGamerecordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFGamerecordHook registers your hook function for all future operations.
func AddCMFGamerecordHook(hookPoint boil.HookPoint, cmfGamerecordHook CMFGamerecordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfGamerecordBeforeInsertHooks = append(cmfGamerecordBeforeInsertHooks, cmfGamerecordHook)
	case boil.BeforeUpdateHook:
		cmfGamerecordBeforeUpdateHooks = append(cmfGamerecordBeforeUpdateHooks, cmfGamerecordHook)
	case boil.BeforeDeleteHook:
		cmfGamerecordBeforeDeleteHooks = append(cmfGamerecordBeforeDeleteHooks, cmfGamerecordHook)
	case boil.BeforeUpsertHook:
		cmfGamerecordBeforeUpsertHooks = append(cmfGamerecordBeforeUpsertHooks, cmfGamerecordHook)
	case boil.AfterInsertHook:
		cmfGamerecordAfterInsertHooks = append(cmfGamerecordAfterInsertHooks, cmfGamerecordHook)
	case boil.AfterSelectHook:
		cmfGamerecordAfterSelectHooks = append(cmfGamerecordAfterSelectHooks, cmfGamerecordHook)
	case boil.AfterUpdateHook:
		cmfGamerecordAfterUpdateHooks = append(cmfGamerecordAfterUpdateHooks, cmfGamerecordHook)
	case boil.AfterDeleteHook:
		cmfGamerecordAfterDeleteHooks = append(cmfGamerecordAfterDeleteHooks, cmfGamerecordHook)
	case boil.AfterUpsertHook:
		cmfGamerecordAfterUpsertHooks = append(cmfGamerecordAfterUpsertHooks, cmfGamerecordHook)
	}
}

// One returns a single cmfGamerecord record from the query.
func (q cmfGamerecordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFGamerecord, error) {
	o := &CMFGamerecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_gamerecord")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFGamerecord records from the query.
func (q cmfGamerecordQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFGamerecordSlice, error) {
	var o []*CMFGamerecord

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFGamerecord slice")
	}

	if len(cmfGamerecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFGamerecord records in the query.
func (q cmfGamerecordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_gamerecord rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfGamerecordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_gamerecord exists")
	}

	return count > 0, nil
}

// CMFGamerecords retrieves all the records using an executor.
func CMFGamerecords(mods ...qm.QueryMod) cmfGamerecordQuery {
	mods = append(mods, qm.From("`cmf_gamerecord`"))
	return cmfGamerecordQuery{NewQuery(mods...)}
}

// FindCMFGamerecord retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFGamerecord(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFGamerecord, error) {
	cmfGamerecordObj := &CMFGamerecord{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_gamerecord` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfGamerecordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_gamerecord")
	}

	return cmfGamerecordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFGamerecord) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_gamerecord provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfGamerecordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfGamerecordInsertCacheMut.RLock()
	cache, cached := cmfGamerecordInsertCache[key]
	cmfGamerecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfGamerecordAllColumns,
			cmfGamerecordColumnsWithDefault,
			cmfGamerecordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfGamerecordType, cmfGamerecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfGamerecordType, cmfGamerecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_gamerecord` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_gamerecord` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_gamerecord` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfGamerecordPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_gamerecord")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfGamerecordMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_gamerecord")
	}

CacheNoHooks:
	if !cached {
		cmfGamerecordInsertCacheMut.Lock()
		cmfGamerecordInsertCache[key] = cache
		cmfGamerecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFGamerecord.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFGamerecord) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfGamerecordUpdateCacheMut.RLock()
	cache, cached := cmfGamerecordUpdateCache[key]
	cmfGamerecordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfGamerecordAllColumns,
			cmfGamerecordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_gamerecord, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_gamerecord` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfGamerecordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfGamerecordType, cmfGamerecordMapping, append(wl, cmfGamerecordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_gamerecord row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_gamerecord")
	}

	if !cached {
		cmfGamerecordUpdateCacheMut.Lock()
		cmfGamerecordUpdateCache[key] = cache
		cmfGamerecordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfGamerecordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_gamerecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_gamerecord")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFGamerecordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGamerecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_gamerecord` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGamerecordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfGamerecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfGamerecord")
	}
	return rowsAff, nil
}

var mySQLCMFGamerecordUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFGamerecord) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_gamerecord provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfGamerecordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFGamerecordUniqueColumns, o)

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

	cmfGamerecordUpsertCacheMut.RLock()
	cache, cached := cmfGamerecordUpsertCache[key]
	cmfGamerecordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfGamerecordAllColumns,
			cmfGamerecordColumnsWithDefault,
			cmfGamerecordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfGamerecordAllColumns,
			cmfGamerecordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_gamerecord, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_gamerecord`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_gamerecord` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfGamerecordType, cmfGamerecordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfGamerecordType, cmfGamerecordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_gamerecord")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfGamerecordMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfGamerecordType, cmfGamerecordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_gamerecord")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_gamerecord")
	}

CacheNoHooks:
	if !cached {
		cmfGamerecordUpsertCacheMut.Lock()
		cmfGamerecordUpsertCache[key] = cache
		cmfGamerecordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFGamerecord record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFGamerecord) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFGamerecord provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfGamerecordPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_gamerecord` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_gamerecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_gamerecord")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfGamerecordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfGamerecordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_gamerecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_gamerecord")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFGamerecordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfGamerecordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGamerecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_gamerecord` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGamerecordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfGamerecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_gamerecord")
	}

	if len(cmfGamerecordAfterDeleteHooks) != 0 {
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
func (o *CMFGamerecord) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFGamerecord(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFGamerecordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFGamerecordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGamerecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_gamerecord`.* FROM `cmf_gamerecord` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGamerecordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFGamerecordSlice")
	}

	*o = slice

	return nil
}

// CMFGamerecordExists checks if the CMFGamerecord row exists.
func CMFGamerecordExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_gamerecord` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_gamerecord exists")
	}

	return exists, nil
}
