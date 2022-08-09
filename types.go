package pay

const (
	BaseUrl      = "https://api.2pay.co"
	SecurePayUrl = "/online/v1/secure-pay"
	RefundPayUrl = "/app-data-search/v1/refund"
	QueryPayUrl  = "/app-data-search/v1/tran-query"
)

type SecurePay struct {
	MerchantNo  string      `json:"merchantNo"`
	VerifySign  string      `json:"verifySign"`
	Currency    string      `json:"currency"`
	Amount      string      `json:"amount"`
	Vendor      string      `json:"vendor"`
	Reference   string      `json:"reference"`
	Terminal    string      `json:"terminal"`
	IpnURL      string      `json:"ipnUrl"`
	CallbackURL string      `json:"callbackUrl"`
	Note        string      `json:"note"`
	Description string      `json:"description"`
	Timeout     string      `json:"timeout"`
	GoodsInfo   []GoodsInfo `json:"goodsInfo"`
	CreditType  string      `json:"creditType"`
	CustomerNo  string      `json:"customerNo"`
}

type GoodsInfo struct {
	GoodsName string `json:"goodsName"`
	Quantity  string `json:"quantity"`
}

type SecurePayResponse struct {
	Result   string       `json:"ret_code"`
	RetMsg   string       `json:"ret_msg"`
	Response SecurePayRet `json:"result,omitempty"`
}

type SecurePayRet struct {
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	TransactionNo  string `json:"transactionNo"`
	Reference      string `json:"reference"`
	CashierURL     string `json:"cashierUrl"`
	SettleCurrency string `json:"settleCurrency"`
}

type BaseResponse struct {
	Result string `json:"ret_code"`
	RetMsg string `json:"ret_msg"`
}

type QueryResp struct {
	BaseResponse
	QueryRet QueryRet `json:"result"`
}

type QueryRet struct {
	TransactionNo  string  `json:"transactionNo"`
	Reference      string  `json:"reference"`
	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
	Currency       string  `json:"currency"`
	SettleCurrency string  `json:"settleCurrency"`
}
