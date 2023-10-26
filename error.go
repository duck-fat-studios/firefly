package firefly

import "fmt"

type FireflyErrorCode uint8

const (
	_ERROR_FRAME_HOUR_OUT_OF_RANGE FireflyErrorCode = iota + 101
	_ERROR_FRAME_MINUTE_OUT_OF_RANGE
	_ERROR_FRAME_SECOND_OUT_OF_RANGE
	_ERROR_FRAME_FRAME_OUT_OF_RANGE
	_ERROR_FRAMERATE_UNKNOWN
	_ERROR_FORMAT_UNKNOWN
	_ERROR_FORMAT_NO_END_OF_SYSEX
	_ERROR_DATA_PACKET_SIZE
)

type fireflyError struct {
	Message string
	Code    FireflyErrorCode
}

func (f fireflyError) Error() string {
	return fmt.Sprintf("Firefly error [%d]: %s", f.Code, f.Message)
}

var ffe = map[FireflyErrorCode]fireflyError{
	_ERROR_FRAME_HOUR_OUT_OF_RANGE:{
		Message: "Hour out of range",
		Code:    _ERROR_FRAME_HOUR_OUT_OF_RANGE,
	},
	_ERROR_FRAME_MINUTE_OUT_OF_RANGE:{
		Message: "Minute our of range",
		Code: _ERROR_FRAME_MINUTE_OUT_OF_RANGE,
	},
}



// var errorHourOutOfRange = fireflyError{
// 	Message: "Hour out of range",
// 	Code:    _ERROR_FRAME_HOUR_OUT_OF_RANGE,
// }
var errorMinuteOutOfRange = fireflyError{
	Message: "Minute out of range",
	Code:    _ERROR_FRAME_MINUTE_OUT_OF_RANGE,
}
var errorSecondOutOfRange = fireflyError{
	Message: "Second out of range",
	Code:    _ERROR_FRAME_SECOND_OUT_OF_RANGE,
}
var errorFrameOutOfRange = fireflyError{
	Message: "Frame out of range",
	Code:    _ERROR_FRAME_FRAME_OUT_OF_RANGE,
}
var errorFramerateUnknown = fireflyError{
	Message: "Frame out of range",
	Code:    _ERROR_FRAMERATE_UNKNOWN,
}
var errorFormatUnknown = fireflyError{
	Message: "Unknown data format",
	Code:    _ERROR_FORMAT_UNKNOWN,
}
var errorFormatNoEndOfSYSEX = fireflyError{
	Message: "End of SYSEX message not found",
	Code:    _ERROR_FORMAT_NO_END_OF_SYSEX,
}
var errorPacketSize = fireflyError{
	Message: "Packet length is incorrect",
	Code:    _ERROR_DATA_PACKET_SIZE,
}

// Returns a slice of error codes and messages
func ShowErrors() {

}
