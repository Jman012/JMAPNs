package JMAPNs

type ResponseStatus int

const (
	Success             ResponseStatus = 200
	BadRequest          ResponseStatus = 400
	CertificateError    ResponseStatus = 403
	BadMethod           ResponseStatus = 405
	TokenNotActive      ResponseStatus = 410
	LargePayload        ResponseStatus = 413
	HighTokenLoad       ResponseStatus = 429
	InternalServerError ResponseStatus = 500
	Unavailable         ResponseStatus = 503
)

const MaximumPayloadSize = 4 * 1024
