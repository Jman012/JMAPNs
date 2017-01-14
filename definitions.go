package JMAPNs

type ResponseStatus int

const (
	RespSuccess             ResponseStatus = 200
	RespBadRequest          ResponseStatus = 400
	RespCertificateError    ResponseStatus = 403
	RespBadMethod           ResponseStatus = 405
	RespTokenNotActive      ResponseStatus = 410
	RespLargePayload        ResponseStatus = 413
	RespHighTokenLoad       ResponseStatus = 429
	RespInternalServerError ResponseStatus = 500
	RespUnavailable         ResponseStatus = 503
)

var ResponseStatusText = map[ResponseStatus]string{
	RespSuccess:             "Success",
	RespBadRequest:          "Bad request",
	RespCertificateError:    "There was an error with the certificate or with the provider authentication token",
	RespBadMethod:           "The request used a bad :method value. Only POST requests are supported.",
	RespTokenNotActive:      "The device token is no longer active for the topic.",
	RespLargePayload:        "The notification payload was too large.",
	RespHighTokenLoad:       "The server received too many requests for the same device token.",
	RespInternalServerError: "Internal server error",
	RespUnavailable:         "The server is shutting down and unavailable.",
}

const MaximumPayloadSize = 4 * 1024

const maximumNumberConns = 8
