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

// CMFAnchorPolicyLog is an object representing the database table.
type CMFAnchorPolicyLog struct {
	ID             int           `boil:"id" json:"id" toml:"id" yaml:"id"`
	FamilyID       int           `boil:"family_id" json:"family_id" toml:"family_id" yaml:"family_id"`
	UID            int64         `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Policy         int16         `boil:"policy" json:"policy" toml:"policy" yaml:"policy"`
	Proportion     types.Decimal `boil:"proportion" json:"proportion" toml:"proportion" yaml:"proportion"`
	PrevPolicy     int16         `boil:"prev_policy" json:"prev_policy" toml:"prev_policy" yaml:"prev_policy"`
	PrevProportion types.Decimal `boil:"prev_proportion" json:"prev_proportion" toml:"prev_proportion" yaml:"prev_proportion"`
	CreatorUID     int           `boil:"creator_uid" json:"creator_uid" toml:"creator_uid" yaml:"creator_uid"`
	EffectedAt     time.Time     `boil:"effected_at" json:"effected_at" toml:"effected_at" yaml:"effected_at"`
	CreatedAt      time.Time     `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *cmfAnchorPolicyLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfAnchorPolicyLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFAnchorPolicyLogColumns = struct {
	ID             string
	FamilyID       string
	UID            string
	Policy         string
	Proportion     string
	PrevPolicy     string
	PrevProportion string
	CreatorUID     string
	EffectedAt     string
	CreatedAt      string
}{
	ID:             "id",
	FamilyID:       "family_id",
	UID:            "uid",
	Policy:         "policy",
	Proportion:     "proportion",
	PrevPolicy:     "prev_policy",
	PrevProportion: "prev_proportion",
	CreatorUID:     "creator_uid",
	EffectedAt:     "effected_at",
	CreatedAt:      "created_at",
}

// Generated where

type whereHelperint16 struct{ field string }

func (w whereHelperint16) EQ(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint16) NEQ(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint16) LT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint16) LTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint16) GT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint16) GTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint16) IN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint16) NIN(slice []int16) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_Decimal struct{ field string }

func (w whereHelpertypes_Decimal) EQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_Decimal) NEQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_Decimal) LT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Decimal) LTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Decimal) GT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Decimal) GTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CMFAnchorPolicyLogWhere = struct {
	ID             whereHelperint
	FamilyID       whereHelperint
	UID            whereHelperint64
	Policy         whereHelperint16
	Proportion     whereHelpertypes_Decimal
	PrevPolicy     whereHelperint16
	PrevProportion whereHelpertypes_Decimal
	CreatorUID     whereHelperint
	EffectedAt     whereHelpertime_Time
	CreatedAt      whereHelpertime_Time
}{
	ID:             whereHelperint{field: "`cmf_anchor_policy_log`.`id`"},
	FamilyID:       whereHelperint{field: "`cmf_anchor_policy_log`.`family_id`"},
	UID:            whereHelperint64{field: "`cmf_anchor_policy_log`.`uid`"},
	Policy:         whereHelperint16{field: "`cmf_anchor_policy_log`.`policy`"},
	Proportion:     whereHelpertypes_Decimal{field: "`cmf_anchor_policy_log`.`proportion`"},
	PrevPolicy:     whereHelperint16{field: "`cmf_anchor_policy_log`.`prev_policy`"},
	PrevProportion: whereHelpertypes_Decimal{field: "`cmf_anchor_policy_log`.`prev_proportion`"},
	CreatorUID:     whereHelperint{field: "`cmf_anchor_policy_log`.`creator_uid`"},
	EffectedAt:     whereHelpertime_Time{field: "`cmf_anchor_policy_log`.`effected_at`"},
	CreatedAt:      whereHelpertime_Time{field: "`cmf_anchor_policy_log`.`created_at`"},
}

// CMFAnchorPolicyLogRels is where relationship names are stored.
var CMFAnchorPolicyLogRels = struct {
}{}

// cmfAnchorPolicyLogR is where relationships are stored.
type cmfAnchorPolicyLogR struct {
}

// NewStruct creates a new relationship struct
func (*cmfAnchorPolicyLogR) NewStruct() *cmfAnchorPolicyLogR {
	return &cmfAnchorPolicyLogR{}
}

// cmfAnchorPolicyLogL is where Load methods for each relationship are stored.
type cmfAnchorPolicyLogL struct{}

