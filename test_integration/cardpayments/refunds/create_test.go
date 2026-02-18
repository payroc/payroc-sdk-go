//go:build smoke
// +build smoke

package refunds_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	payroc "github.com/payroc/payroc-sdk-go"
	"github.com/payroc/payroc-sdk-go/cardpayments"
	"github.com/payroc/payroc-sdk-go/client"
	"github.com/payroc/payroc-sdk-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCreateRefund tests creating a refund against the Payroc UAT environment.
// This test mirrors the functionality of CreateTests.cs from papi-sdk-dotnet.
func TestCreateRefund(t *testing.T) {
	// Load API key from environment variable (same pattern as .NET SDK)
	apiKey := os.Getenv("PAYROC_API_KEY_PAYMENTS")
	if apiKey == "" {
		t.Skip("PAYROC_API_KEY_PAYMENTS environment variable not set")
	}

	// Load processing terminal ID from environment
	processingTerminalId := os.Getenv("TERMINAL_ID_AVS")

	// Prev impl...
	//payrocClient := client.NewPayrocClient(
	//	option.WithApiKey(apiKey),
	//	option.WithBaseURL(payroc.Environments.Uat.Api),
	//)

	// Create client configured for UAT environment
	payrocClient := client.NewPayrocClient(
		option.WithApiKey(apiKey),
		option.WithEnvironment(payroc.Environments.Uat),
	)

	ctx := context.Background()

	// Step 1: Create a payment first (required to have a payment to refund)
	t.Log("Creating a payment to refund...")
	paymentRequest := &cardpayments.PaymentRequest{
		IdempotencyKey:       generateIdempotencyKey(),
		Channel:              cardpayments.PaymentRequestChannelWeb,
		ProcessingTerminalId: processingTerminalId,
		Operator:             payroc.String("IntegrationTest"),
		Order: &payroc.PaymentOrderRequest{
			OrderId:     payroc.String("TEST-" + time.Now().Format("20060102150405")),
			Description: payroc.String("Integration test payment for refund"),
			Amount:      payroc.Int64(1000), // $10.00
			Currency:    payroc.CurrencyUsd.Ptr(),
		},
		Customer: &payroc.Customer{
			FirstName: payroc.String("Test"),
			LastName:  payroc.String("User"),
			BillingAddress: &payroc.Address{
				Address1:   "123 Test St",
				City:       "Test City",
				State:      "CA",
				Country:    "US",
				PostalCode: "90210",
			},
		},
		PaymentMethod: &cardpayments.PaymentRequestPaymentMethod{
			Card: &payroc.CardPayload{
				CardDetails: &payroc.CardPayloadCardDetails{
					Keyed: &payroc.KeyedCardDetails{
						KeyedData: &payroc.KeyedCardDetailsKeyedData{
							PlainText: &payroc.PlainTextKeyedDataFormat{
								CardNumber: "4111111111111111", // Test card
								ExpiryDate: payroc.String("1230"),
								Cvv:        payroc.String("123"),
							},
						},
					},
				},
			},
		},
	}

	payment, err := payrocClient.CardPayments.Payments.Create(ctx, paymentRequest)
	require.NoError(t, err, "Failed to create payment")

	require.NotNil(t, payment, "Payment should not be nil")
	require.NotEmpty(t, payment.PaymentId, "Payment ID should not be empty")
	t.Logf("Created payment with ID: %s", payment.PaymentId)

	// Step 2: Create a refund for the payment
	t.Log("Creating refund...")
	refundRequest := &cardpayments.ReferencedRefund{
		IdempotencyKey: generateIdempotencyKey(),
		PaymentId:      payment.PaymentId,
		Operator:       payroc.String("IntegrationTest"),
		Amount:         500, // Partial refund of $5.00
		Description:    "Integration test refund",
	}

	refund, err := payrocClient.CardPayments.Refunds.CreateReferencedRefund(ctx, refundRequest)
	require.NoError(t, err, "Failed to create refund")

	// Assertions
	require.NotNil(t, refund, "Refund should not be nil")
	require.NotEmpty(t, refund.PaymentId, "Payment ID should not be empty")
	assert.Equal(t, payment.PaymentId, refund.PaymentId, "Refund should be linked to the payment")

	t.Logf("Successfully created refund for payment ID: %s", refund.PaymentId)
	t.Logf("Refund status: %s", refund.TransactionResult.Status)
}

// generateIdempotencyKey generates a unique idempotency key for requests
func generateIdempotencyKey() string {
	return uuid.New().String()
}
