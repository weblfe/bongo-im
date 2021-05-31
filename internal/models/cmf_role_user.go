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

// CMFRoleUser is an object representing the database table.
type CMFRoleUser struct {
	ID     uint64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	RoleID uint   `boil:"role_id" json:"role_id" toml:"role_id" yaml:"role_id"`
	UserID uint64 `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`

	R *cmfRoleUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfRoleUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFRoleUserColumns = struct {
	ID     string
	RoleID string
	UserID string
}{
	ID:     "id",
	RoleID: "role_id",
	UserID: "user_id",
}

// Generated where

var CMFRoleUserWhere = struct {
	ID     whereHelperuint64
	RoleID whereHelperuint
	UserID whereHelperuint64
}{
	ID:     whereHelperuint64{field: "`cmf_role_user`.`id`"},
	RoleID: whereHelperuint{field: "`cmf_role_user`.`role_id`"},
	UserID: whereHelperuint64{field: "`cmf_role_user`.`user_id`"},
}

// CMFRoleUserRels is where relationship names are stored.
var CMFRoleUserRels = struct {
}{}

// cmfRoleUserR is where relationships are stored.
type cmfRoleUserR struct {
}

// NewStruct creates a new relationship struct
func (*cmfRoleUserR) NewStruct() *cmfRoleUserR {
	return &cmfRoleUserR{}
}

// cmfRoleUserL is where Load methods for each relationship are stored.
type cmfRoleUserL struct{}

var (
	cmfRoleUserAllColumns            = []string{"id", "role_id", "user_id"}
	cmfRoleUserColumnsWithoutDefault = []string{}
	cmfRoleUserColumnsWithDefault    = []string{"id", "role_id", "user_id"}
	cmfRoleUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFRoleUserSlice is an alias for a slice of pointers to CMFRoleUser.
	// This should generally be used opposed to []CMFRoleUser.
	CMFRoleUserSlice []*CMFRoleUser
	// CMFRoleUserHook is the signature for custom CMFRoleUser hook methods
	CMFRoleUserHook func(context.Context, boil.ContextExecutor, *CMFRoleUser) error

	cmfRoleUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfRoleUserType                 = reflect.TypeOf(&CMFRoleUser{})
	cmfRoleUserMapping              = queries.MakeStructMapping(cmfRoleUserType)
	cmfRoleUserPrimaryKeyMapping, _ = queries.BindMapping(cmfRoleUserType, cmfRoleUserMapping, cmfRoleUserPrimaryKeyColumns)
	cmfRoleUserInsertCacheMut       sync.RWMutex
	cmfRoleUserInsertCache          = make(map[string]insertCache)
	cmfRoleUserUpdateCacheMut       sync.RWMutex
	cmfRoleUserUpdateCache          = make(map[string]updateCache)
	cmfRoleUserUpsertCacheMut       sync.RWMutex
	cmfRoleUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfRoleUserBeforeInsertHooks []CMFRoleUserHook
var cmfRoleUserBeforeUpdateHooks []CMFRoleUserHook
var cmfRoleUserBeforeDeleteHooks []CMFRoleUserHook
var cmfRoleUserBeforeUpsertHooks []CMFRoleUserHook

var cmfRoleUserAfterInsertHooks []CMFRoleUserHook
var cmfRoleUserAfterSelectHooks []CMFRoleUserHook
var cmfRoleUserAfterUpdateHooks []CMFRoleUserHook
var cmfRoleUserAfterDeleteHooks []CMFRoleUserHook
var cmfRoleUserAfterUpsertHooks []CMFRoleUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFRoleUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFRoleUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFRoleUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFRoleUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFRoleUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFRoleUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFRoleUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFRoleUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFRoleUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfRoleUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFRoleUserHook registers your hook function for all future operations.
func AddCMFRoleUserHook(hookPoint boil.HookPoint, cmfRoleUserHook CMFRoleUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfRoleUserBeforeInsertHooks = append(cmfRoleUserBeforeInsertHooks, cmfRoleUserHook)
	case boil.BeforeUpdateHook:
		cmfRoleUserBeforeUpdateHooks = append(cmfRoleUserBeforeUpdateHooks, cmfRoleUserHook)
	case boil.BeforeDeleteHook:
		cmfRoleUserBeforeDeleteHooks = append(cmfRoleUserBeforeDeleteHooks, cmfRoleUserHook)
	case boil.BeforeUpsertHook:
		cmfRoleUserBeforeUpsertHooks = append(cmfRoleUserBeforeUpsertHooks, cmfRoleUserHook)
	case boil.AfterInsertHook:
		cmfRoleUserAfterInsertHooks = append(cmfRoleUserAfterInsertHooks, cmfRoleUserHook)
	case boil.AfterSelectHook:
		cmfRoleUserAfterSelectHooks = append(cmfRoleUserAfterSelectHooks, cmfRoleUserHook)
	case boil.AfterUpdateHook:
		cmfRoleUserAfterUpdateHooks = append(cmfRoleUserAfterUpdateHooks, cmfRoleUserHook)
	case boil.AfterDeleteHook:
		cmfRoleUserAfterDeleteHooks = append(cmfRoleUserAfterDeleteHooks, cmfRoleUserHook)
	case boil.AfterUpsertHook:
		cmfRoleUserAfterUpsertHooks = append(cmfRoleUserAfterUpsertHooks, cmfRoleUserHook)
	}
}

// One returns a single cmfRoleUser record from the query.
func (q cmfRoleUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFRoleUser, error) {
	o := &CMFRoleUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_role_user")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFRoleUser records from the query.
func (q cmfRoleUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFRoleUserSlice, error) {
	var o []*CMFRoleUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFRoleUser slice")
	}

	if len(cmfRoleUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFRoleUser records in the query.
func (q cmfRoleUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_role_user rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfRoleUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_role_user exists")
	}

	return count > 0, nil
}

// CMFRoleUsers retrieves all the records using an executor.
func CMFRoleUsers(mods ...qm.QueryMod) cmfRoleUserQuery {
	mods = append(mods, qm.From("`cmf_role_user`"))
	return cmfRoleUserQuery{NewQuery(mods...)}
}

// FindCMFRoleUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFRoleUser(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*CMFRoleUser, error) {
	cmfRoleUserObj := &CMFRoleUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_role_user` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfRoleUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_role_user")
	}

	return cmfRoleUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFRoleUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_role_user provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfRoleUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfRoleUserInsertCacheMut.RLock()
	cache, cached := cmfRoleUserInsertCache[key]
	cmfRoleUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfRoleUserAllColumns,
			cmfRoleUserColumnsWithDefault,
			cmfRoleUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfRoleUserType, cmfRoleUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfRoleUserType, cmfRoleUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_role_user` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_role_user` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_role_user` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfRoleUserPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_role_user")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfRoleUserMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_role_user")
	}

