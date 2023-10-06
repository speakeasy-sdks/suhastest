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
                CustomerBankAccountNumber: suhastest.String("North double"),
                CustomerBankCode: suhastest.String("spherical woman burdensome"),
                CustomerBankIfsc: suhastest.String("interfaces Smart"),
                CustomerEmail: suhastest.String("Doyle brown toast"),
                CustomerID: "Bedfordshire",
                CustomerPhone: "Mohr North",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: suhastest.String("2021-07-29T00:00:00Z"),
            OrderID: suhastest.String("deploy South"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: suhastest.String("Road male Berkshire"),
                PaymentMethods: suhastest.String("parsing female middleware"),
                ReturnURL: suhastest.String("Bedfordshire navigating"),
            },
            OrderNote: suhastest.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: suhastest.Float64(5942.72),
                    Percentage: suhastest.Float64(3302.96),
                    VendorID: suhastest.String("dearly remount"),
                },
            },
            OrderTags: map[string]string{
                "expedita": "South",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "Southwest",
                TerminalPhoneNo: "violet Chips Porsche",
                TerminalType: "mobile",
            },
        },
        XAPIVersion: suhastest.String("ROI bypassing vero"),
        XClientID: "Solutions Ferrari Accountability",
        XClientSecret: "Folk ampere",
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