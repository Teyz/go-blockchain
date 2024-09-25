package solana

type Config struct {
	SolanaHost   string `env:"SOL_HOST"`
	SolanaApiKey string `env:"SOL_API_KEY"`
}
