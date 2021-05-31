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

// CMFPlugin is an object representing the database table.
type CMFPlugin struct {
	ID          uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Type        uint8       `boil:"type" json:"type" toml:"type" yaml:"type"`
	HasAdmin    uint8       `boil:"has_admin" json:"has_admin" toml:"has_admin" yaml:"has_admin"`
	Status      uint8       `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreateTime  uint        `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Title       string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	DemoURL     string      `boil:"demo_url" json:"demo_url" toml:"demo_url" yaml:"demo_url"`
	Hooks       string      `boil:"hooks" json:"hooks" toml:"hooks" yaml:"hooks"`
	Author      string      `boil:"author" json:"author" toml:"author" yaml:"author"`
	AuthorURL   string      `boil:"author_url" json:"author_url" toml:"author_url" yaml:"author_url"`
	Version     string      `boil:"version" json:"version" toml:"version" yaml:"version"`
	Description string      `boil:"description" json:"description" toml:"description" yaml:"description"`
	Config      null.String `boil:"config" json:"config,omitempty" toml:"config" yaml:"config,omitempty"`

	R *cmfPluginR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfPluginL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFPluginColumns = struct {
	ID          string
	Type        string
	HasAdmin    string
	Status      string
	CreateTime  string
	Name        string
	Title       string
	DemoURL     string
	Hooks       string
	Author      string
	AuthorURL   string
	Version     string
	Description string
	Config      string
}{
	ID:          "id",
	Type:        "type",
	HasAdmin:    "has_admin",
	Status:      "status",
	CreateTime:  "create_time",
	Name:        "name",
	Title:       "title",
	DemoURL:     "demo_url",
	Hooks:       "hooks",
	Author:      "author",
	AuthorURL:   "author_url",
	Version:     "version",
	Description: "description",
	Config:      "config",
}

// Generated where

var CMFPluginWhere = struct {
	ID          whereHelperuint
	Type        whereHelperuint8
	HasAdmin    whereHelperuint8
	Status      whereHelperuint8
	CreateTime  whereHelperuint
	Name        whereHelperstring
	Title       whereHelperstring
	DemoURL     whereHelperstring
	Hooks       whereHelperstring
	Author      whereHelperstring
	AuthorURL   whereHelperstring
	Version     whereHelperstring
	Description whereHelperstring
	Config      whereHelpernull_String
}{
	ID:          whereHelperuint{field: "`cmf_plugin`.`id`"},
	Type:        whereHelperuint8{field: "`cmf_plugin`.`type`"},
	HasAdmin:    whereHelperuint8{field: "`cmf_plugin`.`has_admin`"},
	Status:      whereHelperuint8{field: "`cmf_plugin`.`status`"},
	CreateTime:  whereHelperuint{field: "`cmf_plugin`.`create_time`"},
	Name:        whereHelperstring{field: "`cmf_plugin`.`name`"},
	Title:       whereHelperstring{field: "`cmf_plugin`.`title`"},
	DemoURL:     whereHelperstring{field: "`cmf_plugin`.`demo_url`"},
	Hooks:       whereHelperstring{field: "`cmf_plugin`.`hooks`"},
	Author:      whereHelperstring{field: "`cmf_plugin`.`author`"},
	AuthorURL:   whereHelperstring{field: "`cmf_plugin`.`author_url`"},
	Version:     whereHelperstring{field: "`cmf_plugin`.`version`"},
	Description: whereHelperstring{field: "`cmf_plugin`.`description`"},
	Config:      whereHelpernull_String{field: "`cmf_plugin`.`config`"},
}

// CMFPluginRels is where relationship names are stored.
var CMFPluginRels = struct {
}{}

// cmfPluginR is where relationships are stored.
type cmfPluginR struct {
}

// NewStruct creates a new relationship struct
func (*cmfPluginR) NewStruct() *cmfPluginR {
	return &cmfPluginR{}
}

// cmfPluginL is where Load methods for each relationship are stored.
type cmfPluginL struct{}

var (
	cmfPluginAllColumns            = []string{"id", "type", "has_admin", "status", "create_time", "name", "title", "demo_url", "hooks", "author", "author_url", "version", "description", "config"}
	cmfPluginColumnsWithoutDefault = []string{"name", "title", "demo_url", "hooks", "author", "author_url", "version", "description", "config"}
	cmfPluginColumnsWithDefault    = []string{"id", "type", "has_admin", "status", "create_time"}
	cmfPluginPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFPluginSlice is an alias for a slice of pointers to CMFPlugin.
	// This should generally be used opposed to []CMFPlugin.
	CMFPluginSlice []*CMFPlugin
	// CMFPluginHook is the signature for custom CMFPlugin hook methods
	CMFPluginHook func(context.Context, boil.ContextExecutor, *CMFPlugin) error

	cmfPluginQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfPluginType                 = reflect.TypeOf(&CMFPlugin{})
	cmfPluginMapping              = queries.MakeStructMapping(cmfPluginType)
	cmfPluginPrimaryKeyMapping, _ = queries.BindMapping(cmfPluginType, cmfPluginMapping, cmfPluginPrimaryKeyColumns)
	cmfPluginInsertCacheMut       sync.RWMutex
	cmfPluginInsertCache          = make(map[string]insertCache)
	cmfPluginUpdateCacheMut       sync.RWMutex
	cmfPluginUpdateCache          = make(map[string]updateCache)
	cmfPluginUpsertCacheMut       sync.RWMutex
	cmfPluginUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfPluginBeforeInsertHooks []CMFPluginHook
var cmfPluginBeforeUpdateHooks []CMFPluginHook
var cmfPluginBeforeDeleteHooks []CMFPluginHook
var cmfPluginBeforeUpsertHooks []CMFPluginHook

var cmfPluginAfterInsertHooks []CMFPluginHook
var cmfPluginAfterSelectHooks []CMFPluginHook
var cmfPluginAfterUpdateHooks []CMFPluginHook
var cmfPluginAfterDeleteHooks []CMFPluginHook
var cmfPluginAfterUpsertHooks []CMFPluginHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFPlugin) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFPlugin) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFPlugin) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFPlugin) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFPlugin) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFPlugin) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFPlugin) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFPlugin) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFPlugin) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPluginAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFPluginHook registers your hook function for all future operations.
func AddCMFPluginHook(hookPoint boil.HookPoint, cmfPluginHook CMFPluginHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfPluginBeforeInsertHooks = append(cmfPluginBeforeInsertHooks, cmfPluginHook)
	case boil.BeforeUpdateHook:
		cmfPluginBeforeUpdateHooks = append(cmfPluginBeforeUpdateHooks, cmfPluginHook)
	case boil.BeforeDeleteHook:
		cmfPluginBeforeDeleteHooks = append(cmfPluginBeforeDeleteHooks, cmfPluginHook)
	case boil.BeforeUpsertHook:
		cmfPluginBeforeUpsertHooks = append(cmfPluginBeforeUpsertHooks, cmfPluginHook)
	case boil.AfterInsertHook:
		cmfPluginAfterInsertHooks = append(cmfPluginAfterInsertHooks, cmfPluginHook)
	case boil.AfterSelectHook:
		cmfPluginAfterSelectHooks = append(cmfPluginAfterSelectHooks, cmfPluginHook)
	case boil.AfterUpdateHook:
		cmfPluginAfterUpdateHooks = append(cmfPluginAfterUpdateHooks, cmfPluginHook)
	case boil.AfterDeleteHook:
		cmfPluginAfterDeleteHooks = append(cmfPluginAfterDeleteHooks, cmfPluginHook)
	case boil.AfterUpsertHook:
		cmfPluginAfterUpsertHooks = append(cmfPluginAfterUpsertHooks, cmfPluginHook)
	}
}

