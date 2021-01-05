package oneinch

import "github.com/ethereum/go-ethereum/common"

var exchangeContractHistoryEventSig = common.HexToHash("0x894dbf1262199c24e1750298a384c709160f49d163422cc6cee694c73713f1d2")

// IsHistoryEvent is exchange contract "History" event
func IsHistoryEvent(topic common.Hash) bool {
	return topic == exchangeContractHistoryEventSig
}
