# Bounded Concurrent Order Processor

Implement `ProcessOrders` in `order_processor.go`. This is a learner-owned exercise: make a meaningful implementation attempt before requesting code changes or a solution.

## Requirements

1. Process all orders using a worker pool.
2. Maximum processing concurrency must be `workerCount`.
3. Use `golang.org/x/sync/errgroup` with `errgroup.WithContext`.
4. Use a jobs channel to distribute `Order` values.
5. Workers consume jobs from the shared jobs channel.
6. Pass the errgroup-derived context to `process(ctx, order)`.
7. If a processor returns an error, cancel remaining work through the errgroup context, stop producing unnecessary jobs, and allow started goroutines to terminate cleanly.
8. Parent-context cancellation must stop the operation.
9. Do not return while goroutines owned by `ProcessOrders` are still running.
10. Avoid goroutine leaks and deadlocks.
11. Return the error produced by the errgroup.
12. Return a validation error when `workerCount <= 0`.
13. Do not use `time.Sleep` for synchronization.

## Run

From this directory:

```sh
go test
go test -race
```

The starter is intentionally incomplete, so behavior tests should fail until you implement `ProcessOrders`.

## Before asking for implementation help

Share your intended invariants, your attempted code, the observed failure, and your current hypothesis. Questions and bounded hints are available before a complete solution.
