package etherscan

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
)

// Envelope is the carrier of nearly every response
type Envelope struct {
	// 1 for good, 0 for error
	Status int `json:"status,string"`
	// OK for good, other words when Status equals 0
	Message string `json:"message"`
	// where response lies
	Result json.RawMessage `json:"result"`
}

// NormalTx holds info from normal tx query
type NormalTx struct {
	BlockNumber       uint64                `json:"blockNumber,string"`
	Timestamp         uint64                `json:"timeStamp,string"`
	Hash              common.Hash           `json:"hash"`
	Nonce             uint64                `json:"nonce,string"`
	BlockHash         common.Hash           `json:"blockHash"`
	TransactionIndex  uint                  `json:"transactionIndex,string"`
	From              string                `json:"from"`
	To                string                `json:"to"`
	Value             *math.HexOrDecimal256 `json:"value"`
	Gas               uint64                `json:"gas,string"`
	GasPrice          *math.HexOrDecimal256 `json:"gasPrice"`
	IsError           int                   `json:"isError,string"`
	TxReceiptStatus   uint64                `json:"txreceipt_status,string"`
	Input             hexutil.Bytes         `json:"input"`
	ContractAddress   string                `json:"contractAddress"`
	CumulativeGasUsed uint64                `json:"cumulativeGasUsed,string"`
	GasUsed           uint64                `json:"gasUsed,string"`
	Confirmations     int64                 `json:"confirmations,string"`
}

// InternalTx holds info from internal tx query
type InternalTx struct {
	BlockNumber     uint64                `json:"blockNumber,string"`
	TimeStamp       uint64                `json:"timeStamp,string"`
	Hash            common.Hash           `json:"hash"`
	From            string                `json:"from"`
	To              string                `json:"to"`
	Value           *math.HexOrDecimal256 `json:"value"`
	ContractAddress string                `json:"contractAddress"`
	Input           string                `json:"input"`
	Type            string                `json:"type"`
	Gas             uint64                `json:"gas,string"`
	GasUsed         uint64                `json:"gasUsed,string"`
	TraceID         string                `json:"traceId"`
	IsError         int                   `json:"isError,string"`
	ErrCode         string                `json:"errCode"`
}

// ERC20Transfer holds info from ERC20 token transfer event query
type ERC20Transfer struct {
	BlockNumber       uint64                `json:"blockNumber,string"`
	TimeStamp         uint64                `json:"timeStamp,string"`
	Hash              common.Hash           `json:"hash"`
	Nonce             uint64                `json:"nonce,string"`
	BlockHash         common.Hash           `json:"blockHash"`
	From              string                `json:"from"`
	ContractAddress   common.Address        `json:"contractAddress"`
	To                string                `json:"to"`
	Value             *math.HexOrDecimal256 `json:"value"`
	TokenName         string                `json:"tokenName"`
	TokenSymbol       string                `json:"tokenSymbol"`
	TokenDecimal      uint8                 `json:"tokenDecimal,string"`
	TransactionIndex  uint                  `json:"transactionIndex,string"`
	Gas               uint64                `json:"gas,string"`
	GasPrice          *math.HexOrDecimal256 `json:"gasPrice"`
	GasUsed           uint64                `json:"gasUsed,string"`
	CumulativeGasUsed uint64                `json:"cumulativeGasUsed,string"`
	Input             string                `json:"input"`
	Confirmations     int64                 `json:"confirmations,string"`
}
