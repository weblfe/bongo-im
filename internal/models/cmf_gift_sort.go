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

// CMFGiftSort is an object representing the database table.
type CMFGiftSort struct {
	ID       int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Sortname string `boil:"sortname" json:"sortname" toml:"sortname" yaml:"sortname"`
	Orderno  int    `boil:"orderno" json:"orderno" toml:"orderno" yaml:"orderno"`
	Addtime  int    `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`

	R *cmfGiftSortR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfGiftSortL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFGiftSortColumns = struct {
	ID       string
	Sortname string
	Orderno  string
	Addtime  string
}{
	ID:       "id",
	Sortname: "sortname",
	Orderno:  "orderno",
	Addtime:  "addtime",
}

// Generated where

var CMFGiftSortWhere = struct {
	ID       whereHelperint
	Sortname whereHelperstring
	Orderno  whereHelperint
	Addtime  whereHelperint
}{
	ID:       whereHelperint{field: "`cmf_gift_sort`.`id`"},
	Sortname: whereHelperstring{field: "`cmf_gift_sort`.`sortname`"},
	Orderno:  whereHelperint{field: "`cmf_gift_sort`.`orderno`"},
	Addtime:  whereHelperint{field: "`cmf_gift_sort`.`addtime`"},
}

// CMFGiftSortRels is where relationship names are stored.
var CMFGiftSortRels = struct {
}{}

// cmfGiftSortR is where relationships are stored.
type cmfGiftSortR struct {
}

// NewStruct creates a new relationship struct
func (*cmfGiftSortR) NewStruct() *cmfGiftSortR {
	return &cmfGiftSortR{}
}

// cmfGiftSortL is where Load methods for each relationship are stored.
type cmfGiftSortL struct{}

var (
	cmfGiftSortAllColumns            = []string{"id", "sortname", "orderno", "addtime"}
	cmfGiftSortColumnsWithoutDefault = []string{"sortname"}
	cmfGiftSortColumnsWithDefault    = []string{"id", "orderno", "addtime"}
	cmfGiftSortPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFGiftSortSlice is an alias for a slice of pointers to CMFGiftSort.
	// This should generally be used opposed to []CMFGiftSort.
	CMFGiftSortSlice []*CMFGiftSort
	// CMFGiftSortHook is the signature for custom CMFGiftSort hook methods
	CMFGiftSortHook func(context.Context, boil.ContextExecutor, *CMFGiftSort) error

	cmfGiftSortQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfGiftSortType                 = reflect.TypeOf(&CMFGiftSort{})
	cmfGiftSortMapping              = queries.MakeStructMapping(cmfGiftSortType)
	cmfGiftSortPrimaryKeyMapping, _ = queries.BindMapping(cmfGiftSortType, cmfGiftSortMapping, cmfGiftSortPrimaryKeyColumns)
	cmfGiftSortInsertCacheMut       sync.RWMutex
	cmfGiftSortInsertCache          = make(map[string]insertCache)
	cmfGiftSortUpdateCacheMut       sync.RWMutex
	cmfGiftSortUpdateCache          = make(map[string]updateCache)
	cmfGiftSortUpsertCacheMut       sync.RWMutex
	cmfGiftSortUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfGiftSortBeforeInsertHooks []CMFGiftSortHook
var cmfGiftSortBeforeUpdateHooks []CMFGiftSortHook
var cmfGiftSortBeforeDeleteHooks []CMFGiftSortHook
var cmfGiftSortBeforeUpsertHooks []CMFGiftSortHook

var cmfGiftSortAfterInsertHooks []CMFGiftSortHook
var cmfGiftSortAfterSelectHooks []CMFGiftSortHook
var cmfGiftSortAfterUpdateHooks []CMFGiftSortHook
var cmfGiftSortAfterDeleteHooks []CMFGiftSortHook
var cmfGiftSortAfterUpsertHooks []CMFGiftSortHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFGiftSort) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFGiftSort) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFGiftSort) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFGiftSort) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFGiftSort) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFGiftSort) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFGiftSort) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFGiftSort) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFGiftSort) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGiftSortAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFGiftSortHook registers your hook function for all future operations.
func AddCMFGiftSortHook(hookPoint boil.HookPoint, cmfGiftSortHook CMFGiftSortHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfGiftSortBeforeInsertHooks = append(cmfGiftSortBeforeInsertHooks, cmfGiftSortHook)
	case boil.BeforeUpdateHook:
		cmfGiftSortBeforeUpdateHooks = append(cmfGiftSortBeforeUpdateHooks, cmfGiftSortHook)
	case boil.BeforeDeleteHook:
		cmfGiftSortBeforeDeleteHooks = append(cmfGiftSortBeforeDeleteHooks, cmfGiftSortHook)
	case boil.BeforeUpsertHook:
		cmfGiftSortBeforeUpsertHooks = append(cmfGiftSortBeforeUpsertHooks, cmfGiftSortHook)
	case boil.AfterInsertHook:
		cmfGiftSortAfterInsertHooks = append(cmfGiftSortAfterInsertHooks, cmfGiftSortHook)
	case boil.AfterSelectHook:
		cmfGiftSortAfterSelectHooks = append(cmfGiftSortAfterSelectHooks, cmfGiftSortHook)
	case boil.AfterUpdateHook:
		cmfGiftSortAfterUpdateHooks = append(cmfGiftSortAfterUpdateHooks, cmfGiftSortHook)
	case boil.AfterDeleteHook:
		cmfGiftSortAfterDeleteHooks = append(cmfGiftSortAfterDeleteHooks, cmfGiftSortHook)
	case boil.AfterUpsertHook:
		cmfGiftSortAfterUpsertHooks = append(cmfGiftSortAfterUpsertHooks, cmfGiftSortHook)
	}
}

