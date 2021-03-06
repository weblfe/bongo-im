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

// CMFGetcodeLimitIP is an object representing the database table.
type CMFGetcodeLimitIP struct {
	IP    int64  `boil:"ip" json:"ip" toml:"ip" yaml:"ip"`
	Date  string `boil:"date" json:"date" toml:"date" yaml:"date"`
	Times int8   `boil:"times" json:"times" toml:"times" yaml:"times"`

	R *cmfGetcodeLimitIPR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfGetcodeLimitIPL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFGetcodeLimitIPColumns = struct {
	IP    string
	Date  string
	Times string
}{
	IP:    "ip",
	Date:  "date",
	Times: "times",
}

// Generated where

type whereHelperint8 struct{ field string }

func (w whereHelperint8) EQ(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint8) NEQ(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint8) LT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint8) LTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint8) GT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint8) GTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint8) IN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint8) NIN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var CMFGetcodeLimitIPWhere = struct {
	IP    whereHelperint64
	Date  whereHelperstring
	Times whereHelperint8
}{
	IP:    whereHelperint64{field: "`cmf_getcode_limit_ip`.`ip`"},
	Date:  whereHelperstring{field: "`cmf_getcode_limit_ip`.`date`"},
	Times: whereHelperint8{field: "`cmf_getcode_limit_ip`.`times`"},
}

// CMFGetcodeLimitIPRels is where relationship names are stored.
var CMFGetcodeLimitIPRels = struct {
}{}

// cmfGetcodeLimitIPR is where relationships are stored.
type cmfGetcodeLimitIPR struct {
}

// NewStruct creates a new relationship struct
func (*cmfGetcodeLimitIPR) NewStruct() *cmfGetcodeLimitIPR {
	return &cmfGetcodeLimitIPR{}
}

// cmfGetcodeLimitIPL is where Load methods for each relationship are stored.
type cmfGetcodeLimitIPL struct{}

var (
	cmfGetcodeLimitIPAllColumns            = []string{"ip", "date", "times"}
	cmfGetcodeLimitIPColumnsWithoutDefault = []string{"ip", "date"}
	cmfGetcodeLimitIPColumnsWithDefault    = []string{"times"}
	cmfGetcodeLimitIPPrimaryKeyColumns     = []string{"ip"}
)

