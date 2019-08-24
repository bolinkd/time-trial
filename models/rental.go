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

// Rental is an object representing the database table.
type Rental struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ShellID   null.Int  `boil:"shell_id" json:"shell_id,omitempty" toml:"shell_id" yaml:"shell_id,omitempty"`
	OutTime   null.Time `boil:"out_time" json:"out_time,omitempty" toml:"out_time" yaml:"out_time,omitempty"`
	InTime    null.Time `boil:"in_time" json:"in_time,omitempty" toml:"in_time" yaml:"in_time,omitempty"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *rentalR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L rentalL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RentalColumns = struct {
	ID        string
	ShellID   string
	OutTime   string
	InTime    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	ShellID:   "shell_id",
	OutTime:   "out_time",
	InTime:    "in_time",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var RentalWhere = struct {
	ID        whereHelperint
	ShellID   whereHelpernull_Int
	OutTime   whereHelpernull_Time
	InTime    whereHelpernull_Time
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: `id`},
	ShellID:   whereHelpernull_Int{field: `shell_id`},
	OutTime:   whereHelpernull_Time{field: `out_time`},
	InTime:    whereHelpernull_Time{field: `in_time`},
	CreatedAt: whereHelpertime_Time{field: `created_at`},
	UpdatedAt: whereHelpertime_Time{field: `updated_at`},
}

// RentalRels is where relationship names are stored.
var RentalRels = struct {
	Shell        string
	RentalRowers string
}{
	Shell:        "Shell",
	RentalRowers: "RentalRowers",
}

// rentalR is where relationships are stored.
type rentalR struct {
	Shell        *Shell
	RentalRowers RentalRowerSlice
}

// NewStruct creates a new relationship struct
func (*rentalR) NewStruct() *rentalR {
	return &rentalR{}
}

// rentalL is where Load methods for each relationship are stored.
type rentalL struct{}

var (
	rentalColumns               = []string{"id", "shell_id", "out_time", "in_time", "created_at", "updated_at"}
	rentalColumnsWithoutDefault = []string{"shell_id", "out_time", "in_time", "created_at", "updated_at"}
	rentalColumnsWithDefault    = []string{"id"}
	rentalPrimaryKeyColumns     = []string{"id"}
)

type (
	// RentalSlice is an alias for a slice of pointers to Rental.
	// This should generally be used opposed to []Rental.
	RentalSlice []*Rental
	// RentalHook is the signature for custom Rental hook methods
	RentalHook func(boil.Executor, *Rental) error

	rentalQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	rentalType                 = reflect.TypeOf(&Rental{})
	rentalMapping              = queries.MakeStructMapping(rentalType)
	rentalPrimaryKeyMapping, _ = queries.BindMapping(rentalType, rentalMapping, rentalPrimaryKeyColumns)
	rentalInsertCacheMut       sync.RWMutex
	rentalInsertCache          = make(map[string]insertCache)
	rentalUpdateCacheMut       sync.RWMutex
	rentalUpdateCache          = make(map[string]updateCache)
	rentalUpsertCacheMut       sync.RWMutex
	rentalUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var rentalBeforeInsertHooks []RentalHook
var rentalBeforeUpdateHooks []RentalHook
var rentalBeforeDeleteHooks []RentalHook
var rentalBeforeUpsertHooks []RentalHook

var rentalAfterInsertHooks []RentalHook
var rentalAfterSelectHooks []RentalHook
var rentalAfterUpdateHooks []RentalHook
var rentalAfterDeleteHooks []RentalHook
var rentalAfterUpsertHooks []RentalHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Rental) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Rental) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Rental) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Rental) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Rental) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Rental) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Rental) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Rental) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Rental) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range rentalAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRentalHook registers your hook function for all future operations.
func AddRentalHook(hookPoint boil.HookPoint, rentalHook RentalHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		rentalBeforeInsertHooks = append(rentalBeforeInsertHooks, rentalHook)
	case boil.BeforeUpdateHook:
		rentalBeforeUpdateHooks = append(rentalBeforeUpdateHooks, rentalHook)
	case boil.BeforeDeleteHook:
		rentalBeforeDeleteHooks = append(rentalBeforeDeleteHooks, rentalHook)
	case boil.BeforeUpsertHook:
		rentalBeforeUpsertHooks = append(rentalBeforeUpsertHooks, rentalHook)
	case boil.AfterInsertHook:
		rentalAfterInsertHooks = append(rentalAfterInsertHooks, rentalHook)
	case boil.AfterSelectHook:
		rentalAfterSelectHooks = append(rentalAfterSelectHooks, rentalHook)
	case boil.AfterUpdateHook:
		rentalAfterUpdateHooks = append(rentalAfterUpdateHooks, rentalHook)
	case boil.AfterDeleteHook:
		rentalAfterDeleteHooks = append(rentalAfterDeleteHooks, rentalHook)
	case boil.AfterUpsertHook:
		rentalAfterUpsertHooks = append(rentalAfterUpsertHooks, rentalHook)
	}
}

// OneG returns a single rental record from the query using the global executor.
func (q rentalQuery) OneG() (*Rental, error) {
	return q.One(boil.GetDB())
}

// One returns a single rental record from the query.
func (q rentalQuery) One(exec boil.Executor) (*Rental, error) {
	o := &Rental{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for rental")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Rental records from the query using the global executor.
func (q rentalQuery) AllG() (RentalSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Rental records from the query.
func (q rentalQuery) All(exec boil.Executor) (RentalSlice, error) {
	var o []*Rental

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Rental slice")
	}

	if len(rentalAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Rental records in the query, and panics on error.
func (q rentalQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Rental records in the query.
func (q rentalQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count rental rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q rentalQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q rentalQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if rental exists")
	}

	return count > 0, nil
}

// Shell pointed to by the foreign key.
func (o *Rental) Shell(mods ...qm.QueryMod) shellQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ShellID),
	}

	queryMods = append(queryMods, mods...)

	query := Shells(queryMods...)
	queries.SetFrom(query.Query, "\"shell\"")

	return query
}

// RentalRowers retrieves all the rental_rower's RentalRowers with an executor.
func (o *Rental) RentalRowers(mods ...qm.QueryMod) rentalRowerQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"rental_rowers\".\"rental_id\"=?", o.ID),
	)

	query := RentalRowers(queryMods...)
	queries.SetFrom(query.Query, "\"rental_rowers\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"rental_rowers\".*"})
	}

	return query
}

// LoadShell allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (rentalL) LoadShell(e boil.Executor, singular bool, maybeRental interface{}, mods queries.Applicator) error {
	var slice []*Rental
	var object *Rental

	if singular {
		object = maybeRental.(*Rental)
	} else {
		slice = *maybeRental.(*[]*Rental)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &rentalR{}
		}
		if !queries.IsNil(object.ShellID) {
			args = append(args, object.ShellID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &rentalR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ShellID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ShellID) {
				args = append(args, obj.ShellID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`shell`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Shell")
	}

	var resultSlice []*Shell
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Shell")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for shell")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for shell")
	}

	if len(rentalAfterSelectHooks) != 0 {
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
		object.R.Shell = foreign
		if foreign.R == nil {
			foreign.R = &shellR{}
		}
		foreign.R.Rentals = append(foreign.R.Rentals, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ShellID, foreign.ID) {
				local.R.Shell = foreign
				if foreign.R == nil {
					foreign.R = &shellR{}
				}
				foreign.R.Rentals = append(foreign.R.Rentals, local)
				break
			}
		}
	}

	return nil
}

// LoadRentalRowers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (rentalL) LoadRentalRowers(e boil.Executor, singular bool, maybeRental interface{}, mods queries.Applicator) error {
	var slice []*Rental
	var object *Rental

	if singular {
		object = maybeRental.(*Rental)
	} else {
		slice = *maybeRental.(*[]*Rental)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &rentalR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &rentalR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`rental_rowers`), qm.WhereIn(`rental_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load rental_rowers")
	}

	var resultSlice []*RentalRower
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice rental_rowers")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on rental_rowers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for rental_rowers")
	}

	if len(rentalRowerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RentalRowers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &rentalRowerR{}
			}
			foreign.R.Rental = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.RentalID) {
				local.R.RentalRowers = append(local.R.RentalRowers, foreign)
				if foreign.R == nil {
					foreign.R = &rentalRowerR{}
				}
				foreign.R.Rental = local
				break
			}
		}
	}

	return nil
}

