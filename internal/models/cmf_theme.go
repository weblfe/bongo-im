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

// CMFTheme is an object representing the database table.
type CMFTheme struct {
	ID          int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreateTime  uint   `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime  uint   `boil:"update_time" json:"update_time" toml:"update_time" yaml:"update_time"`
	Status      uint8  `boil:"status" json:"status" toml:"status" yaml:"status"`
	IsCompiled  uint8  `boil:"is_compiled" json:"is_compiled" toml:"is_compiled" yaml:"is_compiled"`
	Theme       string `boil:"theme" json:"theme" toml:"theme" yaml:"theme"`
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Version     string `boil:"version" json:"version" toml:"version" yaml:"version"`
	DemoURL     string `boil:"demo_url" json:"demo_url" toml:"demo_url" yaml:"demo_url"`
	Thumbnail   string `boil:"thumbnail" json:"thumbnail" toml:"thumbnail" yaml:"thumbnail"`
	Author      string `boil:"author" json:"author" toml:"author" yaml:"author"`
	AuthorURL   string `boil:"author_url" json:"author_url" toml:"author_url" yaml:"author_url"`
	Lang        string `boil:"lang" json:"lang" toml:"lang" yaml:"lang"`
	Keywords    string `boil:"keywords" json:"keywords" toml:"keywords" yaml:"keywords"`
	Description string `boil:"description" json:"description" toml:"description" yaml:"description"`

	R *cmfThemeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfThemeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFThemeColumns = struct {
	ID          string
	CreateTime  string
	UpdateTime  string
	Status      string
	IsCompiled  string
	Theme       string
	Name        string
	Version     string
	DemoURL     string
	Thumbnail   string
	Author      string
	AuthorURL   string
	Lang        string
	Keywords    string
	Description string
}{
	ID:          "id",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	Status:      "status",
	IsCompiled:  "is_compiled",
	Theme:       "theme",
	Name:        "name",
	Version:     "version",
	DemoURL:     "demo_url",
	Thumbnail:   "thumbnail",
	Author:      "author",
	AuthorURL:   "author_url",
	Lang:        "lang",
	Keywords:    "keywords",
	Description: "description",
}

// Generated where

var CMFThemeWhere = struct {
	ID          whereHelperint
	CreateTime  whereHelperuint
	UpdateTime  whereHelperuint
	Status      whereHelperuint8
	IsCompiled  whereHelperuint8
	Theme       whereHelperstring
	Name        whereHelperstring
	Version     whereHelperstring
	DemoURL     whereHelperstring
	Thumbnail   whereHelperstring
	Author      whereHelperstring
	AuthorURL   whereHelperstring
	Lang        whereHelperstring
	Keywords    whereHelperstring
	Description whereHelperstring
}{
	ID:          whereHelperint{field: "`cmf_theme`.`id`"},
	CreateTime:  whereHelperuint{field: "`cmf_theme`.`create_time`"},
	UpdateTime:  whereHelperuint{field: "`cmf_theme`.`update_time`"},
	Status:      whereHelperuint8{field: "`cmf_theme`.`status`"},
	IsCompiled:  whereHelperuint8{field: "`cmf_theme`.`is_compiled`"},
	Theme:       whereHelperstring{field: "`cmf_theme`.`theme`"},
	Name:        whereHelperstring{field: "`cmf_theme`.`name`"},
	Version:     whereHelperstring{field: "`cmf_theme`.`version`"},
	DemoURL:     whereHelperstring{field: "`cmf_theme`.`demo_url`"},
	Thumbnail:   whereHelperstring{field: "`cmf_theme`.`thumbnail`"},
	Author:      whereHelperstring{field: "`cmf_theme`.`author`"},
	AuthorURL:   whereHelperstring{field: "`cmf_theme`.`author_url`"},
	Lang:        whereHelperstring{field: "`cmf_theme`.`lang`"},
	Keywords:    whereHelperstring{field: "`cmf_theme`.`keywords`"},
	Description: whereHelperstring{field: "`cmf_theme`.`description`"},
}

// CMFThemeRels is where relationship names are stored.
var CMFThemeRels = struct {
}{}

// cmfThemeR is where relationships are stored.
type cmfThemeR struct {
}

// NewStruct creates a new relationship struct
func (*cmfThemeR) NewStruct() *cmfThemeR {
	return &cmfThemeR{}
}

// cmfThemeL is where Load methods for each relationship are stored.
type cmfThemeL struct{}

var (
	cmfThemeAllColumns            = []string{"id", "create_time", "update_time", "status", "is_compiled", "theme", "name", "version", "demo_url", "thumbnail", "author", "author_url", "lang", "keywords", "description"}
	cmfThemeColumnsWithoutDefault = []string{"theme", "name", "version", "demo_url", "thumbnail", "author", "author_url", "lang", "keywords", "description"}
	cmfThemeColumnsWithDefault    = []string{"id", "create_time", "update_time", "status", "is_compiled"}
	cmfThemePrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFThemeSlice is an alias for a slice of pointers to CMFTheme.
	// This should generally be used opposed to []CMFTheme.
	CMFThemeSlice []*CMFTheme
	// CMFThemeHook is the signature for custom CMFTheme hook methods
	CMFThemeHook func(context.Context, boil.ContextExecutor, *CMFTheme) error

	cmfThemeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfThemeType                 = reflect.TypeOf(&CMFTheme{})
	cmfThemeMapping              = queries.MakeStructMapping(cmfThemeType)
	cmfThemePrimaryKeyMapping, _ = queries.BindMapping(cmfThemeType, cmfThemeMapping, cmfThemePrimaryKeyColumns)
	cmfThemeInsertCacheMut       sync.RWMutex
	cmfThemeInsertCache          = make(map[string]insertCache)
	cmfThemeUpdateCacheMut       sync.RWMutex
	cmfThemeUpdateCache          = make(map[string]updateCache)
	cmfThemeUpsertCacheMut       sync.RWMutex
	cmfThemeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfThemeBeforeInsertHooks []CMFThemeHook
var cmfThemeBeforeUpdateHooks []CMFThemeHook
var cmfThemeBeforeDeleteHooks []CMFThemeHook
var cmfThemeBeforeUpsertHooks []CMFThemeHook

var cmfThemeAfterInsertHooks []CMFThemeHook
var cmfThemeAfterSelectHooks []CMFThemeHook
var cmfThemeAfterUpdateHooks []CMFThemeHook
var cmfThemeAfterDeleteHooks []CMFThemeHook
var cmfThemeAfterUpsertHooks []CMFThemeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFTheme) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFTheme) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFTheme) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFTheme) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFTheme) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFTheme) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFTheme) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFTheme) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFTheme) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfThemeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFThemeHook registers your hook function for all future operations.
func AddCMFThemeHook(hookPoint boil.HookPoint, cmfThemeHook CMFThemeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfThemeBeforeInsertHooks = append(cmfThemeBeforeInsertHooks, cmfThemeHook)
	case boil.BeforeUpdateHook:
		cmfThemeBeforeUpdateHooks = append(cmfThemeBeforeUpdateHooks, cmfThemeHook)
	case boil.BeforeDeleteHook:
		cmfThemeBeforeDeleteHooks = append(cmfThemeBeforeDeleteHooks, cmfThemeHook)
	case boil.BeforeUpsertHook:
		cmfThemeBeforeUpsertHooks = append(cmfThemeBeforeUpsertHooks, cmfThemeHook)
	case boil.AfterInsertHook:
		cmfThemeAfterInsertHooks = append(cmfThemeAfterInsertHooks, cmfThemeHook)
	case boil.AfterSelectHook:
		cmfThemeAfterSelectHooks = append(cmfThemeAfterSelectHooks, cmfThemeHook)
	case boil.AfterUpdateHook:
		cmfThemeAfterUpdateHooks = append(cmfThemeAfterUpdateHooks, cmfThemeHook)
	case boil.AfterDeleteHook:
		cmfThemeAfterDeleteHooks = append(cmfThemeAfterDeleteHooks, cmfThemeHook)
	case boil.AfterUpsertHook:
		cmfThemeAfterUpsertHooks = append(cmfThemeAfterUpsertHooks, cmfThemeHook)
	}
}

// One returns a single cmfTheme record from the query.
func (q cmfThemeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFTheme, error) {
	o := &CMFTheme{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_theme")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFTheme records from the query.
func (q cmfThemeQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFThemeSlice, error) {
	var o []*CMFTheme

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFTheme slice")
	}

	if len(cmfThemeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFTheme records in the query.
func (q cmfThemeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_theme rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfThemeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_theme exists")
	}

	return count > 0, nil
}

// CMFThemes retrieves all the records using an executor.
func CMFThemes(mods ...qm.QueryMod) cmfThemeQuery {
	mods = append(mods, qm.From("`cmf_theme`"))
	return cmfThemeQuery{NewQuery(mods...)}
}

// FindCMFTheme retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFTheme(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFTheme, error) {
	cmfThemeObj := &CMFTheme{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_theme` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfThemeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_theme")
	}

	return cmfThemeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFTheme) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_theme provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfThemeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfThemeInsertCacheMut.RLock()
	cache, cached := cmfThemeInsertCache[key]
	cmfThemeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfThemeAllColumns,
			cmfThemeColumnsWithDefault,
			cmfThemeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfThemeType, cmfThemeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfThemeType, cmfThemeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_theme` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_theme` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_theme` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfThemePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_theme")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfThemeMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_theme")
	}

CacheNoHooks:
	if !cached {
		cmfThemeInsertCacheMut.Lock()
		cmfThemeInsertCache[key] = cache
		cmfThemeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFTheme.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFTheme) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfThemeUpdateCacheMut.RLock()
	cache, cached := cmfThemeUpdateCache[key]
	cmfThemeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfThemeAllColumns,
			cmfThemePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_theme, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_theme` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfThemePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfThemeType, cmfThemeMapping, append(wl, cmfThemePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_theme row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_theme")
	}

	if !cached {
		cmfThemeUpdateCacheMut.Lock()
		cmfThemeUpdateCache[key] = cache
		cmfThemeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfThemeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_theme")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_theme")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFThemeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfThemePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_theme` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfThemePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfTheme slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfTheme")
	}
	return rowsAff, nil
}

var mySQLCMFThemeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFTheme) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_theme provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfThemeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFThemeUniqueColumns, o)

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

	cmfThemeUpsertCacheMut.RLock()
	cache, cached := cmfThemeUpsertCache[key]
	cmfThemeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfThemeAllColumns,
			cmfThemeColumnsWithDefault,
			cmfThemeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfThemeAllColumns,
			cmfThemePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_theme, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_theme`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_theme` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfThemeType, cmfThemeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfThemeType, cmfThemeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_theme")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfThemeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfThemeType, cmfThemeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_theme")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_theme")
	}

CacheNoHooks:
	if !cached {
		cmfThemeUpsertCacheMut.Lock()
		cmfThemeUpsertCache[key] = cache
		cmfThemeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFTheme record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFTheme) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFTheme provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfThemePrimaryKeyMapping)
	sql := "DELETE FROM `cmf_theme` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_theme")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_theme")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfThemeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfThemeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_theme")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_theme")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFThemeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfThemeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfThemePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_theme` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfThemePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfTheme slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_theme")
	}

	if len(cmfThemeAfterDeleteHooks) != 0 {
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
func (o *CMFTheme) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFTheme(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFThemeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFThemeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfThemePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_theme`.* FROM `cmf_theme` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfThemePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFThemeSlice")
	}

	*o = slice

	return nil
}

// CMFThemeExists checks if the CMFTheme row exists.
func CMFThemeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_theme` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_theme exists")
	}

	return exists, nil
}