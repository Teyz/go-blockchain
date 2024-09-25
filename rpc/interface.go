package rpc

type Rpc interface {
	SolanaCallRPC(method string, blockchainUrl string, params []interface{}, result interface{}) error
}
