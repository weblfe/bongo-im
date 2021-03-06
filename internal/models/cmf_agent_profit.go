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

// CMFAgentProfit is an object representing the database table.
type CMFAgentProfit struct {
	ID        int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID       null.Int          `boil:"uid" json:"uid,omitempty" toml:"uid" yaml:"uid,omitempty"`
	OneProfit types.NullDecimal `boil:"one_profit" json:"one_profit,omitempty" toml:"one_profit" yaml:"one_profit,omitempty"`

	R *cmfAgentProfitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfAgentProfitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFAgentProfitColumns = struct {
	ID        string
	UID       string
	OneProfit string
}{
	ID:        "id",
	UID:       "uid",
	OneProfit: "one_profit",
}

// Generated where

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertypes_NullDecimal struct{ field string }

func (w whereHelpertypes_NullDecimal) EQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_NullDecimal) NEQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_NullDecimal) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_NullDecimal) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}
func (w whereHelpertypes_NullDecimal) LT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_NullDecimal) LTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_NullDecimal) GT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_NullDecimal) GTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CMFAgentProfitWhere = struct {
	ID        whereHelperint
	UID       whereHelpernull_Int
	OneProfit whereHelpertypes_NullDecimal
}{
	ID:        whereHelperint{field: "`cmf_agent_profit`.`id`"},
	UID:       whereHelpernull_Int{field: "`cmf_agent_profit`.`uid`"},
	OneProfit: whereHelpertypes_NullDecimal{field: "`cmf_agent_profit`.`one_profit`"},
}

// CMFAgentProfitRels is where relationship names are stored.
var CMFAgentProfitRels = struct {
}{}

// cmfAgentProfitR is where relationships are stored.
type cmfAgentProfitR struct {
}

// NewStruct creates a new relationship struct
func (*cmfAgentProfitR) NewStruct() *cmfAgentProfitR {
	return &cmfAgentProfitR{}
}

// cmfAgentProfitL is where Load methods for each relationship are stored.
type cmfAgentProfitL struct{}

var (
	cmfAgentProfitAllColumns            = []string{"id", "uid", "one_profit"}
	cmfAgentProfitColumnsWithoutDefault = []string{}
	cmfAgentProfitColumnsWithDefault    = []string{"id", "uid", "one_profit"}
	cmfAgentProfitPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFAgentProfitSlice is an alias for a slice of pointers to CMFAgentProfit.
	// This should generally be used opposed to []CMFAgentProfit.
	CMFAgentProfitSlice []*CMFAgentProfit
	// CMFAgentProfitHook is the signature for custom CMFAgentProfit hook methods
	CMFAgentProfitHook func(context.Context, boil.ContextExecutor, *CMFAgentProfit) error

	cmfAgentProfitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfAgentProfitType                 = reflect.TypeOf(&CMFAgentProfit{})
	cmfAgentProfitMapping              = queries.MakeStructMapping(cmfAgentProfitType)
	cmfAgentProfitPrimaryKeyMapping, _ = queries.BindMapping(cmfAgentProfitType, cmfAgentProfitMapping, cmfAgentProfitPrimaryKeyColumns)
	cmfAgentProfitInsertCacheMut       sync.RWMutex
	cmfAgentProfitInsertCache          = make(map[string]insertCache)
	cmfAgentProfitUpdateCacheMut       sync.RWMutex
	cmfAgentProfitUpdateCache          = make(map[string]updateCache)
	cmfAgentProfitUpsertCacheMut       sync.RWMutex
	cmfAgentProfitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfAgentProfitBeforeInsertHooks []CMFAgentProfitHook
var cmfAgentProfitBeforeUpdateHooks []CMFAgentProfitHook
var cmfAgentProfitBeforeDeleteHooks []CMFAgentProfitHook
var cmfAgentProfitBeforeUpsertHooks []CMFAgentProfitHook

var cmfAgentProfitAfterInsertHooks []CMFAgentProfitHook
var cmfAgentProfitAfterSelectHooks []CMFAgentProfitHook
var cmfAgentProfitAfterUpdateHooks []CMFAgentProfitHook
var cmfAgentProfitAfterDeleteHooks []CMFAgentProfitHook
var cmfAgentProfitAfterUpsertHooks []CMFAgentProfitHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFAgentProfit) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFAgentProfit) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFAgentProfit) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFAgentProfit) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFAgentProfit) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFAgentProfit) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFAgentProfit) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFAgentProfit) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFAgentProfit) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentProfitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFAgentProfitHook registers your hook function for all future operations.
func AddCMFAgentProfitHook(hookPoint boil.HookPoint, cmfAgentProfitHook CMFAgentProfitHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfAgentProfitBeforeInsertHooks = append(cmfAgentProfitBeforeInsertHooks, cmfAgentProfitHook)
	case boil.BeforeUpdateHook:
		cmfAgentProfitBeforeUpdateHooks = append(cmfAgentProfitBeforeUpdateHooks, cmfAgentProfitHook)
	case boil.BeforeDeleteHook:
		cmfAgentProfitBeforeDeleteHooks = append(cmfAgentProfitBeforeDeleteHooks, cmfAgentProfitHook)
	case boil.BeforeUpsertHook:
		cmfAgentProfitBeforeUpsertHooks = append(cmfAgentProfitBeforeUpsertHooks, cmfAgentProfitHook)
	case boil.AfterInsertHook:
		cmfAgentProfitAfterInsertHooks = append(cmfAgentProfitAfterInsertHooks, cmfAgentProfitHook)
	case boil.AfterSelectHook:
		cmfAgentProfitAfterSelectHooks = append(cmfAgentProfitAfterSelectHooks, cmfAgentProfitHook)
	case boil.AfterUpdateHook:
		cmfAgentProfitAfterUpdateHooks = append(cmfAgentProfitAfterUpdateHooks, cmfAgentProfitHook)
	case boil.AfterDeleteHook:
		cmfAgentProfitAfterDeleteHooks = append(cmfAgentProfitAfterDeleteHooks, cmfAgentProfitHook)
	case boil.AfterUpsertHook:
		cmfAgentProfitAfterUpsertHooks = append(cmfAgentProfitAfterUpsertHooks, cmfAgentProfitHook)
	}
}

