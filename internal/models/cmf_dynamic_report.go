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

// CMFDynamicReport is an object representing the database table.
type CMFDynamicReport struct {
	ID          uint   `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID         int    `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Touid       int    `boil:"touid" json:"touid" toml:"touid" yaml:"touid"`
	Dynamicid   int    `boil:"dynamicid" json:"dynamicid" toml:"dynamicid" yaml:"dynamicid"`
	Content     string `boil:"content" json:"content" toml:"content" yaml:"content"`
	Status      bool   `boil:"status" json:"status" toml:"status" yaml:"status"`
	Addtime     int    `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Uptime      int    `boil:"uptime" json:"uptime" toml:"uptime" yaml:"uptime"`
	DynamicType int    `boil:"dynamic_type" json:"dynamic_type" toml:"dynamic_type" yaml:"dynamic_type"`

	R *cmfDynamicReportR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfDynamicReportL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFDynamicReportColumns = struct {
	ID          string
	UID         string
	Touid       string
	Dynamicid   string
	Content     string
	Status      string
	Addtime     string
	Uptime      string
	DynamicType string
}{
	ID:          "id",
	UID:         "uid",
	Touid:       "touid",
	Dynamicid:   "dynamicid",
	Content:     "content",
	Status:      "status",
	Addtime:     "addtime",
	Uptime:      "uptime",
	DynamicType: "dynamic_type",
}

// Generated where

var CMFDynamicReportWhere = struct {
	ID          whereHelperuint
	UID         whereHelperint
	Touid       whereHelperint
	Dynamicid   whereHelperint
	Content     whereHelperstring
	Status      whereHelperbool
	Addtime     whereHelperint
	Uptime      whereHelperint
	DynamicType whereHelperint
}{
	ID:          whereHelperuint{field: "`cmf_dynamic_report`.`id`"},
	UID:         whereHelperint{field: "`cmf_dynamic_report`.`uid`"},
	Touid:       whereHelperint{field: "`cmf_dynamic_report`.`touid`"},
	Dynamicid:   whereHelperint{field: "`cmf_dynamic_report`.`dynamicid`"},
	Content:     whereHelperstring{field: "`cmf_dynamic_report`.`content`"},
	Status:      whereHelperbool{field: "`cmf_dynamic_report`.`status`"},
	Addtime:     whereHelperint{field: "`cmf_dynamic_report`.`addtime`"},
	Uptime:      whereHelperint{field: "`cmf_dynamic_report`.`uptime`"},
	DynamicType: whereHelperint{field: "`cmf_dynamic_report`.`dynamic_type`"},
}

// CMFDynamicReportRels is where relationship names are stored.
var CMFDynamicReportRels = struct {
}{}

// cmfDynamicReportR is where relationships are stored.
type cmfDynamicReportR struct {
}

// NewStruct creates a new relationship struct
func (*cmfDynamicReportR) NewStruct() *cmfDynamicReportR {
	return &cmfDynamicReportR{}
}

// cmfDynamicReportL is where Load methods for each relationship are stored.
type cmfDynamicReportL struct{}

var (
	cmfDynamicReportAllColumns            = []string{"id", "uid", "touid", "dynamicid", "content", "status", "addtime", "uptime", "dynamic_type"}
	cmfDynamicReportColumnsWithoutDefault = []string{"content"}
	cmfDynamicReportColumnsWithDefault    = []string{"id", "uid", "touid", "dynamicid", "status", "addtime", "uptime", "dynamic_type"}
	cmfDynamicReportPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFDynamicReportSlice is an alias for a slice of pointers to CMFDynamicReport.
	// This should generally be used opposed to []CMFDynamicReport.
	CMFDynamicReportSlice []*CMFDynamicReport
	// CMFDynamicReportHook is the signature for custom CMFDynamicReport hook methods
	CMFDynamicReportHook func(context.Context, boil.ContextExecutor, *CMFDynamicReport) error

	cmfDynamicReportQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfDynamicReportType                 = reflect.TypeOf(&CMFDynamicReport{})
	cmfDynamicReportMapping              = queries.MakeStructMapping(cmfDynamicReportType)
	cmfDynamicReportPrimaryKeyMapping, _ = queries.BindMapping(cmfDynamicReportType, cmfDynamicReportMapping, cmfDynamicReportPrimaryKeyColumns)
	cmfDynamicReportInsertCacheMut       sync.RWMutex
	cmfDynamicReportInsertCache          = make(map[string]insertCache)
	cmfDynamicReportUpdateCacheMut       sync.RWMutex
	cmfDynamicReportUpdateCache          = make(map[string]updateCache)
	cmfDynamicReportUpsertCacheMut       sync.RWMutex
	cmfDynamicReportUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfDynamicReportBeforeInsertHooks []CMFDynamicReportHook
var cmfDynamicReportBeforeUpdateHooks []CMFDynamicReportHook
var cmfDynamicReportBeforeDeleteHooks []CMFDynamicReportHook
var cmfDynamicReportBeforeUpsertHooks []CMFDynamicReportHook

var cmfDynamicReportAfterInsertHooks []CMFDynamicReportHook
var cmfDynamicReportAfterSelectHooks []CMFDynamicReportHook
var cmfDynamicReportAfterUpdateHooks []CMFDynamicReportHook
var cmfDynamicReportAfterDeleteHooks []CMFDynamicReportHook
var cmfDynamicReportAfterUpsertHooks []CMFDynamicReportHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFDynamicReport) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFDynamicReport) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFDynamicReport) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFDynamicReport) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFDynamicReport) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFDynamicReport) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFDynamicReport) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFDynamicReport) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFDynamicReport) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfDynamicReportAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFDynamicReportHook registers your hook function for all future operations.
func AddCMFDynamicReportHook(hookPoint boil.HookPoint, cmfDynamicReportHook CMFDynamicReportHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfDynamicReportBeforeInsertHooks = append(cmfDynamicReportBeforeInsertHooks, cmfDynamicReportHook)
	case boil.BeforeUpdateHook:
		cmfDynamicReportBeforeUpdateHooks = append(cmfDynamicReportBeforeUpdateHooks, cmfDynamicReportHook)
	case boil.BeforeDeleteHook:
		cmfDynamicReportBeforeDeleteHooks = append(cmfDynamicReportBeforeDeleteHooks, cmfDynamicReportHook)
	case boil.BeforeUpsertHook:
		cmfDynamicReportBeforeUpsertHooks = append(cmfDynamicReportBeforeUpsertHooks, cmfDynamicReportHook)
	case boil.AfterInsertHook:
		cmfDynamicReportAfterInsertHooks = append(cmfDynamicReportAfterInsertHooks, cmfDynamicReportHook)
	case boil.AfterSelectHook:
		cmfDynamicReportAfterSelectHooks = append(cmfDynamicReportAfterSelectHooks, cmfDynamicReportHook)
	case boil.AfterUpdateHook:
		cmfDynamicReportAfterUpdateHooks = append(cmfDynamicReportAfterUpdateHooks, cmfDynamicReportHook)
	case boil.AfterDeleteHook:
		cmfDynamicReportAfterDeleteHooks = append(cmfDynamicReportAfterDeleteHooks, cmfDynamicReportHook)
	case boil.AfterUpsertHook:
		cmfDynamicReportAfterUpsertHooks = append(cmfDynamicReportAfterUpsertHooks, cmfDynamicReportHook)
	}
}

// One returns a single cmfDynamicReport record from the query.
func (q cmfDynamicReportQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFDynamicReport, error) {
	o := &CMFDynamicReport{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_dynamic_report")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFDynamicReport records from the query.
func (q cmfDynamicReportQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFDynamicReportSlice, error) {
	var o []*CMFDynamicReport

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFDynamicReport slice")
	}

	if len(cmfDynamicReportAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFDynamicReport records in the query.
func (q cmfDynamicReportQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_dynamic_report rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfDynamicReportQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_dynamic_report exists")
	}

	return count > 0, nil
}

// CMFDynamicReports retrieves all the records using an executor.
func CMFDynamicReports(mods ...qm.QueryMod) cmfDynamicReportQuery {
	mods = append(mods, qm.From("`cmf_dynamic_report`"))
	return cmfDynamicReportQuery{NewQuery(mods...)}
}

// FindCMFDynamicReport retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFDynamicReport(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CMFDynamicReport, error) {
	cmfDynamicReportObj := &CMFDynamicReport{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_dynamic_report` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfDynamicReportObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_dynamic_report")
	}

	return cmfDynamicReportObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFDynamicReport) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_dynamic_report provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfDynamicReportColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfDynamicReportInsertCacheMut.RLock()
	cache, cached := cmfDynamicReportInsertCache[key]
	cmfDynamicReportInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfDynamicReportAllColumns,
			cmfDynamicReportColumnsWithDefault,
			cmfDynamicReportColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfDynamicReportType, cmfDynamicReportMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfDynamicReportType, cmfDynamicReportMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_dynamic_report` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_dynamic_report` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_dynamic_report` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfDynamicReportPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_dynamic_report")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfDynamicReportMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_dynamic_report")
	}

CacheNoHooks:
	if !cached {
		cmfDynamicReportInsertCacheMut.Lock()
		cmfDynamicReportInsertCache[key] = cache
		cmfDynamicReportInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFDynamicReport.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFDynamicReport) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfDynamicReportUpdateCacheMut.RLock()
	cache, cached := cmfDynamicReportUpdateCache[key]
	cmfDynamicReportUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfDynamicReportAllColumns,
			cmfDynamicReportPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_dynamic_report, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_dynamic_report` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfDynamicReportPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfDynamicReportType, cmfDynamicReportMapping, append(wl, cmfDynamicReportPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_dynamic_report row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_dynamic_report")
	}

	if !cached {
		cmfDynamicReportUpdateCacheMut.Lock()
		cmfDynamicReportUpdateCache[key] = cache
		cmfDynamicReportUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfDynamicReportQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_dynamic_report")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_dynamic_report")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFDynamicReportSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfDynamicReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_dynamic_report` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfDynamicReportPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfDynamicReport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfDynamicReport")
	}
	return rowsAff, nil
}

var mySQLCMFDynamicReportUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFDynamicReport) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_dynamic_report provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfDynamicReportColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFDynamicReportUniqueColumns, o)

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

	cmfDynamicReportUpsertCacheMut.RLock()
	cache, cached := cmfDynamicReportUpsertCache[key]
	cmfDynamicReportUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfDynamicReportAllColumns,
			cmfDynamicReportColumnsWithDefault,
			cmfDynamicReportColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfDynamicReportAllColumns,
			cmfDynamicReportPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_dynamic_report, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_dynamic_report`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_dynamic_report` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfDynamicReportType, cmfDynamicReportMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfDynamicReportType, cmfDynamicReportMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_dynamic_report")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfDynamicReportMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfDynamicReportType, cmfDynamicReportMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_dynamic_report")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_dynamic_report")
	}

CacheNoHooks:
	if !cached {
		cmfDynamicReportUpsertCacheMut.Lock()
		cmfDynamicReportUpsertCache[key] = cache
		cmfDynamicReportUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFDynamicReport record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFDynamicReport) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFDynamicReport provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfDynamicReportPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_dynamic_report` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_dynamic_report")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_dynamic_report")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfDynamicReportQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfDynamicReportQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_dynamic_report")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_dynamic_report")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFDynamicReportSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfDynamicReportBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfDynamicReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_dynamic_report` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfDynamicReportPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfDynamicReport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_dynamic_report")
	}

	if len(cmfDynamicReportAfterDeleteHooks) != 0 {
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
func (o *CMFDynamicReport) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFDynamicReport(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFDynamicReportSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFDynamicReportSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfDynamicReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_dynamic_report`.* FROM `cmf_dynamic_report` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfDynamicReportPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFDynamicReportSlice")
	}

	*o = slice

	return nil
}

// CMFDynamicReportExists checks if the CMFDynamicReport row exists.
func CMFDynamicReportExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_dynamic_report` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_dynamic_report exists")
	}

	return exists, nil
}