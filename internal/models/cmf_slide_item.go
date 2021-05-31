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

// CMFSlideItem is an object representing the database table.
type CMFSlideItem struct {
	ID          uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	SlideID     int         `boil:"slide_id" json:"slide_id" toml:"slide_id" yaml:"slide_id"`
	Status      uint8       `boil:"status" json:"status" toml:"status" yaml:"status"`
	ListOrder   float32     `boil:"list_order" json:"list_order" toml:"list_order" yaml:"list_order"`
	Title       string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Image       string      `boil:"image" json:"image" toml:"image" yaml:"image"`
	URL         string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	Target      string      `boil:"target" json:"target" toml:"target" yaml:"target"`
	Description string      `boil:"description" json:"description" toml:"description" yaml:"description"`
	Content     null.String `boil:"content" json:"content,omitempty" toml:"content" yaml:"content,omitempty"`
	More        null.String `boil:"more" json:"more,omitempty" toml:"more" yaml:"more,omitempty"`

	R *cmfSlideItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfSlideItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFSlideItemColumns = struct {
	ID          string
	SlideID     string
	Status      string
	ListOrder   string
	Title       string
	Image       string
	URL         string
	Target      string
	Description string
	Content     string
	More        string
}{
	ID:          "id",
	SlideID:     "slide_id",
	Status:      "status",
	ListOrder:   "list_order",
	Title:       "title",
	Image:       "image",
	URL:         "url",
	Target:      "target",
	Description: "description",
	Content:     "content",
	More:        "more",
}

// Generated where

var CMFSlideItemWhere = struct {
	ID          whereHelperuint
	SlideID     whereHelperint
	Status      whereHelperuint8
	ListOrder   whereHelperfloat32
	Title       whereHelperstring
	Image       whereHelperstring
	URL         whereHelperstring
	Target      whereHelperstring
	Description whereHelperstring
	Content     whereHelpernull_String
	More        whereHelpernull_String
}{
	ID:          whereHelperuint{field: "`cmf_slide_item`.`id`"},
	SlideID:     whereHelperint{field: "`cmf_slide_item`.`slide_id`"},
	Status:      whereHelperuint8{field: "`cmf_slide_item`.`status`"},
	ListOrder:   whereHelperfloat32{field: "`cmf_slide_item`.`list_order`"},
	Title:       whereHelperstring{field: "`cmf_slide_item`.`title`"},
	Image:       whereHelperstring{field: "`cmf_slide_item`.`image`"},
	URL:         whereHelperstring{field: "`cmf_slide_item`.`url`"},
	Target:      whereHelperstring{field: "`cmf_slide_item`.`target`"},
	Description: whereHelperstring{field: "`cmf_slide_item`.`description`"},
	Content:     whereHelpernull_String{field: "`cmf_slide_item`.`content`"},
	More:        whereHelpernull_String{field: "`cmf_slide_item`.`more`"},
}

// CMFSlideItemRels is where relationship names are stored.
var CMFSlideItemRels = struct {
}{}

// cmfSlideItemR is where relationships are stored.
type cmfSlideItemR struct {
}

// NewStruct creates a new relationship struct
func (*cmfSlideItemR) NewStruct() *cmfSlideItemR {
	return &cmfSlideItemR{}
}

// cmfSlideItemL is where Load methods for each relationship are stored.
type cmfSlideItemL struct{}

var (
	cmfSlideItemAllColumns            = []string{"id", "slide_id", "status", "list_order", "title", "image", "url", "target", "description", "content", "more"}
	cmfSlideItemColumnsWithoutDefault = []string{"title", "image", "url", "target", "description", "content", "more"}
	cmfSlideItemColumnsWithDefault    = []string{"id", "slide_id", "status", "list_order"}
	cmfSlideItemPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFSlideItemSlice is an alias for a slice of pointers to CMFSlideItem.
	// This should generally be used opposed to []CMFSlideItem.
	CMFSlideItemSlice []*CMFSlideItem
	// CMFSlideItemHook is the signature for custom CMFSlideItem hook methods
	CMFSlideItemHook func(context.Context, boil.ContextExecutor, *CMFSlideItem) error

	cmfSlideItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfSlideItemType                 = reflect.TypeOf(&CMFSlideItem{})
	cmfSlideItemMapping              = queries.MakeStructMapping(cmfSlideItemType)
	cmfSlideItemPrimaryKeyMapping, _ = queries.BindMapping(cmfSlideItemType, cmfSlideItemMapping, cmfSlideItemPrimaryKeyColumns)
	cmfSlideItemInsertCacheMut       sync.RWMutex
	cmfSlideItemInsertCache          = make(map[string]insertCache)
	cmfSlideItemUpdateCacheMut       sync.RWMutex
	cmfSlideItemUpdateCache          = make(map[string]updateCache)
	cmfSlideItemUpsertCacheMut       sync.RWMutex
	cmfSlideItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfSlideItemBeforeInsertHooks []CMFSlideItemHook
var cmfSlideItemBeforeUpdateHooks []CMFSlideItemHook
var cmfSlideItemBeforeDeleteHooks []CMFSlideItemHook
var cmfSlideItemBeforeUpsertHooks []CMFSlideItemHook

var cmfSlideItemAfterInsertHooks []CMFSlideItemHook
var cmfSlideItemAfterSelectHooks []CMFSlideItemHook
var cmfSlideItemAfterUpdateHooks []CMFSlideItemHook
var cmfSlideItemAfterDeleteHooks []CMFSlideItemHook
var cmfSlideItemAfterUpsertHooks []CMFSlideItemHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFSlideItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFSlideItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFSlideItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFSlideItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFSlideItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFSlideItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFSlideItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFSlideItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFSlideItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfSlideItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFSlideItemHook registers your hook function for all future operations.
func AddCMFSlideItemHook(hookPoint boil.HookPoint, cmfSlideItemHook CMFSlideItemHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfSlideItemBeforeInsertHooks = append(cmfSlideItemBeforeInsertHooks, cmfSlideItemHook)
	case boil.BeforeUpdateHook:
		cmfSlideItemBeforeUpdateHooks = append(cmfSlideItemBeforeUpdateHooks, cmfSlideItemHook)
	case boil.BeforeDeleteHook:
		cmfSlideItemBeforeDeleteHooks = append(cmfSlideItemBeforeDeleteHooks, cmfSlideItemHook)
	case boil.BeforeUpsertHook:
		cmfSlideItemBeforeUpsertHooks = append(cmfSlideItemBeforeUpsertHooks, cmfSlideItemHook)
	case boil.AfterInsertHook:
		cmfSlideItemAfterInsertHooks = append(cmfSlideItemAfterInsertHooks, cmfSlideItemHook)
	case boil.AfterSelectHook:
		cmfSlideItemAfterSelectHooks = append(cmfSlideItemAfterSelectHooks, cmfSlideItemHook)
	case boil.AfterUpdateHook:
		cmfSlideItemAfterUpdateHooks = append(cmfSlideItemAfterUpdateHooks, cmfSlideItemHook)
	case boil.AfterDeleteHook:
		cmfSlideItemAfterDeleteHooks = append(cmfSlideItemAfterDeleteHooks, cmfSlideItemHook)
	case boil.AfterUpsertHook:
		cmfSlideItemAfterUpsertHooks = append(cmfSlideItemAfterUpsertHooks, cmfSlideItemHook)
	}
}

// One returns a single cmfSlideItem record from the query.
func (q cmfSlideItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFSlideItem, error) {
	o := &CMFSlideItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_slide_item")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFSlideItem records from the query.
func (q cmfSlideItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFSlideItemSlice, error) {
	var o []*CMFSlideItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFSlideItem slice")
	}

	if len(cmfSlideItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFSlideItem records in the query.
func (q cmfSlideItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_slide_item rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfSlideItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_slide_item exists")
	}

	return count > 0, nil
}

// CMFSlideItems retrieves all the records using an executor.
func CMFSlideItems(mods ...qm.QueryMod) cmfSlideItemQuery {
	mods = append(mods, qm.From("`cmf_slide_item`"))
	return cmfSlideItemQuery{NewQuery(mods...)}
}

// FindCMFSlideItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFSlideItem(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CMFSlideItem, error) {
	cmfSlideItemObj := &CMFSlideItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_slide_item` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfSlideItemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_slide_item")
	}

	return cmfSlideItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFSlideItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_slide_item provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfSlideItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfSlideItemInsertCacheMut.RLock()
	cache, cached := cmfSlideItemInsertCache[key]
	cmfSlideItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfSlideItemAllColumns,
			cmfSlideItemColumnsWithDefault,
			cmfSlideItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfSlideItemType, cmfSlideItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfSlideItemType, cmfSlideItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_slide_item` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_slide_item` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_slide_item` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfSlideItemPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_slide_item")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfSlideItemMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_slide_item")
	}

