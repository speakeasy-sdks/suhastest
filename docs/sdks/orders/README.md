# Orders

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
    s := suhas.New()

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
        CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
            CustomerDetails: shared.CustomerDetails{
                CustomerBankAccountNumber: suhas.String("sapiente"),
                CustomerBankCode: suhas.String("quo"),
                CustomerBankIfsc: suhas.String("odit"),
                CustomerEmail: suhas.String("at"),
                CustomerID: "at",
                CustomerPhone: "maiores",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: suhas.String("2021-07-29T00:00:00Z"),
            OrderID: suhas.String("molestiae"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: suhas.String("quod"),
                PaymentMethods: suhas.String("quod"),
                ReturnURL: suhas.String("esse"),
            },
            OrderNote: suhas.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: suhas.Float64(7805.29),
                    Percentage: suhas.Float64(6788.8),
                    VendorID: suhas.String("dicta"),
                },
                shared.VendorSplit{
                    Amount: suhas.Float64(7206.33),
                    Percentage: suhas.Float64(6399.21),
                    VendorID: suhas.String("occaecati"),
                },
                shared.VendorSplit{
                    Amount: suhas.Float64(1433.53),
                    Percentage: suhas.Float64(5373.73),
                    VendorID: suhas.String("hic"),
                },
            },
            OrderTags: map[string]string{
                "totam": "beatae",
                "commodi": "molestiae",
                "modi": "qui",
                "impedit": "cum",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "esse",
                TerminalPhoneNo: "ipsum",
                TerminalType: "excepturi",
            },
        },
        XAPIVersion: suhas.String("aspernatur"),
        XClientID: "perferendis",
        XClientSecret: "ad",
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

