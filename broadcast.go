package JMAPNs

/*
	Use this function when broadcasting the same payload to
	a large number of devices. It's optimized for sending
	large amounts of the same exact Payload.

	Instead of sending entire Notifications through the normal
	package-wide channels, this function gives you new channels
	to send just Tokens. You supply the Notification up-front
	with all the specified values (DeviceToken will be ignored)
	and it handles pushing them efficiently.
*/
func NewBatchPayload(notification Notification) (SendChannel chan Token, ResponseChannel chan Response, err error) {

	return nil, nil, nil
}
