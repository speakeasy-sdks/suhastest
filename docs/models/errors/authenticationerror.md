# AuthenticationError

Authentication Error


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `RawResponse`                                                                     | [*http.Response](https://pkg.go.dev/net/http#Response)                            | :heavy_minus_sign:                                                                | N/A                                                                               |
| `Code`                                                                            | **string*                                                                         | :heavy_minus_sign:                                                                | N/A                                                                               |
| `Message`                                                                         | **string*                                                                         | :heavy_minus_sign:                                                                | N/A                                                                               |
| `Type`                                                                            | [*shared.AuthenticationErrorType](../../models/shared/authenticationerrortype.md) | :heavy_minus_sign:                                                                | authentication_error                                                              |