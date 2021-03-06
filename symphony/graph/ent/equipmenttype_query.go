// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/equipment"
	"github.com/facebookincubator/symphony/graph/ent/equipmentcategory"
	"github.com/facebookincubator/symphony/graph/ent/equipmentportdefinition"
	"github.com/facebookincubator/symphony/graph/ent/equipmentpositiondefinition"
	"github.com/facebookincubator/symphony/graph/ent/equipmenttype"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/propertytype"
)

// EquipmentTypeQuery is the builder for querying EquipmentType entities.
type EquipmentTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.EquipmentType
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (etq *EquipmentTypeQuery) Where(ps ...predicate.EquipmentType) *EquipmentTypeQuery {
	etq.predicates = append(etq.predicates, ps...)
	return etq
}

// Limit adds a limit step to the query.
func (etq *EquipmentTypeQuery) Limit(limit int) *EquipmentTypeQuery {
	etq.limit = &limit
	return etq
}

// Offset adds an offset step to the query.
func (etq *EquipmentTypeQuery) Offset(offset int) *EquipmentTypeQuery {
	etq.offset = &offset
	return etq
}

// Order adds an order step to the query.
func (etq *EquipmentTypeQuery) Order(o ...Order) *EquipmentTypeQuery {
	etq.order = append(etq.order, o...)
	return etq
}

// QueryPortDefinitions chains the current query on the port_definitions edge.
func (etq *EquipmentTypeQuery) QueryPortDefinitions() *EquipmentPortDefinitionQuery {
	query := &EquipmentPortDefinitionQuery{config: etq.config}
	step := &sql.Step{}
	step.From.V = etq.sqlQuery()
	step.From.Table = equipmenttype.Table
	step.From.Column = equipmenttype.FieldID
	step.To.Table = equipmentportdefinition.Table
	step.To.Column = equipmentportdefinition.FieldID
	step.Edge.Rel = sql.O2M
	step.Edge.Inverse = false
	step.Edge.Table = equipmenttype.PortDefinitionsTable
	step.Edge.Columns = append(step.Edge.Columns, equipmenttype.PortDefinitionsColumn)
	query.sql = sql.SetNeighbors(etq.driver.Dialect(), step)
	return query
}

// QueryPositionDefinitions chains the current query on the position_definitions edge.
func (etq *EquipmentTypeQuery) QueryPositionDefinitions() *EquipmentPositionDefinitionQuery {
	query := &EquipmentPositionDefinitionQuery{config: etq.config}
	step := &sql.Step{}
	step.From.V = etq.sqlQuery()
	step.From.Table = equipmenttype.Table
	step.From.Column = equipmenttype.FieldID
	step.To.Table = equipmentpositiondefinition.Table
	step.To.Column = equipmentpositiondefinition.FieldID
	step.Edge.Rel = sql.O2M
	step.Edge.Inverse = false
	step.Edge.Table = equipmenttype.PositionDefinitionsTable
	step.Edge.Columns = append(step.Edge.Columns, equipmenttype.PositionDefinitionsColumn)
	query.sql = sql.SetNeighbors(etq.driver.Dialect(), step)
	return query
}

// QueryPropertyTypes chains the current query on the property_types edge.
func (etq *EquipmentTypeQuery) QueryPropertyTypes() *PropertyTypeQuery {
	query := &PropertyTypeQuery{config: etq.config}
	step := &sql.Step{}
	step.From.V = etq.sqlQuery()
	step.From.Table = equipmenttype.Table
	step.From.Column = equipmenttype.FieldID
	step.To.Table = propertytype.Table
	step.To.Column = propertytype.FieldID
	step.Edge.Rel = sql.O2M
	step.Edge.Inverse = false
	step.Edge.Table = equipmenttype.PropertyTypesTable
	step.Edge.Columns = append(step.Edge.Columns, equipmenttype.PropertyTypesColumn)
	query.sql = sql.SetNeighbors(etq.driver.Dialect(), step)
	return query
}

