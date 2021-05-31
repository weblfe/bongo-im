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

// CMFVideoView is an object representing the database table.
type CMFVideoView struct {
	ID      int `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID     int `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Videoid int `boil:"videoid" json:"videoid" toml:"videoid" yaml:"videoid"`
	Addtime int `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`

	R *cmfVideoViewR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfVideoViewL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFVideoViewColumns = struct {
	ID      string
	UID     string
	Videoid string
	Addtime string
}{
	ID:      "id",
	UID:     "uid",
	Videoid: "videoid",
	Addtime: "addtime",
}

// Generated where

var CMFVideoViewWhere = struct {
	ID      whereHelperint
	UID     whereHelperint
	Videoid whereHelperint
	Addtime whereHelperint
}{
	ID:      whereHelperint{field: "`cmf_video_view`.`id`"},
	UID:     whereHelperint{field: "`cmf_video_view`.`uid`"},
	Videoid: whereHelperint{field: "`cmf_video_view`.`videoid`"},
	Addtime: whereHelperint{field: "`cmf_video_view`.`addtime`"},
}

// CMFVideoViewRels is where relationship names are stored.
var CMFVideoViewRels = struct {
}{}

// cmfVideoViewR is where relationships are stored.
type cmfVideoViewR struct {
}

// NewStruct creates a new relationship struct
func (*cmfVideoViewR) NewStruct() *cmfVideoViewR {
	return &cmfVideoViewR{}
}

// cmfVideoViewL is where Load methods for each relationship are stored.
type cmfVideoViewL struct{}

var (
	cmfVideoViewAllColumns            = []string{"id", "uid", "videoid", "addtime"}
	cmfVideoViewColumnsWithoutDefault = []string{}
	cmfVideoViewColumnsWithDefault    = []string{"id", "uid", "videoid", "addtime"}
	cmfVideoViewPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFVideoViewSlice is an alias for a slice of pointers to CMFVideoView.
	// This should generally be used opposed to []CMFVideoView.
	CMFVideoViewSlice []*CMFVideoView
	// CMFVideoViewHook is the signature for custom CMFVideoView hook methods
	CMFVideoViewHook func(context.Context, boil.ContextExecutor, *CMFVideoView) error

	cmfVideoViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfVideoViewType                 = reflect.TypeOf(&CMFVideoView{})
	cmfVideoViewMapping              = queries.MakeStructMapping(cmfVideoViewType)
	cmfVideoViewPrimaryKeyMapping, _ = queries.BindMapping(cmfVideoViewType, cmfVideoViewMapping, cmfVideoViewPrimaryKeyColumns)
	cmfVideoViewInsertCacheMut       sync.RWMutex
	cmfVideoViewInsertCache          = make(map[string]insertCache)
	cmfVideoViewUpdateCacheMut       sync.RWMutex
	cmfVideoViewUpdateCache          = make(map[string]updateCache)
	cmfVideoViewUpsertCacheMut       sync.RWMutex
	cmfVideoViewUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfVideoViewBeforeInsertHooks []CMFVideoViewHook
var cmfVideoViewBeforeUpdateHooks []CMFVideoViewHook
var cmfVideoViewBeforeDeleteHooks []CMFVideoViewHook
var cmfVideoViewBeforeUpsertHooks []CMFVideoViewHook

var cmfVideoViewAfterInsertHooks []CMFVideoViewHook
var cmfVideoViewAfterSelectHooks []CMFVideoViewHook
var cmfVideoViewAfterUpdateHooks []CMFVideoViewHook
var cmfVideoViewAfterDeleteHooks []CMFVideoViewHook
var cmfVideoViewAfterUpsertHooks []CMFVideoViewHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFVideoView) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFVideoView) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFVideoView) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFVideoView) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFVideoView) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFVideoView) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFVideoView) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFVideoView) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFVideoView) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoViewAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFVideoViewHook registers your hook function for all future operations.
func AddCMFVideoViewHook(hookPoint boil.HookPoint, cmfVideoViewHook CMFVideoViewHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfVideoViewBeforeInsertHooks = append(cmfVideoViewBeforeInsertHooks, cmfVideoViewHook)
	case boil.BeforeUpdateHook:
		cmfVideoViewBeforeUpdateHooks = append(cmfVideoViewBeforeUpdateHooks, cmfVideoViewHook)
	case boil.BeforeDeleteHook:
		cmfVideoViewBeforeDeleteHooks = append(cmfVideoViewBeforeDeleteHooks, cmfVideoViewHook)
	case boil.BeforeUpsertHook:
		cmfVideoViewBeforeUpsertHooks = append(cmfVideoViewBeforeUpsertHooks, cmfVideoViewHook)
	case boil.AfterInsertHook:
		cmfVideoViewAfterInsertHooks = append(cmfVideoViewAfterInsertHooks, cmfVideoViewHook)
	case boil.AfterSelectHook:
		cmfVideoViewAfterSelectHooks = append(cmfVideoViewAfterSelectHooks, cmfVideoViewHook)
	case boil.AfterUpdateHook:
		cmfVideoViewAfterUpdateHooks = append(cmfVideoViewAfterUpdateHooks, cmfVideoViewHook)
	case boil.AfterDeleteHook:
		cmfVideoViewAfterDeleteHooks = append(cmfVideoViewAfterDeleteHooks, cmfVideoViewHook)
	case boil.AfterUpsertHook:
		cmfVideoViewAfterUpsertHooks = append(cmfVideoViewAfterUpsertHooks, cmfVideoViewHook)
	}
}

