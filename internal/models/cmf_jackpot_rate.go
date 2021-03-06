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

// CMFJackpotRate is an object representing the database table.
type CMFJackpotRate struct {
	ID          uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Giftid      int         `boil:"giftid" json:"giftid" toml:"giftid" yaml:"giftid"`
	Nums        uint        `boil:"nums" json:"nums" toml:"nums" yaml:"nums"`
	RateJackpot null.String `boil:"rate_jackpot" json:"rate_jackpot,omitempty" toml:"rate_jackpot" yaml:"rate_jackpot,omitempty"`

	R *cmfJackpotRateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfJackpotRateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFJackpotRateColumns = struct {
	ID          string
	Giftid      string
	Nums        string
	RateJackpot string
}{
	ID:          "id",
	Giftid:      "giftid",
	Nums:        "nums",
	RateJackpot: "rate_jackpot",
}

// Generated where

var CMFJackpotRateWhere = struct {
	ID          whereHelperuint64
	Giftid      whereHelperint
	Nums        whereHelperuint
	RateJackpot whereHelpernull_String
}{
	ID:          whereHelperuint64{field: "`cmf_jackpot_rate`.`id`"},
	Giftid:      whereHelperint{field: "`cmf_jackpot_rate`.`giftid`"},
	Nums:        whereHelperuint{field: "`cmf_jackpot_rate`.`nums`"},
	RateJackpot: whereHelpernull_String{field: "`cmf_jackpot_rate`.`rate_jackpot`"},
}

// CMFJackpotRateRels is where relationship names are stored.
var CMFJackpotRateRels = struct {
}{}

// cmfJackpotRateR is where relationships are stored.
type cmfJackpotRateR struct {
}

// NewStruct creates a new relationship struct
func (*cmfJackpotRateR) NewStruct() *cmfJackpotRateR {
	return &cmfJackpotRateR{}
}

// cmfJackpotRateL is where Load methods for each relationship are stored.
type cmfJackpotRateL struct{}

var (
	cmfJackpotRateAllColumns            = []string{"id", "giftid", "nums", "rate_jackpot"}
	cmfJackpotRateColumnsWithoutDefault = []string{"nums", "rate_jackpot"}
	cmfJackpotRateColumnsWithDefault    = []string{"id", "giftid"}
	cmfJackpotRatePrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFJackpotRateSlice is an alias for a slice of pointers to CMFJackpotRate.
	// This should generally be used opposed to []CMFJackpotRate.
	CMFJackpotRateSlice []*CMFJackpotRate
	// CMFJackpotRateHook is the signature for custom CMFJackpotRate hook methods
	CMFJackpotRateHook func(context.Context, boil.ContextExecutor, *CMFJackpotRate) error

	cmfJackpotRateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfJackpotRateType                 = reflect.TypeOf(&CMFJackpotRate{})
	cmfJackpotRateMapping              = queries.MakeStructMapping(cmfJackpotRateType)
	cmfJackpotRatePrimaryKeyMapping, _ = queries.BindMapping(cmfJackpotRateType, cmfJackpotRateMapping, cmfJackpotRatePrimaryKeyColumns)
	cmfJackpotRateInsertCacheMut       sync.RWMutex
	cmfJackpotRateInsertCache          = make(map[string]insertCache)
	cmfJackpotRateUpdateCacheMut       sync.RWMutex
	cmfJackpotRateUpdateCache          = make(map[string]updateCache)
	cmfJackpotRateUpsertCacheMut       sync.RWMutex
	cmfJackpotRateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfJackpotRateBeforeInsertHooks []CMFJackpotRateHook
var cmfJackpotRateBeforeUpdateHooks []CMFJackpotRateHook
var cmfJackpotRateBeforeDeleteHooks []CMFJackpotRateHook
var cmfJackpotRateBeforeUpsertHooks []CMFJackpotRateHook

var cmfJackpotRateAfterInsertHooks []CMFJackpotRateHook
var cmfJackpotRateAfterSelectHooks []CMFJackpotRateHook
var cmfJackpotRateAfterUpdateHooks []CMFJackpotRateHook
var cmfJackpotRateAfterDeleteHooks []CMFJackpotRateHook
var cmfJackpotRateAfterUpsertHooks []CMFJackpotRateHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFJackpotRate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFJackpotRate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFJackpotRate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFJackpotRate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFJackpotRate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFJackpotRate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFJackpotRate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFJackpotRate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFJackpotRate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfJackpotRateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFJackpotRateHook registers your hook function for all future operations.
func AddCMFJackpotRateHook(hookPoint boil.HookPoint, cmfJackpotRateHook CMFJackpotRateHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfJackpotRateBeforeInsertHooks = append(cmfJackpotRateBeforeInsertHooks, cmfJackpotRateHook)
	case boil.BeforeUpdateHook:
		cmfJackpotRateBeforeUpdateHooks = append(cmfJackpotRateBeforeUpdateHooks, cmfJackpotRateHook)
	case boil.BeforeDeleteHook:
		cmfJackpotRateBeforeDeleteHooks = append(cmfJackpotRateBeforeDeleteHooks, cmfJackpotRateHook)
	case boil.BeforeUpsertHook:
		cmfJackpotRateBeforeUpsertHooks = append(cmfJackpotRateBeforeUpsertHooks, cmfJackpotRateHook)
	case boil.AfterInsertHook:
		cmfJackpotRateAfterInsertHooks = append(cmfJackpotRateAfterInsertHooks, cmfJackpotRateHook)
	case boil.AfterSelectHook:
		cmfJackpotRateAfterSelectHooks = append(cmfJackpotRateAfterSelectHooks, cmfJackpotRateHook)
	case boil.AfterUpdateHook:
		cmfJackpotRateAfterUpdateHooks = append(cmfJackpotRateAfterUpdateHooks, cmfJackpotRateHook)
	case boil.AfterDeleteHook:
		cmfJackpotRateAfterDeleteHooks = append(cmfJackpotRateAfterDeleteHooks, cmfJackpotRateHook)
	case boil.AfterUpsertHook:
		cmfJackpotRateAfterUpsertHooks = append(cmfJackpotRateAfterUpsertHooks, cmfJackpotRateHook)
	}
}

// One returns a single cmfJackpotRate record from the query.
func (q cmfJackpotRateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFJackpotRate, error) {
	o := &CMFJackpotRate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_jackpot_rate")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFJackpotRate records from the query.
func (q cmfJackpotRateQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFJackpotRateSlice, error) {
	var o []*CMFJackpotRate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFJackpotRate slice")
	}

	if len(cmfJackpotRateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFJackpotRate records in the query.
func (q cmfJackpotRateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_jackpot_rate rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfJackpotRateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_jackpot_rate exists")
	}

	return count > 0, nil
}

// CMFJackpotRates retrieves all the records using an executor.
func CMFJackpotRates(mods ...qm.QueryMod) cmfJackpotRateQuery {
	mods = append(mods, qm.From("`cmf_jackpot_rate`"))
	return cmfJackpotRateQuery{NewQuery(mods...)}
}

// FindCMFJackpotRate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFJackpotRate(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*CMFJackpotRate, error) {
	cmfJackpotRateObj := &CMFJackpotRate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_jackpot_rate` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfJackpotRateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_jackpot_rate")
	}

	return cmfJackpotRateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFJackpotRate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_jackpot_rate provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfJackpotRateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfJackpotRateInsertCacheMut.RLock()
	cache, cached := cmfJackpotRateInsertCache[key]
	cmfJackpotRateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfJackpotRateAllColumns,
			cmfJackpotRateColumnsWithDefault,
			cmfJackpotRateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfJackpotRateType, cmfJackpotRateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfJackpotRateType, cmfJackpotRateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_jackpot_rate` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_jackpot_rate` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_jackpot_rate` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfJackpotRatePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_jackpot_rate")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfJackpotRateMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_jackpot_rate")
	}