// SetShellG of the rental to the related item.
// Sets o.R.Shell to related.
// Adds o to related.R.Rentals.
// Uses the global database handle.
func (o *Rental) SetShellG(insert bool, related *Shell) error {
	return o.SetShell(boil.GetDB(), insert, related)
}

// SetShell of the rental to the related item.
// Sets o.R.Shell to related.
// Adds o to related.R.Rentals.
func (o *Rental) SetShell(exec boil.Executor, insert bool, related *Shell) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"rental\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"shell_id"}),
		strmangle.WhereClause("\"", "\"", 2, rentalPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ShellID, related.ID)
	if o.R == nil {
		o.R = &rentalR{
			Shell: related,
		}
	} else {
		o.R.Shell = related
	}

	if related.R == nil {
		related.R = &shellR{
			Rentals: RentalSlice{o},
		}
	} else {
		related.R.Rentals = append(related.R.Rentals, o)
	}

	return nil
}

// RemoveShellG relationship.
// Sets o.R.Shell to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *Rental) RemoveShellG(related *Shell) error {
	return o.RemoveShell(boil.GetDB(), related)
}

// RemoveShell relationship.
// Sets o.R.Shell to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Rental) RemoveShell(exec boil.Executor, related *Shell) error {
	var err error

	queries.SetScanner(&o.ShellID, nil)
	if _, err = o.Update(exec, boil.Whitelist("shell_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Shell = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Rentals {
		if queries.Equal(o.ShellID, ri.ShellID) {
			continue
		}

		ln := len(related.R.Rentals)
		if ln > 1 && i < ln-1 {
			related.R.Rentals[i] = related.R.Rentals[ln-1]
		}
		related.R.Rentals = related.R.Rentals[:ln-1]
		break
	}
	return nil
}

