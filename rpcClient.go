package rpcClient

import (
	"allTestProject/transaction"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type Client struct {
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	host      string
}

type Clients struct {
	clients []*Client
}

func New() *Clients {
	return &Clients{}
}

func (c *Clients) AddClient(host string) error {
	// new rpc client
	rpcclient, err := rpc.Dial(host)
	if err != nil {
		log.Error("Host connect error...", "err", err)
		return err
	}
	// new eth client
	ethclient := ethclient.NewClient(rpcclient)
	c.clients = append(c.clients, &Client{rpcclient, ethclient, host})

	// test node
	l := len(c.clients)
	if  _, err := c.clients[l-1].GetNodeListening(context.TODO()); err != nil {
		return err
	}

	return nil
}

func (c *Clients) DelClient(host string) error {
	l := len(c.clients)
	if l == 0 {
		return errors.New("node length is zero")
	}

	for i, v := range c.clients {
		if v.host == host {
			c.clients[i] = c.clients[l-1]
			c.clients = c.clients[:l-1]
			return nil
		}
	}
	return errors.New("could not find node")
}

func (c *Clients) GetClients() map[int]string {
	res := make(map[int]string)
	for i, v := range c.clients {
		res[i] = v.host
	}
	return res
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

