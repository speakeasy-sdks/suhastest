# Orders
(*Orders*)

### Available Operations

* [CreateOrder](#createorder) - Create Order

## CreateOrder

Use this API to create orders with Cashfree from your backend and get the payment link. To use this API S2S flag needs to be enabled from the backend. In case you want to use the cards payment option the PCI DSS flag is required, for more information email us at "care@cashfree.com".

### Example Usage

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
                CustomerBankAccountNumber: suhastest.String("quod"),
                CustomerBankCode: suhastest.String("quod"),
                CustomerBankIfsc: suhastest.String("esse"),
                CustomerEmail: suhastest.String("totam"),
                CustomerID: "porro",
                CustomerPhone: "dolorum",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: suhastest.String("2021-07-29T00:00:00Z"),
            OrderID: suhastest.String("dicta"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: suhastest.String("nam"),
                PaymentMethods: suhastest.String("officia"),
                ReturnURL: suhastest.String("occaecati"),
            },
            OrderNote: suhastest.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: suhastest.Float64(1433.53),
                    Percentage: suhastest.Float64(5373.73),
                    VendorID: suhastest.String("hic"),
                },
            },
            OrderTags: map[string]string{
                "optio": "totam",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "beatae",
                TerminalPhoneNo: "commodi",
                TerminalType: "molestiae",
            },
        },
        XAPIVersion: suhastest.String("modi"),
        XClientID: "qui",
        XClientSecret: "impedit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrdersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateOrderRequest](../../models/operations/createorderrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateOrderResponse](../../models/operations/createorderresponse.md), error**

