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
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/projecttype"
	"github.com/facebookincubator/symphony/graph/ent/workorderdefinition"
	"github.com/facebookincubator/symphony/graph/ent/workordertype"
)

// WorkOrderDefinitionQuery is the builder for querying WorkOrderDefinition entities.
type WorkOrderDefinitionQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.WorkOrderDefinition
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (wodq *WorkOrderDefinitionQuery) Where(ps ...predicate.WorkOrderDefinition) *WorkOrderDefinitionQuery {
	wodq.predicates = append(wodq.predicates, ps...)
	return wodq
}

// Limit adds a limit step to the query.
func (wodq *WorkOrderDefinitionQuery) Limit(limit int) *WorkOrderDefinitionQuery {
	wodq.limit = &limit
	return wodq
}

// Offset adds an offset step to the query.
func (wodq *WorkOrderDefinitionQuery) Offset(offset int) *WorkOrderDefinitionQuery {
	wodq.offset = &offset
	return wodq
}

// Order adds an order step to the query.
func (wodq *WorkOrderDefinitionQuery) Order(o ...Order) *WorkOrderDefinitionQuery {
	wodq.order = append(wodq.order, o...)
	return wodq
}

// QueryType chains the current query on the type edge.
func (wodq *WorkOrderDefinitionQuery) QueryType() *WorkOrderTypeQuery {
	query := &WorkOrderTypeQuery{config: wodq.config}
	step := &sql.Step{}
	step.From.V = wodq.sqlQuery()
	step.From.Table = workorderdefinition.Table
	step.From.Column = workorderdefinition.FieldID
	step.To.Table = workordertype.Table
	step.To.Column = workordertype.FieldID
	step.Edge.Rel = sql.M2O
	step.Edge.Inverse = false
	step.Edge.Table = workorderdefinition.TypeTable
	step.Edge.Columns = append(step.Edge.Columns, workorderdefinition.TypeColumn)
	query.sql = sql.SetNeighbors(wodq.driver.Dialect(), step)
	return query
}

// QueryProjectType chains the current query on the project_type edge.
func (wodq *WorkOrderDefinitionQuery) QueryProjectType() *ProjectTypeQuery {
	query := &ProjectTypeQuery{config: wodq.config}
	step := &sql.Step{}
	step.From.V = wodq.sqlQuery()
	step.From.Table = workorderdefinition.Table
	step.From.Column = workorderdefinition.FieldID
	step.To.Table = projecttype.Table
	step.To.Column = projecttype.FieldID
	step.Edge.Rel = sql.M2O
	step.Edge.Inverse = true
	step.Edge.Table = workorderdefinition.ProjectTypeTable
	step.Edge.Columns = append(step.Edge.Columns, workorderdefinition.ProjectTypeColumn)
	query.sql = sql.SetNeighbors(wodq.driver.Dialect(), step)
	return query
}

