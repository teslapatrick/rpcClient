package rpcClient

import (
	"allTestProject/transaction"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestConnect(t *testing.T) {
	// new client
	//client, err := Connect("http://127.0.0.1:8545")

	client := New()
	err := client.AddClient("http://127.0.0.1:8545")
	if err != nil {
		t.Fatalf("Connect error: %v", err)
	}
	// check node client
	listening, _ := client.clients[0].GetNodeListening(context.TODO())
	fmt.Println("listening:", listening)
}

func TestClient_SendTransaction(t *testing.T) {
	// new client
	client := New()
	err := client.AddClient("http://127.0.0.1:8545")
	if err != nil {
		t.Fatalf("Connect error: %v", err)
	}

	// new JSON tx
	from := common.HexToAddress("0xa5aa21a837321563dc6889ddaa1bea808cd5224c")
	to := common.HexToAddress("0xd166bffd7066d989ee21ad0f51271c6e137a70ab")
	rawTx := transaction.NewTransaction(from, &to, big.NewInt(0), big.NewInt(90000), big.NewInt(123), common.Hex2Bytes("0x65"))

	// send TX
	txHash, err := client.clients[0].SendTransaction(context.TODO(), &rawTx)

	if err != nil{
		t.Fatalf("error occur: %v", err)
	}

	// print tx hash
	fmt.Println(">>>>>>>>>>> sendTX", "txHash", txHash.String())
}

func TestClient_SendRawTransaction(t *testing.T) {

	// new client
	client := New()
	err := client.AddClient("http://127.0.0.1:8545")
	if err != nil {
		t.Fatalf("Connect error: %v", err)
	}

	// new JSON rawTx
	rawTx := common.Hex2Bytes("0xf84c83011530850430e2340083015f9094d166bffd7066d989ee21ad0f51271c6e137a70ab820fff80808080a0ffffffffff29b5b42c688e5eb7cae79f1dd5c4073eb380e656e27bf6e1a7d6bb01")

	// send TX
	txHash, err := client.clients[0].SendRawTransaction(context.TODO(), rawTx)

	if err != nil {
		t.Fatalf("error: %v", err)
	}
	// print tx hash
	fmt.Println(">>>>>>>>>>> sendTX", "txHash", txHash.String())

}

func TestDelClient(t *testing.T)  {
	// new client
	client := New()
	err := client.AddClient("http://127.0.0.1:8545")
	if err != nil {
		t.Fatalf("Connect error: %v", err)
	}

	// del client
	err = client.DelClient("http://127.0.0.1:8545")
	if err != nil {
		t.Fatalf("error found: %v", err)
	}
}