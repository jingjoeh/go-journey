package orderprocessor

import "context"

type Order struct {
	ID int
}

type Processor func(ctx context.Context, order Order) error

func ProcessOrders(
	ctx context.Context,
	orders []Order,
	workerCount int,
	process Processor,
) error {
	// TODO: learner implementation
	return nil
}
