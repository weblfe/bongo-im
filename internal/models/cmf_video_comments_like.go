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

// CMFVideoCommentsLike is an object representing the database table.
type CMFVideoCommentsLike struct {
	ID        int `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID       int `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Commentid int `boil:"commentid" json:"commentid" toml:"commentid" yaml:"commentid"`
	Addtime   int `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Touid     int `boil:"touid" json:"touid" toml:"touid" yaml:"touid"`
	Videoid   int `boil:"videoid" json:"videoid" toml:"videoid" yaml:"videoid"`

	R *cmfVideoCommentsLikeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfVideoCommentsLikeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFVideoCommentsLikeColumns = struct {
	ID        string
	UID       string
	Commentid string
	Addtime   string
	Touid     string
	Videoid   string
}{
	ID:        "id",
	UID:       "uid",
	Commentid: "commentid",
	Addtime:   "addtime",
	Touid:     "touid",
	Videoid:   "videoid",
}

// Generated where

var CMFVideoCommentsLikeWhere = struct {
	ID        whereHelperint
	UID       whereHelperint
	Commentid whereHelperint
	Addtime   whereHelperint
	Touid     whereHelperint
	Videoid   whereHelperint
}{
	ID:        whereHelperint{field: "`cmf_video_comments_like`.`id`"},
	UID:       whereHelperint{field: "`cmf_video_comments_like`.`uid`"},
	Commentid: whereHelperint{field: "`cmf_video_comments_like`.`commentid`"},
	Addtime:   whereHelperint{field: "`cmf_video_comments_like`.`addtime`"},
	Touid:     whereHelperint{field: "`cmf_video_comments_like`.`touid`"},
	Videoid:   whereHelperint{field: "`cmf_video_comments_like`.`videoid`"},
}

// CMFVideoCommentsLikeRels is where relationship names are stored.
var CMFVideoCommentsLikeRels = struct {
}{}

// cmfVideoCommentsLikeR is where relationships are stored.
type cmfVideoCommentsLikeR struct {
}

// NewStruct creates a new relationship struct
func (*cmfVideoCommentsLikeR) NewStruct() *cmfVideoCommentsLikeR {
	return &cmfVideoCommentsLikeR{}
}

// cmfVideoCommentsLikeL is where Load methods for each relationship are stored.
type cmfVideoCommentsLikeL struct{}

var (
	cmfVideoCommentsLikeAllColumns            = []string{"id", "uid", "commentid", "addtime", "touid", "videoid"}
	cmfVideoCommentsLikeColumnsWithoutDefault = []string{}
	cmfVideoCommentsLikeColumnsWithDefault    = []string{"id", "uid", "commentid", "addtime", "touid", "videoid"}
	cmfVideoCommentsLikePrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFVideoCommentsLikeSlice is an alias for a slice of pointers to CMFVideoCommentsLike.
	// This should generally be used opposed to []CMFVideoCommentsLike.
	CMFVideoCommentsLikeSlice []*CMFVideoCommentsLike
	// CMFVideoCommentsLikeHook is the signature for custom CMFVideoCommentsLike hook methods
	CMFVideoCommentsLikeHook func(context.Context, boil.ContextExecutor, *CMFVideoCommentsLike) error

	cmfVideoCommentsLikeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfVideoCommentsLikeType                 = reflect.TypeOf(&CMFVideoCommentsLike{})
	cmfVideoCommentsLikeMapping              = queries.MakeStructMapping(cmfVideoCommentsLikeType)
	cmfVideoCommentsLikePrimaryKeyMapping, _ = queries.BindMapping(cmfVideoCommentsLikeType, cmfVideoCommentsLikeMapping, cmfVideoCommentsLikePrimaryKeyColumns)
	cmfVideoCommentsLikeInsertCacheMut       sync.RWMutex
	cmfVideoCommentsLikeInsertCache          = make(map[string]insertCache)
	cmfVideoCommentsLikeUpdateCacheMut       sync.RWMutex
	cmfVideoCommentsLikeUpdateCache          = make(map[string]updateCache)
	cmfVideoCommentsLikeUpsertCacheMut       sync.RWMutex
	cmfVideoCommentsLikeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfVideoCommentsLikeBeforeInsertHooks []CMFVideoCommentsLikeHook
var cmfVideoCommentsLikeBeforeUpdateHooks []CMFVideoCommentsLikeHook
var cmfVideoCommentsLikeBeforeDeleteHooks []CMFVideoCommentsLikeHook
var cmfVideoCommentsLikeBeforeUpsertHooks []CMFVideoCommentsLikeHook

var cmfVideoCommentsLikeAfterInsertHooks []CMFVideoCommentsLikeHook
var cmfVideoCommentsLikeAfterSelectHooks []CMFVideoCommentsLikeHook
var cmfVideoCommentsLikeAfterUpdateHooks []CMFVideoCommentsLikeHook
var cmfVideoCommentsLikeAfterDeleteHooks []CMFVideoCommentsLikeHook
var cmfVideoCommentsLikeAfterUpsertHooks []CMFVideoCommentsLikeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFVideoCommentsLike) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFVideoCommentsLike) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFVideoCommentsLike) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFVideoCommentsLike) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFVideoCommentsLike) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFVideoCommentsLike) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFVideoCommentsLike) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFVideoCommentsLike) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFVideoCommentsLike) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfVideoCommentsLikeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFVideoCommentsLikeHook registers your hook function for all future operations.
func AddCMFVideoCommentsLikeHook(hookPoint boil.HookPoint, cmfVideoCommentsLikeHook CMFVideoCommentsLikeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfVideoCommentsLikeBeforeInsertHooks = append(cmfVideoCommentsLikeBeforeInsertHooks, cmfVideoCommentsLikeHook)
	case boil.BeforeUpdateHook:
		cmfVideoCommentsLikeBeforeUpdateHooks = append(cmfVideoCommentsLikeBeforeUpdateHooks, cmfVideoCommentsLikeHook)
	case boil.BeforeDeleteHook:
		cmfVideoCommentsLikeBeforeDeleteHooks = append(cmfVideoCommentsLikeBeforeDeleteHooks, cmfVideoCommentsLikeHook)
	case boil.BeforeUpsertHook:
		cmfVideoCommentsLikeBeforeUpsertHooks = append(cmfVideoCommentsLikeBeforeUpsertHooks, cmfVideoCommentsLikeHook)
	case boil.AfterInsertHook:
		cmfVideoCommentsLikeAfterInsertHooks = append(cmfVideoCommentsLikeAfterInsertHooks, cmfVideoCommentsLikeHook)
	case boil.AfterSelectHook:
		cmfVideoCommentsLikeAfterSelectHooks = append(cmfVideoCommentsLikeAfterSelectHooks, cmfVideoCommentsLikeHook)
	case boil.AfterUpdateHook:
		cmfVideoCommentsLikeAfterUpdateHooks = append(cmfVideoCommentsLikeAfterUpdateHooks, cmfVideoCommentsLikeHook)
	case boil.AfterDeleteHook:
		cmfVideoCommentsLikeAfterDeleteHooks = append(cmfVideoCommentsLikeAfterDeleteHooks, cmfVideoCommentsLikeHook)
	case boil.AfterUpsertHook:
		cmfVideoCommentsLikeAfterUpsertHooks = append(cmfVideoCommentsLikeAfterUpsertHooks, cmfVideoCommentsLikeHook)
	}
}

// One returns a single cmfVideoCommentsLike record from the query.
func (q cmfVideoCommentsLikeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFVideoCommentsLike, error) {
	o := &CMFVideoCommentsLike{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_video_comments_like")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFVideoCommentsLike records from the query.
func (q cmfVideoCommentsLikeQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFVideoCommentsLikeSlice, error) {
	var o []*CMFVideoCommentsLike

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFVideoCommentsLike slice")
	}

	if len(cmfVideoCommentsLikeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFVideoCommentsLike records in the query.
func (q cmfVideoCommentsLikeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_video_comments_like rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfVideoCommentsLikeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_video_comments_like exists")
	}

	return count > 0, nil
}

// CMFVideoCommentsLikes retrieves all the records using an executor.
func CMFVideoCommentsLikes(mods ...qm.QueryMod) cmfVideoCommentsLikeQuery {
	mods = append(mods, qm.From("`cmf_video_comments_like`"))
	return cmfVideoCommentsLikeQuery{NewQuery(mods...)}
}

// FindCMFVideoCommentsLike retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFVideoCommentsLike(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFVideoCommentsLike, error) {
	cmfVideoCommentsLikeObj := &CMFVideoCommentsLike{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_video_comments_like` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfVideoCommentsLikeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_video_comments_like")
	}

	return cmfVideoCommentsLikeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFVideoCommentsLike) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_video_comments_like provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfVideoCommentsLikeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfVideoCommentsLikeInsertCacheMut.RLock()
	cache, cached := cmfVideoCommentsLikeInsertCache[key]
	cmfVideoCommentsLikeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfVideoCommentsLikeAllColumns,
			cmfVideoCommentsLikeColumnsWithDefault,
			cmfVideoCommentsLikeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfVideoCommentsLikeType, cmfVideoCommentsLikeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfVideoCommentsLikeType, cmfVideoCommentsLikeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_video_comments_like` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_video_comments_like` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_video_comments_like` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfVideoCommentsLikePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_video_comments_like")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfVideoCommentsLikeMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_video_comments_like")
	}

CacheNoHooks:
	if !cached {
		cmfVideoCommentsLikeInsertCacheMut.Lock()
		cmfVideoCommentsLikeInsertCache[key] = cache
		cmfVideoCommentsLikeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFVideoCommentsLike.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFVideoCommentsLike) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfVideoCommentsLikeUpdateCacheMut.RLock()
	cache, cached := cmfVideoCommentsLikeUpdateCache[key]
	cmfVideoCommentsLikeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfVideoCommentsLikeAllColumns,
			cmfVideoCommentsLikePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_video_comments_like, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_video_comments_like` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfVideoCommentsLikePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfVideoCommentsLikeType, cmfVideoCommentsLikeMapping, append(wl, cmfVideoCommentsLikePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_video_comments_like row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_video_comments_like")
	}

	if !cached {
		cmfVideoCommentsLikeUpdateCacheMut.Lock()
		cmfVideoCommentsLikeUpdateCache[key] = cache
		cmfVideoCommentsLikeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfVideoCommentsLikeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_video_comments_like")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_video_comments_like")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFVideoCommentsLikeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoCommentsLikePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_video_comments_like` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoCommentsLikePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfVideoCommentsLike slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfVideoCommentsLike")
	}
	return rowsAff, nil
}

var mySQLCMFVideoCommentsLikeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFVideoCommentsLike) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_video_comments_like provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfVideoCommentsLikeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFVideoCommentsLikeUniqueColumns, o)

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

	cmfVideoCommentsLikeUpsertCacheMut.RLock()
	cache, cached := cmfVideoCommentsLikeUpsertCache[key]
	cmfVideoCommentsLikeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfVideoCommentsLikeAllColumns,
			cmfVideoCommentsLikeColumnsWithDefault,
			cmfVideoCommentsLikeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfVideoCommentsLikeAllColumns,
			cmfVideoCommentsLikePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_video_comments_like, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_video_comments_like`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_video_comments_like` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfVideoCommentsLikeType, cmfVideoCommentsLikeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfVideoCommentsLikeType, cmfVideoCommentsLikeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_video_comments_like")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfVideoCommentsLikeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfVideoCommentsLikeType, cmfVideoCommentsLikeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_video_comments_like")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_video_comments_like")
	}

CacheNoHooks:
	if !cached {
		cmfVideoCommentsLikeUpsertCacheMut.Lock()
		cmfVideoCommentsLikeUpsertCache[key] = cache
		cmfVideoCommentsLikeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFVideoCommentsLike record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFVideoCommentsLike) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFVideoCommentsLike provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfVideoCommentsLikePrimaryKeyMapping)
	sql := "DELETE FROM `cmf_video_comments_like` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_video_comments_like")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_video_comments_like")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfVideoCommentsLikeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfVideoCommentsLikeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_video_comments_like")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_video_comments_like")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFVideoCommentsLikeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfVideoCommentsLikeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoCommentsLikePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_video_comments_like` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoCommentsLikePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfVideoCommentsLike slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_video_comments_like")
	}

	if len(cmfVideoCommentsLikeAfterDeleteHooks) != 0 {
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
func (o *CMFVideoCommentsLike) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFVideoCommentsLike(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFVideoCommentsLikeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFVideoCommentsLikeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfVideoCommentsLikePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_video_comments_like`.* FROM `cmf_video_comments_like` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfVideoCommentsLikePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFVideoCommentsLikeSlice")
	}

	*o = slice

	return nil
}

// CMFVideoCommentsLikeExists checks if the CMFVideoCommentsLike row exists.
func CMFVideoCommentsLikeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_video_comments_like` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_video_comments_like exists")
	}

	return exists, nil
}
