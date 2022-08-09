## 2PAY GOLANG SDK

[2Pay Online Api](https://2pay.gitbook.io/2pay-api-docs-en/)



### Installation


```shell
$ go get github.com/2pay-co/2pay-golang-sdk
```


### Usage

Please see [examples](https://github.com/2pay-co/2pay-golang-sdk/tree/main/tests)

### demo

```go
use Pay\TwoPay;
use Pay\SecurePay;

ssssss

$secure_param = new SecurePay();
$secure_param->amount = "3.2";
$secure_param->callbackUrl = "";
$secure_param->currency ="USD";
$secure_param->goodsInfo = "";
$secure_param->note = "test order";
$secure_param->description = "order test";
$secure_param->terminal = "ONLINE";
$secure_param->vendor = "alipay";
$secure_param->ipnUrl = '';
$secure_param->reference = (string)time();
$secure_param->timeout = "120";

$two_pay = new TwoPay("merchant_no", "token");
// $res is json string
$res = $two_pay->SecurePay($secure_param);

var_dump($res);
```
