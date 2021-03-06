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

// CMFLink is an object representing the database table.
type CMFLink struct {
	ID          uint64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Status      uint8   `boil:"status" json:"status" toml:"status" yaml:"status"`
	Rating      int     `boil:"rating" json:"rating" toml:"rating" yaml:"rating"`
	ListOrder   float32 `boil:"list_order" json:"list_order" toml:"list_order" yaml:"list_order"`
	Description string  `boil:"description" json:"description" toml:"description" yaml:"description"`
	URL         string  `boil:"url" json:"url" toml:"url" yaml:"url"`
	Name        string  `boil:"name" json:"name" toml:"name" yaml:"name"`
	Image       string  `boil:"image" json:"image" toml:"image" yaml:"image"`
	Target      string  `boil:"target" json:"target" toml:"target" yaml:"target"`
	Rel         string  `boil:"rel" json:"rel" toml:"rel" yaml:"rel"`

	R *cmfLinkR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfLinkL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFLinkColumns = struct {
	ID          string
	Status      string
	Rating      string
	ListOrder   string
	Description string
	URL         string
	Name        string
	Image       string
	Target      string
	Rel         string
}{
	ID:          "id",
	Status:      "status",
	Rating:      "rating",
	ListOrder:   "list_order",
	Description: "description",
	URL:         "url",
	Name:        "name",
	Image:       "image",
	Target:      "target",
	Rel:         "rel",
}

// Generated where

var CMFLinkWhere = struct {
	ID          whereHelperuint64
	Status      whereHelperuint8
	Rating      whereHelperint
	ListOrder   whereHelperfloat32
	Description whereHelperstring
	URL         whereHelperstring
	Name        whereHelperstring
	Image       whereHelperstring
	Target      whereHelperstring
	Rel         whereHelperstring
}{
	ID:          whereHelperuint64{field: "`cmf_link`.`id`"},
	Status:      whereHelperuint8{field: "`cmf_link`.`status`"},
	Rating:      whereHelperint{field: "`cmf_link`.`rating`"},
	ListOrder:   whereHelperfloat32{field: "`cmf_link`.`list_order`"},
	Description: whereHelperstring{field: "`cmf_link`.`description`"},
	URL:         whereHelperstring{field: "`cmf_link`.`url`"},
	Name:        whereHelperstring{field: "`cmf_link`.`name`"},
	Image:       whereHelperstring{field: "`cmf_link`.`image`"},
	Target:      whereHelperstring{field: "`cmf_link`.`target`"},
	Rel:         whereHelperstring{field: "`cmf_link`.`rel`"},
}

// CMFLinkRels is where relationship names are stored.
var CMFLinkRels = struct {
}{}

// cmfLinkR is where relationships are stored.
type cmfLinkR struct {
}

// NewStruct creates a new relationship struct
func (*cmfLinkR) NewStruct() *cmfLinkR {
	return &cmfLinkR{}
}

// cmfLinkL is where Load methods for each relationship are stored.
type cmfLinkL struct{}

var (
	cmfLinkAllColumns            = []string{"id", "status", "rating", "list_order", "description", "url", "name", "image", "target", "rel"}
	cmfLinkColumnsWithoutDefault = []string{"description", "url", "name", "image", "target", "rel"}
	cmfLinkColumnsWithDefault    = []string{"id", "status", "rating", "list_order"}
	cmfLinkPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFLinkSlice is an alias for a slice of pointers to CMFLink.
	// This should generally be used opposed to []CMFLink.
	CMFLinkSlice []*CMFLink
	// CMFLinkHook is the signature for custom CMFLink hook methods
	CMFLinkHook func(context.Context, boil.ContextExecutor, *CMFLink) error

	cmfLinkQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfLinkType                 = reflect.TypeOf(&CMFLink{})
	cmfLinkMapping              = queries.MakeStructMapping(cmfLinkType)
	cmfLinkPrimaryKeyMapping, _ = queries.BindMapping(cmfLinkType, cmfLinkMapping, cmfLinkPrimaryKeyColumns)
	cmfLinkInsertCacheMut       sync.RWMutex
	cmfLinkInsertCache          = make(map[string]insertCache)
	cmfLinkUpdateCacheMut       sync.RWMutex
	cmfLinkUpdateCache          = make(map[string]updateCache)
	cmfLinkUpsertCacheMut       sync.RWMutex
	cmfLinkUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfLinkBeforeInsertHooks []CMFLinkHook
var cmfLinkBeforeUpdateHooks []CMFLinkHook
var cmfLinkBeforeDeleteHooks []CMFLinkHook
var cmfLinkBeforeUpsertHooks []CMFLinkHook

var cmfLinkAfterInsertHooks []CMFLinkHook
var cmfLinkAfterSelectHooks []CMFLinkHook
var cmfLinkAfterUpdateHooks []CMFLinkHook
var cmfLinkAfterDeleteHooks []CMFLinkHook
var cmfLinkAfterUpsertHooks []CMFLinkHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFLink) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFLink) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFLink) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFLink) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFLink) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFLink) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFLink) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFLink) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFLink) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLinkAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFLinkHook registers your hook function for all future operations.
func AddCMFLinkHook(hookPoint boil.HookPoint, cmfLinkHook CMFLinkHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfLinkBeforeInsertHooks = append(cmfLinkBeforeInsertHooks, cmfLinkHook)
	case boil.BeforeUpdateHook:
		cmfLinkBeforeUpdateHooks = append(cmfLinkBeforeUpdateHooks, cmfLinkHook)
	case boil.BeforeDeleteHook:
		cmfLinkBeforeDeleteHooks = append(cmfLinkBeforeDeleteHooks, cmfLinkHook)
	case boil.BeforeUpsertHook:
		cmfLinkBeforeUpsertHooks = append(cmfLinkBeforeUpsertHooks, cmfLinkHook)
	case boil.AfterInsertHook:
		cmfLinkAfterInsertHooks = append(cmfLinkAfterInsertHooks, cmfLinkHook)
	case boil.AfterSelectHook:
		cmfLinkAfterSelectHooks = append(cmfLinkAfterSelectHooks, cmfLinkHook)
	case boil.AfterUpdateHook:
		cmfLinkAfterUpdateHooks = append(cmfLinkAfterUpdateHooks, cmfLinkHook)
	case boil.AfterDeleteHook:
		cmfLinkAfterDeleteHooks = append(cmfLinkAfterDeleteHooks, cmfLinkHook)
	case boil.AfterUpsertHook:
		cmfLinkAfterUpsertHooks = append(cmfLinkAfterUpsertHooks, cmfLinkHook)
	}
}

