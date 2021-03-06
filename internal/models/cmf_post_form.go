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

// CMFPostForm is an object representing the database table.
type CMFPostForm struct {
	ID         int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	FormName   string      `boil:"form_name" json:"form_name" toml:"form_name" yaml:"form_name"`
	Source     string      `boil:"source" json:"source" toml:"source" yaml:"source"`
	FieldsJSON string      `boil:"fields_json" json:"fields_json" toml:"fields_json" yaml:"fields_json"`
	Comment    null.String `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`
	CreateTime time.Time   `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime null.Time   `boil:"update_time" json:"update_time,omitempty" toml:"update_time" yaml:"update_time,omitempty"`

	R *cmfPostFormR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfPostFormL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFPostFormColumns = struct {
	ID         string
	FormName   string
	Source     string
	FieldsJSON string
	Comment    string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	FormName:   "form_name",
	Source:     "source",
	FieldsJSON: "fields_json",
	Comment:    "comment",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// Generated where

var CMFPostFormWhere = struct {
	ID         whereHelperint64
	FormName   whereHelperstring
	Source     whereHelperstring
	FieldsJSON whereHelperstring
	Comment    whereHelpernull_String
	CreateTime whereHelpertime_Time
	UpdateTime whereHelpernull_Time
}{
	ID:         whereHelperint64{field: "`cmf_post_form`.`id`"},
	FormName:   whereHelperstring{field: "`cmf_post_form`.`form_name`"},
	Source:     whereHelperstring{field: "`cmf_post_form`.`source`"},
	FieldsJSON: whereHelperstring{field: "`cmf_post_form`.`fields_json`"},
	Comment:    whereHelpernull_String{field: "`cmf_post_form`.`comment`"},
	CreateTime: whereHelpertime_Time{field: "`cmf_post_form`.`create_time`"},
	UpdateTime: whereHelpernull_Time{field: "`cmf_post_form`.`update_time`"},
}

// CMFPostFormRels is where relationship names are stored.
var CMFPostFormRels = struct {
}{}

// cmfPostFormR is where relationships are stored.
type cmfPostFormR struct {
}

// NewStruct creates a new relationship struct
func (*cmfPostFormR) NewStruct() *cmfPostFormR {
	return &cmfPostFormR{}
}

// cmfPostFormL is where Load methods for each relationship are stored.
type cmfPostFormL struct{}

var (
	cmfPostFormAllColumns            = []string{"id", "form_name", "source", "fields_json", "comment", "create_time", "update_time"}
	cmfPostFormColumnsWithoutDefault = []string{"form_name", "fields_json", "comment", "create_time", "update_time"}
	cmfPostFormColumnsWithDefault    = []string{"id", "source"}
	cmfPostFormPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFPostFormSlice is an alias for a slice of pointers to CMFPostForm.
	// This should generally be used opposed to []CMFPostForm.
	CMFPostFormSlice []*CMFPostForm
	// CMFPostFormHook is the signature for custom CMFPostForm hook methods
	CMFPostFormHook func(context.Context, boil.ContextExecutor, *CMFPostForm) error

	cmfPostFormQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfPostFormType                 = reflect.TypeOf(&CMFPostForm{})
	cmfPostFormMapping              = queries.MakeStructMapping(cmfPostFormType)
	cmfPostFormPrimaryKeyMapping, _ = queries.BindMapping(cmfPostFormType, cmfPostFormMapping, cmfPostFormPrimaryKeyColumns)
	cmfPostFormInsertCacheMut       sync.RWMutex
	cmfPostFormInsertCache          = make(map[string]insertCache)
	cmfPostFormUpdateCacheMut       sync.RWMutex
	cmfPostFormUpdateCache          = make(map[string]updateCache)
	cmfPostFormUpsertCacheMut       sync.RWMutex
	cmfPostFormUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfPostFormBeforeInsertHooks []CMFPostFormHook
var cmfPostFormBeforeUpdateHooks []CMFPostFormHook
var cmfPostFormBeforeDeleteHooks []CMFPostFormHook
var cmfPostFormBeforeUpsertHooks []CMFPostFormHook

var cmfPostFormAfterInsertHooks []CMFPostFormHook
var cmfPostFormAfterSelectHooks []CMFPostFormHook
var cmfPostFormAfterUpdateHooks []CMFPostFormHook
var cmfPostFormAfterDeleteHooks []CMFPostFormHook
var cmfPostFormAfterUpsertHooks []CMFPostFormHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFPostForm) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFPostForm) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFPostForm) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFPostForm) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFPostForm) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFPostForm) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFPostForm) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFPostForm) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFPostForm) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPostFormAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFPostFormHook registers your hook function for all future operations.
func AddCMFPostFormHook(hookPoint boil.HookPoint, cmfPostFormHook CMFPostFormHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfPostFormBeforeInsertHooks = append(cmfPostFormBeforeInsertHooks, cmfPostFormHook)
	case boil.BeforeUpdateHook:
		cmfPostFormBeforeUpdateHooks = append(cmfPostFormBeforeUpdateHooks, cmfPostFormHook)
	case boil.BeforeDeleteHook:
		cmfPostFormBeforeDeleteHooks = append(cmfPostFormBeforeDeleteHooks, cmfPostFormHook)
	case boil.BeforeUpsertHook:
		cmfPostFormBeforeUpsertHooks = append(cmfPostFormBeforeUpsertHooks, cmfPostFormHook)
	case boil.AfterInsertHook:
		cmfPostFormAfterInsertHooks = append(cmfPostFormAfterInsertHooks, cmfPostFormHook)
	case boil.AfterSelectHook:
		cmfPostFormAfterSelectHooks = append(cmfPostFormAfterSelectHooks, cmfPostFormHook)
	case boil.AfterUpdateHook:
		cmfPostFormAfterUpdateHooks = append(cmfPostFormAfterUpdateHooks, cmfPostFormHook)
	case boil.AfterDeleteHook:
		cmfPostFormAfterDeleteHooks = append(cmfPostFormAfterDeleteHooks, cmfPostFormHook)
	case boil.AfterUpsertHook:
		cmfPostFormAfterUpsertHooks = append(cmfPostFormAfterUpsertHooks, cmfPostFormHook)
	}
}

// One returns a single cmfPostForm record from the query.
func (q cmfPostFormQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFPostForm, error) {
	o := &CMFPostForm{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_post_form")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFPostForm records from the query.
func (q cmfPostFormQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFPostFormSlice, error) {
	var o []*CMFPostForm

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFPostForm slice")
	}

	if len(cmfPostFormAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFPostForm records in the query.
func (q cmfPostFormQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_post_form rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfPostFormQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_post_form exists")
	}

	return count > 0, nil
}

// CMFPostForms retrieves all the records using an executor.
func CMFPostForms(mods ...qm.QueryMod) cmfPostFormQuery {
	mods = append(mods, qm.From("`cmf_post_form`"))
	return cmfPostFormQuery{NewQuery(mods...)}
}

// FindCMFPostForm retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFPostForm(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CMFPostForm, error) {
	cmfPostFormObj := &CMFPostForm{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_post_form` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfPostFormObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_post_form")
	}

	return cmfPostFormObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFPostForm) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_post_form provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPostFormColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfPostFormInsertCacheMut.RLock()
	cache, cached := cmfPostFormInsertCache[key]
	cmfPostFormInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfPostFormAllColumns,
			cmfPostFormColumnsWithDefault,
			cmfPostFormColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfPostFormType, cmfPostFormMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfPostFormType, cmfPostFormMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_post_form` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_post_form` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_post_form` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfPostFormPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_post_form")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPostFormMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_post_form")
	}

CacheNoHooks:
	if !cached {
		cmfPostFormInsertCacheMut.Lock()
		cmfPostFormInsertCache[key] = cache
		cmfPostFormInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFPostForm.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFPostForm) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfPostFormUpdateCacheMut.RLock()
	cache, cached := cmfPostFormUpdateCache[key]
	cmfPostFormUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfPostFormAllColumns,
			cmfPostFormPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_post_form, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_post_form` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfPostFormPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfPostFormType, cmfPostFormMapping, append(wl, cmfPostFormPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_post_form row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_post_form")
	}

	if !cached {
		cmfPostFormUpdateCacheMut.Lock()
		cmfPostFormUpdateCache[key] = cache
		cmfPostFormUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfPostFormQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_post_form")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_post_form")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFPostFormSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPostFormPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_post_form` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPostFormPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfPostForm slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfPostForm")
	}
	return rowsAff, nil
}

var mySQLCMFPostFormUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFPostForm) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_post_form provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPostFormColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFPostFormUniqueColumns, o)

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

	cmfPostFormUpsertCacheMut.RLock()
	cache, cached := cmfPostFormUpsertCache[key]
	cmfPostFormUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfPostFormAllColumns,
			cmfPostFormColumnsWithDefault,
			cmfPostFormColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfPostFormAllColumns,
			cmfPostFormPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_post_form, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_post_form`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_post_form` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfPostFormType, cmfPostFormMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfPostFormType, cmfPostFormMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_post_form")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPostFormMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfPostFormType, cmfPostFormMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_post_form")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_post_form")
	}

CacheNoHooks:
	if !cached {
		cmfPostFormUpsertCacheMut.Lock()
		cmfPostFormUpsertCache[key] = cache
		cmfPostFormUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFPostForm record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFPostForm) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFPostForm provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfPostFormPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_post_form` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_post_form")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_post_form")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfPostFormQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfPostFormQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_post_form")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_post_form")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFPostFormSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfPostFormBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPostFormPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_post_form` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPostFormPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfPostForm slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_post_form")
	}

	if len(cmfPostFormAfterDeleteHooks) != 0 {
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
func (o *CMFPostForm) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFPostForm(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFPostFormSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFPostFormSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPostFormPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_post_form`.* FROM `cmf_post_form` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPostFormPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFPostFormSlice")
	}

	*o = slice

	return nil
}

// CMFPostFormExists checks if the CMFPostForm row exists.
func CMFPostFormExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_post_form` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_post_form exists")
	}

	return exists, nil
}