type (
	// CMFGetcodeLimitIPSlice is an alias for a slice of pointers to CMFGetcodeLimitIP.
	// This should generally be used opposed to []CMFGetcodeLimitIP.
	CMFGetcodeLimitIPSlice []*CMFGetcodeLimitIP
	// CMFGetcodeLimitIPHook is the signature for custom CMFGetcodeLimitIP hook methods
	CMFGetcodeLimitIPHook func(context.Context, boil.ContextExecutor, *CMFGetcodeLimitIP) error

	cmfGetcodeLimitIPQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfGetcodeLimitIPType                 = reflect.TypeOf(&CMFGetcodeLimitIP{})
	cmfGetcodeLimitIPMapping              = queries.MakeStructMapping(cmfGetcodeLimitIPType)
	cmfGetcodeLimitIPPrimaryKeyMapping, _ = queries.BindMapping(cmfGetcodeLimitIPType, cmfGetcodeLimitIPMapping, cmfGetcodeLimitIPPrimaryKeyColumns)
	cmfGetcodeLimitIPInsertCacheMut       sync.RWMutex
	cmfGetcodeLimitIPInsertCache          = make(map[string]insertCache)
	cmfGetcodeLimitIPUpdateCacheMut       sync.RWMutex
	cmfGetcodeLimitIPUpdateCache          = make(map[string]updateCache)
	cmfGetcodeLimitIPUpsertCacheMut       sync.RWMutex
	cmfGetcodeLimitIPUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfGetcodeLimitIPBeforeInsertHooks []CMFGetcodeLimitIPHook
var cmfGetcodeLimitIPBeforeUpdateHooks []CMFGetcodeLimitIPHook
var cmfGetcodeLimitIPBeforeDeleteHooks []CMFGetcodeLimitIPHook
var cmfGetcodeLimitIPBeforeUpsertHooks []CMFGetcodeLimitIPHook

var cmfGetcodeLimitIPAfterInsertHooks []CMFGetcodeLimitIPHook
var cmfGetcodeLimitIPAfterSelectHooks []CMFGetcodeLimitIPHook
var cmfGetcodeLimitIPAfterUpdateHooks []CMFGetcodeLimitIPHook
var cmfGetcodeLimitIPAfterDeleteHooks []CMFGetcodeLimitIPHook
var cmfGetcodeLimitIPAfterUpsertHooks []CMFGetcodeLimitIPHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFGetcodeLimitIP) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFGetcodeLimitIP) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFGetcodeLimitIP) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFGetcodeLimitIP) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFGetcodeLimitIP) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFGetcodeLimitIP) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFGetcodeLimitIP) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFGetcodeLimitIP) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFGetcodeLimitIP) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfGetcodeLimitIPAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFGetcodeLimitIPHook registers your hook function for all future operations.
func AddCMFGetcodeLimitIPHook(hookPoint boil.HookPoint, cmfGetcodeLimitIPHook CMFGetcodeLimitIPHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfGetcodeLimitIPBeforeInsertHooks = append(cmfGetcodeLimitIPBeforeInsertHooks, cmfGetcodeLimitIPHook)
	case boil.BeforeUpdateHook:
		cmfGetcodeLimitIPBeforeUpdateHooks = append(cmfGetcodeLimitIPBeforeUpdateHooks, cmfGetcodeLimitIPHook)
	case boil.BeforeDeleteHook:
		cmfGetcodeLimitIPBeforeDeleteHooks = append(cmfGetcodeLimitIPBeforeDeleteHooks, cmfGetcodeLimitIPHook)
	case boil.BeforeUpsertHook:
		cmfGetcodeLimitIPBeforeUpsertHooks = append(cmfGetcodeLimitIPBeforeUpsertHooks, cmfGetcodeLimitIPHook)
	case boil.AfterInsertHook:
		cmfGetcodeLimitIPAfterInsertHooks = append(cmfGetcodeLimitIPAfterInsertHooks, cmfGetcodeLimitIPHook)
	case boil.AfterSelectHook:
		cmfGetcodeLimitIPAfterSelectHooks = append(cmfGetcodeLimitIPAfterSelectHooks, cmfGetcodeLimitIPHook)
	case boil.AfterUpdateHook:
		cmfGetcodeLimitIPAfterUpdateHooks = append(cmfGetcodeLimitIPAfterUpdateHooks, cmfGetcodeLimitIPHook)
	case boil.AfterDeleteHook:
		cmfGetcodeLimitIPAfterDeleteHooks = append(cmfGetcodeLimitIPAfterDeleteHooks, cmfGetcodeLimitIPHook)
	case boil.AfterUpsertHook:
		cmfGetcodeLimitIPAfterUpsertHooks = append(cmfGetcodeLimitIPAfterUpsertHooks, cmfGetcodeLimitIPHook)
	}
}

// One returns a single cmfGetcodeLimitIP record from the query.
func (q cmfGetcodeLimitIPQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFGetcodeLimitIP, error) {
	o := &CMFGetcodeLimitIP{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_getcode_limit_ip")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFGetcodeLimitIP records from the query.
func (q cmfGetcodeLimitIPQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFGetcodeLimitIPSlice, error) {
	var o []*CMFGetcodeLimitIP

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFGetcodeLimitIP slice")
	}

	if len(cmfGetcodeLimitIPAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFGetcodeLimitIP records in the query.
func (q cmfGetcodeLimitIPQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_getcode_limit_ip rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfGetcodeLimitIPQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_getcode_limit_ip exists")
	}

	return count > 0, nil
}

// CMFGetcodeLimitIps retrieves all the records using an executor.
func CMFGetcodeLimitIps(mods ...qm.QueryMod) cmfGetcodeLimitIPQuery {
	mods = append(mods, qm.From("`cmf_getcode_limit_ip`"))
	return cmfGetcodeLimitIPQuery{NewQuery(mods...)}
}

// FindCMFGetcodeLimitIP retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFGetcodeLimitIP(ctx context.Context, exec boil.ContextExecutor, iP int64, selectCols ...string) (*CMFGetcodeLimitIP, error) {
	cmfGetcodeLimitIPObj := &CMFGetcodeLimitIP{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_getcode_limit_ip` where `ip`=?", sel,
	)

	q := queries.Raw(query, iP)

	err := q.Bind(ctx, exec, cmfGetcodeLimitIPObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_getcode_limit_ip")
	}

	return cmfGetcodeLimitIPObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFGetcodeLimitIP) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_getcode_limit_ip provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfGetcodeLimitIPColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfGetcodeLimitIPInsertCacheMut.RLock()
	cache, cached := cmfGetcodeLimitIPInsertCache[key]
	cmfGetcodeLimitIPInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfGetcodeLimitIPAllColumns,
			cmfGetcodeLimitIPColumnsWithDefault,
			cmfGetcodeLimitIPColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfGetcodeLimitIPType, cmfGetcodeLimitIPMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfGetcodeLimitIPType, cmfGetcodeLimitIPMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_getcode_limit_ip` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_getcode_limit_ip` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_getcode_limit_ip` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfGetcodeLimitIPPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into cmf_getcode_limit_ip")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IP,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_getcode_limit_ip")
	}

CacheNoHooks:
	if !cached {
		cmfGetcodeLimitIPInsertCacheMut.Lock()
		cmfGetcodeLimitIPInsertCache[key] = cache
		cmfGetcodeLimitIPInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFGetcodeLimitIP.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFGetcodeLimitIP) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfGetcodeLimitIPUpdateCacheMut.RLock()
	cache, cached := cmfGetcodeLimitIPUpdateCache[key]
	cmfGetcodeLimitIPUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfGetcodeLimitIPAllColumns,
			cmfGetcodeLimitIPPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_getcode_limit_ip, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_getcode_limit_ip` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfGetcodeLimitIPPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfGetcodeLimitIPType, cmfGetcodeLimitIPMapping, append(wl, cmfGetcodeLimitIPPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_getcode_limit_ip row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_getcode_limit_ip")
	}

	if !cached {
		cmfGetcodeLimitIPUpdateCacheMut.Lock()
		cmfGetcodeLimitIPUpdateCache[key] = cache
		cmfGetcodeLimitIPUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfGetcodeLimitIPQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_getcode_limit_ip")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_getcode_limit_ip")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFGetcodeLimitIPSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGetcodeLimitIPPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_getcode_limit_ip` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGetcodeLimitIPPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfGetcodeLimitIP slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfGetcodeLimitIP")
	}
	return rowsAff, nil
}

