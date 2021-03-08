// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/mediaevents"
	"github.com/dreamvo/gilfoyle/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// MediaEventsQuery is the builder for querying MediaEvents entities.
type MediaEventsQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.MediaEvents
	// eager-loading edges.
	withMedia *MediaQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (meq *MediaEventsQuery) Where(ps ...predicate.MediaEvents) *MediaEventsQuery {
	meq.predicates = append(meq.predicates, ps...)
	return meq
}

// Limit adds a limit step to the query.
func (meq *MediaEventsQuery) Limit(limit int) *MediaEventsQuery {
	meq.limit = &limit
	return meq
}

// Offset adds an offset step to the query.
func (meq *MediaEventsQuery) Offset(offset int) *MediaEventsQuery {
	meq.offset = &offset
	return meq
}

// Order adds an order step to the query.
func (meq *MediaEventsQuery) Order(o ...OrderFunc) *MediaEventsQuery {
	meq.order = append(meq.order, o...)
	return meq
}

// QueryMedia chains the current query on the media edge.
func (meq *MediaEventsQuery) QueryMedia() *MediaQuery {
	query := &MediaQuery{config: meq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := meq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := meq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mediaevents.Table, mediaevents.FieldID, selector),
			sqlgraph.To(media.Table, media.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, mediaevents.MediaTable, mediaevents.MediaColumn),
		)
		fromU = sqlgraph.SetNeighbors(meq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MediaEvents entity in the query. Returns *NotFoundError when no mediaevents was found.
func (meq *MediaEventsQuery) First(ctx context.Context) (*MediaEvents, error) {
	nodes, err := meq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mediaevents.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (meq *MediaEventsQuery) FirstX(ctx context.Context) *MediaEvents {
	node, err := meq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MediaEvents id in the query. Returns *NotFoundError when no id was found.
func (meq *MediaEventsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = meq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mediaevents.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (meq *MediaEventsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := meq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only MediaEvents entity in the query, returns an error if not exactly one entity was returned.
func (meq *MediaEventsQuery) Only(ctx context.Context) (*MediaEvents, error) {
	nodes, err := meq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mediaevents.Label}
	default:
		return nil, &NotSingularError{mediaevents.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (meq *MediaEventsQuery) OnlyX(ctx context.Context) *MediaEvents {
	node, err := meq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only MediaEvents id in the query, returns an error if not exactly one id was returned.
func (meq *MediaEventsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = meq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = &NotSingularError{mediaevents.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (meq *MediaEventsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := meq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MediaEventsSlice.
func (meq *MediaEventsQuery) All(ctx context.Context) ([]*MediaEvents, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return meq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (meq *MediaEventsQuery) AllX(ctx context.Context) []*MediaEvents {
	nodes, err := meq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MediaEvents ids.
func (meq *MediaEventsQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := meq.Select(mediaevents.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (meq *MediaEventsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := meq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (meq *MediaEventsQuery) Count(ctx context.Context) (int, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return meq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (meq *MediaEventsQuery) CountX(ctx context.Context) int {
	count, err := meq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (meq *MediaEventsQuery) Exist(ctx context.Context) (bool, error) {
	if err := meq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return meq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (meq *MediaEventsQuery) ExistX(ctx context.Context) bool {
	exist, err := meq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (meq *MediaEventsQuery) Clone() *MediaEventsQuery {
	if meq == nil {
		return nil
	}
	return &MediaEventsQuery{
		config:     meq.config,
		limit:      meq.limit,
		offset:     meq.offset,
		order:      append([]OrderFunc{}, meq.order...),
		predicates: append([]predicate.MediaEvents{}, meq.predicates...),
		withMedia:  meq.withMedia.Clone(),
		// clone intermediate query.
		sql:  meq.sql.Clone(),
		path: meq.path,
	}
}

//  WithMedia tells the query-builder to eager-loads the nodes that are connected to
// the "media" edge. The optional arguments used to configure the query builder of the edge.
func (meq *MediaEventsQuery) WithMedia(opts ...func(*MediaQuery)) *MediaEventsQuery {
	query := &MediaQuery{config: meq.config}
	for _, opt := range opts {
		opt(query)
	}
	meq.withMedia = query
	return meq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Level mediaevents.Level `json:"level,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MediaEvents.Query().
//		GroupBy(mediaevents.FieldLevel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (meq *MediaEventsQuery) GroupBy(field string, fields ...string) *MediaEventsGroupBy {
	group := &MediaEventsGroupBy{config: meq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := meq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return meq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Level mediaevents.Level `json:"level,omitempty"`
//	}
//
//	client.MediaEvents.Query().
//		Select(mediaevents.FieldLevel).
//		Scan(ctx, &v)
//
func (meq *MediaEventsQuery) Select(field string, fields ...string) *MediaEventsSelect {
	selector := &MediaEventsSelect{config: meq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := meq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return meq.sqlQuery(), nil
	}
	return selector
}

func (meq *MediaEventsQuery) prepareQuery(ctx context.Context) error {
	if meq.path != nil {
		prev, err := meq.path(ctx)
		if err != nil {
			return err
		}
		meq.sql = prev
	}
	return nil
}

func (meq *MediaEventsQuery) sqlAll(ctx context.Context) ([]*MediaEvents, error) {
	var (
		nodes       = []*MediaEvents{}
		withFKs     = meq.withFKs
		_spec       = meq.querySpec()
		loadedTypes = [1]bool{
			meq.withMedia != nil,
		}
	)
	if meq.withMedia != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, mediaevents.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &MediaEvents{config: meq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, meq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := meq.withMedia; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*MediaEvents)
		for i := range nodes {
			if fk := nodes[i].media; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(media.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "media" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Media = n
			}
		}
	}

	return nodes, nil
}

func (meq *MediaEventsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := meq.querySpec()
	return sqlgraph.CountNodes(ctx, meq.driver, _spec)
}

func (meq *MediaEventsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := meq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (meq *MediaEventsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mediaevents.Table,
			Columns: mediaevents.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mediaevents.FieldID,
			},
		},
		From:   meq.sql,
		Unique: true,
	}
	if ps := meq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := meq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := meq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := meq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, mediaevents.ValidColumn)
			}
		}
	}
	return _spec
}

func (meq *MediaEventsQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(meq.driver.Dialect())
	t1 := builder.Table(mediaevents.Table)
	selector := builder.Select(t1.Columns(mediaevents.Columns...)...).From(t1)
	if meq.sql != nil {
		selector = meq.sql
		selector.Select(selector.Columns(mediaevents.Columns...)...)
	}
	for _, p := range meq.predicates {
		p(selector)
	}
	for _, p := range meq.order {
		p(selector, mediaevents.ValidColumn)
	}
	if offset := meq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := meq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MediaEventsGroupBy is the builder for group-by MediaEvents entities.
type MediaEventsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (megb *MediaEventsGroupBy) Aggregate(fns ...AggregateFunc) *MediaEventsGroupBy {
	megb.fns = append(megb.fns, fns...)
	return megb
}

// Scan applies the group-by query and scan the result into the given value.
func (megb *MediaEventsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := megb.path(ctx)
	if err != nil {
		return err
	}
	megb.sql = query
	return megb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (megb *MediaEventsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := megb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (megb *MediaEventsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(megb.fields) > 1 {
		return nil, errors.New("ent: MediaEventsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := megb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (megb *MediaEventsGroupBy) StringsX(ctx context.Context) []string {
	v, err := megb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (megb *MediaEventsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = megb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = fmt.Errorf("ent: MediaEventsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (megb *MediaEventsGroupBy) StringX(ctx context.Context) string {
	v, err := megb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (megb *MediaEventsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(megb.fields) > 1 {
		return nil, errors.New("ent: MediaEventsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := megb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (megb *MediaEventsGroupBy) IntsX(ctx context.Context) []int {
	v, err := megb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (megb *MediaEventsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = megb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = fmt.Errorf("ent: MediaEventsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (megb *MediaEventsGroupBy) IntX(ctx context.Context) int {
	v, err := megb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (megb *MediaEventsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(megb.fields) > 1 {
		return nil, errors.New("ent: MediaEventsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := megb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (megb *MediaEventsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := megb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (megb *MediaEventsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = megb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = fmt.Errorf("ent: MediaEventsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (megb *MediaEventsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := megb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (megb *MediaEventsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(megb.fields) > 1 {
		return nil, errors.New("ent: MediaEventsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := megb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (megb *MediaEventsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := megb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (megb *MediaEventsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = megb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = fmt.Errorf("ent: MediaEventsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (megb *MediaEventsGroupBy) BoolX(ctx context.Context) bool {
	v, err := megb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (megb *MediaEventsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range megb.fields {
		if !mediaevents.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := megb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := megb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (megb *MediaEventsGroupBy) sqlQuery() *sql.Selector {
	selector := megb.sql
	columns := make([]string, 0, len(megb.fields)+len(megb.fns))
	columns = append(columns, megb.fields...)
	for _, fn := range megb.fns {
		columns = append(columns, fn(selector, mediaevents.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(megb.fields...)
}

// MediaEventsSelect is the builder for select fields of MediaEvents entities.
type MediaEventsSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (mes *MediaEventsSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := mes.path(ctx)
	if err != nil {
		return err
	}
	mes.sql = query
	return mes.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mes *MediaEventsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mes.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (mes *MediaEventsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mes.fields) > 1 {
		return nil, errors.New("ent: MediaEventsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mes *MediaEventsSelect) StringsX(ctx context.Context) []string {
	v, err := mes.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (mes *MediaEventsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mes.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = fmt.Errorf("ent: MediaEventsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mes *MediaEventsSelect) StringX(ctx context.Context) string {
	v, err := mes.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (mes *MediaEventsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mes.fields) > 1 {
		return nil, errors.New("ent: MediaEventsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mes *MediaEventsSelect) IntsX(ctx context.Context) []int {
	v, err := mes.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (mes *MediaEventsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mes.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = fmt.Errorf("ent: MediaEventsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mes *MediaEventsSelect) IntX(ctx context.Context) int {
	v, err := mes.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (mes *MediaEventsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mes.fields) > 1 {
		return nil, errors.New("ent: MediaEventsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mes *MediaEventsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mes.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (mes *MediaEventsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mes.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = fmt.Errorf("ent: MediaEventsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mes *MediaEventsSelect) Float64X(ctx context.Context) float64 {
	v, err := mes.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (mes *MediaEventsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mes.fields) > 1 {
		return nil, errors.New("ent: MediaEventsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mes *MediaEventsSelect) BoolsX(ctx context.Context) []bool {
	v, err := mes.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (mes *MediaEventsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mes.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mediaevents.Label}
	default:
		err = fmt.Errorf("ent: MediaEventsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mes *MediaEventsSelect) BoolX(ctx context.Context) bool {
	v, err := mes.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mes *MediaEventsSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mes.fields {
		if !mediaevents.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := mes.sqlQuery().Query()
	if err := mes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mes *MediaEventsSelect) sqlQuery() sql.Querier {
	selector := mes.sql
	selector.Select(selector.Columns(mes.fields...)...)
	return selector
}