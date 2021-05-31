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

// CMFPaidprogram is an object representing the database table.
type CMFPaidprogram struct {
	ID            int64         `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID           int64         `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Classid       int           `boil:"classid" json:"classid" toml:"classid" yaml:"classid"`
	Title         string        `boil:"title" json:"title" toml:"title" yaml:"title"`
	Thumb         string        `boil:"thumb" json:"thumb" toml:"thumb" yaml:"thumb"`
	Content       string        `boil:"content" json:"content" toml:"content" yaml:"content"`
	PersonalDesc  string        `boil:"personal_desc" json:"personal_desc" toml:"personal_desc" yaml:"personal_desc"`
	Money         types.Decimal `boil:"money" json:"money" toml:"money" yaml:"money"`
	Type          bool          `boil:"type" json:"type" toml:"type" yaml:"type"`
	Videos        string        `boil:"videos" json:"videos" toml:"videos" yaml:"videos"`
	SaleNums      int64         `boil:"sale_nums" json:"sale_nums" toml:"sale_nums" yaml:"sale_nums"`
	Status        bool          `boil:"status" json:"status" toml:"status" yaml:"status"`
	EvaluateNums  int64         `boil:"evaluate_nums" json:"evaluate_nums" toml:"evaluate_nums" yaml:"evaluate_nums"`
	EvaluateTotal int64         `boil:"evaluate_total" json:"evaluate_total" toml:"evaluate_total" yaml:"evaluate_total"`
	Addtime       int           `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`
	Edittime      int           `boil:"edittime" json:"edittime" toml:"edittime" yaml:"edittime"`

	R *cmfPaidprogramR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfPaidprogramL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFPaidprogramColumns = struct {
	ID            string
	UID           string
	Classid       string
	Title         string
	Thumb         string
	Content       string
	PersonalDesc  string
	Money         string
	Type          string
	Videos        string
	SaleNums      string
	Status        string
	EvaluateNums  string
	EvaluateTotal string
	Addtime       string
	Edittime      string
}{
	ID:            "id",
	UID:           "uid",
	Classid:       "classid",
	Title:         "title",
	Thumb:         "thumb",
	Content:       "content",
	PersonalDesc:  "personal_desc",
	Money:         "money",
	Type:          "type",
	Videos:        "videos",
	SaleNums:      "sale_nums",
	Status:        "status",
	EvaluateNums:  "evaluate_nums",
	EvaluateTotal: "evaluate_total",
	Addtime:       "addtime",
	Edittime:      "edittime",
}

// Generated where

var CMFPaidprogramWhere = struct {
	ID            whereHelperint64
	UID           whereHelperint64
	Classid       whereHelperint
	Title         whereHelperstring
	Thumb         whereHelperstring
	Content       whereHelperstring
	PersonalDesc  whereHelperstring
	Money         whereHelpertypes_Decimal
	Type          whereHelperbool
	Videos        whereHelperstring
	SaleNums      whereHelperint64
	Status        whereHelperbool
	EvaluateNums  whereHelperint64
	EvaluateTotal whereHelperint64
	Addtime       whereHelperint
	Edittime      whereHelperint
}{
	ID:            whereHelperint64{field: "`cmf_paidprogram`.`id`"},
	UID:           whereHelperint64{field: "`cmf_paidprogram`.`uid`"},
	Classid:       whereHelperint{field: "`cmf_paidprogram`.`classid`"},
	Title:         whereHelperstring{field: "`cmf_paidprogram`.`title`"},
	Thumb:         whereHelperstring{field: "`cmf_paidprogram`.`thumb`"},
	Content:       whereHelperstring{field: "`cmf_paidprogram`.`content`"},
	PersonalDesc:  whereHelperstring{field: "`cmf_paidprogram`.`personal_desc`"},
	Money:         whereHelpertypes_Decimal{field: "`cmf_paidprogram`.`money`"},
	Type:          whereHelperbool{field: "`cmf_paidprogram`.`type`"},
	Videos:        whereHelperstring{field: "`cmf_paidprogram`.`videos`"},
	SaleNums:      whereHelperint64{field: "`cmf_paidprogram`.`sale_nums`"},
	Status:        whereHelperbool{field: "`cmf_paidprogram`.`status`"},
	EvaluateNums:  whereHelperint64{field: "`cmf_paidprogram`.`evaluate_nums`"},
	EvaluateTotal: whereHelperint64{field: "`cmf_paidprogram`.`evaluate_total`"},
	Addtime:       whereHelperint{field: "`cmf_paidprogram`.`addtime`"},
	Edittime:      whereHelperint{field: "`cmf_paidprogram`.`edittime`"},
}

// CMFPaidprogramRels is where relationship names are stored.
var CMFPaidprogramRels = struct {
}{}

// cmfPaidprogramR is where relationships are stored.
type cmfPaidprogramR struct {
}

// NewStruct creates a new relationship struct
func (*cmfPaidprogramR) NewStruct() *cmfPaidprogramR {
	return &cmfPaidprogramR{}
}

// cmfPaidprogramL is where Load methods for each relationship are stored.
type cmfPaidprogramL struct{}

var (
	cmfPaidprogramAllColumns            = []string{"id", "uid", "classid", "title", "thumb", "content", "personal_desc", "money", "type", "videos", "sale_nums", "status", "evaluate_nums", "evaluate_total", "addtime", "edittime"}
	cmfPaidprogramColumnsWithoutDefault = []string{"title", "thumb", "content", "personal_desc", "videos"}
	cmfPaidprogramColumnsWithDefault    = []string{"id", "uid", "classid", "money", "type", "sale_nums", "status", "evaluate_nums", "evaluate_total", "addtime", "edittime"}
	cmfPaidprogramPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFPaidprogramSlice is an alias for a slice of pointers to CMFPaidprogram.
	// This should generally be used opposed to []CMFPaidprogram.
	CMFPaidprogramSlice []*CMFPaidprogram
	// CMFPaidprogramHook is the signature for custom CMFPaidprogram hook methods
	CMFPaidprogramHook func(context.Context, boil.ContextExecutor, *CMFPaidprogram) error

	cmfPaidprogramQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfPaidprogramType                 = reflect.TypeOf(&CMFPaidprogram{})
	cmfPaidprogramMapping              = queries.MakeStructMapping(cmfPaidprogramType)
	cmfPaidprogramPrimaryKeyMapping, _ = queries.BindMapping(cmfPaidprogramType, cmfPaidprogramMapping, cmfPaidprogramPrimaryKeyColumns)
	cmfPaidprogramInsertCacheMut       sync.RWMutex
	cmfPaidprogramInsertCache          = make(map[string]insertCache)
	cmfPaidprogramUpdateCacheMut       sync.RWMutex
	cmfPaidprogramUpdateCache          = make(map[string]updateCache)
	cmfPaidprogramUpsertCacheMut       sync.RWMutex
	cmfPaidprogramUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfPaidprogramBeforeInsertHooks []CMFPaidprogramHook
var cmfPaidprogramBeforeUpdateHooks []CMFPaidprogramHook
var cmfPaidprogramBeforeDeleteHooks []CMFPaidprogramHook
var cmfPaidprogramBeforeUpsertHooks []CMFPaidprogramHook

var cmfPaidprogramAfterInsertHooks []CMFPaidprogramHook
var cmfPaidprogramAfterSelectHooks []CMFPaidprogramHook
var cmfPaidprogramAfterUpdateHooks []CMFPaidprogramHook
var cmfPaidprogramAfterDeleteHooks []CMFPaidprogramHook
var cmfPaidprogramAfterUpsertHooks []CMFPaidprogramHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFPaidprogram) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFPaidprogram) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFPaidprogram) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFPaidprogram) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFPaidprogram) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFPaidprogram) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFPaidprogram) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFPaidprogram) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFPaidprogram) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfPaidprogramAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFPaidprogramHook registers your hook function for all future operations.
func AddCMFPaidprogramHook(hookPoint boil.HookPoint, cmfPaidprogramHook CMFPaidprogramHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfPaidprogramBeforeInsertHooks = append(cmfPaidprogramBeforeInsertHooks, cmfPaidprogramHook)
	case boil.BeforeUpdateHook:
		cmfPaidprogramBeforeUpdateHooks = append(cmfPaidprogramBeforeUpdateHooks, cmfPaidprogramHook)
	case boil.BeforeDeleteHook:
		cmfPaidprogramBeforeDeleteHooks = append(cmfPaidprogramBeforeDeleteHooks, cmfPaidprogramHook)
	case boil.BeforeUpsertHook:
		cmfPaidprogramBeforeUpsertHooks = append(cmfPaidprogramBeforeUpsertHooks, cmfPaidprogramHook)
	case boil.AfterInsertHook:
		cmfPaidprogramAfterInsertHooks = append(cmfPaidprogramAfterInsertHooks, cmfPaidprogramHook)
	case boil.AfterSelectHook:
		cmfPaidprogramAfterSelectHooks = append(cmfPaidprogramAfterSelectHooks, cmfPaidprogramHook)
	case boil.AfterUpdateHook:
		cmfPaidprogramAfterUpdateHooks = append(cmfPaidprogramAfterUpdateHooks, cmfPaidprogramHook)
	case boil.AfterDeleteHook:
		cmfPaidprogramAfterDeleteHooks = append(cmfPaidprogramAfterDeleteHooks, cmfPaidprogramHook)
	case boil.AfterUpsertHook:
		cmfPaidprogramAfterUpsertHooks = append(cmfPaidprogramAfterUpsertHooks, cmfPaidprogramHook)
	}
}