// AddRentalRowersG adds the given related objects to the existing relationships
// of the rental, optionally inserting them as new records.
// Appends related to o.R.RentalRowers.
// Sets related.R.Rental appropriately.
// Uses the global database handle.
func (o *Rental) AddRentalRowersG(insert bool, related ...*RentalRower) error {
	return o.AddRentalRowers(boil.GetDB(), insert, related...)
}

// AddRentalRowers adds the given related objects to the existing relationships
// of the rental, optionally inserting them as new records.
// Appends related to o.R.RentalRowers.
// Sets related.R.Rental appropriately.
func (o *Rental) AddRentalRowers(exec boil.Executor, insert bool, related ...*RentalRower) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.RentalID, o.ID)
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"rental_rowers\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"rental_id"}),
				strmangle.WhereClause("\"", "\"", 2, rentalRowerPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.RentalID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &rentalR{
			RentalRowers: related,
		}
	} else {
		o.R.RentalRowers = append(o.R.RentalRowers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &rentalRowerR{
				Rental: o,
			}
		} else {
			rel.R.Rental = o
		}
	}
	return nil
}

// SetRentalRowersG removes all previously related items of the
// rental replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Rental's RentalRowers accordingly.
// Replaces o.R.RentalRowers with related.
// Sets related.R.Rental's RentalRowers accordingly.
// Uses the global database handle.
func (o *Rental) SetRentalRowersG(insert bool, related ...*RentalRower) error {
	return o.SetRentalRowers(boil.GetDB(), insert, related...)
}

// SetRentalRowers removes all previously related items of the
// rental replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Rental's RentalRowers accordingly.
// Replaces o.R.RentalRowers with related.
// Sets related.R.Rental's RentalRowers accordingly.
func (o *Rental) SetRentalRowers(exec boil.Executor, insert bool, related ...*RentalRower) error {
	query := "update \"rental_rowers\" set \"rental_id\" = null where \"rental_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.RentalRowers {
			queries.SetScanner(&rel.RentalID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Rental = nil
		}

		o.R.RentalRowers = nil
	}
	return o.AddRentalRowers(exec, insert, related...)
}

// RemoveRentalRowersG relationships from objects passed in.
// Removes related items from R.RentalRowers (uses pointer comparison, removal does not keep order)
// Sets related.R.Rental.
// Uses the global database handle.
func (o *Rental) RemoveRentalRowersG(related ...*RentalRower) error {
	return o.RemoveRentalRowers(boil.GetDB(), related...)
}

