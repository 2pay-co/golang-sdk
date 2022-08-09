package pay

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
	"sort"
)

type TwoPay struct {
	MerchantNo string
	Token      string
}

func New(merchantNo, token string) *TwoPay {
	return &TwoPay{
		MerchantNo: merchantNo,
		Token:      token,
	}
}

func (t *TwoPay) SecurePay(param *SecurePay) (*SecurePayResponse, error) {

	paramMap := gconv.MapDeep(param)
	paramMap["verifySign"] = t.sign(paramMap)

	securePayUrl := fmt.Sprintf("%s%s", BaseUrl, SecurePayUrl)
	body, err := PostJson(securePayUrl, paramMap)
	if err != nil {
		return nil, err
	}

	var resp SecurePayResponse
	err = json.Unmarshal(body, &resp)
	return &resp, err
}

func (t *TwoPay) Refund(orderNo string, amount float64) (*BaseResponse, error) {

	var paramMap = g.MapStrAny{
		"reference": orderNo,
		"amount":    decimal.NewFromFloat(amount).String(),
	}

	paramMap["verifySign"] = t.sign(paramMap)
	securePayUrl := fmt.Sprintf("%s%s", BaseUrl, RefundPayUrl)
	body, err := PostJson(securePayUrl, paramMap)
	if err != nil {
		return nil, err
	}

	var resp BaseResponse
	err = json.Unmarshal(body, &resp)
	return &resp, err
}

func (t *TwoPay) Query(orderNo string) (*QueryResp, error) {

	var paramMap = g.MapStrAny{
		"reference": orderNo,
	}

	paramMap["verifySign"] = t.sign(paramMap)
	securePayUrl := fmt.Sprintf("%s%s", BaseUrl, QueryPayUrl)
	body, err := PostJson(securePayUrl, paramMap)
	if err != nil {
		return nil, err
	}

	var resp QueryResp
	err = json.Unmarshal(body, &resp)
	return &resp, err
}

func (t *TwoPay) sign(param map[string]interface{}) string {

	param["merchantNo"] = t.MerchantNo

	var keys []string
	for key, _ := range param {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var signStr string
	for _, key := range keys {

		if key == "verifySign" {
			continue
		}

		valStr := gconv.String(param[key])
		if valStr == "" {
			continue
		}

		if key == "goodsInfo" && valStr != "" {
			param[key] = valStr
		}

		signStr += fmt.Sprintf("%s=%s&", key, valStr)
	}

	signStr += gmd5.MustEncryptString(t.Token)
	return gmd5.MustEncryptString(signStr)
}
