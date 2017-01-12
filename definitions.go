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

const MaximumPayloadSize = 4 * 1024