// First returns the first WorkOrderDefinition entity in the query. Returns *ErrNotFound when no workorderdefinition was found.
func (wodq *WorkOrderDefinitionQuery) First(ctx context.Context) (*WorkOrderDefinition, error) {
	wods, err := wodq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(wods) == 0 {
		return nil, &ErrNotFound{workorderdefinition.Label}
	}
	return wods[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wodq *WorkOrderDefinitionQuery) FirstX(ctx context.Context) *WorkOrderDefinition {
	wod, err := wodq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return wod
}

// FirstID returns the first WorkOrderDefinition id in the query. Returns *ErrNotFound when no id was found.
func (wodq *WorkOrderDefinitionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wodq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{workorderdefinition.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (wodq *WorkOrderDefinitionQuery) FirstXID(ctx context.Context) string {
	id, err := wodq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only WorkOrderDefinition entity in the query, returns an error if not exactly one entity was returned.
func (wodq *WorkOrderDefinitionQuery) Only(ctx context.Context) (*WorkOrderDefinition, error) {
	wods, err := wodq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(wods) {
	case 1:
		return wods[0], nil
	case 0:
		return nil, &ErrNotFound{workorderdefinition.Label}
	default:
		return nil, &ErrNotSingular{workorderdefinition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wodq *WorkOrderDefinitionQuery) OnlyX(ctx context.Context) *WorkOrderDefinition {
	wod, err := wodq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return wod
}

// OnlyID returns the only WorkOrderDefinition id in the query, returns an error if not exactly one id was returned.
func (wodq *WorkOrderDefinitionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wodq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{workorderdefinition.Label}
	default:
		err = &ErrNotSingular{workorderdefinition.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (wodq *WorkOrderDefinitionQuery) OnlyXID(ctx context.Context) string {
	id, err := wodq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkOrderDefinitions.
func (wodq *WorkOrderDefinitionQuery) All(ctx context.Context) ([]*WorkOrderDefinition, error) {
	return wodq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wodq *WorkOrderDefinitionQuery) AllX(ctx context.Context) []*WorkOrderDefinition {
	wods, err := wodq.All(ctx)
	if err != nil {
		panic(err)
	}
	return wods
}

// IDs executes the query and returns a list of WorkOrderDefinition ids.
func (wodq *WorkOrderDefinitionQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := wodq.Select(workorderdefinition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wodq *WorkOrderDefinitionQuery) IDsX(ctx context.Context) []string {
	ids, err := wodq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wodq *WorkOrderDefinitionQuery) Count(ctx context.Context) (int, error) {
	return wodq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wodq *WorkOrderDefinitionQuery) CountX(ctx context.Context) int {
	count, err := wodq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wodq *WorkOrderDefinitionQuery) Exist(ctx context.Context) (bool, error) {
	return wodq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wodq *WorkOrderDefinitionQuery) ExistX(ctx context.Context) bool {
	exist, err := wodq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wodq *WorkOrderDefinitionQuery) Clone() *WorkOrderDefinitionQuery {
	return &WorkOrderDefinitionQuery{
		config:     wodq.config,
		limit:      wodq.limit,
		offset:     wodq.offset,
		order:      append([]Order{}, wodq.order...),
		unique:     append([]string{}, wodq.unique...),
		predicates: append([]predicate.WorkOrderDefinition{}, wodq.predicates...),
		// clone intermediate queries.
		sql: wodq.sql.Clone(),
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
//	client.WorkOrderDefinition.Query().
//		GroupBy(workorderdefinition.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wodq *WorkOrderDefinitionQuery) GroupBy(field string, fields ...string) *WorkOrderDefinitionGroupBy {
	group := &WorkOrderDefinitionGroupBy{config: wodq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = wodq.sqlQuery()
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
//	client.WorkOrderDefinition.Query().
//		Select(workorderdefinition.FieldCreateTime).
//		Scan(ctx, &v)
//
func (wodq *WorkOrderDefinitionQuery) Select(field string, fields ...string) *WorkOrderDefinitionSelect {
	selector := &WorkOrderDefinitionSelect{config: wodq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = wodq.sqlQuery()
	return selector
}

func (wodq *WorkOrderDefinitionQuery) sqlAll(ctx context.Context) ([]*WorkOrderDefinition, error) {
	rows := &sql.Rows{}
	selector := wodq.sqlQuery()
	if unique := wodq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := wodq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var wods WorkOrderDefinitions
	if err := wods.FromRows(rows); err != nil {
		return nil, err
	}
	wods.config(wodq.config)
	return wods, nil
}

func (wodq *WorkOrderDefinitionQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := wodq.sqlQuery()
	unique := []string{workorderdefinition.FieldID}
	if len(wodq.unique) > 0 {
		unique = wodq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := wodq.driver.Query(ctx, query, args, rows); err != nil {
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

func (wodq *WorkOrderDefinitionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wodq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (wodq *WorkOrderDefinitionQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wodq.driver.Dialect())
	t1 := builder.Table(workorderdefinition.Table)
	selector := builder.Select(t1.Columns(workorderdefinition.Columns...)...).From(t1)
	if wodq.sql != nil {
		selector = wodq.sql
		selector.Select(selector.Columns(workorderdefinition.Columns...)...)
	}
	for _, p := range wodq.predicates {
		p(selector)
	}
	for _, p := range wodq.order {
		p(selector)
	}
	if offset := wodq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wodq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkOrderDefinitionGroupBy is the builder for group-by WorkOrderDefinition entities.
type WorkOrderDefinitionGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wodgb *WorkOrderDefinitionGroupBy) Aggregate(fns ...Aggregate) *WorkOrderDefinitionGroupBy {
	wodgb.fns = append(wodgb.fns, fns...)
	return wodgb
}

// Scan applies the group-by query and scan the result into the given value.
func (wodgb *WorkOrderDefinitionGroupBy) Scan(ctx context.Context, v interface{}) error {
	return wodgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wodgb *WorkOrderDefinitionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wodgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wodgb *WorkOrderDefinitionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wodgb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderDefinitionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wodgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wodgb *WorkOrderDefinitionGroupBy) StringsX(ctx context.Context) []string {
	v, err := wodgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wodgb *WorkOrderDefinitionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wodgb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderDefinitionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wodgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wodgb *WorkOrderDefinitionGroupBy) IntsX(ctx context.Context) []int {
	v, err := wodgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wodgb *WorkOrderDefinitionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wodgb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderDefinitionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wodgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wodgb *WorkOrderDefinitionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wodgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wodgb *WorkOrderDefinitionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wodgb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderDefinitionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wodgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wodgb *WorkOrderDefinitionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wodgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wodgb *WorkOrderDefinitionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wodgb.sqlQuery().Query()
	if err := wodgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wodgb *WorkOrderDefinitionGroupBy) sqlQuery() *sql.Selector {
	selector := wodgb.sql
	columns := make([]string, 0, len(wodgb.fields)+len(wodgb.fns))
	columns = append(columns, wodgb.fields...)
	for _, fn := range wodgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(wodgb.fields...)
}

// WorkOrderDefinitionSelect is the builder for select fields of WorkOrderDefinition entities.
type WorkOrderDefinitionSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (wods *WorkOrderDefinitionSelect) Scan(ctx context.Context, v interface{}) error {
	return wods.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wods *WorkOrderDefinitionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wods.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (wods *WorkOrderDefinitionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wods.fields) > 1 {
		return nil, errors.New("ent: WorkOrderDefinitionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wods *WorkOrderDefinitionSelect) StringsX(ctx context.Context) []string {
	v, err := wods.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (wods *WorkOrderDefinitionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wods.fields) > 1 {
		return nil, errors.New("ent: WorkOrderDefinitionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wods *WorkOrderDefinitionSelect) IntsX(ctx context.Context) []int {
	v, err := wods.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (wods *WorkOrderDefinitionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wods.fields) > 1 {
		return nil, errors.New("ent: WorkOrderDefinitionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wods *WorkOrderDefinitionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wods.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (wods *WorkOrderDefinitionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wods.fields) > 1 {
		return nil, errors.New("ent: WorkOrderDefinitionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wods *WorkOrderDefinitionSelect) BoolsX(ctx context.Context) []bool {
	v, err := wods.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wods *WorkOrderDefinitionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wods.sqlQuery().Query()
	if err := wods.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wods *WorkOrderDefinitionSelect) sqlQuery() sql.Querier {
	view := "workorderdefinition_view"
	return sql.Dialect(wods.driver.Dialect()).
		Select(wods.fields...).From(wods.sql.As(view))
}
