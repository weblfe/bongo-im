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

// CMFReport is an object representing the database table.
type CMFReport struct {
	ID      uint   `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID     int    `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Touid   int    `boil:"touid" json:"touid" toml:"touid" yaml:"touid"`
	Content string `boil:"content" json:"content" toml:"content" yaml:"content"`
	Status  bool   `boil:"status" json:"status" toml:"status" yaml:"status"`
	Addtime int    `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Uptime  int    `boil:"uptime" json:"uptime" toml:"uptime" yaml:"uptime"`

	R *cmfReportR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfReportL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFReportColumns = struct {
	ID      string
	UID     string
	Touid   string
	Content string
	Status  string
	Addtime string
	Uptime  string
}{
	ID:      "id",
	UID:     "uid",
	Touid:   "touid",
	Content: "content",
	Status:  "status",
	Addtime: "addtime",
	Uptime:  "uptime",
}

// Generated where

var CMFReportWhere = struct {
	ID      whereHelperuint
	UID     whereHelperint
	Touid   whereHelperint
	Content whereHelperstring
	Status  whereHelperbool
	Addtime whereHelperint
	Uptime  whereHelperint
}{
	ID:      whereHelperuint{field: "`cmf_report`.`id`"},
	UID:     whereHelperint{field: "`cmf_report`.`uid`"},
	Touid:   whereHelperint{field: "`cmf_report`.`touid`"},
	Content: whereHelperstring{field: "`cmf_report`.`content`"},
	Status:  whereHelperbool{field: "`cmf_report`.`status`"},
	Addtime: whereHelperint{field: "`cmf_report`.`addtime`"},
	Uptime:  whereHelperint{field: "`cmf_report`.`uptime`"},
}

// CMFReportRels is where relationship names are stored.
var CMFReportRels = struct {
}{}

// cmfReportR is where relationships are stored.
type cmfReportR struct {
}

// NewStruct creates a new relationship struct
func (*cmfReportR) NewStruct() *cmfReportR {
	return &cmfReportR{}
}

// cmfReportL is where Load methods for each relationship are stored.
type cmfReportL struct{}

var (
	cmfReportAllColumns            = []string{"id", "uid", "touid", "content", "status", "addtime", "uptime"}
	cmfReportColumnsWithoutDefault = []string{"content"}
	cmfReportColumnsWithDefault    = []string{"id", "uid", "touid", "status", "addtime", "uptime"}
	cmfReportPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFReportSlice is an alias for a slice of pointers to CMFReport.
	// This should generally be used opposed to []CMFReport.
	CMFReportSlice []*CMFReport
	// CMFReportHook is the signature for custom CMFReport hook methods
	CMFReportHook func(context.Context, boil.ContextExecutor, *CMFReport) error

	cmfReportQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfReportType                 = reflect.TypeOf(&CMFReport{})
	cmfReportMapping              = queries.MakeStructMapping(cmfReportType)
	cmfReportPrimaryKeyMapping, _ = queries.BindMapping(cmfReportType, cmfReportMapping, cmfReportPrimaryKeyColumns)
	cmfReportInsertCacheMut       sync.RWMutex
	cmfReportInsertCache          = make(map[string]insertCache)
	cmfReportUpdateCacheMut       sync.RWMutex
	cmfReportUpdateCache          = make(map[string]updateCache)
	cmfReportUpsertCacheMut       sync.RWMutex
	cmfReportUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfReportBeforeInsertHooks []CMFReportHook
var cmfReportBeforeUpdateHooks []CMFReportHook
var cmfReportBeforeDeleteHooks []CMFReportHook
var cmfReportBeforeUpsertHooks []CMFReportHook

var cmfReportAfterInsertHooks []CMFReportHook
var cmfReportAfterSelectHooks []CMFReportHook
var cmfReportAfterUpdateHooks []CMFReportHook
var cmfReportAfterDeleteHooks []CMFReportHook
var cmfReportAfterUpsertHooks []CMFReportHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFReport) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFReport) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFReport) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFReport) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFReport) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFReport) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFReport) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFReport) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFReport) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfReportAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFReportHook registers your hook function for all future operations.
func AddCMFReportHook(hookPoint boil.HookPoint, cmfReportHook CMFReportHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfReportBeforeInsertHooks = append(cmfReportBeforeInsertHooks, cmfReportHook)
	case boil.BeforeUpdateHook:
		cmfReportBeforeUpdateHooks = append(cmfReportBeforeUpdateHooks, cmfReportHook)
	case boil.BeforeDeleteHook:
		cmfReportBeforeDeleteHooks = append(cmfReportBeforeDeleteHooks, cmfReportHook)
	case boil.BeforeUpsertHook:
		cmfReportBeforeUpsertHooks = append(cmfReportBeforeUpsertHooks, cmfReportHook)
	case boil.AfterInsertHook:
		cmfReportAfterInsertHooks = append(cmfReportAfterInsertHooks, cmfReportHook)
	case boil.AfterSelectHook:
		cmfReportAfterSelectHooks = append(cmfReportAfterSelectHooks, cmfReportHook)
	case boil.AfterUpdateHook:
		cmfReportAfterUpdateHooks = append(cmfReportAfterUpdateHooks, cmfReportHook)
	case boil.AfterDeleteHook:
		cmfReportAfterDeleteHooks = append(cmfReportAfterDeleteHooks, cmfReportHook)
	case boil.AfterUpsertHook:
		cmfReportAfterUpsertHooks = append(cmfReportAfterUpsertHooks, cmfReportHook)
	}
}

// One returns a single cmfReport record from the query.
func (q cmfReportQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFReport, error) {
	o := &CMFReport{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_report")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFReport records from the query.
func (q cmfReportQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFReportSlice, error) {
	var o []*CMFReport

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFReport slice")
	}

	if len(cmfReportAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFReport records in the query.
func (q cmfReportQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_report rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfReportQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_report exists")
	}

	return count > 0, nil
}

// CMFReports retrieves all the records using an executor.
func CMFReports(mods ...qm.QueryMod) cmfReportQuery {
	mods = append(mods, qm.From("`cmf_report`"))
	return cmfReportQuery{NewQuery(mods...)}
}

// FindCMFReport retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFReport(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CMFReport, error) {
	cmfReportObj := &CMFReport{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_report` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfReportObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_report")
	}

	return cmfReportObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFReport) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_report provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfReportColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfReportInsertCacheMut.RLock()
	cache, cached := cmfReportInsertCache[key]
	cmfReportInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfReportAllColumns,
			cmfReportColumnsWithDefault,
			cmfReportColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfReportType, cmfReportMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfReportType, cmfReportMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_report` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_report` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_report` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfReportPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_report")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfReportMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_report")
	}

CacheNoHooks:
	if !cached {
		cmfReportInsertCacheMut.Lock()
		cmfReportInsertCache[key] = cache
		cmfReportInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFReport.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFReport) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfReportUpdateCacheMut.RLock()
	cache, cached := cmfReportUpdateCache[key]
	cmfReportUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfReportAllColumns,
			cmfReportPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_report, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_report` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfReportPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfReportType, cmfReportMapping, append(wl, cmfReportPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_report row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_report")
	}

	if !cached {
		cmfReportUpdateCacheMut.Lock()
		cmfReportUpdateCache[key] = cache
		cmfReportUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfReportQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_report")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_report")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFReportSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_report` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfReportPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfReport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfReport")
	}
	return rowsAff, nil
}

var mySQLCMFReportUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFReport) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_report provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfReportColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFReportUniqueColumns, o)

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

	cmfReportUpsertCacheMut.RLock()
	cache, cached := cmfReportUpsertCache[key]
	cmfReportUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfReportAllColumns,
			cmfReportColumnsWithDefault,
			cmfReportColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfReportAllColumns,
			cmfReportPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_report, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_report`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_report` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfReportType, cmfReportMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfReportType, cmfReportMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_report")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfReportMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfReportType, cmfReportMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_report")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_report")
	}

CacheNoHooks:
	if !cached {
		cmfReportUpsertCacheMut.Lock()
		cmfReportUpsertCache[key] = cache
		cmfReportUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFReport record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFReport) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFReport provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfReportPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_report` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_report")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_report")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfReportQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfReportQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_report")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_report")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFReportSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfReportBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_report` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfReportPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfReport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_report")
	}

	if len(cmfReportAfterDeleteHooks) != 0 {
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
func (o *CMFReport) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFReport(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFReportSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFReportSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_report`.* FROM `cmf_report` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfReportPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFReportSlice")
	}

	*o = slice

	return nil
}

// CMFReportExists checks if the CMFReport row exists.
func CMFReportExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_report` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_report exists")
	}

	return exists, nil
}