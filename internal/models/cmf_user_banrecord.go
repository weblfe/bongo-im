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

// CMFUserBanrecord is an object representing the database table.
type CMFUserBanrecord struct {
	ID        uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	BanReason null.String `boil:"ban_reason" json:"ban_reason,omitempty" toml:"ban_reason" yaml:"ban_reason,omitempty"`
	BanLong   null.Int    `boil:"ban_long" json:"ban_long,omitempty" toml:"ban_long" yaml:"ban_long,omitempty"`
	UID       null.Int    `boil:"uid" json:"uid,omitempty" toml:"uid" yaml:"uid,omitempty"`
	Addtime   null.Int    `boil:"addtime" json:"addtime,omitempty" toml:"addtime" yaml:"addtime,omitempty"`
	EndTime   null.Int    `boil:"end_time" json:"end_time,omitempty" toml:"end_time" yaml:"end_time,omitempty"`

	R *cmfUserBanrecordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfUserBanrecordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFUserBanrecordColumns = struct {
	ID        string
	BanReason string
	BanLong   string
	UID       string
	Addtime   string
	EndTime   string
}{
	ID:        "id",
	BanReason: "ban_reason",
	BanLong:   "ban_long",
	UID:       "uid",
	Addtime:   "addtime",
	EndTime:   "end_time",
}

// Generated where

var CMFUserBanrecordWhere = struct {
	ID        whereHelperuint
	BanReason whereHelpernull_String
	BanLong   whereHelpernull_Int
	UID       whereHelpernull_Int
	Addtime   whereHelpernull_Int
	EndTime   whereHelpernull_Int
}{
	ID:        whereHelperuint{field: "`cmf_user_banrecord`.`id`"},
	BanReason: whereHelpernull_String{field: "`cmf_user_banrecord`.`ban_reason`"},
	BanLong:   whereHelpernull_Int{field: "`cmf_user_banrecord`.`ban_long`"},
	UID:       whereHelpernull_Int{field: "`cmf_user_banrecord`.`uid`"},
	Addtime:   whereHelpernull_Int{field: "`cmf_user_banrecord`.`addtime`"},
	EndTime:   whereHelpernull_Int{field: "`cmf_user_banrecord`.`end_time`"},
}

// CMFUserBanrecordRels is where relationship names are stored.
var CMFUserBanrecordRels = struct {
}{}

// cmfUserBanrecordR is where relationships are stored.
type cmfUserBanrecordR struct {
}

// NewStruct creates a new relationship struct
func (*cmfUserBanrecordR) NewStruct() *cmfUserBanrecordR {
	return &cmfUserBanrecordR{}
}

// cmfUserBanrecordL is where Load methods for each relationship are stored.
type cmfUserBanrecordL struct{}

var (
	cmfUserBanrecordAllColumns            = []string{"id", "ban_reason", "ban_long", "uid", "addtime", "end_time"}
	cmfUserBanrecordColumnsWithoutDefault = []string{"ban_reason"}
	cmfUserBanrecordColumnsWithDefault    = []string{"id", "ban_long", "uid", "addtime", "end_time"}
	cmfUserBanrecordPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFUserBanrecordSlice is an alias for a slice of pointers to CMFUserBanrecord.
	// This should generally be used opposed to []CMFUserBanrecord.
	CMFUserBanrecordSlice []*CMFUserBanrecord
	// CMFUserBanrecordHook is the signature for custom CMFUserBanrecord hook methods
	CMFUserBanrecordHook func(context.Context, boil.ContextExecutor, *CMFUserBanrecord) error

	cmfUserBanrecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfUserBanrecordType                 = reflect.TypeOf(&CMFUserBanrecord{})
	cmfUserBanrecordMapping              = queries.MakeStructMapping(cmfUserBanrecordType)
	cmfUserBanrecordPrimaryKeyMapping, _ = queries.BindMapping(cmfUserBanrecordType, cmfUserBanrecordMapping, cmfUserBanrecordPrimaryKeyColumns)
	cmfUserBanrecordInsertCacheMut       sync.RWMutex
	cmfUserBanrecordInsertCache          = make(map[string]insertCache)
	cmfUserBanrecordUpdateCacheMut       sync.RWMutex
	cmfUserBanrecordUpdateCache          = make(map[string]updateCache)
	cmfUserBanrecordUpsertCacheMut       sync.RWMutex
	cmfUserBanrecordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfUserBanrecordBeforeInsertHooks []CMFUserBanrecordHook
var cmfUserBanrecordBeforeUpdateHooks []CMFUserBanrecordHook
var cmfUserBanrecordBeforeDeleteHooks []CMFUserBanrecordHook
var cmfUserBanrecordBeforeUpsertHooks []CMFUserBanrecordHook

var cmfUserBanrecordAfterInsertHooks []CMFUserBanrecordHook
var cmfUserBanrecordAfterSelectHooks []CMFUserBanrecordHook
var cmfUserBanrecordAfterUpdateHooks []CMFUserBanrecordHook
var cmfUserBanrecordAfterDeleteHooks []CMFUserBanrecordHook
var cmfUserBanrecordAfterUpsertHooks []CMFUserBanrecordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFUserBanrecord) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFUserBanrecord) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFUserBanrecord) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFUserBanrecord) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFUserBanrecord) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFUserBanrecord) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFUserBanrecord) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFUserBanrecord) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFUserBanrecord) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfUserBanrecordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFUserBanrecordHook registers your hook function for all future operations.
func AddCMFUserBanrecordHook(hookPoint boil.HookPoint, cmfUserBanrecordHook CMFUserBanrecordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfUserBanrecordBeforeInsertHooks = append(cmfUserBanrecordBeforeInsertHooks, cmfUserBanrecordHook)
	case boil.BeforeUpdateHook:
		cmfUserBanrecordBeforeUpdateHooks = append(cmfUserBanrecordBeforeUpdateHooks, cmfUserBanrecordHook)
	case boil.BeforeDeleteHook:
		cmfUserBanrecordBeforeDeleteHooks = append(cmfUserBanrecordBeforeDeleteHooks, cmfUserBanrecordHook)
	case boil.BeforeUpsertHook:
		cmfUserBanrecordBeforeUpsertHooks = append(cmfUserBanrecordBeforeUpsertHooks, cmfUserBanrecordHook)
	case boil.AfterInsertHook:
		cmfUserBanrecordAfterInsertHooks = append(cmfUserBanrecordAfterInsertHooks, cmfUserBanrecordHook)
	case boil.AfterSelectHook:
		cmfUserBanrecordAfterSelectHooks = append(cmfUserBanrecordAfterSelectHooks, cmfUserBanrecordHook)
	case boil.AfterUpdateHook:
		cmfUserBanrecordAfterUpdateHooks = append(cmfUserBanrecordAfterUpdateHooks, cmfUserBanrecordHook)
	case boil.AfterDeleteHook:
		cmfUserBanrecordAfterDeleteHooks = append(cmfUserBanrecordAfterDeleteHooks, cmfUserBanrecordHook)
	case boil.AfterUpsertHook:
		cmfUserBanrecordAfterUpsertHooks = append(cmfUserBanrecordAfterUpsertHooks, cmfUserBanrecordHook)
	}
}