// One returns a single cmfAgentProfit record from the query.
func (q cmfAgentProfitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFAgentProfit, error) {
	o := &CMFAgentProfit{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_agent_profit")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFAgentProfit records from the query.
func (q cmfAgentProfitQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFAgentProfitSlice, error) {
	var o []*CMFAgentProfit

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFAgentProfit slice")
	}

	if len(cmfAgentProfitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFAgentProfit records in the query.
func (q cmfAgentProfitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_agent_profit rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfAgentProfitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_agent_profit exists")
	}

	return count > 0, nil
}

// CMFAgentProfits retrieves all the records using an executor.
func CMFAgentProfits(mods ...qm.QueryMod) cmfAgentProfitQuery {
	mods = append(mods, qm.From("`cmf_agent_profit`"))
	return cmfAgentProfitQuery{NewQuery(mods...)}
}

// FindCMFAgentProfit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFAgentProfit(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFAgentProfit, error) {
	cmfAgentProfitObj := &CMFAgentProfit{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_agent_profit` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfAgentProfitObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_agent_profit")
	}

	return cmfAgentProfitObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFAgentProfit) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_agent_profit provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfAgentProfitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfAgentProfitInsertCacheMut.RLock()
	cache, cached := cmfAgentProfitInsertCache[key]
	cmfAgentProfitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfAgentProfitAllColumns,
			cmfAgentProfitColumnsWithDefault,
			cmfAgentProfitColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfAgentProfitType, cmfAgentProfitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfAgentProfitType, cmfAgentProfitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_agent_profit` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_agent_profit` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_agent_profit` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfAgentProfitPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_agent_profit")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfAgentProfitMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_agent_profit")
	}

CacheNoHooks:
	if !cached {
		cmfAgentProfitInsertCacheMut.Lock()
		cmfAgentProfitInsertCache[key] = cache
		cmfAgentProfitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFAgentProfit.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFAgentProfit) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfAgentProfitUpdateCacheMut.RLock()
	cache, cached := cmfAgentProfitUpdateCache[key]
	cmfAgentProfitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfAgentProfitAllColumns,
			cmfAgentProfitPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_agent_profit, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_agent_profit` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfAgentProfitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfAgentProfitType, cmfAgentProfitMapping, append(wl, cmfAgentProfitPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_agent_profit row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_agent_profit")
	}

	if !cached {
		cmfAgentProfitUpdateCacheMut.Lock()
		cmfAgentProfitUpdateCache[key] = cache
		cmfAgentProfitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfAgentProfitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_agent_profit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_agent_profit")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFAgentProfitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAgentProfitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_agent_profit` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAgentProfitPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfAgentProfit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfAgentProfit")
	}
	return rowsAff, nil
}

var mySQLCMFAgentProfitUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFAgentProfit) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_agent_profit provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfAgentProfitColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFAgentProfitUniqueColumns, o)

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

	cmfAgentProfitUpsertCacheMut.RLock()
	cache, cached := cmfAgentProfitUpsertCache[key]
	cmfAgentProfitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfAgentProfitAllColumns,
			cmfAgentProfitColumnsWithDefault,
			cmfAgentProfitColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfAgentProfitAllColumns,
			cmfAgentProfitPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_agent_profit, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_agent_profit`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_agent_profit` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfAgentProfitType, cmfAgentProfitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfAgentProfitType, cmfAgentProfitMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_agent_profit")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfAgentProfitMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfAgentProfitType, cmfAgentProfitMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_agent_profit")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_agent_profit")
	}

CacheNoHooks:
	if !cached {
		cmfAgentProfitUpsertCacheMut.Lock()
		cmfAgentProfitUpsertCache[key] = cache
		cmfAgentProfitUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFAgentProfit record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFAgentProfit) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFAgentProfit provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfAgentProfitPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_agent_profit` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_agent_profit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_agent_profit")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfAgentProfitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfAgentProfitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_agent_profit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_agent_profit")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFAgentProfitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfAgentProfitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAgentProfitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_agent_profit` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAgentProfitPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfAgentProfit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_agent_profit")
	}

	if len(cmfAgentProfitAfterDeleteHooks) != 0 {
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
func (o *CMFAgentProfit) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFAgentProfit(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFAgentProfitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFAgentProfitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAgentProfitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_agent_profit`.* FROM `cmf_agent_profit` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAgentProfitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFAgentProfitSlice")
	}

	*o = slice

	return nil
}

// CMFAgentProfitExists checks if the CMFAgentProfit row exists.
func CMFAgentProfitExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_agent_profit` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_agent_profit exists")
	}

	return exists, nil
}