// RemoveRentalRowers relationships from objects passed in.
// Removes related items from R.RentalRowers (uses pointer comparison, removal does not keep order)
// Sets related.R.Rental.
func (o *Rental) RemoveRentalRowers(exec boil.Executor, related ...*RentalRower) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.RentalID, nil)
		if rel.R != nil {
			rel.R.Rental = nil
		}
		if _, err = rel.Update(exec, boil.Whitelist("rental_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.RentalRowers {
			if rel != ri {
				continue
			}

			ln := len(o.R.RentalRowers)
			if ln > 1 && i < ln-1 {
				o.R.RentalRowers[i] = o.R.RentalRowers[ln-1]
			}
			o.R.RentalRowers = o.R.RentalRowers[:ln-1]
			break
		}
	}

	return nil
}

// Rentals retrieves all the records using an executor.
func Rentals(mods ...qm.QueryMod) rentalQuery {
	mods = append(mods, qm.From("\"rental\""))
	return rentalQuery{NewQuery(mods...)}
}

// FindRentalG retrieves a single record by ID.
func FindRentalG(iD int, selectCols ...string) (*Rental, error) {
	return FindRental(boil.GetDB(), iD, selectCols...)
}

// FindRental retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRental(exec boil.Executor, iD int, selectCols ...string) (*Rental, error) {
	rentalObj := &Rental{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"rental\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, rentalObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from rental")
	}

	return rentalObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Rental) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Rental) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rental provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(rentalColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	rentalInsertCacheMut.RLock()
	cache, cached := rentalInsertCache[key]
	rentalInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			rentalColumns,
			rentalColumnsWithDefault,
			rentalColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(rentalType, rentalMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(rentalType, rentalMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"rental\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"rental\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into rental")
	}

	if !cached {
		rentalInsertCacheMut.Lock()
		rentalInsertCache[key] = cache
		rentalInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Rental record using the global executor.
// See Update for more documentation.
func (o *Rental) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Rental.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Rental) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	rentalUpdateCacheMut.RLock()
	cache, cached := rentalUpdateCache[key]
	rentalUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			rentalColumns,
			rentalPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update rental, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"rental\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rentalPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(rentalType, rentalMapping, append(wl, rentalPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update rental row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for rental")
	}

	if !cached {
		rentalUpdateCacheMut.Lock()
		rentalUpdateCache[key] = cache
		rentalUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q rentalQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q rentalQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for rental")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for rental")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o RentalSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RentalSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rentalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"rental\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rentalPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in rental slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all rental")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Rental) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Rental) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rental provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rentalColumnsWithDefault, o)

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

	rentalUpsertCacheMut.RLock()
	cache, cached := rentalUpsertCache[key]
	rentalUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			rentalColumns,
			rentalColumnsWithDefault,
			rentalColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			rentalColumns,
			rentalPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert rental, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rentalPrimaryKeyColumns))
			copy(conflict, rentalPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"rental\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(rentalType, rentalMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(rentalType, rentalMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert rental")
	}

	if !cached {
		rentalUpsertCacheMut.Lock()
		rentalUpsertCache[key] = cache
		rentalUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Rental record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Rental) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Rental record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Rental) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Rental provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rentalPrimaryKeyMapping)
	sql := "DELETE FROM \"rental\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from rental")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for rental")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q rentalQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no rentalQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rental")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for rental")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o RentalSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RentalSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Rental slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(rentalBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rentalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"rental\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rentalPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rental slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for rental")
	}

	if len(rentalAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Rental) ReloadG() error {
	if o == nil {
		return errors.New("models: no Rental provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Rental) Reload(exec boil.Executor) error {
	ret, err := FindRental(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RentalSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty RentalSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RentalSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RentalSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rentalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"rental\".* FROM \"rental\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rentalPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RentalSlice")
	}

	*o = slice

	return nil
}

// RentalExistsG checks if the Rental row exists.
func RentalExistsG(iD int) (bool, error) {
	return RentalExists(boil.GetDB(), iD)
}

// RentalExists checks if the Rental row exists.
func RentalExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"rental\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if rental exists")
	}

	return exists, nil
}