// One returns a single cmfUserBanrecord record from the query.
func (q cmfUserBanrecordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFUserBanrecord, error) {
	o := &CMFUserBanrecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_user_banrecord")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFUserBanrecord records from the query.
func (q cmfUserBanrecordQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFUserBanrecordSlice, error) {
	var o []*CMFUserBanrecord

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFUserBanrecord slice")
	}

	if len(cmfUserBanrecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFUserBanrecord records in the query.
func (q cmfUserBanrecordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_user_banrecord rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfUserBanrecordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_user_banrecord exists")
	}

	return count > 0, nil
}

// CMFUserBanrecords retrieves all the records using an executor.
func CMFUserBanrecords(mods ...qm.QueryMod) cmfUserBanrecordQuery {
	mods = append(mods, qm.From("`cmf_user_banrecord`"))
	return cmfUserBanrecordQuery{NewQuery(mods...)}
}

// FindCMFUserBanrecord retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFUserBanrecord(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CMFUserBanrecord, error) {
	cmfUserBanrecordObj := &CMFUserBanrecord{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_user_banrecord` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfUserBanrecordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_user_banrecord")
	}

	return cmfUserBanrecordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFUserBanrecord) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_banrecord provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserBanrecordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfUserBanrecordInsertCacheMut.RLock()
	cache, cached := cmfUserBanrecordInsertCache[key]
	cmfUserBanrecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfUserBanrecordAllColumns,
			cmfUserBanrecordColumnsWithDefault,
			cmfUserBanrecordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserBanrecordType, cmfUserBanrecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfUserBanrecordType, cmfUserBanrecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_user_banrecord` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_user_banrecord` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_user_banrecord` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfUserBanrecordPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_user_banrecord")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfUserBanrecordMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_banrecord")
	}

