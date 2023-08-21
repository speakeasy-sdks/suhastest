# RateLimitError

Rate Limit Error


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Code`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Message`                                                               | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Type`                                                                  | [*shared.RateLimitErrorType](../../models/shared/ratelimiterrortype.md) | :heavy_minus_sign:                                                      | rate_limit_error                                                        |