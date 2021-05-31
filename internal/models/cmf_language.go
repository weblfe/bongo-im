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

// CMFLanguage is an object representing the database table.
type CMFLanguage struct {
	ID            int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	FromLang      null.String `boil:"from_lang" json:"from_lang,omitempty" toml:"from_lang" yaml:"from_lang,omitempty"`
	TargetLang    null.Time   `boil:"target_lang" json:"target_lang,omitempty" toml:"target_lang" yaml:"target_lang,omitempty"`
	Tag           null.String `boil:"tag" json:"tag,omitempty" toml:"tag" yaml:"tag,omitempty"`
	TargetContent null.String `boil:"target_content" json:"target_content,omitempty" toml:"target_content" yaml:"target_content,omitempty"`
	FromContent   null.String `boil:"from_content" json:"from_content,omitempty" toml:"from_content" yaml:"from_content,omitempty"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *cmfLanguageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfLanguageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFLanguageColumns = struct {
	ID            string
	FromLang      string
	TargetLang    string
	Tag           string
	TargetContent string
	FromContent   string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	FromLang:      "from_lang",
	TargetLang:    "target_lang",
	Tag:           "tag",
	TargetContent: "target_content",
	FromContent:   "from_content",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// Generated where

var CMFLanguageWhere = struct {
	ID            whereHelperint64
	FromLang      whereHelpernull_String
	TargetLang    whereHelpernull_Time
	Tag           whereHelpernull_String
	TargetContent whereHelpernull_String
	FromContent   whereHelpernull_String
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
}{
	ID:            whereHelperint64{field: "`cmf_language`.`id`"},
	FromLang:      whereHelpernull_String{field: "`cmf_language`.`from_lang`"},
	TargetLang:    whereHelpernull_Time{field: "`cmf_language`.`target_lang`"},
	Tag:           whereHelpernull_String{field: "`cmf_language`.`tag`"},
	TargetContent: whereHelpernull_String{field: "`cmf_language`.`target_content`"},
	FromContent:   whereHelpernull_String{field: "`cmf_language`.`from_content`"},
	CreatedAt:     whereHelpertime_Time{field: "`cmf_language`.`created_at`"},
	UpdatedAt:     whereHelpertime_Time{field: "`cmf_language`.`updated_at`"},
}

// CMFLanguageRels is where relationship names are stored.
var CMFLanguageRels = struct {
}{}

// cmfLanguageR is where relationships are stored.
type cmfLanguageR struct {
}

// NewStruct creates a new relationship struct
func (*cmfLanguageR) NewStruct() *cmfLanguageR {
	return &cmfLanguageR{}
}

// cmfLanguageL is where Load methods for each relationship are stored.
type cmfLanguageL struct{}

var (
	cmfLanguageAllColumns            = []string{"id", "from_lang", "target_lang", "tag", "target_content", "from_content", "created_at", "updated_at"}
	cmfLanguageColumnsWithoutDefault = []string{"from_lang", "target_lang", "tag", "target_content", "from_content", "created_at", "updated_at"}
	cmfLanguageColumnsWithDefault    = []string{"id"}
	cmfLanguagePrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFLanguageSlice is an alias for a slice of pointers to CMFLanguage.
	// This should generally be used opposed to []CMFLanguage.
	CMFLanguageSlice []*CMFLanguage
	// CMFLanguageHook is the signature for custom CMFLanguage hook methods
	CMFLanguageHook func(context.Context, boil.ContextExecutor, *CMFLanguage) error

	cmfLanguageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfLanguageType                 = reflect.TypeOf(&CMFLanguage{})
	cmfLanguageMapping              = queries.MakeStructMapping(cmfLanguageType)
	cmfLanguagePrimaryKeyMapping, _ = queries.BindMapping(cmfLanguageType, cmfLanguageMapping, cmfLanguagePrimaryKeyColumns)
	cmfLanguageInsertCacheMut       sync.RWMutex
	cmfLanguageInsertCache          = make(map[string]insertCache)
	cmfLanguageUpdateCacheMut       sync.RWMutex
	cmfLanguageUpdateCache          = make(map[string]updateCache)
	cmfLanguageUpsertCacheMut       sync.RWMutex
	cmfLanguageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfLanguageBeforeInsertHooks []CMFLanguageHook
var cmfLanguageBeforeUpdateHooks []CMFLanguageHook
var cmfLanguageBeforeDeleteHooks []CMFLanguageHook
var cmfLanguageBeforeUpsertHooks []CMFLanguageHook

var cmfLanguageAfterInsertHooks []CMFLanguageHook
var cmfLanguageAfterSelectHooks []CMFLanguageHook
var cmfLanguageAfterUpdateHooks []CMFLanguageHook
var cmfLanguageAfterDeleteHooks []CMFLanguageHook
var cmfLanguageAfterUpsertHooks []CMFLanguageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFLanguage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFLanguage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFLanguage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFLanguage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFLanguage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFLanguage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFLanguage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFLanguage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFLanguage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLanguageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFLanguageHook registers your hook function for all future operations.
func AddCMFLanguageHook(hookPoint boil.HookPoint, cmfLanguageHook CMFLanguageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfLanguageBeforeInsertHooks = append(cmfLanguageBeforeInsertHooks, cmfLanguageHook)
	case boil.BeforeUpdateHook:
		cmfLanguageBeforeUpdateHooks = append(cmfLanguageBeforeUpdateHooks, cmfLanguageHook)
	case boil.BeforeDeleteHook:
		cmfLanguageBeforeDeleteHooks = append(cmfLanguageBeforeDeleteHooks, cmfLanguageHook)
	case boil.BeforeUpsertHook:
		cmfLanguageBeforeUpsertHooks = append(cmfLanguageBeforeUpsertHooks, cmfLanguageHook)
	case boil.AfterInsertHook:
		cmfLanguageAfterInsertHooks = append(cmfLanguageAfterInsertHooks, cmfLanguageHook)
	case boil.AfterSelectHook:
		cmfLanguageAfterSelectHooks = append(cmfLanguageAfterSelectHooks, cmfLanguageHook)
	case boil.AfterUpdateHook:
		cmfLanguageAfterUpdateHooks = append(cmfLanguageAfterUpdateHooks, cmfLanguageHook)
	case boil.AfterDeleteHook:
		cmfLanguageAfterDeleteHooks = append(cmfLanguageAfterDeleteHooks, cmfLanguageHook)
	case boil.AfterUpsertHook:
		cmfLanguageAfterUpsertHooks = append(cmfLanguageAfterUpsertHooks, cmfLanguageHook)
	}
}

// One returns a single cmfLanguage record from the query.
func (q cmfLanguageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFLanguage, error) {
	o := &CMFLanguage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_language")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFLanguage records from the query.
func (q cmfLanguageQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFLanguageSlice, error) {
	var o []*CMFLanguage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFLanguage slice")
	}

	if len(cmfLanguageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFLanguage records in the query.
func (q cmfLanguageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_language rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfLanguageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_language exists")
	}

	return count > 0, nil
}

// CMFLanguages retrieves all the records using an executor.
func CMFLanguages(mods ...qm.QueryMod) cmfLanguageQuery {
	mods = append(mods, qm.From("`cmf_language`"))
	return cmfLanguageQuery{NewQuery(mods...)}
}

// FindCMFLanguage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFLanguage(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CMFLanguage, error) {
	cmfLanguageObj := &CMFLanguage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_language` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfLanguageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_language")
	}

	return cmfLanguageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFLanguage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_language provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfLanguageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfLanguageInsertCacheMut.RLock()
	cache, cached := cmfLanguageInsertCache[key]
	cmfLanguageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfLanguageAllColumns,
			cmfLanguageColumnsWithDefault,
			cmfLanguageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfLanguageType, cmfLanguageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfLanguageType, cmfLanguageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_language` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_language` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_language` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfLanguagePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_language")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfLanguageMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_language")
	}

CacheNoHooks:
	if !cached {
		cmfLanguageInsertCacheMut.Lock()
		cmfLanguageInsertCache[key] = cache
		cmfLanguageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFLanguage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFLanguage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfLanguageUpdateCacheMut.RLock()
	cache, cached := cmfLanguageUpdateCache[key]
	cmfLanguageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfLanguageAllColumns,
			cmfLanguagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_language, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_language` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfLanguagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfLanguageType, cmfLanguageMapping, append(wl, cmfLanguagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_language row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_language")
	}

	if !cached {
		cmfLanguageUpdateCacheMut.Lock()
		cmfLanguageUpdateCache[key] = cache
		cmfLanguageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfLanguageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_language")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_language")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFLanguageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_language` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLanguagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfLanguage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfLanguage")
	}
	return rowsAff, nil
}

var mySQLCMFLanguageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFLanguage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_language provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfLanguageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFLanguageUniqueColumns, o)

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

	cmfLanguageUpsertCacheMut.RLock()
	cache, cached := cmfLanguageUpsertCache[key]
	cmfLanguageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfLanguageAllColumns,
			cmfLanguageColumnsWithDefault,
			cmfLanguageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfLanguageAllColumns,
			cmfLanguagePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_language, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_language`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_language` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfLanguageType, cmfLanguageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfLanguageType, cmfLanguageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_language")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfLanguageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfLanguageType, cmfLanguageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_language")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_language")
	}

CacheNoHooks:
	if !cached {
		cmfLanguageUpsertCacheMut.Lock()
		cmfLanguageUpsertCache[key] = cache
		cmfLanguageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFLanguage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFLanguage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFLanguage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfLanguagePrimaryKeyMapping)
	sql := "DELETE FROM `cmf_language` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_language")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_language")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfLanguageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfLanguageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_language")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_language")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFLanguageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfLanguageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_language` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLanguagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfLanguage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_language")
	}

	if len(cmfLanguageAfterDeleteHooks) != 0 {
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
func (o *CMFLanguage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFLanguage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFLanguageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFLanguageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLanguagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_language`.* FROM `cmf_language` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLanguagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFLanguageSlice")
	}

	*o = slice

	return nil
}

// CMFLanguageExists checks if the CMFLanguage row exists.
func CMFLanguageExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_language` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_language exists")
	}

	return exists, nil
}