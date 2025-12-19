# Payroc API Go SDK

The Payroc API Go SDK provides convenient access to the Payroc API from Go.

## Contents

- [Installation](#installation)
- [Usage](#usage)
  - [API Key](#api-key)
  - [PayrocClient](#payrocclient)
  - [Advanced Usage with Custom Environment](#advanced-usage-with-custom-environment)
- [Exception Handling](#exception-handling)
- [Logging](#logging)
- [Pagination](#pagination)
  - [Pagination Gotcha](#pagination-gotcha)
- [Polymorphism](#polymorphism)
  - [Understanding Discriminated Unions](#understanding-discriminated-unions)
  - [Working with Polymorphic Types](#working-with-polymorphic-types)
  - [Common Polymorphic Types](#common-polymorphic-types)
- [Request Parameters](#request-parameters)
- [Request Options](#request-options)
- [Advanced](#advanced)
  - [Response Headers](#response-headers)
  - [Retries](#retries)
  - [Timeouts](#timeouts)
  - [Explicit Null](#explicit-null)
- [Contributing](#contributing)
- [References](#references)

## Installation

```sh
go get github.com/payroc/payroc-sdk-go
```

## Reference

A full reference for this library is available [here](https://github.com/payroc/payroc-sdk-go/blob/HEAD/./reference.md).

## Usage

### API Key

You need to provide your API Key to the `PayrocClient` constructor. In this example we read it from an environment variable named `PAYROC_API_KEY`. In your own code you should consider security and compliance best practices, likely retrieving this value from a secure vault on demand.

### PayrocClient

Instantiate and use the client with the following:

```go
package example

import (
    "context"
    "os"
    
    payroc "github.com/payroc/payroc-sdk-go"
    "github.com/payroc/payroc-sdk-go/cardpayments"
    "github.com/payroc/payroc-sdk-go/client"
    "github.com/payroc/payroc-sdk-go/option"
)

func do() {
    apiKey := os.Getenv("PAYROC_API_KEY")
    if apiKey == "" {
        panic("Payroc API Key not found")
    }
    
    payrocClient := client.NewPayrocClient(
        option.WithApiKey(apiKey),
    )
    request := &cardpayments.PaymentRequest{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Channel: cardpayments.PaymentRequestChannelWeb,
        ProcessingTerminalId: "1234001",
        Operator: payroc.String(
            "Jane",
        ),
        Order: &payroc.PaymentOrderRequest{
            OrderId: payroc.String(
                "OrderRef6543",
            ),
            Description: payroc.String(
                "Large Pepperoni Pizza",
            ),
            Amount: payroc.Int64(
                4999,
            ),
            Currency: payroc.CurrencyUsd.Ptr(),
        },
        Customer: &payroc.Customer{
            FirstName: payroc.String(
                "Sarah",
            ),
            LastName: payroc.String(
                "Hopper",
            ),
            BillingAddress: &payroc.Address{
                Address1: "1 Example Ave.",
                Address2: payroc.String(
                    "Example Address Line 2",
                ),
                Address3: payroc.String(
                    "Example Address Line 3",
                ),
                City: "Chicago",
                State: "Illinois",
                Country: "US",
                PostalCode: "60056",
            },
            ShippingAddress: &payroc.Shipping{
                RecipientName: payroc.String(
                    "Sarah Hopper",
                ),
                Address: &payroc.Address{
                    Address1: "1 Example Ave.",
                    Address2: payroc.String(
                        "Example Address Line 2",
                    ),
                    Address3: payroc.String(
                        "Example Address Line 3",
                    ),
                    City: "Chicago",
                    State: "Illinois",
                    Country: "US",
                    PostalCode: "60056",
                },
            },
        },
        PaymentMethod: &cardpayments.PaymentRequestPaymentMethod{
            Card: &payroc.CardPayload{
                CardDetails: &payroc.CardPayloadCardDetails{
                    Raw: &payroc.RawCardDetails{
                        Device: &payroc.Device{
                            Model: payroc.DeviceModelBbposChp,
                            SerialNumber: "1850010868",
                        },
                        RawData: "A1B2C3D4E5F67890ABCD1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
                    },
                },
            },
        },
        CustomFields: []*payroc.CustomField{
            &payroc.CustomField{
                Name: "yourCustomField",
                Value: "abc123",
            },
        },
    }
    response, err := payrocClient.CardPayments.Payments.Create(
        context.TODO(),
        request,
    )
    if err != nil {
        // handle error
    }
    _ = response
}

### Advanced Usage with Custom Environment

If you wish to use the SDK against a custom URL, such as a mock API server, you can provide a custom base URL to the `PayrocClient` constructor:

```go
import (
    "github.com/payroc/payroc-sdk-go/client"
    "github.com/payroc/payroc-sdk-go/option"
)

func main() {
    client := client.NewPayrocClient(
        option.WithApiKey(apiKey),
        option.WithBaseURL("http://localhost:3000"),
    )
}

## Exception Handling

When the API returns a non-success status code (4xx or 5xx response), an error will be returned. Structured error types are compatible with the `errors.Is` and `errors.As` APIs:

```go
import (
    "errors"
    
    "github.com/payroc/payroc-sdk-go/core"
)

response, err := client.CardPayments.Payments.Create(ctx, request)
if err != nil {
    var apiError *core.APIError
    if errors.As(err, &apiError) {
        // Access error details
        fmt.Printf("Status Code: %d\n", apiError.StatusCode)
        fmt.Printf("Error: %v\n", apiError)
    }
    return err
}
```

## Logging

> [!WARNING]  
> Be careful when configuring your logging not to log the headers of outbound HTTP requests, lest you leak an API key or access token.

## Pagination

List endpoints are paginated. The SDK provides three approaches for iterating through pages:

### Approach 1: Manual Pagination

```go
import (
    "github.com/payroc/payroc-sdk-go/cardpayments"
)

pager, err := client.CardPayments.Payments.List(
    ctx,
    &cardpayments.ListPaymentsRequest{
        ProcessingTerminalId: "1234001",
    },
)
if err != nil {
    // handle error
}

// Access current page
for _, payment := range pager.Current().GetData() {
    // process payment
}

// Get next page
if pager.HasNextPage() {
    nextPage, err := pager.GetNextPage(ctx)
    if err != nil {
        // handle error
    }
    _ = nextPage
}
```

### Approach 2: Channel-Based Iteration (All Go Versions)

```go
pager, err := client.CardPayments.Payments.List(
    ctx,
    &cardpayments.ListPaymentsRequest{
        ProcessingTerminalId: "1234001",
    },
)
if err != nil {
    // handle error
}

for page := range pager.Iter(ctx) {
    for _, payment := range page.GetData() {
        // process payment
    }
}
```

### Approach 3: Range-Over-Func (Go 1.23+)

```go
pager, err := client.CardPayments.Payments.List(
    ctx,
    &cardpayments.ListPaymentsRequest{
        ProcessingTerminalId: "1234001",
    },
)
if err != nil {
    // handle error
}

for page := range pager.Seq(ctx) {
    for _, payment := range page.GetData() {
        // process payment
    }
}
```

### Pagination Gotcha

Beware of iterating the items on a single page and thinking that they are all there are. In the following example, there are only items from the first page, not all available items:

```go
pager, err := client.CardPayments.Payments.List(ctx, &cardpayments.ListPaymentsRequest{})
if err != nil {
    // handle error
}

// ⚠️ This only processes the FIRST page!
for _, payment := range pager.Current().GetData() {
    // This will miss items on subsequent pages
}
```

This might be helpful when you only want to process the first few results, but to iterate all items, use one of the iteration approaches shown above.

## Polymorphism

The Payroc API is heavily polymorphic, meaning many request and response types can take multiple forms. The SDK represents these polymorphic types as discriminated unions with type-safe handling through the visitor pattern.

### Understanding Discriminated Unions

A discriminated union is a type that can hold one of several possible variants, distinguished by a discriminator field. For example, `CardPayloadCardDetails` can be one of:
- `Raw` - Raw card data from a device
- `Icc` - ICC/chip card data
- `Keyed` - Manually keyed card data

### Working with Polymorphic Types

There are two main approaches to working with polymorphic types:

#### Approach 1: Direct Field Access

You can construct polymorphic types by setting the appropriate variant field:

```go
// Create a payment with raw card details
paymentMethod := &cardpayments.PaymentRequestPaymentMethod{
    Card: &payroc.CardPayload{
        CardDetails: &payroc.CardPayloadCardDetails{
            Raw: &payroc.RawCardDetails{
                Device: &payroc.Device{
                    Model: payroc.DeviceModelBbposChp,
                    SerialNumber: "1850010868",
                },
                RawData: "A1B2C3D4E5F67890...",
            },
        },
    },
}

// Or with keyed card details
paymentMethod := &cardpayments.PaymentRequestPaymentMethod{
    Card: &payroc.CardPayload{
        CardDetails: &payroc.CardPayloadCardDetails{
            Keyed: &payroc.KeyedCardDetails{
                KeyedData: &payroc.KeyedCardDetailsKeyedData{
                    PlainText: &payroc.PlainTextKeyedDataFormat{
                        CardNumber: "4111111111111111",
                        ExpiryMonth: "12",
                        ExpiryYear: "2025",
                        Cvv: payroc.String("123"),
                    },
                },
            },
        },
    },
}
```

When reading polymorphic response fields, check which variant is set:

```go
response, err := client.CardPayments.Payments.Get(ctx, paymentId)
if err != nil {
    // handle error
}

cardDetails := response.PaymentMethod.Card.CardDetails
if cardDetails.Raw != nil {
    fmt.Printf("Raw card data from device: %s\n", cardDetails.Raw.Device.SerialNumber)
} else if cardDetails.Icc != nil {
    fmt.Printf("ICC card data\n")
} else if cardDetails.Keyed != nil {
    fmt.Printf("Keyed card data\n")
}
```

#### Approach 2: Visitor Pattern (Type-Safe)

For more robust handling, use the visitor pattern. This ensures you handle all possible variants at compile time:

```go
type CardDetailsHandler struct {
    result string
}

func (h *CardDetailsHandler) VisitRaw(raw *payroc.RawCardDetails) error {
    h.result = fmt.Sprintf("Raw card from device: %s", raw.Device.SerialNumber)
    return nil
}

func (h *CardDetailsHandler) VisitIcc(icc *payroc.IccCardDetails) error {
    h.result = "ICC/Chip card"
    return nil
}

func (h *CardDetailsHandler) VisitKeyed(keyed *payroc.KeyedCardDetails) error {
    h.result = "Manually keyed card"
    return nil
}

// Use the visitor
handler := &CardDetailsHandler{}
err := cardDetails.Accept(handler)
if err != nil {
    // handle error
}
fmt.Println(handler.result)
```

The visitor pattern is particularly useful when:
- You need to ensure all variants are handled
- The handling logic is complex
- You want compile-time safety when new variants are added

### Common Polymorphic Types

The SDK includes many polymorphic types. Here are some commonly used ones:

- **`CardPayloadCardDetails`** - Card data entry method (Raw, ICC, Keyed)
- **`KeyedCardDetailsKeyedData`** - Keyed data encryption (FullyEncrypted, PartiallyEncrypted, PlainText)
- **`SecureTokenSource`** - Token source type (ACH, PAD, Card)
- **`PatchDocument`** - JSON Patch operations (Add, Remove, Replace, Move, Copy, Test)
- **`ContactMethod`** - Contact method type (Email, Phone, Mobile)

Each polymorphic type includes a `Type` field (the discriminator) that indicates which variant is active, though you typically don't need to set this manually—it's handled automatically during serialization.

## Request Parameters

Sometimes you need to filter results, for example, retrieving results from a given date. The SDK uses request objects to set query parameters.

Examples of setting different query parameters:

```go
import (
    "time"
    
    "github.com/payroc/payroc-sdk-go/cardpayments"
)

// Filter by date range
request := &cardpayments.ListPaymentsRequest{
    ProcessingTerminalId: "1234001",
    DateFrom: payroc.Time(time.Date(2024, 7, 1, 15, 30, 0, 0, time.UTC)),
    DateTo: payroc.Time(time.Date(2024, 7, 3, 15, 30, 0, 0, time.UTC)),
}

// Pagination with cursor
request := &cardpayments.ListPaymentsRequest{
    ProcessingTerminalId: "1234001",
    After: payroc.String("8516"),
}

request := &cardpayments.ListPaymentsRequest{
    ProcessingTerminalId: "1234001",
    Before: payroc.String("2571"),
}
```

Inspect the code definition of your particular `...Request` type in your IDE to see what properties can be used for filtering.

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes configuring
authorization tokens, or providing your own instrumented `*http.Client`.

These request options can either be
specified on the client so that they're applied on every request, or for an individual request, like so:

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

```go
import (
    "net/http"
    "time"
    
    "github.com/payroc/payroc-sdk-go/client"
    "github.com/payroc/payroc-sdk-go/option"
)

// Specify default options applied on every request.
client := client.NewPayrocClient(
    option.WithApiKey("<YOUR_API_KEY>"),
    option.WithHTTPClient(
        &http.Client{
            Timeout: 5 * time.Second,
        },
    ),
)

// Specify options for an individual request.
response, err := client.CardPayments.Payments.Create(
    ctx,
    request,
    option.WithApiKey("<YOUR_API_KEY>"),
)
```

## Advanced

### Response Headers

You can access the raw HTTP response data by using the `WithRawResponse` field on the client. This is useful
when you need to examine the response headers received from the API call. (When the endpoint is paginated,
the raw HTTP response data will be included automatically in the Page response object.)

```go
response, err := client.CardPayments.Payments.WithRawResponse.Create(ctx, request)
if err != nil {
    return err
}
fmt.Printf("Got response headers: %v\n", response.Header)
fmt.Printf("Got status code: %d\n", response.StatusCode)
```

### Retries

The SDK is instrumented with automatic retries with exponential backoff. A request will be retried as long
as the request is deemed retryable and the number of retry attempts has not grown larger than the configured
retry limit (default: 2).

A request is deemed retryable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

If the `Retry-After` header is present in the response, the SDK will prioritize respecting its value exactly
over the default exponential backoff.

Use the `option.WithMaxAttempts` option to configure this behavior for the entire client or an individual request:

```go
client := client.NewPayrocClient(
    option.WithApiKey(apiKey),
    option.WithMaxAttempts(1),
)

response, err := client.CardPayments.Payments.Create(
    ctx,
    request,
    option.WithMaxAttempts(1),
)
```

### Timeouts

Setting a timeout for each individual request is as simple as using the standard context library. Setting a one second timeout for an individual API call looks like the following:

```go
import (
    "context"
    "time"
)

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

response, err := client.CardPayments.Payments.Create(ctx, request)
```

### Explicit Null

If you want to send the explicit `null` JSON value through an optional parameter, you can use the setters\
that come with every object. Calling a setter method for a property will flip a bit in the `explicitFields`
bitfield for that setter's object; during serialization, any property with a flipped bit will have its
omittable status stripped, so zero or `nil` values will be sent explicitly rather than omitted altogether:

```go
type ExampleRequest struct {
    // An optional string parameter.
    Name *string `json:"name,omitempty" url:"-"`

    // Private bitmask of fields set to an explicit value and therefore not to be omitted
    explicitFields *big.Int `json:"-" url:"-"`
}

request := &ExampleRequest{}
request.SetName(nil)

response, err := client.CardPayments.Payments.Create(ctx, request)
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!

For details on setting up your development environment, running tests, and code quality standards, please see [CONTRIBUTING.md](./CONTRIBUTING.md).

## References

The Payroc API SDK is generated via [Fern](https://www.buildwithfern.com/).

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=https%3A%2F%2Fgithub.com%2Fpayroc%2Fpayroc-sdk-go)