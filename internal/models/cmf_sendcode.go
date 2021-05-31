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

// CMFSendcode is an object representing the database table.
type CMFSendcode struct {
	ID      int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Type    bool   `boil:"type" json:"type" toml:"type" yaml:"type"`
	Account string `boil:"account" json:"account" toml:"account" yaml:"account"`
	Content string `boil:"content" json:"content" toml:"content" yaml:"content"`
	Addtime int    `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`

	R *cmfSendcodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfSendcodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFSendcodeColumns = struct {
	ID      string
	Type    string
	Account string
	Content string
	Addtime string
}{
	ID:      "id",
	Type:    "type",
	Account: "account",
	Content: "content",
	Addtime: "addtime",
}

// Generated where

var CMFSendcodeWhere = struct {
	ID      whereHelperint
	Type    whereHelperbool
	Account whereHelperstring
	Content whereHelperstring
	Addtime whereHelperint
}{
	ID:      whereHelperint{field: "`cmf_sendcode`.`id`"},
	Type:    whereHelperbool{field: "`cmf_sendcode`.`type`"},
	Account: whereHelperstring{field: "`cmf_sendcode`.`account`"},
	Content: whereHelperstring{field: "`cmf_sendcode`.`content`"},
	Addtime: whereHelperint{field: "`cmf_sendcode`.`addtime`"},
}

// CMFSendcodeRels is where relationship names are stored.
var CMFSendcodeRels = struct {
}{}

// cmfSendcodeR is where relationships are stored.
type cmfSendcodeR struct {
}

// NewStruct creates a new relationship struct
func (*cmfSendcodeR) NewStruct() *cmfSendcodeR {
	return &cmfSendcodeR{}
}

// cmfSendcodeL is where Load methods for each relationship are stored.
type cmfSendcodeL struct{}

var (
	cmfSendcodeAllColumns            = []string{"id", "type", "account", "content", "addtime"}
	cmfSendcodeColumnsWithoutDefault = []string{"account", "content", "addtime"}
	cmfSendcodeColumnsWithDefault    = []string{"id", "type"}
	cmfSendcodePrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFSendcodeSlice is an alias for a slice of pointers to CMFSendcode.
	// This should generally be used opposed to []CMFSendcode.
	CMFSendcodeSlice []*CMFSendcode
	// CMFSendcodeHook is the signature for custom CMFSendcode hook methods
	CMFSendcodeHook func(context.Context, boil.ContextExecutor, *CMFSendcode) error

	cmfSendcodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfSendcodeType                 = reflect.TypeOf(&CMFSendcode{})
	cmfSendcodeMapping              = queries.MakeStructMapping(cmfSendcodeType)
	cmfSendcodePrimaryKeyMapping, _ = queries.BindMapping(cmfSendcodeType, cmfSendcodeMapping, cmfSendcodePrimaryKeyColumns)
	cmfSendcodeInsertCacheMut       sync.RWMutex
	cmfSendcodeInsertCache          = make(map[string]insertCache)
	cmfSendcodeUpdateCacheMut       sync.RWMutex
	cmfSendcodeUpdateCache          = make(map[string]updateCache)
	cmfSendcodeUpsertCacheMut       sync.RWMutex
	cmfSendcodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfSendcodeBeforeInsertHooks []CMFSendcodeHook
var cmfSendcodeBeforeUpdateHooks []CMFSendcodeHook
var cmfSendcodeBeforeDeleteHooks []CMFSendcodeHook
var cmfSendcodeBeforeUpsertHooks []CMFSendcodeHook

var cmfSendcodeAfterInsertHooks []CMFSendcodeHook
var cmfSendcodeAfterSelectHooks []CMFSendcodeHook
var cmfSendcodeAfterUpdateHooks []CMFSendcodeHook
var cmfSendcodeAfterDeleteHooks []CMFSendcodeHook
var cmfSendcodeAfterUpsertHooks []CMFSendcodeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFSendcode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFSendcode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFSendcode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFSendcode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFSendcode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFSendcode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFSendcode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFSendcode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFSendcode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSendcodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFSendcodeHook registers your hook function for all future operations.
func AddCMFSendcodeHook(hookPoint boil.HookPoint, cmfSendcodeHook CMFSendcodeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfSendcodeBeforeInsertHooks = append(cmfSendcodeBeforeInsertHooks, cmfSendcodeHook)
	case boil.BeforeUpdateHook:
		cmfSendcodeBeforeUpdateHooks = append(cmfSendcodeBeforeUpdateHooks, cmfSendcodeHook)
	case boil.BeforeDeleteHook:
		cmfSendcodeBeforeDeleteHooks = append(cmfSendcodeBeforeDeleteHooks, cmfSendcodeHook)
	case boil.BeforeUpsertHook:
		cmfSendcodeBeforeUpsertHooks = append(cmfSendcodeBeforeUpsertHooks, cmfSendcodeHook)
	case boil.AfterInsertHook:
		cmfSendcodeAfterInsertHooks = append(cmfSendcodeAfterInsertHooks, cmfSendcodeHook)
	case boil.AfterSelectHook:
		cmfSendcodeAfterSelectHooks = append(cmfSendcodeAfterSelectHooks, cmfSendcodeHook)
	case boil.AfterUpdateHook:
		cmfSendcodeAfterUpdateHooks = append(cmfSendcodeAfterUpdateHooks, cmfSendcodeHook)
	case boil.AfterDeleteHook:
		cmfSendcodeAfterDeleteHooks = append(cmfSendcodeAfterDeleteHooks, cmfSendcodeHook)
	case boil.AfterUpsertHook:
		cmfSendcodeAfterUpsertHooks = append(cmfSendcodeAfterUpsertHooks, cmfSendcodeHook)
	}
}

// One returns a single cmfSendcode record from the query.
func (q cmfSendcodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFSendcode, error) {
	o := &CMFSendcode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_sendcode")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFSendcode records from the query.
func (q cmfSendcodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFSendcodeSlice, error) {
	var o []*CMFSendcode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFSendcode slice")
	}

	if len(cmfSendcodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFSendcode records in the query.
func (q cmfSendcodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_sendcode rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfSendcodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_sendcode exists")
	}

	return count > 0, nil
}

// CMFSendcodes retrieves all the records using an executor.
func CMFSendcodes(mods ...qm.QueryMod) cmfSendcodeQuery {
	mods = append(mods, qm.From("`cmf_sendcode`"))
	return cmfSendcodeQuery{NewQuery(mods...)}
}

// FindCMFSendcode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFSendcode(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFSendcode, error) {
	cmfSendcodeObj := &CMFSendcode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_sendcode` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfSendcodeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_sendcode")
	}

	return cmfSendcodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFSendcode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_sendcode provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfSendcodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfSendcodeInsertCacheMut.RLock()
	cache, cached := cmfSendcodeInsertCache[key]
	cmfSendcodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfSendcodeAllColumns,
			cmfSendcodeColumnsWithDefault,
			cmfSendcodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfSendcodeType, cmfSendcodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfSendcodeType, cmfSendcodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_sendcode` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_sendcode` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_sendcode` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfSendcodePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_sendcode")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfSendcodeMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_sendcode")
	}