// One returns a single cmfVideoView record from the query.
func (q cmfVideoViewQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFVideoView, error) {
	o := &CMFVideoView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_video_view")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFVideoView records from the query.
func (q cmfVideoViewQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFVideoViewSlice, error) {
	var o []*CMFVideoView

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFVideoView slice")
	}

	if len(cmfVideoViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFVideoView records in the query.
func (q cmfVideoViewQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_video_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfVideoViewQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_video_view exists")
	}

	return count > 0, nil
}

// CMFVideoViews retrieves all the records using an executor.
func CMFVideoViews(mods ...qm.QueryMod) cmfVideoViewQuery {
	mods = append(mods, qm.From("`cmf_video_view`"))
	return cmfVideoViewQuery{NewQuery(mods...)}
}

// FindCMFVideoView retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFVideoView(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFVideoView, error) {
	cmfVideoViewObj := &CMFVideoView{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_video_view` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfVideoViewObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_video_view")
	}

	return cmfVideoViewObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFVideoView) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_video_view provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfVideoViewColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfVideoViewInsertCacheMut.RLock()
	cache, cached := cmfVideoViewInsertCache[key]
	cmfVideoViewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfVideoViewAllColumns,
			cmfVideoViewColumnsWithDefault,
			cmfVideoViewColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfVideoViewType, cmfVideoViewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfVideoViewType, cmfVideoViewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_video_view` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_video_view` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_video_view` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfVideoViewPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_video_view")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfVideoViewMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_video_view")
	}

CacheNoHooks:
	if !cached {
		cmfVideoViewInsertCacheMut.Lock()
		cmfVideoViewInsertCache[key] = cache
		cmfVideoViewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFVideoView.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFVideoView) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfVideoViewUpdateCacheMut.RLock()
	cache, cached := cmfVideoViewUpdateCache[key]
	cmfVideoViewUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfVideoViewAllColumns,
			cmfVideoViewPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_video_view, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_video_view` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfVideoViewPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfVideoViewType, cmfVideoViewMapping, append(wl, cmfVideoViewPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_video_view row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_video_view")
	}

	if !cached {
		cmfVideoViewUpdateCacheMut.Lock()
		cmfVideoViewUpdateCache[key] = cache
		cmfVideoViewUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfVideoViewQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_video_view")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_video_view")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFVideoViewSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoViewPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_video_view` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoViewPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfVideoView slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfVideoView")
	}
	return rowsAff, nil
}

var mySQLCMFVideoViewUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFVideoView) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_video_view provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfVideoViewColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFVideoViewUniqueColumns, o)

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

	cmfVideoViewUpsertCacheMut.RLock()
	cache, cached := cmfVideoViewUpsertCache[key]
	cmfVideoViewUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfVideoViewAllColumns,
			cmfVideoViewColumnsWithDefault,
			cmfVideoViewColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfVideoViewAllColumns,
			cmfVideoViewPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_video_view, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_video_view`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_video_view` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfVideoViewType, cmfVideoViewMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfVideoViewType, cmfVideoViewMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_video_view")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfVideoViewMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfVideoViewType, cmfVideoViewMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_video_view")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_video_view")
	}

CacheNoHooks:
	if !cached {
		cmfVideoViewUpsertCacheMut.Lock()
		cmfVideoViewUpsertCache[key] = cache
		cmfVideoViewUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFVideoView record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFVideoView) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFVideoView provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfVideoViewPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_video_view` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_video_view")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_video_view")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfVideoViewQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfVideoViewQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_video_view")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_video_view")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFVideoViewSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfVideoViewBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoViewPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_video_view` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoViewPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfVideoView slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_video_view")
	}

	if len(cmfVideoViewAfterDeleteHooks) != 0 {
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
func (o *CMFVideoView) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFVideoView(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFVideoViewSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFVideoViewSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoViewPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_video_view`.* FROM `cmf_video_view` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoViewPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFVideoViewSlice")
	}

	*o = slice

	return nil
}

// CMFVideoViewExists checks if the CMFVideoView row exists.
func CMFVideoViewExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_video_view` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_video_view exists")
	}

	return exists, nil
}