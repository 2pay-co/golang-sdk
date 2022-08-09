## 2PAY GOLANG SDK

[2Pay Online Api](https://2pay.gitbook.io/2pay-api-docs-en/)



### Installation

```shell
$ go get github.com/2pay-co/golang-sdk
```

### Usage

Please see [examples](https://github.com/2pay-co/2pay-golang-sdk/tree/main/tests)

### demo

```go
package main

import (
	"fmt"
	pay "github.com/2pay-co/golang-sdk"
	"github.com/gogf/gf/v2/os/gtime"
)

func main() {

	t := pay.New("", "")
	param := pay.SecurePay{
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
		GoodsInfo: []pay.GoodsInfo{
			{
				"test goods 1", "1",
			},
			{
				"test goods 2", "1",
			},
		},
	}

	resp, err := t.SecurePay(&param)
	fmt.Println(resp, err)
}
```
