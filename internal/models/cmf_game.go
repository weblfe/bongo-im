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

// CMFGame is an object representing the database table.
type CMFGame struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Action       null.Bool   `boil:"action" json:"action,omitempty" toml:"action" yaml:"action,omitempty"`
	Liveuid      null.Int    `boil:"liveuid" json:"liveuid,omitempty" toml:"liveuid" yaml:"liveuid,omitempty"`
	Bankerid     null.Int    `boil:"bankerid" json:"bankerid,omitempty" toml:"bankerid" yaml:"bankerid,omitempty"`
	Stream       null.String `boil:"stream" json:"stream,omitempty" toml:"stream" yaml:"stream,omitempty"`
	Starttime    null.Int    `boil:"starttime" json:"starttime,omitempty" toml:"starttime" yaml:"starttime,omitempty"`
	Endtime      null.Int    `boil:"endtime" json:"endtime,omitempty" toml:"endtime" yaml:"endtime,omitempty"`
	Result       null.String `boil:"result" json:"result,omitempty" toml:"result" yaml:"result,omitempty"`
	State        null.Bool   `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
	BankerProfit null.Int    `boil:"banker_profit" json:"banker_profit,omitempty" toml:"banker_profit" yaml:"banker_profit,omitempty"`
	BankerCard   null.String `boil:"banker_card" json:"banker_card,omitempty" toml:"banker_card" yaml:"banker_card,omitempty"`
	Isintervene  null.Bool   `boil:"isintervene" json:"isintervene,omitempty" toml:"isintervene" yaml:"isintervene,omitempty"`

	R *cmfGameR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfGameL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFGameColumns = struct {
	ID           string
	Action       string
	Liveuid      string
	Bankerid     string
	Stream       string
	Starttime    string
	Endtime      string
	Result       string
	State        string
	BankerProfit string
	BankerCard   string
	Isintervene  string
}{
	ID:           "id",
	Action:       "action",
	Liveuid:      "liveuid",
	Bankerid:     "bankerid",
	Stream:       "stream",
	Starttime:    "starttime",
	Endtime:      "endtime",
	Result:       "result",
	State:        "state",
	BankerProfit: "banker_profit",
	BankerCard:   "banker_card",
	Isintervene:  "isintervene",
}

// Generated where

var CMFGameWhere = struct {
	ID           whereHelperint
	Action       whereHelpernull_Bool
	Liveuid      whereHelpernull_Int
	Bankerid     whereHelpernull_Int
	Stream       whereHelpernull_String
	Starttime    whereHelpernull_Int
	Endtime      whereHelpernull_Int
	Result       whereHelpernull_String
	State        whereHelpernull_Bool
	BankerProfit whereHelpernull_Int
	BankerCard   whereHelpernull_String
	Isintervene  whereHelpernull_Bool
}{
	ID:           whereHelperint{field: "`cmf_game`.`id`"},
	Action:       whereHelpernull_Bool{field: "`cmf_game`.`action`"},
	Liveuid:      whereHelpernull_Int{field: "`cmf_game`.`liveuid`"},
	Bankerid:     whereHelpernull_Int{field: "`cmf_game`.`bankerid`"},
	Stream:       whereHelpernull_String{field: "`cmf_game`.`stream`"},
	Starttime:    whereHelpernull_Int{field: "`cmf_game`.`starttime`"},
	Endtime:      whereHelpernull_Int{field: "`cmf_game`.`endtime`"},
	Result:       whereHelpernull_String{field: "`cmf_game`.`result`"},
	State:        whereHelpernull_Bool{field: "`cmf_game`.`state`"},
	BankerProfit: whereHelpernull_Int{field: "`cmf_game`.`banker_profit`"},
	BankerCard:   whereHelpernull_String{field: "`cmf_game`.`banker_card`"},
	Isintervene:  whereHelpernull_Bool{field: "`cmf_game`.`isintervene`"},
}

// CMFGameRels is where relationship names are stored.
var CMFGameRels = struct {
}{}

// cmfGameR is where relationships are stored.
type cmfGameR struct {
}

// NewStruct creates a new relationship struct
func (*cmfGameR) NewStruct() *cmfGameR {
	return &cmfGameR{}
}

// cmfGameL is where Load methods for each relationship are stored.
type cmfGameL struct{}

var (
	cmfGameAllColumns            = []string{"id", "action", "liveuid", "bankerid", "stream", "starttime", "endtime", "result", "state", "banker_profit", "banker_card", "isintervene"}
	cmfGameColumnsWithoutDefault = []string{"stream", "banker_card"}
	cmfGameColumnsWithDefault    = []string{"id", "action", "liveuid", "bankerid", "starttime", "endtime", "result", "state", "banker_profit", "isintervene"}
	cmfGamePrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFGameSlice is an alias for a slice of pointers to CMFGame.
	// This should generally be used opposed to []CMFGame.
	CMFGameSlice []*CMFGame
	// CMFGameHook is the signature for custom CMFGame hook methods
	CMFGameHook func(context.Context, boil.ContextExecutor, *CMFGame) error

	cmfGameQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfGameType                 = reflect.TypeOf(&CMFGame{})
	cmfGameMapping              = queries.MakeStructMapping(cmfGameType)
	cmfGamePrimaryKeyMapping, _ = queries.BindMapping(cmfGameType, cmfGameMapping, cmfGamePrimaryKeyColumns)
	cmfGameInsertCacheMut       sync.RWMutex
	cmfGameInsertCache          = make(map[string]insertCache)
	cmfGameUpdateCacheMut       sync.RWMutex
	cmfGameUpdateCache          = make(map[string]updateCache)
	cmfGameUpsertCacheMut       sync.RWMutex
	cmfGameUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfGameBeforeInsertHooks []CMFGameHook
var cmfGameBeforeUpdateHooks []CMFGameHook
var cmfGameBeforeDeleteHooks []CMFGameHook
var cmfGameBeforeUpsertHooks []CMFGameHook

var cmfGameAfterInsertHooks []CMFGameHook
var cmfGameAfterSelectHooks []CMFGameHook
var cmfGameAfterUpdateHooks []CMFGameHook
var cmfGameAfterDeleteHooks []CMFGameHook
var cmfGameAfterUpsertHooks []CMFGameHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFGame) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFGame) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFGame) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFGame) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFGame) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFGame) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFGame) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFGame) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFGame) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGameAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFGameHook registers your hook function for all future operations.
func AddCMFGameHook(hookPoint boil.HookPoint, cmfGameHook CMFGameHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfGameBeforeInsertHooks = append(cmfGameBeforeInsertHooks, cmfGameHook)
	case boil.BeforeUpdateHook:
		cmfGameBeforeUpdateHooks = append(cmfGameBeforeUpdateHooks, cmfGameHook)
	case boil.BeforeDeleteHook:
		cmfGameBeforeDeleteHooks = append(cmfGameBeforeDeleteHooks, cmfGameHook)
	case boil.BeforeUpsertHook:
		cmfGameBeforeUpsertHooks = append(cmfGameBeforeUpsertHooks, cmfGameHook)
	case boil.AfterInsertHook:
		cmfGameAfterInsertHooks = append(cmfGameAfterInsertHooks, cmfGameHook)
	case boil.AfterSelectHook:
		cmfGameAfterSelectHooks = append(cmfGameAfterSelectHooks, cmfGameHook)
	case boil.AfterUpdateHook:
		cmfGameAfterUpdateHooks = append(cmfGameAfterUpdateHooks, cmfGameHook)
	case boil.AfterDeleteHook:
		cmfGameAfterDeleteHooks = append(cmfGameAfterDeleteHooks, cmfGameHook)
	case boil.AfterUpsertHook:
		cmfGameAfterUpsertHooks = append(cmfGameAfterUpsertHooks, cmfGameHook)
	}
}

// One returns a single cmfGame record from the query.
func (q cmfGameQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFGame, error) {
	o := &CMFGame{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_game")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFGame records from the query.
func (q cmfGameQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFGameSlice, error) {
	var o []*CMFGame

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFGame slice")
	}

	if len(cmfGameAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFGame records in the query.
func (q cmfGameQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_game rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfGameQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_game exists")
	}

	return count > 0, nil
}

// CMFGames retrieves all the records using an executor.
func CMFGames(mods ...qm.QueryMod) cmfGameQuery {
	mods = append(mods, qm.From("`cmf_game`"))
	return cmfGameQuery{NewQuery(mods...)}
}

// FindCMFGame retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFGame(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFGame, error) {
	cmfGameObj := &CMFGame{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_game` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfGameObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_game")
	}

	return cmfGameObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFGame) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_game provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfGameColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfGameInsertCacheMut.RLock()
	cache, cached := cmfGameInsertCache[key]
	cmfGameInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfGameAllColumns,
			cmfGameColumnsWithDefault,
			cmfGameColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfGameType, cmfGameMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfGameType, cmfGameMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_game` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_game` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_game` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfGamePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_game")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfGameMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_game")
	}

CacheNoHooks:
	if !cached {
		cmfGameInsertCacheMut.Lock()
		cmfGameInsertCache[key] = cache
		cmfGameInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFGame.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFGame) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfGameUpdateCacheMut.RLock()
	cache, cached := cmfGameUpdateCache[key]
	cmfGameUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfGameAllColumns,
			cmfGamePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_game, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_game` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfGamePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfGameType, cmfGameMapping, append(wl, cmfGamePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_game row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_game")
	}

	if !cached {
		cmfGameUpdateCacheMut.Lock()
		cmfGameUpdateCache[key] = cache
		cmfGameUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfGameQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_game")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_game")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFGameSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_game` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGamePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfGame slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfGame")
	}
	return rowsAff, nil
}

var mySQLCMFGameUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFGame) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_game provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfGameColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFGameUniqueColumns, o)

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

	cmfGameUpsertCacheMut.RLock()
	cache, cached := cmfGameUpsertCache[key]
	cmfGameUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfGameAllColumns,
			cmfGameColumnsWithDefault,
			cmfGameColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfGameAllColumns,
			cmfGamePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_game, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_game`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_game` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfGameType, cmfGameMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfGameType, cmfGameMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_game")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfGameMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfGameType, cmfGameMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_game")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_game")
	}

CacheNoHooks:
	if !cached {
		cmfGameUpsertCacheMut.Lock()
		cmfGameUpsertCache[key] = cache
		cmfGameUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFGame record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFGame) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFGame provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfGamePrimaryKeyMapping)
	sql := "DELETE FROM `cmf_game` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_game")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_game")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfGameQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfGameQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_game")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_game")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFGameSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfGameBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_game` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGamePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfGame slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_game")
	}

	if len(cmfGameAfterDeleteHooks) != 0 {
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
func (o *CMFGame) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFGame(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFGameSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFGameSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_game`.* FROM `cmf_game` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGamePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFGameSlice")
	}

	*o = slice

	return nil
}

// CMFGameExists checks if the CMFGame row exists.
func CMFGameExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_game` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_game exists")
	}

	return exists, nil
}
