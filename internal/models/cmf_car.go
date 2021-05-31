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

// CMFCar is an object representing the database table.
type CMFCar struct {
	ID        int           `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string        `boil:"name" json:"name" toml:"name" yaml:"name"`
	Thumb     string        `boil:"thumb" json:"thumb" toml:"thumb" yaml:"thumb"`
	SWF       string        `boil:"swf" json:"swf" toml:"swf" yaml:"swf"`
	Swftime   types.Decimal `boil:"swftime" json:"swftime" toml:"swftime" yaml:"swftime"`
	Needcoin  int           `boil:"needcoin" json:"needcoin" toml:"needcoin" yaml:"needcoin"`
	Score     int           `boil:"score" json:"score" toml:"score" yaml:"score"`
	ListOrder int           `boil:"list_order" json:"list_order" toml:"list_order" yaml:"list_order"`
	Addtime   int           `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Words     string        `boil:"words" json:"words" toml:"words" yaml:"words"`
	Uptime    int           `boil:"uptime" json:"uptime" toml:"uptime" yaml:"uptime"`

	R *cmfCarR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfCarL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFCarColumns = struct {
	ID        string
	Name      string
	Thumb     string
	SWF       string
	Swftime   string
	Needcoin  string
	Score     string
	ListOrder string
	Addtime   string
	Words     string
	Uptime    string
}{
	ID:        "id",
	Name:      "name",
	Thumb:     "thumb",
	SWF:       "swf",
	Swftime:   "swftime",
	Needcoin:  "needcoin",
	Score:     "score",
	ListOrder: "list_order",
	Addtime:   "addtime",
	Words:     "words",
	Uptime:    "uptime",
}

// Generated where

var CMFCarWhere = struct {
	ID        whereHelperint
	Name      whereHelperstring
	Thumb     whereHelperstring
	SWF       whereHelperstring
	Swftime   whereHelpertypes_Decimal
	Needcoin  whereHelperint
	Score     whereHelperint
	ListOrder whereHelperint
	Addtime   whereHelperint
	Words     whereHelperstring
	Uptime    whereHelperint
}{
	ID:        whereHelperint{field: "`cmf_car`.`id`"},
	Name:      whereHelperstring{field: "`cmf_car`.`name`"},
	Thumb:     whereHelperstring{field: "`cmf_car`.`thumb`"},
	SWF:       whereHelperstring{field: "`cmf_car`.`swf`"},
	Swftime:   whereHelpertypes_Decimal{field: "`cmf_car`.`swftime`"},
	Needcoin:  whereHelperint{field: "`cmf_car`.`needcoin`"},
	Score:     whereHelperint{field: "`cmf_car`.`score`"},
	ListOrder: whereHelperint{field: "`cmf_car`.`list_order`"},
	Addtime:   whereHelperint{field: "`cmf_car`.`addtime`"},
	Words:     whereHelperstring{field: "`cmf_car`.`words`"},
	Uptime:    whereHelperint{field: "`cmf_car`.`uptime`"},
}

// CMFCarRels is where relationship names are stored.
var CMFCarRels = struct {
}{}

// cmfCarR is where relationships are stored.
type cmfCarR struct {
}

// NewStruct creates a new relationship struct
func (*cmfCarR) NewStruct() *cmfCarR {
	return &cmfCarR{}
}

// cmfCarL is where Load methods for each relationship are stored.
type cmfCarL struct{}

var (
	cmfCarAllColumns            = []string{"id", "name", "thumb", "swf", "swftime", "needcoin", "score", "list_order", "addtime", "words", "uptime"}
	cmfCarColumnsWithoutDefault = []string{"name", "thumb", "swf", "words"}
	cmfCarColumnsWithDefault    = []string{"id", "swftime", "needcoin", "score", "list_order", "addtime", "uptime"}
	cmfCarPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFCarSlice is an alias for a slice of pointers to CMFCar.
	// This should generally be used opposed to []CMFCar.
	CMFCarSlice []*CMFCar
	// CMFCarHook is the signature for custom CMFCar hook methods
	CMFCarHook func(context.Context, boil.ContextExecutor, *CMFCar) error

	cmfCarQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfCarType                 = reflect.TypeOf(&CMFCar{})
	cmfCarMapping              = queries.MakeStructMapping(cmfCarType)
	cmfCarPrimaryKeyMapping, _ = queries.BindMapping(cmfCarType, cmfCarMapping, cmfCarPrimaryKeyColumns)
	cmfCarInsertCacheMut       sync.RWMutex
	cmfCarInsertCache          = make(map[string]insertCache)
	cmfCarUpdateCacheMut       sync.RWMutex
	cmfCarUpdateCache          = make(map[string]updateCache)
	cmfCarUpsertCacheMut       sync.RWMutex
	cmfCarUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfCarBeforeInsertHooks []CMFCarHook
var cmfCarBeforeUpdateHooks []CMFCarHook
var cmfCarBeforeDeleteHooks []CMFCarHook
var cmfCarBeforeUpsertHooks []CMFCarHook

var cmfCarAfterInsertHooks []CMFCarHook
var cmfCarAfterSelectHooks []CMFCarHook
var cmfCarAfterUpdateHooks []CMFCarHook
var cmfCarAfterDeleteHooks []CMFCarHook
var cmfCarAfterUpsertHooks []CMFCarHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFCar) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFCar) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFCar) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFCar) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFCar) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFCar) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFCar) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFCar) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFCar) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfCarAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFCarHook registers your hook function for all future operations.
func AddCMFCarHook(hookPoint boil.HookPoint, cmfCarHook CMFCarHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfCarBeforeInsertHooks = append(cmfCarBeforeInsertHooks, cmfCarHook)
	case boil.BeforeUpdateHook:
		cmfCarBeforeUpdateHooks = append(cmfCarBeforeUpdateHooks, cmfCarHook)
	case boil.BeforeDeleteHook:
		cmfCarBeforeDeleteHooks = append(cmfCarBeforeDeleteHooks, cmfCarHook)
	case boil.BeforeUpsertHook:
		cmfCarBeforeUpsertHooks = append(cmfCarBeforeUpsertHooks, cmfCarHook)
	case boil.AfterInsertHook:
		cmfCarAfterInsertHooks = append(cmfCarAfterInsertHooks, cmfCarHook)
	case boil.AfterSelectHook:
		cmfCarAfterSelectHooks = append(cmfCarAfterSelectHooks, cmfCarHook)
	case boil.AfterUpdateHook:
		cmfCarAfterUpdateHooks = append(cmfCarAfterUpdateHooks, cmfCarHook)
	case boil.AfterDeleteHook:
		cmfCarAfterDeleteHooks = append(cmfCarAfterDeleteHooks, cmfCarHook)
	case boil.AfterUpsertHook:
		cmfCarAfterUpsertHooks = append(cmfCarAfterUpsertHooks, cmfCarHook)
	}
}

