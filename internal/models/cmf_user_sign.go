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

// CMFUserSign is an object representing the database table.
type CMFUserSign struct {
	UID       int `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	BonusDay  int `boil:"bonus_day" json:"bonus_day" toml:"bonus_day" yaml:"bonus_day"`
	BonusTime int `boil:"bonus_time" json:"bonus_time" toml:"bonus_time" yaml:"bonus_time"`
	CountDay  int `boil:"count_day" json:"count_day" toml:"count_day" yaml:"count_day"`

	R *cmfUserSignR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfUserSignL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFUserSignColumns = struct {
	UID       string
	BonusDay  string
	BonusTime string
	CountDay  string
}{
	UID:       "uid",
	BonusDay:  "bonus_day",
	BonusTime: "bonus_time",
	CountDay:  "count_day",
}

// Generated where

var CMFUserSignWhere = struct {
	UID       whereHelperint
	BonusDay  whereHelperint
	BonusTime whereHelperint
	CountDay  whereHelperint
}{
	UID:       whereHelperint{field: "`cmf_user_sign`.`uid`"},
	BonusDay:  whereHelperint{field: "`cmf_user_sign`.`bonus_day`"},
	BonusTime: whereHelperint{field: "`cmf_user_sign`.`bonus_time`"},
	CountDay:  whereHelperint{field: "`cmf_user_sign`.`count_day`"},
}

// CMFUserSignRels is where relationship names are stored.
var CMFUserSignRels = struct {
}{}

// cmfUserSignR is where relationships are stored.
type cmfUserSignR struct {
}

// NewStruct creates a new relationship struct
func (*cmfUserSignR) NewStruct() *cmfUserSignR {
	return &cmfUserSignR{}
}

// cmfUserSignL is where Load methods for each relationship are stored.
type cmfUserSignL struct{}

var (
	cmfUserSignAllColumns            = []string{"uid", "bonus_day", "bonus_time", "count_day"}
	cmfUserSignColumnsWithoutDefault = []string{"uid"}
	cmfUserSignColumnsWithDefault    = []string{"bonus_day", "bonus_time", "count_day"}
	cmfUserSignPrimaryKeyColumns     = []string{"uid"}
)

type (
	// CMFUserSignSlice is an alias for a slice of pointers to CMFUserSign.
	// This should generally be used opposed to []CMFUserSign.
	CMFUserSignSlice []*CMFUserSign
	// CMFUserSignHook is the signature for custom CMFUserSign hook methods
	CMFUserSignHook func(context.Context, boil.ContextExecutor, *CMFUserSign) error

	cmfUserSignQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfUserSignType                 = reflect.TypeOf(&CMFUserSign{})
	cmfUserSignMapping              = queries.MakeStructMapping(cmfUserSignType)
	cmfUserSignPrimaryKeyMapping, _ = queries.BindMapping(cmfUserSignType, cmfUserSignMapping, cmfUserSignPrimaryKeyColumns)
	cmfUserSignInsertCacheMut       sync.RWMutex
	cmfUserSignInsertCache          = make(map[string]insertCache)
	cmfUserSignUpdateCacheMut       sync.RWMutex
	cmfUserSignUpdateCache          = make(map[string]updateCache)
	cmfUserSignUpsertCacheMut       sync.RWMutex
	cmfUserSignUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfUserSignBeforeInsertHooks []CMFUserSignHook
var cmfUserSignBeforeUpdateHooks []CMFUserSignHook
var cmfUserSignBeforeDeleteHooks []CMFUserSignHook
var cmfUserSignBeforeUpsertHooks []CMFUserSignHook

var cmfUserSignAfterInsertHooks []CMFUserSignHook
var cmfUserSignAfterSelectHooks []CMFUserSignHook
var cmfUserSignAfterUpdateHooks []CMFUserSignHook
var cmfUserSignAfterDeleteHooks []CMFUserSignHook
var cmfUserSignAfterUpsertHooks []CMFUserSignHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFUserSign) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFUserSign) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFUserSign) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFUserSign) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFUserSign) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFUserSign) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFUserSign) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFUserSign) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFUserSign) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserSignAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFUserSignHook registers your hook function for all future operations.
func AddCMFUserSignHook(hookPoint boil.HookPoint, cmfUserSignHook CMFUserSignHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfUserSignBeforeInsertHooks = append(cmfUserSignBeforeInsertHooks, cmfUserSignHook)
	case boil.BeforeUpdateHook:
		cmfUserSignBeforeUpdateHooks = append(cmfUserSignBeforeUpdateHooks, cmfUserSignHook)
	case boil.BeforeDeleteHook:
		cmfUserSignBeforeDeleteHooks = append(cmfUserSignBeforeDeleteHooks, cmfUserSignHook)
	case boil.BeforeUpsertHook:
		cmfUserSignBeforeUpsertHooks = append(cmfUserSignBeforeUpsertHooks, cmfUserSignHook)
	case boil.AfterInsertHook:
		cmfUserSignAfterInsertHooks = append(cmfUserSignAfterInsertHooks, cmfUserSignHook)
	case boil.AfterSelectHook:
		cmfUserSignAfterSelectHooks = append(cmfUserSignAfterSelectHooks, cmfUserSignHook)
	case boil.AfterUpdateHook:
		cmfUserSignAfterUpdateHooks = append(cmfUserSignAfterUpdateHooks, cmfUserSignHook)
	case boil.AfterDeleteHook:
		cmfUserSignAfterDeleteHooks = append(cmfUserSignAfterDeleteHooks, cmfUserSignHook)
	case boil.AfterUpsertHook:
		cmfUserSignAfterUpsertHooks = append(cmfUserSignAfterUpsertHooks, cmfUserSignHook)
	}
}

// One returns a single cmfUserSign record from the query.
func (q cmfUserSignQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFUserSign, error) {
	o := &CMFUserSign{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_user_sign")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFUserSign records from the query.
func (q cmfUserSignQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFUserSignSlice, error) {
	var o []*CMFUserSign

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFUserSign slice")
	}

	if len(cmfUserSignAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFUserSign records in the query.
func (q cmfUserSignQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_user_sign rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfUserSignQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_user_sign exists")
	}

	return count > 0, nil
}

// CMFUserSigns retrieves all the records using an executor.
func CMFUserSigns(mods ...qm.QueryMod) cmfUserSignQuery {
	mods = append(mods, qm.From("`cmf_user_sign`"))
	return cmfUserSignQuery{NewQuery(mods...)}
}

// FindCMFUserSign retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFUserSign(ctx context.Context, exec boil.ContextExecutor, uID int, selectCols ...string) (*CMFUserSign, error) {
	cmfUserSignObj := &CMFUserSign{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_user_sign` where `uid`=?", sel,
	)

	q := queries.Raw(query, uID)

	err := q.Bind(ctx, exec, cmfUserSignObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_user_sign")
	}

	return cmfUserSignObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFUserSign) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_sign provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserSignColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfUserSignInsertCacheMut.RLock()
	cache, cached := cmfUserSignInsertCache[key]
	cmfUserSignInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfUserSignAllColumns,
			cmfUserSignColumnsWithDefault,
			cmfUserSignColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserSignType, cmfUserSignMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfUserSignType, cmfUserSignMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_user_sign` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_user_sign` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_user_sign` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfUserSignPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_user_sign")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.UID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_sign")
	}

CacheNoHooks:
	if !cached {
		cmfUserSignInsertCacheMut.Lock()
		cmfUserSignInsertCache[key] = cache
		cmfUserSignInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFUserSign.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFUserSign) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfUserSignUpdateCacheMut.RLock()
	cache, cached := cmfUserSignUpdateCache[key]
	cmfUserSignUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfUserSignAllColumns,
			cmfUserSignPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_user_sign, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_user_sign` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfUserSignPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfUserSignType, cmfUserSignMapping, append(wl, cmfUserSignPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_user_sign row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_user_sign")
	}

	if !cached {
		cmfUserSignUpdateCacheMut.Lock()
		cmfUserSignUpdateCache[key] = cache
		cmfUserSignUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfUserSignQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_user_sign")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_user_sign")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFUserSignSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserSignPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_user_sign` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserSignPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfUserSign slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfUserSign")
	}
	return rowsAff, nil
}

var mySQLCMFUserSignUniqueColumns = []string{
	"uid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFUserSign) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_sign provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserSignColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFUserSignUniqueColumns, o)

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

	cmfUserSignUpsertCacheMut.RLock()
	cache, cached := cmfUserSignUpsertCache[key]
	cmfUserSignUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfUserSignAllColumns,
			cmfUserSignColumnsWithDefault,
			cmfUserSignColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfUserSignAllColumns,
			cmfUserSignPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_user_sign, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_user_sign`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_user_sign` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserSignType, cmfUserSignMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfUserSignType, cmfUserSignMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_user_sign")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfUserSignType, cmfUserSignMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_user_sign")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_sign")
	}

CacheNoHooks:
	if !cached {
		cmfUserSignUpsertCacheMut.Lock()
		cmfUserSignUpsertCache[key] = cache
		cmfUserSignUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFUserSign record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFUserSign) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFUserSign provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfUserSignPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_user_sign` WHERE `uid`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_user_sign")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_user_sign")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfUserSignQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfUserSignQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_user_sign")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_sign")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFUserSignSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfUserSignBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserSignPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_user_sign` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserSignPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfUserSign slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_sign")
	}

	if len(cmfUserSignAfterDeleteHooks) != 0 {
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
func (o *CMFUserSign) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFUserSign(ctx, exec, o.UID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFUserSignSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFUserSignSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserSignPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_user_sign`.* FROM `cmf_user_sign` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserSignPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFUserSignSlice")
	}

	*o = slice

	return nil
}

// CMFUserSignExists checks if the CMFUserSign row exists.
func CMFUserSignExists(ctx context.Context, exec boil.ContextExecutor, uID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_user_sign` where `uid`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, uID)
	}
	row := exec.QueryRowContext(ctx, sql, uID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_user_sign exists")
	}

	return exists, nil
}