// QueryEquipment chains the current query on the equipment edge.
func (etq *EquipmentTypeQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: etq.config}
	step := &sql.Step{}
	step.From.V = etq.sqlQuery()
	step.From.Table = equipmenttype.Table
	step.From.Column = equipmenttype.FieldID
	step.To.Table = equipment.Table
	step.To.Column = equipment.FieldID
	step.Edge.Rel = sql.O2M
	step.Edge.Inverse = true
	step.Edge.Table = equipmenttype.EquipmentTable
	step.Edge.Columns = append(step.Edge.Columns, equipmenttype.EquipmentColumn)
	query.sql = sql.SetNeighbors(etq.driver.Dialect(), step)
	return query
}

// QueryCategory chains the current query on the category edge.
func (etq *EquipmentTypeQuery) QueryCategory() *EquipmentCategoryQuery {
	query := &EquipmentCategoryQuery{config: etq.config}
	step := &sql.Step{}
	step.From.V = etq.sqlQuery()
	step.From.Table = equipmenttype.Table
	step.From.Column = equipmenttype.FieldID
	step.To.Table = equipmentcategory.Table
	step.To.Column = equipmentcategory.FieldID
	step.Edge.Rel = sql.M2O
	step.Edge.Inverse = false
	step.Edge.Table = equipmenttype.CategoryTable
	step.Edge.Columns = append(step.Edge.Columns, equipmenttype.CategoryColumn)
	query.sql = sql.SetNeighbors(etq.driver.Dialect(), step)
	return query
}

