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

// CMFUserCoinrecord is an object representing the database table.
type CMFUserCoinrecord struct {
	ID        int  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Type      bool `boil:"type" json:"type" toml:"type" yaml:"type"`
	Action    bool `boil:"action" json:"action" toml:"action" yaml:"action"`
	UID       int  `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Touid     int  `boil:"touid" json:"touid" toml:"touid" yaml:"touid"`
	Giftid    int  `boil:"giftid" json:"giftid" toml:"giftid" yaml:"giftid"`
	Giftcount int  `boil:"giftcount" json:"giftcount" toml:"giftcount" yaml:"giftcount"`
	Totalcoin int  `boil:"totalcoin" json:"totalcoin" toml:"totalcoin" yaml:"totalcoin"`
	Showid    int  `boil:"showid" json:"showid" toml:"showid" yaml:"showid"`
	Addtime   int  `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Mark      bool `boil:"mark" json:"mark" toml:"mark" yaml:"mark"`

	R *cmfUserCoinrecordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfUserCoinrecordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFUserCoinrecordColumns = struct {
	ID        string
	Type      string
	Action    string
	UID       string
	Touid     string
	Giftid    string
	Giftcount string
	Totalcoin string
	Showid    string
	Addtime   string
	Mark      string
}{
	ID:        "id",
	Type:      "type",
	Action:    "action",
	UID:       "uid",
	Touid:     "touid",
	Giftid:    "giftid",
	Giftcount: "giftcount",
	Totalcoin: "totalcoin",
	Showid:    "showid",
	Addtime:   "addtime",
	Mark:      "mark",
}

// Generated where

var CMFUserCoinrecordWhere = struct {
	ID        whereHelperint
	Type      whereHelperbool
	Action    whereHelperbool
	UID       whereHelperint
	Touid     whereHelperint
	Giftid    whereHelperint
	Giftcount whereHelperint
	Totalcoin whereHelperint
	Showid    whereHelperint
	Addtime   whereHelperint
	Mark      whereHelperbool
}{
	ID:        whereHelperint{field: "`cmf_user_coinrecord`.`id`"},
	Type:      whereHelperbool{field: "`cmf_user_coinrecord`.`type`"},
	Action:    whereHelperbool{field: "`cmf_user_coinrecord`.`action`"},
	UID:       whereHelperint{field: "`cmf_user_coinrecord`.`uid`"},
	Touid:     whereHelperint{field: "`cmf_user_coinrecord`.`touid`"},
	Giftid:    whereHelperint{field: "`cmf_user_coinrecord`.`giftid`"},
	Giftcount: whereHelperint{field: "`cmf_user_coinrecord`.`giftcount`"},
	Totalcoin: whereHelperint{field: "`cmf_user_coinrecord`.`totalcoin`"},
	Showid:    whereHelperint{field: "`cmf_user_coinrecord`.`showid`"},
	Addtime:   whereHelperint{field: "`cmf_user_coinrecord`.`addtime`"},
	Mark:      whereHelperbool{field: "`cmf_user_coinrecord`.`mark`"},
}

// CMFUserCoinrecordRels is where relationship names are stored.
var CMFUserCoinrecordRels = struct {
}{}

// cmfUserCoinrecordR is where relationships are stored.
type cmfUserCoinrecordR struct {
}

// NewStruct creates a new relationship struct
func (*cmfUserCoinrecordR) NewStruct() *cmfUserCoinrecordR {
	return &cmfUserCoinrecordR{}
}

// cmfUserCoinrecordL is where Load methods for each relationship are stored.
type cmfUserCoinrecordL struct{}

var (
	cmfUserCoinrecordAllColumns            = []string{"id", "type", "action", "uid", "touid", "giftid", "giftcount", "totalcoin", "showid", "addtime", "mark"}
	cmfUserCoinrecordColumnsWithoutDefault = []string{}
	cmfUserCoinrecordColumnsWithDefault    = []string{"id", "type", "action", "uid", "touid", "giftid", "giftcount", "totalcoin", "showid", "addtime", "mark"}
	cmfUserCoinrecordPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFUserCoinrecordSlice is an alias for a slice of pointers to CMFUserCoinrecord.
	// This should generally be used opposed to []CMFUserCoinrecord.
	CMFUserCoinrecordSlice []*CMFUserCoinrecord
	// CMFUserCoinrecordHook is the signature for custom CMFUserCoinrecord hook methods
	CMFUserCoinrecordHook func(context.Context, boil.ContextExecutor, *CMFUserCoinrecord) error

	cmfUserCoinrecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfUserCoinrecordType                 = reflect.TypeOf(&CMFUserCoinrecord{})
	cmfUserCoinrecordMapping              = queries.MakeStructMapping(cmfUserCoinrecordType)
	cmfUserCoinrecordPrimaryKeyMapping, _ = queries.BindMapping(cmfUserCoinrecordType, cmfUserCoinrecordMapping, cmfUserCoinrecordPrimaryKeyColumns)
	cmfUserCoinrecordInsertCacheMut       sync.RWMutex
	cmfUserCoinrecordInsertCache          = make(map[string]insertCache)
	cmfUserCoinrecordUpdateCacheMut       sync.RWMutex
	cmfUserCoinrecordUpdateCache          = make(map[string]updateCache)
	cmfUserCoinrecordUpsertCacheMut       sync.RWMutex
	cmfUserCoinrecordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfUserCoinrecordBeforeInsertHooks []CMFUserCoinrecordHook
var cmfUserCoinrecordBeforeUpdateHooks []CMFUserCoinrecordHook
var cmfUserCoinrecordBeforeDeleteHooks []CMFUserCoinrecordHook
var cmfUserCoinrecordBeforeUpsertHooks []CMFUserCoinrecordHook

var cmfUserCoinrecordAfterInsertHooks []CMFUserCoinrecordHook
var cmfUserCoinrecordAfterSelectHooks []CMFUserCoinrecordHook
var cmfUserCoinrecordAfterUpdateHooks []CMFUserCoinrecordHook
var cmfUserCoinrecordAfterDeleteHooks []CMFUserCoinrecordHook
var cmfUserCoinrecordAfterUpsertHooks []CMFUserCoinrecordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFUserCoinrecord) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFUserCoinrecord) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFUserCoinrecord) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFUserCoinrecord) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFUserCoinrecord) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFUserCoinrecord) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFUserCoinrecord) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFUserCoinrecord) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFUserCoinrecord) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserCoinrecordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFUserCoinrecordHook registers your hook function for all future operations.
func AddCMFUserCoinrecordHook(hookPoint boil.HookPoint, cmfUserCoinrecordHook CMFUserCoinrecordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfUserCoinrecordBeforeInsertHooks = append(cmfUserCoinrecordBeforeInsertHooks, cmfUserCoinrecordHook)
	case boil.BeforeUpdateHook:
		cmfUserCoinrecordBeforeUpdateHooks = append(cmfUserCoinrecordBeforeUpdateHooks, cmfUserCoinrecordHook)
	case boil.BeforeDeleteHook:
		cmfUserCoinrecordBeforeDeleteHooks = append(cmfUserCoinrecordBeforeDeleteHooks, cmfUserCoinrecordHook)
	case boil.BeforeUpsertHook:
		cmfUserCoinrecordBeforeUpsertHooks = append(cmfUserCoinrecordBeforeUpsertHooks, cmfUserCoinrecordHook)
	case boil.AfterInsertHook:
		cmfUserCoinrecordAfterInsertHooks = append(cmfUserCoinrecordAfterInsertHooks, cmfUserCoinrecordHook)
	case boil.AfterSelectHook:
		cmfUserCoinrecordAfterSelectHooks = append(cmfUserCoinrecordAfterSelectHooks, cmfUserCoinrecordHook)
	case boil.AfterUpdateHook:
		cmfUserCoinrecordAfterUpdateHooks = append(cmfUserCoinrecordAfterUpdateHooks, cmfUserCoinrecordHook)
	case boil.AfterDeleteHook:
		cmfUserCoinrecordAfterDeleteHooks = append(cmfUserCoinrecordAfterDeleteHooks, cmfUserCoinrecordHook)
	case boil.AfterUpsertHook:
		cmfUserCoinrecordAfterUpsertHooks = append(cmfUserCoinrecordAfterUpsertHooks, cmfUserCoinrecordHook)
	}
}

