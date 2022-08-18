//shared.go is a file containing definitions that are shared across the other files.
//It includes custom structs, custom types, data, etc.
package signals

var SignalChannels = struct {
	CONTINUE_CHANNEL string
	CANCEL_CHANNEL   string
}{
	CONTINUE_CHANNEL: "CONTINUE_CHANNEL",
	CANCEL_CHANNEL:   "CANCEL_CHANNEL",
}

var CommandTypes = struct {
	CONTINUE string
	CANCEL   string
}{
	CONTINUE: "continue",
	CANCEL:   "cancel",
}

type ContinueSignal struct {
	Command string
}

type CancelSignal struct {
	Command string
}
