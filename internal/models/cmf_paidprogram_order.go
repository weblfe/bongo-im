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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// CMFPaidprogramOrder is an object representing the database table.
type CMFPaidprogramOrder struct {
	ID       int64         `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID      int64         `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Touid    int64         `boil:"touid" json:"touid" toml:"touid" yaml:"touid"`
	ObjectID int64         `boil:"object_id" json:"object_id" toml:"object_id" yaml:"object_id"`
	Type     bool          `boil:"type" json:"type" toml:"type" yaml:"type"`
	Status   bool          `boil:"status" json:"status" toml:"status" yaml:"status"`
	Orderno  string        `boil:"orderno" json:"orderno" toml:"orderno" yaml:"orderno"`
	TradeNo  string        `boil:"trade_no" json:"trade_no" toml:"trade_no" yaml:"trade_no"`
	Money    types.Decimal `boil:"money" json:"money" toml:"money" yaml:"money"`
	Addtime  int           `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Edittime int           `boil:"edittime" json:"edittime" toml:"edittime" yaml:"edittime"`
	Isdel    bool          `boil:"isdel" json:"isdel" toml:"isdel" yaml:"isdel"`

	R *cmfPaidprogramOrderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfPaidprogramOrderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFPaidprogramOrderColumns = struct {
	ID       string
	UID      string
	Touid    string
	ObjectID string
	Type     string
	Status   string
	Orderno  string
	TradeNo  string
	Money    string
	Addtime  string
	Edittime string
	Isdel    string
}{
	ID:       "id",
	UID:      "uid",
	Touid:    "touid",
	ObjectID: "object_id",
	Type:     "type",
	Status:   "status",
	Orderno:  "orderno",
	TradeNo:  "trade_no",
	Money:    "money",
	Addtime:  "addtime",
	Edittime: "edittime",
	Isdel:    "isdel",
}

// Generated where

var CMFPaidprogramOrderWhere = struct {
	ID       whereHelperint64
	UID      whereHelperint64
	Touid    whereHelperint64
	ObjectID whereHelperint64
	Type     whereHelperbool
	Status   whereHelperbool
	Orderno  whereHelperstring
	TradeNo  whereHelperstring
	Money    whereHelpertypes_Decimal
	Addtime  whereHelperint
	Edittime whereHelperint
	Isdel    whereHelperbool
}{
	ID:       whereHelperint64{field: "`cmf_paidprogram_order`.`id`"},
	UID:      whereHelperint64{field: "`cmf_paidprogram_order`.`uid`"},
	Touid:    whereHelperint64{field: "`cmf_paidprogram_order`.`touid`"},
	ObjectID: whereHelperint64{field: "`cmf_paidprogram_order`.`object_id`"},
	Type:     whereHelperbool{field: "`cmf_paidprogram_order`.`type`"},
	Status:   whereHelperbool{field: "`cmf_paidprogram_order`.`status`"},
	Orderno:  whereHelperstring{field: "`cmf_paidprogram_order`.`orderno`"},
	TradeNo:  whereHelperstring{field: "`cmf_paidprogram_order`.`trade_no`"},
	Money:    whereHelpertypes_Decimal{field: "`cmf_paidprogram_order`.`money`"},
	Addtime:  whereHelperint{field: "`cmf_paidprogram_order`.`addtime`"},
	Edittime: whereHelperint{field: "`cmf_paidprogram_order`.`edittime`"},
	Isdel:    whereHelperbool{field: "`cmf_paidprogram_order`.`isdel`"},
}

// CMFPaidprogramOrderRels is where relationship names are stored.
var CMFPaidprogramOrderRels = struct {
}{}

// cmfPaidprogramOrderR is where relationships are stored.
type cmfPaidprogramOrderR struct {
}

// NewStruct creates a new relationship struct
func (*cmfPaidprogramOrderR) NewStruct() *cmfPaidprogramOrderR {
	return &cmfPaidprogramOrderR{}
}

// cmfPaidprogramOrderL is where Load methods for each relationship are stored.
type cmfPaidprogramOrderL struct{}

var (
	cmfPaidprogramOrderAllColumns            = []string{"id", "uid", "touid", "object_id", "type", "status", "orderno", "trade_no", "money", "addtime", "edittime", "isdel"}
	cmfPaidprogramOrderColumnsWithoutDefault = []string{"type", "status", "orderno", "trade_no"}
	cmfPaidprogramOrderColumnsWithDefault    = []string{"id", "uid", "touid", "object_id", "money", "addtime", "edittime", "isdel"}
	cmfPaidprogramOrderPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFPaidprogramOrderSlice is an alias for a slice of pointers to CMFPaidprogramOrder.
	// This should generally be used opposed to []CMFPaidprogramOrder.
	CMFPaidprogramOrderSlice []*CMFPaidprogramOrder
	// CMFPaidprogramOrderHook is the signature for custom CMFPaidprogramOrder hook methods
	CMFPaidprogramOrderHook func(context.Context, boil.ContextExecutor, *CMFPaidprogramOrder) error

	cmfPaidprogramOrderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfPaidprogramOrderType                 = reflect.TypeOf(&CMFPaidprogramOrder{})
	cmfPaidprogramOrderMapping              = queries.MakeStructMapping(cmfPaidprogramOrderType)
	cmfPaidprogramOrderPrimaryKeyMapping, _ = queries.BindMapping(cmfPaidprogramOrderType, cmfPaidprogramOrderMapping, cmfPaidprogramOrderPrimaryKeyColumns)
	cmfPaidprogramOrderInsertCacheMut       sync.RWMutex
	cmfPaidprogramOrderInsertCache          = make(map[string]insertCache)
	cmfPaidprogramOrderUpdateCacheMut       sync.RWMutex
	cmfPaidprogramOrderUpdateCache          = make(map[string]updateCache)
	cmfPaidprogramOrderUpsertCacheMut       sync.RWMutex
	cmfPaidprogramOrderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfPaidprogramOrderBeforeInsertHooks []CMFPaidprogramOrderHook
var cmfPaidprogramOrderBeforeUpdateHooks []CMFPaidprogramOrderHook
var cmfPaidprogramOrderBeforeDeleteHooks []CMFPaidprogramOrderHook
var cmfPaidprogramOrderBeforeUpsertHooks []CMFPaidprogramOrderHook

var cmfPaidprogramOrderAfterInsertHooks []CMFPaidprogramOrderHook
var cmfPaidprogramOrderAfterSelectHooks []CMFPaidprogramOrderHook
var cmfPaidprogramOrderAfterUpdateHooks []CMFPaidprogramOrderHook
var cmfPaidprogramOrderAfterDeleteHooks []CMFPaidprogramOrderHook
var cmfPaidprogramOrderAfterUpsertHooks []CMFPaidprogramOrderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFPaidprogramOrder) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFPaidprogramOrder) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFPaidprogramOrder) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFPaidprogramOrder) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFPaidprogramOrder) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFPaidprogramOrder) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFPaidprogramOrder) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFPaidprogramOrder) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFPaidprogramOrder) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramOrderAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFPaidprogramOrderHook registers your hook function for all future operations.
func AddCMFPaidprogramOrderHook(hookPoint boil.HookPoint, cmfPaidprogramOrderHook CMFPaidprogramOrderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfPaidprogramOrderBeforeInsertHooks = append(cmfPaidprogramOrderBeforeInsertHooks, cmfPaidprogramOrderHook)
	case boil.BeforeUpdateHook:
		cmfPaidprogramOrderBeforeUpdateHooks = append(cmfPaidprogramOrderBeforeUpdateHooks, cmfPaidprogramOrderHook)
	case boil.BeforeDeleteHook:
		cmfPaidprogramOrderBeforeDeleteHooks = append(cmfPaidprogramOrderBeforeDeleteHooks, cmfPaidprogramOrderHook)
	case boil.BeforeUpsertHook:
		cmfPaidprogramOrderBeforeUpsertHooks = append(cmfPaidprogramOrderBeforeUpsertHooks, cmfPaidprogramOrderHook)
	case boil.AfterInsertHook:
		cmfPaidprogramOrderAfterInsertHooks = append(cmfPaidprogramOrderAfterInsertHooks, cmfPaidprogramOrderHook)
	case boil.AfterSelectHook:
		cmfPaidprogramOrderAfterSelectHooks = append(cmfPaidprogramOrderAfterSelectHooks, cmfPaidprogramOrderHook)
	case boil.AfterUpdateHook:
		cmfPaidprogramOrderAfterUpdateHooks = append(cmfPaidprogramOrderAfterUpdateHooks, cmfPaidprogramOrderHook)
	case boil.AfterDeleteHook:
		cmfPaidprogramOrderAfterDeleteHooks = append(cmfPaidprogramOrderAfterDeleteHooks, cmfPaidprogramOrderHook)
	case boil.AfterUpsertHook:
		cmfPaidprogramOrderAfterUpsertHooks = append(cmfPaidprogramOrderAfterUpsertHooks, cmfPaidprogramOrderHook)
	}
}

// One returns a single cmfPaidprogramOrder record from the query.
func (q cmfPaidprogramOrderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFPaidprogramOrder, error) {
	o := &CMFPaidprogramOrder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_paidprogram_order")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFPaidprogramOrder records from the query.
func (q cmfPaidprogramOrderQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFPaidprogramOrderSlice, error) {
	var o []*CMFPaidprogramOrder

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFPaidprogramOrder slice")
	}

	if len(cmfPaidprogramOrderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFPaidprogramOrder records in the query.
func (q cmfPaidprogramOrderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_paidprogram_order rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfPaidprogramOrderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_paidprogram_order exists")
	}

	return count > 0, nil
}

// CMFPaidprogramOrders retrieves all the records using an executor.
func CMFPaidprogramOrders(mods ...qm.QueryMod) cmfPaidprogramOrderQuery {
	mods = append(mods, qm.From("`cmf_paidprogram_order`"))
	return cmfPaidprogramOrderQuery{NewQuery(mods...)}
}

// FindCMFPaidprogramOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFPaidprogramOrder(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CMFPaidprogramOrder, error) {
	cmfPaidprogramOrderObj := &CMFPaidprogramOrder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_paidprogram_order` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfPaidprogramOrderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_paidprogram_order")
	}

	return cmfPaidprogramOrderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFPaidprogramOrder) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_paidprogram_order provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPaidprogramOrderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfPaidprogramOrderInsertCacheMut.RLock()
	cache, cached := cmfPaidprogramOrderInsertCache[key]
	cmfPaidprogramOrderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfPaidprogramOrderAllColumns,
			cmfPaidprogramOrderColumnsWithDefault,
			cmfPaidprogramOrderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfPaidprogramOrderType, cmfPaidprogramOrderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfPaidprogramOrderType, cmfPaidprogramOrderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_paidprogram_order` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_paidprogram_order` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_paidprogram_order` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfPaidprogramOrderPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_paidprogram_order")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPaidprogramOrderMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_paidprogram_order")
	}

CacheNoHooks:
	if !cached {
		cmfPaidprogramOrderInsertCacheMut.Lock()
		cmfPaidprogramOrderInsertCache[key] = cache
		cmfPaidprogramOrderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFPaidprogramOrder.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFPaidprogramOrder) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfPaidprogramOrderUpdateCacheMut.RLock()
	cache, cached := cmfPaidprogramOrderUpdateCache[key]
	cmfPaidprogramOrderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfPaidprogramOrderAllColumns,
			cmfPaidprogramOrderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_paidprogram_order, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_paidprogram_order` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfPaidprogramOrderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfPaidprogramOrderType, cmfPaidprogramOrderMapping, append(wl, cmfPaidprogramOrderPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_paidprogram_order row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_paidprogram_order")
	}

	if !cached {
		cmfPaidprogramOrderUpdateCacheMut.Lock()
		cmfPaidprogramOrderUpdateCache[key] = cache
		cmfPaidprogramOrderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfPaidprogramOrderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_paidprogram_order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_paidprogram_order")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFPaidprogramOrderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPaidprogramOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_paidprogram_order` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPaidprogramOrderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfPaidprogramOrder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfPaidprogramOrder")
	}
	return rowsAff, nil
}

var mySQLCMFPaidprogramOrderUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFPaidprogramOrder) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_paidprogram_order provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPaidprogramOrderColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFPaidprogramOrderUniqueColumns, o)

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

	cmfPaidprogramOrderUpsertCacheMut.RLock()
	cache, cached := cmfPaidprogramOrderUpsertCache[key]
	cmfPaidprogramOrderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfPaidprogramOrderAllColumns,
			cmfPaidprogramOrderColumnsWithDefault,
			cmfPaidprogramOrderColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfPaidprogramOrderAllColumns,
			cmfPaidprogramOrderPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_paidprogram_order, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_paidprogram_order`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_paidprogram_order` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfPaidprogramOrderType, cmfPaidprogramOrderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfPaidprogramOrderType, cmfPaidprogramOrderMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_paidprogram_order")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPaidprogramOrderMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfPaidprogramOrderType, cmfPaidprogramOrderMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_paidprogram_order")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_paidprogram_order")
	}

CacheNoHooks:
	if !cached {
		cmfPaidprogramOrderUpsertCacheMut.Lock()
		cmfPaidprogramOrderUpsertCache[key] = cache
		cmfPaidprogramOrderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFPaidprogramOrder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFPaidprogramOrder) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFPaidprogramOrder provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfPaidprogramOrderPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_paidprogram_order` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_paidprogram_order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_paidprogram_order")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfPaidprogramOrderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfPaidprogramOrderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_paidprogram_order")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_paidprogram_order")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFPaidprogramOrderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfPaidprogramOrderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPaidprogramOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_paidprogram_order` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPaidprogramOrderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfPaidprogramOrder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_paidprogram_order")
	}

	if len(cmfPaidprogramOrderAfterDeleteHooks) != 0 {
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
func (o *CMFPaidprogramOrder) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFPaidprogramOrder(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFPaidprogramOrderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFPaidprogramOrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPaidprogramOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_paidprogram_order`.* FROM `cmf_paidprogram_order` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPaidprogramOrderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFPaidprogramOrderSlice")
	}

	*o = slice

	return nil
}

// CMFPaidprogramOrderExists checks if the CMFPaidprogramOrder row exists.
func CMFPaidprogramOrderExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_paidprogram_order` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_paidprogram_order exists")
	}

	return exists, nil
}