// One returns a single cmfLink record from the query.
func (q cmfLinkQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFLink, error) {
	o := &CMFLink{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_link")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFLink records from the query.
func (q cmfLinkQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFLinkSlice, error) {
	var o []*CMFLink

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFLink slice")
	}

	if len(cmfLinkAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFLink records in the query.
func (q cmfLinkQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_link rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfLinkQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_link exists")
	}

	return count > 0, nil
}

// CMFLinks retrieves all the records using an executor.
func CMFLinks(mods ...qm.QueryMod) cmfLinkQuery {
	mods = append(mods, qm.From("`cmf_link`"))
	return cmfLinkQuery{NewQuery(mods...)}
}

// FindCMFLink retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFLink(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*CMFLink, error) {
	cmfLinkObj := &CMFLink{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_link` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfLinkObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_link")
	}

	return cmfLinkObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFLink) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_link provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfLinkColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfLinkInsertCacheMut.RLock()
	cache, cached := cmfLinkInsertCache[key]
	cmfLinkInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfLinkAllColumns,
			cmfLinkColumnsWithDefault,
			cmfLinkColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfLinkType, cmfLinkMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfLinkType, cmfLinkMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_link` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_link` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_link` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfLinkPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_link")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfLinkMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_link")
	}

CacheNoHooks:
	if !cached {
		cmfLinkInsertCacheMut.Lock()
		cmfLinkInsertCache[key] = cache
		cmfLinkInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFLink.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFLink) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfLinkUpdateCacheMut.RLock()
	cache, cached := cmfLinkUpdateCache[key]
	cmfLinkUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfLinkAllColumns,
			cmfLinkPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_link, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_link` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfLinkPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfLinkType, cmfLinkMapping, append(wl, cmfLinkPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_link row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_link")
	}

	if !cached {
		cmfLinkUpdateCacheMut.Lock()
		cmfLinkUpdateCache[key] = cache
		cmfLinkUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfLinkQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_link")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_link")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFLinkSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLinkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_link` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLinkPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfLink slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfLink")
	}
	return rowsAff, nil
}

var mySQLCMFLinkUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFLink) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_link provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfLinkColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFLinkUniqueColumns, o)

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

	cmfLinkUpsertCacheMut.RLock()
	cache, cached := cmfLinkUpsertCache[key]
	cmfLinkUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfLinkAllColumns,
			cmfLinkColumnsWithDefault,
			cmfLinkColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfLinkAllColumns,
			cmfLinkPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_link, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_link`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_link` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfLinkType, cmfLinkMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfLinkType, cmfLinkMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_link")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfLinkMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfLinkType, cmfLinkMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_link")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_link")
	}

CacheNoHooks:
	if !cached {
		cmfLinkUpsertCacheMut.Lock()
		cmfLinkUpsertCache[key] = cache
		cmfLinkUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFLink record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFLink) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFLink provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfLinkPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_link` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_link")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_link")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfLinkQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfLinkQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_link")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_link")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFLinkSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfLinkBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLinkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_link` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLinkPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfLink slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_link")
	}

	if len(cmfLinkAfterDeleteHooks) != 0 {
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
func (o *CMFLink) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFLink(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFLinkSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFLinkSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLinkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_link`.* FROM `cmf_link` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLinkPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFLinkSlice")
	}

	*o = slice

	return nil
}

// CMFLinkExists checks if the CMFLink row exists.
func CMFLinkExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_link` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_link exists")
	}

	return exists, nil
}
