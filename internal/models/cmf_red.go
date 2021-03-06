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

// CMFRed is an object representing the database table.
type CMFRed struct {
	ID         uint   `boil:"id" json:"id" toml:"id" yaml:"id"`
	Showid     int    `boil:"showid" json:"showid" toml:"showid" yaml:"showid"`
	UID        int    `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Liveuid    int    `boil:"liveuid" json:"liveuid" toml:"liveuid" yaml:"liveuid"`
	Type       bool   `boil:"type" json:"type" toml:"type" yaml:"type"`
	TypeGrant  bool   `boil:"type_grant" json:"type_grant" toml:"type_grant" yaml:"type_grant"`
	Coin       int    `boil:"coin" json:"coin" toml:"coin" yaml:"coin"`
	Nums       int    `boil:"nums" json:"nums" toml:"nums" yaml:"nums"`
	Des        string `boil:"des" json:"des" toml:"des" yaml:"des"`
	Effecttime int    `boil:"effecttime" json:"effecttime" toml:"effecttime" yaml:"effecttime"`
	Addtime    int    `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Status     int8   `boil:"status" json:"status" toml:"status" yaml:"status"`
	CoinRob    int    `boil:"coin_rob" json:"coin_rob" toml:"coin_rob" yaml:"coin_rob"`
	NumsRob    int    `boil:"nums_rob" json:"nums_rob" toml:"nums_rob" yaml:"nums_rob"`

	R *cmfRedR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfRedL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFRedColumns = struct {
	ID         string
	Showid     string
	UID        string
	Liveuid    string
	Type       string
	TypeGrant  string
	Coin       string
	Nums       string
	Des        string
	Effecttime string
	Addtime    string
	Status     string
	CoinRob    string
	NumsRob    string
}{
	ID:         "id",
	Showid:     "showid",
	UID:        "uid",
	Liveuid:    "liveuid",
	Type:       "type",
	TypeGrant:  "type_grant",
	Coin:       "coin",
	Nums:       "nums",
	Des:        "des",
	Effecttime: "effecttime",
	Addtime:    "addtime",
	Status:     "status",
	CoinRob:    "coin_rob",
	NumsRob:    "nums_rob",
}

// Generated where

var CMFRedWhere = struct {
	ID         whereHelperuint
	Showid     whereHelperint
	UID        whereHelperint
	Liveuid    whereHelperint
	Type       whereHelperbool
	TypeGrant  whereHelperbool
	Coin       whereHelperint
	Nums       whereHelperint
	Des        whereHelperstring
	Effecttime whereHelperint
	Addtime    whereHelperint
	Status     whereHelperint8
	CoinRob    whereHelperint
	NumsRob    whereHelperint
}{
	ID:         whereHelperuint{field: "`cmf_red`.`id`"},
	Showid:     whereHelperint{field: "`cmf_red`.`showid`"},
	UID:        whereHelperint{field: "`cmf_red`.`uid`"},
	Liveuid:    whereHelperint{field: "`cmf_red`.`liveuid`"},
	Type:       whereHelperbool{field: "`cmf_red`.`type`"},
	TypeGrant:  whereHelperbool{field: "`cmf_red`.`type_grant`"},
	Coin:       whereHelperint{field: "`cmf_red`.`coin`"},
	Nums:       whereHelperint{field: "`cmf_red`.`nums`"},
	Des:        whereHelperstring{field: "`cmf_red`.`des`"},
	Effecttime: whereHelperint{field: "`cmf_red`.`effecttime`"},
	Addtime:    whereHelperint{field: "`cmf_red`.`addtime`"},
	Status:     whereHelperint8{field: "`cmf_red`.`status`"},
	CoinRob:    whereHelperint{field: "`cmf_red`.`coin_rob`"},
	NumsRob:    whereHelperint{field: "`cmf_red`.`nums_rob`"},
}

// CMFRedRels is where relationship names are stored.
var CMFRedRels = struct {
}{}

// cmfRedR is where relationships are stored.
type cmfRedR struct {
}

// NewStruct creates a new relationship struct
func (*cmfRedR) NewStruct() *cmfRedR {
	return &cmfRedR{}
}

// cmfRedL is where Load methods for each relationship are stored.
type cmfRedL struct{}

var (
	cmfRedAllColumns            = []string{"id", "showid", "uid", "liveuid", "type", "type_grant", "coin", "nums", "des", "effecttime", "addtime", "status", "coin_rob", "nums_rob"}
	cmfRedColumnsWithoutDefault = []string{"des"}
	cmfRedColumnsWithDefault    = []string{"id", "showid", "uid", "liveuid", "type", "type_grant", "coin", "nums", "effecttime", "addtime", "status", "coin_rob", "nums_rob"}
	cmfRedPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFRedSlice is an alias for a slice of pointers to CMFRed.
	// This should generally be used opposed to []CMFRed.
	CMFRedSlice []*CMFRed
	// CMFRedHook is the signature for custom CMFRed hook methods
	CMFRedHook func(context.Context, boil.ContextExecutor, *CMFRed) error

	cmfRedQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfRedType                 = reflect.TypeOf(&CMFRed{})
	cmfRedMapping              = queries.MakeStructMapping(cmfRedType)
	cmfRedPrimaryKeyMapping, _ = queries.BindMapping(cmfRedType, cmfRedMapping, cmfRedPrimaryKeyColumns)
	cmfRedInsertCacheMut       sync.RWMutex
	cmfRedInsertCache          = make(map[string]insertCache)
	cmfRedUpdateCacheMut       sync.RWMutex
	cmfRedUpdateCache          = make(map[string]updateCache)
	cmfRedUpsertCacheMut       sync.RWMutex
	cmfRedUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfRedBeforeInsertHooks []CMFRedHook
var cmfRedBeforeUpdateHooks []CMFRedHook
var cmfRedBeforeDeleteHooks []CMFRedHook
var cmfRedBeforeUpsertHooks []CMFRedHook

var cmfRedAfterInsertHooks []CMFRedHook
var cmfRedAfterSelectHooks []CMFRedHook
var cmfRedAfterUpdateHooks []CMFRedHook
var cmfRedAfterDeleteHooks []CMFRedHook
var cmfRedAfterUpsertHooks []CMFRedHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFRed) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFRed) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFRed) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFRed) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFRed) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFRed) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFRed) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFRed) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFRed) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRedAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFRedHook registers your hook function for all future operations.
func AddCMFRedHook(hookPoint boil.HookPoint, cmfRedHook CMFRedHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfRedBeforeInsertHooks = append(cmfRedBeforeInsertHooks, cmfRedHook)
	case boil.BeforeUpdateHook:
		cmfRedBeforeUpdateHooks = append(cmfRedBeforeUpdateHooks, cmfRedHook)
	case boil.BeforeDeleteHook:
		cmfRedBeforeDeleteHooks = append(cmfRedBeforeDeleteHooks, cmfRedHook)
	case boil.BeforeUpsertHook:
		cmfRedBeforeUpsertHooks = append(cmfRedBeforeUpsertHooks, cmfRedHook)
	case boil.AfterInsertHook:
		cmfRedAfterInsertHooks = append(cmfRedAfterInsertHooks, cmfRedHook)
	case boil.AfterSelectHook:
		cmfRedAfterSelectHooks = append(cmfRedAfterSelectHooks, cmfRedHook)
	case boil.AfterUpdateHook:
		cmfRedAfterUpdateHooks = append(cmfRedAfterUpdateHooks, cmfRedHook)
	case boil.AfterDeleteHook:
		cmfRedAfterDeleteHooks = append(cmfRedAfterDeleteHooks, cmfRedHook)
	case boil.AfterUpsertHook:
		cmfRedAfterUpsertHooks = append(cmfRedAfterUpsertHooks, cmfRedHook)
	}
}

