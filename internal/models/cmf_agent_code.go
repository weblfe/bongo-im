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

// CMFAgentCode is an object representing the database table.
type CMFAgentCode struct {
	UID  int    `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Code string `boil:"code" json:"code" toml:"code" yaml:"code"`

	R *cmfAgentCodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfAgentCodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFAgentCodeColumns = struct {
	UID  string
	Code string
}{
	UID:  "uid",
	Code: "code",
}

// Generated where

var CMFAgentCodeWhere = struct {
	UID  whereHelperint
	Code whereHelperstring
}{
	UID:  whereHelperint{field: "`cmf_agent_code`.`uid`"},
	Code: whereHelperstring{field: "`cmf_agent_code`.`code`"},
}

// CMFAgentCodeRels is where relationship names are stored.
var CMFAgentCodeRels = struct {
}{}

// cmfAgentCodeR is where relationships are stored.
type cmfAgentCodeR struct {
}

// NewStruct creates a new relationship struct
func (*cmfAgentCodeR) NewStruct() *cmfAgentCodeR {
	return &cmfAgentCodeR{}
}

// cmfAgentCodeL is where Load methods for each relationship are stored.
type cmfAgentCodeL struct{}

var (
	cmfAgentCodeAllColumns            = []string{"uid", "code"}
	cmfAgentCodeColumnsWithoutDefault = []string{"code"}
	cmfAgentCodeColumnsWithDefault    = []string{"uid"}
	cmfAgentCodePrimaryKeyColumns     = []string{"uid"}
)

type (
	// CMFAgentCodeSlice is an alias for a slice of pointers to CMFAgentCode.
	// This should generally be used opposed to []CMFAgentCode.
	CMFAgentCodeSlice []*CMFAgentCode
	// CMFAgentCodeHook is the signature for custom CMFAgentCode hook methods
	CMFAgentCodeHook func(context.Context, boil.ContextExecutor, *CMFAgentCode) error

	cmfAgentCodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfAgentCodeType                 = reflect.TypeOf(&CMFAgentCode{})
	cmfAgentCodeMapping              = queries.MakeStructMapping(cmfAgentCodeType)
	cmfAgentCodePrimaryKeyMapping, _ = queries.BindMapping(cmfAgentCodeType, cmfAgentCodeMapping, cmfAgentCodePrimaryKeyColumns)
	cmfAgentCodeInsertCacheMut       sync.RWMutex
	cmfAgentCodeInsertCache          = make(map[string]insertCache)
	cmfAgentCodeUpdateCacheMut       sync.RWMutex
	cmfAgentCodeUpdateCache          = make(map[string]updateCache)
	cmfAgentCodeUpsertCacheMut       sync.RWMutex
	cmfAgentCodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfAgentCodeBeforeInsertHooks []CMFAgentCodeHook
var cmfAgentCodeBeforeUpdateHooks []CMFAgentCodeHook
var cmfAgentCodeBeforeDeleteHooks []CMFAgentCodeHook
var cmfAgentCodeBeforeUpsertHooks []CMFAgentCodeHook

var cmfAgentCodeAfterInsertHooks []CMFAgentCodeHook
var cmfAgentCodeAfterSelectHooks []CMFAgentCodeHook
var cmfAgentCodeAfterUpdateHooks []CMFAgentCodeHook
var cmfAgentCodeAfterDeleteHooks []CMFAgentCodeHook
var cmfAgentCodeAfterUpsertHooks []CMFAgentCodeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFAgentCode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFAgentCode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFAgentCode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFAgentCode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFAgentCode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFAgentCode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFAgentCode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFAgentCode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFAgentCode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAgentCodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFAgentCodeHook registers your hook function for all future operations.
func AddCMFAgentCodeHook(hookPoint boil.HookPoint, cmfAgentCodeHook CMFAgentCodeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfAgentCodeBeforeInsertHooks = append(cmfAgentCodeBeforeInsertHooks, cmfAgentCodeHook)
	case boil.BeforeUpdateHook:
		cmfAgentCodeBeforeUpdateHooks = append(cmfAgentCodeBeforeUpdateHooks, cmfAgentCodeHook)
	case boil.BeforeDeleteHook:
		cmfAgentCodeBeforeDeleteHooks = append(cmfAgentCodeBeforeDeleteHooks, cmfAgentCodeHook)
	case boil.BeforeUpsertHook:
		cmfAgentCodeBeforeUpsertHooks = append(cmfAgentCodeBeforeUpsertHooks, cmfAgentCodeHook)
	case boil.AfterInsertHook:
		cmfAgentCodeAfterInsertHooks = append(cmfAgentCodeAfterInsertHooks, cmfAgentCodeHook)
	case boil.AfterSelectHook:
		cmfAgentCodeAfterSelectHooks = append(cmfAgentCodeAfterSelectHooks, cmfAgentCodeHook)
	case boil.AfterUpdateHook:
		cmfAgentCodeAfterUpdateHooks = append(cmfAgentCodeAfterUpdateHooks, cmfAgentCodeHook)
	case boil.AfterDeleteHook:
		cmfAgentCodeAfterDeleteHooks = append(cmfAgentCodeAfterDeleteHooks, cmfAgentCodeHook)
	case boil.AfterUpsertHook:
		cmfAgentCodeAfterUpsertHooks = append(cmfAgentCodeAfterUpsertHooks, cmfAgentCodeHook)
	}
}

// One returns a single cmfAgentCode record from the query.
func (q cmfAgentCodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFAgentCode, error) {
	o := &CMFAgentCode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_agent_code")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFAgentCode records from the query.
func (q cmfAgentCodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFAgentCodeSlice, error) {
	var o []*CMFAgentCode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFAgentCode slice")
	}

	if len(cmfAgentCodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFAgentCode records in the query.
func (q cmfAgentCodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_agent_code rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfAgentCodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_agent_code exists")
	}

	return count > 0, nil
}

// CMFAgentCodes retrieves all the records using an executor.
func CMFAgentCodes(mods ...qm.QueryMod) cmfAgentCodeQuery {
	mods = append(mods, qm.From("`cmf_agent_code`"))
	return cmfAgentCodeQuery{NewQuery(mods...)}
}

// FindCMFAgentCode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFAgentCode(ctx context.Context, exec boil.ContextExecutor, uID int, selectCols ...string) (*CMFAgentCode, error) {
	cmfAgentCodeObj := &CMFAgentCode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_agent_code` where `uid`=?", sel,
	)

	q := queries.Raw(query, uID)

	err := q.Bind(ctx, exec, cmfAgentCodeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_agent_code")
	}

	return cmfAgentCodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFAgentCode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_agent_code provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfAgentCodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfAgentCodeInsertCacheMut.RLock()
	cache, cached := cmfAgentCodeInsertCache[key]
	cmfAgentCodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfAgentCodeAllColumns,
			cmfAgentCodeColumnsWithDefault,
			cmfAgentCodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfAgentCodeType, cmfAgentCodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfAgentCodeType, cmfAgentCodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_agent_code` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_agent_code` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_agent_code` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfAgentCodePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_agent_code")
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

	o.UID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfAgentCodeMapping["uid"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_agent_code")
	}

CacheNoHooks:
	if !cached {
		cmfAgentCodeInsertCacheMut.Lock()
		cmfAgentCodeInsertCache[key] = cache
		cmfAgentCodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFAgentCode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFAgentCode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfAgentCodeUpdateCacheMut.RLock()
	cache, cached := cmfAgentCodeUpdateCache[key]
	cmfAgentCodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfAgentCodeAllColumns,
			cmfAgentCodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_agent_code, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_agent_code` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfAgentCodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfAgentCodeType, cmfAgentCodeMapping, append(wl, cmfAgentCodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_agent_code row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_agent_code")
	}

	if !cached {
		cmfAgentCodeUpdateCacheMut.Lock()
		cmfAgentCodeUpdateCache[key] = cache
		cmfAgentCodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfAgentCodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_agent_code")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_agent_code")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFAgentCodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAgentCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_agent_code` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAgentCodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfAgentCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfAgentCode")
	}
	return rowsAff, nil
}

var mySQLCMFAgentCodeUniqueColumns = []string{
	"uid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFAgentCode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_agent_code provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfAgentCodeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFAgentCodeUniqueColumns, o)

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

	cmfAgentCodeUpsertCacheMut.RLock()
	cache, cached := cmfAgentCodeUpsertCache[key]
	cmfAgentCodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfAgentCodeAllColumns,
			cmfAgentCodeColumnsWithDefault,
			cmfAgentCodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfAgentCodeAllColumns,
			cmfAgentCodePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_agent_code, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_agent_code`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_agent_code` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfAgentCodeType, cmfAgentCodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfAgentCodeType, cmfAgentCodeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_agent_code")
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

	o.UID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfAgentCodeMapping["uid"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfAgentCodeType, cmfAgentCodeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_agent_code")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_agent_code")
	}

CacheNoHooks:
	if !cached {
		cmfAgentCodeUpsertCacheMut.Lock()
		cmfAgentCodeUpsertCache[key] = cache
		cmfAgentCodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFAgentCode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFAgentCode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFAgentCode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfAgentCodePrimaryKeyMapping)
	sql := "DELETE FROM `cmf_agent_code` WHERE `uid`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_agent_code")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_agent_code")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfAgentCodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfAgentCodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_agent_code")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_agent_code")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFAgentCodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfAgentCodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAgentCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_agent_code` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAgentCodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfAgentCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_agent_code")
	}

	if len(cmfAgentCodeAfterDeleteHooks) != 0 {
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
func (o *CMFAgentCode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFAgentCode(ctx, exec, o.UID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFAgentCodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFAgentCodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAgentCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_agent_code`.* FROM `cmf_agent_code` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAgentCodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFAgentCodeSlice")
	}

	*o = slice

	return nil
}

// CMFAgentCodeExists checks if the CMFAgentCode row exists.
func CMFAgentCodeExists(ctx context.Context, exec boil.ContextExecutor, uID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_agent_code` where `uid`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, uID)
	}
	row := exec.QueryRowContext(ctx, sql, uID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_agent_code exists")
	}

	return exists, nil
}
