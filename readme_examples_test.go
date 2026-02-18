package payroc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	payroc "github.com/payroc/payroc-sdk-go"
	"github.com/payroc/payroc-sdk-go/cardpayments"
	"github.com/payroc/payroc-sdk-go/client"
	"github.com/payroc/payroc-sdk-go/core"
	"github.com/payroc/payroc-sdk-go/option"
)

func TestReadmeExample_BasicUsage(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	apiKey := os.Getenv("PAYROC_API_KEY")
	if apiKey == "" {
		t.Skip("Payroc API Key not found")
	}

	payrocClient := client.NewPayrocClient(
		option.WithApiKey(apiKey),
	)
	request := &cardpayments.PaymentRequest{
		IdempotencyKey:       "8e03978e-40d5-43e8-bc93-6894a57f9324",
		Channel:              cardpayments.PaymentRequestChannelWeb,
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
				City:       "Chicago",
				State:      "Illinois",
				Country:    "US",
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
					City:       "Chicago",
					State:      "Illinois",
					Country:    "US",
					PostalCode: "60056",
				},
			},
		},
		PaymentMethod: &cardpayments.PaymentRequestPaymentMethod{
			Card: &payroc.CardPayload{
				CardDetails: &payroc.CardPayloadCardDetails{
					Raw: &payroc.RawCardDetails{
						Device: &payroc.Device{
							Model:        payroc.DeviceModelBbposChp,
							SerialNumber: "1850010868",
						},
						RawData: "A1B2C3D4E5F67890ABCD1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					},
				},
			},
		},
		CustomFields: []*payroc.CustomField{
			&payroc.CustomField{
				Name:  "yourCustomField",
				Value: "abc123",
			},
		},
	}
	response, err := payrocClient.CardPayments.Payments.Create(
		context.TODO(),
		request,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_ = response
}

func TestReadmeExample_CustomEnvironment(t *testing.T) {
	t.Skip("Skipping compilation test")

	apiKey := "test-key"
	c := client.NewPayrocClient(
		option.WithApiKey(apiKey),
		option.WithBaseURL("http://localhost:3000"),
	)
	_ = c
}

func TestReadmeExample_ExceptionHandling(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	ctx := context.Background()
	c := client.NewPayrocClient(option.WithApiKey("test-key"))
	request := &cardpayments.PaymentRequest{}

	response, err := c.CardPayments.Payments.Create(ctx, request)
	if err != nil {
		var apiError *core.APIError
		if errors.As(err, &apiError) {
			fmt.Printf("Status Code: %d\n", apiError.StatusCode)
			fmt.Printf("Error: %v\n", apiError)
		}
		_ = response
		return
	}
}

