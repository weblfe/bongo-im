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

// CMFUserScorerecord is an object representing the database table.
type CMFUserScorerecord struct {
	ID         int  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Type       bool `boil:"type" json:"type" toml:"type" yaml:"type"`
	Action     bool `boil:"action" json:"action" toml:"action" yaml:"action"`
	UID        int  `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Touid      int  `boil:"touid" json:"touid" toml:"touid" yaml:"touid"`
	Giftid     int  `boil:"giftid" json:"giftid" toml:"giftid" yaml:"giftid"`
	Giftcount  int  `boil:"giftcount" json:"giftcount" toml:"giftcount" yaml:"giftcount"`
	Totalcoin  int  `boil:"totalcoin" json:"totalcoin" toml:"totalcoin" yaml:"totalcoin"`
	Showid     int  `boil:"showid" json:"showid" toml:"showid" yaml:"showid"`
	Addtime    int  `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	GameAction bool `boil:"game_action" json:"game_action" toml:"game_action" yaml:"game_action"`

	R *cmfUserScorerecordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfUserScorerecordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFUserScorerecordColumns = struct {
	ID         string
	Type       string
	Action     string
	UID        string
	Touid      string
	Giftid     string
	Giftcount  string
	Totalcoin  string
	Showid     string
	Addtime    string
	GameAction string
}{
	ID:         "id",
	Type:       "type",
	Action:     "action",
	UID:        "uid",
	Touid:      "touid",
	Giftid:     "giftid",
	Giftcount:  "giftcount",
	Totalcoin:  "totalcoin",
	Showid:     "showid",
	Addtime:    "addtime",
	GameAction: "game_action",
}

// Generated where

var CMFUserScorerecordWhere = struct {
	ID         whereHelperint
	Type       whereHelperbool
	Action     whereHelperbool
	UID        whereHelperint
	Touid      whereHelperint
	Giftid     whereHelperint
	Giftcount  whereHelperint
	Totalcoin  whereHelperint
	Showid     whereHelperint
	Addtime    whereHelperint
	GameAction whereHelperbool
}{
	ID:         whereHelperint{field: "`cmf_user_scorerecord`.`id`"},
	Type:       whereHelperbool{field: "`cmf_user_scorerecord`.`type`"},
	Action:     whereHelperbool{field: "`cmf_user_scorerecord`.`action`"},
	UID:        whereHelperint{field: "`cmf_user_scorerecord`.`uid`"},
	Touid:      whereHelperint{field: "`cmf_user_scorerecord`.`touid`"},
	Giftid:     whereHelperint{field: "`cmf_user_scorerecord`.`giftid`"},
	Giftcount:  whereHelperint{field: "`cmf_user_scorerecord`.`giftcount`"},
	Totalcoin:  whereHelperint{field: "`cmf_user_scorerecord`.`totalcoin`"},
	Showid:     whereHelperint{field: "`cmf_user_scorerecord`.`showid`"},
	Addtime:    whereHelperint{field: "`cmf_user_scorerecord`.`addtime`"},
	GameAction: whereHelperbool{field: "`cmf_user_scorerecord`.`game_action`"},
}

// CMFUserScorerecordRels is where relationship names are stored.
var CMFUserScorerecordRels = struct {
}{}

// cmfUserScorerecordR is where relationships are stored.
type cmfUserScorerecordR struct {
}

// NewStruct creates a new relationship struct
func (*cmfUserScorerecordR) NewStruct() *cmfUserScorerecordR {
	return &cmfUserScorerecordR{}
}

// cmfUserScorerecordL is where Load methods for each relationship are stored.
type cmfUserScorerecordL struct{}

var (
	cmfUserScorerecordAllColumns            = []string{"id", "type", "action", "uid", "touid", "giftid", "giftcount", "totalcoin", "showid", "addtime", "game_action"}
	cmfUserScorerecordColumnsWithoutDefault = []string{}
	cmfUserScorerecordColumnsWithDefault    = []string{"id", "type", "action", "uid", "touid", "giftid", "giftcount", "totalcoin", "showid", "addtime", "game_action"}
	cmfUserScorerecordPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFUserScorerecordSlice is an alias for a slice of pointers to CMFUserScorerecord.
	// This should generally be used opposed to []CMFUserScorerecord.
	CMFUserScorerecordSlice []*CMFUserScorerecord
	// CMFUserScorerecordHook is the signature for custom CMFUserScorerecord hook methods
	CMFUserScorerecordHook func(context.Context, boil.ContextExecutor, *CMFUserScorerecord) error

	cmfUserScorerecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfUserScorerecordType                 = reflect.TypeOf(&CMFUserScorerecord{})
	cmfUserScorerecordMapping              = queries.MakeStructMapping(cmfUserScorerecordType)
	cmfUserScorerecordPrimaryKeyMapping, _ = queries.BindMapping(cmfUserScorerecordType, cmfUserScorerecordMapping, cmfUserScorerecordPrimaryKeyColumns)
	cmfUserScorerecordInsertCacheMut       sync.RWMutex
	cmfUserScorerecordInsertCache          = make(map[string]insertCache)
	cmfUserScorerecordUpdateCacheMut       sync.RWMutex
	cmfUserScorerecordUpdateCache          = make(map[string]updateCache)
	cmfUserScorerecordUpsertCacheMut       sync.RWMutex
	cmfUserScorerecordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfUserScorerecordBeforeInsertHooks []CMFUserScorerecordHook
var cmfUserScorerecordBeforeUpdateHooks []CMFUserScorerecordHook
var cmfUserScorerecordBeforeDeleteHooks []CMFUserScorerecordHook
var cmfUserScorerecordBeforeUpsertHooks []CMFUserScorerecordHook

var cmfUserScorerecordAfterInsertHooks []CMFUserScorerecordHook
var cmfUserScorerecordAfterSelectHooks []CMFUserScorerecordHook
var cmfUserScorerecordAfterUpdateHooks []CMFUserScorerecordHook
var cmfUserScorerecordAfterDeleteHooks []CMFUserScorerecordHook
var cmfUserScorerecordAfterUpsertHooks []CMFUserScorerecordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFUserScorerecord) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFUserScorerecord) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFUserScorerecord) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFUserScorerecord) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFUserScorerecord) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFUserScorerecord) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFUserScorerecord) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFUserScorerecord) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFUserScorerecord) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserScorerecordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFUserScorerecordHook registers your hook function for all future operations.
func AddCMFUserScorerecordHook(hookPoint boil.HookPoint, cmfUserScorerecordHook CMFUserScorerecordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfUserScorerecordBeforeInsertHooks = append(cmfUserScorerecordBeforeInsertHooks, cmfUserScorerecordHook)
	case boil.BeforeUpdateHook:
		cmfUserScorerecordBeforeUpdateHooks = append(cmfUserScorerecordBeforeUpdateHooks, cmfUserScorerecordHook)
	case boil.BeforeDeleteHook:
		cmfUserScorerecordBeforeDeleteHooks = append(cmfUserScorerecordBeforeDeleteHooks, cmfUserScorerecordHook)
	case boil.BeforeUpsertHook:
		cmfUserScorerecordBeforeUpsertHooks = append(cmfUserScorerecordBeforeUpsertHooks, cmfUserScorerecordHook)
	case boil.AfterInsertHook:
		cmfUserScorerecordAfterInsertHooks = append(cmfUserScorerecordAfterInsertHooks, cmfUserScorerecordHook)
	case boil.AfterSelectHook:
		cmfUserScorerecordAfterSelectHooks = append(cmfUserScorerecordAfterSelectHooks, cmfUserScorerecordHook)
	case boil.AfterUpdateHook:
		cmfUserScorerecordAfterUpdateHooks = append(cmfUserScorerecordAfterUpdateHooks, cmfUserScorerecordHook)
	case boil.AfterDeleteHook:
		cmfUserScorerecordAfterDeleteHooks = append(cmfUserScorerecordAfterDeleteHooks, cmfUserScorerecordHook)
	case boil.AfterUpsertHook:
		cmfUserScorerecordAfterUpsertHooks = append(cmfUserScorerecordAfterUpsertHooks, cmfUserScorerecordHook)
	}
}

