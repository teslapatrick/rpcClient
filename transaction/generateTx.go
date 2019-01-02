package transaction

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type Message struct {
	To       *common.Address  `json:"to"`
	From     common.Address   `json:"from"`
	Value    string           `json:"value"`
	GasLimit string           `json:"gas"`
	GasPrice string           `json:"gasPrice"`
	Data     []byte           `json:"data"`
}

func NewTransaction(from common.Address, to *common.Address, value *big.Int, gasLimit *big.Int, gasPrice *big.Int, data []byte) Message {

	return Message{
		From: from,
		To: to,
		Value: toJsonArg(value),
		GasLimit: toJsonArg(gasLimit),
		GasPrice: toJsonArg(gasPrice),
		Data: data,
	}
}

func toJsonArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}