CacheNoHooks:
	if !cached {
		cmfSlideItemInsertCacheMut.Lock()
		cmfSlideItemInsertCache[key] = cache
		cmfSlideItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFSlideItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFSlideItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfSlideItemUpdateCacheMut.RLock()
	cache, cached := cmfSlideItemUpdateCache[key]
	cmfSlideItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfSlideItemAllColumns,
			cmfSlideItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_slide_item, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_slide_item` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfSlideItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfSlideItemType, cmfSlideItemMapping, append(wl, cmfSlideItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_slide_item row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_slide_item")
	}

	if !cached {
		cmfSlideItemUpdateCacheMut.Lock()
		cmfSlideItemUpdateCache[key] = cache
		cmfSlideItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfSlideItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_slide_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_slide_item")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFSlideItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfSlideItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_slide_item` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfSlideItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfSlideItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfSlideItem")
	}
	return rowsAff, nil
}

var mySQLCMFSlideItemUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFSlideItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_slide_item provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfSlideItemColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFSlideItemUniqueColumns, o)

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

	cmfSlideItemUpsertCacheMut.RLock()
	cache, cached := cmfSlideItemUpsertCache[key]
	cmfSlideItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfSlideItemAllColumns,
			cmfSlideItemColumnsWithDefault,
			cmfSlideItemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfSlideItemAllColumns,
			cmfSlideItemPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_slide_item, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_slide_item`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_slide_item` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfSlideItemType, cmfSlideItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfSlideItemType, cmfSlideItemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_slide_item")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfSlideItemMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfSlideItemType, cmfSlideItemMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_slide_item")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_slide_item")
	}

CacheNoHooks:
	if !cached {
		cmfSlideItemUpsertCacheMut.Lock()
		cmfSlideItemUpsertCache[key] = cache
		cmfSlideItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFSlideItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFSlideItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFSlideItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfSlideItemPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_slide_item` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_slide_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_slide_item")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfSlideItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfSlideItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_slide_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_slide_item")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFSlideItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfSlideItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfSlideItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_slide_item` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfSlideItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfSlideItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_slide_item")
	}

	if len(cmfSlideItemAfterDeleteHooks) != 0 {
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
func (o *CMFSlideItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFSlideItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFSlideItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFSlideItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfSlideItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_slide_item`.* FROM `cmf_slide_item` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfSlideItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFSlideItemSlice")
	}

	*o = slice

	return nil
}

// CMFSlideItemExists checks if the CMFSlideItem row exists.
func CMFSlideItemExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_slide_item` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_slide_item exists")
	}

	return exists, nil
}