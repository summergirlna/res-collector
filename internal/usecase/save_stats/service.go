package save_stats

import "context"

type Service interface {
	Save(ctx context.Context, in *Input) (*Output, error)
}