// One returns a single cmfUserScorerecord record from the query.
func (q cmfUserScorerecordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFUserScorerecord, error) {
	o := &CMFUserScorerecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_user_scorerecord")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFUserScorerecord records from the query.
func (q cmfUserScorerecordQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFUserScorerecordSlice, error) {
	var o []*CMFUserScorerecord

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFUserScorerecord slice")
	}

	if len(cmfUserScorerecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFUserScorerecord records in the query.
func (q cmfUserScorerecordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_user_scorerecord rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfUserScorerecordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_user_scorerecord exists")
	}

	return count > 0, nil
}

// CMFUserScorerecords retrieves all the records using an executor.
func CMFUserScorerecords(mods ...qm.QueryMod) cmfUserScorerecordQuery {
	mods = append(mods, qm.From("`cmf_user_scorerecord`"))
	return cmfUserScorerecordQuery{NewQuery(mods...)}
}

// FindCMFUserScorerecord retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFUserScorerecord(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFUserScorerecord, error) {
	cmfUserScorerecordObj := &CMFUserScorerecord{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_user_scorerecord` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfUserScorerecordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_user_scorerecord")
	}

	return cmfUserScorerecordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFUserScorerecord) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_scorerecord provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserScorerecordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfUserScorerecordInsertCacheMut.RLock()
	cache, cached := cmfUserScorerecordInsertCache[key]
	cmfUserScorerecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfUserScorerecordAllColumns,
			cmfUserScorerecordColumnsWithDefault,
			cmfUserScorerecordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserScorerecordType, cmfUserScorerecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfUserScorerecordType, cmfUserScorerecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_user_scorerecord` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_user_scorerecord` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_user_scorerecord` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfUserScorerecordPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_user_scorerecord")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfUserScorerecordMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_scorerecord")
	}

