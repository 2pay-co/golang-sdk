package pay

import (
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
)

func NewPay() *TwoPay {
	return New("test_merchant_no", "test_token")
}

func TestTwoPay_SecurePay(t1 *testing.T) {

	t := NewPay()
	param := SecurePay{
		Currency:    "USD",
		Amount:      "10",
		Vendor:      "alipay",
		Reference:   gtime.Datetime(),
		Terminal:    "ONLINE",
		IpnURL:      "https://google.com",
		CallbackURL: "https://google.com",
		Note:        "test order",
		Description: "test order",
		Timeout:     "120",
		GoodsInfo: []GoodsInfo{
			{
				"test goods 1", "1",
			},
			{
				"test goods 2", "1",
			},
		},
	}

	resp, err := t.SecurePay(&param)
	t1.Log(resp, err)

}

func TestTwoPay_Query(t *testing.T) {
	pay := NewPay()
	resp, err := pay.Query("2022-08-09 18:39:03")

	t.Logf("resp: %#v, err: %v", resp, err)
}
