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

// CMFMusicCollection is an object representing the database table.
type CMFMusicCollection struct {
	ID         int  `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID        int  `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	MusicID    int  `boil:"music_id" json:"music_id" toml:"music_id" yaml:"music_id"`
	Addtime    int  `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Updatetime int  `boil:"updatetime" json:"updatetime" toml:"updatetime" yaml:"updatetime"`
	Status     bool `boil:"status" json:"status" toml:"status" yaml:"status"`

	R *cmfMusicCollectionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfMusicCollectionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFMusicCollectionColumns = struct {
	ID         string
	UID        string
	MusicID    string
	Addtime    string
	Updatetime string
	Status     string
}{
	ID:         "id",
	UID:        "uid",
	MusicID:    "music_id",
	Addtime:    "addtime",
	Updatetime: "updatetime",
	Status:     "status",
}

// Generated where

var CMFMusicCollectionWhere = struct {
	ID         whereHelperint
	UID        whereHelperint
	MusicID    whereHelperint
	Addtime    whereHelperint
	Updatetime whereHelperint
	Status     whereHelperbool
}{
	ID:         whereHelperint{field: "`cmf_music_collection`.`id`"},
	UID:        whereHelperint{field: "`cmf_music_collection`.`uid`"},
	MusicID:    whereHelperint{field: "`cmf_music_collection`.`music_id`"},
	Addtime:    whereHelperint{field: "`cmf_music_collection`.`addtime`"},
	Updatetime: whereHelperint{field: "`cmf_music_collection`.`updatetime`"},
	Status:     whereHelperbool{field: "`cmf_music_collection`.`status`"},
}

// CMFMusicCollectionRels is where relationship names are stored.
var CMFMusicCollectionRels = struct {
}{}

// cmfMusicCollectionR is where relationships are stored.
type cmfMusicCollectionR struct {
}

// NewStruct creates a new relationship struct
func (*cmfMusicCollectionR) NewStruct() *cmfMusicCollectionR {
	return &cmfMusicCollectionR{}
}

// cmfMusicCollectionL is where Load methods for each relationship are stored.
type cmfMusicCollectionL struct{}

var (
	cmfMusicCollectionAllColumns            = []string{"id", "uid", "music_id", "addtime", "updatetime", "status"}
	cmfMusicCollectionColumnsWithoutDefault = []string{}
	cmfMusicCollectionColumnsWithDefault    = []string{"id", "uid", "music_id", "addtime", "updatetime", "status"}
	cmfMusicCollectionPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFMusicCollectionSlice is an alias for a slice of pointers to CMFMusicCollection.
	// This should generally be used opposed to []CMFMusicCollection.
	CMFMusicCollectionSlice []*CMFMusicCollection
	// CMFMusicCollectionHook is the signature for custom CMFMusicCollection hook methods
	CMFMusicCollectionHook func(context.Context, boil.ContextExecutor, *CMFMusicCollection) error

	cmfMusicCollectionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfMusicCollectionType                 = reflect.TypeOf(&CMFMusicCollection{})
	cmfMusicCollectionMapping              = queries.MakeStructMapping(cmfMusicCollectionType)
	cmfMusicCollectionPrimaryKeyMapping, _ = queries.BindMapping(cmfMusicCollectionType, cmfMusicCollectionMapping, cmfMusicCollectionPrimaryKeyColumns)
	cmfMusicCollectionInsertCacheMut       sync.RWMutex
	cmfMusicCollectionInsertCache          = make(map[string]insertCache)
	cmfMusicCollectionUpdateCacheMut       sync.RWMutex
	cmfMusicCollectionUpdateCache          = make(map[string]updateCache)
	cmfMusicCollectionUpsertCacheMut       sync.RWMutex
	cmfMusicCollectionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfMusicCollectionBeforeInsertHooks []CMFMusicCollectionHook
var cmfMusicCollectionBeforeUpdateHooks []CMFMusicCollectionHook
var cmfMusicCollectionBeforeDeleteHooks []CMFMusicCollectionHook
var cmfMusicCollectionBeforeUpsertHooks []CMFMusicCollectionHook

var cmfMusicCollectionAfterInsertHooks []CMFMusicCollectionHook
var cmfMusicCollectionAfterSelectHooks []CMFMusicCollectionHook
var cmfMusicCollectionAfterUpdateHooks []CMFMusicCollectionHook
var cmfMusicCollectionAfterDeleteHooks []CMFMusicCollectionHook
var cmfMusicCollectionAfterUpsertHooks []CMFMusicCollectionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFMusicCollection) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFMusicCollection) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFMusicCollection) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFMusicCollection) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFMusicCollection) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFMusicCollection) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFMusicCollection) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFMusicCollection) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFMusicCollection) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicCollectionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFMusicCollectionHook registers your hook function for all future operations.
func AddCMFMusicCollectionHook(hookPoint boil.HookPoint, cmfMusicCollectionHook CMFMusicCollectionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfMusicCollectionBeforeInsertHooks = append(cmfMusicCollectionBeforeInsertHooks, cmfMusicCollectionHook)
	case boil.BeforeUpdateHook:
		cmfMusicCollectionBeforeUpdateHooks = append(cmfMusicCollectionBeforeUpdateHooks, cmfMusicCollectionHook)
	case boil.BeforeDeleteHook:
		cmfMusicCollectionBeforeDeleteHooks = append(cmfMusicCollectionBeforeDeleteHooks, cmfMusicCollectionHook)
	case boil.BeforeUpsertHook:
		cmfMusicCollectionBeforeUpsertHooks = append(cmfMusicCollectionBeforeUpsertHooks, cmfMusicCollectionHook)
	case boil.AfterInsertHook:
		cmfMusicCollectionAfterInsertHooks = append(cmfMusicCollectionAfterInsertHooks, cmfMusicCollectionHook)
	case boil.AfterSelectHook:
		cmfMusicCollectionAfterSelectHooks = append(cmfMusicCollectionAfterSelectHooks, cmfMusicCollectionHook)
	case boil.AfterUpdateHook:
		cmfMusicCollectionAfterUpdateHooks = append(cmfMusicCollectionAfterUpdateHooks, cmfMusicCollectionHook)
	case boil.AfterDeleteHook:
		cmfMusicCollectionAfterDeleteHooks = append(cmfMusicCollectionAfterDeleteHooks, cmfMusicCollectionHook)
	case boil.AfterUpsertHook:
		cmfMusicCollectionAfterUpsertHooks = append(cmfMusicCollectionAfterUpsertHooks, cmfMusicCollectionHook)
	}
}

// One returns a single cmfMusicCollection record from the query.
func (q cmfMusicCollectionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFMusicCollection, error) {
	o := &CMFMusicCollection{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_music_collection")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFMusicCollection records from the query.
func (q cmfMusicCollectionQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFMusicCollectionSlice, error) {
	var o []*CMFMusicCollection

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFMusicCollection slice")
	}

	if len(cmfMusicCollectionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFMusicCollection records in the query.
func (q cmfMusicCollectionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_music_collection rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfMusicCollectionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_music_collection exists")
	}

	return count > 0, nil
}

// CMFMusicCollections retrieves all the records using an executor.
func CMFMusicCollections(mods ...qm.QueryMod) cmfMusicCollectionQuery {
	mods = append(mods, qm.From("`cmf_music_collection`"))
	return cmfMusicCollectionQuery{NewQuery(mods...)}
}

// FindCMFMusicCollection retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFMusicCollection(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFMusicCollection, error) {
	cmfMusicCollectionObj := &CMFMusicCollection{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_music_collection` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfMusicCollectionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_music_collection")
	}

	return cmfMusicCollectionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFMusicCollection) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_music_collection provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfMusicCollectionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfMusicCollectionInsertCacheMut.RLock()
	cache, cached := cmfMusicCollectionInsertCache[key]
	cmfMusicCollectionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfMusicCollectionAllColumns,
			cmfMusicCollectionColumnsWithDefault,
			cmfMusicCollectionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfMusicCollectionType, cmfMusicCollectionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfMusicCollectionType, cmfMusicCollectionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_music_collection` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_music_collection` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_music_collection` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfMusicCollectionPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_music_collection")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfMusicCollectionMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_music_collection")
	}

CacheNoHooks:
	if !cached {
		cmfMusicCollectionInsertCacheMut.Lock()
		cmfMusicCollectionInsertCache[key] = cache
		cmfMusicCollectionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFMusicCollection.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFMusicCollection) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfMusicCollectionUpdateCacheMut.RLock()
	cache, cached := cmfMusicCollectionUpdateCache[key]
	cmfMusicCollectionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfMusicCollectionAllColumns,
			cmfMusicCollectionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_music_collection, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_music_collection` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfMusicCollectionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfMusicCollectionType, cmfMusicCollectionMapping, append(wl, cmfMusicCollectionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_music_collection row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_music_collection")
	}

	if !cached {
		cmfMusicCollectionUpdateCacheMut.Lock()
		cmfMusicCollectionUpdateCache[key] = cache
		cmfMusicCollectionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfMusicCollectionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_music_collection")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_music_collection")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFMusicCollectionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfMusicCollectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_music_collection` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfMusicCollectionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfMusicCollection slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfMusicCollection")
	}
	return rowsAff, nil
}

var mySQLCMFMusicCollectionUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFMusicCollection) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_music_collection provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfMusicCollectionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFMusicCollectionUniqueColumns, o)

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

	cmfMusicCollectionUpsertCacheMut.RLock()
	cache, cached := cmfMusicCollectionUpsertCache[key]
	cmfMusicCollectionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfMusicCollectionAllColumns,
			cmfMusicCollectionColumnsWithDefault,
			cmfMusicCollectionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfMusicCollectionAllColumns,
			cmfMusicCollectionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_music_collection, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_music_collection`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_music_collection` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfMusicCollectionType, cmfMusicCollectionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfMusicCollectionType, cmfMusicCollectionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_music_collection")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfMusicCollectionMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfMusicCollectionType, cmfMusicCollectionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_music_collection")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_music_collection")
	}

CacheNoHooks:
	if !cached {
		cmfMusicCollectionUpsertCacheMut.Lock()
		cmfMusicCollectionUpsertCache[key] = cache
		cmfMusicCollectionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFMusicCollection record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFMusicCollection) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFMusicCollection provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfMusicCollectionPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_music_collection` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_music_collection")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_music_collection")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfMusicCollectionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfMusicCollectionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_music_collection")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_music_collection")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFMusicCollectionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfMusicCollectionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfMusicCollectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_music_collection` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfMusicCollectionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfMusicCollection slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_music_collection")
	}

	if len(cmfMusicCollectionAfterDeleteHooks) != 0 {
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
func (o *CMFMusicCollection) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFMusicCollection(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFMusicCollectionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFMusicCollectionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfMusicCollectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_music_collection`.* FROM `cmf_music_collection` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfMusicCollectionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFMusicCollectionSlice")
	}

	*o = slice

	return nil
}

// CMFMusicCollectionExists checks if the CMFMusicCollection row exists.
func CMFMusicCollectionExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_music_collection` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_music_collection exists")
	}

	return exists, nil
}
