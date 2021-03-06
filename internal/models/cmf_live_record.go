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

// CMFLiveRecord is an object representing the database table.
type CMFLiveRecord struct {
	ID          uint64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID         int    `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Showid      int64  `boil:"showid" json:"showid" toml:"showid" yaml:"showid"`
	Nums        int    `boil:"nums" json:"nums" toml:"nums" yaml:"nums"`
	Starttime   int    `boil:"starttime" json:"starttime" toml:"starttime" yaml:"starttime"`
	Endtime     int    `boil:"endtime" json:"endtime" toml:"endtime" yaml:"endtime"`
	Title       string `boil:"title" json:"title" toml:"title" yaml:"title"`
	Province    string `boil:"province" json:"province" toml:"province" yaml:"province"`
	City        string `boil:"city" json:"city" toml:"city" yaml:"city"`
	Stream      string `boil:"stream" json:"stream" toml:"stream" yaml:"stream"`
	Thumb       string `boil:"thumb" json:"thumb" toml:"thumb" yaml:"thumb"`
	LNG         string `boil:"lng" json:"lng" toml:"lng" yaml:"lng"`
	Lat         string `boil:"lat" json:"lat" toml:"lat" yaml:"lat"`
	Type        bool   `boil:"type" json:"type" toml:"type" yaml:"type"`
	TypeVal     string `boil:"type_val" json:"type_val" toml:"type_val" yaml:"type_val"`
	Votes       string `boil:"votes" json:"votes" toml:"votes" yaml:"votes"`
	Time        string `boil:"time" json:"time" toml:"time" yaml:"time"`
	Liveclassid int    `boil:"liveclassid" json:"liveclassid" toml:"liveclassid" yaml:"liveclassid"`
	VideoURL    string `boil:"video_url" json:"video_url" toml:"video_url" yaml:"video_url"`

	R *cmfLiveRecordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cmfLiveRecordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CMFLiveRecordColumns = struct {
	ID          string
	UID         string
	Showid      string
	Nums        string
	Starttime   string
	Endtime     string
	Title       string
	Province    string
	City        string
	Stream      string
	Thumb       string
	LNG         string
	Lat         string
	Type        string
	TypeVal     string
	Votes       string
	Time        string
	Liveclassid string
	VideoURL    string
}{
	ID:          "id",
	UID:         "uid",
	Showid:      "showid",
	Nums:        "nums",
	Starttime:   "starttime",
	Endtime:     "endtime",
	Title:       "title",
	Province:    "province",
	City:        "city",
	Stream:      "stream",
	Thumb:       "thumb",
	LNG:         "lng",
	Lat:         "lat",
	Type:        "type",
	TypeVal:     "type_val",
	Votes:       "votes",
	Time:        "time",
	Liveclassid: "liveclassid",
	VideoURL:    "video_url",
}

// Generated where

var CMFLiveRecordWhere = struct {
	ID          whereHelperuint64
	UID         whereHelperint
	Showid      whereHelperint64
	Nums        whereHelperint
	Starttime   whereHelperint
	Endtime     whereHelperint
	Title       whereHelperstring
	Province    whereHelperstring
	City        whereHelperstring
	Stream      whereHelperstring
	Thumb       whereHelperstring
	LNG         whereHelperstring
	Lat         whereHelperstring
	Type        whereHelperbool
	TypeVal     whereHelperstring
	Votes       whereHelperstring
	Time        whereHelperstring
	Liveclassid whereHelperint
	VideoURL    whereHelperstring
}{
	ID:          whereHelperuint64{field: "`cmf_live_record`.`id`"},
	UID:         whereHelperint{field: "`cmf_live_record`.`uid`"},
	Showid:      whereHelperint64{field: "`cmf_live_record`.`showid`"},
	Nums:        whereHelperint{field: "`cmf_live_record`.`nums`"},
	Starttime:   whereHelperint{field: "`cmf_live_record`.`starttime`"},
	Endtime:     whereHelperint{field: "`cmf_live_record`.`endtime`"},
	Title:       whereHelperstring{field: "`cmf_live_record`.`title`"},
	Province:    whereHelperstring{field: "`cmf_live_record`.`province`"},
	City:        whereHelperstring{field: "`cmf_live_record`.`city`"},
	Stream:      whereHelperstring{field: "`cmf_live_record`.`stream`"},
	Thumb:       whereHelperstring{field: "`cmf_live_record`.`thumb`"},
	LNG:         whereHelperstring{field: "`cmf_live_record`.`lng`"},
	Lat:         whereHelperstring{field: "`cmf_live_record`.`lat`"},
	Type:        whereHelperbool{field: "`cmf_live_record`.`type`"},
	TypeVal:     whereHelperstring{field: "`cmf_live_record`.`type_val`"},
	Votes:       whereHelperstring{field: "`cmf_live_record`.`votes`"},
	Time:        whereHelperstring{field: "`cmf_live_record`.`time`"},
	Liveclassid: whereHelperint{field: "`cmf_live_record`.`liveclassid`"},
	VideoURL:    whereHelperstring{field: "`cmf_live_record`.`video_url`"},
}

// CMFLiveRecordRels is where relationship names are stored.
var CMFLiveRecordRels = struct {
}{}

// cmfLiveRecordR is where relationships are stored.
type cmfLiveRecordR struct {
}

// NewStruct creates a new relationship struct
func (*cmfLiveRecordR) NewStruct() *cmfLiveRecordR {
	return &cmfLiveRecordR{}
}

// cmfLiveRecordL is where Load methods for each relationship are stored.
type cmfLiveRecordL struct{}

var (
	cmfLiveRecordAllColumns            = []string{"id", "uid", "showid", "nums", "starttime", "endtime", "title", "province", "city", "stream", "thumb", "lng", "lat", "type", "type_val", "votes", "time", "liveclassid", "video_url"}
	cmfLiveRecordColumnsWithoutDefault = []string{"showid", "title", "province", "city", "stream", "thumb", "lng", "lat", "type_val", "time", "video_url"}
	cmfLiveRecordColumnsWithDefault    = []string{"id", "uid", "nums", "starttime", "endtime", "type", "votes", "liveclassid"}
	cmfLiveRecordPrimaryKeyColumns     = []string{"id"}
)

type (
	// CMFLiveRecordSlice is an alias for a slice of pointers to CMFLiveRecord.
	// This should generally be used opposed to []CMFLiveRecord.
	CMFLiveRecordSlice []*CMFLiveRecord
	// CMFLiveRecordHook is the signature for custom CMFLiveRecord hook methods
	CMFLiveRecordHook func(context.Context, boil.ContextExecutor, *CMFLiveRecord) error

	cmfLiveRecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cmfLiveRecordType                 = reflect.TypeOf(&CMFLiveRecord{})
	cmfLiveRecordMapping              = queries.MakeStructMapping(cmfLiveRecordType)
	cmfLiveRecordPrimaryKeyMapping, _ = queries.BindMapping(cmfLiveRecordType, cmfLiveRecordMapping, cmfLiveRecordPrimaryKeyColumns)
	cmfLiveRecordInsertCacheMut       sync.RWMutex
	cmfLiveRecordInsertCache          = make(map[string]insertCache)
	cmfLiveRecordUpdateCacheMut       sync.RWMutex
	cmfLiveRecordUpdateCache          = make(map[string]updateCache)
	cmfLiveRecordUpsertCacheMut       sync.RWMutex
	cmfLiveRecordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cmfLiveRecordBeforeInsertHooks []CMFLiveRecordHook
var cmfLiveRecordBeforeUpdateHooks []CMFLiveRecordHook
var cmfLiveRecordBeforeDeleteHooks []CMFLiveRecordHook
var cmfLiveRecordBeforeUpsertHooks []CMFLiveRecordHook

var cmfLiveRecordAfterInsertHooks []CMFLiveRecordHook
var cmfLiveRecordAfterSelectHooks []CMFLiveRecordHook
var cmfLiveRecordAfterUpdateHooks []CMFLiveRecordHook
var cmfLiveRecordAfterDeleteHooks []CMFLiveRecordHook
var cmfLiveRecordAfterUpsertHooks []CMFLiveRecordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CMFLiveRecord) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CMFLiveRecord) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CMFLiveRecord) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CMFLiveRecord) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CMFLiveRecord) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CMFLiveRecord) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CMFLiveRecord) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CMFLiveRecord) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CMFLiveRecord) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cmfLiveRecordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCMFLiveRecordHook registers your hook function for all future operations.
func AddCMFLiveRecordHook(hookPoint boil.HookPoint, cmfLiveRecordHook CMFLiveRecordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cmfLiveRecordBeforeInsertHooks = append(cmfLiveRecordBeforeInsertHooks, cmfLiveRecordHook)
	case boil.BeforeUpdateHook:
		cmfLiveRecordBeforeUpdateHooks = append(cmfLiveRecordBeforeUpdateHooks, cmfLiveRecordHook)
	case boil.BeforeDeleteHook:
		cmfLiveRecordBeforeDeleteHooks = append(cmfLiveRecordBeforeDeleteHooks, cmfLiveRecordHook)
	case boil.BeforeUpsertHook:
		cmfLiveRecordBeforeUpsertHooks = append(cmfLiveRecordBeforeUpsertHooks, cmfLiveRecordHook)
	case boil.AfterInsertHook:
		cmfLiveRecordAfterInsertHooks = append(cmfLiveRecordAfterInsertHooks, cmfLiveRecordHook)
	case boil.AfterSelectHook:
		cmfLiveRecordAfterSelectHooks = append(cmfLiveRecordAfterSelectHooks, cmfLiveRecordHook)
	case boil.AfterUpdateHook:
		cmfLiveRecordAfterUpdateHooks = append(cmfLiveRecordAfterUpdateHooks, cmfLiveRecordHook)
	case boil.AfterDeleteHook:
		cmfLiveRecordAfterDeleteHooks = append(cmfLiveRecordAfterDeleteHooks, cmfLiveRecordHook)
	case boil.AfterUpsertHook:
		cmfLiveRecordAfterUpsertHooks = append(cmfLiveRecordAfterUpsertHooks, cmfLiveRecordHook)
	}
}

// One returns a single cmfLiveRecord record from the query.
func (q cmfLiveRecordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CMFLiveRecord, error) {
	o := &CMFLiveRecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cmf_live_record")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CMFLiveRecord records from the query.
func (q cmfLiveRecordQuery) All(ctx context.Context, exec boil.ContextExecutor) (CMFLiveRecordSlice, error) {
	var o []*CMFLiveRecord

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CMFLiveRecord slice")
	}

	if len(cmfLiveRecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CMFLiveRecord records in the query.
func (q cmfLiveRecordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cmf_live_record rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cmfLiveRecordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cmf_live_record exists")
	}

	return count > 0, nil
}

// CMFLiveRecords retrieves all the records using an executor.
func CMFLiveRecords(mods ...qm.QueryMod) cmfLiveRecordQuery {
	mods = append(mods, qm.From("`cmf_live_record`"))
	return cmfLiveRecordQuery{NewQuery(mods...)}
}

// FindCMFLiveRecord retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCMFLiveRecord(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*CMFLiveRecord, error) {
	cmfLiveRecordObj := &CMFLiveRecord{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cmf_live_record` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cmfLiveRecordObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cmf_live_record")
	}

	return cmfLiveRecordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CMFLiveRecord) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_live_record provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfLiveRecordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cmfLiveRecordInsertCacheMut.RLock()
	cache, cached := cmfLiveRecordInsertCache[key]
	cmfLiveRecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cmfLiveRecordAllColumns,
			cmfLiveRecordColumnsWithDefault,
			cmfLiveRecordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cmfLiveRecordType, cmfLiveRecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cmfLiveRecordType, cmfLiveRecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cmf_live_record` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cmf_live_record` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cmf_live_record` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cmfLiveRecordPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into cmf_live_record")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfLiveRecordMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for cmf_live_record")
	}

CacheNoHooks:
	if !cached {
		cmfLiveRecordInsertCacheMut.Lock()
		cmfLiveRecordInsertCache[key] = cache
		cmfLiveRecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CMFLiveRecord.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CMFLiveRecord) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cmfLiveRecordUpdateCacheMut.RLock()
	cache, cached := cmfLiveRecordUpdateCache[key]
	cmfLiveRecordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cmfLiveRecordAllColumns,
			cmfLiveRecordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cmf_live_record, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cmf_live_record` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cmfLiveRecordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cmfLiveRecordType, cmfLiveRecordMapping, append(wl, cmfLiveRecordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cmf_live_record row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cmf_live_record")
	}

	if !cached {
		cmfLiveRecordUpdateCacheMut.Lock()
		cmfLiveRecordUpdateCache[key] = cache
		cmfLiveRecordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cmfLiveRecordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cmf_live_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cmf_live_record")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CMFLiveRecordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLiveRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cmf_live_record` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLiveRecordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cmfLiveRecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cmfLiveRecord")
	}
	return rowsAff, nil
}

var mySQLCMFLiveRecordUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CMFLiveRecord) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cmf_live_record provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cmfLiveRecordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCMFLiveRecordUniqueColumns, o)

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

	cmfLiveRecordUpsertCacheMut.RLock()
	cache, cached := cmfLiveRecordUpsertCache[key]
	cmfLiveRecordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cmfLiveRecordAllColumns,
			cmfLiveRecordColumnsWithDefault,
			cmfLiveRecordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cmfLiveRecordAllColumns,
			cmfLiveRecordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert cmf_live_record, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cmf_live_record`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cmf_live_record` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cmfLiveRecordType, cmfLiveRecordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cmfLiveRecordType, cmfLiveRecordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cmf_live_record")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cmfLiveRecordMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cmfLiveRecordType, cmfLiveRecordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for cmf_live_record")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for cmf_live_record")
	}

CacheNoHooks:
	if !cached {
		cmfLiveRecordUpsertCacheMut.Lock()
		cmfLiveRecordUpsertCache[key] = cache
		cmfLiveRecordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CMFLiveRecord record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CMFLiveRecord) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CMFLiveRecord provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cmfLiveRecordPrimaryKeyMapping)
	sql := "DELETE FROM `cmf_live_record` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cmf_live_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cmf_live_record")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cmfLiveRecordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cmfLiveRecordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmf_live_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_live_record")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CMFLiveRecordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cmfLiveRecordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLiveRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cmf_live_record` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLiveRecordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cmfLiveRecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cmf_live_record")
	}

	if len(cmfLiveRecordAfterDeleteHooks) != 0 {
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
func (o *CMFLiveRecord) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCMFLiveRecord(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CMFLiveRecordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CMFLiveRecordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cmfLiveRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cmf_live_record`.* FROM `cmf_live_record` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cmfLiveRecordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CMFLiveRecordSlice")
	}

	*o = slice

	return nil
}

// CMFLiveRecordExists checks if the CMFLiveRecord row exists.
func CMFLiveRecordExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cmf_live_record` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cmf_live_record exists")
	}

	return exists, nil
}
