# Senior Go Backend Testing Workshop

> You own the tests in this lab. Production code and test infrastructure are provided; the exercise assertions remain TODOs and there is no solution directory.

## Goal

Practice testing a small order flow at two boundaries:

- unit-test `OrderService` with `testify/mock`
- integration-test a real PostgreSQL repository
- verify error identity through wrapping
- prove transaction rollback and test isolation
- exercise a concurrent last-item purchase

The implementation intentionally stays small: one model, one service, one repository interface, and one PostgreSQL repository.

## Directory structure

```text
15-go-testing-workshop/
├── compose.yml                         # local PostgreSQL for integration tests
├── migrations/
│   └── 001_create_orders.sql           # products and orders schema
├── mocks/
│   └── order_repository.go             # testify/mock repository
├── starter/
│   ├── order.go                        # model, errors, and repository interface
│   ├── postgres_repository.go          # transactional PostgreSQL implementation
│   └── service.go                      # PlaceOrder business rules
└── tests/
    ├── service_test.go                 # unit-test exercises
    ├── integration_helpers_test.go     # setup, migration, fixtures, and cleanup
    └── postgres_repository_integration_test.go
                                           # integration-test exercises
```

## Run the unit tests

From this lab directory:

```sh
go test ./...
```

The learner tests initially report `SKIP`. Remove one exercise's `t.Skip` only when you start implementing that test. Do not remove every skip at once.

## Prepare PostgreSQL

The simplest local setup uses the included Docker Compose service:

```sh
docker compose up -d
docker compose ps
```

The default test connection is:

```text
postgres://postgres:postgres@localhost:55432/go_testing_workshop?sslmode=disable
```

To use another PostgreSQL instance, set `TEST_DATABASE_URL` to a disposable test database. The integration helpers truncate `orders` and `products`, so never point this variable at a database containing valuable data.

The helpers automatically apply `migrations/001_create_orders.sql`, clean the tables before each exercise, and register cleanup after it. If PostgreSQL is unavailable, an enabled integration exercise is skipped with the connection error.

## Run the integration tests

Integration tests are excluded from ordinary test runs by the `integration` build tag:

```sh
go test -tags=integration ./tests
```

To run one exercise while developing it:

```sh
go test -tags=integration ./tests -run TestPostgreSQLCreateOrderSuccess -count=1
```

After the concurrency exercise is implemented, also run:

```sh
go test -tags=integration -race ./tests -run TestPostgreSQLCreateOrderConcurrentLastItem -count=1
```

Stop the workshop database when finished:

```sh
docker compose down
```

## Exercise order

1. **Invalid quantity** — verify `ErrInvalidQuantity` and prove the mock repository received no call.
2. **Successful unit request** — return the repository result and verify the mock arguments.
3. **Repository error** — verify the service wraps the error without breaking `ErrorIs`.
4. **Successful integration request** — prove the returned order and database state agree.
5. **Out of stock** — prove `ErrOutOfStock` and unchanged database state.
6. **Insert failure** — force a failure after the stock update and prove rollback removed all partial effects.
7. **Concurrent last item** — with stock `1`, prove exactly one of two competing orders succeeds.

## Expected behavior

- Quantities of zero or less are rejected before repository access.
- Valid service calls delegate the original context, product ID, and quantity.
- Repository errors retain their identity when crossing the service boundary.
- Stock is decreased only when enough stock exists.
- Stock decrease and order insertion succeed or fail together.
- An out-of-stock attempt changes neither stock nor orders.
- Two concurrent requests cannot both purchase the final item.
- Every integration exercise starts from data it creates and cleans up afterward.

The TODOs deliberately do not prescribe exact assertions, goroutine coordination, or error-counting strategy. Those choices are part of the workshop.
