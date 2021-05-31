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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// CMFPostFormRecord is an object representing the database table.
type CMFPostFormRecord struct {
	ID            int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	FormID        int64             `boil:"form_id" json:"form_id" toml:"form_id" yaml:"form_id"`
	FieldName     string            `boil:"field_name" json:"field_name" toml:"field_name" yaml:"field_name"`
	FieldValue    null.String       `boil:"field_value" json:"field_value,omitempty" toml:"field_value" yaml:"field_value,omitempty"`
	FieldLenValue null.String       `boil:"field_len_value" json:"field_len_value,omitempty" toml:"field_len_value" yaml:"field_len_value,omitempty"`
	FieldInt      null.Int64        `boil:"field_int" json:"field_int,omitempty" toml:"field_int" yaml:"field_int,omitempty"`
	PostID        int64             `boil:"post_id" json:"post_id" toml:"post_id" yaml:"post_id"`
	FieldDecimals types.NullDecimal `boil:"field_decimals" json:"field_decimals,omitempty" toml:"field_decimals" yaml:"field_decimals,omitempty"`
	CreateTime    time.Time         `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`

	R *cmfPostFormRecordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfPostFormRecordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFPostFormRecordColumns = struct {
	ID            string
	FormID        string
	FieldName     string
	FieldValue    string
	FieldLenValue string
	FieldInt      string
	PostID        string
	FieldDecimals string
	CreateTime    string
}{
	ID:            "id",
	FormID:        "form_id",
	FieldName:     "field_name",
	FieldValue:    "field_value",
	FieldLenValue: "field_len_value",
	FieldInt:      "field_int",
	PostID:        "post_id",
	FieldDecimals: "field_decimals",
	CreateTime:    "create_time",
}

// Generated where

var CMFPostFormRecordWhere = struct {
	ID            whereHelperint64
	FormID        whereHelperint64
	FieldName     whereHelperstring
	FieldValue    whereHelpernull_String
	FieldLenValue whereHelpernull_String
	FieldInt      whereHelpernull_Int64
	PostID        whereHelperint64
	FieldDecimals whereHelpertypes_NullDecimal
	CreateTime    whereHelpertime_Time
}{
	ID:            whereHelperint64{field: "`cmf_post_form_records`.`id`"},
	FormID:        whereHelperint64{field: "`cmf_post_form_records`.`form_id`"},
	FieldName:     whereHelperstring{field: "`cmf_post_form_records`.`field_name`"},
	FieldValue:    whereHelpernull_String{field: "`cmf_post_form_records`.`field_value`"},
	FieldLenValue: whereHelpernull_String{field: "`cmf_post_form_records`.`field_len_value`"},
	FieldInt:      whereHelpernull_Int64{field: "`cmf_post_form_records`.`field_int`"},
	PostID:        whereHelperint64{field: "`cmf_post_form_records`.`post_id`"},
	FieldDecimals: whereHelpertypes_NullDecimal{field: "`cmf_post_form_records`.`field_decimals`"},
	CreateTime:    whereHelpertime_Time{field: "`cmf_post_form_records`.`create_time`"},
}

// CMFPostFormRecordRels is where relationship names are stored.
var CMFPostFormRecordRels = struct {
}{}

// cmfPostFormRecordR is where relationships are stored.
type cmfPostFormRecordR struct {
}

// NewStruct creates a new relationship struct
func (*cmfPostFormRecordR) NewStruct() *cmfPostFormRecordR {
	return &cmfPostFormRecordR{}
}

// cmfPostFormRecordL is where Load methods for each relationship are stored.
type cmfPostFormRecordL struct{}

var (
	cmfPostFormRecordAllColumns            = []string{"id", "form_id", "field_name", "field_value", "field_len_value", "field_int", "post_id", "field_decimals", "create_time"}
	cmfPostFormRecordColumnsWithoutDefault = []string{"form_id", "field_name", "field_value", "field_len_value", "field_int", "post_id", "field_decimals", "create_time"}
	cmfPostFormRecordColumnsWithDefault    = []string{"id"}
	cmfPostFormRecordPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFPostFormRecordSlice is an alias for a slice of pointers to CMFPostFormRecord.
	// This should generally be used opposed to []CMFPostFormRecord.
	CMFPostFormRecordSlice []*CMFPostFormRecord
	// CMFPostFormRecordHook is the signature for custom CMFPostFormRecord hook methods
	CMFPostFormRecordHook func(context.Context, boil.ContextExecutor, *CMFPostFormRecord) error

	cmfPostFormRecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfPostFormRecordType                 = reflect.TypeOf(&CMFPostFormRecord{})
	cmfPostFormRecordMapping              = queries.MakeStructMapping(cmfPostFormRecordType)
	cmfPostFormRecordPrimaryKeyMapping, _ = queries.BindMapping(cmfPostFormRecordType, cmfPostFormRecordMapping, cmfPostFormRecordPrimaryKeyColumns)
	cmfPostFormRecordInsertCacheMut       sync.RWMutex
	cmfPostFormRecordInsertCache          = make(map[string]insertCache)
	cmfPostFormRecordUpdateCacheMut       sync.RWMutex
	cmfPostFormRecordUpdateCache          = make(map[string]updateCache)
	cmfPostFormRecordUpsertCacheMut       sync.RWMutex
	cmfPostFormRecordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfPostFormRecordBeforeInsertHooks []CMFPostFormRecordHook
var cmfPostFormRecordBeforeUpdateHooks []CMFPostFormRecordHook
var cmfPostFormRecordBeforeDeleteHooks []CMFPostFormRecordHook
var cmfPostFormRecordBeforeUpsertHooks []CMFPostFormRecordHook

var cmfPostFormRecordAfterInsertHooks []CMFPostFormRecordHook
var cmfPostFormRecordAfterSelectHooks []CMFPostFormRecordHook
var cmfPostFormRecordAfterUpdateHooks []CMFPostFormRecordHook
var cmfPostFormRecordAfterDeleteHooks []CMFPostFormRecordHook
var cmfPostFormRecordAfterUpsertHooks []CMFPostFormRecordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFPostFormRecord) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFPostFormRecord) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFPostFormRecord) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFPostFormRecord) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFPostFormRecord) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFPostFormRecord) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFPostFormRecord) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFPostFormRecord) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFPostFormRecord) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormRecordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFPostFormRecordHook registers your hook function for all future operations.
func AddCMFPostFormRecordHook(hookPoint boil.HookPoint, cmfPostFormRecordHook CMFPostFormRecordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfPostFormRecordBeforeInsertHooks = append(cmfPostFormRecordBeforeInsertHooks, cmfPostFormRecordHook)
	case boil.BeforeUpdateHook:
		cmfPostFormRecordBeforeUpdateHooks = append(cmfPostFormRecordBeforeUpdateHooks, cmfPostFormRecordHook)
	case boil.BeforeDeleteHook:
		cmfPostFormRecordBeforeDeleteHooks = append(cmfPostFormRecordBeforeDeleteHooks, cmfPostFormRecordHook)
	case boil.BeforeUpsertHook:
		cmfPostFormRecordBeforeUpsertHooks = append(cmfPostFormRecordBeforeUpsertHooks, cmfPostFormRecordHook)
	case boil.AfterInsertHook:
		cmfPostFormRecordAfterInsertHooks = append(cmfPostFormRecordAfterInsertHooks, cmfPostFormRecordHook)
	case boil.AfterSelectHook:
		cmfPostFormRecordAfterSelectHooks = append(cmfPostFormRecordAfterSelectHooks, cmfPostFormRecordHook)
	case boil.AfterUpdateHook:
		cmfPostFormRecordAfterUpdateHooks = append(cmfPostFormRecordAfterUpdateHooks, cmfPostFormRecordHook)
	case boil.AfterDeleteHook:
		cmfPostFormRecordAfterDeleteHooks = append(cmfPostFormRecordAfterDeleteHooks, cmfPostFormRecordHook)
	case boil.AfterUpsertHook:
		cmfPostFormRecordAfterUpsertHooks = append(cmfPostFormRecordAfterUpsertHooks, cmfPostFormRecordHook)
	}
}

