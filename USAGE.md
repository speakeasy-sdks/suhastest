<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/suhastest"
	"github.com/speakeasy-sdks/suhastest/pkg/models/operations"
	"log"
)

func main() {
	s := suhastest.New()

	ctx := context.Background()
	res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
		XClientID:     "<value>",
		XClientSecret: "<value>",
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