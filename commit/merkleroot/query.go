package merkleroot

import (
	"context"
)

func (p *Processor) Query(_ context.Context, _ Outcome) (Query, error) {
	return Query{}, nil
}
