# APIError

API related Errors


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `RawResponse`                                               | [*http.Response](https://pkg.go.dev/net/http#Response)      | :heavy_minus_sign:                                          | N/A                                                         |
| `Code`                                                      | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `Message`                                                   | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `Type`                                                      | [*shared.APIErrorType](../../models/shared/apierrortype.md) | :heavy_minus_sign:                                          | api_error                                                   |