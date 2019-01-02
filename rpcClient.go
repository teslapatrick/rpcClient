package rpcClient

import (
	"allTestProject/transaction"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/nknorg/nnet/log"
	"math/big"

)

type Client struct {
	rpcClient *rpc.Client
	ethClient *ethclient.Client
}

func Connect(host string) (*Client, error) {
	rpcClient, err := rpc.Dial(host)
	if err != nil {
		log.Error("Host connect error...")
	}
	ethClient := ethclient.NewClient(rpcClient)
	return &Client{rpcClient, ethClient}, err
}

func (ec *Client) GetBlockNumber(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := ec.rpcClient.CallContext(ctx, &result, "eth_blockNumber")
	return (*big.Int)(&result), err
}

func (ec *Client) GetBlockHeaderByNumber(ctx context.Context, blkNum *big.Int) (*types.Header, error){
	var head *types.Header
	err := ec.rpcClient.CallContext(ctx, &head, "eth_getBlockByNumber", toJsonArg(blkNum), true)
	return head, err
}

func (ec *Client) GetNonceByAddress(ctx context.Context, address common.Address) (uint64, error) {
	var nonce uint64
	err := ec.rpcClient.CallContext(ctx, &nonce, "eth_getTransactionCount", address, "latest")
	return nonce, err
}

func (ec *Client) GetNodeListening(ctx context.Context) (bool, error) {
	var isOK bool
	err := ec.rpcClient.CallContext(ctx, &isOK, "net_listening")
	return isOK, err
}

func (ec *Client) SendTransaction(ctx context.Context, tx *transaction.Message) (common.Hash, error) {
	var TxHash common.Hash
	err := ec.rpcClient.CallContext(ctx, &TxHash, "eth_sendTransaction", tx)
	return TxHash, err
}

func (ec *Client) SendRawTransaction(ctx context.Context, rawTxData []byte) (common.Hash, error) {
	var TxHash common.Hash
	err := ec.rpcClient.CallContext(ctx, &TxHash, "eth_sendRawTransaction", rawTxData)
	return TxHash, err
}

func toJsonArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