func TestReadmeExample_PaginationManual(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	ctx := context.Background()
	c := client.NewPayrocClient(option.WithApiKey("test-key"))

	pager, err := c.CardPayments.Payments.List(
		ctx,
		&cardpayments.ListPaymentsRequest{
			ProcessingTerminalId: payroc.String("1234001"),
		},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, payment := range pager.Current().GetData() {
		_ = payment
	}

	if pager.HasNextPage() {
		nextPage, err := pager.GetNextPage(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_ = nextPage
	}
}

func TestReadmeExample_PaginationChannelBased(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	ctx := context.Background()
	c := client.NewPayrocClient(option.WithApiKey("test-key"))

	pager, err := c.CardPayments.Payments.List(
		ctx,
		&cardpayments.ListPaymentsRequest{
			ProcessingTerminalId: payroc.String("1234001"),
		},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for page := range pager.Iter(ctx) {
		for _, payment := range page.GetData() {
			_ = payment
		}
	}
}

func TestReadmeExample_PolymorphicDirectAccess(t *testing.T) {
	t.Skip("Skipping compilation test")

	paymentMethod := &cardpayments.PaymentRequestPaymentMethod{
		Card: &payroc.CardPayload{
			CardDetails: &payroc.CardPayloadCardDetails{
				Raw: &payroc.RawCardDetails{
					Device: &payroc.Device{
						Model:        payroc.DeviceModelBbposChp,
						SerialNumber: "1850010868",
					},
					RawData: "A1B2C3D4E5F67890...",
				},
			},
		},
	}
	_ = paymentMethod

	paymentMethod2 := &cardpayments.PaymentRequestPaymentMethod{
		Card: &payroc.CardPayload{
			CardDetails: &payroc.CardPayloadCardDetails{
				Keyed: &payroc.KeyedCardDetails{
					KeyedData: &payroc.KeyedCardDetailsKeyedData{
						PlainText: &payroc.PlainTextKeyedDataFormat{
							CardNumber: "4111111111111111",
							ExpiryDate: payroc.String("1230"),
							Cvv:        payroc.String("123"),
						},
					},
				},
			},
		},
	}
	_ = paymentMethod2
}

func TestReadmeExample_PolymorphicReading(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	ctx := context.Background()
	c := client.NewPayrocClient(option.WithApiKey("test-key"))
	paymentId := "test-payment-id"

	response, err := c.CardPayments.Payments.Retrieve(ctx, &cardpayments.RetrievePaymentsRequest{
		PaymentId: paymentId,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Note: RetrievedPayment.Card is a RetrievedCard which has flattened fields,
	// not the polymorphic CardDetails structure shown in README.
	// The README example is conceptual for request structures.
	_ = response
}

func TestReadmeExample_PolymorphicVisitor(t *testing.T) {
	t.Skip("Skipping compilation test")

	type CardDetailsHandler struct {
		result string
	}

	handler := &CardDetailsHandler{}
	
	// Note: The visitor pattern example in README is conceptual.
	// In actual usage, you would implement the CardPayloadCardDetailsVisitor interface
	// and call cardDetails.Accept(handler) on a real CardPayloadCardDetails instance.
	_ = handler
}

func TestReadmeExample_RequestParameters(t *testing.T) {
	t.Skip("Skipping compilation test")

	request1 := &cardpayments.ListPaymentsRequest{
		ProcessingTerminalId: payroc.String("1234001"),
		DateFrom:             payroc.Time(time.Date(2024, 7, 1, 15, 30, 0, 0, time.UTC)),
		DateTo:               payroc.Time(time.Date(2024, 7, 3, 15, 30, 0, 0, time.UTC)),
	}
	_ = request1

	request2 := &cardpayments.ListPaymentsRequest{
		ProcessingTerminalId: payroc.String("1234001"),
		After:                payroc.String("8516"),
	}
	_ = request2

	request3 := &cardpayments.ListPaymentsRequest{
		ProcessingTerminalId: payroc.String("1234001"),
		Before:               payroc.String("2571"),
	}
	_ = request3
}

func TestReadmeExample_RequestOptions(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	ctx := context.Background()

	c := client.NewPayrocClient(
		option.WithApiKey("<YOUR_API_KEY>"),
		option.WithHTTPClient(
			&http.Client{
				Timeout: 5 * time.Second,
			},
		),
	)

	request := &cardpayments.PaymentRequest{}
	response, err := c.CardPayments.Payments.Create(
		ctx,
		request,
		option.WithApiKey("<YOUR_API_KEY>"),
	)
	_ = response
	_ = err
}

func TestReadmeExample_ResponseHeaders(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	ctx := context.Background()
	c := client.NewPayrocClient(option.WithApiKey("test-key"))
	request := &cardpayments.PaymentRequest{}

	response, err := c.CardPayments.Payments.WithRawResponse.Create(ctx, request)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fmt.Printf("Got response headers: %v\n", response.Header)
	fmt.Printf("Got status code: %d\n", response.StatusCode)
}

func TestReadmeExample_Retries(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	ctx := context.Background()
	apiKey := "test-key"

	c := client.NewPayrocClient(
		option.WithApiKey(apiKey),
		option.WithMaxAttempts(1),
	)

	request := &cardpayments.PaymentRequest{}
	response, err := c.CardPayments.Payments.Create(
		ctx,
		request,
		option.WithMaxAttempts(1),
	)
	_ = response
	_ = err
}

func TestReadmeExample_Timeouts(t *testing.T) {
	t.Skip("Skipping compilation test - requires valid API key")

	c := client.NewPayrocClient(option.WithApiKey("test-key"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &cardpayments.PaymentRequest{}
	response, err := c.CardPayments.Payments.Create(ctx, request)
	_ = response
	_ = err
}
