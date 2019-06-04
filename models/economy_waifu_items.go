// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// EconomyWaifuItem is an object representing the database table.
type EconomyWaifuItem struct {
	GuildID int64  `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	LocalID int64  `boil:"local_id" json:"local_id" toml:"local_id" yaml:"local_id"`
	Name    string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Icon    string `boil:"icon" json:"icon" toml:"icon" yaml:"icon"`
	Price   int    `boil:"price" json:"price" toml:"price" yaml:"price"`

	R *economyWaifuItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L economyWaifuItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EconomyWaifuItemColumns = struct {
	GuildID string
	LocalID string
	Name    string
	Icon    string
	Price   string
}{
	GuildID: "guild_id",
	LocalID: "local_id",
	Name:    "name",
	Icon:    "icon",
	Price:   "price",
}

// Generated where

var EconomyWaifuItemWhere = struct {
	GuildID whereHelperint64
	LocalID whereHelperint64
	Name    whereHelperstring
	Icon    whereHelperstring
	Price   whereHelperint
}{
	GuildID: whereHelperint64{field: `guild_id`},
	LocalID: whereHelperint64{field: `local_id`},
	Name:    whereHelperstring{field: `name`},
	Icon:    whereHelperstring{field: `icon`},
	Price:   whereHelperint{field: `price`},
}

// EconomyWaifuItemRels is where relationship names are stored.
var EconomyWaifuItemRels = struct {
}{}

// economyWaifuItemR is where relationships are stored.
type economyWaifuItemR struct {
}

// NewStruct creates a new relationship struct
func (*economyWaifuItemR) NewStruct() *economyWaifuItemR {
	return &economyWaifuItemR{}
}

// economyWaifuItemL is where Load methods for each relationship are stored.
type economyWaifuItemL struct{}

var (
	economyWaifuItemColumns               = []string{"guild_id", "local_id", "name", "icon", "price"}
	economyWaifuItemColumnsWithoutDefault = []string{"guild_id", "local_id", "name", "icon", "price"}
	economyWaifuItemColumnsWithDefault    = []string{}
	economyWaifuItemPrimaryKeyColumns     = []string{"guild_id", "local_id"}
)

type (
	// EconomyWaifuItemSlice is an alias for a slice of pointers to EconomyWaifuItem.
	// This should generally be used opposed to []EconomyWaifuItem.
	EconomyWaifuItemSlice []*EconomyWaifuItem

	economyWaifuItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	economyWaifuItemType                 = reflect.TypeOf(&EconomyWaifuItem{})
	economyWaifuItemMapping              = queries.MakeStructMapping(economyWaifuItemType)
	economyWaifuItemPrimaryKeyMapping, _ = queries.BindMapping(economyWaifuItemType, economyWaifuItemMapping, economyWaifuItemPrimaryKeyColumns)
	economyWaifuItemInsertCacheMut       sync.RWMutex
	economyWaifuItemInsertCache          = make(map[string]insertCache)
	economyWaifuItemUpdateCacheMut       sync.RWMutex
	economyWaifuItemUpdateCache          = make(map[string]updateCache)
	economyWaifuItemUpsertCacheMut       sync.RWMutex
	economyWaifuItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single economyWaifuItem record from the query using the global executor.
func (q economyWaifuItemQuery) OneG(ctx context.Context) (*EconomyWaifuItem, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single economyWaifuItem record from the query.
func (q economyWaifuItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EconomyWaifuItem, error) {
	o := &EconomyWaifuItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for economy_waifu_items")
	}

	return o, nil
}

// AllG returns all EconomyWaifuItem records from the query using the global executor.
func (q economyWaifuItemQuery) AllG(ctx context.Context) (EconomyWaifuItemSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all EconomyWaifuItem records from the query.
func (q economyWaifuItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (EconomyWaifuItemSlice, error) {
	var o []*EconomyWaifuItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EconomyWaifuItem slice")
	}

	return o, nil
}

// CountG returns the count of all EconomyWaifuItem records in the query, and panics on error.
func (q economyWaifuItemQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all EconomyWaifuItem records in the query.
func (q economyWaifuItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count economy_waifu_items rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q economyWaifuItemQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q economyWaifuItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if economy_waifu_items exists")
	}

	return count > 0, nil
}

// EconomyWaifuItems retrieves all the records using an executor.
func EconomyWaifuItems(mods ...qm.QueryMod) economyWaifuItemQuery {
	mods = append(mods, qm.From("\"economy_waifu_items\""))
	return economyWaifuItemQuery{NewQuery(mods...)}
}

// FindEconomyWaifuItemG retrieves a single record by ID.
func FindEconomyWaifuItemG(ctx context.Context, guildID int64, localID int64, selectCols ...string) (*EconomyWaifuItem, error) {
	return FindEconomyWaifuItem(ctx, boil.GetContextDB(), guildID, localID, selectCols...)
}

// FindEconomyWaifuItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEconomyWaifuItem(ctx context.Context, exec boil.ContextExecutor, guildID int64, localID int64, selectCols ...string) (*EconomyWaifuItem, error) {
	economyWaifuItemObj := &EconomyWaifuItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"economy_waifu_items\" where \"guild_id\"=$1 AND \"local_id\"=$2", sel,
	)

	q := queries.Raw(query, guildID, localID)

	err := q.Bind(ctx, exec, economyWaifuItemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from economy_waifu_items")
	}

	return economyWaifuItemObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *EconomyWaifuItem) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EconomyWaifuItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_waifu_items provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(economyWaifuItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	economyWaifuItemInsertCacheMut.RLock()
	cache, cached := economyWaifuItemInsertCache[key]
	economyWaifuItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			economyWaifuItemColumns,
			economyWaifuItemColumnsWithDefault,
			economyWaifuItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(economyWaifuItemType, economyWaifuItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(economyWaifuItemType, economyWaifuItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"economy_waifu_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"economy_waifu_items\" %sDEFAULT VALUES%s"
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
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into economy_waifu_items")
	}

	if !cached {
		economyWaifuItemInsertCacheMut.Lock()
		economyWaifuItemInsertCache[key] = cache
		economyWaifuItemInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single EconomyWaifuItem record using the global executor.
// See Update for more documentation.
func (o *EconomyWaifuItem) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the EconomyWaifuItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EconomyWaifuItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	economyWaifuItemUpdateCacheMut.RLock()
	cache, cached := economyWaifuItemUpdateCache[key]
	economyWaifuItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			economyWaifuItemColumns,
			economyWaifuItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update economy_waifu_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"economy_waifu_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, economyWaifuItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(economyWaifuItemType, economyWaifuItemMapping, append(wl, economyWaifuItemPrimaryKeyColumns...))
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
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update economy_waifu_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for economy_waifu_items")
	}

	if !cached {
		economyWaifuItemUpdateCacheMut.Lock()
		economyWaifuItemUpdateCache[key] = cache
		economyWaifuItemUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q economyWaifuItemQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q economyWaifuItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for economy_waifu_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for economy_waifu_items")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EconomyWaifuItemSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EconomyWaifuItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyWaifuItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"economy_waifu_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, economyWaifuItemPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in economyWaifuItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all economyWaifuItem")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *EconomyWaifuItem) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EconomyWaifuItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_waifu_items provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(economyWaifuItemColumnsWithDefault, o)

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

	economyWaifuItemUpsertCacheMut.RLock()
	cache, cached := economyWaifuItemUpsertCache[key]
	economyWaifuItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			economyWaifuItemColumns,
			economyWaifuItemColumnsWithDefault,
			economyWaifuItemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			economyWaifuItemColumns,
			economyWaifuItemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert economy_waifu_items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(economyWaifuItemPrimaryKeyColumns))
			copy(conflict, economyWaifuItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"economy_waifu_items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(economyWaifuItemType, economyWaifuItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(economyWaifuItemType, economyWaifuItemMapping, ret)
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
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert economy_waifu_items")
	}

	if !cached {
		economyWaifuItemUpsertCacheMut.Lock()
		economyWaifuItemUpsertCache[key] = cache
		economyWaifuItemUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single EconomyWaifuItem record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *EconomyWaifuItem) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single EconomyWaifuItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EconomyWaifuItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EconomyWaifuItem provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), economyWaifuItemPrimaryKeyMapping)
	sql := "DELETE FROM \"economy_waifu_items\" WHERE \"guild_id\"=$1 AND \"local_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from economy_waifu_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for economy_waifu_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q economyWaifuItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no economyWaifuItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economy_waifu_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_waifu_items")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o EconomyWaifuItemSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EconomyWaifuItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EconomyWaifuItem slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyWaifuItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"economy_waifu_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyWaifuItemPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economyWaifuItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_waifu_items")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *EconomyWaifuItem) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no EconomyWaifuItem provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EconomyWaifuItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEconomyWaifuItem(ctx, exec, o.GuildID, o.LocalID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyWaifuItemSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty EconomyWaifuItemSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyWaifuItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EconomyWaifuItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyWaifuItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"economy_waifu_items\".* FROM \"economy_waifu_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyWaifuItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EconomyWaifuItemSlice")
	}

	*o = slice

	return nil
}

// EconomyWaifuItemExistsG checks if the EconomyWaifuItem row exists.
func EconomyWaifuItemExistsG(ctx context.Context, guildID int64, localID int64) (bool, error) {
	return EconomyWaifuItemExists(ctx, boil.GetContextDB(), guildID, localID)
}

// EconomyWaifuItemExists checks if the EconomyWaifuItem row exists.
func EconomyWaifuItemExists(ctx context.Context, exec boil.ContextExecutor, guildID int64, localID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"economy_waifu_items\" where \"guild_id\"=$1 AND \"local_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID, localID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID, localID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if economy_waifu_items exists")
	}

	return exists, nil
}