// One returns a single cmfPaidprogram record from the query.
func (q cmfPaidprogramQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFPaidprogram, error) {
	o := &CMFPaidprogram{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_paidprogram")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFPaidprogram records from the query.
func (q cmfPaidprogramQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFPaidprogramSlice, error) {
	var o []*CMFPaidprogram

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFPaidprogram slice")
	}

	if len(cmfPaidprogramAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFPaidprogram records in the query.
func (q cmfPaidprogramQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_paidprogram rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfPaidprogramQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_paidprogram exists")
	}

	return count > 0, nil
}

// CMFPaidprograms retrieves all the records using an executor.
func CMFPaidprograms(mods ...qm.QueryMod) cmfPaidprogramQuery {
	mods = append(mods, qm.From("`cmf_paidprogram`"))
	return cmfPaidprogramQuery{NewQuery(mods...)}
}

// FindCMFPaidprogram retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFPaidprogram(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CMFPaidprogram, error) {
	cmfPaidprogramObj := &CMFPaidprogram{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_paidprogram` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfPaidprogramObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_paidprogram")
	}

	return cmfPaidprogramObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFPaidprogram) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_paidprogram provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPaidprogramColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfPaidprogramInsertCacheMut.RLock()
	cache, cached := cmfPaidprogramInsertCache[key]
	cmfPaidprogramInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfPaidprogramAllColumns,
			cmfPaidprogramColumnsWithDefault,
			cmfPaidprogramColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfPaidprogramType, cmfPaidprogramMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfPaidprogramType, cmfPaidprogramMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_paidprogram` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_paidprogram` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_paidprogram` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfPaidprogramPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_paidprogram")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPaidprogramMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_paidprogram")
	}

CacheNoHooks:
	if !cached {
		cmfPaidprogramInsertCacheMut.Lock()
		cmfPaidprogramInsertCache[key] = cache
		cmfPaidprogramInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFPaidprogram.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFPaidprogram) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfPaidprogramUpdateCacheMut.RLock()
	cache, cached := cmfPaidprogramUpdateCache[key]
	cmfPaidprogramUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfPaidprogramAllColumns,
			cmfPaidprogramPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_paidprogram, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_paidprogram` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfPaidprogramPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfPaidprogramType, cmfPaidprogramMapping, append(wl, cmfPaidprogramPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_paidprogram row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_paidprogram")
	}

	if !cached {
		cmfPaidprogramUpdateCacheMut.Lock()
		cmfPaidprogramUpdateCache[key] = cache
		cmfPaidprogramUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfPaidprogramQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_paidprogram")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_paidprogram")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFPaidprogramSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPaidprogramPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_paidprogram` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPaidprogramPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfPaidprogram slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfPaidprogram")
	}
	return rowsAff, nil
}

var mySQLCMFPaidprogramUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFPaidprogram) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_paidprogram provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfPaidprogramColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFPaidprogramUniqueColumns, o)

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

	cmfPaidprogramUpsertCacheMut.RLock()
	cache, cached := cmfPaidprogramUpsertCache[key]
	cmfPaidprogramUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfPaidprogramAllColumns,
			cmfPaidprogramColumnsWithDefault,
			cmfPaidprogramColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfPaidprogramAllColumns,
			cmfPaidprogramPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_paidprogram, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_paidprogram`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_paidprogram` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfPaidprogramType, cmfPaidprogramMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfPaidprogramType, cmfPaidprogramMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_paidprogram")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfPaidprogramMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfPaidprogramType, cmfPaidprogramMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_paidprogram")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_paidprogram")
	}

CacheNoHooks:
	if !cached {
		cmfPaidprogramUpsertCacheMut.Lock()
		cmfPaidprogramUpsertCache[key] = cache
		cmfPaidprogramUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFPaidprogram record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFPaidprogram) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFPaidprogram provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfPaidprogramPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_paidprogram` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_paidprogram")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_paidprogram")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfPaidprogramQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfPaidprogramQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_paidprogram")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_paidprogram")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFPaidprogramSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfPaidprogramBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPaidprogramPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_paidprogram` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPaidprogramPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfPaidprogram slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_paidprogram")
	}

	if len(cmfPaidprogramAfterDeleteHooks) != 0 {
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
func (o *CMFPaidprogram) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFPaidprogram(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFPaidprogramSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFPaidprogramSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfPaidprogramPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_paidprogram`.* FROM `cmf_paidprogram` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfPaidprogramPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFPaidprogramSlice")
	}

	*o = slice

	return nil
}

// CMFPaidprogramExists checks if the CMFPaidprogram row exists.
func CMFPaidprogramExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_paidprogram` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_paidprogram exists")
	}

	return exists, nil
}