CacheNoHooks:
	if !cached {
		cmfJackpotRateInsertCacheMut.Lock()
		cmfJackpotRateInsertCache[key] = cache
		cmfJackpotRateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFJackpotRate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFJackpotRate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfJackpotRateUpdateCacheMut.RLock()
	cache, cached := cmfJackpotRateUpdateCache[key]
	cmfJackpotRateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfJackpotRateAllColumns,
			cmfJackpotRatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_jackpot_rate, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_jackpot_rate` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfJackpotRatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfJackpotRateType, cmfJackpotRateMapping, append(wl, cmfJackpotRatePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_jackpot_rate row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_jackpot_rate")
	}

	if !cached {
		cmfJackpotRateUpdateCacheMut.Lock()
		cmfJackpotRateUpdateCache[key] = cache
		cmfJackpotRateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfJackpotRateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_jackpot_rate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_jackpot_rate")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFJackpotRateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfJackpotRatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_jackpot_rate` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfJackpotRatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfJackpotRate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfJackpotRate")
	}
	return rowsAff, nil
}

var mySQLCMFJackpotRateUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFJackpotRate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_jackpot_rate provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfJackpotRateColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFJackpotRateUniqueColumns, o)

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

	cmfJackpotRateUpsertCacheMut.RLock()
	cache, cached := cmfJackpotRateUpsertCache[key]
	cmfJackpotRateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfJackpotRateAllColumns,
			cmfJackpotRateColumnsWithDefault,
			cmfJackpotRateColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfJackpotRateAllColumns,
			cmfJackpotRatePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_jackpot_rate, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_jackpot_rate`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_jackpot_rate` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfJackpotRateType, cmfJackpotRateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfJackpotRateType, cmfJackpotRateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_jackpot_rate")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfJackpotRateMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfJackpotRateType, cmfJackpotRateMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_jackpot_rate")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_jackpot_rate")
	}

CacheNoHooks:
	if !cached {
		cmfJackpotRateUpsertCacheMut.Lock()
		cmfJackpotRateUpsertCache[key] = cache
		cmfJackpotRateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFJackpotRate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFJackpotRate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFJackpotRate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfJackpotRatePrimaryKeyMapping)
	sql := "DELETE FROM `cmf_jackpot_rate` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_jackpot_rate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_jackpot_rate")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfJackpotRateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfJackpotRateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_jackpot_rate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_jackpot_rate")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFJackpotRateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfJackpotRateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfJackpotRatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_jackpot_rate` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfJackpotRatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfJackpotRate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_jackpot_rate")
	}

	if len(cmfJackpotRateAfterDeleteHooks) != 0 {
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
func (o *CMFJackpotRate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFJackpotRate(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFJackpotRateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFJackpotRateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfJackpotRatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_jackpot_rate`.* FROM `cmf_jackpot_rate` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfJackpotRatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFJackpotRateSlice")
	}

	*o = slice

	return nil
}

// CMFJackpotRateExists checks if the CMFJackpotRate row exists.
func CMFJackpotRateExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_jackpot_rate` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_jackpot_rate exists")
	}

	return exists, nil
}
