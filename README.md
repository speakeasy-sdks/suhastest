# github.com/speakeasy-sdks/suhastest

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/suhastest
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/suhastest"
	"github.com/speakeasy-sdks/suhastest/pkg/models/operations"
	"github.com/speakeasy-sdks/suhastest/pkg/models/shared"
	"github.com/speakeasy-sdks/suhastest/pkg/models/callbacks"
	"net/http"
)

func main() {
    s := suhastest.New()

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
        CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
            CustomerDetails: shared.CustomerDetails{
                CustomerBankAccountNumber: suhastest.String("corrupti"),
                CustomerBankCode: suhastest.String("provident"),
                CustomerBankIfsc: suhastest.String("distinctio"),
                CustomerEmail: suhastest.String("quibusdam"),
                CustomerID: "unde",
                CustomerPhone: "nulla",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: suhastest.String("2021-07-29T00:00:00Z"),
            OrderID: suhastest.String("corrupti"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: suhastest.String("illum"),
                PaymentMethods: suhastest.String("vel"),
                ReturnURL: suhastest.String("error"),
            },
            OrderNote: suhastest.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: suhastest.Float64(6458.94),
                    Percentage: suhastest.Float64(3843.82),
                    VendorID: suhastest.String("iure"),
                },
            },
            OrderTags: map[string]string{
                "magnam": "debitis",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "ipsa",
                TerminalPhoneNo: "delectus",
                TerminalType: "tempora",
            },
        },
        XAPIVersion: suhastest.String("suscipit"),
        XClientID: "molestiae",
        XClientSecret: "minus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Orders](docs/sdks/orders/README.md)

* [CreateOrder](docs/sdks/orders/README.md#createorder) - Create Order
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
