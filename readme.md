## RPC-Client

> An ethereum like JSON-RPC Client

* **Client struct**

```go
type Client struct {
	rpcClient *rpc.Client
	ethClient *ethclient.Client
}
```



* **Connect**
```go
func Connect(host string) (*Client, error)
```

​	`eg:  client, err := Connect("http://127.0.0.1:8545")`

---

* **GetBlockNumber**

```go
func (ec *Client) GetBlockNumber(ctx context.Context) (*big.Int, error)
```

​	Get current blocknumber.

---

* **GetBlockHeaderByNumber**

```go
func (ec *Client) GetBlockHeaderByNumber(ctx context.Context, blkNum *big.Int) (*types.Header, error)
```

​	Get block header by block number.

---

* **GetNonceByAddress**

```go
func (ec *Client) GetNonceByAddress(ctx context.Context, address common.Address) (uint64, error)
```

​	Get transaction count of address

---

* **GetNodeListening**

```go
func (ec *Client) GetNodeListening(ctx context.Context) (bool, error)
```

​	Get node status.

---

* **SendTransaction**

```go

type Message struct {
	To 	   *common.Address 	`json:"to"`
	From   common.Address 	`json:"from"`
	Value  string			`json:"value"`
	GasLimit string			`json:"gas"`
	GasPrice string			`json:"gasPrice"`
	Data     []byte			`json:"data"`

}

func (ec *Client) SendTransaction(ctx context.Context, tx *transaction.Message) (common.Hash, error)
```

​	Send transaction.

---

* **SendRawTransaction**

```go
func (ec *Client) SendRawTransaction(ctx context.Context, rawTxData []byte) (common.Hash, error)
```

​	Send Raw transaction.