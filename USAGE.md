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
                    Amount: suhas.Float64(3843.82),
                    Percentage: suhas.Float64(4375.87),
                    VendorID: suhas.String("magnam"),
                },
                shared.VendorSplit{
                    Amount: suhas.Float64(8917.73),
                    Percentage: suhas.Float64(567.13),
                    VendorID: suhas.String("delectus"),
                },
                shared.VendorSplit{
                    Amount: suhas.Float64(2726.56),
                    Percentage: suhas.Float64(3834.41),
                    VendorID: suhas.String("molestiae"),
                },
            },
            OrderTags: map[string]string{
                "placeat": "voluptatum",
                "iusto": "excepturi",
                "nisi": "recusandae",
                "temporibus": "ab",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "quis",
                TerminalPhoneNo: "veritatis",
                TerminalType: "deserunt",
            },
        },
        XAPIVersion: suhas.String("perferendis"),
        XClientID: "ipsam",
        XClientSecret: "repellendus",
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