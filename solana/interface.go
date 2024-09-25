package solana

type Solana interface {
	GetBalance(solAddress string) (uint64, error)
}