// One returns a single cmfPostFormRecord record from the query.
func (q cmfPostFormRecordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFPostFormRecord, error) {
	o := &CMFPostFormRecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_post_form_records")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFPostFormRecord records from the query.
func (q cmfPostFormRecordQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFPostFormRecordSlice, error) {
	var o []*CMFPostFormRecord

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFPostFormRecord slice")
	}

	if len(cmfPostFormRecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFPostFormRecord records in the query.
func (q cmfPostFormRecordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_post_form_records rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfPostFormRecordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_post_form_records exists")
	}

	return count > 0, nil
}

// CMFPostFormRecords retrieves all the records using an executor.
func CMFPostFormRecords(mods ...qm.QueryMod) cmfPostFormRecordQuery {
	mods = append(mods, qm.From("`cmf_post_form_records`"))
	return cmfPostFormRecordQuery{NewQuery(mods...)}
}

// FindCMFPostFormRecord retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFPostFormRecord(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CMFPostFormRecord, error) {
	cmfPostFormRecordObj := &CMFPostFormRecord{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_post_form_records` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfPostFormRecordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_post_form_records")
	}

	return cmfPostFormRecordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFPostFormRecord) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_post_form_records provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPostFormRecordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfPostFormRecordInsertCacheMut.RLock()
	cache, cached := cmfPostFormRecordInsertCache[key]
	cmfPostFormRecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfPostFormRecordAllColumns,
			cmfPostFormRecordColumnsWithDefault,
			cmfPostFormRecordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfPostFormRecordType, cmfPostFormRecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfPostFormRecordType, cmfPostFormRecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_post_form_records` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_post_form_records` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_post_form_records` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfPostFormRecordPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_post_form_records")
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

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPostFormRecordMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_post_form_records")
	}

CacheNoHooks:
	if !cached {
		cmfPostFormRecordInsertCacheMut.Lock()
		cmfPostFormRecordInsertCache[key] = cache
		cmfPostFormRecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFPostFormRecord.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFPostFormRecord) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfPostFormRecordUpdateCacheMut.RLock()
	cache, cached := cmfPostFormRecordUpdateCache[key]
	cmfPostFormRecordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfPostFormRecordAllColumns,
			cmfPostFormRecordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_post_form_records, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_post_form_records` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfPostFormRecordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfPostFormRecordType, cmfPostFormRecordMapping, append(wl, cmfPostFormRecordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_post_form_records row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_post_form_records")
	}

	if !cached {
		cmfPostFormRecordUpdateCacheMut.Lock()
		cmfPostFormRecordUpdateCache[key] = cache
		cmfPostFormRecordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfPostFormRecordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_post_form_records")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_post_form_records")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFPostFormRecordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPostFormRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_post_form_records` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPostFormRecordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfPostFormRecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfPostFormRecord")
	}
	return rowsAff, nil
}

var mySQLCMFPostFormRecordUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFPostFormRecord) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_post_form_records provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPostFormRecordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFPostFormRecordUniqueColumns, o)

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

	cmfPostFormRecordUpsertCacheMut.RLock()
	cache, cached := cmfPostFormRecordUpsertCache[key]
	cmfPostFormRecordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfPostFormRecordAllColumns,
			cmfPostFormRecordColumnsWithDefault,
			cmfPostFormRecordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfPostFormRecordAllColumns,
			cmfPostFormRecordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_post_form_records, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_post_form_records`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_post_form_records` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfPostFormRecordType, cmfPostFormRecordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfPostFormRecordType, cmfPostFormRecordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_post_form_records")
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

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPostFormRecordMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfPostFormRecordType, cmfPostFormRecordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_post_form_records")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_post_form_records")
	}

CacheNoHooks:
	if !cached {
		cmfPostFormRecordUpsertCacheMut.Lock()
		cmfPostFormRecordUpsertCache[key] = cache
		cmfPostFormRecordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFPostFormRecord record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFPostFormRecord) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFPostFormRecord provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfPostFormRecordPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_post_form_records` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_post_form_records")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_post_form_records")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfPostFormRecordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfPostFormRecordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_post_form_records")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_post_form_records")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFPostFormRecordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfPostFormRecordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPostFormRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_post_form_records` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPostFormRecordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfPostFormRecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_post_form_records")
	}

	if len(cmfPostFormRecordAfterDeleteHooks) != 0 {
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
func (o *CMFPostFormRecord) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFPostFormRecord(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFPostFormRecordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFPostFormRecordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPostFormRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_post_form_records`.* FROM `cmf_post_form_records` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPostFormRecordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFPostFormRecordSlice")
	}

	*o = slice

	return nil
}

// CMFPostFormRecordExists checks if the CMFPostFormRecord row exists.
func CMFPostFormRecordExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_post_form_records` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_post_form_records exists")
	}

	return exists, nil
}