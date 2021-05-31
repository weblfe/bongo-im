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

// CMFHookPlugin is an object representing the database table.
type CMFHookPlugin struct {
	ID        uint    `boil:"id" json:"id" toml:"id" yaml:"id"`
	ListOrder float32 `boil:"list_order" json:"list_order" toml:"list_order" yaml:"list_order"`
	Status    int8    `boil:"status" json:"status" toml:"status" yaml:"status"`
	Hook      string  `boil:"hook" json:"hook" toml:"hook" yaml:"hook"`
	Plugin    string  `boil:"plugin" json:"plugin" toml:"plugin" yaml:"plugin"`

	R *cmfHookPluginR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfHookPluginL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFHookPluginColumns = struct {
	ID        string
	ListOrder string
	Status    string
	Hook      string
	Plugin    string
}{
	ID:        "id",
	ListOrder: "list_order",
	Status:    "status",
	Hook:      "hook",
	Plugin:    "plugin",
}

// Generated where

var CMFHookPluginWhere = struct {
	ID        whereHelperuint
	ListOrder whereHelperfloat32
	Status    whereHelperint8
	Hook      whereHelperstring
	Plugin    whereHelperstring
}{
	ID:        whereHelperuint{field: "`cmf_hook_plugin`.`id`"},
	ListOrder: whereHelperfloat32{field: "`cmf_hook_plugin`.`list_order`"},
	Status:    whereHelperint8{field: "`cmf_hook_plugin`.`status`"},
	Hook:      whereHelperstring{field: "`cmf_hook_plugin`.`hook`"},
	Plugin:    whereHelperstring{field: "`cmf_hook_plugin`.`plugin`"},
}

// CMFHookPluginRels is where relationship names are stored.
var CMFHookPluginRels = struct {
}{}

// cmfHookPluginR is where relationships are stored.
type cmfHookPluginR struct {
}

// NewStruct creates a new relationship struct
func (*cmfHookPluginR) NewStruct() *cmfHookPluginR {
	return &cmfHookPluginR{}
}

// cmfHookPluginL is where Load methods for each relationship are stored.
type cmfHookPluginL struct{}

var (
	cmfHookPluginAllColumns            = []string{"id", "list_order", "status", "hook", "plugin"}
	cmfHookPluginColumnsWithoutDefault = []string{"hook", "plugin"}
	cmfHookPluginColumnsWithDefault    = []string{"id", "list_order", "status"}
	cmfHookPluginPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFHookPluginSlice is an alias for a slice of pointers to CMFHookPlugin.
	// This should generally be used opposed to []CMFHookPlugin.
	CMFHookPluginSlice []*CMFHookPlugin
	// CMFHookPluginHook is the signature for custom CMFHookPlugin hook methods
	CMFHookPluginHook func(context.Context, boil.ContextExecutor, *CMFHookPlugin) error

	cmfHookPluginQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfHookPluginType                 = reflect.TypeOf(&CMFHookPlugin{})
	cmfHookPluginMapping              = queries.MakeStructMapping(cmfHookPluginType)
	cmfHookPluginPrimaryKeyMapping, _ = queries.BindMapping(cmfHookPluginType, cmfHookPluginMapping, cmfHookPluginPrimaryKeyColumns)
	cmfHookPluginInsertCacheMut       sync.RWMutex
	cmfHookPluginInsertCache          = make(map[string]insertCache)
	cmfHookPluginUpdateCacheMut       sync.RWMutex
	cmfHookPluginUpdateCache          = make(map[string]updateCache)
	cmfHookPluginUpsertCacheMut       sync.RWMutex
	cmfHookPluginUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfHookPluginBeforeInsertHooks []CMFHookPluginHook
var cmfHookPluginBeforeUpdateHooks []CMFHookPluginHook
var cmfHookPluginBeforeDeleteHooks []CMFHookPluginHook
var cmfHookPluginBeforeUpsertHooks []CMFHookPluginHook

var cmfHookPluginAfterInsertHooks []CMFHookPluginHook
var cmfHookPluginAfterSelectHooks []CMFHookPluginHook
var cmfHookPluginAfterUpdateHooks []CMFHookPluginHook
var cmfHookPluginAfterDeleteHooks []CMFHookPluginHook
var cmfHookPluginAfterUpsertHooks []CMFHookPluginHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFHookPlugin) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFHookPlugin) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFHookPlugin) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFHookPlugin) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFHookPlugin) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFHookPlugin) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFHookPlugin) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFHookPlugin) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFHookPlugin) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfHookPluginAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFHookPluginHook registers your hook function for all future operations.
func AddCMFHookPluginHook(hookPoint boil.HookPoint, cmfHookPluginHook CMFHookPluginHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfHookPluginBeforeInsertHooks = append(cmfHookPluginBeforeInsertHooks, cmfHookPluginHook)
	case boil.BeforeUpdateHook:
		cmfHookPluginBeforeUpdateHooks = append(cmfHookPluginBeforeUpdateHooks, cmfHookPluginHook)
	case boil.BeforeDeleteHook:
		cmfHookPluginBeforeDeleteHooks = append(cmfHookPluginBeforeDeleteHooks, cmfHookPluginHook)
	case boil.BeforeUpsertHook:
		cmfHookPluginBeforeUpsertHooks = append(cmfHookPluginBeforeUpsertHooks, cmfHookPluginHook)
	case boil.AfterInsertHook:
		cmfHookPluginAfterInsertHooks = append(cmfHookPluginAfterInsertHooks, cmfHookPluginHook)
	case boil.AfterSelectHook:
		cmfHookPluginAfterSelectHooks = append(cmfHookPluginAfterSelectHooks, cmfHookPluginHook)
	case boil.AfterUpdateHook:
		cmfHookPluginAfterUpdateHooks = append(cmfHookPluginAfterUpdateHooks, cmfHookPluginHook)
	case boil.AfterDeleteHook:
		cmfHookPluginAfterDeleteHooks = append(cmfHookPluginAfterDeleteHooks, cmfHookPluginHook)
	case boil.AfterUpsertHook:
		cmfHookPluginAfterUpsertHooks = append(cmfHookPluginAfterUpsertHooks, cmfHookPluginHook)
	}
}