// One returns a single cmfUserCoinrecord record from the query.
func (q cmfUserCoinrecordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFUserCoinrecord, error) {
	o := &CMFUserCoinrecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_user_coinrecord")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFUserCoinrecord records from the query.
func (q cmfUserCoinrecordQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFUserCoinrecordSlice, error) {
	var o []*CMFUserCoinrecord

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFUserCoinrecord slice")
	}

	if len(cmfUserCoinrecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFUserCoinrecord records in the query.
func (q cmfUserCoinrecordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_user_coinrecord rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfUserCoinrecordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_user_coinrecord exists")
	}

	return count > 0, nil
}

// CMFUserCoinrecords retrieves all the records using an executor.
func CMFUserCoinrecords(mods ...qm.QueryMod) cmfUserCoinrecordQuery {
	mods = append(mods, qm.From("`cmf_user_coinrecord`"))
	return cmfUserCoinrecordQuery{NewQuery(mods...)}
}

// FindCMFUserCoinrecord retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFUserCoinrecord(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFUserCoinrecord, error) {
	cmfUserCoinrecordObj := &CMFUserCoinrecord{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_user_coinrecord` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfUserCoinrecordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_user_coinrecord")
	}

	return cmfUserCoinrecordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFUserCoinrecord) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_coinrecord provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserCoinrecordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfUserCoinrecordInsertCacheMut.RLock()
	cache, cached := cmfUserCoinrecordInsertCache[key]
	cmfUserCoinrecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfUserCoinrecordAllColumns,
			cmfUserCoinrecordColumnsWithDefault,
			cmfUserCoinrecordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserCoinrecordType, cmfUserCoinrecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfUserCoinrecordType, cmfUserCoinrecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_user_coinrecord` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_user_coinrecord` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_user_coinrecord` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfUserCoinrecordPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_user_coinrecord")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfUserCoinrecordMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_coinrecord")
	}