// One returns a single cmfPlugin record from the query.
func (q cmfPluginQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFPlugin, error) {
	o := &CMFPlugin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_plugin")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFPlugin records from the query.
func (q cmfPluginQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFPluginSlice, error) {
	var o []*CMFPlugin

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFPlugin slice")
	}

	if len(cmfPluginAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFPlugin records in the query.
func (q cmfPluginQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_plugin rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfPluginQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_plugin exists")
	}

	return count > 0, nil
}

// CMFPlugins retrieves all the records using an executor.
func CMFPlugins(mods ...qm.QueryMod) cmfPluginQuery {
	mods = append(mods, qm.From("`cmf_plugin`"))
	return cmfPluginQuery{NewQuery(mods...)}
}

// FindCMFPlugin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFPlugin(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CMFPlugin, error) {
	cmfPluginObj := &CMFPlugin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_plugin` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfPluginObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_plugin")
	}

	return cmfPluginObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFPlugin) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_plugin provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPluginColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfPluginInsertCacheMut.RLock()
	cache, cached := cmfPluginInsertCache[key]
	cmfPluginInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfPluginAllColumns,
			cmfPluginColumnsWithDefault,
			cmfPluginColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfPluginType, cmfPluginMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfPluginType, cmfPluginMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_plugin` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_plugin` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_plugin` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfPluginPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_plugin")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPluginMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_plugin")
	}

CacheNoHooks:
	if !cached {
		cmfPluginInsertCacheMut.Lock()
		cmfPluginInsertCache[key] = cache
		cmfPluginInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFPlugin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFPlugin) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfPluginUpdateCacheMut.RLock()
	cache, cached := cmfPluginUpdateCache[key]
	cmfPluginUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfPluginAllColumns,
			cmfPluginPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_plugin, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_plugin` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfPluginPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfPluginType, cmfPluginMapping, append(wl, cmfPluginPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_plugin row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_plugin")
	}

	if !cached {
		cmfPluginUpdateCacheMut.Lock()
		cmfPluginUpdateCache[key] = cache
		cmfPluginUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfPluginQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_plugin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_plugin")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFPluginSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPluginPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_plugin` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPluginPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfPlugin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfPlugin")
	}
	return rowsAff, nil
}

var mySQLCMFPluginUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFPlugin) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_plugin provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPluginColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFPluginUniqueColumns, o)

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

	cmfPluginUpsertCacheMut.RLock()
	cache, cached := cmfPluginUpsertCache[key]
	cmfPluginUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfPluginAllColumns,
			cmfPluginColumnsWithDefault,
			cmfPluginColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfPluginAllColumns,
			cmfPluginPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_plugin, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_plugin`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_plugin` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfPluginType, cmfPluginMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfPluginType, cmfPluginMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_plugin")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPluginMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfPluginType, cmfPluginMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_plugin")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_plugin")
	}

CacheNoHooks:
	if !cached {
		cmfPluginUpsertCacheMut.Lock()
		cmfPluginUpsertCache[key] = cache
		cmfPluginUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFPlugin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFPlugin) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFPlugin provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfPluginPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_plugin` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_plugin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_plugin")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfPluginQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfPluginQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_plugin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_plugin")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFPluginSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfPluginBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPluginPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_plugin` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPluginPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfPlugin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_plugin")
	}

	if len(cmfPluginAfterDeleteHooks) != 0 {
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
func (o *CMFPlugin) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFPlugin(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFPluginSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFPluginSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPluginPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_plugin`.* FROM `cmf_plugin` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPluginPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFPluginSlice")
	}

	*o = slice

	return nil
}

// CMFPluginExists checks if the CMFPlugin row exists.
func CMFPluginExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_plugin` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_plugin exists")
	}

	return exists, nil
}