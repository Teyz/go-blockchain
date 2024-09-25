package solana

import "github.com/teyz/go-blockchain/rpc"

func (s *solanaClient) GetBalance(solAddress string) (uint64, error) {
	var result struct {
		Value uint64 `json:"value"`
	}
	err := rpc.SolanaCallRPC("getBalance", s.solClient, []interface{}{solAddress}, &result)
	if err != nil {
		return 0, err
	}
	return result.Value, nil
}