// One returns a single cmfRed record from the query.
func (q cmfRedQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFRed, error) {
	o := &CMFRed{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_red")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFRed records from the query.
func (q cmfRedQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFRedSlice, error) {
	var o []*CMFRed

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFRed slice")
	}

	if len(cmfRedAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFRed records in the query.
func (q cmfRedQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_red rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfRedQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_red exists")
	}

	return count > 0, nil
}

// CMFReds retrieves all the records using an executor.
func CMFReds(mods ...qm.QueryMod) cmfRedQuery {
	mods = append(mods, qm.From("`cmf_red`"))
	return cmfRedQuery{NewQuery(mods...)}
}

// FindCMFRed retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFRed(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CMFRed, error) {
	cmfRedObj := &CMFRed{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_red` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfRedObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_red")
	}

	return cmfRedObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFRed) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_red provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfRedColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfRedInsertCacheMut.RLock()
	cache, cached := cmfRedInsertCache[key]
	cmfRedInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfRedAllColumns,
			cmfRedColumnsWithDefault,
			cmfRedColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfRedType, cmfRedMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfRedType, cmfRedMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_red` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_red` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_red` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfRedPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_red")
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

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfRedMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_red")
	}

CacheNoHooks:
	if !cached {
		cmfRedInsertCacheMut.Lock()
		cmfRedInsertCache[key] = cache
		cmfRedInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFRed.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFRed) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfRedUpdateCacheMut.RLock()
	cache, cached := cmfRedUpdateCache[key]
	cmfRedUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfRedAllColumns,
			cmfRedPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_red, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_red` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfRedPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfRedType, cmfRedMapping, append(wl, cmfRedPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_red row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_red")
	}

	if !cached {
		cmfRedUpdateCacheMut.Lock()
		cmfRedUpdateCache[key] = cache
		cmfRedUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfRedQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_red")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_red")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFRedSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfRedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_red` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfRedPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfRed slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfRed")
	}
	return rowsAff, nil
}

var mySQLCMFRedUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFRed) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_red provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfRedColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFRedUniqueColumns, o)

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

	cmfRedUpsertCacheMut.RLock()
	cache, cached := cmfRedUpsertCache[key]
	cmfRedUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfRedAllColumns,
			cmfRedColumnsWithDefault,
			cmfRedColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfRedAllColumns,
			cmfRedPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_red, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_red`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_red` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfRedType, cmfRedMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfRedType, cmfRedMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_red")
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

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfRedMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfRedType, cmfRedMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_red")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_red")
	}

CacheNoHooks:
	if !cached {
		cmfRedUpsertCacheMut.Lock()
		cmfRedUpsertCache[key] = cache
		cmfRedUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFRed record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFRed) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFRed provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfRedPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_red` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_red")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_red")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfRedQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfRedQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_red")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_red")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFRedSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfRedBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfRedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_red` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfRedPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfRed slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_red")
	}

	if len(cmfRedAfterDeleteHooks) != 0 {
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
func (o *CMFRed) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFRed(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFRedSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFRedSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfRedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_red`.* FROM `cmf_red` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfRedPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFRedSlice")
	}

	*o = slice

	return nil
}

// CMFRedExists checks if the CMFRed row exists.
func CMFRedExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_red` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_red exists")
	}

	return exists, nil
}
