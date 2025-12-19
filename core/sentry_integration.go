// Sentry integration for Payroc SDK with opt-out mechanism.
package core

import (
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/getsentry/sentry-go"
)

const (
	sentryDSN = "https://75e90fb06d4a55c53c95c96c305be89f@o4505201678483456.ingest.us.sentry.io/4510516981334016"
	redacted  = "[REDACTED]"
)

var (
	initialized     bool
	initializeMutex sync.Mutex

	sensitiveHeaders = map[string]bool{
		"authorization":  true,
		"x-api-key":      true,
		"api-key":        true,
		"token":          true,
		"auth-token":     true,
		"access-token":   true,
		"client-secret":  true,
		"secret":         true,
		"x-access-token": true,
	}

	sensitiveExtraKeywords = []string{"token", "key", "secret", "auth", "password"}

	bearerPattern  = regexp.MustCompile(`(?i)(Authorization['"]?\s*:\s*['"]?Bearer\s+)[^\s'"]+`)
	xAPIKeyPattern = regexp.MustCompile(`(?i)(x-api-key['"]?\s*:\s*['"]?)[^\s'"]+`)
	apiKeyPattern  = regexp.MustCompile(`(?i)(api[_-]?key['"]?\s*[:=]\s*['"]?)[^\s'"]+`)
)

// InitializeSentry initializes Sentry for error tracking and monitoring.
// Users can opt-out by setting the PAYROC_DISABLE_SENTRY environment variable
// to 'true', '1', or 'yes'.
func InitializeSentry() {
	initializeMutex.Lock()
	defer initializeMutex.Unlock()

	if initialized {
		return
	}

	disableSentry := os.Getenv("PAYROC_DISABLE_SENTRY")
	if disableSentry != "" {
		lower := strings.ToLower(disableSentry)
		if lower == "true" || lower == "1" || lower == "yes" {
			return
		}
	}

	environment := os.Getenv("PAYROC_ENVIRONMENT")
	if environment == "" {
		environment = "production"
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
		SendDefaultPII:   false,
		TracesSampleRate: 1.0,
		EnableTracing:    true,
		Release:          getSDKVersion(),
		Environment:      environment,
		BeforeSend:       beforeSend,
	})
	if err != nil {
		// Initialization error - don't break SDK
		return
	}

	initialized = true
}

func beforeSend(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	// Redact sensitive headers from request
	if event.Request != nil && event.Request.Headers != nil {
		for key := range event.Request.Headers {
			if sensitiveHeaders[strings.ToLower(key)] {
				event.Request.Headers[key] = redacted
			}
		}
	}

	// Redact sensitive data from exception messages
	for i := range event.Exception {
		if event.Exception[i].Value != "" {
			value := event.Exception[i].Value
			value = bearerPattern.ReplaceAllString(value, "${1}"+redacted)
			value = xAPIKeyPattern.ReplaceAllString(value, "${1}"+redacted)
			value = apiKeyPattern.ReplaceAllString(value, "${1}"+redacted)
			event.Exception[i].Value = value
		}
	}

	// Redact sensitive data from extra context
	for key := range event.Extra {
		lowerKey := strings.ToLower(key)
		for _, sensitive := range sensitiveExtraKeywords {
			if strings.Contains(lowerKey, sensitive) {
				event.Extra[key] = redacted
				break
			}
		}
	}

	// Redact sensitive data from message
	if event.Message != "" {
		event.Message = bearerPattern.ReplaceAllString(event.Message, "${1}"+redacted)
		event.Message = xAPIKeyPattern.ReplaceAllString(event.Message, "${1}"+redacted)
		event.Message = apiKeyPattern.ReplaceAllString(event.Message, "${1}"+redacted)
	}

	return event
}

func getSDKVersion() string {
	return "payroc-go-sdk@0.0.1"
}

func init() {
	InitializeSentry()
}