// One returns a single cmfHookPlugin record from the query.
func (q cmfHookPluginQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFHookPlugin, error) {
	o := &CMFHookPlugin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_hook_plugin")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFHookPlugin records from the query.
func (q cmfHookPluginQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFHookPluginSlice, error) {
	var o []*CMFHookPlugin

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFHookPlugin slice")
	}

	if len(cmfHookPluginAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFHookPlugin records in the query.
func (q cmfHookPluginQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_hook_plugin rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfHookPluginQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_hook_plugin exists")
	}

	return count > 0, nil
}

// CMFHookPlugins retrieves all the records using an executor.
func CMFHookPlugins(mods ...qm.QueryMod) cmfHookPluginQuery {
	mods = append(mods, qm.From("`cmf_hook_plugin`"))
	return cmfHookPluginQuery{NewQuery(mods...)}
}

// FindCMFHookPlugin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFHookPlugin(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CMFHookPlugin, error) {
	cmfHookPluginObj := &CMFHookPlugin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_hook_plugin` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfHookPluginObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_hook_plugin")
	}

	return cmfHookPluginObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFHookPlugin) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_hook_plugin provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfHookPluginColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfHookPluginInsertCacheMut.RLock()
	cache, cached := cmfHookPluginInsertCache[key]
	cmfHookPluginInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfHookPluginAllColumns,
			cmfHookPluginColumnsWithDefault,
			cmfHookPluginColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfHookPluginType, cmfHookPluginMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfHookPluginType, cmfHookPluginMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_hook_plugin` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_hook_plugin` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_hook_plugin` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfHookPluginPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_hook_plugin")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfHookPluginMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_hook_plugin")
	}

CacheNoHooks:
	if !cached {
		cmfHookPluginInsertCacheMut.Lock()
		cmfHookPluginInsertCache[key] = cache
		cmfHookPluginInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFHookPlugin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFHookPlugin) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfHookPluginUpdateCacheMut.RLock()
	cache, cached := cmfHookPluginUpdateCache[key]
	cmfHookPluginUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfHookPluginAllColumns,
			cmfHookPluginPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_hook_plugin, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_hook_plugin` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfHookPluginPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfHookPluginType, cmfHookPluginMapping, append(wl, cmfHookPluginPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_hook_plugin row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_hook_plugin")
	}

	if !cached {
		cmfHookPluginUpdateCacheMut.Lock()
		cmfHookPluginUpdateCache[key] = cache
		cmfHookPluginUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfHookPluginQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_hook_plugin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_hook_plugin")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFHookPluginSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfHookPluginPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_hook_plugin` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfHookPluginPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfHookPlugin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfHookPlugin")
	}
	return rowsAff, nil
}

var mySQLCMFHookPluginUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFHookPlugin) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_hook_plugin provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfHookPluginColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFHookPluginUniqueColumns, o)

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

	cmfHookPluginUpsertCacheMut.RLock()
	cache, cached := cmfHookPluginUpsertCache[key]
	cmfHookPluginUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfHookPluginAllColumns,
			cmfHookPluginColumnsWithDefault,
			cmfHookPluginColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfHookPluginAllColumns,
			cmfHookPluginPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_hook_plugin, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_hook_plugin`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_hook_plugin` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfHookPluginType, cmfHookPluginMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfHookPluginType, cmfHookPluginMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_hook_plugin")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfHookPluginMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfHookPluginType, cmfHookPluginMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_hook_plugin")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_hook_plugin")
	}

CacheNoHooks:
	if !cached {
		cmfHookPluginUpsertCacheMut.Lock()
		cmfHookPluginUpsertCache[key] = cache
		cmfHookPluginUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFHookPlugin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFHookPlugin) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFHookPlugin provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfHookPluginPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_hook_plugin` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_hook_plugin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_hook_plugin")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfHookPluginQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfHookPluginQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_hook_plugin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_hook_plugin")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFHookPluginSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfHookPluginBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfHookPluginPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_hook_plugin` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfHookPluginPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfHookPlugin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_hook_plugin")
	}

	if len(cmfHookPluginAfterDeleteHooks) != 0 {
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
func (o *CMFHookPlugin) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFHookPlugin(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFHookPluginSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFHookPluginSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfHookPluginPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_hook_plugin`.* FROM `cmf_hook_plugin` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfHookPluginPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFHookPluginSlice")
	}

	*o = slice

	return nil
}

// CMFHookPluginExists checks if the CMFHookPlugin row exists.
func CMFHookPluginExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_hook_plugin` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_hook_plugin exists")
	}

	return exists, nil
}
