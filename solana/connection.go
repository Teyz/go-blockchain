package solana

import (
	"context"
	"fmt"
)

func NewSolanaConnection(ctx context.Context, cfg *Config) string {
	return fmt.Sprintf("%s/%s", cfg.SolanaHost, cfg.SolanaApiKey)
}
