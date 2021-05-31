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

// CMFVideoClass is an object representing the database table.
type CMFVideoClass struct {
	ID        uint   `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ListOrder int    `boil:"list_order" json:"list_order" toml:"list_order" yaml:"list_order"`

	R *cmfVideoClassR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfVideoClassL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFVideoClassColumns = struct {
	ID        string
	Name      string
	ListOrder string
}{
	ID:        "id",
	Name:      "name",
	ListOrder: "list_order",
}

// Generated where

var CMFVideoClassWhere = struct {
	ID        whereHelperuint
	Name      whereHelperstring
	ListOrder whereHelperint
}{
	ID:        whereHelperuint{field: "`cmf_video_class`.`id`"},
	Name:      whereHelperstring{field: "`cmf_video_class`.`name`"},
	ListOrder: whereHelperint{field: "`cmf_video_class`.`list_order`"},
}

// CMFVideoClassRels is where relationship names are stored.
var CMFVideoClassRels = struct {
}{}

// cmfVideoClassR is where relationships are stored.
type cmfVideoClassR struct {
}

// NewStruct creates a new relationship struct
func (*cmfVideoClassR) NewStruct() *cmfVideoClassR {
	return &cmfVideoClassR{}
}

// cmfVideoClassL is where Load methods for each relationship are stored.
type cmfVideoClassL struct{}

var (
	cmfVideoClassAllColumns            = []string{"id", "name", "list_order"}
	cmfVideoClassColumnsWithoutDefault = []string{"name"}
	cmfVideoClassColumnsWithDefault    = []string{"id", "list_order"}
	cmfVideoClassPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFVideoClassSlice is an alias for a slice of pointers to CMFVideoClass.
	// This should generally be used opposed to []CMFVideoClass.
	CMFVideoClassSlice []*CMFVideoClass
	// CMFVideoClassHook is the signature for custom CMFVideoClass hook methods
	CMFVideoClassHook func(context.Context, boil.ContextExecutor, *CMFVideoClass) error

	cmfVideoClassQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfVideoClassType                 = reflect.TypeOf(&CMFVideoClass{})
	cmfVideoClassMapping              = queries.MakeStructMapping(cmfVideoClassType)
	cmfVideoClassPrimaryKeyMapping, _ = queries.BindMapping(cmfVideoClassType, cmfVideoClassMapping, cmfVideoClassPrimaryKeyColumns)
	cmfVideoClassInsertCacheMut       sync.RWMutex
	cmfVideoClassInsertCache          = make(map[string]insertCache)
	cmfVideoClassUpdateCacheMut       sync.RWMutex
	cmfVideoClassUpdateCache          = make(map[string]updateCache)
	cmfVideoClassUpsertCacheMut       sync.RWMutex
	cmfVideoClassUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfVideoClassBeforeInsertHooks []CMFVideoClassHook
var cmfVideoClassBeforeUpdateHooks []CMFVideoClassHook
var cmfVideoClassBeforeDeleteHooks []CMFVideoClassHook
var cmfVideoClassBeforeUpsertHooks []CMFVideoClassHook

var cmfVideoClassAfterInsertHooks []CMFVideoClassHook
var cmfVideoClassAfterSelectHooks []CMFVideoClassHook
var cmfVideoClassAfterUpdateHooks []CMFVideoClassHook
var cmfVideoClassAfterDeleteHooks []CMFVideoClassHook
var cmfVideoClassAfterUpsertHooks []CMFVideoClassHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFVideoClass) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFVideoClass) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFVideoClass) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFVideoClass) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFVideoClass) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFVideoClass) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFVideoClass) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFVideoClass) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFVideoClass) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoClassAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFVideoClassHook registers your hook function for all future operations.
func AddCMFVideoClassHook(hookPoint boil.HookPoint, cmfVideoClassHook CMFVideoClassHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfVideoClassBeforeInsertHooks = append(cmfVideoClassBeforeInsertHooks, cmfVideoClassHook)
	case boil.BeforeUpdateHook:
		cmfVideoClassBeforeUpdateHooks = append(cmfVideoClassBeforeUpdateHooks, cmfVideoClassHook)
	case boil.BeforeDeleteHook:
		cmfVideoClassBeforeDeleteHooks = append(cmfVideoClassBeforeDeleteHooks, cmfVideoClassHook)
	case boil.BeforeUpsertHook:
		cmfVideoClassBeforeUpsertHooks = append(cmfVideoClassBeforeUpsertHooks, cmfVideoClassHook)
	case boil.AfterInsertHook:
		cmfVideoClassAfterInsertHooks = append(cmfVideoClassAfterInsertHooks, cmfVideoClassHook)
	case boil.AfterSelectHook:
		cmfVideoClassAfterSelectHooks = append(cmfVideoClassAfterSelectHooks, cmfVideoClassHook)
	case boil.AfterUpdateHook:
		cmfVideoClassAfterUpdateHooks = append(cmfVideoClassAfterUpdateHooks, cmfVideoClassHook)
	case boil.AfterDeleteHook:
		cmfVideoClassAfterDeleteHooks = append(cmfVideoClassAfterDeleteHooks, cmfVideoClassHook)
	case boil.AfterUpsertHook:
		cmfVideoClassAfterUpsertHooks = append(cmfVideoClassAfterUpsertHooks, cmfVideoClassHook)
	}
}

// One returns a single cmfVideoClass record from the query.
func (q cmfVideoClassQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFVideoClass, error) {
	o := &CMFVideoClass{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_video_class")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFVideoClass records from the query.
func (q cmfVideoClassQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFVideoClassSlice, error) {
	var o []*CMFVideoClass

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFVideoClass slice")
	}

	if len(cmfVideoClassAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFVideoClass records in the query.
func (q cmfVideoClassQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_video_class rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfVideoClassQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_video_class exists")
	}

	return count > 0, nil
}

// CMFVideoClasses retrieves all the records using an executor.
func CMFVideoClasses(mods ...qm.QueryMod) cmfVideoClassQuery {
	mods = append(mods, qm.From("`cmf_video_class`"))
	return cmfVideoClassQuery{NewQuery(mods...)}
}

// FindCMFVideoClass retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFVideoClass(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CMFVideoClass, error) {
	cmfVideoClassObj := &CMFVideoClass{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_video_class` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfVideoClassObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_video_class")
	}

	return cmfVideoClassObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFVideoClass) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_video_class provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfVideoClassColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfVideoClassInsertCacheMut.RLock()
	cache, cached := cmfVideoClassInsertCache[key]
	cmfVideoClassInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfVideoClassAllColumns,
			cmfVideoClassColumnsWithDefault,
			cmfVideoClassColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfVideoClassType, cmfVideoClassMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfVideoClassType, cmfVideoClassMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_video_class` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_video_class` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_video_class` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfVideoClassPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_video_class")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfVideoClassMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_video_class")
	}

CacheNoHooks:
	if !cached {
		cmfVideoClassInsertCacheMut.Lock()
		cmfVideoClassInsertCache[key] = cache
		cmfVideoClassInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFVideoClass.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFVideoClass) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfVideoClassUpdateCacheMut.RLock()
	cache, cached := cmfVideoClassUpdateCache[key]
	cmfVideoClassUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfVideoClassAllColumns,
			cmfVideoClassPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_video_class, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_video_class` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfVideoClassPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfVideoClassType, cmfVideoClassMapping, append(wl, cmfVideoClassPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_video_class row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_video_class")
	}

	if !cached {
		cmfVideoClassUpdateCacheMut.Lock()
		cmfVideoClassUpdateCache[key] = cache
		cmfVideoClassUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfVideoClassQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_video_class")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_video_class")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFVideoClassSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoClassPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_video_class` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoClassPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfVideoClass slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfVideoClass")
	}
	return rowsAff, nil
}

var mySQLCMFVideoClassUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFVideoClass) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_video_class provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfVideoClassColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFVideoClassUniqueColumns, o)

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

	cmfVideoClassUpsertCacheMut.RLock()
	cache, cached := cmfVideoClassUpsertCache[key]
	cmfVideoClassUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfVideoClassAllColumns,
			cmfVideoClassColumnsWithDefault,
			cmfVideoClassColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfVideoClassAllColumns,
			cmfVideoClassPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_video_class, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_video_class`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_video_class` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfVideoClassType, cmfVideoClassMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfVideoClassType, cmfVideoClassMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_video_class")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfVideoClassMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfVideoClassType, cmfVideoClassMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_video_class")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_video_class")
	}

CacheNoHooks:
	if !cached {
		cmfVideoClassUpsertCacheMut.Lock()
		cmfVideoClassUpsertCache[key] = cache
		cmfVideoClassUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFVideoClass record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFVideoClass) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFVideoClass provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfVideoClassPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_video_class` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_video_class")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_video_class")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfVideoClassQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfVideoClassQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_video_class")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_video_class")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFVideoClassSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfVideoClassBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoClassPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_video_class` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoClassPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfVideoClass slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_video_class")
	}

	if len(cmfVideoClassAfterDeleteHooks) != 0 {
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
func (o *CMFVideoClass) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFVideoClass(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFVideoClassSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFVideoClassSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoClassPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_video_class`.* FROM `cmf_video_class` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoClassPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFVideoClassSlice")
	}

	*o = slice

	return nil
}

// CMFVideoClassExists checks if the CMFVideoClass row exists.
func CMFVideoClassExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_video_class` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_video_class exists")
	}

	return exists, nil
}