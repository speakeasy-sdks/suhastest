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