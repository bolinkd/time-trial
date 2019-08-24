// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Boat is an object representing the database table.
type Boat struct {
	ID          int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TimeTrialID null.Int  `boil:"time_trial_id" json:"time_trial_id,omitempty" toml:"time_trial_id" yaml:"time_trial_id,omitempty"`
	BowMarker   null.Int  `boil:"bow_marker" json:"bow_marker,omitempty" toml:"bow_marker" yaml:"bow_marker,omitempty"`
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Start       null.Int  `boil:"start" json:"start,omitempty" toml:"start" yaml:"start,omitempty"`
	End         null.Int  `boil:"end" json:"end,omitempty" toml:"end" yaml:"end,omitempty"`
	Time        null.Int  `boil:"time" json:"time,omitempty" toml:"time" yaml:"time,omitempty"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *boatR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L boatL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BoatColumns = struct {
	ID          string
	TimeTrialID string
	BowMarker   string
	Name        string
	Start       string
	End         string
	Time        string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	TimeTrialID: "time_trial_id",
	BowMarker:   "bow_marker",
	Name:        "name",
	Start:       "start",
	End:         "end",
	Time:        "time",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var BoatWhere = struct {
	ID          whereHelperint
	TimeTrialID whereHelpernull_Int
	BowMarker   whereHelpernull_Int
	Name        whereHelperstring
	Start       whereHelpernull_Int
	End         whereHelpernull_Int
	Time        whereHelpernull_Int
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint{field: `id`},
	TimeTrialID: whereHelpernull_Int{field: `time_trial_id`},
	BowMarker:   whereHelpernull_Int{field: `bow_marker`},
	Name:        whereHelperstring{field: `name`},
	Start:       whereHelpernull_Int{field: `start`},
	End:         whereHelpernull_Int{field: `end`},
	Time:        whereHelpernull_Int{field: `time`},
	CreatedAt:   whereHelpertime_Time{field: `created_at`},
	UpdatedAt:   whereHelpertime_Time{field: `updated_at`},
}

// BoatRels is where relationship names are stored.
var BoatRels = struct {
	TimeTrial string
}{
	TimeTrial: "TimeTrial",
}

// boatR is where relationships are stored.
type boatR struct {
	TimeTrial *TimeTrial
}

// NewStruct creates a new relationship struct
func (*boatR) NewStruct() *boatR {
	return &boatR{}
}

// boatL is where Load methods for each relationship are stored.
type boatL struct{}

var (
	boatColumns               = []string{"id", "time_trial_id", "bow_marker", "name", "start", "end", "time", "created_at", "updated_at"}
	boatColumnsWithoutDefault = []string{"time_trial_id", "bow_marker", "name", "start", "end", "time", "created_at", "updated_at"}
	boatColumnsWithDefault    = []string{"id"}
	boatPrimaryKeyColumns     = []string{"id"}
)

type (
	// BoatSlice is an alias for a slice of pointers to Boat.
	// This should generally be used opposed to []Boat.
	BoatSlice []*Boat
	// BoatHook is the signature for custom Boat hook methods
	BoatHook func(boil.Executor, *Boat) error

	boatQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	boatType                 = reflect.TypeOf(&Boat{})
	boatMapping              = queries.MakeStructMapping(boatType)
	boatPrimaryKeyMapping, _ = queries.BindMapping(boatType, boatMapping, boatPrimaryKeyColumns)
	boatInsertCacheMut       sync.RWMutex
	boatInsertCache          = make(map[string]insertCache)
	boatUpdateCacheMut       sync.RWMutex
	boatUpdateCache          = make(map[string]updateCache)
	boatUpsertCacheMut       sync.RWMutex
	boatUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var boatBeforeInsertHooks []BoatHook
var boatBeforeUpdateHooks []BoatHook
var boatBeforeDeleteHooks []BoatHook
var boatBeforeUpsertHooks []BoatHook

var boatAfterInsertHooks []BoatHook
var boatAfterSelectHooks []BoatHook
var boatAfterUpdateHooks []BoatHook
var boatAfterDeleteHooks []BoatHook
var boatAfterUpsertHooks []BoatHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Boat) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boatBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Boat) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range boatBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Boat) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range boatBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Boat) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boatBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Boat) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boatAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Boat) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range boatAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Boat) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range boatAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Boat) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range boatAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Boat) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boatAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBoatHook registers your hook function for all future operations.
func AddBoatHook(hookPoint boil.HookPoint, boatHook BoatHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		boatBeforeInsertHooks = append(boatBeforeInsertHooks, boatHook)
	case boil.BeforeUpdateHook:
		boatBeforeUpdateHooks = append(boatBeforeUpdateHooks, boatHook)
	case boil.BeforeDeleteHook:
		boatBeforeDeleteHooks = append(boatBeforeDeleteHooks, boatHook)
	case boil.BeforeUpsertHook:
		boatBeforeUpsertHooks = append(boatBeforeUpsertHooks, boatHook)
	case boil.AfterInsertHook:
		boatAfterInsertHooks = append(boatAfterInsertHooks, boatHook)
	case boil.AfterSelectHook:
		boatAfterSelectHooks = append(boatAfterSelectHooks, boatHook)
	case boil.AfterUpdateHook:
		boatAfterUpdateHooks = append(boatAfterUpdateHooks, boatHook)
	case boil.AfterDeleteHook:
		boatAfterDeleteHooks = append(boatAfterDeleteHooks, boatHook)
	case boil.AfterUpsertHook:
		boatAfterUpsertHooks = append(boatAfterUpsertHooks, boatHook)
	}
}

// OneG returns a single boat record from the query using the global executor.
func (q boatQuery) OneG() (*Boat, error) {
	return q.One(boil.GetDB())
}

// One returns a single boat record from the query.
func (q boatQuery) One(exec boil.Executor) (*Boat, error) {
	o := &Boat{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for boat")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Boat records from the query using the global executor.
func (q boatQuery) AllG() (BoatSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Boat records from the query.
func (q boatQuery) All(exec boil.Executor) (BoatSlice, error) {
	var o []*Boat

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Boat slice")
	}

	if len(boatAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Boat records in the query, and panics on error.
func (q boatQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Boat records in the query.
func (q boatQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count boat rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q boatQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q boatQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if boat exists")
	}

	return count > 0, nil
}

// TimeTrial pointed to by the foreign key.
func (o *Boat) TimeTrial(mods ...qm.QueryMod) timeTrialQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TimeTrialID),
	}

	queryMods = append(queryMods, mods...)

	query := TimeTrials(queryMods...)
	queries.SetFrom(query.Query, "\"time_trial\"")

	return query
}

// LoadTimeTrial allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (boatL) LoadTimeTrial(e boil.Executor, singular bool, maybeBoat interface{}, mods queries.Applicator) error {
	var slice []*Boat
	var object *Boat

	if singular {
		object = maybeBoat.(*Boat)
	} else {
		slice = *maybeBoat.(*[]*Boat)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &boatR{}
		}
		if !queries.IsNil(object.TimeTrialID) {
			args = append(args, object.TimeTrialID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &boatR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.TimeTrialID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.TimeTrialID) {
				args = append(args, obj.TimeTrialID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`time_trial`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TimeTrial")
	}

	var resultSlice []*TimeTrial
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TimeTrial")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for time_trial")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for time_trial")
	}

	if len(boatAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.TimeTrial = foreign
		if foreign.R == nil {
			foreign.R = &timeTrialR{}
		}
		foreign.R.Boats = append(foreign.R.Boats, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.TimeTrialID, foreign.ID) {
				local.R.TimeTrial = foreign
				if foreign.R == nil {
					foreign.R = &timeTrialR{}
				}
				foreign.R.Boats = append(foreign.R.Boats, local)
				break
			}
		}
	}

	return nil
}

// SetTimeTrialG of the boat to the related item.
// Sets o.R.TimeTrial to related.
// Adds o to related.R.Boats.
// Uses the global database handle.
func (o *Boat) SetTimeTrialG(insert bool, related *TimeTrial) error {
	return o.SetTimeTrial(boil.GetDB(), insert, related)
}

// SetTimeTrial of the boat to the related item.
// Sets o.R.TimeTrial to related.
// Adds o to related.R.Boats.
func (o *Boat) SetTimeTrial(exec boil.Executor, insert bool, related *TimeTrial) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"boat\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"time_trial_id"}),
		strmangle.WhereClause("\"", "\"", 2, boatPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.TimeTrialID, related.ID)
	if o.R == nil {
		o.R = &boatR{
			TimeTrial: related,
		}
	} else {
		o.R.TimeTrial = related
	}

	if related.R == nil {
		related.R = &timeTrialR{
			Boats: BoatSlice{o},
		}
	} else {
		related.R.Boats = append(related.R.Boats, o)
	}

	return nil
}

// RemoveTimeTrialG relationship.
// Sets o.R.TimeTrial to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *Boat) RemoveTimeTrialG(related *TimeTrial) error {
	return o.RemoveTimeTrial(boil.GetDB(), related)
}

// RemoveTimeTrial relationship.
// Sets o.R.TimeTrial to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Boat) RemoveTimeTrial(exec boil.Executor, related *TimeTrial) error {
	var err error

	queries.SetScanner(&o.TimeTrialID, nil)
	if _, err = o.Update(exec, boil.Whitelist("time_trial_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.TimeTrial = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Boats {
		if queries.Equal(o.TimeTrialID, ri.TimeTrialID) {
			continue
		}

		ln := len(related.R.Boats)
		if ln > 1 && i < ln-1 {
			related.R.Boats[i] = related.R.Boats[ln-1]
		}
		related.R.Boats = related.R.Boats[:ln-1]
		break
	}
	return nil
}

// Boats retrieves all the records using an executor.
func Boats(mods ...qm.QueryMod) boatQuery {
	mods = append(mods, qm.From("\"boat\""))
	return boatQuery{NewQuery(mods...)}
}

// FindBoatG retrieves a single record by ID.
func FindBoatG(iD int, selectCols ...string) (*Boat, error) {
	return FindBoat(boil.GetDB(), iD, selectCols...)
}

// FindBoat retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBoat(exec boil.Executor, iD int, selectCols ...string) (*Boat, error) {
	boatObj := &Boat{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"boat\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, boatObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from boat")
	}

	return boatObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Boat) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Boat) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no boat provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(boatColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	boatInsertCacheMut.RLock()
	cache, cached := boatInsertCache[key]
	boatInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			boatColumns,
			boatColumnsWithDefault,
			boatColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(boatType, boatMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(boatType, boatMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"boat\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"boat\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into boat")
	}

	if !cached {
		boatInsertCacheMut.Lock()
		boatInsertCache[key] = cache
		boatInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Boat record using the global executor.
// See Update for more documentation.
func (o *Boat) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Boat.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Boat) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	boatUpdateCacheMut.RLock()
	cache, cached := boatUpdateCache[key]
	boatUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			boatColumns,
			boatPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update boat, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"boat\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, boatPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(boatType, boatMapping, append(wl, boatPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update boat row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for boat")
	}

	if !cached {
		boatUpdateCacheMut.Lock()
		boatUpdateCache[key] = cache
		boatUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q boatQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q boatQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for boat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for boat")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BoatSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BoatSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"boat\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, boatPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in boat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all boat")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Boat) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Boat) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no boat provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(boatColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	boatUpsertCacheMut.RLock()
	cache, cached := boatUpsertCache[key]
	boatUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			boatColumns,
			boatColumnsWithDefault,
			boatColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			boatColumns,
			boatPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert boat, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(boatPrimaryKeyColumns))
			copy(conflict, boatPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"boat\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(boatType, boatMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(boatType, boatMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert boat")
	}

	if !cached {
		boatUpsertCacheMut.Lock()
		boatUpsertCache[key] = cache
		boatUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Boat record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Boat) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Boat record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Boat) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Boat provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), boatPrimaryKeyMapping)
	sql := "DELETE FROM \"boat\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from boat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for boat")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q boatQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no boatQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from boat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for boat")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BoatSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BoatSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Boat slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(boatBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"boat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, boatPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from boat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for boat")
	}

	if len(boatAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Boat) ReloadG() error {
	if o == nil {
		return errors.New("models: no Boat provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Boat) Reload(exec boil.Executor) error {
	ret, err := FindBoat(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BoatSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty BoatSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BoatSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BoatSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"boat\".* FROM \"boat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, boatPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BoatSlice")
	}

	*o = slice

	return nil
}

// BoatExistsG checks if the Boat row exists.
func BoatExistsG(iD int) (bool, error) {
	return BoatExists(boil.GetDB(), iD)
}

// BoatExists checks if the Boat row exists.
func BoatExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"boat\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if boat exists")
	}

	return exists, nil
}