CacheNoHooks:
	if !cached {
		cmfUserBanrecordInsertCacheMut.Lock()
		cmfUserBanrecordInsertCache[key] = cache
		cmfUserBanrecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFUserBanrecord.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFUserBanrecord) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfUserBanrecordUpdateCacheMut.RLock()
	cache, cached := cmfUserBanrecordUpdateCache[key]
	cmfUserBanrecordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfUserBanrecordAllColumns,
			cmfUserBanrecordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_user_banrecord, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_user_banrecord` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfUserBanrecordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfUserBanrecordType, cmfUserBanrecordMapping, append(wl, cmfUserBanrecordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_user_banrecord row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_user_banrecord")
	}

	if !cached {
		cmfUserBanrecordUpdateCacheMut.Lock()
		cmfUserBanrecordUpdateCache[key] = cache
		cmfUserBanrecordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfUserBanrecordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_user_banrecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_user_banrecord")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFUserBanrecordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserBanrecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_user_banrecord` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserBanrecordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfUserBanrecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfUserBanrecord")
	}
	return rowsAff, nil
}

var mySQLCMFUserBanrecordUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFUserBanrecord) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_user_banrecord provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfUserBanrecordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFUserBanrecordUniqueColumns, o)

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

	cmfUserBanrecordUpsertCacheMut.RLock()
	cache, cached := cmfUserBanrecordUpsertCache[key]
	cmfUserBanrecordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfUserBanrecordAllColumns,
			cmfUserBanrecordColumnsWithDefault,
			cmfUserBanrecordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfUserBanrecordAllColumns,
			cmfUserBanrecordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_user_banrecord, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_user_banrecord`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_user_banrecord` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfUserBanrecordType, cmfUserBanrecordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfUserBanrecordType, cmfUserBanrecordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_user_banrecord")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfUserBanrecordMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfUserBanrecordType, cmfUserBanrecordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_user_banrecord")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_user_banrecord")
	}

CacheNoHooks:
	if !cached {
		cmfUserBanrecordUpsertCacheMut.Lock()
		cmfUserBanrecordUpsertCache[key] = cache
		cmfUserBanrecordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFUserBanrecord record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFUserBanrecord) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFUserBanrecord provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfUserBanrecordPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_user_banrecord` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_user_banrecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_user_banrecord")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfUserBanrecordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfUserBanrecordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_user_banrecord")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_banrecord")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFUserBanrecordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfUserBanrecordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserBanrecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_user_banrecord` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserBanrecordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfUserBanrecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_user_banrecord")
	}

	if len(cmfUserBanrecordAfterDeleteHooks) != 0 {
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
func (o *CMFUserBanrecord) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFUserBanrecord(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFUserBanrecordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFUserBanrecordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfUserBanrecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_user_banrecord`.* FROM `cmf_user_banrecord` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfUserBanrecordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFUserBanrecordSlice")
	}

	*o = slice

	return nil
}

// CMFUserBanrecordExists checks if the CMFUserBanrecord row exists.
func CMFUserBanrecordExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_user_banrecord` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_user_banrecord exists")
	}

	return exists, nil
}
