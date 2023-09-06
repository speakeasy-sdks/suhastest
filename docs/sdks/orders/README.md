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
                CustomerBankAccountNumber: suhas.String("placeat"),
                CustomerBankCode: suhas.String("voluptatum"),
                CustomerBankIfsc: suhas.String("iusto"),
                CustomerEmail: suhas.String("excepturi"),
                CustomerID: "nisi",
                CustomerPhone: "recusandae",
            },
            OrderAmount: 10.15,
            OrderCurrency: "INR",
            OrderExpiryTime: suhas.String("2021-07-29T00:00:00Z"),
            OrderID: suhas.String("temporibus"),
            OrderMeta: &shared.OrderMeta{
                NotifyURL: suhas.String("ab"),
                PaymentMethods: suhas.String("quis"),
                ReturnURL: suhas.String("veritatis"),
            },
            OrderNote: suhas.String("Test order"),
            OrderSplits: []shared.VendorSplit{
                shared.VendorSplit{
                    Amount: suhas.Float64(6481.72),
                    Percentage: suhas.Float64(202.18),
                    VendorID: suhas.String("ipsam"),
                },
            },
            OrderTags: map[string]string{
                "repellendus": "sapiente",
            },
            Terminal: &shared.TerminalDetails{
                TerminalID: "quo",
                TerminalPhoneNo: "odit",
                TerminalType: "at",
            },
        },
        XAPIVersion: suhas.String("at"),
        XClientID: "maiores",
        XClientSecret: "molestiae",
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
