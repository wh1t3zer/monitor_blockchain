package apis

const (
	// 通过钱包地址查交易(仅代币)
	BSC_TRANSACTION = "https://api.bscscan.com/api?module=account&action=tokentx&page=1&offset=1&startblock=0&endblock=999999999&sort=desc"
)

type BSC_Transaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

type BSC_TXResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Result  []BSC_Transaction `json:"result"`
}