// First returns the first EquipmentType entity in the query. Returns *ErrNotFound when no equipmenttype was found.
func (etq *EquipmentTypeQuery) First(ctx context.Context) (*EquipmentType, error) {
	ets, err := etq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ets) == 0 {
		return nil, &ErrNotFound{equipmenttype.Label}
	}
	return ets[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (etq *EquipmentTypeQuery) FirstX(ctx context.Context) *EquipmentType {
	et, err := etq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return et
}

// FirstID returns the first EquipmentType id in the query. Returns *ErrNotFound when no id was found.
func (etq *EquipmentTypeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = etq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{equipmenttype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (etq *EquipmentTypeQuery) FirstXID(ctx context.Context) string {
	id, err := etq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only EquipmentType entity in the query, returns an error if not exactly one entity was returned.
func (etq *EquipmentTypeQuery) Only(ctx context.Context) (*EquipmentType, error) {
	ets, err := etq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ets) {
	case 1:
		return ets[0], nil
	case 0:
		return nil, &ErrNotFound{equipmenttype.Label}
	default:
		return nil, &ErrNotSingular{equipmenttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (etq *EquipmentTypeQuery) OnlyX(ctx context.Context) *EquipmentType {
	et, err := etq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return et
}

// OnlyID returns the only EquipmentType id in the query, returns an error if not exactly one id was returned.
func (etq *EquipmentTypeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = etq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{equipmenttype.Label}
	default:
		err = &ErrNotSingular{equipmenttype.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (etq *EquipmentTypeQuery) OnlyXID(ctx context.Context) string {
	id, err := etq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentTypes.
func (etq *EquipmentTypeQuery) All(ctx context.Context) ([]*EquipmentType, error) {
	return etq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (etq *EquipmentTypeQuery) AllX(ctx context.Context) []*EquipmentType {
	ets, err := etq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ets
}

// IDs executes the query and returns a list of EquipmentType ids.
func (etq *EquipmentTypeQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := etq.Select(equipmenttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (etq *EquipmentTypeQuery) IDsX(ctx context.Context) []string {
	ids, err := etq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (etq *EquipmentTypeQuery) Count(ctx context.Context) (int, error) {
	return etq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (etq *EquipmentTypeQuery) CountX(ctx context.Context) int {
	count, err := etq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (etq *EquipmentTypeQuery) Exist(ctx context.Context) (bool, error) {
	return etq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (etq *EquipmentTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := etq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (etq *EquipmentTypeQuery) Clone() *EquipmentTypeQuery {
	return &EquipmentTypeQuery{
		config:     etq.config,
		limit:      etq.limit,
		offset:     etq.offset,
		order:      append([]Order{}, etq.order...),
		unique:     append([]string{}, etq.unique...),
		predicates: append([]predicate.EquipmentType{}, etq.predicates...),
		// clone intermediate queries.
		sql: etq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EquipmentType.Query().
//		GroupBy(equipmenttype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (etq *EquipmentTypeQuery) GroupBy(field string, fields ...string) *EquipmentTypeGroupBy {
	group := &EquipmentTypeGroupBy{config: etq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = etq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.EquipmentType.Query().
//		Select(equipmenttype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (etq *EquipmentTypeQuery) Select(field string, fields ...string) *EquipmentTypeSelect {
	selector := &EquipmentTypeSelect{config: etq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = etq.sqlQuery()
	return selector
}

func (etq *EquipmentTypeQuery) sqlAll(ctx context.Context) ([]*EquipmentType, error) {
	rows := &sql.Rows{}
	selector := etq.sqlQuery()
	if unique := etq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := etq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ets EquipmentTypes
	if err := ets.FromRows(rows); err != nil {
		return nil, err
	}
	ets.config(etq.config)
	return ets, nil
}

func (etq *EquipmentTypeQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := etq.sqlQuery()
	unique := []string{equipmenttype.FieldID}
	if len(etq.unique) > 0 {
		unique = etq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := etq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (etq *EquipmentTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := etq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (etq *EquipmentTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(etq.driver.Dialect())
	t1 := builder.Table(equipmenttype.Table)
	selector := builder.Select(t1.Columns(equipmenttype.Columns...)...).From(t1)
	if etq.sql != nil {
		selector = etq.sql
		selector.Select(selector.Columns(equipmenttype.Columns...)...)
	}
	for _, p := range etq.predicates {
		p(selector)
	}
	for _, p := range etq.order {
		p(selector)
	}
	if offset := etq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := etq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentTypeGroupBy is the builder for group-by EquipmentType entities.
type EquipmentTypeGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (etgb *EquipmentTypeGroupBy) Aggregate(fns ...Aggregate) *EquipmentTypeGroupBy {
	etgb.fns = append(etgb.fns, fns...)
	return etgb
}

// Scan applies the group-by query and scan the result into the given value.
func (etgb *EquipmentTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	return etgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (etgb *EquipmentTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := etgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (etgb *EquipmentTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (etgb *EquipmentTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := etgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (etgb *EquipmentTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (etgb *EquipmentTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := etgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (etgb *EquipmentTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (etgb *EquipmentTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := etgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (etgb *EquipmentTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EquipmentTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (etgb *EquipmentTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := etgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (etgb *EquipmentTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := etgb.sqlQuery().Query()
	if err := etgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (etgb *EquipmentTypeGroupBy) sqlQuery() *sql.Selector {
	selector := etgb.sql
	columns := make([]string, 0, len(etgb.fields)+len(etgb.fns))
	columns = append(columns, etgb.fields...)
	for _, fn := range etgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(etgb.fields...)
}

// EquipmentTypeSelect is the builder for select fields of EquipmentType entities.
type EquipmentTypeSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ets *EquipmentTypeSelect) Scan(ctx context.Context, v interface{}) error {
	return ets.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ets *EquipmentTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ets.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ets *EquipmentTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EquipmentTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ets *EquipmentTypeSelect) StringsX(ctx context.Context) []string {
	v, err := ets.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ets *EquipmentTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EquipmentTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ets *EquipmentTypeSelect) IntsX(ctx context.Context) []int {
	v, err := ets.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ets *EquipmentTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EquipmentTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ets *EquipmentTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ets.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ets *EquipmentTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EquipmentTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ets *EquipmentTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := ets.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ets *EquipmentTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ets.sqlQuery().Query()
	if err := ets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ets *EquipmentTypeSelect) sqlQuery() sql.Querier {
	view := "equipmenttype_view"
	return sql.Dialect(ets.driver.Dialect()).
		Select(ets.fields...).From(ets.sql.As(view))
}