CacheNoHooks:
	if !cached {
		cmfUserCoinrecordInsertCacheMut.Lock()
		cmfUserCoinrecordInsertCache[key] = cache
		cmfUserCoinrecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFUserCoinrecord.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFUserCoinrecord) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfUserCoinrecordUpdateCacheMut.RLock()
	cache, cached := cmfUserCoinrecordUpdateCache[key]
	cmfUserCoinrecordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfUserCoinrecordAllColumns,
			cmfUserCoinrecordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_user_coinrecord, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_user_coinrecord` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfUserCoinrecordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfUserCoinrecordType, cmfUserCoinrecordMapping, append(wl, cmfUserCoinrecordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_user_coinrecord row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_user_coinrecord")
	}

	if !cached {
		cmfUserCoinrecordUpdateCacheMut.Lock()
		cmfUserCoinrecordUpdateCache[key] = cache
		cmfUserCoinrecordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfUserCoinrecordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_user_coinrecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_user_coinrecord")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFUserCoinrecordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserCoinrecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_user_coinrecord` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserCoinrecordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfUserCoinrecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfUserCoinrecord")
	}
	return rowsAff, nil
}

var mySQLCMFUserCoinrecordUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFUserCoinrecord) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_coinrecord provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserCoinrecordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFUserCoinrecordUniqueColumns, o)

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

	cmfUserCoinrecordUpsertCacheMut.RLock()
	cache, cached := cmfUserCoinrecordUpsertCache[key]
	cmfUserCoinrecordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfUserCoinrecordAllColumns,
			cmfUserCoinrecordColumnsWithDefault,
			cmfUserCoinrecordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfUserCoinrecordAllColumns,
			cmfUserCoinrecordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_user_coinrecord, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_user_coinrecord`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_user_coinrecord` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserCoinrecordType, cmfUserCoinrecordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfUserCoinrecordType, cmfUserCoinrecordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_user_coinrecord")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfUserCoinrecordMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfUserCoinrecordType, cmfUserCoinrecordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_user_coinrecord")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_coinrecord")
	}

CacheNoHooks:
	if !cached {
		cmfUserCoinrecordUpsertCacheMut.Lock()
		cmfUserCoinrecordUpsertCache[key] = cache
		cmfUserCoinrecordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFUserCoinrecord record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFUserCoinrecord) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFUserCoinrecord provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfUserCoinrecordPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_user_coinrecord` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_user_coinrecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_user_coinrecord")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfUserCoinrecordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfUserCoinrecordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_user_coinrecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_coinrecord")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFUserCoinrecordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfUserCoinrecordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserCoinrecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_user_coinrecord` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserCoinrecordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfUserCoinrecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_coinrecord")
	}

	if len(cmfUserCoinrecordAfterDeleteHooks) != 0 {
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
func (o *CMFUserCoinrecord) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFUserCoinrecord(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFUserCoinrecordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFUserCoinrecordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserCoinrecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_user_coinrecord`.* FROM `cmf_user_coinrecord` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserCoinrecordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFUserCoinrecordSlice")
	}

	*o = slice

	return nil
}

// CMFUserCoinrecordExists checks if the CMFUserCoinrecord row exists.
func CMFUserCoinrecordExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_user_coinrecord` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_user_coinrecord exists")
	}

	return exists, nil
}