var mySQLCMFGetcodeLimitIPUniqueColumns = []string{
	"ip",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFGetcodeLimitIP) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_getcode_limit_ip provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfGetcodeLimitIPColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFGetcodeLimitIPUniqueColumns, o)

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

	cmfGetcodeLimitIPUpsertCacheMut.RLock()
	cache, cached := cmfGetcodeLimitIPUpsertCache[key]
	cmfGetcodeLimitIPUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfGetcodeLimitIPAllColumns,
			cmfGetcodeLimitIPColumnsWithDefault,
			cmfGetcodeLimitIPColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfGetcodeLimitIPAllColumns,
			cmfGetcodeLimitIPPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_getcode_limit_ip, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_getcode_limit_ip`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_getcode_limit_ip` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfGetcodeLimitIPType, cmfGetcodeLimitIPMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfGetcodeLimitIPType, cmfGetcodeLimitIPMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for cmf_getcode_limit_ip")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfGetcodeLimitIPType, cmfGetcodeLimitIPMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_getcode_limit_ip")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_getcode_limit_ip")
	}

CacheNoHooks:
	if !cached {
		cmfGetcodeLimitIPUpsertCacheMut.Lock()
		cmfGetcodeLimitIPUpsertCache[key] = cache
		cmfGetcodeLimitIPUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFGetcodeLimitIP record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFGetcodeLimitIP) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFGetcodeLimitIP provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfGetcodeLimitIPPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_getcode_limit_ip` WHERE `ip`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_getcode_limit_ip")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_getcode_limit_ip")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfGetcodeLimitIPQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfGetcodeLimitIPQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_getcode_limit_ip")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_getcode_limit_ip")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFGetcodeLimitIPSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfGetcodeLimitIPBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGetcodeLimitIPPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_getcode_limit_ip` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGetcodeLimitIPPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfGetcodeLimitIP slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_getcode_limit_ip")
	}

	if len(cmfGetcodeLimitIPAfterDeleteHooks) != 0 {
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
func (o *CMFGetcodeLimitIP) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFGetcodeLimitIP(ctx, exec, o.IP)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFGetcodeLimitIPSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFGetcodeLimitIPSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfGetcodeLimitIPPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_getcode_limit_ip`.* FROM `cmf_getcode_limit_ip` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfGetcodeLimitIPPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFGetcodeLimitIPSlice")
	}

	*o = slice

	return nil
}

// CMFGetcodeLimitIPExists checks if the CMFGetcodeLimitIP row exists.
func CMFGetcodeLimitIPExists(ctx context.Context, exec boil.ContextExecutor, iP int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_getcode_limit_ip` where `ip`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iP)
	}
	row := exec.QueryRowContext(ctx, sql, iP)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_getcode_limit_ip exists")
	}

	return exists, nil
}
