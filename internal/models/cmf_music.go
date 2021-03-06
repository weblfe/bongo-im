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

// CMFMusic is an object representing the database table.
type CMFMusic struct {
	ID         int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title      string `boil:"title" json:"title" toml:"title" yaml:"title"`
	Author     string `boil:"author" json:"author" toml:"author" yaml:"author"`
	Uploader   int    `boil:"uploader" json:"uploader" toml:"uploader" yaml:"uploader"`
	UploadType bool   `boil:"upload_type" json:"upload_type" toml:"upload_type" yaml:"upload_type"`
	ImgURL     string `boil:"img_url" json:"img_url" toml:"img_url" yaml:"img_url"`
	Length     string `boil:"length" json:"length" toml:"length" yaml:"length"`
	FileURL    string `boil:"file_url" json:"file_url" toml:"file_url" yaml:"file_url"`
	UseNums    int    `boil:"use_nums" json:"use_nums" toml:"use_nums" yaml:"use_nums"`
	Isdel      bool   `boil:"isdel" json:"isdel" toml:"isdel" yaml:"isdel"`
	Addtime    int    `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Updatetime int    `boil:"updatetime" json:"updatetime" toml:"updatetime" yaml:"updatetime"`
	ClassifyID int    `boil:"classify_id" json:"classify_id" toml:"classify_id" yaml:"classify_id"`

	R *cmfMusicR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfMusicL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFMusicColumns = struct {
	ID         string
	Title      string
	Author     string
	Uploader   string
	UploadType string
	ImgURL     string
	Length     string
	FileURL    string
	UseNums    string
	Isdel      string
	Addtime    string
	Updatetime string
	ClassifyID string
}{
	ID:         "id",
	Title:      "title",
	Author:     "author",
	Uploader:   "uploader",
	UploadType: "upload_type",
	ImgURL:     "img_url",
	Length:     "length",
	FileURL:    "file_url",
	UseNums:    "use_nums",
	Isdel:      "isdel",
	Addtime:    "addtime",
	Updatetime: "updatetime",
	ClassifyID: "classify_id",
}

// Generated where

var CMFMusicWhere = struct {
	ID         whereHelperint
	Title      whereHelperstring
	Author     whereHelperstring
	Uploader   whereHelperint
	UploadType whereHelperbool
	ImgURL     whereHelperstring
	Length     whereHelperstring
	FileURL    whereHelperstring
	UseNums    whereHelperint
	Isdel      whereHelperbool
	Addtime    whereHelperint
	Updatetime whereHelperint
	ClassifyID whereHelperint
}{
	ID:         whereHelperint{field: "`cmf_music`.`id`"},
	Title:      whereHelperstring{field: "`cmf_music`.`title`"},
	Author:     whereHelperstring{field: "`cmf_music`.`author`"},
	Uploader:   whereHelperint{field: "`cmf_music`.`uploader`"},
	UploadType: whereHelperbool{field: "`cmf_music`.`upload_type`"},
	ImgURL:     whereHelperstring{field: "`cmf_music`.`img_url`"},
	Length:     whereHelperstring{field: "`cmf_music`.`length`"},
	FileURL:    whereHelperstring{field: "`cmf_music`.`file_url`"},
	UseNums:    whereHelperint{field: "`cmf_music`.`use_nums`"},
	Isdel:      whereHelperbool{field: "`cmf_music`.`isdel`"},
	Addtime:    whereHelperint{field: "`cmf_music`.`addtime`"},
	Updatetime: whereHelperint{field: "`cmf_music`.`updatetime`"},
	ClassifyID: whereHelperint{field: "`cmf_music`.`classify_id`"},
}

// CMFMusicRels is where relationship names are stored.
var CMFMusicRels = struct {
}{}

// cmfMusicR is where relationships are stored.
type cmfMusicR struct {
}

// NewStruct creates a new relationship struct
func (*cmfMusicR) NewStruct() *cmfMusicR {
	return &cmfMusicR{}
}

// cmfMusicL is where Load methods for each relationship are stored.
type cmfMusicL struct{}

var (
	cmfMusicAllColumns            = []string{"id", "title", "author", "uploader", "upload_type", "img_url", "length", "file_url", "use_nums", "isdel", "addtime", "updatetime", "classify_id"}
	cmfMusicColumnsWithoutDefault = []string{"title", "author", "img_url", "length", "file_url"}
	cmfMusicColumnsWithDefault    = []string{"id", "uploader", "upload_type", "use_nums", "isdel", "addtime", "updatetime", "classify_id"}
	cmfMusicPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFMusicSlice is an alias for a slice of pointers to CMFMusic.
	// This should generally be used opposed to []CMFMusic.
	CMFMusicSlice []*CMFMusic
	// CMFMusicHook is the signature for custom CMFMusic hook methods
	CMFMusicHook func(context.Context, boil.ContextExecutor, *CMFMusic) error

	cmfMusicQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfMusicType                 = reflect.TypeOf(&CMFMusic{})
	cmfMusicMapping              = queries.MakeStructMapping(cmfMusicType)
	cmfMusicPrimaryKeyMapping, _ = queries.BindMapping(cmfMusicType, cmfMusicMapping, cmfMusicPrimaryKeyColumns)
	cmfMusicInsertCacheMut       sync.RWMutex
	cmfMusicInsertCache          = make(map[string]insertCache)
	cmfMusicUpdateCacheMut       sync.RWMutex
	cmfMusicUpdateCache          = make(map[string]updateCache)
	cmfMusicUpsertCacheMut       sync.RWMutex
	cmfMusicUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfMusicBeforeInsertHooks []CMFMusicHook
var cmfMusicBeforeUpdateHooks []CMFMusicHook
var cmfMusicBeforeDeleteHooks []CMFMusicHook
var cmfMusicBeforeUpsertHooks []CMFMusicHook

var cmfMusicAfterInsertHooks []CMFMusicHook
var cmfMusicAfterSelectHooks []CMFMusicHook
var cmfMusicAfterUpdateHooks []CMFMusicHook
var cmfMusicAfterDeleteHooks []CMFMusicHook
var cmfMusicAfterUpsertHooks []CMFMusicHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFMusic) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFMusic) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFMusic) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFMusic) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFMusic) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFMusic) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFMusic) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFMusic) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFMusic) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfMusicAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFMusicHook registers your hook function for all future operations.
func AddCMFMusicHook(hookPoint boil.HookPoint, cmfMusicHook CMFMusicHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfMusicBeforeInsertHooks = append(cmfMusicBeforeInsertHooks, cmfMusicHook)
	case boil.BeforeUpdateHook:
		cmfMusicBeforeUpdateHooks = append(cmfMusicBeforeUpdateHooks, cmfMusicHook)
	case boil.BeforeDeleteHook:
		cmfMusicBeforeDeleteHooks = append(cmfMusicBeforeDeleteHooks, cmfMusicHook)
	case boil.BeforeUpsertHook:
		cmfMusicBeforeUpsertHooks = append(cmfMusicBeforeUpsertHooks, cmfMusicHook)
	case boil.AfterInsertHook:
		cmfMusicAfterInsertHooks = append(cmfMusicAfterInsertHooks, cmfMusicHook)
	case boil.AfterSelectHook:
		cmfMusicAfterSelectHooks = append(cmfMusicAfterSelectHooks, cmfMusicHook)
	case boil.AfterUpdateHook:
		cmfMusicAfterUpdateHooks = append(cmfMusicAfterUpdateHooks, cmfMusicHook)
	case boil.AfterDeleteHook:
		cmfMusicAfterDeleteHooks = append(cmfMusicAfterDeleteHooks, cmfMusicHook)
	case boil.AfterUpsertHook:
		cmfMusicAfterUpsertHooks = append(cmfMusicAfterUpsertHooks, cmfMusicHook)
	}
}

// One returns a single cmfMusic record from the query.
func (q cmfMusicQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFMusic, error) {
	o := &CMFMusic{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_music")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFMusic records from the query.
func (q cmfMusicQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFMusicSlice, error) {
	var o []*CMFMusic

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFMusic slice")
	}

	if len(cmfMusicAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFMusic records in the query.
func (q cmfMusicQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_music rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfMusicQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_music exists")
	}

	return count > 0, nil
}

// CMFMusics retrieves all the records using an executor.
func CMFMusics(mods ...qm.QueryMod) cmfMusicQuery {
	mods = append(mods, qm.From("`cmf_music`"))
	return cmfMusicQuery{NewQuery(mods...)}
}

// FindCMFMusic retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFMusic(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFMusic, error) {
	cmfMusicObj := &CMFMusic{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_music` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfMusicObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_music")
	}

	return cmfMusicObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFMusic) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_music provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfMusicColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfMusicInsertCacheMut.RLock()
	cache, cached := cmfMusicInsertCache[key]
	cmfMusicInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfMusicAllColumns,
			cmfMusicColumnsWithDefault,
			cmfMusicColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfMusicType, cmfMusicMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfMusicType, cmfMusicMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_music` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_music` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_music` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfMusicPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_music")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfMusicMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_music")
	}

CacheNoHooks:
	if !cached {
		cmfMusicInsertCacheMut.Lock()
		cmfMusicInsertCache[key] = cache
		cmfMusicInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFMusic.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFMusic) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfMusicUpdateCacheMut.RLock()
	cache, cached := cmfMusicUpdateCache[key]
	cmfMusicUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfMusicAllColumns,
			cmfMusicPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_music, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_music` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfMusicPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfMusicType, cmfMusicMapping, append(wl, cmfMusicPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_music row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_music")
	}

	if !cached {
		cmfMusicUpdateCacheMut.Lock()
		cmfMusicUpdateCache[key] = cache
		cmfMusicUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfMusicQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_music")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_music")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFMusicSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfMusicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_music` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfMusicPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfMusic slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfMusic")
	}
	return rowsAff, nil
}

var mySQLCMFMusicUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFMusic) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_music provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfMusicColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFMusicUniqueColumns, o)

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

	cmfMusicUpsertCacheMut.RLock()
	cache, cached := cmfMusicUpsertCache[key]
	cmfMusicUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfMusicAllColumns,
			cmfMusicColumnsWithDefault,
			cmfMusicColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfMusicAllColumns,
			cmfMusicPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_music, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_music`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_music` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfMusicType, cmfMusicMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfMusicType, cmfMusicMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_music")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfMusicMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfMusicType, cmfMusicMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_music")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_music")
	}

CacheNoHooks:
	if !cached {
		cmfMusicUpsertCacheMut.Lock()
		cmfMusicUpsertCache[key] = cache
		cmfMusicUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFMusic record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFMusic) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFMusic provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfMusicPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_music` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_music")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_music")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfMusicQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfMusicQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_music")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_music")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFMusicSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfMusicBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfMusicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_music` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfMusicPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfMusic slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_music")
	}

	if len(cmfMusicAfterDeleteHooks) != 0 {
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
func (o *CMFMusic) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFMusic(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFMusicSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFMusicSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfMusicPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_music`.* FROM `cmf_music` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfMusicPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFMusicSlice")
	}

	*o = slice

	return nil
}

// CMFMusicExists checks if the CMFMusic row exists.
func CMFMusicExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_music` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_music exists")
	}

	return exists, nil
}
