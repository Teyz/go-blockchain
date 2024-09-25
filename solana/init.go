package solana

import (
	"context"
)

type solanaClient struct {
	solClient string
}

func NewSolanaClient(ctx context.Context, url string) *solanaClient {
	return &solanaClient{
		solClient: url,
	}
}
