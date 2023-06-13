// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"code-connect/ent/migrate"

	"code-connect/ent/problem"
	"code-connect/ent/record"
	"code-connect/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Problem is the client for interacting with the Problem builders.
	Problem *ProblemClient
	// Record is the client for interacting with the Record builders.
	Record *RecordClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Problem = NewProblemClient(c.config)
	c.Record = NewRecordClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Problem: NewProblemClient(cfg),
		Record:  NewRecordClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Problem: NewProblemClient(cfg),
		Record:  NewRecordClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Problem.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Problem.Use(hooks...)
	c.Record.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Problem.Intercept(interceptors...)
	c.Record.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ProblemMutation:
		return c.Problem.mutate(ctx, m)
	case *RecordMutation:
		return c.Record.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ProblemClient is a client for the Problem schema.
type ProblemClient struct {
	config
}

// NewProblemClient returns a client for the Problem from the given config.
func NewProblemClient(c config) *ProblemClient {
	return &ProblemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `problem.Hooks(f(g(h())))`.
func (c *ProblemClient) Use(hooks ...Hook) {
	c.hooks.Problem = append(c.hooks.Problem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `problem.Intercept(f(g(h())))`.
func (c *ProblemClient) Intercept(interceptors ...Interceptor) {
	c.inters.Problem = append(c.inters.Problem, interceptors...)
}

// Create returns a builder for creating a Problem entity.
func (c *ProblemClient) Create() *ProblemCreate {
	mutation := newProblemMutation(c.config, OpCreate)
	return &ProblemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Problem entities.
func (c *ProblemClient) CreateBulk(builders ...*ProblemCreate) *ProblemCreateBulk {
	return &ProblemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Problem.
func (c *ProblemClient) Update() *ProblemUpdate {
	mutation := newProblemMutation(c.config, OpUpdate)
	return &ProblemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProblemClient) UpdateOne(pr *Problem) *ProblemUpdateOne {
	mutation := newProblemMutation(c.config, OpUpdateOne, withProblem(pr))
	return &ProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProblemClient) UpdateOneID(id int) *ProblemUpdateOne {
	mutation := newProblemMutation(c.config, OpUpdateOne, withProblemID(id))
	return &ProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Problem.
func (c *ProblemClient) Delete() *ProblemDelete {
	mutation := newProblemMutation(c.config, OpDelete)
	return &ProblemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProblemClient) DeleteOne(pr *Problem) *ProblemDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProblemClient) DeleteOneID(id int) *ProblemDeleteOne {
	builder := c.Delete().Where(problem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProblemDeleteOne{builder}
}

// Query returns a query builder for Problem.
func (c *ProblemClient) Query() *ProblemQuery {
	return &ProblemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProblem},
		inters: c.Interceptors(),
	}
}

// Get returns a Problem entity by its id.
func (c *ProblemClient) Get(ctx context.Context, id int) (*Problem, error) {
	return c.Query().Where(problem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProblemClient) GetX(ctx context.Context, id int) *Problem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRecords queries the records edge of a Problem.
func (c *ProblemClient) QueryRecords(pr *Problem) *RecordQuery {
	query := (&RecordClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, id),
			sqlgraph.To(record.Table, record.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, problem.RecordsTable, problem.RecordsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProblemClient) Hooks() []Hook {
	return c.hooks.Problem
}

// Interceptors returns the client interceptors.
func (c *ProblemClient) Interceptors() []Interceptor {
	return c.inters.Problem
}

func (c *ProblemClient) mutate(ctx context.Context, m *ProblemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProblemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProblemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProblemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Problem mutation op: %q", m.Op())
	}
}

// RecordClient is a client for the Record schema.
type RecordClient struct {
	config
}

// NewRecordClient returns a client for the Record from the given config.
func NewRecordClient(c config) *RecordClient {
	return &RecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `record.Hooks(f(g(h())))`.
func (c *RecordClient) Use(hooks ...Hook) {
	c.hooks.Record = append(c.hooks.Record, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `record.Intercept(f(g(h())))`.
func (c *RecordClient) Intercept(interceptors ...Interceptor) {
	c.inters.Record = append(c.inters.Record, interceptors...)
}

// Create returns a builder for creating a Record entity.
func (c *RecordClient) Create() *RecordCreate {
	mutation := newRecordMutation(c.config, OpCreate)
	return &RecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Record entities.
func (c *RecordClient) CreateBulk(builders ...*RecordCreate) *RecordCreateBulk {
	return &RecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Record.
func (c *RecordClient) Update() *RecordUpdate {
	mutation := newRecordMutation(c.config, OpUpdate)
	return &RecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RecordClient) UpdateOne(r *Record) *RecordUpdateOne {
	mutation := newRecordMutation(c.config, OpUpdateOne, withRecord(r))
	return &RecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RecordClient) UpdateOneID(id int) *RecordUpdateOne {
	mutation := newRecordMutation(c.config, OpUpdateOne, withRecordID(id))
	return &RecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Record.
func (c *RecordClient) Delete() *RecordDelete {
	mutation := newRecordMutation(c.config, OpDelete)
	return &RecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RecordClient) DeleteOne(r *Record) *RecordDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RecordClient) DeleteOneID(id int) *RecordDeleteOne {
	builder := c.Delete().Where(record.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RecordDeleteOne{builder}
}

// Query returns a query builder for Record.
func (c *RecordClient) Query() *RecordQuery {
	return &RecordQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRecord},
		inters: c.Interceptors(),
	}
}

// Get returns a Record entity by its id.
func (c *RecordClient) Get(ctx context.Context, id int) (*Record, error) {
	return c.Query().Where(record.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RecordClient) GetX(ctx context.Context, id int) *Record {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProblem queries the problem edge of a Record.
func (c *RecordClient) QueryProblem(r *Record) *ProblemQuery {
	query := (&ProblemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(record.Table, record.FieldID, id),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, record.ProblemTable, record.ProblemPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RecordClient) Hooks() []Hook {
	return c.hooks.Record
}

// Interceptors returns the client interceptors.
func (c *RecordClient) Interceptors() []Interceptor {
	return c.inters.Record
}

func (c *RecordClient) mutate(ctx context.Context, m *RecordMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RecordCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RecordUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RecordDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Record mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Problem, Record, User []ent.Hook
	}
	inters struct {
		Problem, Record, User []ent.Interceptor
	}
)