var (
	cmfAnchorPolicyLogAllColumns            = []string{"id", "family_id", "uid", "policy", "proportion", "prev_policy", "prev_proportion", "creator_uid", "effected_at", "created_at"}
	cmfAnchorPolicyLogColumnsWithoutDefault = []string{"family_id", "uid", "policy", "prev_policy", "creator_uid", "effected_at", "created_at"}
	cmfAnchorPolicyLogColumnsWithDefault    = []string{"id", "proportion", "prev_proportion"}
	cmfAnchorPolicyLogPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFAnchorPolicyLogSlice is an alias for a slice of pointers to CMFAnchorPolicyLog.
	// This should generally be used opposed to []CMFAnchorPolicyLog.
	CMFAnchorPolicyLogSlice []*CMFAnchorPolicyLog
	// CMFAnchorPolicyLogHook is the signature for custom CMFAnchorPolicyLog hook methods
	CMFAnchorPolicyLogHook func(context.Context, boil.ContextExecutor, *CMFAnchorPolicyLog) error

	cmfAnchorPolicyLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfAnchorPolicyLogType                 = reflect.TypeOf(&CMFAnchorPolicyLog{})
	cmfAnchorPolicyLogMapping              = queries.MakeStructMapping(cmfAnchorPolicyLogType)
	cmfAnchorPolicyLogPrimaryKeyMapping, _ = queries.BindMapping(cmfAnchorPolicyLogType, cmfAnchorPolicyLogMapping, cmfAnchorPolicyLogPrimaryKeyColumns)
	cmfAnchorPolicyLogInsertCacheMut       sync.RWMutex
	cmfAnchorPolicyLogInsertCache          = make(map[string]insertCache)
	cmfAnchorPolicyLogUpdateCacheMut       sync.RWMutex
	cmfAnchorPolicyLogUpdateCache          = make(map[string]updateCache)
	cmfAnchorPolicyLogUpsertCacheMut       sync.RWMutex
	cmfAnchorPolicyLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfAnchorPolicyLogBeforeInsertHooks []CMFAnchorPolicyLogHook
var cmfAnchorPolicyLogBeforeUpdateHooks []CMFAnchorPolicyLogHook
var cmfAnchorPolicyLogBeforeDeleteHooks []CMFAnchorPolicyLogHook
var cmfAnchorPolicyLogBeforeUpsertHooks []CMFAnchorPolicyLogHook

var cmfAnchorPolicyLogAfterInsertHooks []CMFAnchorPolicyLogHook
var cmfAnchorPolicyLogAfterSelectHooks []CMFAnchorPolicyLogHook
var cmfAnchorPolicyLogAfterUpdateHooks []CMFAnchorPolicyLogHook
var cmfAnchorPolicyLogAfterDeleteHooks []CMFAnchorPolicyLogHook
var cmfAnchorPolicyLogAfterUpsertHooks []CMFAnchorPolicyLogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFAnchorPolicyLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFAnchorPolicyLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFAnchorPolicyLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFAnchorPolicyLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFAnchorPolicyLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFAnchorPolicyLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFAnchorPolicyLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFAnchorPolicyLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFAnchorPolicyLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfAnchorPolicyLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFAnchorPolicyLogHook registers your hook function for all future operations.
func AddCMFAnchorPolicyLogHook(hookPoint boil.HookPoint, cmfAnchorPolicyLogHook CMFAnchorPolicyLogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfAnchorPolicyLogBeforeInsertHooks = append(cmfAnchorPolicyLogBeforeInsertHooks, cmfAnchorPolicyLogHook)
	case boil.BeforeUpdateHook:
		cmfAnchorPolicyLogBeforeUpdateHooks = append(cmfAnchorPolicyLogBeforeUpdateHooks, cmfAnchorPolicyLogHook)
	case boil.BeforeDeleteHook:
		cmfAnchorPolicyLogBeforeDeleteHooks = append(cmfAnchorPolicyLogBeforeDeleteHooks, cmfAnchorPolicyLogHook)
	case boil.BeforeUpsertHook:
		cmfAnchorPolicyLogBeforeUpsertHooks = append(cmfAnchorPolicyLogBeforeUpsertHooks, cmfAnchorPolicyLogHook)
	case boil.AfterInsertHook:
		cmfAnchorPolicyLogAfterInsertHooks = append(cmfAnchorPolicyLogAfterInsertHooks, cmfAnchorPolicyLogHook)
	case boil.AfterSelectHook:
		cmfAnchorPolicyLogAfterSelectHooks = append(cmfAnchorPolicyLogAfterSelectHooks, cmfAnchorPolicyLogHook)
	case boil.AfterUpdateHook:
		cmfAnchorPolicyLogAfterUpdateHooks = append(cmfAnchorPolicyLogAfterUpdateHooks, cmfAnchorPolicyLogHook)
	case boil.AfterDeleteHook:
		cmfAnchorPolicyLogAfterDeleteHooks = append(cmfAnchorPolicyLogAfterDeleteHooks, cmfAnchorPolicyLogHook)
	case boil.AfterUpsertHook:
		cmfAnchorPolicyLogAfterUpsertHooks = append(cmfAnchorPolicyLogAfterUpsertHooks, cmfAnchorPolicyLogHook)
	}
}

