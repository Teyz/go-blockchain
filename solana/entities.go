package solana

type AccountInfoResponse struct {
	Lamports   uint64   `json:"lamports"`
	Owner      string   `json:"owner"`
	Data       []string `json:"data"`
	Executable bool     `json:"executable"`
}

type ConfirmedBlockResponse struct {
	BlockHeight  uint64        `json:"blockHeight"`
	BlockTime    int64         `json:"blockTime"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Signature string `json:"signature"`
}

type RecentBlockhashResponse struct {
	Blockhash     string `json:"blockhash"`
	FeeCalculator struct {
		LamportsPerSignature uint64 `json:"lamportsPerSignature"`
	} `json:"feeCalculator"`
}

type EpochInfoResponse struct {
	AbsoluteSlot     uint64 `json:"absoluteSlot"`
	Epoch            uint64 `json:"epoch"`
	SlotIndex        uint64 `json:"slotIndex"`
	SlotsInEpoch     uint64 `json:"slotsInEpoch"`
	TransactionCount uint64 `json:"transactionCount"`
}

type SupplyResponse struct {
	Total          uint64 `json:"total"`
	Circulating    uint64 `json:"circulating"`
	NonCirculating uint64 `json:"nonCirculating"`
}

type SignatureStatusResponse struct {
	Slot          uint64 `json:"slot"`
	Confirmations uint64 `json:"confirmations"`
	Status        struct {
		Ok interface{} `json:"Ok"`
	} `json:"status"`
}

type VoteAccountsResponse struct {
	Current    []ValidatorInfo `json:"current"`
	Delinquent []ValidatorInfo `json:"delinquent"`
}

type ValidatorInfo struct {
	VotePubkey string `json:"votePubkey"`
	NodePubkey string `json:"nodePubkey"`
	Commission int    `json:"commission"`
}

type ClusterNodeResponse struct {
	Pubkey  string `json:"pubkey"`
	Gossip  string `json:"gossip"`
	Rpc     string `json:"rpc"`
	Version string `json:"version"`
}

type VersionResponse struct {
	SolanaCore string `json:"solana-core"`
}

type TokenAccountResponse struct {
	Account string `json:"account"`
	Mint    string `json:"mint"`
	Owner   string `json:"owner"`
	Balance uint64 `json:"balance"`
}

type TokenSupplyResponse struct {
	Amount uint64 `json:"amount"`
}

type TokenAccountBalanceResponse struct {
	Amount string `json:"amount"`
}

type ProgramAccountResponse struct {
	Pubkey  string `json:"pubkey"`
	Account struct {
		Lamports   uint64 `json:"lamports"`
		Owner      string `json:"owner"`
		Executable bool   `json:"executable"`
		RentEpoch  uint64 `json:"rentEpoch"`
		Data       string `json:"data"`
	} `json:"account"`
}

type ConfirmedTransactionResponse struct {
	Slot        uint64          `json:"slot"`
	Transaction Transaction     `json:"transaction"`
	Meta        TransactionMeta `json:"meta"`
	BlockTime   *int64          `json:"blockTime,omitempty"`
}

type TransactionMessage struct {
	AccountKeys     []string      `json:"accountKeys"`
	Header          MessageHeader `json:"header"`
	Instructions    []Instruction `json:"instructions"`
	RecentBlockhash string        `json:"recentBlockhash"`
}

type MessageHeader struct {
	NumRequiredSignatures       uint8 `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}

type Instruction struct {
	ProgramIDIndex uint8   `json:"programIdIndex"`
	Accounts       []uint8 `json:"accounts"`
	Data           string  `json:"data"` // Data encodée en base64
}

type TransactionMeta struct {
	Fee               uint64             `json:"fee"`
	PreBalances       []uint64           `json:"preBalances"`
	PostBalances      []uint64           `json:"postBalances"`
	InnerInstructions []InnerInstruction `json:"innerInstructions,omitempty"`
	LogMessages       []string           `json:"logMessages,omitempty"`
	Status            struct {
		Ok interface{} `json:"Ok"`
	} `json:"status"`
}

type InnerInstruction struct {
	Index        uint64        `json:"index"`
	Instructions []Instruction `json:"instructions"`
}
