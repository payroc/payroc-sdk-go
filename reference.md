# Reference
## Payment links
<details><summary><code>client.PaymentLinks.List(ProcessingTerminalId) -> *core.PayrocPager[*payroc.PaymentLinkPaginatedList, []*payroc.PaymentLinkPaginatedListDataItem, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of payment links linked to a processing terminal.  

**Note:** If you want to view the details of a specific payment link and you have its paymentLinkId, use our [Retrieve Payment Link](https://docs.payroc.com/api/schema/payment-links/retrieve) method.  

Use query parameters to filter the list of results that we return, for example, to search for only active links or multi-use links.  

Our gateway returns the following information about each payment link in the list:  
- **type** - Indicates whether the link can be used only once or if it can be used multiple times.  
- **authType** - Indicates whether the transaction is a sale or a pre-authorization.  
- **paymentMethods** - Indicates the payment method that the merchant accepts.  
- **charge** - Indicates whether the merchant or the customer enters the amount for the transaction.  
- **status** - Indicates if the payment link is active.  

For each payment link, we also return a paymentLinkId, which you can use for follow-on actions. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroc.ListPaymentLinksRequest{
        ProcessingTerminalId: "1234001",
        MerchantReference: payroc.String(
            "LinkRef6543",
        ),
        LinkType: payroc.ListPaymentLinksRequestLinkTypeMultiUse.Ptr(),
        ChargeType: payroc.ListPaymentLinksRequestChargeTypePreset.Ptr(),
        Status: payroc.ListPaymentLinksRequestStatusActive.Ptr(),
        RecipientName: payroc.String(
            "Sarah Hazel Hopper",
        ),
        RecipientEmail: payroc.String(
            "sarah.hopper@example.com",
        ),
        CreatedOn: payroc.Time(
            payroc.MustParseDate(
                "2024-07-02",
            ),
        ),
        ExpiresOn: payroc.Time(
            payroc.MustParseDate(
                "2024-08-02",
            ),
        ),
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.PaymentLinks.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**merchantReference:** `*string` — Filter results by the unique identifier that the merchant assigned to the payment link.
    
</dd>
</dl>

<dl>
<dd>

**linkType:** `*payroc.ListPaymentLinksRequestLinkType` — Filter results by the type of link. Send a value of <code>singleUse</code> or <code>multiUse</code>.
    
</dd>
</dl>

<dl>
<dd>

**chargeType:** `*payroc.ListPaymentLinksRequestChargeType` 

Filter results by the user that entered the amount. Send one of the following values:
- <code>prompt</code> - Customer entered the amount.
- <code>preset</code> - Merchant entered the amount.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*payroc.ListPaymentLinksRequestStatus` — Filter results by the status of the payment link. Send a value of <code>active</code>, <code>completed</code>,<code>deactived</code>, or <code>expired</code>.
    
</dd>
</dl>

<dl>
<dd>

**recipientName:** `*string` — Filter results by the customer's name.
    
</dd>
</dl>

<dl>
<dd>

**recipientEmail:** `*string` — Filter results by the customer's email address.
    
</dd>
</dl>

<dl>
<dd>

**createdOn:** `*time.Time` — Filter results by the link's created date. Send a value in **YYYY-MM-DD** format.
    
</dd>
</dl>

<dl>
<dd>

**expiresOn:** `*time.Time` — Filter results by the link's expiry date. Send a value in **YYYY-MM-DD** format.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentLinks.Create(ProcessingTerminalId, request) -> *payroc.CreatePaymentLinksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a payment link that a customer can use to make a payment for goods or services.  

The request includes the following settings:
- **type** - Indicates whether the link can be used only once or if it can be used multiple times.
- **authType** - Indicates whether the transaction is a sale or a pre-authorization.
- **paymentMethod** - Indicates the payment methods that the merchant accepts.
- **charge** - Indicates whether the merchant or the customer enters the amount for the transaction.  

If your request is successful, our gateway returns a paymentLinkId, which you can use to perform follow-on actions.  

**Note:** To share the payment link with a customer, use our [Share Payment Link](https://docs.payroc.com/api/schema/payment-links/sharing-events/share) method.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroc.CreatePaymentLinksRequest{
        ProcessingTerminalId: "1234001",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.CreatePaymentLinksRequestBody{
            MultiUse: &payroc.MultiUsePaymentLink{
                MerchantReference: "LinkRef6543",
                Order: &payroc.MultiUsePaymentLinkOrder{
                    Charge: &payroc.MultiUsePaymentLinkOrderCharge{
                        Prompt: &payroc.PromptPaymentLinkCharge{
                            Currency: payroc.CurrencyAed,
                        },
                    },
                },
                AuthType: payroc.MultiUsePaymentLinkAuthTypeSale,
                PaymentMethods: []payroc.MultiUsePaymentLinkPaymentMethodsItem{
                    payroc.MultiUsePaymentLinkPaymentMethodsItemCard,
                },
            },
        },
    }
client.PaymentLinks.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.CreatePaymentLinksRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentLinks.Retrieve(PaymentLinkId) -> *payroc.RetrievePaymentLinksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a payment link.  

To retrieve a payment link, you need its paymentLinkId. Our gateway returned the paymentLinkId in the response of the [Create Payment Link](https://docs.payroc.com/api/schema/payment-links/create) method.  

**Note:** If you don't have the paymentLinkId, use our [List Payment Links](https://docs.payroc.com/api/schema/payment-links/list) method to search for the payment link.  

Our gateway returns the following information about the payment link:  
- **type** - Indicates whether the link can be used only once or if it can be used multiple times.  
- **authType** - Indicates whether the transaction is a sale or a pre-authorization.  
- **paymentMethods** - Indicates the payment method that the merchant accepts.  
- **charge** - Indicates whether the merchant or the customer enters the amount for the transaction.
- **status** - Indicates if the payment link is active.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroc.RetrievePaymentLinksRequest{
        PaymentLinkId: "JZURRJBUPS",
    }
client.PaymentLinks.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentLinkId:** `string` — Unique identifier that we assigned to the payment link.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentLinks.PartiallyUpdate(PaymentLinkId, request) -> *payroc.PartiallyUpdatePaymentLinksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to partially update a payment link. Structure your request to follow the [RFC 6902](https://datatracker.ietf.org/doc/html/rfc6902) standard.

To update a payment link, you need its paymentLinkId, which we sent you in the response of the [Create Payment Link](https://docs.payroc.com/api/schema/payment-links/create) method.  

**Note:** If you don't have the paymentLinkId, use our [List Payment Links](https://docs.payroc.com/api/schema/payment-links/list) method to search for the payment link.  

You can update the following properties of a multi-use link:
- **expiresOn parameter** - Expiration date of the link.
- **customLabels object** - Label for the payment button.
- **credentialOnFile object** - Settings for saving the customer's payment details.

You can update the following properties of a single-use link:
- **expiresOn parameter** - Expiration date of the link.
- **authType parameter** - Transaction type of the payment link.
- **amount parameter** - Total amount of the transaction.
- **currency parameter** - Currency of the transaction.
- **description parameter** - Brief description of the transaction.
- **customLabels object** - Label for the payment button.
- **credentialOnFile object** - Settings for saving the customer's payment details.

**Note:** When a merchant updates a single-use link, we update the payment URL and HTML code in the assets object. The customer can't use the original link to make a payment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroc.PartiallyUpdatePaymentLinksRequest{
        PaymentLinkId: "JZURRJBUPS",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: []*payroc.PatchDocument{
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
        },
    }
client.PaymentLinks.PartiallyUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentLinkId:** `string` — Unique identifier that we assigned to the payment link.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `payroc.PatchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentLinks.Deactivate(PaymentLinkId) -> *payroc.DeactivatePaymentLinksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to deactivate a payment link.  

To deactivate a payment link, you need its paymentLinkId. Our gateway returned the paymentLinkId in the response of the [Create Payment Link](https://docs.payroc.com/api/schema/payment-links/create) method.  

**Note:** If you don't have the paymentLinkId, use our [List Payment Links](https://docs.payroc.com/api/schema/payment-links/list) method to search for the payment link.  

If your request is successful, our gateway deactivates the payment link. The customer can't use the link to make a payment, and you can't reactivate the payment link.    
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroc.DeactivatePaymentLinksRequest{
        PaymentLinkId: "JZURRJBUPS",
    }
client.PaymentLinks.Deactivate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentLinkId:** `string` — Unique identifier that we assigned to the payment link.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hosted Fields
<details><summary><code>client.HostedFields.Create(ProcessingTerminalId, request) -> *payroc.HostedFieldsCreateSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a Hosted Fields session token. You need to generate a new session token each time you load Hosted Fields on a webpage.  

In your request, you need to indicate whether the merchant is using Hosted Fields to run a sale, save payment details, or update saved payment details.  

In the response, our gateway returns the session token and the time that it expires. You need the session token when you configure the JavaScript for Hosted Fields.  

For more information about adding Hosted Fields to a webpage, go to [Hosted Fields](https://docs.payroc.com/guides/integrate/hosted-fields). 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroc.HostedFieldsCreateSessionRequest{
        ProcessingTerminalId: "1234001",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        LibVersion: "1.1.0.123456",
        Scenario: payroc.HostedFieldsCreateSessionRequestScenarioPayment,
    }
client.HostedFields.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**libVersion:** `string` 

Version of the Hosted Fields JavaScript library that you are using.  

The current production version is `1.6.0.172441`.
    
</dd>
</dl>

<dl>
<dd>

**scenario:** `*payroc.HostedFieldsCreateSessionRequestScenario` 

Indicates if a merchant wants to take a payment or tokenize a customer's payment details:  

- `payment` - The merchant wants to run a sale or run a sale and tokenize in the same transaction.  
- `tokenization` - The merchant wants to save the customer's payment details to take a payment later or to update a customer's payment details that they've already saved.  
    
</dd>
</dl>

<dl>
<dd>

**secureTokenId:** `*string` 

Unique identifier that represents a customer's payment details.  

If a merchant wants to update a customer's payment details that are linked to a secure token, include the secureTokenId in your request.  
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ApplePaySessions
<details><summary><code>client.ApplePaySessions.Create(ProcessingTerminalId, request) -> *payroc.ApplePayResponseSession</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to start an Apple Pay session for your merchant.  

In the response, we return the startSessionObject that you send to Apple when you retrieve the cardholder's encrypted payment details.  

**Note:** For more information about how to integrate with Apple Pay, go to [Apple Pay](https://docs.payroc.com/guides/integrate/apple-pay).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroc.ApplePaySessions{
        ProcessingTerminalId: "1234001",
        AppleDomainId: "DUHDZJHGYY",
        AppleValidationUrl: "https://apple-pay-gateway.apple.com/paymentservices/startSession",
    }
client.ApplePaySessions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**appleDomainId:** `string` — Unique appleDomainId of the merchant's domain that we assigned when you added their domain to our Self-Care Portal.
    
</dd>
</dl>

<dl>
<dd>

**appleValidationUrl:** `string` — Validation URL from the Apple Pay JS API.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Auth
<details><summary><code>client.Auth.RetrieveToken(request) -> *payroc.GetTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Obtain an access token using client credentials
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroc.RetrieveTokenAuthRequest{
        ApiKey: "x-api-key",
        ClientId: "client_id",
        ClientSecret: "client_secret",
    }
client.Auth.RetrieveToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKey:** `string` — The API key of the application
    
</dd>
</dl>

<dl>
<dd>

**clientId:** `string` — The client ID of the application
    
</dd>
</dl>

<dl>
<dd>

**clientSecret:** `string` — The client secret of the application
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BankTransferPayments Payments
<details><summary><code>client.BankTransferPayments.Payments.List() -> *core.PayrocPager[*payroc.BankTransferPaymentPaginatedList, []*payroc.BankTransferPayment, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of payments.  

**Note:** If you want to view the details of a specific payment and you have its paymentId, use our [Retrieve Payment](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/retrieve) method.  

Use query parameters to filter the list of results that we return, for example, to search for payments for a customer, a date range, or a settlement state.  

Our gateway returns the following information about each payment in the list:  

- Order details, including the transaction amount and when it was processed.  
- Bank account details, including the customer’s name and account number.  
- Customer's details, including the customer’s phone number.  
- Transaction details, including any refunds or re-presentments.  

For each transaction, we also return the paymentId and an optional secureTokenId, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.ListPaymentsRequest{
        ProcessingTerminalId: "1234001",
        OrderId: payroc.String(
            "OrderRef6543",
        ),
        NameOnAccount: payroc.String(
            "Sarah%20Hazel%20Hopper",
        ),
        Last4: payroc.String(
            "7890",
        ),
        DateFrom: payroc.Time(
            payroc.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        ),
        DateTo: payroc.Time(
            payroc.MustParseDateTime(
                "2024-07-31T23:59:59Z",
            ),
        ),
        SettlementState: banktransferpayments.ListPaymentsRequestSettlementStateSettled.Ptr(),
        SettlementDate: payroc.Time(
            payroc.MustParseDate(
                "2024-07-15",
            ),
        ),
        PaymentLinkId: payroc.String(
            "JZURRJBUPS",
        ),
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.BankTransferPayments.Payments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Filter results by the unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**orderId:** `*string` — Filter results by the order ID of the payment.
    
</dd>
</dl>

<dl>
<dd>

**nameOnAccount:** `*string` — Filter results by the account holder's name.
    
</dd>
</dl>

<dl>
<dd>

**last4:** `*string` — Filter results by the last four digits of the account number.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*banktransferpayments.ListPaymentsRequestTypeItem` — Filter results by transaction type.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*banktransferpayments.ListPaymentsRequestStatusItem` — Filter results by the status of the payment.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `*time.Time` — Filter results by payments that the merchant ran after a specific date. The value follows the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard.
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `*time.Time` — Filter results by payments that the merchant ran before a specific date. The value follows the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard.
    
</dd>
</dl>

<dl>
<dd>

**settlementState:** `*banktransferpayments.ListPaymentsRequestSettlementState` — Filter results by the settlement status.
    
</dd>
</dl>

<dl>
<dd>

**settlementDate:** `*time.Time` — Filter results by the settlement date. Send a value in **YYYY-MM-DD** format.
    
</dd>
</dl>

<dl>
<dd>

**paymentLinkId:** `*string` — Filter results by the paymentLinkId.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankTransferPayments.Payments.Create(request) -> *payroc.BankTransferPayment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to run a sale with a customer's bank account details.  

In the response, our gateway returns information about the bank transfer payment and a paymentId, which you need for the following methods:  
-	[Retrieve payment](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/retrieve) - View the details of the bank transfer payment.
-	[Reverse payment](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/reverse-payment) - Cancel the bank transfer payment if it's an open batch.
-	[Refund payment](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/refund) - Run a referenced refund to return funds to the customer's bank account.

**Payment methods**  

Our gateway accepts the following payment methods:  
-	Automated clearing house (ACH) details
-	Pre-authorized debit (PAD) details  

You can also use [secure tokens](https://docs.payroc.com/api/schema/payments/secure-tokens/overview) and [single-use tokens](https://docs.payroc.com/api/schema/tokenization/single-use-tokens/create) that you created from ACH details or PAD details. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.BankTransferPaymentRequest{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        ProcessingTerminalId: "1234001",
        Order: &payroc.BankTransferPaymentRequestOrder{
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
            Breakdown: &payroc.BankTransferRequestBreakdown{
                Subtotal: 4347,
                Tip: &payroc.Tip{
                    Type: payroc.TipTypePercentage,
                    Percentage: payroc.Float64(
                        10,
                    ),
                },
                Taxes: []*payroc.TaxRate{
                    &payroc.TaxRate{
                        Type: payroc.TaxRateTypeRate,
                        Rate: 5,
                        Name: "Sales Tax",
                    },
                },
            },
        },
        Customer: &payroc.BankTransferCustomer{
            NotificationLanguage: payroc.BankTransferCustomerNotificationLanguageEn.Ptr(),
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
        },
        CredentialOnFile: &payroc.SchemasCredentialOnFile{
            Tokenize: payroc.Bool(
                true,
            ),
        },
        PaymentMethod: &banktransferpayments.BankTransferPaymentRequestPaymentMethod{
            Ach: &payroc.AchPayload{
                NameOnAccount: "Shara Hazel Hopper",
                AccountNumber: "1234567890",
                RoutingNumber: "123456789",
            },
        },
        CustomFields: []*payroc.CustomField{
            &payroc.CustomField{
                Name: "yourCustomField",
                Value: "abc123",
            },
        },
    }
client.BankTransferPayments.Payments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*payroc.BankTransferPaymentRequestOrder` 
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.BankTransferCustomer` 
    
</dd>
</dl>

<dl>
<dd>

**credentialOnFile:** `*payroc.SchemasCredentialOnFile` 
    
</dd>
</dl>

<dl>
<dd>

**paymentMethod:** `*banktransferpayments.BankTransferPaymentRequestPaymentMethod` — Object that contains information about the customer's payment details.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*payroc.CustomField` — Array of customField objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankTransferPayments.Payments.Retrieve(PaymentId) -> *payroc.BankTransferPayment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a bank transfer payment.  

To retrieve a payment, you need its paymentId. Our gateway returned the paymentId in the response of the [Create Payment](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/create) method.  

Note: If you don’t have the paymentId, use our [List Payments](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/list) method to search for the payment.  

Our gateway returns the following information about the payment:  

-	Order details, including the transaction amount and when it was processed.  
-	Bank account details, including the customer’s name and account number.  
-	Customer’s details, including the customer’s phone number.  
-	Transaction details, including any refunds or re-presentments.  

If the merchant saved the customer’s bank account details, our gateway returns a secureTokenID, which you can use to perform follow-on actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.RetrievePaymentsRequest{
        PaymentId: "M2MJOG6O2Y",
    }
client.BankTransferPayments.Payments.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier that our gateway assigned to the payment.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankTransferPayments.Payments.Represent(PaymentId, request) -> *payroc.BankTransferPayment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to re-present an ACH payment.  

To re-present a payment, you need the paymentId of the return. To get the paymentId of the return, complete the following steps:  

1.	Use our [Retrieve Payment](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/retrieve) method  to view the details of the original payment.  
2.	From the [returns object](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/retrieve#response.body.returns) in the response, get the paymentId of the return.  

Our gateway uses the bank account details from the original payment. If you want to update the customer's bank account details, send the new bank account details in the request.  

If your request is successful, our gateway re-presents the payment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.Representment{
        PaymentId: "M2MJOG6O2Y",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        PaymentMethod: &banktransferpayments.RepresentmentPaymentMethod{
            Ach: &payroc.AchPayload{
                NameOnAccount: "Shara Hazel Hopper",
                AccountNumber: "1234567890",
                RoutingNumber: "123456789",
            },
        },
    }
client.BankTransferPayments.Payments.Represent(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier that our gateway assigned to the payment.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**paymentMethod:** `*banktransferpayments.RepresentmentPaymentMethod` — Object that contains information about the customer's payment details.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BankTransferPayments Refunds
<details><summary><code>client.BankTransferPayments.Refunds.ReversePayment(PaymentId) -> *payroc.BankTransferPayment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to cancel a bank transfer payment in an open batch. This is also known as voiding a payment.  

To cancel a bank transfer payment, you need its paymentId. Our gateway returned the paymentId in the response of the [Create Payment](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/create) method.  

**Note:** If you don't have the paymentId, use our [List Payments](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/list) method to search for the bank transfer payment.  

If your request is successful, our gateway removes the bank transfer payment from the merchant’s open batch and no funds are taken from the customer's bank account.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.ReversePaymentRefundsRequest{
        PaymentId: "M2MJOG6O2Y",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
    }
client.BankTransferPayments.Refunds.ReversePayment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier that our gateway assigned to the payment.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankTransferPayments.Refunds.Refund(PaymentId, request) -> *payroc.BankTransferPayment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to refund a bank transfer payment that is in a closed batch.  

To refund a bank transfer payment, you need its paymentId. Our gateway returned the paymentId in the response of the [Create Payment](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/create) method.  

**Note:** If you don’t have the paymentId, use our [List Payments](https://docs.payroc.com/api/schema/bank-transfer-payments/payments/list) method to search for the bank transfer payment.  

If your refund is successful, our gateway returns the payment amount to the customer's account.  

**Things to consider**  
- If the merchant refunds a bank transfer payment that is in an open batch, our gateway reverses the bank transfer payment.  
- Some merchants can run unreferenced refunds, which means that they don’t need a paymentId to return an amount to a customer. For more information about how to run an unreferenced refund, go to [Create Refund](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/create).  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.BankTransferReferencedRefund{
        PaymentId: "M2MJOG6O2Y",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Amount: 4999,
        Description: "amount to refund",
    }
client.BankTransferPayments.Refunds.Refund(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier that our gateway assigned to the payment.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**amount:** `int64` — Total amount of the refund. The value is in the currency's lowest denomination, for example, cents.
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` — Description of the refund.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankTransferPayments.Refunds.List() -> *core.PayrocPager[*payroc.BankTransferRefundPaginatedList, []*payroc.BankTransferRefund, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of bank transfer refunds.  

**Note:** If you want to view the details of a specific refund and you have its refundId, use our [Retrieve Refund](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/retrieve) method.  

Use query parameters to filter the list of results that we return, for example, to search for refunds for a customer, an orderId, or a date range.  

Our gateway returns the following information about each refund in the list:  

-	Order details, including the refund amount and when it was processed.  
-	Bank account details, including the customer’s name and account number.  

For referenced refunds, our gateway also returns details about the payment that the refund is linked to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.ListRefundsRequest{
        ProcessingTerminalId: "1234001",
        OrderId: payroc.String(
            "OrderRef6543",
        ),
        NameOnAccount: payroc.String(
            "Sarah%20Hazel%20Hopper",
        ),
        Last4: payroc.String(
            "7062",
        ),
        DateFrom: payroc.Time(
            payroc.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        ),
        DateTo: payroc.Time(
            payroc.MustParseDateTime(
                "2024-07-31T23:59:59Z",
            ),
        ),
        SettlementState: banktransferpayments.ListRefundsRequestSettlementStateSettled.Ptr(),
        SettlementDate: payroc.Time(
            payroc.MustParseDate(
                "2024-07-15",
            ),
        ),
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.BankTransferPayments.Refunds.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Filter results by the unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**orderId:** `*string` — Filter results by the order ID of the refund.
    
</dd>
</dl>

<dl>
<dd>

**nameOnAccount:** `*string` — Filter results by the accountholder's name.
    
</dd>
</dl>

<dl>
<dd>

**last4:** `*string` — Filter results by the last four digits of the account number.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*banktransferpayments.ListRefundsRequestTypeItem` — Filter results by transaction type.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*banktransferpayments.ListRefundsRequestStatusItem` — Filter results by the status of the refund.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `*time.Time` — Filter results by refunds that the merchant ran after a specific date. The value follows the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard.
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `*time.Time` — Filter results by refunds that the merchant ran before a specific date. The value follows the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard.
    
</dd>
</dl>

<dl>
<dd>

**settlementState:** `*banktransferpayments.ListRefundsRequestSettlementState` — Filter results by the settlement status.
    
</dd>
</dl>

<dl>
<dd>

**settlementDate:** `*time.Time` — Filter results by the settlement date. Send a value in **YYYY-MM-DD** format.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankTransferPayments.Refunds.Create(request) -> *payroc.BankTransferRefund</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create an unreferenced refund. An unreferenced refund is a refund that isn’t linked to a bank transfer payment.  

**Note:** If you have the paymentId of the payment you want to refund, use our [Refund Payment](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/refund) method. If you use our Refund Payment method, our gateway sends the refund amount to the customer’s original payment method and links the refund to the payment.  

In the request, you must provide the customer’s payment method and information about the order including the refund amount.  

In the response, our gateway returns information about the refund and a refundId, which you need for the following methods:  

-	[Retrieve refund](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/retrieve) – View the details of the refund.  
-	[Reverse refund](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/reverse-refund) – Cancel the refund if it’s in an open batch.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.BankTransferUnreferencedRefund{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        ProcessingTerminalId: "1234001",
        Order: &payroc.BankTransferRefundOrder{
            OrderId: payroc.String(
                "OrderRef6543",
            ),
            Description: payroc.String(
                "Refund for order OrderRef6543",
            ),
            Amount: payroc.Int64(
                4999,
            ),
            Currency: payroc.CurrencyUsd.Ptr(),
        },
        Customer: &payroc.BankTransferCustomer{
            NotificationLanguage: payroc.BankTransferCustomerNotificationLanguageEn.Ptr(),
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
        },
        RefundMethod: &banktransferpayments.BankTransferUnreferencedRefundRefundMethod{
            Ach: &payroc.AchPayload{
                NameOnAccount: "Shara Hazel Hopper",
                AccountNumber: "1234567890",
                RoutingNumber: "123456789",
            },
        },
        CustomFields: []*payroc.CustomField{
            &payroc.CustomField{
                Name: "yourCustomField",
                Value: "abc123",
            },
        },
    }
client.BankTransferPayments.Refunds.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*payroc.BankTransferRefundOrder` 
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.BankTransferCustomer` 
    
</dd>
</dl>

<dl>
<dd>

**refundMethod:** `*banktransferpayments.BankTransferUnreferencedRefundRefundMethod` — Object that contains information about how the merchant refunds the customer.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*payroc.CustomField` — Array of customField objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankTransferPayments.Refunds.Retrieve(RefundId) -> *payroc.BankTransferRefund</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a refund.  

To retrieve a refund, you need its refundId. Our gateway returned the refundId in the response of the [Refund Payment](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/refund) method or the [Create Refund](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/create) method.  

**Note:** If you don’t have the refundId, use our [List Refunds](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/list) method to search for the refund.  

Our gateway returns the following information about the refund:  

- Order details, including the refund amount and when it was processed.  
- Bank account details, including the customer’s name and account number.  

If the refund is a referenced refund, our gateway also returns details about the payment that the refund is linked to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.RetrieveRefundsRequest{
        RefundId: "CD3HN88U9F",
    }
client.BankTransferPayments.Refunds.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refundId:** `string` — Unique identifier that our gateway assigned to the refund.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BankTransferPayments.Refunds.ReverseRefund(RefundId) -> *payroc.BankTransferRefund</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to cancel a bank transfer refund in an open batch.  

To cancel a refund, you need its refundId. Our gateway returned the refundId in the response of the [Refund Payment](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/refund) or [Create Refund](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/create) method.  

**Note:** If you don’t have the refundId, use our [List Refunds](https://docs.payroc.com/api/schema/bank-transfer-payments/refunds/list) method to search for the refund.  

If your request is successful, the gateway removes the refund from the merchant’s open batch, and no funds are returned to the cardholder’s account.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &banktransferpayments.ReverseRefundRefundsRequest{
        RefundId: "CD3HN88U9F",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
    }
client.BankTransferPayments.Refunds.ReverseRefund(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refundId:** `string` — Unique identifier that our gateway assigned to the refund.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Boarding Owners
<details><summary><code>client.Boarding.Owners.Retrieve(OwnerId) -> *payroc.Owner</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve details about an owner of a processing account or an owner associated with a funding recipient.  

To retrieve an owner, you need their ownerId. Our gateway returned the ownerId in the response of the [Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) method or the [Create Funding Recipient Owner](https://docs.payroc.com/api/schema/funding/funding-recipients/create-owner) method.  

**Note:** If you don't have the ownerId, use the [Retrieve Processing Account](https://docs.payroc.com/api/schema/boarding/processing-accounts/retrieve) method if you are searching for a processing account owner, or use the [List Funding Recipient Owners](https://docs.payroc.com/api/schema/funding/funding-recipients/list-owners) method if you are searching for a funding recipient owner.  

Our gateway returns the following information about an owner:  
- Name, date of birth, and address.  
- Contact details, including their email address.  
- Relationship to the business, including whether they are a control prong or authorized signatory, and their equity stake in the business. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.RetrieveOwnersRequest{
        OwnerId: 1,
    }
client.Boarding.Owners.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `int` — Unique identifier that we assigned to the owner.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.Owners.Update(OwnerId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Important:** You can't update the details of an owner of a processing account.  

Use this method to update the details of an owner associated with a funding recipient.  

To update an owner, you need their ownerId. Our gateway returned the ownerId in the response of the [Create Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/create) method and the [Create Funding Recipient Owner](https://docs.payroc.com/api/schema/funding/funding-recipients/create-owner) method.   

**Note:** If you don't have the ownerId, use the [List Funding Recipient Owners](https://docs.payroc.com/api/schema/funding/funding-recipients/list-owners) method, the [Retrieve Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/retrieve) method, or the [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient owner.  

You can update the following details about an owner:  

- Personal details, including their name, date of birth, and address.  
- Identification details, including their identification type and number.  
- Contact details, including their email address.  
- Relationship to the business, including whether they are a control prong.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.UpdateOwnersRequest{
        OwnerId: 1,
        Body: &payroc.Owner{
            FirstName: "Jane",
            MiddleName: payroc.String(
                "Helen",
            ),
            LastName: "Doe",
            DateOfBirth: payroc.MustParseDate(
                "1964-03-22",
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
            Identifiers: []*payroc.Identifier{
                &payroc.Identifier{
                    Type: payroc.IdentifierTypeNationalId,
                    Value: "000-00-4320",
                },
            },
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
            Relationship: &payroc.OwnerRelationship{
                EquityPercentage: payroc.Float64(
                    48.5,
                ),
                Title: payroc.String(
                    "CFO",
                ),
                IsControlProng: true,
                IsAuthorizedSignatory: payroc.Bool(
                    false,
                ),
            },
        },
    }
client.Boarding.Owners.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `int` — Unique identifier that we assigned to the owner.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.Owner` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.Owners.Delete(OwnerId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Important:** You can't delete an owner of a processing account. 

Use this method to delete an owner associated with a funding recipient. You can delete an owner only if the funding recipient has more than one owner.  

To delete an owner, you need their ownerId. Our gateway returned the ownerId in the response of the [Create Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/create) method and the [Create Funding Recipient Owner](https://docs.payroc.com/api/schema/funding/funding-recipients/create-owner) method.  

**Note:** If you don't have the ownerId, use the [List Funding Recipient Owners](https://docs.payroc.com/api/schema/funding/funding-recipients/list-owners) method, the [Retrieve Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/retrieve) method, or the [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient owner.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.DeleteOwnersRequest{
        OwnerId: 1,
    }
client.Boarding.Owners.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ownerId:** `int` — Unique identifier that we assigned to the owner.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Boarding PricingIntents
<details><summary><code>client.Boarding.PricingIntents.List() -> *core.PayrocPager[*payroc.PaginatedPricingIntent, []payroc.PricingIntent, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of pricing intents associated with the ISV.  

**Note:** If you want to view the details of a specific pricing intent and you have its pricingIntentId, use our [Retrieve Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/retrieve) method.  

Our gateway returns the following information about each pricing intent in the list:  

- Information about the fees, including the base fees, gateway fees, and processor fees.  
- Status of the pricing intent, including whether we approved the pricing intent.  

For each pricing intent, we also return its pricingIntentId which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.ListPricingIntentsRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.Boarding.PricingIntents.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.PricingIntents.Create(request) -> payroc.PricingIntent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a pricing intent that you can assign to a processing account.  

In the request, you must provide the following:
-	Processing fees, including the pricing program and the fee to process each transaction.
-	Gateway fees, including the fee for each transaction handled by our gateway.
-	Base fees, including maintenance and PCI fees.

In the response, our gateway returns information about the pricing intent and the pricingIntentId, which you need for the following methods:
-	[Create Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create) - Assign the pricing intent to a processing account, when you create the merchant platform and its processing accounts.
-	[Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) - Assign the pricing intent to a processing account.
-	[Retrieve Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/retrieve) - Retrieve information about a pricing intent.
-	[Update Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/update) - Update the details of a pricing intent. 
-	[Delete Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/delete) - Delete a pricing intent.
-	[Partially Update Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/partially-update) - Partially update the details of a pricing intent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.CreatePricingIntentsRequest{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.PricingIntent50{
            Country: payroc.PricingAgreementUs50CountryUs,
            Version: payroc.PricingAgreementUs50VersionFive0,
            Base: &payroc.BaseUs{
                AddressVerification: payroc.Int(
                    5,
                ),
                AnnualFee: &payroc.BaseUsAnnualFee{
                    BillInMonth: payroc.BaseUsAnnualFeeBillInMonthJune.Ptr(),
                    Amount: 9900,
                },
                RegulatoryAssistanceProgram: payroc.Int(
                    15,
                ),
                PciNonCompliance: payroc.Int(
                    4995,
                ),
                MerchantAdvantage: payroc.Int(
                    10,
                ),
                PlatinumSecurity: &payroc.BaseUsPlatinumSecurity{
                    Monthly: &payroc.BaseUsMonthly{},
                },
                Maintenance: 500,
                Minimum: 100,
                VoiceAuthorization: payroc.Int(
                    95,
                ),
                Chargeback: payroc.Int(
                    2500,
                ),
                Retrieval: payroc.Int(
                    1500,
                ),
                Batch: 1500,
                EarlyTermination: payroc.Int(
                    57500,
                ),
            },
            Processor: &payroc.PricingAgreementUs50Processor{
                Card: &payroc.PricingAgreementUs50ProcessorCard{
                    InterchangePlus: &payroc.InterchangePlus{
                        Fees: &payroc.InterchangePlusFees{
                            MastercardVisaDiscover: &payroc.ProcessorFee{},
                        },
                    },
                },
            },
            Services: &payroc.ServicesUs50{
                &payroc.ServiceUs50{
                    HardwareAdvantagePlan: &payroc.HardwareAdvantagePlan{
                        Enabled: true,
                    },
                },
            },
            Key: "Your-Unique-Identifier",
            Metadata: map[string]string{
                "yourCustomField": "abc123",
            },
        },
    }
client.Boarding.PricingIntents.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `payroc.PricingIntent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.PricingIntents.Retrieve(PricingIntentId) -> payroc.PricingIntent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a pricing intent.  

To retrieve a pricing intent, you need its pricingIntentId. Our gateway returned the pricingIntentId in the response of the [Create Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/create) method.  

**Note:** If you don't have the pricingIntentId, use our [List Pricing Intents](https://docs.payroc.com/api/schema/boarding/pricing-intents/list) method to search for the pricing intent.  

Our gateway returns the following information about the pricing intent:  

- Information about the fees, including the base fees, gateway fees, and processor fees.  
- Status of the pricing intent, including whether we approved the pricing intent.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.RetrievePricingIntentsRequest{
        PricingIntentId: "5",
    }
client.Boarding.PricingIntents.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pricingIntentId:** `string` — Unique identifier that we assigned to the pricing intent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.PricingIntents.Update(PricingIntentId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to update the details of a pricing intent. If you update a pricing intent, it won't affect merchant that you've previously onboarded.  

To update a pricing intent, you need its pricingIntentId. Our gateway returned the pricingIntentId in the response of the [Create Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/create) method.  

**Note:** If you don't have the pricingIntentId, use our [List Pricing Intents](https://docs.payroc.com/api/schema/boarding/pricing-intents/list) method to search for the pricing intent.  

You can update the following details about a pricing intent:  

- Fees, including the base fees, processor fees, and gateway fees.  
- Custom name for the pricing intent.  
- Additional services that merchants can sign up for.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.UpdatePricingIntentsRequest{
        PricingIntentId: "5",
        Body: &payroc.PricingIntent50{
            Country: payroc.PricingAgreementUs50CountryUs,
            Version: payroc.PricingAgreementUs50VersionFive0,
            Base: &payroc.BaseUs{
                AddressVerification: payroc.Int(
                    5,
                ),
                AnnualFee: &payroc.BaseUsAnnualFee{
                    BillInMonth: payroc.BaseUsAnnualFeeBillInMonthJune.Ptr(),
                    Amount: 9900,
                },
                RegulatoryAssistanceProgram: payroc.Int(
                    15,
                ),
                PciNonCompliance: payroc.Int(
                    4995,
                ),
                MerchantAdvantage: payroc.Int(
                    10,
                ),
                PlatinumSecurity: &payroc.BaseUsPlatinumSecurity{
                    Monthly: &payroc.BaseUsMonthly{},
                },
                Maintenance: 500,
                Minimum: 100,
                VoiceAuthorization: payroc.Int(
                    95,
                ),
                Chargeback: payroc.Int(
                    2500,
                ),
                Retrieval: payroc.Int(
                    1500,
                ),
                Batch: 1500,
                EarlyTermination: payroc.Int(
                    57500,
                ),
            },
            Processor: &payroc.PricingAgreementUs50Processor{
                Card: &payroc.PricingAgreementUs50ProcessorCard{
                    InterchangePlus: &payroc.InterchangePlus{
                        Fees: &payroc.InterchangePlusFees{
                            MastercardVisaDiscover: &payroc.ProcessorFee{},
                        },
                    },
                },
                Ach: &payroc.Ach{
                    Fees: &payroc.AchFees{
                        Transaction: 50,
                        Batch: 5,
                        Returns: 400,
                        UnauthorizedReturn: 1999,
                        Statement: 800,
                        MonthlyMinimum: 20000,
                        AccountVerification: 10,
                        DiscountRateUnder10000: 5.25,
                        DiscountRateAbove10000: 10,
                    },
                },
            },
            Gateway: &payroc.GatewayUs50{
                Fees: &payroc.GatewayUs50Fees{
                    Monthly: 2000,
                    Setup: 5000,
                    PerTransaction: 2000,
                    PerDeviceMonthly: 10,
                },
            },
            Services: &payroc.ServicesUs50{
                &payroc.ServiceUs50{
                    HardwareAdvantagePlan: &payroc.HardwareAdvantagePlan{
                        Enabled: true,
                    },
                },
            },
            Key: "Your-Unique-Identifier",
            Metadata: map[string]string{
                "yourCustomField": "abc123",
            },
        },
    }
client.Boarding.PricingIntents.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pricingIntentId:** `string` — Unique identifier that we assigned to the pricing intent.
    
</dd>
</dl>

<dl>
<dd>

**request:** `payroc.PricingIntent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.PricingIntents.Delete(PricingIntentId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to delete a pricing intent.  

> **Important:** When you delete a pricing intent, you can't recover it. You also won't be able to assign the pricing intent to a merchant's boarding application.  

To delete a pricing intent, you need its pricingIntentId. Our gateway returned the pricingIntentId in the response of the [Create Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/create) method.  

**Note:** If you don't have the pricingIntentId, use our [List Pricing Intents](https://docs.payroc.com/api/schema/boarding/pricing-intents/list) method to search for the pricing intent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.DeletePricingIntentsRequest{
        PricingIntentId: "5",
    }
client.Boarding.PricingIntents.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pricingIntentId:** `string` — Unique identifier that we assigned to the pricing intent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.PricingIntents.PartiallyUpdate(PricingIntentId, request) -> payroc.PricingIntent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to partially update a pricing intent. Structure your request to follow the [RFC 6902](https://datatracker.ietf.org/doc/html/rfc6902) standard.  

If you update a pricing intent, it won't affect merchants you've previously onboarded.  

To update a pricing intent, you need its pricingIntentId. Our gateway returned the pricingIntentId in the response of the [Create Pricing Intent](https://docs.payroc.com/api/schema/boarding/pricing-intents/create) method.  

**Note:** If you don't have the pricingIntentId, use our [List Pricing Intents](https://docs.payroc.com/api/schema/boarding/pricing-intents/list) method to search for the pricing intent.  

You can update the following details about a pricing intent:  

- Fees, including the base fees, processor fees, and gateway fees.  
- Custom name for the pricing intent.  
- Additional services that merchants can sign up for.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.PartiallyUpdatePricingIntentsRequest{
        PricingIntentId: "5",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: []*payroc.PatchDocument{
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
        },
    }
client.Boarding.PricingIntents.PartiallyUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**pricingIntentId:** `string` — Unique identifier that we assigned to the pricing intent.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `payroc.PatchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Boarding MerchantPlatforms
<details><summary><code>client.Boarding.MerchantPlatforms.List() -> *core.PayrocPager[*payroc.PaginatedMerchants, []*payroc.MerchantPlatform, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of merchant platforms that are linked to your ISV account.  

**Note**: If you want to view the details of a specific merchant platform and you have its merchantPlatformId, use our [Retrieve Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/retrieve) method.  

Our gateway returns the following information about each merchant platform in the list:  
- Legal information, including its legal name and address.  
- Contact information, including the email address for the business.  
- Processing  account information, including the processingAccountId and status of each processing account that's linked to the merchant platform.  

For each merchant platform, we also return its merchantPlatformId and its linked processingAccountIds, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.ListMerchantPlatformsRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.Boarding.MerchantPlatforms.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.MerchantPlatforms.Create(request) -> *payroc.MerchantPlatform</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to board a merchant with Payroc.  

**Note**: This method is part of our Boarding solution. To help you understand how this method works with other Boarding methods, go to [Board a Merchant](https://docs.payroc.com/guides/integrate/boarding).  

In the request, include the following information:  
- Legal information, including its legal name and address.  
- Contact information, including the email address for the business.  
- Processing account information, including the pricing model, owners, and contacts for the processing account.  

When you send a successful request, we review the merchant's information. After we complete our review and approve the merchant, we assign:  
- **merchantPlatformId** - Unique identifier for the merchant platform.  
- **processingAccountId** - Unique identifier for each processing account linked to the merchant platform.  

You need to keep these to perform follow-on actions, for example, you need the processingAccountId to [order terminals](https://docs.payroc.com/api/schema/boarding/processing-accounts/create-terminal-order) for the processing account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.CreateMerchantAccount{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Business: &payroc.Business{
            Name: "Example Corp",
            TaxId: "12-3456789",
            OrganizationType: payroc.BusinessOrganizationTypePrivateCorporation,
            CountryOfOperation: payroc.BusinessCountryOfOperationUs.Ptr(),
            Addresses: []*payroc.LegalAddress{
                &payroc.LegalAddress{
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
                    Type: payroc.AddressTypeTypeLegalAddress,
                },
            },
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
        },
        ProcessingAccounts: []*payroc.CreateProcessingAccount{
            &payroc.CreateProcessingAccount{
                DoingBusinessAs: "Pizza Doe",
                Owners: []*payroc.Owner{
                    &payroc.Owner{
                        FirstName: "Jane",
                        MiddleName: payroc.String(
                            "Helen",
                        ),
                        LastName: "Doe",
                        DateOfBirth: payroc.MustParseDate(
                            "1964-03-22",
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
                        Identifiers: []*payroc.Identifier{
                            &payroc.Identifier{
                                Type: payroc.IdentifierTypeNationalId,
                                Value: "000-00-4320",
                            },
                        },
                        ContactMethods: []*payroc.ContactMethod{
                            &payroc.ContactMethod{
                                Email: &payroc.ContactMethodEmail{
                                    Value: "jane.doe@example.com",
                                },
                            },
                        },
                        Relationship: &payroc.OwnerRelationship{
                            EquityPercentage: payroc.Float64(
                                48.5,
                            ),
                            Title: payroc.String(
                                "CFO",
                            ),
                            IsControlProng: true,
                            IsAuthorizedSignatory: payroc.Bool(
                                false,
                            ),
                        },
                    },
                },
                Website: payroc.String(
                    "www.example.com",
                ),
                BusinessType: payroc.CreateProcessingAccountBusinessTypeRestaurant,
                CategoryCode: 5999,
                MerchandiseOrServiceSold: "Pizza",
                BusinessStartDate: payroc.MustParseDate(
                    "2020-01-01",
                ),
                Timezone: payroc.TimezoneAmericaChicago,
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
                ContactMethods: []*payroc.ContactMethod{
                    &payroc.ContactMethod{
                        Email: &payroc.ContactMethodEmail{
                            Value: "jane.doe@example.com",
                        },
                    },
                },
                Processing: &payroc.Processing{
                    TransactionAmounts: &payroc.ProcessingTransactionAmounts{
                        Average: 5000,
                        Highest: 10000,
                    },
                    MonthlyAmounts: &payroc.ProcessingMonthlyAmounts{
                        Average: 50000,
                        Highest: 100000,
                    },
                    VolumeBreakdown: &payroc.ProcessingVolumeBreakdown{
                        CardPresent: 77,
                        MailOrTelephone: 3,
                        Ecommerce: 20,
                    },
                    IsSeasonal: payroc.Bool(
                        true,
                    ),
                    MonthsOfOperation: []payroc.ProcessingMonthsOfOperationItem{
                        payroc.ProcessingMonthsOfOperationItemJan,
                        payroc.ProcessingMonthsOfOperationItemFeb,
                    },
                    Ach: &payroc.ProcessingAch{
                        Naics: payroc.String(
                            "5812",
                        ),
                        PreviouslyTerminatedForAch: payroc.Bool(
                            false,
                        ),
                        Refunds: &payroc.ProcessingAchRefunds{
                            WrittenRefundPolicy: true,
                            RefundPolicyUrl: payroc.String(
                                "www.example.com/refund-poilcy-url",
                            ),
                        },
                        EstimatedMonthlyTransactions: 3000,
                        Limits: &payroc.ProcessingAchLimits{
                            SingleTransaction: 10000,
                            DailyDeposit: 200000,
                            MonthlyDeposit: 6000000,
                        },
                        TransactionTypes: []payroc.ProcessingAchTransactionTypesItem{
                            payroc.ProcessingAchTransactionTypesItemPrearrangedPayment,
                            payroc.ProcessingAchTransactionTypesItemOther,
                        },
                        TransactionTypesOther: payroc.String(
                            "anotherTransactionType",
                        ),
                    },
                    CardAcceptance: &payroc.ProcessingCardAcceptance{
                        DebitOnly: payroc.Bool(
                            false,
                        ),
                        HsaFsa: payroc.Bool(
                            false,
                        ),
                        CardsAccepted: []payroc.ProcessingCardAcceptanceCardsAcceptedItem{
                            payroc.ProcessingCardAcceptanceCardsAcceptedItemVisa,
                            payroc.ProcessingCardAcceptanceCardsAcceptedItemMastercard,
                        },
                        SpecialityCards: &payroc.ProcessingCardAcceptanceSpecialityCards{
                            AmericanExpressDirect: &payroc.ProcessingCardAcceptanceSpecialityCardsAmericanExpressDirect{
                                Enabled: payroc.Bool(
                                    true,
                                ),
                                MerchantNumber: payroc.String(
                                    "abc1234567",
                                ),
                            },
                            ElectronicBenefitsTransfer: &payroc.ProcessingCardAcceptanceSpecialityCardsElectronicBenefitsTransfer{
                                Enabled: payroc.Bool(
                                    true,
                                ),
                                FnsNumber: payroc.String(
                                    "6789012",
                                ),
                            },
                            Other: &payroc.ProcessingCardAcceptanceSpecialityCardsOther{
                                WexMerchantNumber: payroc.String(
                                    "abc1234567",
                                ),
                                VoyagerMerchantId: payroc.String(
                                    "abc1234567",
                                ),
                                FleetMerchantId: payroc.String(
                                    "abc1234567",
                                ),
                            },
                        },
                    },
                },
                Funding: &payroc.CreateFunding{
                    FundingSchedule: payroc.CommonFundingFundingScheduleNextday.Ptr(),
                    AcceleratedFundingFee: payroc.Int(
                        1999,
                    ),
                    DailyDiscount: payroc.Bool(
                        false,
                    ),
                    FundingAccounts: []*payroc.FundingAccount{
                        &payroc.FundingAccount{
                            Type: payroc.FundingAccountTypeChecking,
                            Use: payroc.FundingAccountUseCreditAndDebit,
                            NameOnAccount: "Jane Doe",
                            PaymentMethods: []*payroc.PaymentMethodsItem{
                                &payroc.PaymentMethodsItem{
                                    Ach: &payroc.PaymentMethodAch{},
                                },
                            },
                            Metadata: map[string]string{
                                "yourCustomField": "abc123",
                            },
                        },
                    },
                },
                Pricing: &payroc.Pricing{
                    Intent: &payroc.PricingTemplate{
                        PricingIntentId: "6123",
                    },
                },
                Signature: &payroc.Signature{
                    RequestedViaDirectLink: &payroc.SignatureByDirectLink{},
                },
                Contacts: []*payroc.Contact{
                    &payroc.Contact{
                        Type: payroc.ContactTypeManager,
                        FirstName: "Jane",
                        MiddleName: payroc.String(
                            "Helen",
                        ),
                        LastName: "Doe",
                        Identifiers: []*payroc.Identifier{
                            &payroc.Identifier{
                                Type: payroc.IdentifierTypeNationalId,
                                Value: "000-00-4320",
                            },
                        },
                        ContactMethods: []*payroc.ContactMethod{
                            &payroc.ContactMethod{
                                Email: &payroc.ContactMethodEmail{
                                    Value: "jane.doe@example.com",
                                },
                            },
                        },
                    },
                },
                Metadata: map[string]string{
                    "customerId": "2345",
                },
            },
        },
        Metadata: map[string]string{
            "customerId": "2345",
        },
    }
client.Boarding.MerchantPlatforms.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**business:** `*payroc.Business` 
    
</dd>
</dl>

<dl>
<dd>

**processingAccounts:** `[]*payroc.CreateProcessingAccount` — Array of processingAccounts objects.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]string` — Object that you can send to include custom data in the request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.MerchantPlatforms.Retrieve(MerchantPlatformId) -> *payroc.MerchantPlatform</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a merchant platform.  

To retrieve a merchant platform, you need its merchantPlatformId. Our gateway returned the merchantPlatformId in the response of the [Create Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create) method.  

**Note:** If you don't have the merchantPlatformId, use our [List Merchant Platforms](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list) method to search for the merchant platform.

Our gateway returns the following information about the merchant platform:
-	Legal information, including its legal name and address.
-	Contact information, including the email address for the business. 
-	Processing account information, including the processingAccountId and status of each processing account that's linked to the merchant platform.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.RetrieveMerchantPlatformsRequest{
        MerchantPlatformId: "12345",
    }
client.Boarding.MerchantPlatforms.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantPlatformId:** `string` — Unique identifier of the merchant platform that we sent to you when you created the merchant platform.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.MerchantPlatforms.ListProcessingAccounts(MerchantPlatformId) -> *core.PayrocPager[*payroc.PaginatedProcessingAccounts, []*payroc.ProcessingAccount, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of processing accounts linked to a merchant platform.  

**Note**: If you want to view the details of a specific processing account and you have its processingAccountId, use our [Retrieve Processing Account](https://docs.payroc.com/api/schema/boarding/processing-accounts/retrieve) method.  

Use the query parameters to filter the list of results that we return, for example, to search for only closed processing accounts.  

To list the processing accounts for a merchant platform, you need its merchantPlatformId. If you don't have the merchantPlatformId, use our [List Merchant Platforms](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list) method to search for the merchant platform.

Our gateway returns the following information about eahc processing account in the list:  
- Business details, including its status, time zone, and address.  
- Owners' details, including their contact details.  
- Funding, pricing, and processing information, including its pricing model and funding accounts.  

For each processing account, we also return its processingAccountId, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.ListBoardingMerchantPlatformProcessingAccountsRequest{
        MerchantPlatformId: "12345",
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        IncludeClosed: payroc.Bool(
            true,
        ),
    }
client.Boarding.MerchantPlatforms.ListProcessingAccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantPlatformId:** `string` — Unique identifier of the merchant platform that we sent to you when you created the merchant platform.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**includeClosed:** `*bool` 

Indicates if you want to return closed processing accounts. This includes processing accounts that have a status of `terminated`, `cancelled`, or `rejected`.  
**Note**: By default, we return only open processing accounts.  
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.MerchantPlatforms.CreateProcessingAccount(MerchantPlatformId, request) -> *payroc.ProcessingAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to add an additional processing account to a merchant platform.  

To add a processing account to a merchant platform, you need the merchantPlatformId. Our gateway returned the merchantPlatformId in the response of the [Create Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create) method.  

**Note**: If you don't have the merchantPlatformId, use our [List Merchant Platforms](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list) method to search for the merchant platform.  

In the request, include the following information:  
- Business details, including its business type, time zone, and address.  
- Owners' details, including their contact details.  
- Funding, pricing, and processing information, including its pricing model and funding accounts.  

When you send a successful request, we review the information about the processing account. After we complete our review and approve the processing account, we assign a processingAccountId, which you need to perform follow-on actions.  

**Note**: You can subscribe to our processingAccount.status.changed event to get notifications when we update the status of a processing account. For more information about how to subscribe to events, go to [Events List](https://docs.payroc.com/knowledge/events/events-list).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.CreateProcessingAccountMerchantPlatformsRequest{
        MerchantPlatformId: "12345",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.CreateProcessingAccount{
            DoingBusinessAs: "Pizza Doe",
            Owners: []*payroc.Owner{
                &payroc.Owner{
                    FirstName: "Jane",
                    MiddleName: payroc.String(
                        "Helen",
                    ),
                    LastName: "Doe",
                    DateOfBirth: payroc.MustParseDate(
                        "1964-03-22",
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
                    Identifiers: []*payroc.Identifier{
                        &payroc.Identifier{
                            Type: payroc.IdentifierTypeNationalId,
                            Value: "000-00-4320",
                        },
                    },
                    ContactMethods: []*payroc.ContactMethod{
                        &payroc.ContactMethod{
                            Email: &payroc.ContactMethodEmail{
                                Value: "jane.doe@example.com",
                            },
                        },
                    },
                    Relationship: &payroc.OwnerRelationship{
                        EquityPercentage: payroc.Float64(
                            51.5,
                        ),
                        Title: payroc.String(
                            "CFO",
                        ),
                        IsControlProng: true,
                        IsAuthorizedSignatory: payroc.Bool(
                            false,
                        ),
                    },
                },
            },
            Website: payroc.String(
                "www.example.com",
            ),
            BusinessType: payroc.CreateProcessingAccountBusinessTypeRestaurant,
            CategoryCode: 5999,
            MerchandiseOrServiceSold: "Pizza",
            BusinessStartDate: payroc.MustParseDate(
                "2020-01-01",
            ),
            Timezone: payroc.TimezoneAmericaChicago,
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
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
            Processing: &payroc.Processing{
                TransactionAmounts: &payroc.ProcessingTransactionAmounts{
                    Average: 5000,
                    Highest: 10000,
                },
                MonthlyAmounts: &payroc.ProcessingMonthlyAmounts{
                    Average: 50000,
                    Highest: 100000,
                },
                VolumeBreakdown: &payroc.ProcessingVolumeBreakdown{
                    CardPresent: 77,
                    MailOrTelephone: 3,
                    Ecommerce: 20,
                },
                IsSeasonal: payroc.Bool(
                    true,
                ),
                MonthsOfOperation: []payroc.ProcessingMonthsOfOperationItem{
                    payroc.ProcessingMonthsOfOperationItemJan,
                    payroc.ProcessingMonthsOfOperationItemFeb,
                },
                Ach: &payroc.ProcessingAch{
                    Naics: payroc.String(
                        "5812",
                    ),
                    PreviouslyTerminatedForAch: payroc.Bool(
                        false,
                    ),
                    Refunds: &payroc.ProcessingAchRefunds{
                        WrittenRefundPolicy: true,
                        RefundPolicyUrl: payroc.String(
                            "www.example.com/refund-poilcy-url",
                        ),
                    },
                    EstimatedMonthlyTransactions: 3000,
                    Limits: &payroc.ProcessingAchLimits{
                        SingleTransaction: 10000,
                        DailyDeposit: 200000,
                        MonthlyDeposit: 6000000,
                    },
                    TransactionTypes: []payroc.ProcessingAchTransactionTypesItem{
                        payroc.ProcessingAchTransactionTypesItemPrearrangedPayment,
                        payroc.ProcessingAchTransactionTypesItemOther,
                    },
                    TransactionTypesOther: payroc.String(
                        "anotherTransactionType",
                    ),
                },
                CardAcceptance: &payroc.ProcessingCardAcceptance{
                    DebitOnly: payroc.Bool(
                        false,
                    ),
                    HsaFsa: payroc.Bool(
                        false,
                    ),
                    CardsAccepted: []payroc.ProcessingCardAcceptanceCardsAcceptedItem{
                        payroc.ProcessingCardAcceptanceCardsAcceptedItemVisa,
                        payroc.ProcessingCardAcceptanceCardsAcceptedItemMastercard,
                    },
                    SpecialityCards: &payroc.ProcessingCardAcceptanceSpecialityCards{
                        AmericanExpressDirect: &payroc.ProcessingCardAcceptanceSpecialityCardsAmericanExpressDirect{
                            Enabled: payroc.Bool(
                                true,
                            ),
                            MerchantNumber: payroc.String(
                                "abc1234567",
                            ),
                        },
                        ElectronicBenefitsTransfer: &payroc.ProcessingCardAcceptanceSpecialityCardsElectronicBenefitsTransfer{
                            Enabled: payroc.Bool(
                                true,
                            ),
                            FnsNumber: payroc.String(
                                "6789012",
                            ),
                        },
                        Other: &payroc.ProcessingCardAcceptanceSpecialityCardsOther{
                            WexMerchantNumber: payroc.String(
                                "abc1234567",
                            ),
                            VoyagerMerchantId: payroc.String(
                                "abc1234567",
                            ),
                            FleetMerchantId: payroc.String(
                                "abc1234567",
                            ),
                        },
                    },
                },
            },
            Funding: &payroc.CreateFunding{
                FundingSchedule: payroc.CommonFundingFundingScheduleNextday.Ptr(),
                AcceleratedFundingFee: payroc.Int(
                    1999,
                ),
                DailyDiscount: payroc.Bool(
                    false,
                ),
                FundingAccounts: []*payroc.FundingAccount{
                    &payroc.FundingAccount{
                        Type: payroc.FundingAccountTypeChecking,
                        Use: payroc.FundingAccountUseCreditAndDebit,
                        NameOnAccount: "Jane Doe",
                        PaymentMethods: []*payroc.PaymentMethodsItem{
                            &payroc.PaymentMethodsItem{
                                Ach: &payroc.PaymentMethodAch{},
                            },
                        },
                        Metadata: map[string]string{
                            "yourCustomField": "abc123",
                        },
                    },
                },
            },
            Pricing: &payroc.Pricing{
                Intent: &payroc.PricingTemplate{
                    PricingIntentId: "6123",
                },
            },
            Signature: &payroc.Signature{
                RequestedViaDirectLink: &payroc.SignatureByDirectLink{},
            },
            Contacts: []*payroc.Contact{
                &payroc.Contact{
                    Type: payroc.ContactTypeManager,
                    FirstName: "Jane",
                    MiddleName: payroc.String(
                        "Helen",
                    ),
                    LastName: "Doe",
                    Identifiers: []*payroc.Identifier{
                        &payroc.Identifier{
                            Type: payroc.IdentifierTypeNationalId,
                            Value: "000-00-4320",
                        },
                    },
                    ContactMethods: []*payroc.ContactMethod{
                        &payroc.ContactMethod{
                            Email: &payroc.ContactMethodEmail{
                                Value: "jane.doe@example.com",
                            },
                        },
                    },
                },
            },
            Metadata: map[string]string{
                "customerId": "2345",
            },
        },
    }
client.Boarding.MerchantPlatforms.CreateProcessingAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**merchantPlatformId:** `string` — Unique identifier of the merchant platform that we sent to you when you created the merchant platform.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.CreateProcessingAccount` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Boarding ProcessingAccounts
<details><summary><code>client.Boarding.ProcessingAccounts.Retrieve(ProcessingAccountId) -> *payroc.ProcessingAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a specific processing account.  

To retrieve a processing account, you need its processingAccountId. Our gateway returned the processingAccountId in the response of the [Create Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create) method or the [Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) method.  

**Note:** If you don't have the processingAccountId, use our [List Merchant Platform's Processing Accounts](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list-processing-accounts) method to search for the processing account.  

Our gateway returns the following information about the processing account:  

-	Business information, including the Merchant Category Code (MCC), status of the processing account, and address of the business.  
-	Processing information, including the merchant’s refund policies and card types that the merchant accepts.  
-	Funding information, including funding schedules, funding fees, and details for the merchant’s funding accounts.  
-	Pricing information, including [HATEOAS](https://docs.payroc.com/knowledge/basic-concepts/hypermedia-as-the-engine-of-application-state-hateoas) links to retrieve the pricing program for the processing account.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.RetrieveProcessingAccountsRequest{
        ProcessingAccountId: "38765",
    }
client.Boarding.ProcessingAccounts.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingAccounts.ListProcessingAccountFundingAccounts(ProcessingAccountId) -> []*payroc.FundingAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of funding accounts associated with a processing account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.ListProcessingAccountFundingAccountsRequest{
        ProcessingAccountId: "38765",
    }
client.Boarding.ProcessingAccounts.ListProcessingAccountFundingAccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingAccounts.ListContacts(ProcessingAccountId) -> *payroc.PaginatedContacts</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a list of contacts for a processing account.    

**Note:** If you want to view the details of a specific contact and you have their contactId, use our [Retrieve Contact](https://docs.payroc.com/api/schema/boarding/contacts/retrieve) method.  

To list contacts for a processing account, you need the processingAccountId. Our gateway returned the processingAccountId in the response of the [Create Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create) method or the [Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) method.  

Our gateway returns the following information about each contact:  

- Name and contact method, including their phone number or mobile number.  
- Role within the business, for example, if they are a manager.  

For each contact, we also return a contactId, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.ListContactsProcessingAccountsRequest{
        ProcessingAccountId: "38765",
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.Boarding.ProcessingAccounts.ListContacts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingAccounts.GetProcessingAccountPricingAgreement(ProcessingAccountId) -> *boarding.GetProcessingAccountPricingAgreementProcessingAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve the pricing agreement that we apply to a processing account.  

To retrieve the pricing agreement of a processing account, you need the processingAccountId. Our gateway returned the processingAccountId in the response to the [Create Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create) method and [Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) method.  

**Note:** If you don't have the processingAccountId, use our [List Merchant Platform’s Processing Accounts](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list-processing-accounts) method to search for the processing account.  

Our gateway returns the following information about the pricing agreement that we apply to the processing account:  

- Base fees, including the annual fee and the fees for each chargeback and retrieval.  
- Processor fees, including the fees that we apply for processing card and ACH payments.  
- Gateway fees, including the setup fee and the fees for each transaction.  
- Service fees, including the fee that we apply if the merchant has signed up to a Hardware Advantage Plan.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.GetProcessingAccountPricingAgreementProcessingAccountsRequest{
        ProcessingAccountId: "38765",
    }
client.Boarding.ProcessingAccounts.GetProcessingAccountPricingAgreement(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingAccounts.ListOwners(ProcessingAccountId) -> *core.PayrocPager[*payroc.PaginatedOwners, []*payroc.Owner, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a list of owners of a processing account.  

**Note:** If you want to view the details of a specific owner and you have the ownerId, go to our [Retrieve Owner](https://docs.payroc.com/api/schema/boarding/owners/retrieve) method.  

To list the owners of a processing account, you need its processingAccountId. If you don't have the processingAccountId, use our [List Merchant Platform's Processing Accounts](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list-processing-accounts) method to search for the processing account.  

Our gateway returns the following information about each owner in the list: 

- Name, date of birth, and address.  
- Contact details, including their email address.  
- Relationship to the business, including whether they are a control prong or authorized signatory, and their equity stake in the business.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.ListProcessingAccountOwnersRequest{
        ProcessingAccountId: "38765",
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.Boarding.ProcessingAccounts.ListOwners(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingAccounts.CreateReminder(ProcessingAccountId, request) -> *boarding.CreateReminderProcessingAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to prompt a merchant to sign their pricing agreement.  

You can create a reminder only if you requested the merchant’s signature by email when you created the processing account for the merchant.  

To create a reminder, you need the processingAccountId. Our gateway returned the processingAccountId in the response of the [Create Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create) method or [Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) method. 

**Note:** If you don’t know the processingAccountId, use our [List Merchant Platform’s Processing Accounts](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list-processing-accounts) method to search for the processing account.  

When you send a successful request, we send an email to the merchant that prompts them to sign their pricing agreement.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.CreateReminderProcessingAccountsRequest{
        ProcessingAccountId: "38765",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &boarding.CreateReminderProcessingAccountsRequestBody{
            PricingAgreement: &payroc.PricingAgreementReminder{},
        },
    }
client.Boarding.ProcessingAccounts.CreateReminder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*boarding.CreateReminderProcessingAccountsRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingAccounts.ListTerminalOrders(ProcessingAccountId) -> []*payroc.TerminalOrder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of terminal orders associated with a processing account.  

**Note:** If you want to view the details of a specific terminal order and you have its terminalOrderId, use our [Retrieve Terminal Order](https://docs.payroc.com/api/schema/boarding/terminal-orders/retrieve) method.  

Use the query parameters to filter the list of results that we return, for example, to search for terminal orders by their status.  

To list the terminal orders for a processing account, you need its processingAccountId. If you don't have the processingAccountId, use our [List Merchant Platforms](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list) method to search for a merchant platform and its processing accounts.

Our gateway returns the following information for each terminal order in the list:  

- Status of the order  
- Items in the order  
- Training provider  
- Shipping information  

For each terminal order, we also return its terminalOrderId, which you can use to perform follow-on actions.   
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.ListTerminalOrdersProcessingAccountsRequest{
        ProcessingAccountId: "38765",
        Status: boarding.ListTerminalOrdersProcessingAccountsRequestStatusOpen.Ptr(),
        FromDateTime: payroc.Time(
            payroc.MustParseDateTime(
                "2024-09-08T12:00:00Z",
            ),
        ),
        ToDateTime: payroc.Time(
            payroc.MustParseDateTime(
                "2024-12-08T11:00:00Z",
            ),
        ),
    }
client.Boarding.ProcessingAccounts.ListTerminalOrders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*boarding.ListTerminalOrdersProcessingAccountsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**fromDateTime:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**toDateTime:** `*time.Time` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingAccounts.CreateTerminalOrder(ProcessingAccountId, request) -> *payroc.TerminalOrder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to order and configure terminals for a processing account.  

**Note**: You need the ID of the processing account before you can create an order. If you don't know the processingAccountId, go to the [Retrieve a Merchant Platform](https://docs.payroc.com/api/schema/boarding/merchant-platforms/retrieve) method.  

In the request, specify the gateway settings, device settings, and application settings for the terminal.  

In the response, our gateway returns information about the terminal order including its status and terminalOrderId that you can use to [retrieve the terminal order](https://docs.payroc.com/api/schema/boarding/terminal-orders/retrieve).  

**Note**: You can subscribe to the terminalOrder.status.changed event to get notifications when we update the status of a terminal order. For more information about how to subscribe to events, go to [Events Subscriptions](https://docs.payroc.com/guides/integrate/event-subscriptions).  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.CreateTerminalOrder{
        ProcessingAccountId: "38765",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        TrainingProvider: payroc.TrainingProviderPayroc.Ptr(),
        Shipping: &boarding.CreateTerminalOrderShipping{
            Preferences: &boarding.CreateTerminalOrderShippingPreferences{
                Method: boarding.CreateTerminalOrderShippingPreferencesMethodNextDay.Ptr(),
                SaturdayDelivery: payroc.Bool(
                    true,
                ),
            },
            Address: &boarding.CreateTerminalOrderShippingAddress{
                RecipientName: "Recipient Name",
                BusinessName: payroc.String(
                    "Company Ltd",
                ),
                AddressLine1: "1 Example Ave.",
                AddressLine2: payroc.String(
                    "Example Address Line 2",
                ),
                City: "Chicago",
                State: "Illinois",
                PostalCode: "60056",
                Email: "example@mail.com",
                Phone: payroc.String(
                    "2025550164",
                ),
            },
        },
        OrderItems: []*payroc.OrderItem{
            &payroc.OrderItem{
                Type: payroc.OrderItemTypeSolution,
                SolutionTemplateId: "Roc Services_DX8000",
                SolutionQuantity: payroc.Float64(
                    1,
                ),
                DeviceCondition: payroc.OrderItemDeviceConditionNew.Ptr(),
                SolutionSetup: &payroc.OrderItemSolutionSetup{
                    Timezone: payroc.SchemasTimezoneAmericaChicago.Ptr(),
                    IndustryTemplateId: payroc.String(
                        "Retail",
                    ),
                    GatewaySettings: &payroc.OrderItemSolutionSetupGatewaySettings{
                        MerchantPortfolioId: payroc.String(
                            "Company Ltd",
                        ),
                        MerchantTemplateId: payroc.String(
                            "Company Ltd Merchant Template",
                        ),
                        UserTemplateId: payroc.String(
                            "Company Ltd User Template",
                        ),
                        TerminalTemplateId: payroc.String(
                            "Company Ltd Terminal Template",
                        ),
                    },
                    ApplicationSettings: &payroc.OrderItemSolutionSetupApplicationSettings{
                        ClerkPrompt: payroc.Bool(
                            false,
                        ),
                        Security: &payroc.OrderItemSolutionSetupApplicationSettingsSecurity{
                            RefundPassword: payroc.Bool(
                                true,
                            ),
                            KeyedSalePassword: payroc.Bool(
                                false,
                            ),
                            ReversalPassword: payroc.Bool(
                                true,
                            ),
                        },
                    },
                    DeviceSettings: &payroc.OrderItemSolutionSetupDeviceSettings{
                        NumberOfMobileUsers: payroc.Float64(
                            2,
                        ),
                        CommunicationType: payroc.OrderItemSolutionSetupDeviceSettingsCommunicationTypeWifi.Ptr(),
                    },
                    BatchClosure: &payroc.OrderItemSolutionSetupBatchClosure{
                        Automatic: &payroc.AutomaticBatchClose{},
                    },
                    ReceiptNotifications: &payroc.OrderItemSolutionSetupReceiptNotifications{
                        EmailReceipt: payroc.Bool(
                            true,
                        ),
                        SmsReceipt: payroc.Bool(
                            false,
                        ),
                    },
                    Taxes: []*payroc.OrderItemSolutionSetupTaxesItem{
                        &payroc.OrderItemSolutionSetupTaxesItem{
                            TaxRate: 6,
                            TaxLabel: "Sales Tax",
                        },
                    },
                    Tips: &payroc.OrderItemSolutionSetupTips{
                        Enabled: payroc.Bool(
                            false,
                        ),
                    },
                    Tokenization: payroc.Bool(
                        true,
                    ),
                },
            },
        },
    }
client.Boarding.ProcessingAccounts.CreateTerminalOrder(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**trainingProvider:** `*payroc.TrainingProvider` 
    
</dd>
</dl>

<dl>
<dd>

**shipping:** `*boarding.CreateTerminalOrderShipping` — Object that contains the shipping details for the terminal order. If you don't provide a shipping address, we use the Doing Business As (DBA) address of the processing account.
    
</dd>
</dl>

<dl>
<dd>

**orderItems:** `[]*payroc.OrderItem` — Array of order items. Provide a minimum of 1 order item and a maximum of 10 order items.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingAccounts.ListProcessingTerminals(ProcessingAccountId) -> *core.PayrocPager[*payroc.PaginatedProcessingTerminals, []*payroc.ProcessingTerminal, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of processing terminals associated with a processing account.  

**Note:** If you want to view the details of a specific processing terminal and you have its processingTerminalId, use our [Retrieve Processing Terminal](https://docs.payroc.com/api/schema/boarding/processing-terminals/retrieve) method.  

To list the terminals for a processing account, you need its processingAccountId. If you don't have the processingAccountId, use our [List Merchant Platforms](https://docs.payroc.com/api/schema/boarding/merchant-platforms/list) method to search for a merchant platform and its processing accounts.   

Our gateway returns the following information for each processing terminal in the list:  

- Status indicating whether the terminal is active or inactive.  
- Configuration settings, including gateway settings and application settings.  
- Features, receipt settings, and security settings.  
- Devices that use the processing terminal's configuration.  

For each processing terminal, we also return its processingTerminalId, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.ListProcessingTerminalsProcessingAccountsRequest{
        ProcessingAccountId: "38765",
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.Boarding.ProcessingAccounts.ListProcessingTerminals(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingAccountId:** `string` — Unique identifier that we assigned to the processing account.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Boarding ProcessingTerminals
<details><summary><code>client.Boarding.ProcessingTerminals.Retrieve(ProcessingTerminalId) -> *payroc.ProcessingTerminal</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Important:** You can retrieve a processing terminal only if the terminal order was created using the Payroc API.

Use this method to retrieve information about a processing terminal.  

To retrieve a processing terminal, you need its processingTerminalId. Our gateway returned the processingTerminalId in the response of the [Create Terminal Order](https://docs.payroc.com/api/schema/boarding/processing-accounts/create-terminal-order) method.  

**Note:** If you don't have the processingTerminalId, use our [Retrieve Terminal Order](https://docs.payroc.com/api/schema/boarding/terminal-orders/retrieve) method or our [List Processing Terminals](https://docs.payroc.com/api/schema/boarding/processing-accounts/list-processing-terminals) method to search for the processing terminal.  

Our gateway returns the following information about the processing terminal:  

- Status indicating whether the terminal is active or inactive.  
- Configuration settings, including gateway settings and application settings.  
- Features, receipt settings, and security settings.  
- Devices that use the processing terminal's configuration.  
  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.RetrieveProcessingTerminalsRequest{
        ProcessingTerminalId: "1234001",
    }
client.Boarding.ProcessingTerminals.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.ProcessingTerminals.RetrieveHostConfiguration(ProcessingTerminalId) -> *payroc.HostConfiguration</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve the host processor configuration of a processing terminal. Integrate with this method only if you use your own gateway and want to validate the processor configuration.  

Our gateway returns the configuration settings for the merchant and the payment terminal.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.RetrieveHostConfigurationProcessingTerminalsRequest{
        ProcessingTerminalId: "1234001",
    }
client.Boarding.ProcessingTerminals.RetrieveHostConfiguration(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Boarding Contacts
<details><summary><code>client.Boarding.Contacts.Retrieve(ContactId) -> *payroc.Contact</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve details about a contact.  

To retrieve a contact, you need its contactId. Our gateway returned the contactId in the [Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) method.  

**Note:** If you don't have the contactId, use the [List Contacts](https://docs.payroc.com/api/schema/boarding/processing-accounts/list-contacts) method to search for the contact.  

Our gateway returns the following information about a contact:  

-	Name and contact method, including their phone number or mobile number.  
-	Role within the business, for example, if they are a manager. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.RetrieveContactsRequest{
        ContactId: 1,
    }
client.Boarding.Contacts.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**contactId:** `int` — Unique identifier for the contact.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.Contacts.Update(ContactId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to update a contact of a processing account.  

To update a contact, you need its contactId. Our gateway returned the contactId in the response of the [Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) method.  

**Note:** If you don't have the contactId, use our [List Contacts](https://docs.payroc.com/api/schema/boarding/processing-accounts/list-contacts) method to search for the contact.  

You can update the following details about a contact:  

-	First name and last name.  
-	Contact details, including their phone number or mobile number.  
-	Identification details, including their identification type and number.  
-	Role within the business, for example, if they are a manager.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.UpdateContactsRequest{
        ContactId: 1,
        Body: &payroc.Contact{
            Type: payroc.ContactTypeManager,
            FirstName: "Jane",
            MiddleName: payroc.String(
                "Helen",
            ),
            LastName: "Doe",
            Identifiers: []*payroc.Identifier{
                &payroc.Identifier{
                    Type: payroc.IdentifierTypeNationalId,
                    Value: "000-00-4320",
                },
            },
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
        },
    }
client.Boarding.Contacts.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**contactId:** `int` — Unique identifier for the contact.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.Contact` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Boarding.Contacts.Delete(ContactId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to delete a contact associated with a processing account.  

To delete a contact, you need their contactId. Our gateway returned the contactId in the response of the [Create Processing Account](https://docs.payroc.com/api/schema/boarding/merchant-platforms/create-processing-account) method.  

**Note:** If you don’t have the contactId, use our [Retrieve Processing Account](https://docs.payroc.com/api/schema/boarding/processing-accounts/retrieve) method or our [List Contacts](https://docs.payroc.com/api/schema/boarding/processing-accounts/list-contacts) method to search for the contact.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.DeleteContactsRequest{
        ContactId: 1,
    }
client.Boarding.Contacts.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**contactId:** `int` — Unique identifier for the contact.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Boarding TerminalOrders
<details><summary><code>client.Boarding.TerminalOrders.Retrieve(TerminalOrderId) -> *payroc.TerminalOrder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a terminal order.  

To retrieve a terminal order, you need it's terminalOrderId. Our gateway returned the terminalOrderId in the response of the [Create Terminal Order](https://docs.payroc.com/api/schema/boarding/processing-accounts/create-terminal-order) method.  

**Note**: If you don't have the terminalOrderId, use our [List Terminal Orders](https://docs.payroc.com/api/schema/boarding/processing-accounts/list-terminal-orders) method to search for the terminal order.  

Our gateway returns the following information about the terminal order:  
- Status of the order  
- Items in the order  
- Training provider  
- Shipping information  

**Note**: You can subscribe to our terminalOrder.status.changed event to get notifications when we update the status of a terminal order. For more information about how to subscribe to events, go to [Events Subscriptions](https://docs.payroc.com/guides/integrate/event-subscriptions).  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &boarding.RetrieveTerminalOrdersRequest{
        TerminalOrderId: "12345",
    }
client.Boarding.TerminalOrders.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**terminalOrderId:** `string` — Unique identifier of the terminal order.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CardPayments Payments
<details><summary><code>client.CardPayments.Payments.List() -> *core.PayrocPager[*payroc.PaymentPaginatedListForRead, []*payroc.RetrievedPayment, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of payments. 

**Note:** If you want to view the details of a specific payment and you have its paymentId, use our [Retrieve Payment](https://docs.payroc.com/api/schema/card-payments/payments/retrieve) method.  

Use query parameters to filter the list of results that we return, for example, to search for payments for a customer, a tip mode, or a date range.  

Our gateway returns the following information about each payment in the list:  

- Order details, including the transaction amount and when it was processed.  
- Payment card details, including the masked card number, expiry date, and payment method. 
- Cardholder details, including their contact information and shipping address. 
- Payment details, including the payment type, status, and response. 
 
For each transaction, we also return the paymentId and an optional secureTokenId, which you can use to perform follow-on actions. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.ListPaymentsRequest{
        ProcessingTerminalId: payroc.String(
            "1234001",
        ),
        OrderId: payroc.String(
            "OrderRef6543",
        ),
        Operator: payroc.String(
            "Jane",
        ),
        CardholderName: payroc.String(
            "Sarah%20Hazel%20Hopper",
        ),
        First6: payroc.String(
            "453985",
        ),
        Last4: payroc.String(
            "7062",
        ),
        Tender: cardpayments.ListPaymentsRequestTenderEbt.Ptr(),
        DateFrom: payroc.Time(
            payroc.MustParseDateTime(
                "2024-07-01T15:30:00Z",
            ),
        ),
        DateTo: payroc.Time(
            payroc.MustParseDateTime(
                "2024-07-03T15:30:00Z",
            ),
        ),
        SettlementState: cardpayments.ListPaymentsRequestSettlementStateSettled.Ptr(),
        SettlementDate: payroc.Time(
            payroc.MustParseDate(
                "2024-07-02",
            ),
        ),
        PaymentLinkId: payroc.String(
            "JZURRJBUPS",
        ),
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.CardPayments.Payments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `*string` — Filter by terminal ID.
    
</dd>
</dl>

<dl>
<dd>

**orderId:** `*string` — Filter payments by order ID.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Filter payments by operator.
    
</dd>
</dl>

<dl>
<dd>

**cardholderName:** `*string` — Filter payments by the cardholder’s name.
    
</dd>
</dl>

<dl>
<dd>

**first6:** `*string` — Filter payments by the first six digits of the card number that the customer used in the transaction.
    
</dd>
</dl>

<dl>
<dd>

**last4:** `*string` — Filter payments by the last four digits of the card number that the customer used in the transaction.
    
</dd>
</dl>

<dl>
<dd>

**tender:** `*cardpayments.ListPaymentsRequestTender` — Filter by tender type.
    
</dd>
</dl>

<dl>
<dd>

**tipMode:** `*cardpayments.ListPaymentsRequestTipModeItem` — Filter payments by tip.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*cardpayments.ListPaymentsRequestTypeItem` — Filter payments by transaction type.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*cardpayments.ListPaymentsRequestStatusItem` — Filter payments by the status of the transaction.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `*time.Time` — Filter by payments that the processor processed after a specific date. The date format follows the ISO 8601 standard.
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `*time.Time` — Filter by payments that the processer processed before a specific date. The date format follows the ISO 8601 standard.
    
</dd>
</dl>

<dl>
<dd>

**settlementState:** `*cardpayments.ListPaymentsRequestSettlementState` — Filter payments by the settlement status of the transaction.
    
</dd>
</dl>

<dl>
<dd>

**settlementDate:** `*time.Time` — Filter by payments that the processor settled on a specific date in the format **YYYY-MM-DD**.
    
</dd>
</dl>

<dl>
<dd>

**paymentLinkId:** `*string` — Unique identifier that our gateway assigned to the payment link.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Payments.Create(request) -> *payroc.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to run a sale or a pre-authorization with a customer's payment card. 

In the response, our gateway returns information about the card payment and a paymentId, which you need for the following methods:

-	[Retrieve payment](https://docs.payroc.com/api/schema/card-payments/payments/retrieve) - View the details of the card payment.
-	[Adjust payment](https://docs.payroc.com/api/schema/card-payments/payments/adjust) - Update the details of the card payment.
-	[Capture payment](https://docs.payroc.com/api/schema/card-payments/payments/capture)  - Capture the pre-authorization.
-	[Reverse payment](https://docs.payroc.com/api/schema/card-payments/refunds/reverse)  - Cancel the card payment if it's in an open batch.
-	[Refund payment](https://docs.payroc.com/api/schema/card-payments/refunds/create-referenced-refund)  - Run a referenced refund to return funds to the payment card.

**Payment methods** 

- **Cards** - Credit, debit, and EBT
- **Digital wallets** - [Apple Pay®](https://docs.payroc.com/guides/integrate/apple-pay) and [Google Pay®](https://docs.payroc.com/guides/integrate/google-pay) 
- **Tokens** - Secure tokens and single-use tokens

**Features** 

Our Create Payment method also supports the following features: 

- [Repeat payments](https://docs.payroc.com/guides/integrate/repeat-payments/use-your-own-software) - Run multiple payments as part of a payment schedule that you manage with your own software. 
- **Offline sales** - Run a sale or a pre-authorization if the terminal loses its connection to our gateway. 
- [Tokenization](https://docs.payroc.com/guides/integrate/save-payment-details) - Save card details to use in future transactions. 
- [3-D Secure](https://docs.payroc.com/guides/integrate/3-d-secure) - Verify the identity of the cardholder. 
- [Custom fields](https://docs.payroc.com/guides/integrate/add-custom-fields) - Add your own data to a payment. 
- **Tips** - Add tips to the card payment.  
- **Taxes** - Add local taxes to the card payment. 
- **Surcharging** - Add a surcharge to the card payment. 
- **Dual pricing** - Offer different prices based on payment method, for example, if you use our RewardPay Choice pricing program. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
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
client.CardPayments.Payments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**channel:** `*cardpayments.PaymentRequestChannel` — Channel that the merchant used to receive the payment details.
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who ran the transaction.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*payroc.PaymentOrderRequest` 
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.Customer` 
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `*payroc.IpAddress` 
    
</dd>
</dl>

<dl>
<dd>

**paymentMethod:** `*cardpayments.PaymentRequestPaymentMethod` — Object that contains information about the customer's payment details.
    
</dd>
</dl>

<dl>
<dd>

**threeDSecure:** `*cardpayments.PaymentRequestThreeDSecure` — Object that contains information for an authentication check on the customer's payment details using the 3-D Secure protocol.
    
</dd>
</dl>

<dl>
<dd>

**credentialOnFile:** `*payroc.SchemasCredentialOnFile` 
    
</dd>
</dl>

<dl>
<dd>

**offlineProcessing:** `*payroc.OfflineProcessing` 
    
</dd>
</dl>

<dl>
<dd>

**autoCapture:** `*bool` 

Indicates if we should automatically capture the payment amount.  

- `true` - Run a sale and automatically capture the transaction.
- `false`- Run a pre-authorization and capture the transaction later.  
  
**Note:** If you send `false` and the terminal doesn't support pre-authorization, we set the transaction's status to pending. The merchant must capture the transaction to take payment from the customer.
    
</dd>
</dl>

<dl>
<dd>

**processAsSale:** `*bool` 

Indicates if we should immediately settle the sale transaction. The merchant cannot adjust the transaction if we immediately settle it.  
**Note:** If the value for **processAsSale** is `true`, the gateway ignores the value in **autoCapture**.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*payroc.CustomField` — Array of customField objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Payments.Retrieve(PaymentId) -> *payroc.RetrievedPayment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a card payment.  

To retrieve a payment, you need its paymentId. Our gateway returned the paymentId in the response of the [Create Payment](https://docs.payroc.com/api/schema/card-payments/payments/create) method.  

**Note:** If you don't have the paymentId, use our [List Payments](https://docs.payroc.com/api/schema/card-payments/payments/list) method to search for the payment.  

Our gateway returns the following information about the payment:  

- Order details, including the transaction amount and when it was processed.  
- Payment card details, including the masked card number, expiry date, and payment method.  
- Cardholder details, including their contact information and shipping address.  
- Payment details, including the payment type, status, and response.  

If the merchant saved the customer's card details, our gateway returns a secureTokenID, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.RetrievePaymentsRequest{
        PaymentId: "M2MJOG6O2Y",
    }
client.CardPayments.Payments.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier of the payment that the merchant wants to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Payments.Adjust(PaymentId, request) -> *payroc.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to adjust a payment in an open batch. 

To adjust a payment, you need its paymentId. Our gateway returned the paymentId in the response of the [Create Payment](https://docs.payroc.com/api/schema/card-payments/payments/create) method.

**Note:** If you don't have the paymentId, use our [List Payments](https://docs.payroc.com/api/schema/card-payments/payments/list) method to search for the payment. 

You can adjust the following details of the payment:
- Sale amount and tip amount
- Payment status
- Cardholder shipping address and contact information
- Cardholder signature data

Our gateway returns information about the adjusted payment, including information about the payment card and the cardholder.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.PaymentAdjustment{
        PaymentId: "M2MJOG6O2Y",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Adjustments: []*cardpayments.PaymentAdjustmentAdjustmentsItem{
            &cardpayments.PaymentAdjustmentAdjustmentsItem{
                Customer: &payroc.CustomerAdjustment{},
            },
            &cardpayments.PaymentAdjustmentAdjustmentsItem{
                Order: &payroc.OrderAdjustment{
                    Amount: 4999,
                },
            },
        },
    }
client.CardPayments.Payments.Adjust(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier of the payment that the merchant wants to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who adjusted the payment.
    
</dd>
</dl>

<dl>
<dd>

**adjustments:** `[]*cardpayments.PaymentAdjustmentAdjustmentsItem` — Array of objects that contain information about the adjustments to the payment.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Payments.Capture(PaymentId, request) -> *payroc.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to capture a pre-authorization. 

To capture a pre-authorization, you need its paymentId. Our gateway returned the paymentId in the response of the [Create Payment](https://docs.payroc.com/api/schema/card-payments/payments/create) method.

**Note:** If you don't have the paymentId, use our [List Payments](https://docs.payroc.com/api/schema/card-payments/payments/list) method to search for the payment.

Depending on the amount you want to capture, complete the following:
-	**Capture the full amount of the pre-authorization** - Don't send a value for the amount parameter in your request.
-	**Capture less than the amount of the pre-authorization** - Send a value for the amount parameter in your request. 
-	**Capture more than the amount of the pre-authorization** - Adjust the pre-authorization before you capture it. For more information about adjusting a pre-authorization, go to [Adjust Payment](https://docs.payroc.com/api/schema/card-payments/payments/adjust).

If your request is successful, our gateway takes the amount from the payment card. 

**Note:** For more information about pre-authorizations and captures, go to [Run a pre-authorization](https://docs.payroc.com/guides/integrate/run-a-pre-authorization).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.PaymentCapture{
        PaymentId: "M2MJOG6O2Y",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        ProcessingTerminalId: payroc.String(
            "1234001",
        ),
        Operator: payroc.String(
            "Jane",
        ),
        Amount: payroc.Int64(
            4999,
        ),
        Breakdown: &payroc.ItemizedBreakdownRequest{
            Subtotal: 4999,
            DutyAmount: payroc.Int64(
                499,
            ),
            FreightAmount: payroc.Int64(
                500,
            ),
            Items: []*payroc.LineItemRequest{
                &payroc.LineItemRequest{
                    UnitPrice: 4000,
                    Quantity: 1,
                },
            },
        },
    }
client.CardPayments.Payments.Capture(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier of the payment that the merchant wants to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `*string` — Unique identifier that our gateway assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who captured the payment.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*int64` 

Amount that the merchant wants to capture. The value is in the currency's lowest denomination, for example, cents.  
**Note:** If the merchant does not send an amount, we capture the total amount of the transaction.
    
</dd>
</dl>

<dl>
<dd>

**breakdown:** `*payroc.ItemizedBreakdownRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CardPayments Refunds
<details><summary><code>client.CardPayments.Refunds.Reverse(PaymentId, request) -> *payroc.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to cancel or to partially cancel a payment in an open batch. This is also known as voiding a payment.  

To cancel a payment, you need its paymentId. Our gateway returned the paymentId in the response of the [Create Payment](https://docs.payroc.com/api/schema/card-payments/payments/create) method.  

**Note:** If you don't have the paymentId, use our [List Payments](https://docs.payroc.com/api/schema/card-payments/payments/list) method to search for the payment.  

If your request is successful, our gateway removes the payment from the merchant's open batch and no funds are taken from the cardholder's account. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.PaymentReversal{
        PaymentId: "M2MJOG6O2Y",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Amount: payroc.Int64(
            4999,
        ),
    }
client.CardPayments.Refunds.Reverse(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier of the payment that the merchant wants to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who reversed the payment.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*int64` 

Amount of the payment that the merchant wants to reverse. The value is in the currency’s lowest denomination, for example, cents.  
**Note:** If the merchant doesn’t send an amount, we reverse the total amount of the transaction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Refunds.CreateReferencedRefund(PaymentId, request) -> *payroc.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to refund a payment that is in a closed batch.  

To refund a payment, you need its paymentId. Our gateway returned the paymentId in the response of the [Create Payment](https://docs.payroc.com/api/schema/card-payments/payments/create) method.  

**Note:** If you don't have the paymentId, use our [List Payments](https://docs.payroc.com/api/schema/card-payments/payments/list) method to search for the payment.  

If your refund is successful, our gateway returns the payment amount to the cardholder's account.  

**Things to consider**  

- If the merchant refunds a payment that is in an open batch, our gateway reverses the payment.
- Some merchants can run unreferenced refunds, which means that they don't need a paymentId to return an amount to a customer. For more information about how to run an unreferenced refund, go to [Create Refund](https://docs.payroc.com/api/schema/card-payments/refunds/create-unreferenced-refund).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.ReferencedRefund{
        PaymentId: "M2MJOG6O2Y",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Amount: 4999,
        Description: "Refund for order OrderRef6543",
    }
client.CardPayments.Refunds.CreateReferencedRefund(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentId:** `string` — Unique identifier of the payment that the merchant wants to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who refunded the payment.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `int64` — Amount of the payment that the merchant wants to refund. The value is in the currency’s lowest denomination, for example, cents.
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` — Reason for the refund.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Refunds.List() -> *core.PayrocPager[*payroc.RefundPaginatedList, []*payroc.RetrievedRefund, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of refunds.  

**Note:** If you want to view the details of a specific refund and you have its refundId, use our [Retrieve Refund](https://docs.payroc.com/api/schema/card-payments/refunds/retrieve) method.  

Use query parameters to filter the list of results that we return, for example, to search for refunds for a customer, a tender type, or a date range.
Our gateway returns the following information about each refund in the list:  
- Order details, including the refund amount and when we processed the refund.
- Payment card details, including the masked card number, expiry date, and payment method.
- Cardholder details, including their contact information and shipping address.  

For referenced refunds, our gateway also returns details about the payment that the refund is linked to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.ListRefundsRequest{
        ProcessingTerminalId: payroc.String(
            "1234001",
        ),
        OrderId: payroc.String(
            "OrderRef6543",
        ),
        Operator: payroc.String(
            "Jane",
        ),
        CardholderName: payroc.String(
            "Sarah%20Hazel%20Hopper",
        ),
        First6: payroc.String(
            "453985",
        ),
        Last4: payroc.String(
            "7062",
        ),
        Tender: cardpayments.ListRefundsRequestTenderEbt.Ptr(),
        DateFrom: payroc.Time(
            payroc.MustParseDateTime(
                "2024-07-01T15:30:00Z",
            ),
        ),
        DateTo: payroc.Time(
            payroc.MustParseDateTime(
                "2024-07-03T15:30:00Z",
            ),
        ),
        SettlementState: cardpayments.ListRefundsRequestSettlementStateSettled.Ptr(),
        SettlementDate: payroc.Time(
            payroc.MustParseDate(
                "2024-07-02",
            ),
        ),
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.CardPayments.Refunds.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `*string` — Filter by terminal ID.
    
</dd>
</dl>

<dl>
<dd>

**orderId:** `*string` — Filter refunds by the unique identifier that the merchant assigned to the order.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Filter refunds by the operator who initiated the request.
    
</dd>
</dl>

<dl>
<dd>

**cardholderName:** `*string` — Filter refunds by cardholder name.
    
</dd>
</dl>

<dl>
<dd>

**first6:** `*string` — Filter refunds by the first six digits of the card number.
    
</dd>
</dl>

<dl>
<dd>

**last4:** `*string` — Filter refunds by the last four digits of the card number.
    
</dd>
</dl>

<dl>
<dd>

**tender:** `*cardpayments.ListRefundsRequestTender` — Filter by tender type.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*cardpayments.ListRefundsRequestStatusItem` — Filter refunds by the current status of the refund.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `*time.Time` — Filter by refunds processed after a specific date. The date format follows the ISO 8601 standard.
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `*time.Time` — Filter by refunds processed before a specific date. The date format follows the ISO 8601 standard.
    
</dd>
</dl>

<dl>
<dd>

**settlementState:** `*cardpayments.ListRefundsRequestSettlementState` — Status of the settlement.
    
</dd>
</dl>

<dl>
<dd>

**settlementDate:** `*time.Time` — Date the transaction was settled.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Refunds.CreateUnreferencedRefund(request) -> *payroc.RetrievedRefund</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create an unreferenced refund. An unreferenced refund is a refund that isn't linked to a payment.  

**Note:** If you have the paymentId of the payment you want to refund, use our [Refund Payment](https://docs.payroc.com/api/schema/card-payments/refunds/create-referenced-refund) method. If you use our Refund Payment method, our gateway sends the refund amount to the customer's original payment method and links the refund to the payment.  

In the request, you must provide the customer's payment details and the refund amount.  

In the response, our gateway returns information about the refund and a refundId, which you need for the following methods:  

- [Retrieve refund](https://docs.payroc.com/api/schema/card-payments/refunds/retrieve) - View the details of the refund.  
- [Adjust refund](https://docs.payroc.com/api/schema/card-payments/refunds/adjust) - Update the details of the refund.  
- [Reverse refund](https://docs.payroc.com/api/schema/card-payments/refunds/reverse-refund) - Cancel the refund if it's in an open batch.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.UnreferencedRefund{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Channel: cardpayments.UnreferencedRefundChannelPos,
        ProcessingTerminalId: "1234001",
        Order: &payroc.RefundOrder{
            OrderId: payroc.String(
                "OrderRef6543",
            ),
            Description: payroc.String(
                "Refund for order OrderRef6543",
            ),
            Amount: payroc.Int64(
                4999,
            ),
            Currency: payroc.CurrencyUsd.Ptr(),
        },
        RefundMethod: &cardpayments.UnreferencedRefundRefundMethod{
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
client.CardPayments.Refunds.CreateUnreferencedRefund(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**channel:** `*cardpayments.UnreferencedRefundChannel` — Channel that the merchant used to request the refund.
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who initiated the request.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*payroc.RefundOrder` 
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.Customer` 
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `*payroc.IpAddress` 
    
</dd>
</dl>

<dl>
<dd>

**refundMethod:** `*cardpayments.UnreferencedRefundRefundMethod` — Object that contains information about how the merchant refunds the customer.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*payroc.CustomField` — Array of customField objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Refunds.Retrieve(RefundId) -> *payroc.RetrievedRefund</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a refund.  

To retrieve a refund, you need its refundId. Our gateway returned the refundId in the response of the [Refund Payment](https://docs.payroc.com/api/schema/card-payments/refunds/create-referenced-refund) method or the [Create Refund](https://docs.payroc.com/api/schema/card-payments/refunds/create-unreferenced-refund) method.  

**Note:** If you don't have the refundId, use our [List Refunds](https://docs.payroc.com/api/schema/card-payments/refunds/list) method to search for the refund.  

Our gateway returns the following information about the refund:  
- Order details, including the refund amount and when we processed the refund.
- Payment card details, including the masked card number, expiry date, and payment method.
- Cardholder details, including their contact information and shipping address.  

If the refund is a referenced refund, our gateway also returns details about the payment that the refund is linked to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.RetrieveRefundsRequest{
        RefundId: "CD3HN88U9F",
    }
client.CardPayments.Refunds.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refundId:** `string` — Unique identifier that our gateway assigned to the refund.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Refunds.Adjust(RefundId, request) -> *payroc.RetrievedRefund</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to adjust a refund in an open batch.  

To adjust a refund, you need its refundId. Our gateway returned the refundId in the response of the [Refund Payment](https://docs.payroc.com/api/schema/card-payments/refunds/create-referenced-refund) method or the [Create Refund](https://docs.payroc.com/api/schema/card-payments/refunds/create-unreferenced-refund) method.  

**Note:** If you don’t have the refundId, use our [List Refunds](https://docs.payroc.com/api/schema/card-payments/refunds/list) method to search for the refund.  

You can adjust the following details of the refund:
- Customer details, including shipping address and contact information.
- Status of the refund.  

Our gateway returns information about the adjusted refund, including:
- Order details, including the refund amount and when we processed the refund.
- Payment card details, including the masked card number, expiry date, and payment method.
- Cardholder details, including their contact information and shipping address.  

If the refund is a referenced refund, our gateway also returns details about the payment that the refund is linked to.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.RefundAdjustment{
        RefundId: "CD3HN88U9F",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Operator: payroc.String(
            "Jane",
        ),
        Adjustments: []*cardpayments.RefundAdjustmentAdjustmentsItem{
            &cardpayments.RefundAdjustmentAdjustmentsItem{
                Customer: &payroc.CustomerAdjustment{},
            },
        },
    }
client.CardPayments.Refunds.Adjust(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refundId:** `string` — Unique identifier that our gateway assigned to the refund.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who requested the adjustment to the refund.
    
</dd>
</dl>

<dl>
<dd>

**adjustments:** `[]*cardpayments.RefundAdjustmentAdjustmentsItem` — Array of objects that contain information about the adjustments to the refund.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CardPayments.Refunds.ReverseRefund(RefundId) -> *payroc.RetrievedRefund</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to cancel a refund in an open batch.  

To cancel a refund, you need its refundId. Our gateway returned the refundId in the response of the [Refund Payment](https://docs.payroc.com/api/schema/card-payments/refunds/create-referenced-refund) or [Create Refund](https://docs.payroc.com/api/schema/card-payments/refunds/create-unreferenced-refund) method.  

**Note:** If you don’t have the refundId, use our [List Refunds](https://docs.payroc.com/api/schema/card-payments/refunds/list) method to search for the refund.  

If your request is successful, the gateway removes the refund from the merchant’s open batch and no funds are returned to the cardholder’s account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &cardpayments.ReverseRefundRefundsRequest{
        RefundId: "CD3HN88U9F",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
    }
client.CardPayments.Refunds.ReverseRefund(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refundId:** `string` — Unique identifier that our gateway assigned to the refund.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Funding FundingRecipients
<details><summary><code>client.Funding.FundingRecipients.List() -> *core.PayrocPager[*payroc.PaginatedFundRecipients, []*payroc.FundingRecipient, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of funding recipients linked to your account.  

Note: If you want to view the details of a specific funding recipient and you have its recipientId, use our [Retrieve Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/retrieve) method.  

Our gateway returns the following information about each funding recipient in the list:  
- Tax ID and Doing Business As (DBA) name.  
- Address and contact details.  
- Funding accounts linked to the funding recipient.  

For each funding recipient, we also return the recipientId, which you can use to perform follow-on actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.ListFundingRecipientsRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.Funding.FundingRecipients.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingRecipients.Create(request) -> *payroc.FundingRecipient</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a funding recipient. 

A funding recipient is a business or organization that can receive funds but can't run transactions, for example, a charity.  

In the request, include the following information:  
-	Legal information, including its tax ID, Doing Business As (DBA) name, and address.  
-	Contact information, including the email address.  
-	Owners' details, including their contact details. 
-	Funding account details.  

Our gateway returns the recipientId of the funding recipient, which you can use to run follow-on actions. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.CreateFundingRecipient{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        RecipientType: funding.CreateFundingRecipientRecipientTypePrivateCorporation,
        TaxId: "123456789",
        DoingBusinessAs: "doingBusinessAs",
        Address: &payroc.Address{
            Address1: "1 Example Ave.",
            City: "Chicago",
            State: "Illinois",
            Country: "US",
            PostalCode: "60056",
        },
        ContactMethods: []*payroc.ContactMethod{
            &payroc.ContactMethod{
                Email: &payroc.ContactMethodEmail{
                    Value: "jane.doe@example.com",
                },
            },
        },
        Owners: []*payroc.Owner{
            &payroc.Owner{
                FirstName: "Jane",
                LastName: "Doe",
                DateOfBirth: payroc.MustParseDate(
                    "1964-03-22",
                ),
                Address: &payroc.Address{
                    Address1: "1 Example Ave.",
                    City: "Chicago",
                    State: "Illinois",
                    Country: "US",
                    PostalCode: "60056",
                },
                Identifiers: []*payroc.Identifier{
                    &payroc.Identifier{
                        Type: payroc.IdentifierTypeNationalId,
                        Value: "xxxxx4320",
                    },
                },
                ContactMethods: []*payroc.ContactMethod{
                    &payroc.ContactMethod{
                        Email: &payroc.ContactMethodEmail{
                            Value: "jane.doe@example.com",
                        },
                    },
                },
                Relationship: &payroc.OwnerRelationship{
                    IsControlProng: true,
                },
            },
        },
        FundingAccounts: []*payroc.FundingAccount{
            &payroc.FundingAccount{
                Type: payroc.FundingAccountTypeChecking,
                Use: payroc.FundingAccountUseCredit,
                NameOnAccount: "Jane Doe",
                PaymentMethods: []*payroc.PaymentMethodsItem{
                    &payroc.PaymentMethodsItem{
                        Ach: &payroc.PaymentMethodAch{},
                    },
                },
            },
        },
    }
client.Funding.FundingRecipients.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**recipientType:** `*funding.CreateFundingRecipientRecipientType` — Type or legal structure of the funding recipient.
    
</dd>
</dl>

<dl>
<dd>

**taxId:** `string` — Employer identification number (EIN) or Social Security number (SSN).
    
</dd>
</dl>

<dl>
<dd>

**charityId:** `*string` — Government identifier of the charity.
    
</dd>
</dl>

<dl>
<dd>

**doingBusinessAs:** `string` — Trading name of the business or organization.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*payroc.Address` — Address of the funding recipient.
    
</dd>
</dl>

<dl>
<dd>

**contactMethods:** `[]*payroc.ContactMethod` — Array of contactMethod objects that you can use to add contact methods for the funding recipient. You must provide at least an email address.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]string` — [Metadata](https://docs.payroc.com/api/metadata) object you can use to include custom data with your request.
    
</dd>
</dl>

<dl>
<dd>

**owners:** `[]*payroc.Owner` — Array of owner objects. Each object contains information about an individual who owns or manages the funding recipient.
    
</dd>
</dl>

<dl>
<dd>

**fundingAccounts:** `[]*payroc.FundingAccount` — Array of fundingAccount objects that you can use to add funding accounts to the funding recipient.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingRecipients.Retrieve(RecipientId) -> *payroc.FundingRecipient</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a funding recipient.  

To retrieve a funding recipient, you need its recipientId. Our gateway returned the recipientId in the response of the [Create Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/create) method.  

**Note:** If you don't have the recipientId, use our [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient.  

Our gateway returns the following information about the funding recipient:  

- Tax ID and Doing Business As (DBA) name.  
- Address and contact details.  
- Funding accounts linked to the funding recipient.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.RetrieveFundingRecipientsRequest{
        RecipientId: 1,
    }
client.Funding.FundingRecipients.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**recipientId:** `int` — Unique identifier that we assigned to the funding recipient.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingRecipients.Update(RecipientId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to update the details of a funding recipient. If a request contains significant changes, we might need to re-approve the funding recipient.  

To update a funding recipient, you need it's recipientId. Our gateway returned the recipientId in the response of the [Create Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/create) method.  

**Note**: If you don't have the recipientId, use our [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient.  

You can update the following details of a funding recipient:  
- Doing Business As (DBA) name  
- Tax ID and charity ID  
- Address and contact methods  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.UpdateFundingRecipientsRequest{
        RecipientId: 1,
        Body: &payroc.FundingRecipient{
            RecipientType: payroc.FundingRecipientRecipientTypePrivateCorporation,
            TaxId: "123456789",
            DoingBusinessAs: "doingBusinessAs",
            Address: &payroc.Address{
                Address1: "1 Example Ave.",
                City: "Chicago",
                State: "Illinois",
                Country: "US",
                PostalCode: "60056",
            },
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
        },
    }
client.Funding.FundingRecipients.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**recipientId:** `int` — Unique identifier that we assigned to the funding recipient.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.FundingRecipient` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingRecipients.Delete(RecipientId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to delete a funding recipient, including its funding accounts and owners.  

To delete a funding recipient, you need its recipientId. Our gateway returned the recipientId in the response of the [Create Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/create) method.   

**Note**: If you don't have the recipientId, use our [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.DeleteFundingRecipientsRequest{
        RecipientId: 1,
    }
client.Funding.FundingRecipients.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**recipientId:** `int` — Unique identifier that we assigned to the funding recipient.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingRecipients.ListAccounts(RecipientId) -> []*payroc.FundingAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use  this method to return a list of funding accounts associated with a funding recipient.  

**Note:** If you want to view the details of a specific funding account and you have its fundingAccountId, use our [Retrieve Funding Account](https://docs.payroc.com/api/schema/funding/funding-accounts/retrieve) method.  

To retrieve the funding accounts associated with a funding recipient, you need the recipientId. If you don't have the recipientId, use our [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient.  
        
Our gateway returns the following information about each funding account:  
-	Name of the account holder.  
-	ACH details for the account.  
-	Status of the account.  

Our gateway also returns the fundingAccountId, which you can use to run follow-on actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.ListFundingRecipientFundingAccountsRequest{
        RecipientId: 1,
    }
client.Funding.FundingRecipients.ListAccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**recipientId:** `int` — Unique identifier that we assigned to the funding recipient.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingRecipients.CreateAccount(RecipientId, request) -> *payroc.FundingAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a funding account and add it to a funding recipient.  

To add a funding account to a funding recipient, you need the recipientId. Our gateway returned the recipientId in the response of the [Create Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/create) method.  

**Note:** If you don't have the recipientId, use our [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient.  

In the request, include the following information:  
-	Account type, for example, if the account is a checking or savings account.  
-	Account holder's name.  
-	ACH information, including the routing number and account number of the account.  

Our gateway returns the fundingAccountId, which you can use to run follow-on actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.CreateAccountFundingRecipientsRequest{
        RecipientId: 1,
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.FundingAccount{
            Type: payroc.FundingAccountTypeChecking,
            Use: payroc.FundingAccountUseCredit,
            NameOnAccount: "Jane Doe",
            PaymentMethods: []*payroc.PaymentMethodsItem{
                &payroc.PaymentMethodsItem{
                    Ach: &payroc.PaymentMethodAch{},
                },
            },
        },
    }
client.Funding.FundingRecipients.CreateAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**recipientId:** `int` — Unique identifier that we assigned to the funding recipient.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.FundingAccount` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingRecipients.ListOwners(RecipientId) -> []*payroc.Owner</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a list of owners of a funding recipient.  

**Note:** If you want to view the details of a specific owner and you have their ownerId, use our [Retrieve Owner](https://docs.payroc.com/api/schema/boarding/owners/retrieve) method.  

To list the owners of a funding recipient, you need its recipientId. Our gateway returned the recipientId in the response of the [Create Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/create) method. If you don't have the recipientId, use our [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient.  

Our gateway returns the following information about each owner in the list:  
-	Name, date of birth, and address.  
-	Contact details, including their email address.  
-	Relationship to the funding recipient.  

Our gateway also returns the ownerId, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.ListFundingRecipientOwnersRequest{
        RecipientId: 1,
    }
client.Funding.FundingRecipients.ListOwners(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**recipientId:** `int` — Unique identifier that we assigned to the funding recipient.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingRecipients.CreateOwner(RecipientId, request) -> *payroc.Owner</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to add an additional owner to a funding recipient.  

To add an owner to a funding recipient, you need the recipientId. Our gateway returned the recipientId in the response of the [Create Funding Recipient](https://docs.payroc.com/api/schema/funding/funding-recipients/create) method.  

**Note:** If you don't have the recipientId, use our [List Funding Recipients](https://docs.payroc.com/api/schema/funding/funding-recipients/list) method to search for the funding recipient.  

In the request, include the following information about the owner:  

- Name, date of birth, and address.  
- Contact details, including their email address.  
- Relationship to the funding recipient, including whether they are a control prong.  

In the response, our gateway returns the ownerId, which you can use to run follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.CreateOwnerFundingRecipientsRequest{
        RecipientId: 1,
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.Owner{
            FirstName: "Jane",
            LastName: "Doe",
            DateOfBirth: payroc.MustParseDate(
                "1964-03-22",
            ),
            Address: &payroc.Address{
                Address1: "1 Example Ave.",
                City: "Chicago",
                State: "Illinois",
                Country: "US",
                PostalCode: "60056",
            },
            Identifiers: []*payroc.Identifier{
                &payroc.Identifier{
                    Type: payroc.IdentifierTypeNationalId,
                    Value: "xxxxx4320",
                },
            },
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
            Relationship: &payroc.OwnerRelationship{
                IsControlProng: true,
            },
        },
    }
client.Funding.FundingRecipients.CreateOwner(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**recipientId:** `int` — Unique identifier that we assigned to the funding recipient.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.Owner` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Funding FundingAccounts
<details><summary><code>client.Funding.FundingAccounts.List() -> *core.PayrocPager[*payroc.ListFundingAccounts, []*payroc.FundingAccount, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of funding accounts associated with your account.  

**Note:** If you want to view the details of a specific funding account and you have its fundingAccountId, use our [Retrieve Funding Account](https://docs.payroc.com/api/schema/funding/funding-accounts/retrieve) method.  

Our gateway returns the following information about each funding account in the list:  
- Name of the account holder and ACH details for the account.  
- Status of the account.  
- Whether we send funds to the account, withdraw funds from the account, or both.  

For each funding account, we also return the fundingAccountId, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.ListFundingAccountsRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.Funding.FundingAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingAccounts.Retrieve(FundingAccountId) -> *payroc.FundingAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a funding account.  

To retrieve a funding account, you need its fundingAccountId. Our gateway returned the fundingAccountId when you created the funding account.  

**Note:** If you don't have the fundingAccountId, use our [List Funding Accounts](https://docs.payroc.com/api/schema/funding/funding-accounts/list) method to search for the account.  

Our gateway returns the following information about the funding account:  
- Name of the account holder and ACH details for the account.  
- Status of the account.  
- Whether we send funds to the account, withdraw funds from the account, or both.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.RetrieveFundingAccountsRequest{
        FundingAccountId: 1,
    }
client.Funding.FundingAccounts.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fundingAccountId:** `int` — Unique identifier of the funding account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingAccounts.Update(FundingAccountId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Important:** You can't update the details of a funding account that is associated with a processing account.  

Use this method to update the details of a funding account that is associated with a funding recipient.  

To update a funding account, you need its fundingAccountId. Our gateway returned the fundingAccountId when you created the funding account.  

**Note:** If you don’t have the fundingAccountId, use our [List Funding Accounts](https://docs.payroc.com/api/schema/funding/funding-accounts/list) method to search for the funding account.  

You can update the following details about the funding account: 
-	Account type. 
-	Account holder's name. 
-	ACH information for the account. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.UpdateFundingAccountsRequest{
        FundingAccountId: 1,
        Body: &payroc.FundingAccount{
            Type: payroc.FundingAccountTypeChecking,
            Use: payroc.FundingAccountUseCredit,
            NameOnAccount: "Jane Doe",
            PaymentMethods: []*payroc.PaymentMethodsItem{
                &payroc.PaymentMethodsItem{
                    Ach: &payroc.PaymentMethodAch{},
                },
            },
        },
    }
client.Funding.FundingAccounts.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fundingAccountId:** `int` — Unique identifier of the funding account.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.FundingAccount` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingAccounts.Delete(FundingAccountId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Important:** You can't delete a funding account that is associated with a processing account.  

Use this method to delete a funding account that is associated with a funding recipient.  

To delete a funding account, you need its fundingAccountId. Our gateway returned the fundingAccountId when you created the funding account.  

**Note:** If you don't have the fundingAccountId, use our [List Funding Accounts](https://docs.payroc.com/api/schema/funding/funding-accounts/list) method to search for the funding account.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.DeleteFundingAccountsRequest{
        FundingAccountId: 1,
    }
client.Funding.FundingAccounts.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fundingAccountId:** `int` — Unique identifier of the funding account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Funding FundingInstructions
<details><summary><code>client.Funding.FundingInstructions.List() -> *core.PayrocPager[*funding.ListFundingInstructionsResponse, []*funding.ListFundingInstructionsResponseDataItem, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> Important: You can return a list of funding instructions from only the previous two years. If you want to view a funding instruction from more than two years ago and you have its instructionId, use our [Retrieve Funding Instruction](https://docs.payroc.com/api/schema/funding/funding-instructions/retrieve) method.  

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of funding instructions within a specific date range.  

**Note:** If you want to view the details of a specific funding instruction and you have its instructionId, use our [Retrieve Funding Instruction](https://docs.payroc.com/api/schema/funding/funding-instructions/retrieve) method.  

Our gateway returns the following information for each instruction in the list:  
-	Status of the funding instruction.  
-	Funding information, including which merchant's funding balance we distribute and the funding account that we send the balance to.  

For each funding instruction, we also return the instructionId, which you can use to perform follow-on actions. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.ListFundingInstructionsRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        DateFrom: payroc.MustParseDate(
            "2024-07-01",
        ),
        DateTo: payroc.MustParseDate(
            "2024-07-03",
        ),
    }
client.Funding.FundingInstructions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `time.Time` — Filter by funding instructions sent after a specific date. Send a value in **YYYY-MM-DD** format.
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `time.Time` — Filter by funding instructions sent before a specific date. Send a value in **YYYY-MM-DD** format.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingInstructions.Create(request) -> *payroc.Instruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a funding instruction that tells us how to distribute the funds from your merchants' transactions.  

**Note:** Before you create a funding instruction, you can use our [List Funding Balances](https://docs.payroc.com/api/schema/funding/funding-activity/retrieve-balance) method to view the amount of available funds that a merchant has.  

In your request, include an array of merchantInstruction objects. Each merchantInstruction object contains the following:  
-	Merchant ID (MID) of the merchant whose funding balance you want to distribute.  
-	Funding account that you want to send funds to.  
-	Amount that you want to send to the funding account.  

Our gateway returns the instructionId, which you can use to run follow-on actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.CreateFundingInstructionsRequest{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.Instruction{},
    }
client.Funding.FundingInstructions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.Instruction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingInstructions.Retrieve(InstructionId) -> *payroc.Instruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a funding instruction.  

To retrieve a funding instruction, you need its instructionId. Our gateway returned the instructionId in the response of the [Create Funding Instruction](https://docs.payroc.com/api/schema/funding/funding-instructions/create) method. 

**Note:** If you don't have the instructionId, use our [List Funding Instructions](https://docs.payroc.com/api/schema/funding/funding-instructions/list) method to search for the funding instruction.  

Our gateway returns the following information about the funding instruction:  
-	Status of the funding instruction.  
-	Funding information, including which merchant's funding balance we distribute and the funding account that we send the balance to.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.RetrieveFundingInstructionsRequest{
        InstructionId: 1,
    }
client.Funding.FundingInstructions.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**instructionId:** `int` — Unique identifier that we assigned to the funding instruction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingInstructions.Update(InstructionId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Important:** You can update a funding instruction only if its status is `accepted`. To view the status of a funding instruction, use our [Retrieve Funding Instruction](https://docs.payroc.com/api/schema/funding/funding-instructions/retrieve) method. 

Use this method to update the details of a funding instruction.  

To update a funding instruction, you need its instructionId. Our gateway returned the instructionId in the response of the [Create Funding Instruction](https://docs.payroc.com/api/schema/funding/funding-instructions/create) method.  

**Note:** If you don't have the fundingInstructionId, use our [List Funding Instructions](https://docs.payroc.com/api/schema/funding/funding-instructions/list) method to search for the funding instruction.  

You can modify the following information for the funding instruction:  
-	Merchant ID (MID) of the merchant whose funding balance you want to distribute.  
-	Funding account that you want to send funds to.  
-	Amount that you want to send to the funding account.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.UpdateFundingInstructionsRequest{
        InstructionId: 1,
        Body: &payroc.Instruction{},
    }
client.Funding.FundingInstructions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**instructionId:** `int` — Unique identifier that we assigned to the funding instruction.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.Instruction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingInstructions.Delete(InstructionId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Important:** You can delete a funding instruction only if its status is `accepted`. To view the status of a funding instruction, use our [Retrieve Funding Instruction](https://docs.payroc.com/api/schema/funding/funding-instructions/retrieve) method.  

Use this method to delete a funding instruction.  

To delete a funding instruction, you need its instructionId. Our gateway returned the instructionId in the response of the [Create Funding Instruction](https://docs.payroc.com/api/schema/funding/funding-instructions/create) method.  

**Note:** If you don't have the instructionId, use our [List Funding Instructions](https://docs.payroc.com/api/schema/funding/funding-instructions/list) method to search for the funding instruction.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.DeleteFundingInstructionsRequest{
        InstructionId: 1,
    }
client.Funding.FundingInstructions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**instructionId:** `int` — Unique identifier that we assigned to the funding instruction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Funding FundingActivity
<details><summary><code>client.Funding.FundingActivity.RetrieveBalance() -> *funding.RetrieveBalanceFundingActivityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of funding balances available for each merchant linked to your account.  

Use query parameters to filter the list of results we return, for example, to search for the funding balance for a specific merchant.  

Our gateway returns the following information about each merchant in the list:  
- Total funds for the merchant.  
- Available funds that you can use for funding instructions.  
- Pending funds that we have not yet sent to funding accounts.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.RetrieveBalanceFundingActivityRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        MerchantId: payroc.String(
            "4525644354",
        ),
    }
client.Funding.FundingActivity.RetrieveBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**merchantId:** `*string` — Filter results by the unique identifier that the processor assigned to the merchant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Funding.FundingActivity.List() -> *core.PayrocPager[*funding.ListFundingActivityResponse, []*payroc.ActivityRecord, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of activity associated with your merchants' funding balances within a specific date range.  

Use query parameters to filter the list of results we return, for example, to view the activity for a specific merchant's funding balance.  

Our gateway returns the following information about each activity in the list:
- Name of the merchant who owns the funding balance.
-	Amount of funds added or removed from the funding balance.
-	Funding account that received funds from the funding balance.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &funding.ListFundingActivityRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        DateFrom: payroc.MustParseDate(
            "2024-07-02",
        ),
        DateTo: payroc.MustParseDate(
            "2024-07-03",
        ),
        MerchantId: payroc.String(
            "4525644354",
        ),
    }
client.Funding.FundingActivity.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `time.Time` — Filter by activity after a specific date. Send a value in **YYYY-MM-DD** format.
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `time.Time` — Filter by activity before a specific date. Send a value in **YYYY-MM-DD** format.
    
</dd>
</dl>

<dl>
<dd>

**merchantId:** `*string` — Filter results by the unique identifier that the processor assigned to the merchant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Notifications EventSubscriptions
<details><summary><code>client.Notifications.EventSubscriptions.List() -> *core.PayrocPager[*payroc.PaginatedEventSubscriptions, []*payroc.EventSubscription, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of event subscriptions that are linked to your ISV account.  

**Note:** If you want to view the details of a specific event subscription and you have its id, use our [Retrieve Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/retrieve) method.  

Use query parameters to filter the list of results that we return, for example, to search for subscriptions with a specific status or an event type.  

Our gateway returns the following information about each subscription in the list:  
- Event types that you have subscribed to.  
- Whether you have enabled notifications for the subscription.  
- How we contact you when an event occurs, including the endpoint that send notifications to.  
- If there are any issues when we try to send you a notification, for example, if we can't contact your endpoint.  

For each event subscription, we also return its id, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.ListEventSubscriptionsRequest{
        Status: notifications.ListEventSubscriptionsRequestStatusRegistered.Ptr(),
        Event: payroc.String(
            "processingAccount.status.changed",
        ),
    }
client.Notifications.EventSubscriptions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**status:** `*notifications.ListEventSubscriptionsRequestStatus` — Filter event subscriptions by subscription status.
    
</dd>
</dl>

<dl>
<dd>

**event:** `*string` — Filter event subscriptions by an event type. For a list of event types, go to [Events List](https://docs.payroc.com/knowledge/events/events-list).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.EventSubscriptions.Create(request) -> *payroc.EventSubscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create an event subscription that we use to notify you when an event occurs, for example, when we change the status of a processing account.  

In the request, include the events that you want to subscribe to and the public endpoint that we send event notifications to. For a complete list of events that you can subscribe to, go to [Events List](https://docs.payroc.com/knowledge/events/events-list).  

In the response, our gateway returns the id of the event subscription, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.CreateEventSubscriptionsRequest{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.EventSubscription{
            Enabled: true,
            EventTypes: []string{
                "processingAccount.status.changed",
            },
            Notifications: []*payroc.Notification{
                &payroc.Notification{
                    Webhook: &payroc.Webhook{
                        Uri: "https://my-server/notification/endpoint",
                        Secret: "aBcD1234eFgH5678iJkL9012mNoP3456",
                        SupportEmailAddress: "supportEmailAddress",
                    },
                },
            },
            Metadata: map[string]any{
                "yourCustomField": "abc123",
            },
        },
    }
client.Notifications.EventSubscriptions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.EventSubscription` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.EventSubscriptions.Retrieve(SubscriptionId) -> *payroc.EventSubscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve the details of an event subscription.  

In your request, include the subscriptionId that we sent to you when we created the event subscription.  
  
**Note:** If you don't know the subscriptionId of the event subscription, go to [List event subscriptions](#listEventSubscriptions).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.RetrieveEventSubscriptionsRequest{
        SubscriptionId: 1,
    }
client.Notifications.EventSubscriptions.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionId:** `int` 

Unique identifier that we assigned to the event subscription.  
**Note:** Our gateway returned the subscriptionId in the id field in the response of the [Create Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/create) method.  
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.EventSubscriptions.Update(SubscriptionId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to update the details of an event subscription.  

To update an event subscription, you need its subscriptionId. Our gateway returned the subscriptionId in the response of the [Create Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/create) method.  

**Note:** If you don’t have the subscriptionId, use our [List Event Subscriptions](https://docs.payroc.com/api/schema/notifications/event-subscriptions/list) method to search for the event subscription.  

You can update the following details about an event subscription:  

- Status of the event subscription.  
- Events that you have subscribed to. For a list of events that you can subscribe to, go to [Events list](https://docs.payroc.com/knowledge/events/events-list).  
- Information about how we contact you when an event occurs.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.UpdateEventSubscriptionsRequest{
        SubscriptionId: 1,
        Body: &payroc.EventSubscription{
            Enabled: true,
            EventTypes: []string{
                "processingAccount.status.changed",
            },
            Notifications: []*payroc.Notification{
                &payroc.Notification{
                    Webhook: &payroc.Webhook{
                        Uri: "https://my-server/notification/endpoint",
                        Secret: "aBcD1234eFgH5678iJkL9012mNoP3456",
                        SupportEmailAddress: "supportEmailAddress",
                    },
                },
            },
            Metadata: map[string]any{
                "yourCustomField": "abc123",
            },
        },
    }
client.Notifications.EventSubscriptions.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionId:** `int` 

Unique identifier that we assigned to the event subscription.  
**Note:** Our gateway returned the subscriptionId in the id field in the response of the [Create Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/create) method.  
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.EventSubscription` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.EventSubscriptions.Delete(SubscriptionId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to delete an event subscription.  

> **Important:** After you delete an event subscription, you can’t recover it. You won't receive event notifications from the event subscription.  

To delete an event subscription, you need its subscriptionId. Our gateway returned the subscriptionId in the response of the [Create Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/create) method.  

If you want to stop receiving event notifications but don't want to delete the event subscription, use our [Update Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/update) method to deactivate it.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.DeleteEventSubscriptionsRequest{
        SubscriptionId: 1,
    }
client.Notifications.EventSubscriptions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionId:** `int` 

Unique identifier that we assigned to the event subscription.  
**Note:** Our gateway returned the subscriptionId in the id field in the response of the [Create Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/create) method.  
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.EventSubscriptions.PartiallyUpdate(SubscriptionId, request) -> *payroc.EventSubscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to partially update an event subscription. Structure your request to follow the [RFC 6902](https://datatracker.ietf.org/doc/html/rfc6902) standard.  

To update an event subscription, you need its subscriptionId. Our gateway returned the subscriptionId in the id field in the response of the [Create Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/create) method.  

**Note:** If you don't have the subscriptionId, use our [List Event Subscriptions](https://docs.payroc.com/api/schema/notifications/event-subscriptions/list) method to search for the subscription.  

You can update the following properties of an event subscription:  
- **eventTypes** - Subscribe to new events or remove events that you are subscribed to.  
- **notifications** - Information about your endpoint and who we email if we can't contact your endpoint.  
- **enabled** - Turn on or turn off notifications for the subscription.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.PartiallyUpdateEventSubscriptionsRequest{
        SubscriptionId: 1,
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: []*payroc.PatchDocument{
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
        },
    }
client.Notifications.EventSubscriptions.PartiallyUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**subscriptionId:** `int` 

Unique identifier that we assigned to the event subscription.  
**Note:** Our gateway returned the subscriptionId in the id field in the response of the [Create Event Subscription](https://docs.payroc.com/api/schema/notifications/event-subscriptions/create) method.  
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `payroc.PatchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PaymentFeatures Cards
<details><summary><code>client.PaymentFeatures.Cards.VerifyCard(request) -> *payroc.CardVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to verify a customer’s card details.  

In the request, send the customer’s card details.  

In the response, our gateway indicates if the card details are valid and if you should use them in follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &paymentfeatures.CardVerificationRequest{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        ProcessingTerminalId: "1234001",
        Operator: payroc.String(
            "Jane",
        ),
        Card: &paymentfeatures.CardVerificationRequestCard{
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
    }
client.PaymentFeatures.Cards.VerifyCard(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who requested to verify the card.
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.Customer` 
    
</dd>
</dl>

<dl>
<dd>

**card:** `*paymentfeatures.CardVerificationRequestCard` — Object that contains information about the card.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentFeatures.Cards.ViewEbtBalance(request) -> *payroc.Balance</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to view the balance of an Electronic Benefit Transfer (EBT) card.  

If the request is successful, our gateway returns the current balance of an EBT card. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &paymentfeatures.BalanceInquiry{
        ProcessingTerminalId: "1234001",
        Operator: payroc.String(
            "Jane",
        ),
        Currency: payroc.CurrencyUsd,
        Card: &paymentfeatures.BalanceInquiryCard{
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
    }
client.PaymentFeatures.Cards.ViewEbtBalance(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who requested the balance inquiry.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*payroc.Currency` 
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.Customer` 
    
</dd>
</dl>

<dl>
<dd>

**card:** `*paymentfeatures.BalanceInquiryCard` — Object that contains information about the card.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentFeatures.Cards.LookupBin(request) -> *payroc.CardInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a debit card, a credit card, or an EBT card. If you apply surcharges to transactions, you can also check if the card supports surcharging.  

In the response, our gateway returns the following information about the card:  

- **Card details** - Information about the card, for example, the issuing bank and the masked card number.  

- **Surcharging information** - If you apply a surcharge to transactions, our gateway checks that the card supports surcharging and returns information about the surcharge. For more information about surcharging, go to [Credit card surcharging](https://docs.payroc.com/knowledge/card-payments/credit-card-surcharging). 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &paymentfeatures.BinLookup{
        ProcessingTerminalId: payroc.String(
            "1234001",
        ),
        Card: &paymentfeatures.BinLookupCard{
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
    }
client.PaymentFeatures.Cards.LookupBin(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `*string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*int64` — Transaction amount that you send to check the surcharge amount. The value is in the currency's lowest denomination, for example, cents.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*payroc.Currency` 
    
</dd>
</dl>

<dl>
<dd>

**card:** `*paymentfeatures.BinLookupCard` — Object that contains information about the card.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentFeatures.Cards.RetrieveFxRates(request) -> *payroc.FxRate</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

> **Important:** There are restrictions on which merchants can use this method. For more information, go to [Dynamic Currency Conversion](https://docs.payroc.com/knowledge/card-payments/dynamic-currency-conversion).  

Use this method to check if a card is eligible for Dynamic Currency Conversion (DCC) and to retrieve the conversion rate for a transaction amount. DCC provides a customer with the option to use their card's currency instead of the merchant's currency, for example, in Ireland, an American customer can pay in US dollars instead of Euros.  

The request includes the following:  

- **Payment method** - Card information, a secure token, or digital wallet.  
- **Transaction information** - Amount and currency of the transaction in the merchant's currency.  

If the card is eligible for DCC, our gateway returns the transaction amount in the card's currency and a dccOffer object that contains information about the conversion rate. The dccOffer object contains the following fields that you need when you [run a sale](https://docs.payroc.com/api/schema/card-payments/payments/create) or [unreferenced refund](https://docs.payroc.com/api/schema/card-payments/refunds/create-unreferenced-refund) with DCC:  
- fxAmount  
- fxCurrency  
- fxRate  
- markup  
- accepted  
- offerReference  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &paymentfeatures.FxRateInquiry{
        Channel: paymentfeatures.FxRateInquiryChannelWeb,
        ProcessingTerminalId: "1234001",
        Operator: payroc.String(
            "Jane",
        ),
        BaseAmount: 10000,
        BaseCurrency: payroc.CurrencyUsd,
        PaymentMethod: &paymentfeatures.FxRateInquiryPaymentMethod{
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
    }
client.PaymentFeatures.Cards.RetrieveFxRates(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**channel:** `*paymentfeatures.FxRateInquiryChannel` — Channel that the merchant used to receive payment details for the transaction.
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who ran the transaction.
    
</dd>
</dl>

<dl>
<dd>

**baseAmount:** `int64` — Total amount of the transaction in the merchant’s currency. The value is in the currency’s lowest denomination, for example, cents.
    
</dd>
</dl>

<dl>
<dd>

**baseCurrency:** `*payroc.Currency` 
    
</dd>
</dl>

<dl>
<dd>

**paymentMethod:** `*paymentfeatures.FxRateInquiryPaymentMethod` — Object that contains information about the customer's payment details.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PaymentFeatures Bank
<details><summary><code>client.PaymentFeatures.Bank.Verify(request) -> *payroc.BankAccountVerificationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to verify a customer's bank account details.  

In the request, send the customer's bank account details. Our gateway can verify the following types of bank details:  
- Automated Clearing House (ACH) details  
- Pre-Authorized Debit (PAD) details  

In the response, our gateway indicates if the account details are valid and if you should use them in follow-on actions.  
  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &paymentfeatures.BankAccountVerificationRequest{
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        ProcessingTerminalId: "1234001",
        BankAccount: &paymentfeatures.BankAccountVerificationRequestBankAccount{
            Pad: &payroc.PadPayload{
                NameOnAccount: "Sarah Hazel Hopper",
                AccountNumber: "1234567890",
                TransitNumber: "76543",
                InstitutionNumber: "543",
            },
        },
    }
client.PaymentFeatures.Bank.Verify(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**bankAccount:** `*paymentfeatures.BankAccountVerificationRequestBankAccount` — Object that contains information about the bank account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PaymentLinks SharingEvents
<details><summary><code>client.PaymentLinks.SharingEvents.List(PaymentLinkId) -> *core.PayrocPager[*payroc.SharingEventPaginatedList, []*payroc.PaymentLinkEmailShareEvent, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of sharing events for a payment link. A sharing event occurs when a merchant shares a payment link with a customer.  

To list the sharing events for a payment link, you need its paymentLinkId. Our gateway returned the paymentLinkId in the response of the [Create Payment Link](https://docs.payroc.com/api/schema/payment-links/create) method.  

**Note:** If you don't have the paymentLinkId, use our [List Payment Links](https://docs.payroc.com/api/schema/payment-links/list) method to search for the payment link.  

Use query parameters to filter the list of results that we return, for example, to search for links sent to a specific customer.  

Our gateway returns the following information for each sharing event in the list:  
- Customer that the merchant sent the link to.  
- Date that the merchant sent the link.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &paymentlinks.ListSharingEventsRequest{
        PaymentLinkId: "JZURRJBUPS",
        RecipientName: payroc.String(
            "Sarah Hazel Hopper",
        ),
        RecipientEmail: payroc.String(
            "sarah.hopper@example.com",
        ),
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.PaymentLinks.SharingEvents.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentLinkId:** `string` — Unique identifier that we assigned to the payment link.
    
</dd>
</dl>

<dl>
<dd>

**recipientName:** `*string` — Filter results by the customer's name.
    
</dd>
</dl>

<dl>
<dd>

**recipientEmail:** `*string` — Filter results by the customer's email address.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PaymentLinks.SharingEvents.Share(PaymentLinkId, request) -> *payroc.PaymentLinkEmailShareEvent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to email a payment link to a customer.  

To email a payment link, you need its paymentLinkId. Our gateway returned the paymentLinkId in the response of the [Create Payment Link](https://docs.payroc.com/api/schema/payment-links/create) method.  

**Note:** If you don't have the paymentLinkId, use our [List Payment Links](https://docs.payroc.com/api/schema/payment-links/list) method to search for the payment link.  

In the request, you must provide the recipient's name and email address.  

In the response, our gateway returns a sharingEventId, which you can use to [List Payment Link Sharing Events](https://docs.payroc.com/api/schema/payment-links/sharing-events/list).  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &paymentlinks.ShareSharingEventsRequest{
        PaymentLinkId: "JZURRJBUPS",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.PaymentLinkEmailShareEvent{
            SharingMethod: payroc.PaymentLinkEmailShareEventSharingMethodEmail,
            MerchantCopy: payroc.Bool(
                true,
            ),
            Message: payroc.String(
                `Dear Sarah,
                
                Your insurance is expiring this month.
                Please, pay the renewal fee by the end of the month to renew it.
                `,
            ),
            Recipients: []*payroc.PaymentLinkEmailRecipient{
                &payroc.PaymentLinkEmailRecipient{
                    Name: "Sarah Hazel Hopper",
                    Email: "sarah.hopper@example.com",
                },
            },
        },
    }
client.PaymentLinks.SharingEvents.Share(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentLinkId:** `string` — Unique identifier that we assigned to the payment link.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.PaymentLinkEmailShareEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PayrocCloud PaymentInstructions
<details><summary><code>client.PayrocCloud.PaymentInstructions.Submit(SerialNumber, request) -> *payroc.PaymentInstruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to submit an instruction request to initiate a sale on a payment device.  

In the request, include the order amount and currency.  

When you send a successful request, our gateway returns information about the payment instruction and a paymentInstructionId, which you need for the following methods:
- [Retrieve payment instruction](https://docs.payroc.com/api/schema/payroc-cloud/payment-instructions/retrieve) - View the details of the payment instruction.
- [Cancel payment instruction](https://docs.payroc.com/api/schema/payroc-cloud/payment-instructions/delete) - Cancel the payment instruction.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.PaymentInstructionRequest{
        SerialNumber: "1850010868",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Operator: payroc.String(
            "Jane",
        ),
        ProcessingTerminalId: "1234001",
        Order: &payroc.PaymentInstructionOrder{
            OrderId: "OrderRef6543",
            Amount: 4999,
            Currency: payroc.CurrencyUsd,
        },
        CustomizationOptions: &payroc.CustomizationOptions{
            EntryMethod: payroc.CustomizationOptionsEntryMethodDeviceRead.Ptr(),
        },
        AutoCapture: payroc.Bool(
            true,
        ),
    }
client.PayrocCloud.PaymentInstructions.Submit(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serialNumber:** `string` — Serial number of the merchant’s payment device.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who initiated the request.
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*payroc.PaymentInstructionOrder` 
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.Customer` 
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `*payroc.IpAddress` 
    
</dd>
</dl>

<dl>
<dd>

**credentialOnFile:** `*payroc.SchemasCredentialOnFile` 
    
</dd>
</dl>

<dl>
<dd>

**customizationOptions:** `*payroc.CustomizationOptions` 
    
</dd>
</dl>

<dl>
<dd>

**autoCapture:** `*bool` 

Indicates if we should automatically capture the payment amount.  

- `true` - Run a sale and automatically capture the transaction.
- `false`- Run a pre-authorization and capture the transaction later.  

**Note:** If you send `false` and the terminal doesn't support pre-authorization, we set the transaction's status to pending. The merchant must capture the transaction to take payment from the customer.
    
</dd>
</dl>

<dl>
<dd>

**processAsSale:** `*bool` 

Indicates if we should immediately settle the sale transaction. The merchant cannot adjust the transaction if we immediately settle it.  
**Note:** If the value for **processAsSale** is `true`, the gateway ignores the value in **autoCapture**.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PayrocCloud.PaymentInstructions.Retrieve(PaymentInstructionId) -> *payroc.PaymentInstruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a payment instruction.  

To retrieve a payment instruction, you need its paymentInstructionId. Our gateway returned the paymentInstructionId in the response of the [Submit Payment Instruction](https://docs.payroc.com/api/schema/payroc-cloud/payment-instructions/submit) method.  

Our gateway returns the status of the payment instruction. If the payment device completed the payment instruction, the response also includes a link to the payment.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.RetrievePaymentInstructionsRequest{
        PaymentInstructionId: "e743a9165d134678a9100ebba3b29597",
    }
client.PayrocCloud.PaymentInstructions.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentInstructionId:** `string` — Unique identifier that we assigned to the payment instruction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PayrocCloud.PaymentInstructions.Delete(PaymentInstructionId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to cancel a payment instruction.  

You can cancel a payment instruction only if its status is `inProgress`. To retrieve the status of a payment instruction, use our [Retrieve Payment Instruction](https://docs.payroc.com/api/schema/payroc-cloud/payment-instructions/retrieve) method.  

To cancel a payment instruction, you need its paymentInstructionId. Our gateway returned the paymentInstructionId in the response of the [Submit Payment Instruction](https://docs.payroc.com/api/schema/payroc-cloud/payment-instructions/submit) method.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.DeletePaymentInstructionsRequest{
        PaymentInstructionId: "e743a9165d134678a9100ebba3b29597",
    }
client.PayrocCloud.PaymentInstructions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**paymentInstructionId:** `string` — Unique identifier that we assigned to the payment instruction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PayrocCloud RefundInstructions
<details><summary><code>client.PayrocCloud.RefundInstructions.Submit(SerialNumber, request) -> *payroc.RefundInstruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to submit an instruction request to initiate a refund on a payment device.  

In the request, include the refund amount and currency.  

If the request is successful, our gateway returns information about the refund instruction and a refundInstructionId, which you need for the following methods:
- [Retrieve refund instruction](https://docs.payroc.com/api/schema/payroc-cloud/refund-instructions/retrieve) - View the details of the refund instruction.
- [Cancel refund instruction](https://docs.payroc.com/api/schema/payroc-cloud/refund-instructions/delete) - Cancel the refund instruction.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.RefundInstructionRequest{
        SerialNumber: "1850010868",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Operator: payroc.String(
            "Jane",
        ),
        ProcessingTerminalId: "1234001",
        Order: &payroc.RefundInstructionOrder{
            OrderId: "OrderRef6543",
            Description: payroc.String(
                "Refund for order OrderRef6543",
            ),
            Amount: 4999,
            Currency: payroc.CurrencyUsd,
        },
        CustomizationOptions: &payroc.CustomizationOptions{
            EntryMethod: payroc.CustomizationOptionsEntryMethodManualEntry.Ptr(),
        },
    }
client.PayrocCloud.RefundInstructions.Submit(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serialNumber:** `string` — Serial number that identifies the merchant’s payment device.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who initiated the request.
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*payroc.RefundInstructionOrder` 
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.Customer` 
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `*payroc.IpAddress` 
    
</dd>
</dl>

<dl>
<dd>

**customizationOptions:** `*payroc.CustomizationOptions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PayrocCloud.RefundInstructions.Retrieve(RefundInstructionId) -> *payroc.RefundInstruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a refund instruction.  

To retrieve a refund instruction, you need its refundInstructionId. Our gateway returned the refundInstructionId in the response of the [Submit Refund Instruction](https://docs.payroc.com/api/schema/payroc-cloud/refund-instructions/submit) method.  

Our gateway returns the status of the refund instruction. If the payment device completed the refund instruction, the response also includes a link to the refund.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.RetrieveRefundInstructionsRequest{
        RefundInstructionId: "a37439165d134678a9100ebba3b29597",
    }
client.PayrocCloud.RefundInstructions.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refundInstructionId:** `string` — Unique identifier that we assigned to the refund instruction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PayrocCloud.RefundInstructions.Delete(RefundInstructionId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to cancel a refund instruction.  

You can cancel a refund instruction only if its status is `inProgress`. To retrieve the status of a refund instruction, use our [Retrieve Refund Instruction](https://docs.payroc.com/api/schema/payroc-cloud/refund-instructions/retrieve) method.  

To cancel a refund instruction, you need its refundInstructionId. Our gateway returned the refundInstructionId in the response of the [Submit Refund Instruction](https://docs.payroc.com/api/schema/payroc-cloud/refund-instructions/submit) method. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.DeleteRefundInstructionsRequest{
        RefundInstructionId: "a37439165d134678a9100ebba3b29597",
    }
client.PayrocCloud.RefundInstructions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**refundInstructionId:** `string` — Unique identifier that we assigned to the refund instruction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PayrocCloud SignatureInstructions
<details><summary><code>client.PayrocCloud.SignatureInstructions.Submit(SerialNumber, request) -> *payroc.SignatureInstruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to submit an instruction to capture a customer's signature on a payment device.  

Our gateway returns information about the signature instruction and a signatureInstructionId, which you need for the following methods:
- [Retrieve signature instruction](https://docs.payroc.com/api/schema/payroc-cloud/signature-instructions/retrieve) - View the details of the signature instruction.
- [Cancel signature instruction](https://docs.payroc.com/api/schema/payroc-cloud/signature-instructions/delete) - Cancel the signature instruction.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.SignatureInstructionRequest{
        SerialNumber: "1850010868",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        ProcessingTerminalId: "1234001",
    }
client.PayrocCloud.SignatureInstructions.Submit(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serialNumber:** `string` — Serial number that identifies the merchant’s payment device.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PayrocCloud.SignatureInstructions.Retrieve(SignatureInstructionId) -> *payroc.SignatureInstruction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a signature instruction.  

To retrieve a signature instruction, you need its signatureInstructionId. Our gateway returned the signatureInstructionId in the response of the [Submit Signature Instruction](https://docs.payroc.com/api/schema/payroc-cloud/signature-instructions/submit) method.  

Our gateway returns the status of the instruction. If the payment device completed the instruction, the response also includes a link to retrieve the signature.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.RetrieveSignatureInstructionsRequest{
        SignatureInstructionId: "a37439165d134678a9100ebba3b29597",
    }
client.PayrocCloud.SignatureInstructions.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**signatureInstructionId:** `string` — Unique identifier that our gateway assigned to the signature instruction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PayrocCloud.SignatureInstructions.Delete(SignatureInstructionId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to cancel a signature instruction.  

To cancel a signature instruction, you need its signatureInstructionId. Our gateway returned the signatureInstructionId in the response of the [Submit signature instruction](https://docs.payroc.com/api/schema/payroc-cloud/signature-instructions/submit) method.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.DeleteSignatureInstructionsRequest{
        SignatureInstructionId: "a37439165d134678a9100ebba3b29597",
    }
client.PayrocCloud.SignatureInstructions.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**signatureInstructionId:** `string` — Unique identifier that our gateway assigned to the signature instruction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PayrocCloud Signatures
<details><summary><code>client.PayrocCloud.Signatures.Retrieve(SignatureId) -> *payroccloud.RetrieveSignaturesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve a signature that a payment device captured using Payroc Cloud. 

Our gateway returns the following information about the signature:
- Image of the signature
- Format of the image
- Date that the device captured the image
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &payroccloud.RetrieveSignaturesRequest{
        SignatureId: "JDN4ILZB0T",
    }
client.PayrocCloud.Signatures.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**signatureId:** `string` — Unique identifier that we assigned to the signature.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## RepeatPayments PaymentPlans
<details><summary><code>client.RepeatPayments.PaymentPlans.List(ProcessingTerminalId) -> *core.PayrocPager[*payroc.PaymentPlanPaginatedList, []*payroc.PaymentPlan, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of payment plans for a processing terminal.  

**Note:** If you want to view the details of a specific payment plan and you have its paymentPlanId, use our [Retrieve Payment Plan](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/retrieve) method.  

Our gateway returns the following information about each payment plan in the list:  

  -	Name, length, and currency of the plan  
  -	How often our gateway collects each payment  
  -	How much our gateway collects for each payment  
  -	What happens if the merchant updates or deletes the plan  

For each payment plan, we return the paymentPlanId, which you can use to perform follow-on actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.ListPaymentPlansRequest{
        ProcessingTerminalId: "1234001",
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.RepeatPayments.PaymentPlans.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.PaymentPlans.Create(ProcessingTerminalId, request) -> *payroc.PaymentPlan</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a payment schedule that you can assign customers to.  

**Note:** This method is part of our Repeat Payments feature. To help you understand how this method works with our Subscriptions endpoints, go to [Repeat Payments](https://docs.payroc.com/guides/integrate/repeat-payments).  

When you create a payment plan you need to provide a unique paymentPlanId that you use to run follow-on actions:  

-	[Retrieve Payment Plan](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/retrieve)  - View the details of the payment plan.  
-	[Update Payment Plan](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/partially-update)  - Update the details of the payment plan.  
-	[Delete Payment Plan](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/delete)  - Delete the payment plan.  
-	[Create Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/create)  - Subscribe a customer to the payment plan.  

The request includes the following settings:  

-	**type** - Indicates if our gateway or the merchant collects payments. If the merchant manually collects payments, integrate with the [Pay Manual Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/pay) method.  
-	**recurringOrder** - Amount of each payment if the gateway automatically collect payments.  
-	**setupOrder** - Setup fee that our gateway immediately collects from the customer's payment method.  
-	**onUpdate and onDelete** - Indicates what happens to associated subscriptions if the merchant updates or deletes the payment plan.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.CreatePaymentPlansRequest{
        ProcessingTerminalId: "1234001",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.PaymentPlan{
            PaymentPlanId: "PlanRef8765",
            Name: "Premium Club",
            Description: payroc.String(
                "Monthly Premium Club subscription",
            ),
            Currency: payroc.CurrencyUsd,
            Length: payroc.Int(
                12,
            ),
            Type: payroc.PaymentPlanBaseTypeAutomatic,
            Frequency: payroc.PaymentPlanBaseFrequencyMonthly,
            OnUpdate: payroc.PaymentPlanBaseOnUpdateContinue,
            OnDelete: payroc.PaymentPlanBaseOnDeleteComplete,
            CustomFieldNames: []string{
                "yourCustomField",
            },
            SetupOrder: &payroc.PaymentPlanSetupOrder{
                Amount: payroc.Int64(
                    4999,
                ),
                Description: payroc.String(
                    "Initial setup fee for Premium Club subscription",
                ),
                Breakdown: &payroc.PaymentPlanOrderBreakdown{
                    Subtotal: 4347,
                    Taxes: []*payroc.RetrievedTax{
                        &payroc.RetrievedTax{
                            Name: "Sales Tax",
                            Rate: 5,
                        },
                    },
                },
            },
            RecurringOrder: &payroc.PaymentPlanRecurringOrder{
                Amount: payroc.Int64(
                    4999,
                ),
                Description: payroc.String(
                    "Monthly Premium Club subscription",
                ),
                Breakdown: &payroc.PaymentPlanOrderBreakdown{
                    Subtotal: 4347,
                    Taxes: []*payroc.RetrievedTax{
                        &payroc.RetrievedTax{
                            Name: "Sales Tax",
                            Rate: 5,
                        },
                    },
                },
            },
        },
    }
client.RepeatPayments.PaymentPlans.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.PaymentPlan` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.PaymentPlans.Retrieve(ProcessingTerminalId, PaymentPlanId) -> *payroc.PaymentPlan</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a payment plan.  

To retrieve a payment plan, you need its paymentPlanId. Our gateway returned the paymentPlanId in the response of the [Create Payment Plan](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/create) method.  

**Note:** If you don't have the paymentPlanId, use our [List Payment Plans](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/list) method to search for the payment plan.  

Our gateway returns the following information about the payment plan:  

  -	Name, length, and currency of the plan  
  -	How often our gateway collects each payment  
  -	How much our gateway collects for each payment  
  -	What happens if the merchant updates or deletes the plan  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.RetrievePaymentPlansRequest{
        ProcessingTerminalId: "1234001",
        PaymentPlanId: "PlanRef8765",
    }
client.RepeatPayments.PaymentPlans.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**paymentPlanId:** `string` — Unique identifier that the merchant assigned to the payment plan.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.PaymentPlans.Delete(ProcessingTerminalId, PaymentPlanId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to delete a payment plan.  

> **Important:** When you delete a payment plan, you can’t recover it. You also won’t be able to add subscriptions to the payment plan.  

To delete a payment plan, you need its paymentPlanId, which you sent in the request of the [Create Payment Plan](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/create) method.  

**Note:** If you don't have the paymentPlanId, use our [List Payment Plans](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/list) method to search for the payment plan.  

The value you sent for the onDelete parameter when you created the payment plan indicates what happens to associated subscriptions when you delete the plan:  

  -	`complete` - Our gateway stops taking payments for the subscriptions associated with the payment plan.  
  -	`continue` - Our gateway continues to take payments for the subscriptions associated with the payment plan. To stop a subscription for a cancelled payment plan, go to the [Deactivate Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/deactivate) method.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.DeletePaymentPlansRequest{
        ProcessingTerminalId: "1234001",
        PaymentPlanId: "PlanRef8765",
    }
client.RepeatPayments.PaymentPlans.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**paymentPlanId:** `string` — Unique identifier that the merchant assigned to the payment plan.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.PaymentPlans.PartiallyUpdate(ProcessingTerminalId, PaymentPlanId, request) -> *payroc.PaymentPlan</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to partially update a payment plan. Structure your request to follow the [RFC 6902](https://datatracker.ietf.org/doc/html/rfc6902) standard.  

To update a payment plan, you need its paymentPlanId, which you sent in the request of the [Create Payment Plan](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/create) method.  

**Note:** If you don't have the paymentPlanId, use our [List Payment Plans](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/list) method to search for the payment plan.  

You can update all of the properties of the payment plan except for the paymentPlanId.  

The value you sent for the onUpdate parameter when you created the payment plan indicates what happens to the associated subscriptions when you update the plan:  
- `update` - Our gateway updates the subscriptions associated with the payment plan.
- `continue` - Our  gateway doesn't update the subscriptions associated with the payment plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.PartiallyUpdatePaymentPlansRequest{
        ProcessingTerminalId: "1234001",
        PaymentPlanId: "PlanRef8765",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: []*payroc.PatchDocument{
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
        },
    }
client.RepeatPayments.PaymentPlans.PartiallyUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**paymentPlanId:** `string` — Unique identifier that the merchant assigned to the payment plan.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `payroc.PatchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## RepeatPayments Subscriptions
<details><summary><code>client.RepeatPayments.Subscriptions.List(ProcessingTerminalId) -> *core.PayrocPager[*payroc.SubscriptionPaginatedList, []*payroc.Subscription, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of subscriptions.  

Note: If you want to view the details of a specific subscription and you have its subscriptionId, use our [Retrieve subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/retrieve) method.  

Use query parameters to filter the list of results that we return, for example, to search for subscriptions for a customer, a payment plan, or frequency.  

Our gateway returns information about the following for each subscription in the list:  

-	Payment plan the subscription is linked to.  
-	Secure token that represents cardholder’s payment details.  
-	Current state of the subscription, including its status, next due date, and invoices.  
-	Fees for setup and the cost of the recurring order.  
-	Subscription length, end date, and frequency.  

For each subscription, we also return the subscriptionId, the paymentPlanId, and the secureTokenId, which you can use to perform follow-actions. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.ListSubscriptionsRequest{
        ProcessingTerminalId: "1234001",
        CustomerName: payroc.String(
            "Sarah%20Hazel%20Hopper",
        ),
        Last4: payroc.String(
            "7062",
        ),
        PaymentPlan: payroc.String(
            "Premium%20Club",
        ),
        Frequency: repeatpayments.ListSubscriptionsRequestFrequencyWeekly.Ptr(),
        Status: repeatpayments.ListSubscriptionsRequestStatusActive.Ptr(),
        EndDate: payroc.Time(
            payroc.MustParseDate(
                "2025-07-01",
            ),
        ),
        NextDueDate: payroc.Time(
            payroc.MustParseDate(
                "2024-08-01",
            ),
        ),
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.RepeatPayments.Subscriptions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**customerName:** `*string` — Filter by the customer's name.
    
</dd>
</dl>

<dl>
<dd>

**last4:** `*string` — Filter by the last four digits of the card or account number.
    
</dd>
</dl>

<dl>
<dd>

**paymentPlan:** `*string` — Filter by the name of the payment plan.
    
</dd>
</dl>

<dl>
<dd>

**frequency:** `*repeatpayments.ListSubscriptionsRequestFrequency` — Filter by the frequency of subscription payments.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*repeatpayments.ListSubscriptionsRequestStatus` — Filter by the current status of the subscription.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` 

Format: `YYYY-MM-DD`  
Filter subscriptions that end on a specific date.
    
</dd>
</dl>

<dl>
<dd>

**nextDueDate:** `*time.Time` 

Format: `YYYY-MM-DD`  
Filter subscriptions by the date that the next payment is collected.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.Subscriptions.Create(ProcessingTerminalId, request) -> *payroc.Subscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to assign a customer to a payment plan.  

**Note:** This method is part of our Repeat Payments feature. To help you understand how this method works with our Payment plans endpoints, go to [Repeat Payments](https://docs.payroc.com/guides/integrate/repeat-payments).  

When you create a subscription you need to provide a unique subscriptionId that you use to run follow-on actions:  

- [Retrieve Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/retrieve) - View the details of the subscription.
- [Update Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/partially-update) - Update the details of the subscription.
- [Deactivate Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/deactivate) - Stop taking payments for the subscription.
- [Re-activate Subscription](https://docs.payroc.com/api/schema/payments/subscriptions/reactivate) - Start taking payments again for the subscription.
- [Pay Manual Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/pay) - Manually collect a payment for the subscription.

The request includes the following settings:
- **paymentPlanId** - Unique identifier of the payment plan that the merchant wants to use. If you don't have the paymentPlanId, use our [List Payment Plans](https://docs.payroc.com/api/schema/repeat-payments/payment-plans/list) method to search for the payment plan.
- **paymentMethod** - Object that contains information about the secure token, which represents the customer's card details or bank account details.
- **startDate** - Date that you want to start to take payments.

You can also update the settings that the subscription inherited from the payment plan, for example, you can change the amount for each payment. If you change the settings for the subscription, it doesn't change the settings in the payment plan that it's linked to. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.SubscriptionRequest{
        ProcessingTerminalId: "1234001",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        SubscriptionId: "SubRef7654",
        PaymentPlanId: "PlanRef8765",
        PaymentMethod: &repeatpayments.SubscriptionRequestPaymentMethod{
            SecureToken: &payroc.SecureTokenPayload{
                Token: "1234567890123456789",
            },
        },
        Name: payroc.String(
            "Premium Club",
        ),
        Description: payroc.String(
            "Premium Club subscription",
        ),
        SetupOrder: &payroc.SubscriptionPaymentOrderRequest{
            OrderId: payroc.String(
                "OrderRef6543",
            ),
            Amount: payroc.Int64(
                4999,
            ),
            Description: payroc.String(
                "Initial setup fee for Premium Club subscription",
            ),
        },
        RecurringOrder: &payroc.SubscriptionRecurringOrderRequest{
            Amount: payroc.Int64(
                4999,
            ),
            Description: payroc.String(
                "Monthly Premium Club subscription",
            ),
            Breakdown: &payroc.SubscriptionOrderBreakdownRequest{
                Subtotal: 4347,
                Taxes: []*payroc.TaxRate{
                    &payroc.TaxRate{
                        Type: payroc.TaxRateTypeRate,
                        Rate: 5,
                        Name: "Sales Tax",
                    },
                },
            },
        },
        StartDate: payroc.MustParseDate(
            "2024-07-02",
        ),
        EndDate: payroc.Time(
            payroc.MustParseDate(
                "2025-07-01",
            ),
        ),
        Length: payroc.Int(
            12,
        ),
        PauseCollectionFor: payroc.Int(
            0,
        ),
    }
client.RepeatPayments.Subscriptions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**subscriptionId:** `string` — Unique identifier that the merchant assigned to the subscription.
    
</dd>
</dl>

<dl>
<dd>

**paymentPlanId:** `string` — Unique identifier that the merchant assigned to the payment plan.
    
</dd>
</dl>

<dl>
<dd>

**paymentMethod:** `*repeatpayments.SubscriptionRequestPaymentMethod` — Object that contains information about the customer's payment details.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 

Name of the subscription. 
This value replaces the name inherited from the payment plan.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 

Description of the subscription. 
This value replaces the description inherited from the payment plan.
    
</dd>
</dl>

<dl>
<dd>

**setupOrder:** `*payroc.SubscriptionPaymentOrderRequest` 
    
</dd>
</dl>

<dl>
<dd>

**recurringOrder:** `*payroc.SubscriptionRecurringOrderRequest` 
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` 

Format: **YYYY-MM-DD**  
Subscription's start date.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*time.Time` 

Format: **YYYY-MM-DD**  
Subscription's end date.  
**Note:** If you provide values for both **length** and **endDate**, 
our gateway uses the value for **endDate** to determine when the subscription should end.
    
</dd>
</dl>

<dl>
<dd>

**length:** `*int` 

Total number of billing cycles. To indicate that the subscription should run indefinitely, send a value of `0`.
This value replaces the **length** inherited from the payment plan.  
**Note:** If you provide values for both **length** and **endDate**, 
our gateway uses the value for **endDate** to determine when the subscription should end.
    
</dd>
</dl>

<dl>
<dd>

**pauseCollectionFor:** `*int` 

Number of billing cycles that the merchant wants to pause payments for. 
For example, if the merchant wants to offer a free trial period.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*payroc.CustomField` — Array of customField objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.Subscriptions.Retrieve(ProcessingTerminalId, SubscriptionId) -> *payroc.Subscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a subscription.  

To retrieve a subscription, you need its subscriptionId. You sent the subscriptionId in the request of the [Create subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/create) method.  

**Note:** If you don't have the subscriptionId, use our [List subscriptions](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/list) method to search for the subscription.  

Our gateway returns information about the following for the subscription:  

-	Payment plan the subscription is linked to.  
-	Secure token that represents cardholder’s payment details.  
-	Current state of the subscription, including its status, next due date, and invoices.  
-	Fees for setup and the cost of the recurring order.  
-	Subscription length, end date, and frequency.  

We also return the paymentPlanId and the secureTokenId, which you can use to perform follow-on actions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.RetrieveSubscriptionsRequest{
        ProcessingTerminalId: "1234001",
        SubscriptionId: "SubRef7654",
    }
client.RepeatPayments.Subscriptions.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**subscriptionId:** `string` — Unique identifier that the merchant assigned to the subscription.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.Subscriptions.PartiallyUpdate(ProcessingTerminalId, SubscriptionId, request) -> *payroc.Subscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to partially update a subscription. Structure your request to follow the [RFC 6902](https://datatracker.ietf.org/doc/html/rfc6902) standard.  

To update a subscription, you need its subscriptionId, which you sent in the request of the [Create subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/create) method.  

**Note:** If you don't have the subscriptionId, use our [List subscriptions](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/list) method to search for the payment.  

You can update all of the properties of the subscription except for the following:  

**Can't delete**  
- recurringOrder
- description
- name

**Can't perform any PATCH operation**  
- currentState
- type
- frequency
- paymentPlan
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.PartiallyUpdateSubscriptionsRequest{
        ProcessingTerminalId: "1234001",
        SubscriptionId: "SubRef7654",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: []*payroc.PatchDocument{
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
        },
    }
client.RepeatPayments.Subscriptions.PartiallyUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**subscriptionId:** `string` — Unique identifier that the merchant assigned to the subscription.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `payroc.PatchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.Subscriptions.Deactivate(ProcessingTerminalId, SubscriptionId) -> *payroc.Subscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to deactivate a subscription.  

To deactivate a subscription, you need its subscriptionId, which you sent in the request of the [Create Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/create) method.  

**Note:** If you don't have the subscriptionId, use our [List Subscriptions](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/list) method to search for the subscription.  

If your request is successful, our gateway stops taking payments from the customer.  

To reactivate the subscription, use our [Reactivate Subscription](https://docs.payroc.com/api/schema/payments/subscriptions/reactivate) method.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.DeactivateSubscriptionsRequest{
        ProcessingTerminalId: "1234001",
        SubscriptionId: "SubRef7654",
    }
client.RepeatPayments.Subscriptions.Deactivate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**subscriptionId:** `string` — Unique identifier that the merchant assigned to the subscription.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.Subscriptions.Reactivate(ProcessingTerminalId, SubscriptionId) -> *payroc.Subscription</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to reactivate a subscription.  

To reactivate a subscription, you need its subscriptionId, which you sent in the request of the [Create Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/create) method.  

**Note:** If you don't have the subscriptionId, use our [List Subscriptions](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/list) method to search for the subscription.  

If your request is successful, our gateway restarts taking payments from the customer.  

To deactivate the subscription, use our [Deactivate Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/deactivate) method.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.ReactivateSubscriptionsRequest{
        ProcessingTerminalId: "1234001",
        SubscriptionId: "SubRef7654",
    }
client.RepeatPayments.Subscriptions.Reactivate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**subscriptionId:** `string` — Unique identifier that the merchant assigned to the subscription.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RepeatPayments.Subscriptions.Pay(ProcessingTerminalId, SubscriptionId, request) -> *payroc.SubscriptionPayment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to manually collect a payment linked to a subscription. You can manually collect a payment only if the merchant chose not to let our gateway automatically collect each payment.  

To manually collect a payment, you need the subscriptionId of the subscription that's linked to the payment. You sent the subscriptionId in the request of the [Create Subscription](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/create) method.  

**Note:** If you don't have the subscriptionId, use our [List Subscriptions](https://docs.payroc.com/api/schema/repeat-payments/subscriptions/list) method to search for the subscription.  

The request includes an order object that contains information about the amount that you want to collect.  

In the response, our gateway returns information about the payment and a paymentId. You can use the paymentId in follow-on actions with the [Payments](https://docs.payroc.com/api/schema/card-payments/payments) endpoints or [Bank Transfer Payments](https://docs.payroc.com/api/schema/bank-transfer-payments/payments) endpoints.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &repeatpayments.SubscriptionPaymentRequest{
        ProcessingTerminalId: "1234001",
        SubscriptionId: "SubRef7654",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Operator: payroc.String(
            "Jane",
        ),
        Order: &payroc.SubscriptionPaymentOrder{
            OrderId: payroc.String(
                "OrderRef6543",
            ),
            Amount: payroc.Int64(
                4999,
            ),
            Description: payroc.String(
                "Monthly Premium Club subscription",
            ),
        },
    }
client.RepeatPayments.Subscriptions.Pay(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**subscriptionId:** `string` — Unique identifier that the merchant assigned to the subscription.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who initiated the request.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*payroc.SubscriptionPaymentOrder` — Object that contains information about the payment.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*payroc.CustomField` — Array of customField objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Reporting Settlement
<details><summary><code>client.Reporting.Settlement.ListBatches() -> *core.PayrocPager[*reporting.ListBatchesSettlementResponse, []*payroc.Batch, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of batches that your merchants submitted to the processor on a specific date.  

**Note:** If you want to view the details of a specific batch and you have its batchId, use our [Retrieve Batch](https://docs.payroc.com/api/schema/reporting/settlement/retrieve-batch) method.  

Use query parameters to filter the list of results that we return, for example, to search for batches that were submitted by a specific merchant.  

> **Important:** You must provide a value for the date query parameter.  

Our gateway returns the following information about each batch in the list:  
-	Transaction information, including the number of transactions and total value of sales.  
-	Merchant information, including the merchant ID (MID) and the processing account that the batch is associated with.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.ListReportingSettlementBatchesRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        Date: payroc.MustParseDate(
            "2027-07-02",
        ),
        MerchantId: payroc.String(
            "4525644354",
        ),
    }
client.Reporting.Settlement.ListBatches(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**date:** `time.Time` — Filter batches by the date that they were submitted. The format of this value is **YYYY-MM-DD**.
    
</dd>
</dl>

<dl>
<dd>

**merchantId:** `*string` — Filter results by the unique identifier that the processor assigned to the merchant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.RetrieveBatch(BatchId) -> *payroc.Batch</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a batch.  

**Note:** To retrieve a batch, you need its batchId. If you don't have the batchId, use our [List Batches](https://docs.payroc.com/api/schema/reporting/settlement/list-batches) method to search for the batch.  

Our gateway returns the following information about the batch:  

-	Transaction information, including the number of transactions and total value of sales.  
-	Merchant information, including the merchant ID (MID) and the processing account that the batch is associated with.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.RetrieveBatchSettlementRequest{
        BatchId: 1,
    }
client.Reporting.Settlement.RetrieveBatch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchId:** `int` — Unique identifier that we assigned to the batch.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.ListTransactions() -> *core.PayrocPager[*reporting.ListTransactionsSettlementResponse, []*payroc.Transaction, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a paginated list of your merchants’ transactions.  

**Note:** If you want to view the details of a specific transaction and you have its transactionId, use our [Retrieve Transaction](https://docs.payroc.com/api/schema/reporting/settlement/retrieve-transaction) method.  

Use query parameters to filter the list of results that we return, for example, to search for transactions for a specific merchant.  

> **Important:** You must provide a value for either the date query parameter or the batchId query parameter.  

Our gateway returns the following information about each transaction in the list:  

-	Merchant and processing account that ran the transaction.  
-	Transaction type, date, amount, and the payment method that the customer used.  
-	Batch that contains the transaction, and authorization details for the transaction.  
-	Processor that settled the transaction and the ACH deposit containing the transaction.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.ListReportingSettlementTransactionsRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        Date: payroc.Time(
            payroc.MustParseDate(
                "2024-07-02",
            ),
        ),
        BatchId: payroc.Int(
            1,
        ),
        MerchantId: payroc.String(
            "4525644354",
        ),
        TransactionType: reporting.ListTransactionsSettlementRequestTransactionTypeCapture.Ptr(),
    }
client.Reporting.Settlement.ListTransactions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**date:** `*time.Time` 

Filter transactions by the date that the merchant submitted the batch that contains the transaction. The format of this value is **YYYY-MM-DD**.  

You must provide either the batchId or the date. 
    
</dd>
</dl>

<dl>
<dd>

**batchId:** `*int` 

Filter transactions by the unique identifier of the batch that contains the transaction.  

You must provide either the batchId or the date. 
    
</dd>
</dl>

<dl>
<dd>

**merchantId:** `*string` — Filter results by the unique identifier that the processor assigned to the merchant.
    
</dd>
</dl>

<dl>
<dd>

**transactionType:** `*reporting.ListTransactionsSettlementRequestTransactionType` — Filter transactions by transaction type.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.RetrieveTransaction(TransactionId) -> *payroc.Transaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a transaction.  

**Note:** To retrieve a transaction, you need its transactionId. If you don't have the transactionId, use our [List Transactions](https://docs.payroc.com/api/schema/reporting/settlement/list-transactions) method to search for the transaction.  

Our gateway returns the following information about the transaction:  

-	Merchant and processing account that ran the transaction.  
-	Transaction type, date, amount, and the payment method that the customer used.  
-	Batch that contains the transaction, and authorization details for the transaction.   
-	Processor that settled the transaction and the ACH deposit containing the transaction.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.RetrieveTransactionSettlementRequest{
        TransactionId: 1,
    }
client.Reporting.Settlement.RetrieveTransaction(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transactionId:** `int` — Unique identifier of the transaction.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.ListAuthorizations() -> *core.PayrocPager[*reporting.ListAuthorizationsSettlementResponse, []*payroc.Authorization, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve a [paginated](https://docs.payroc.com/api/pagination) list of authorizations.  

Use query parameters to filter the list of results that we return, for example, to search for authorizations linked to a specific merchant.  

> **Important:** You must provide a value for either the date query parameter or the batchId query parameter.  

Our gateway returns the following information about each authorization in the list:
- Authorization response from the issuing bank.
- Amount that the issuing bank authorized.
- Merchant that ran the authorization.
- Details about the customer's card, the transaction, and the batch.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.ListReportingSettlementAuthorizationsRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        Date: payroc.Time(
            payroc.MustParseDate(
                "2024-07-02",
            ),
        ),
        BatchId: payroc.Int(
            1,
        ),
        MerchantId: payroc.String(
            "4525644354",
        ),
    }
client.Reporting.Settlement.ListAuthorizations(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**date:** `*time.Time` 

Filter transactions by the date that the merchant submitted the batch that contains the transaction. The format of this value is **YYYY-MM-DD**.  

You must provide either the batchId or the date. 
    
</dd>
</dl>

<dl>
<dd>

**batchId:** `*int` 

Filter transactions by the unique identifier of the batch that contains the transaction.  

You must provide either the batchId or the date. 
    
</dd>
</dl>

<dl>
<dd>

**merchantId:** `*string` — Filter results by the unique identifier that the processor assigned to the merchant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.RetrieveAuthorization(AuthorizationId) -> *payroc.Authorization</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about an authorization.  

**Note:** To retrieve an authorization, you need its authorizationId. If you don't have the authorizationId, use our [List Authorizations](https://docs.payroc.com/api/schema/reporting/settlement/list-authorizations) method to search for the authorization.  

Our gateway returns the following information about the authorization:
- Authorization response from the issuing bank.
- Amount that the issuing bank authorized.
- Merchant that ran the authorization.
- Details about the customer's card, the transaction, and the batch.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.RetrieveAuthorizationSettlementRequest{
        AuthorizationId: 1,
    }
client.Reporting.Settlement.RetrieveAuthorization(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authorizationId:** `int` — Unique identifier of the authorization.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.ListDisputes() -> *core.PayrocPager[*reporting.ListDisputesSettlementResponse, []*payroc.Dispute, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of disputes.  

Use query parameters to filter the list of results that we return, for example, to search for disputes linked to a specific merchant.  

> **Important:** You must provide a value for the date query parameter.  

Our gateway returns the following information about each dispute in the list:  
- Its status, type, and description.  
- Transaction that the dispute is linked to, including the transaction date, merchant who ran the transaction, and the payment method that the cardholder used.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.ListReportingSettlementDisputesRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        Date: payroc.MustParseDate(
            "2024-07-02",
        ),
        MerchantId: payroc.String(
            "4525644354",
        ),
    }
client.Reporting.Settlement.ListDisputes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**date:** `time.Time` — Filter results by the date that the dispute was submitted.
    
</dd>
</dl>

<dl>
<dd>

**merchantId:** `*string` — Filter results by the unique identifier that the processor assigned to the merchant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.ListDisputesStatuses(DisputeId) -> []*payroc.DisputeStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return the status history of a dispute.  

To view the status history of a dispute, you need its disputeId. If you don't have the disputeId, use our [List Disputes](https://docs.payroc.com/api/schema/reporting/settlement/list-disputes) method to search for the dispute. 

Our gateway returns a list that contains each status change, the date it was changed, and its updated status.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.ListDisputesStatusesSettlementRequest{
        DisputeId: 1,
    }
client.Reporting.Settlement.ListDisputesStatuses(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disputeId:** `int` — Unique identifier of the dispute.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.ListAchDeposits() -> *core.PayrocPager[*reporting.ListAchDepositsSettlementResponse, []*payroc.AchDeposit, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of ACH deposits that we paid to your merchants.  

**Note:** If you want to view the details of a specific ACH deposit and you have its achDepositId, use our [Retrieve ACH Deposit](https://docs.payroc.com/api/schema/reporting/settlement/retrieve-ach-deposit) method.  

Use query parameters to filter the list of results that we return, for example, to search for ACH deposits that we paid to a specific merchant.  

> **Important:** You must provide a value for the date query parameter.  

Our gateway returns the following information about each ACH deposit in the list: 
- Merchant that we sent the ACH deposit to.
- Total amount that we paid the merchant.
- Breakdown of sales, returns, and fees.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.ListReportingSettlementAchDepositsRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        Date: payroc.MustParseDate(
            "2024-07-02",
        ),
        MerchantId: payroc.String(
            "4525644354",
        ),
    }
client.Reporting.Settlement.ListAchDeposits(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**date:** `time.Time` — Filter results by the date that the merchant received the ACH deposit.
    
</dd>
</dl>

<dl>
<dd>

**merchantId:** `*string` — Filter results by the unique identifier that the processor assigned to the merchant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.RetrieveAchDeposit(AchDepositId) -> *payroc.AchDeposit</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about an ACH deposit that we paid to a merchant.  

**Note:** To retrieve an ACH deposit, you need its achDepositId. If you don't have the achDepositId, use our [List ACH Deposits](https://docs.payroc.com/api/schema/reporting/settlement/list-ach-deposits) method to search for the ACH deposit.  

Our gateway returns the following information about the ACH deposit:  

- Merchant that we sent the ACH deposit to.  
- Total amount that we paid the merchant.  
- Breakdown of sales, returns, and fees.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.RetrieveAchDepositSettlementRequest{
        AchDepositId: 1,
    }
client.Reporting.Settlement.RetrieveAchDeposit(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**achDepositId:** `int` — Unique identifier of the ACH deposit.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Reporting.Settlement.ListAchDepositFees() -> *core.PayrocPager[*reporting.ListAchDepositFeesSettlementResponse, []*payroc.AchDepositFee, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of ACH deposit fees.

> **Important:** You must provide a value for either the 'date' query parameter or the 'achDepositId' query parameter.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &reporting.ListReportingSettlementAchDepositFeesRequest{
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
        Date: payroc.Time(
            payroc.MustParseDate(
                "2024-07-02",
            ),
        ),
        AchDepositId: payroc.Int(
            1,
        ),
        MerchantId: payroc.String(
            "4525644354",
        ),
    }
client.Reporting.Settlement.ListAchDepositFees(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>

<dl>
<dd>

**date:** `*time.Time` — Date to retrieve results from. You must provide either the 'achDepositId' or the 'date'.
    
</dd>
</dl>

<dl>
<dd>

**achDepositId:** `*int` — Unique identifier of the ACH deposit. You must provide either the 'achDepositId' or the 'date'.
    
</dd>
</dl>

<dl>
<dd>

**merchantId:** `*string` — Filter results by the unique identifier that the processor assigned to the merchant.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tokenization SecureTokens
<details><summary><code>client.Tokenization.SecureTokens.List(ProcessingTerminalId) -> *core.PayrocPager[*payroc.SecureTokenPaginatedListWithAccountType, []*payroc.SecureTokenWithAccountType, *payroc.Link]</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to return a [paginated](https://docs.payroc.com/api/pagination) list of secure tokens.  

**Note:** If you want to view the details of a specific secure token and you have its secureTokenId, use our [Retrieve Secure Token](https://docs.payroc.com/api/schema/tokenization/secure-tokens/retrieve) method.  

Use query parameters to filter the list of results that we return, for example, to search for secure tokens by customer or by the first four digits of a card number.  

Our gateway returns information about the following for each secure token in the list:  

  -	Payment details that the secure token represents.  
  -	Customer details, including shipping and billing addresses.  
  -	Secure token that you can use to carry out transactions.  

  For each secure token, we also return the secureTokenId, which you can use to perform follow-on actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tokenization.ListSecureTokensRequest{
        ProcessingTerminalId: "1234001",
        SecureTokenId: payroc.String(
            "MREF_abc1de23-f4a5-6789-bcd0-12e345678901fa",
        ),
        CustomerName: payroc.String(
            "Sarah%20Hazel%20Hopper",
        ),
        Phone: payroc.String(
            "2025550165",
        ),
        Email: payroc.String(
            "sarah.hopper@example.com",
        ),
        Token: payroc.String(
            "296753123456",
        ),
        First6: payroc.String(
            "453985",
        ),
        Last4: payroc.String(
            "7062",
        ),
        Before: payroc.String(
            "2571",
        ),
        After: payroc.String(
            "8516",
        ),
        Limit: payroc.Int(
            1,
        ),
    }
client.Tokenization.SecureTokens.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**secureTokenId:** `*string` — Unique identifier that the merchant assigned to the secure token.
    
</dd>
</dl>

<dl>
<dd>

**customerName:** `*string` — Filter by the customer's name.
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` — Filter by the customer's phone number.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Filter by the customer's email address.
    
</dd>
</dl>

<dl>
<dd>

**token:** `*string` — Filter by the token that the merchant used in a transaction to represent the customer's payment details.
    
</dd>
</dl>

<dl>
<dd>

**first6:** `*string` — Filter by the first six digits of the card number.
    
</dd>
</dl>

<dl>
<dd>

**last4:** `*string` — Filter by the last four digits of the card or account number.
    
</dd>
</dl>

<dl>
<dd>

**before:** `*string` 

Return the previous page of results before the value that you specify.  

You can’t send the before parameter in the same request as the after parameter. 
    
</dd>
</dl>

<dl>
<dd>

**after:** `*string` 

Return the next page of results after the value that you specify.  

You can’t send the after parameter in the same request as the before parameter. 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the maximum number of results that we return for each page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokenization.SecureTokens.Create(ProcessingTerminalId, request) -> *payroc.SecureToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a secure token that represents a customer's payment details.  

When you create a secure token, you need to generate and provide a secureTokenId that you use to run follow-on actions:  
- [Retrieve Secure Token](https://docs.payroc.com/api/schema/tokenization/secure-tokens/retrieve) – View the details of the secure token.  
- [Delete Secure Token](https://docs.payroc.com/api/schema/tokenization/secure-tokens/delete) – Delete the secure token.  
- [Update Secure Token](https://docs.payroc.com/api/schema/tokenization/secure-tokens/partially-update) – Update the details of the secure token.  
- [Update Account Details](https://docs.payroc.com/api/schema/tokenization/secure-tokens/update-account) – Update the secure token with the details from a single-use token.  

**Note:** If you don't generate a secureTokenId to identify the token, our gateway generates a unique identifier and returns it in the response.  

If the request is successful, our gateway returns a token that the merchant can use in transactions instead of the customer's sensitive payment details, for example, when they [run a sale](https://docs.payroc.com/api/schema/card-payments/payments/create).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tokenization.TokenizationRequest{
        ProcessingTerminalId: "1234001",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Operator: payroc.String(
            "Jane",
        ),
        MitAgreement: tokenization.TokenizationRequestMitAgreementUnscheduled.Ptr(),
        Customer: &payroc.Customer{
            FirstName: payroc.String(
                "Sarah",
            ),
            LastName: payroc.String(
                "Hopper",
            ),
            DateOfBirth: payroc.Time(
                payroc.MustParseDate(
                    "1990-07-15",
                ),
            ),
            ReferenceNumber: payroc.String(
                "Customer-12",
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
            ContactMethods: []*payroc.ContactMethod{
                &payroc.ContactMethod{
                    Email: &payroc.ContactMethodEmail{
                        Value: "jane.doe@example.com",
                    },
                },
            },
            NotificationLanguage: payroc.CustomerNotificationLanguageEn.Ptr(),
        },
        IpAddress: &payroc.IpAddress{
            Type: payroc.IpAddressTypeIpv4,
            Value: "104.18.24.203",
        },
        Source: &tokenization.TokenizationRequestSource{
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
client.Tokenization.SecureTokens.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**secureTokenId:** `*string` 

Unique identifier that the merchant created for the secure token that represents the customer's payment details. 
If the merchant doesn't create a secureTokenId, the gateway generates one and returns it in the response.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who saved the customer's payment details.
    
</dd>
</dl>

<dl>
<dd>

**mitAgreement:** `*tokenization.TokenizationRequestMitAgreement` 

Indicates how the merchant can use the customer's card details, as agreed by the customer:  

- `unscheduled` - Transactions for a fixed or variable amount that are run at a certain pre-defined event.  
- `recurring` - Transactions for a fixed amount that are run at regular intervals, for example, monthly. Recurring transactions don't have a fixed duration and run until the customer cancels the agreement.  
- `installment` - Transactions for a fixed amount that are run at regular intervals, for example, monthly. Installment transactions have a fixed duration.
    
</dd>
</dl>

<dl>
<dd>

**customer:** `*payroc.Customer` 
    
</dd>
</dl>

<dl>
<dd>

**ipAddress:** `*payroc.IpAddress` 
    
</dd>
</dl>

<dl>
<dd>

**source:** `*tokenization.TokenizationRequestSource` — Object that contains information about the payment method to tokenize.
    
</dd>
</dl>

<dl>
<dd>

**threeDSecure:** `*tokenization.TokenizationRequestThreeDSecure` — Object that contains information for an authentication check on the customer's payment details using the 3-D Secure protocol.
    
</dd>
</dl>

<dl>
<dd>

**customFields:** `[]*payroc.CustomField` — Array of customField objects.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokenization.SecureTokens.Retrieve(ProcessingTerminalId, SecureTokenId) -> *payroc.SecureTokenWithAccountType</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to retrieve information about a secure token.  

To retrieve a secure token, you need its secureTokenID, which you sent in the request of the [Create Secure Token](https://docs.payroc.com/api/schema/tokenization/secure-tokens/create) method.  

**Note:** If you don't have the secureTokenId, use our [List Secure Tokens](https://docs.payroc.com/api/schema/tokenization/secure-tokens/list) method to search for the secure token.  

Our gateway returns the following information about the secure token:  

  -	Payment details that the secure token represents.  
  -	Customer details, including shipping and billing addresses.  
  -	Secure token that you can use to carry out transactions.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tokenization.RetrieveSecureTokensRequest{
        ProcessingTerminalId: "1234001",
        SecureTokenId: "MREF_abc1de23-f4a5-6789-bcd0-12e345678901fa",
    }
client.Tokenization.SecureTokens.Retrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**secureTokenId:** `string` — Unique identifier that the merchant assigned to the secure token.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokenization.SecureTokens.Delete(ProcessingTerminalId, SecureTokenId) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to delete a secure token and its related payment details from our vault.  

To delete a secure token, you need its secureTokenId, which you sent in the request of the [Create Secure Token](https://docs.payroc.com/api/schema/tokenization/secure-tokens/create) method.  

**Note:** If you don’t have the secureTokenId, use our [List Secure Tokens](https://docs.payroc.com/api/schema/tokenization/secure-tokens/list) method to search for the secure token.  

When you delete a secure token, you can’t recover it, and you can’t reuse its identifier for a new token.  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tokenization.DeleteSecureTokensRequest{
        ProcessingTerminalId: "1234001",
        SecureTokenId: "MREF_abc1de23-f4a5-6789-bcd0-12e345678901fa",
    }
client.Tokenization.SecureTokens.Delete(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**secureTokenId:** `string` — Unique identifier that the merchant assigned to the secure token.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokenization.SecureTokens.PartiallyUpdate(ProcessingTerminalId, SecureTokenId, request) -> *payroc.SecureToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to partially update a secure token. Structure your request to follow the [RFC 6902](https://datatracker.ietf.org/doc/html/rfc6902) standard.  

To update a secure token, you need its secureTokenId, which you sent in the request of the [Create Secure Token](https://docs.payroc.com/api/schema/tokenization/secure-tokens/create) method.  

**Note:** If you don't have the secureTokenId, use our [List Secure Tokens](https://docs.payroc.com/api/schema/tokenization/secure-tokens/list) method to search  for the payment.  

You can update all of the properties of the secure token, except the following:  
- processingTerminalId  
- type  
- token  
- status  
- source/Card  
  - type  
  - cardNumber  
  - cardType  
  - currency  
  - debit  
  - surcharging  
- source/ACH account  
  - accountNumber  
  - routingNumber  
- source/PAD account  
  - type  
  - accountNumber  
  - transitNumber  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tokenization.PartiallyUpdateSecureTokensRequest{
        ProcessingTerminalId: "1234001",
        SecureTokenId: "MREF_abc1de23-f4a5-6789-bcd0-12e345678901fa",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: []*payroc.PatchDocument{
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
            &payroc.PatchDocument{
                Remove: &payroc.PatchRemove{
                    Path: "path",
                },
            },
        },
    }
client.Tokenization.SecureTokens.PartiallyUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**secureTokenId:** `string` — Unique identifier that the merchant assigned to the secure token.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `payroc.PatchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tokenization.SecureTokens.UpdateAccount(ProcessingTerminalId, SecureTokenId, request) -> *payroc.SecureToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to update a secure token if you have a single-use token from Hosted Fields.  

**Note:** If you don't have a single-use token, you can update saved payment details with our [Update Secure Token](https://docs.payroc.com/api/resources#updateSecureToken) method. For more information about our two options to update a secure token, go to [Update saved payment details](https://docs.payroc.com/guides/integrate/update-saved-payment-details).  
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tokenization.UpdateAccountSecureTokensRequest{
        ProcessingTerminalId: "1234001",
        SecureTokenId: "MREF_abc1de23-f4a5-6789-bcd0-12e345678901fa",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Body: &payroc.AccountUpdate{
            SingleUseToken: &payroc.SingleUseTokenAccountUpdate{
                Token: "abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
            },
        },
    }
client.Tokenization.SecureTokens.UpdateAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that we assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**secureTokenId:** `string` — Unique identifier that the merchant assigned to the secure token.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**request:** `*payroc.AccountUpdate` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tokenization SingleUseTokens
<details><summary><code>client.Tokenization.SingleUseTokens.Create(ProcessingTerminalId, request) -> *payroc.SingleUseToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this method to create a single-use token that represents a customer’s payment details.  

A single-use token expires after 30 minutes and merchants can use them only once.  

**Note:** To create a reusable permanent token, go to [Create Secure Token](https://docs.payroc.com/api/schema/tokenization/secure-tokens/create).  

In the request, send the customer’s payment details. If the request is successful, our gateway returns a token that you can use in a follow-on action, for example, [run a sale](https://docs.payroc.com/api/schema/card-payments/payments/create).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &tokenization.SingleUseTokenRequest{
        ProcessingTerminalId: "1234001",
        IdempotencyKey: "8e03978e-40d5-43e8-bc93-6894a57f9324",
        Channel: tokenization.SingleUseTokenRequestChannelWeb,
        Operator: payroc.String(
            "Jane",
        ),
        Source: &tokenization.SingleUseTokenRequestSource{
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
    }
client.Tokenization.SingleUseTokens.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**processingTerminalId:** `string` — Unique identifier that our gateway assigned to the terminal.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — Unique identifier that you generate for each request. You must use the [UUID v4 format](https://www.rfc-editor.org/rfc/rfc4122) for the identifier. For more information about the idempotency key, go to [Idempotency](https://docs.payroc.com/api/idempotency).
    
</dd>
</dl>

<dl>
<dd>

**channel:** `*tokenization.SingleUseTokenRequestChannel` — Channel that the merchant used to receive the payment details.
    
</dd>
</dl>

<dl>
<dd>

**operator:** `*string` — Operator who initiated the request.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*tokenization.SingleUseTokenRequestSource` — Object that contains information about the payment method to tokenize.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