CacheNoHooks:
	if !cached {
		cmfUserScorerecordInsertCacheMut.Lock()
		cmfUserScorerecordInsertCache[key] = cache
		cmfUserScorerecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFUserScorerecord.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFUserScorerecord) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfUserScorerecordUpdateCacheMut.RLock()
	cache, cached := cmfUserScorerecordUpdateCache[key]
	cmfUserScorerecordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfUserScorerecordAllColumns,
			cmfUserScorerecordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_user_scorerecord, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_user_scorerecord` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfUserScorerecordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfUserScorerecordType, cmfUserScorerecordMapping, append(wl, cmfUserScorerecordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_user_scorerecord row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_user_scorerecord")
	}

	if !cached {
		cmfUserScorerecordUpdateCacheMut.Lock()
		cmfUserScorerecordUpdateCache[key] = cache
		cmfUserScorerecordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfUserScorerecordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_user_scorerecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_user_scorerecord")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFUserScorerecordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserScorerecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_user_scorerecord` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserScorerecordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfUserScorerecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfUserScorerecord")
	}
	return rowsAff, nil
}

var mySQLCMFUserScorerecordUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFUserScorerecord) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_scorerecord provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserScorerecordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFUserScorerecordUniqueColumns, o)

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

	cmfUserScorerecordUpsertCacheMut.RLock()
	cache, cached := cmfUserScorerecordUpsertCache[key]
	cmfUserScorerecordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfUserScorerecordAllColumns,
			cmfUserScorerecordColumnsWithDefault,
			cmfUserScorerecordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfUserScorerecordAllColumns,
			cmfUserScorerecordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_user_scorerecord, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_user_scorerecord`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_user_scorerecord` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserScorerecordType, cmfUserScorerecordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfUserScorerecordType, cmfUserScorerecordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_user_scorerecord")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfUserScorerecordMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfUserScorerecordType, cmfUserScorerecordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_user_scorerecord")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_scorerecord")
	}

CacheNoHooks:
	if !cached {
		cmfUserScorerecordUpsertCacheMut.Lock()
		cmfUserScorerecordUpsertCache[key] = cache
		cmfUserScorerecordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFUserScorerecord record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFUserScorerecord) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFUserScorerecord provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfUserScorerecordPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_user_scorerecord` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_user_scorerecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_user_scorerecord")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfUserScorerecordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfUserScorerecordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_user_scorerecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_scorerecord")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFUserScorerecordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfUserScorerecordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserScorerecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_user_scorerecord` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserScorerecordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfUserScorerecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_scorerecord")
	}

	if len(cmfUserScorerecordAfterDeleteHooks) != 0 {
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
func (o *CMFUserScorerecord) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFUserScorerecord(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFUserScorerecordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFUserScorerecordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserScorerecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_user_scorerecord`.* FROM `cmf_user_scorerecord` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserScorerecordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFUserScorerecordSlice")
	}

	*o = slice

	return nil
}

// CMFUserScorerecordExists checks if the CMFUserScorerecord row exists.
func CMFUserScorerecordExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_user_scorerecord` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_user_scorerecord exists")
	}

	return exists, nil
}
