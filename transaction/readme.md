## Transaction

> Transaction

* **Json-rpc call message: TX**
```go
type Message struct {
	To 	     *common.Address 	`json:"to"`
	From     common.Address 	`json:"from"`
	Value    string				`json:"value"`
	GasLimit string				`json:"gas"`
	GasPrice string				`json:"gasPrice"`
	Data     []byte				`json:"data"`
}
```

---

* **NewTransaction**

```go
func NewTransaction(from common.Address, to *common.Address, value *big.Int, gasLimit *big.Int, gasPrice *big.Int, data []byte) Message
```

​	New Tx message.

