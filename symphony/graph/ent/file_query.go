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
	"github.com/facebookincubator/symphony/graph/ent/file"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// FileQuery is the builder for querying File entities.
type FileQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.File
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (fq *FileQuery) Where(ps ...predicate.File) *FileQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FileQuery) Limit(limit int) *FileQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FileQuery) Offset(offset int) *FileQuery {
	fq.offset = &offset
	return fq
}

// Order adds an order step to the query.
func (fq *FileQuery) Order(o ...Order) *FileQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// First returns the first File entity in the query. Returns *ErrNotFound when no file was found.
func (fq *FileQuery) First(ctx context.Context) (*File, error) {
	fs, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(fs) == 0 {
		return nil, &ErrNotFound{file.Label}
	}
	return fs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FileQuery) FirstX(ctx context.Context) *File {
	f, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return f
}

// FirstID returns the first File id in the query. Returns *ErrNotFound when no id was found.
func (fq *FileQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{file.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (fq *FileQuery) FirstXID(ctx context.Context) string {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only File entity in the query, returns an error if not exactly one entity was returned.
func (fq *FileQuery) Only(ctx context.Context) (*File, error) {
	fs, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(fs) {
	case 1:
		return fs[0], nil
	case 0:
		return nil, &ErrNotFound{file.Label}
	default:
		return nil, &ErrNotSingular{file.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FileQuery) OnlyX(ctx context.Context) *File {
	f, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// OnlyID returns the only File id in the query, returns an error if not exactly one id was returned.
func (fq *FileQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{file.Label}
	default:
		err = &ErrNotSingular{file.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (fq *FileQuery) OnlyXID(ctx context.Context) string {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Files.
func (fq *FileQuery) All(ctx context.Context) ([]*File, error) {
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FileQuery) AllX(ctx context.Context) []*File {
	fs, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return fs
}

// IDs executes the query and returns a list of File ids.
func (fq *FileQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := fq.Select(file.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FileQuery) IDsX(ctx context.Context) []string {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FileQuery) Count(ctx context.Context) (int, error) {
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FileQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FileQuery) Exist(ctx context.Context) (bool, error) {
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FileQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FileQuery) Clone() *FileQuery {
	return &FileQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]Order{}, fq.order...),
		unique:     append([]string{}, fq.unique...),
		predicates: append([]predicate.File{}, fq.predicates...),
		// clone intermediate queries.
		sql: fq.sql.Clone(),
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
//	client.File.Query().
//		GroupBy(file.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fq *FileQuery) GroupBy(field string, fields ...string) *FileGroupBy {
	group := &FileGroupBy{config: fq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = fq.sqlQuery()
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
//	client.File.Query().
//		Select(file.FieldCreateTime).
//		Scan(ctx, &v)
//
func (fq *FileQuery) Select(field string, fields ...string) *FileSelect {
	selector := &FileSelect{config: fq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = fq.sqlQuery()
	return selector
}

func (fq *FileQuery) sqlAll(ctx context.Context) ([]*File, error) {
	rows := &sql.Rows{}
	selector := fq.sqlQuery()
	if unique := fq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := fq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var fs Files
	if err := fs.FromRows(rows); err != nil {
		return nil, err
	}
	fs.config(fq.config)
	return fs, nil
}

func (fq *FileQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := fq.sqlQuery()
	unique := []string{file.FieldID}
	if len(fq.unique) > 0 {
		unique = fq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := fq.driver.Query(ctx, query, args, rows); err != nil {
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

func (fq *FileQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (fq *FileQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(file.Table)
	selector := builder.Select(t1.Columns(file.Columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(file.Columns...)...)
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FileGroupBy is the builder for group-by File entities.
type FileGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FileGroupBy) Aggregate(fns ...Aggregate) *FileGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scan the result into the given value.
func (fgb *FileGroupBy) Scan(ctx context.Context, v interface{}) error {
	return fgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fgb *FileGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (fgb *FileGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FileGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgb *FileGroupBy) StringsX(ctx context.Context) []string {
	v, err := fgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (fgb *FileGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FileGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgb *FileGroupBy) IntsX(ctx context.Context) []int {
	v, err := fgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (fgb *FileGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FileGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgb *FileGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (fgb *FileGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FileGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgb *FileGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgb *FileGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fgb.sqlQuery().Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FileGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql
	columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
	columns = append(columns, fgb.fields...)
	for _, fn := range fgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(fgb.fields...)
}

// FileSelect is the builder for select fields of File entities.
type FileSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (fs *FileSelect) Scan(ctx context.Context, v interface{}) error {
	return fs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fs *FileSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (fs *FileSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FileSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fs *FileSelect) StringsX(ctx context.Context) []string {
	v, err := fs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (fs *FileSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FileSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fs *FileSelect) IntsX(ctx context.Context) []int {
	v, err := fs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (fs *FileSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FileSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fs *FileSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (fs *FileSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FileSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fs *FileSelect) BoolsX(ctx context.Context) []bool {
	v, err := fs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fs *FileSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fs.sqlQuery().Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fs *FileSelect) sqlQuery() sql.Querier {
	view := "file_view"
	return sql.Dialect(fs.driver.Dialect()).
		Select(fs.fields...).From(fs.sql.As(view))
}
