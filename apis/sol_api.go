package apis

const (
	SOL_ACCOUNT  = ""
	SOL_TRANSFER = "https://pro-api.solscan.io/v2.0/account/transfer?page=1&page_size=10&sort_by=block_time&sort_order=desc"
)

type SOL_TXResponse struct {
	Success bool              `json:"success"`
	Data    []SOL_Transaction `json:"data"`
}
type SOL_Transaction struct {
	BlockId      int64   `json:"block_id"`
	TransId      string  `json:"trans_id"`
	BlockTime    int64   `json:"block_time"`
	ActivityType string  `json:"activity_type"`
	FromAddress  string  `json:"from_address"`
	ToAddress    string  `json:"to_address"`
	TokenAddress string  `json:"token_address"`
	TokenDecimal int64   `json:"token_decimals"`
	Amount       int64   `json:"amount"`
	Flow         string  `json:"flow"`
	Value        float64 `json:"value"`
	Time         string  `json:"time"`
}
