package transaction

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	// new JSON tx
	from := common.HexToAddress("0xa5aa21a837321563dc6889ddaa1bea808cd5224c")
	to := common.HexToAddress("0xd166bffd7066d989ee21ad0f51271c6e137a70ab")
	rawTx := NewTransaction(from, &to, big.NewInt(0), big.NewInt(22000), big.NewInt(123), common.Hex2Bytes("0x0656678909"))

	// print
	fmt.Println(">>>>>>>>>>> rawTx:", "\n",
		"From:", rawTx.From.String(), "\n",
		"To:", rawTx.To.String(), "\n",
		"Value:", rawTx.Value, "\n",
		"gasLimit:", rawTx.GasLimit, "\n",
		"gasPrice:", rawTx.GasPrice,
	)

	/*kv := make(map[string]interface{})
	value := reflect.ValueOf(rawTx)
	types  := reflect.TypeOf(rawTx)

	for i:=0; i<value.NumField(); i++ {
		kv[types.Field(i).Name] = value.Field(i)
	}

	for k, v := range kv {
		fmt.Println(k, ": ", v)
	}*/
}