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
    s := suhas.New()

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
        CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
            CustomerDetails: shared.CustomerDetails{
                CustomerBankAccountNumber: suhas.String("corrupti"),
                CustomerBankCode: suhas.String("provident"),
                CustomerBankIfsc: suhas.String("distinctio"),
                CustomerEmail: suhas.String("quibusdam"),
                CustomerID: "unde",
                CustomerPhone: "nulla",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: suhas.String("2021-07-29T00:00:00Z"),
            OrderID: suhas.String("corrupti"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: suhas.String("illum"),
                PaymentMethods: suhas.String("vel"),
                ReturnURL: suhas.String("error"),
            },
            OrderNote: suhas.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: suhas.Float64(6458.94),
                    Percentage: suhas.Float64(3843.82),
                    VendorID: suhas.String("iure"),
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
        XAPIVersion: suhas.String("suscipit"),
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