<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/suhastest"
	"github.com/speakeasy-sdks/suhastest/pkg/models/operations"
	"github.com/speakeasy-sdks/suhastest/pkg/models/shared"
	"log"
)

func main() {
	s := suhastest.New()

	ctx := context.Background()
	res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
		CreateOrderBackendRequest: &shared.CreateOrderBackendRequest{
			CustomerDetails: shared.CustomerDetails{
				CustomerID:    "string",
				CustomerPhone: "string",
			},
			OrderAmount:     10.15,
			OrderCurrency:   "INR",
			OrderExpiryTime: suhastest.String("2021-07-29T00:00:00Z"),
			OrderMeta:       &shared.OrderMeta{},
			OrderNote:       suhastest.String("Test order"),
			OrderSplits: []shared.VendorSplit{
				shared.VendorSplit{},
			},
			OrderTags: map[string]string{
				"key": "string",
			},
			Terminal: &shared.TerminalDetails{
				TerminalID:      "string",
				TerminalPhoneNo: "string",
				TerminalType:    "string",
			},
		},
		XClientID:     "string",
		XClientSecret: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.OrdersEntity != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->