// One returns a single cmfCar record from the query.
func (q cmfCarQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFCar, error) {
	o := &CMFCar{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_car")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFCar records from the query.
func (q cmfCarQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFCarSlice, error) {
	var o []*CMFCar

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFCar slice")
	}

	if len(cmfCarAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFCar records in the query.
func (q cmfCarQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_car rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfCarQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_car exists")
	}

	return count > 0, nil
}

// CMFCars retrieves all the records using an executor.
func CMFCars(mods ...qm.QueryMod) cmfCarQuery {
	mods = append(mods, qm.From("`cmf_car`"))
	return cmfCarQuery{NewQuery(mods...)}
}

// FindCMFCar retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFCar(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFCar, error) {
	cmfCarObj := &CMFCar{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_car` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfCarObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_car")
	}

	return cmfCarObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFCar) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_car provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfCarColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfCarInsertCacheMut.RLock()
	cache, cached := cmfCarInsertCache[key]
	cmfCarInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfCarAllColumns,
			cmfCarColumnsWithDefault,
			cmfCarColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfCarType, cmfCarMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfCarType, cmfCarMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_car` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_car` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_car` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfCarPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_car")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfCarMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_car")
	}

CacheNoHooks:
	if !cached {
		cmfCarInsertCacheMut.Lock()
		cmfCarInsertCache[key] = cache
		cmfCarInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFCar.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFCar) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfCarUpdateCacheMut.RLock()
	cache, cached := cmfCarUpdateCache[key]
	cmfCarUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfCarAllColumns,
			cmfCarPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_car, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_car` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfCarPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfCarType, cmfCarMapping, append(wl, cmfCarPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_car row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_car")
	}

	if !cached {
		cmfCarUpdateCacheMut.Lock()
		cmfCarUpdateCache[key] = cache
		cmfCarUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfCarQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_car")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_car")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFCarSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfCarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_car` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfCarPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfCar slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfCar")
	}
	return rowsAff, nil
}

var mySQLCMFCarUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFCar) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_car provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfCarColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFCarUniqueColumns, o)

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

	cmfCarUpsertCacheMut.RLock()
	cache, cached := cmfCarUpsertCache[key]
	cmfCarUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfCarAllColumns,
			cmfCarColumnsWithDefault,
			cmfCarColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfCarAllColumns,
			cmfCarPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_car, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_car`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_car` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfCarType, cmfCarMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfCarType, cmfCarMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_car")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfCarMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfCarType, cmfCarMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_car")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_car")
	}

CacheNoHooks:
	if !cached {
		cmfCarUpsertCacheMut.Lock()
		cmfCarUpsertCache[key] = cache
		cmfCarUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFCar record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFCar) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFCar provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfCarPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_car` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_car")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_car")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfCarQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfCarQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_car")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_car")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFCarSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfCarBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfCarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_car` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfCarPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfCar slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_car")
	}

	if len(cmfCarAfterDeleteHooks) != 0 {
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
func (o *CMFCar) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFCar(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFCarSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFCarSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfCarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_car`.* FROM `cmf_car` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfCarPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFCarSlice")
	}

	*o = slice

	return nil
}

// CMFCarExists checks if the CMFCar row exists.
func CMFCarExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_car` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_car exists")
	}

	return exists, nil
}