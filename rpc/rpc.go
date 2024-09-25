package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func SolanaCallRPC(method string, blockchainUrl string, params []interface{}, result interface{}) error {
	reqBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}

	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		log.Error().Err(err).
			Msg("pkg.blockchain.rpc.SolanaCallRPC: unable to marshal request")
		return err
	}

	resp, err := http.Post(blockchainUrl, "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		log.Error().Err(err).
			Msg("pkg.blockchain.rpc.SolanaCallRPC: unable to post request")
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).
			Msg("pkg.blockchain.rpc.SolanaCallRPC: unable to read resp.Body")
		return err
	}

	var rpcResponse struct {
		Jsonrpc string          `json:"jsonrpc"`
		Result  json.RawMessage `json:"result"`
		Error   *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		log.Error().Err(err).
			Msg("pkg.blockchain.rpc.SolanaCallRPC: unable to unmarshal body")
		return err
	}

	if rpcResponse.Error != nil {
		log.Error().Err(err).
			Msg("pkg.blockchain.rpc.SolanaCallRPC: error in rpcResp")
		return fmt.Errorf("pkg.blockchain.rpc.SolanaCallRPC: error in rpcResp: %s", rpcResponse.Error.Message)
	}

	if err := json.Unmarshal(rpcResponse.Result, &result); err != nil {
		log.Error().Err(err).
			Msg("pkg.blockchain.rpc.SolanaCallRPC: unable to unmarshal result")
		return err
	}

	return nil
}