// One returns a single cmfGiftSort record from the query.
func (q cmfGiftSortQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFGiftSort, error) {
	o := &CMFGiftSort{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_gift_sort")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFGiftSort records from the query.
func (q cmfGiftSortQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFGiftSortSlice, error) {
	var o []*CMFGiftSort

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFGiftSort slice")
	}

	if len(cmfGiftSortAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFGiftSort records in the query.
func (q cmfGiftSortQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_gift_sort rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfGiftSortQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_gift_sort exists")
	}

	return count > 0, nil
}

// CMFGiftSorts retrieves all the records using an executor.
func CMFGiftSorts(mods ...qm.QueryMod) cmfGiftSortQuery {
	mods = append(mods, qm.From("`cmf_gift_sort`"))
	return cmfGiftSortQuery{NewQuery(mods...)}
}

// FindCMFGiftSort retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFGiftSort(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFGiftSort, error) {
	cmfGiftSortObj := &CMFGiftSort{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_gift_sort` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfGiftSortObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_gift_sort")
	}

	return cmfGiftSortObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFGiftSort) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_gift_sort provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfGiftSortColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfGiftSortInsertCacheMut.RLock()
	cache, cached := cmfGiftSortInsertCache[key]
	cmfGiftSortInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfGiftSortAllColumns,
			cmfGiftSortColumnsWithDefault,
			cmfGiftSortColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfGiftSortType, cmfGiftSortMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfGiftSortType, cmfGiftSortMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_gift_sort` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_gift_sort` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_gift_sort` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfGiftSortPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_gift_sort")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfGiftSortMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_gift_sort")
	}

CacheNoHooks:
	if !cached {
		cmfGiftSortInsertCacheMut.Lock()
		cmfGiftSortInsertCache[key] = cache
		cmfGiftSortInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFGiftSort.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFGiftSort) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfGiftSortUpdateCacheMut.RLock()
	cache, cached := cmfGiftSortUpdateCache[key]
	cmfGiftSortUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfGiftSortAllColumns,
			cmfGiftSortPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_gift_sort, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_gift_sort` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfGiftSortPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfGiftSortType, cmfGiftSortMapping, append(wl, cmfGiftSortPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_gift_sort row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_gift_sort")
	}

	if !cached {
		cmfGiftSortUpdateCacheMut.Lock()
		cmfGiftSortUpdateCache[key] = cache
		cmfGiftSortUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfGiftSortQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_gift_sort")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_gift_sort")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFGiftSortSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGiftSortPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_gift_sort` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGiftSortPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfGiftSort slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfGiftSort")
	}
	return rowsAff, nil
}

var mySQLCMFGiftSortUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFGiftSort) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_gift_sort provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfGiftSortColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFGiftSortUniqueColumns, o)

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

	cmfGiftSortUpsertCacheMut.RLock()
	cache, cached := cmfGiftSortUpsertCache[key]
	cmfGiftSortUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfGiftSortAllColumns,
			cmfGiftSortColumnsWithDefault,
			cmfGiftSortColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfGiftSortAllColumns,
			cmfGiftSortPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_gift_sort, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_gift_sort`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_gift_sort` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfGiftSortType, cmfGiftSortMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfGiftSortType, cmfGiftSortMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_gift_sort")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfGiftSortMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfGiftSortType, cmfGiftSortMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_gift_sort")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_gift_sort")
	}

CacheNoHooks:
	if !cached {
		cmfGiftSortUpsertCacheMut.Lock()
		cmfGiftSortUpsertCache[key] = cache
		cmfGiftSortUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFGiftSort record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFGiftSort) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFGiftSort provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfGiftSortPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_gift_sort` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_gift_sort")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_gift_sort")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfGiftSortQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfGiftSortQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_gift_sort")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_gift_sort")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFGiftSortSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfGiftSortBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGiftSortPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_gift_sort` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGiftSortPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfGiftSort slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_gift_sort")
	}

	if len(cmfGiftSortAfterDeleteHooks) != 0 {
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
func (o *CMFGiftSort) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFGiftSort(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFGiftSortSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFGiftSortSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGiftSortPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_gift_sort`.* FROM `cmf_gift_sort` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGiftSortPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFGiftSortSlice")
	}

	*o = slice

	return nil
}

// CMFGiftSortExists checks if the CMFGiftSort row exists.
func CMFGiftSortExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_gift_sort` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_gift_sort exists")
	}

	return exists, nil
}
