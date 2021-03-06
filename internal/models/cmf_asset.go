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

// CMFAsset is an object representing the database table.
type CMFAsset struct {
	ID            uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID        uint64      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	FileSize      uint64      `boil:"file_size" json:"file_size" toml:"file_size" yaml:"file_size"`
	CreateTime    uint        `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	Status        uint8       `boil:"status" json:"status" toml:"status" yaml:"status"`
	DownloadTimes uint        `boil:"download_times" json:"download_times" toml:"download_times" yaml:"download_times"`
	FileKey       string      `boil:"file_key" json:"file_key" toml:"file_key" yaml:"file_key"`
	Filename      string      `boil:"filename" json:"filename" toml:"filename" yaml:"filename"`
	FilePath      string      `boil:"file_path" json:"file_path" toml:"file_path" yaml:"file_path"`
	FileMD5       string      `boil:"file_md5" json:"file_md5" toml:"file_md5" yaml:"file_md5"`
	FileSha1      string      `boil:"file_sha1" json:"file_sha1" toml:"file_sha1" yaml:"file_sha1"`
	Suffix        string      `boil:"suffix" json:"suffix" toml:"suffix" yaml:"suffix"`
	More          null.String `boil:"more" json:"more,omitempty" toml:"more" yaml:"more,omitempty"`

	R *cmfAssetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfAssetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFAssetColumns = struct {
	ID            string
	UserID        string
	FileSize      string
	CreateTime    string
	Status        string
	DownloadTimes string
	FileKey       string
	Filename      string
	FilePath      string
	FileMD5       string
	FileSha1      string
	Suffix        string
	More          string
}{
	ID:            "id",
	UserID:        "user_id",
	FileSize:      "file_size",
	CreateTime:    "create_time",
	Status:        "status",
	DownloadTimes: "download_times",
	FileKey:       "file_key",
	Filename:      "filename",
	FilePath:      "file_path",
	FileMD5:       "file_md5",
	FileSha1:      "file_sha1",
	Suffix:        "suffix",
	More:          "more",
}

// Generated where

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint64) NIN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var CMFAssetWhere = struct {
	ID            whereHelperuint64
	UserID        whereHelperuint64
	FileSize      whereHelperuint64
	CreateTime    whereHelperuint
	Status        whereHelperuint8
	DownloadTimes whereHelperuint
	FileKey       whereHelperstring
	Filename      whereHelperstring
	FilePath      whereHelperstring
	FileMD5       whereHelperstring
	FileSha1      whereHelperstring
	Suffix        whereHelperstring
	More          whereHelpernull_String
}{
	ID:            whereHelperuint64{field: "`cmf_asset`.`id`"},
	UserID:        whereHelperuint64{field: "`cmf_asset`.`user_id`"},
	FileSize:      whereHelperuint64{field: "`cmf_asset`.`file_size`"},
	CreateTime:    whereHelperuint{field: "`cmf_asset`.`create_time`"},
	Status:        whereHelperuint8{field: "`cmf_asset`.`status`"},
	DownloadTimes: whereHelperuint{field: "`cmf_asset`.`download_times`"},
	FileKey:       whereHelperstring{field: "`cmf_asset`.`file_key`"},
	Filename:      whereHelperstring{field: "`cmf_asset`.`filename`"},
	FilePath:      whereHelperstring{field: "`cmf_asset`.`file_path`"},
	FileMD5:       whereHelperstring{field: "`cmf_asset`.`file_md5`"},
	FileSha1:      whereHelperstring{field: "`cmf_asset`.`file_sha1`"},
	Suffix:        whereHelperstring{field: "`cmf_asset`.`suffix`"},
	More:          whereHelpernull_String{field: "`cmf_asset`.`more`"},
}

// CMFAssetRels is where relationship names are stored.
var CMFAssetRels = struct {
}{}

// cmfAssetR is where relationships are stored.
type cmfAssetR struct {
}

// NewStruct creates a new relationship struct
func (*cmfAssetR) NewStruct() *cmfAssetR {
	return &cmfAssetR{}
}

// cmfAssetL is where Load methods for each relationship are stored.
type cmfAssetL struct{}

var (
	cmfAssetAllColumns            = []string{"id", "user_id", "file_size", "create_time", "status", "download_times", "file_key", "filename", "file_path", "file_md5", "file_sha1", "suffix", "more"}
	cmfAssetColumnsWithoutDefault = []string{"file_key", "filename", "file_path", "file_md5", "file_sha1", "suffix", "more"}
	cmfAssetColumnsWithDefault    = []string{"id", "user_id", "file_size", "create_time", "status", "download_times"}
	cmfAssetPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFAssetSlice is an alias for a slice of pointers to CMFAsset.
	// This should generally be used opposed to []CMFAsset.
	CMFAssetSlice []*CMFAsset
	// CMFAssetHook is the signature for custom CMFAsset hook methods
	CMFAssetHook func(context.Context, boil.ContextExecutor, *CMFAsset) error

	cmfAssetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfAssetType                 = reflect.TypeOf(&CMFAsset{})
	cmfAssetMapping              = queries.MakeStructMapping(cmfAssetType)
	cmfAssetPrimaryKeyMapping, _ = queries.BindMapping(cmfAssetType, cmfAssetMapping, cmfAssetPrimaryKeyColumns)
	cmfAssetInsertCacheMut       sync.RWMutex
	cmfAssetInsertCache          = make(map[string]insertCache)
	cmfAssetUpdateCacheMut       sync.RWMutex
	cmfAssetUpdateCache          = make(map[string]updateCache)
	cmfAssetUpsertCacheMut       sync.RWMutex
	cmfAssetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfAssetBeforeInsertHooks []CMFAssetHook
var cmfAssetBeforeUpdateHooks []CMFAssetHook
var cmfAssetBeforeDeleteHooks []CMFAssetHook
var cmfAssetBeforeUpsertHooks []CMFAssetHook

var cmfAssetAfterInsertHooks []CMFAssetHook
var cmfAssetAfterSelectHooks []CMFAssetHook
var cmfAssetAfterUpdateHooks []CMFAssetHook
var cmfAssetAfterDeleteHooks []CMFAssetHook
var cmfAssetAfterUpsertHooks []CMFAssetHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFAsset) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFAsset) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFAsset) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFAsset) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFAsset) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFAsset) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFAsset) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFAsset) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFAsset) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAssetAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFAssetHook registers your hook function for all future operations.
func AddCMFAssetHook(hookPoint boil.HookPoint, cmfAssetHook CMFAssetHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfAssetBeforeInsertHooks = append(cmfAssetBeforeInsertHooks, cmfAssetHook)
	case boil.BeforeUpdateHook:
		cmfAssetBeforeUpdateHooks = append(cmfAssetBeforeUpdateHooks, cmfAssetHook)
	case boil.BeforeDeleteHook:
		cmfAssetBeforeDeleteHooks = append(cmfAssetBeforeDeleteHooks, cmfAssetHook)
	case boil.BeforeUpsertHook:
		cmfAssetBeforeUpsertHooks = append(cmfAssetBeforeUpsertHooks, cmfAssetHook)
	case boil.AfterInsertHook:
		cmfAssetAfterInsertHooks = append(cmfAssetAfterInsertHooks, cmfAssetHook)
	case boil.AfterSelectHook:
		cmfAssetAfterSelectHooks = append(cmfAssetAfterSelectHooks, cmfAssetHook)
	case boil.AfterUpdateHook:
		cmfAssetAfterUpdateHooks = append(cmfAssetAfterUpdateHooks, cmfAssetHook)
	case boil.AfterDeleteHook:
		cmfAssetAfterDeleteHooks = append(cmfAssetAfterDeleteHooks, cmfAssetHook)
	case boil.AfterUpsertHook:
		cmfAssetAfterUpsertHooks = append(cmfAssetAfterUpsertHooks, cmfAssetHook)
	}
}

// One returns a single cmfAsset record from the query.
func (q cmfAssetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFAsset, error) {
	o := &CMFAsset{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_asset")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFAsset records from the query.
func (q cmfAssetQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFAssetSlice, error) {
	var o []*CMFAsset

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFAsset slice")
	}

	if len(cmfAssetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFAsset records in the query.
func (q cmfAssetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_asset rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfAssetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_asset exists")
	}

	return count > 0, nil
}

// CMFAssets retrieves all the records using an executor.
func CMFAssets(mods ...qm.QueryMod) cmfAssetQuery {
	mods = append(mods, qm.From("`cmf_asset`"))
	return cmfAssetQuery{NewQuery(mods...)}
}

// FindCMFAsset retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFAsset(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*CMFAsset, error) {
	cmfAssetObj := &CMFAsset{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_asset` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfAssetObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_asset")
	}

	return cmfAssetObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFAsset) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_asset provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfAssetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfAssetInsertCacheMut.RLock()
	cache, cached := cmfAssetInsertCache[key]
	cmfAssetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfAssetAllColumns,
			cmfAssetColumnsWithDefault,
			cmfAssetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfAssetType, cmfAssetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfAssetType, cmfAssetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_asset` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_asset` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_asset` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfAssetPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_asset")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfAssetMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_asset")
	}

CacheNoHooks:
	if !cached {
		cmfAssetInsertCacheMut.Lock()
		cmfAssetInsertCache[key] = cache
		cmfAssetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFAsset.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFAsset) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfAssetUpdateCacheMut.RLock()
	cache, cached := cmfAssetUpdateCache[key]
	cmfAssetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfAssetAllColumns,
			cmfAssetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_asset, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_asset` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfAssetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfAssetType, cmfAssetMapping, append(wl, cmfAssetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_asset row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_asset")
	}

	if !cached {
		cmfAssetUpdateCacheMut.Lock()
		cmfAssetUpdateCache[key] = cache
		cmfAssetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfAssetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_asset")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFAssetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAssetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_asset` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAssetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfAsset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfAsset")
	}
	return rowsAff, nil
}

var mySQLCMFAssetUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFAsset) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_asset provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfAssetColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFAssetUniqueColumns, o)

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

	cmfAssetUpsertCacheMut.RLock()
	cache, cached := cmfAssetUpsertCache[key]
	cmfAssetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfAssetAllColumns,
			cmfAssetColumnsWithDefault,
			cmfAssetColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfAssetAllColumns,
			cmfAssetPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_asset, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_asset`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_asset` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfAssetType, cmfAssetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfAssetType, cmfAssetMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_asset")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfAssetMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfAssetType, cmfAssetMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_asset")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_asset")
	}

CacheNoHooks:
	if !cached {
		cmfAssetUpsertCacheMut.Lock()
		cmfAssetUpsertCache[key] = cache
		cmfAssetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFAsset record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFAsset) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFAsset provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfAssetPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_asset` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_asset")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfAssetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfAssetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_asset")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFAssetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfAssetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAssetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_asset` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAssetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfAsset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_asset")
	}

	if len(cmfAssetAfterDeleteHooks) != 0 {
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
func (o *CMFAsset) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFAsset(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFAssetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFAssetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAssetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_asset`.* FROM `cmf_asset` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAssetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFAssetSlice")
	}

	*o = slice

	return nil
}

// CMFAssetExists checks if the CMFAsset row exists.
func CMFAssetExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_asset` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_asset exists")
	}

	return exists, nil
}