CacheNoHooks:
	if !cached {
		cmfRoleUserInsertCacheMut.Lock()
		cmfRoleUserInsertCache[key] = cache
		cmfRoleUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFRoleUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFRoleUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfRoleUserUpdateCacheMut.RLock()
	cache, cached := cmfRoleUserUpdateCache[key]
	cmfRoleUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfRoleUserAllColumns,
			cmfRoleUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_role_user, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_role_user` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfRoleUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfRoleUserType, cmfRoleUserMapping, append(wl, cmfRoleUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_role_user row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_role_user")
	}

	if !cached {
		cmfRoleUserUpdateCacheMut.Lock()
		cmfRoleUserUpdateCache[key] = cache
		cmfRoleUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfRoleUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_role_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_role_user")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFRoleUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfRoleUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_role_user` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfRoleUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfRoleUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfRoleUser")
	}
	return rowsAff, nil
}

var mySQLCMFRoleUserUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFRoleUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_role_user provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfRoleUserColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFRoleUserUniqueColumns, o)

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

	cmfRoleUserUpsertCacheMut.RLock()
	cache, cached := cmfRoleUserUpsertCache[key]
	cmfRoleUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfRoleUserAllColumns,
			cmfRoleUserColumnsWithDefault,
			cmfRoleUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfRoleUserAllColumns,
			cmfRoleUserPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_role_user, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_role_user`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_role_user` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfRoleUserType, cmfRoleUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfRoleUserType, cmfRoleUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_role_user")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfRoleUserMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfRoleUserType, cmfRoleUserMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_role_user")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_role_user")
	}

CacheNoHooks:
	if !cached {
		cmfRoleUserUpsertCacheMut.Lock()
		cmfRoleUserUpsertCache[key] = cache
		cmfRoleUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFRoleUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFRoleUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFRoleUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfRoleUserPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_role_user` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_role_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_role_user")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfRoleUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfRoleUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_role_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_role_user")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFRoleUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfRoleUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfRoleUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_role_user` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfRoleUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfRoleUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_role_user")
	}

	if len(cmfRoleUserAfterDeleteHooks) != 0 {
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
func (o *CMFRoleUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFRoleUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFRoleUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFRoleUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfRoleUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_role_user`.* FROM `cmf_role_user` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfRoleUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFRoleUserSlice")
	}

	*o = slice

	return nil
}

// CMFRoleUserExists checks if the CMFRoleUser row exists.
func CMFRoleUserExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_role_user` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_role_user exists")
	}

	return exists, nil
}