// One returns a single cmfAnchorPolicyLog record from the query.
func (q cmfAnchorPolicyLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFAnchorPolicyLog, error) {
	o := &CMFAnchorPolicyLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_anchor_policy_log")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFAnchorPolicyLog records from the query.
func (q cmfAnchorPolicyLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFAnchorPolicyLogSlice, error) {
	var o []*CMFAnchorPolicyLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFAnchorPolicyLog slice")
	}

	if len(cmfAnchorPolicyLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFAnchorPolicyLog records in the query.
func (q cmfAnchorPolicyLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_anchor_policy_log rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfAnchorPolicyLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_anchor_policy_log exists")
	}

	return count > 0, nil
}

// CMFAnchorPolicyLogs retrieves all the records using an executor.
func CMFAnchorPolicyLogs(mods ...qm.QueryMod) cmfAnchorPolicyLogQuery {
	mods = append(mods, qm.From("`cmf_anchor_policy_log`"))
	return cmfAnchorPolicyLogQuery{NewQuery(mods...)}
}

// FindCMFAnchorPolicyLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFAnchorPolicyLog(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CMFAnchorPolicyLog, error) {
	cmfAnchorPolicyLogObj := &CMFAnchorPolicyLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_anchor_policy_log` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfAnchorPolicyLogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_anchor_policy_log")
	}

	return cmfAnchorPolicyLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFAnchorPolicyLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_anchor_policy_log provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfAnchorPolicyLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfAnchorPolicyLogInsertCacheMut.RLock()
	cache, cached := cmfAnchorPolicyLogInsertCache[key]
	cmfAnchorPolicyLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfAnchorPolicyLogAllColumns,
			cmfAnchorPolicyLogColumnsWithDefault,
			cmfAnchorPolicyLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfAnchorPolicyLogType, cmfAnchorPolicyLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfAnchorPolicyLogType, cmfAnchorPolicyLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_anchor_policy_log` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_anchor_policy_log` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_anchor_policy_log` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfAnchorPolicyLogPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_anchor_policy_log")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfAnchorPolicyLogMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_anchor_policy_log")
	}

CacheNoHooks:
	if !cached {
		cmfAnchorPolicyLogInsertCacheMut.Lock()
		cmfAnchorPolicyLogInsertCache[key] = cache
		cmfAnchorPolicyLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFAnchorPolicyLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFAnchorPolicyLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfAnchorPolicyLogUpdateCacheMut.RLock()
	cache, cached := cmfAnchorPolicyLogUpdateCache[key]
	cmfAnchorPolicyLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfAnchorPolicyLogAllColumns,
			cmfAnchorPolicyLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_anchor_policy_log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_anchor_policy_log` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfAnchorPolicyLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfAnchorPolicyLogType, cmfAnchorPolicyLogMapping, append(wl, cmfAnchorPolicyLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_anchor_policy_log row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_anchor_policy_log")
	}

	if !cached {
		cmfAnchorPolicyLogUpdateCacheMut.Lock()
		cmfAnchorPolicyLogUpdateCache[key] = cache
		cmfAnchorPolicyLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfAnchorPolicyLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_anchor_policy_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_anchor_policy_log")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFAnchorPolicyLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAnchorPolicyLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_anchor_policy_log` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAnchorPolicyLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfAnchorPolicyLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfAnchorPolicyLog")
	}
	return rowsAff, nil
}

var mySQLCMFAnchorPolicyLogUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFAnchorPolicyLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_anchor_policy_log provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfAnchorPolicyLogColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFAnchorPolicyLogUniqueColumns, o)

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

	cmfAnchorPolicyLogUpsertCacheMut.RLock()
	cache, cached := cmfAnchorPolicyLogUpsertCache[key]
	cmfAnchorPolicyLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfAnchorPolicyLogAllColumns,
			cmfAnchorPolicyLogColumnsWithDefault,
			cmfAnchorPolicyLogColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfAnchorPolicyLogAllColumns,
			cmfAnchorPolicyLogPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_anchor_policy_log, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_anchor_policy_log`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_anchor_policy_log` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfAnchorPolicyLogType, cmfAnchorPolicyLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfAnchorPolicyLogType, cmfAnchorPolicyLogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_anchor_policy_log")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfAnchorPolicyLogMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfAnchorPolicyLogType, cmfAnchorPolicyLogMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_anchor_policy_log")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_anchor_policy_log")
	}

CacheNoHooks:
	if !cached {
		cmfAnchorPolicyLogUpsertCacheMut.Lock()
		cmfAnchorPolicyLogUpsertCache[key] = cache
		cmfAnchorPolicyLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFAnchorPolicyLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFAnchorPolicyLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFAnchorPolicyLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfAnchorPolicyLogPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_anchor_policy_log` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_anchor_policy_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_anchor_policy_log")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfAnchorPolicyLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfAnchorPolicyLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_anchor_policy_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_anchor_policy_log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFAnchorPolicyLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfAnchorPolicyLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAnchorPolicyLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_anchor_policy_log` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAnchorPolicyLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfAnchorPolicyLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_anchor_policy_log")
	}

	if len(cmfAnchorPolicyLogAfterDeleteHooks) != 0 {
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
func (o *CMFAnchorPolicyLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFAnchorPolicyLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFAnchorPolicyLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFAnchorPolicyLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfAnchorPolicyLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_anchor_policy_log`.* FROM `cmf_anchor_policy_log` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfAnchorPolicyLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFAnchorPolicyLogSlice")
	}

	*o = slice

	return nil
}

// CMFAnchorPolicyLogExists checks if the CMFAnchorPolicyLog row exists.
func CMFAnchorPolicyLogExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_anchor_policy_log` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_anchor_policy_log exists")
	}

	return exists, nil
}