CacheNoHooks:
	if !cached {
		cmfSendcodeInsertCacheMut.Lock()
		cmfSendcodeInsertCache[key] = cache
		cmfSendcodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFSendcode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFSendcode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfSendcodeUpdateCacheMut.RLock()
	cache, cached := cmfSendcodeUpdateCache[key]
	cmfSendcodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfSendcodeAllColumns,
			cmfSendcodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_sendcode, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_sendcode` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfSendcodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfSendcodeType, cmfSendcodeMapping, append(wl, cmfSendcodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_sendcode row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_sendcode")
	}

	if !cached {
		cmfSendcodeUpdateCacheMut.Lock()
		cmfSendcodeUpdateCache[key] = cache
		cmfSendcodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfSendcodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_sendcode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_sendcode")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFSendcodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfSendcodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_sendcode` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfSendcodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfSendcode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfSendcode")
	}
	return rowsAff, nil
}

var mySQLCMFSendcodeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFSendcode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_sendcode provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfSendcodeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFSendcodeUniqueColumns, o)

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

	cmfSendcodeUpsertCacheMut.RLock()
	cache, cached := cmfSendcodeUpsertCache[key]
	cmfSendcodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfSendcodeAllColumns,
			cmfSendcodeColumnsWithDefault,
			cmfSendcodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfSendcodeAllColumns,
			cmfSendcodePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_sendcode, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_sendcode`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_sendcode` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfSendcodeType, cmfSendcodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfSendcodeType, cmfSendcodeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_sendcode")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfSendcodeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfSendcodeType, cmfSendcodeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_sendcode")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_sendcode")
	}

CacheNoHooks:
	if !cached {
		cmfSendcodeUpsertCacheMut.Lock()
		cmfSendcodeUpsertCache[key] = cache
		cmfSendcodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFSendcode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFSendcode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFSendcode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfSendcodePrimaryKeyMapping)
	sql := "DELETE FROM `cmf_sendcode` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_sendcode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_sendcode")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfSendcodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfSendcodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_sendcode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_sendcode")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFSendcodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfSendcodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfSendcodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_sendcode` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfSendcodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfSendcode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_sendcode")
	}

	if len(cmfSendcodeAfterDeleteHooks) != 0 {
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
func (o *CMFSendcode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFSendcode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFSendcodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFSendcodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfSendcodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_sendcode`.* FROM `cmf_sendcode` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfSendcodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFSendcodeSlice")
	}

	*o = slice

	return nil
}

// CMFSendcodeExists checks if the CMFSendcode row exists.
func CMFSendcodeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_sendcode` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_sendcode exists")
	}

	return exists, nil
}