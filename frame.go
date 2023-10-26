package firefly

import (
	"reflect"
)

// The base Frame we will be working from. Default values should error out.
var (
	outFrame [10]byte = [10]byte{
		0xF0, // Start of SysEx message
		0x7F, // Universal Message
		0x7F, // Broadcast
		0x01, // Timecode Type Message
		0x01, // Full Frame
		0xFF, // Hours ^ Framerate
		0xFF, // Minutes
		0xFF, // Seconds
		0xFF, // Frames
		0xF7, // End of SysEx message
	}
)

type EncodedFrame [10]byte

type Frame struct {
	Framerate Framerate
	Hour      uint8
	Minute    uint8
	Second    uint8
	Frame     uint8
}

func (f Frame) Encode() (EncodedFrame, error) {

	if f.Hour > 23 {
		return outFrame, ffe[_ERROR_FRAME_HOUR_OUT_OF_RANGE]
	}
	if f.Minute > 59 {
		return outFrame, errorMinuteOutOfRange
	}
	if f.Second > 59 {
		return outFrame, errorSecondOutOfRange
	}

	var secondCheck uint8 = 0

	switch f.Framerate {
	case FR24:
		secondCheck = 23
		break
	case FR25:
		secondCheck = 24
		break
	case FR29:
		secondCheck = 28
		break
	case FR30:
		secondCheck = 29
	default:
		return outFrame, errorFramerateUnknown
	}

	if f.Frame > secondCheck {
		return outFrame, errorFrameOutOfRange
	}

	outputFrame := outFrame

	outputFrame[5] = f.MergeHourFramerate()
	outputFrame[6] = f.Minute
	outputFrame[7] = f.Second
	outputFrame[8] = f.Frame

	return outputFrame, nil

}

func (f *Frame) MergeHourFramerate() uint8 {
	return f.Hour ^ uint8(f.Framerate)
}

func Decode(tcBytes EncodedFrame) (Frame, error) {
	// We should check to make sure its a valid timecode message
	newFrame := Frame{
		Framerate: FR24,
		Hour:      0,
		Minute:    0,
		Second:    0,
		Frame:     0,
	}

	// Yes all the data SHOULD be the correct length, but not knowing where the data is coming from it could be a shot in the dark
	if len(tcBytes) != 10 {
		return newFrame, errorPacketSize
	}

	if !(reflect.DeepEqual(tcBytes[0:3], outFrame[0:3])) {
		return newFrame, errorFormatUnknown
	}

	if !(reflect.DeepEqual(tcBytes[9], outFrame[9])) {
		return newFrame, errorFormatNoEndOfSYSEX
	}

	fr, hour := SplitHourFramerate(tcBytes)


	newFrame.Framerate = fr
	newFrame.Hour = hour
	newFrame.Minute = tcBytes[6]
	newFrame.Second = tcBytes[7]
	newFrame.Frame = tcBytes[8]

	return newFrame, nil
}

func SplitHourFramerate(frame EncodedFrame) (Framerate, uint8) {

	mergedByte := frame[5]

	frr := mergedByte & uint8(FR30)
	hr := mergedByte & 0b00011111

	var fr Framerate = FR30

	switch frr {
	case 0b00000000:
		fr = FR24
	case 0b01000000:
		fr = FR25
	case 0b10000000:
		fr = FR29
	case 0b11000000:
		fr = FR30
	default:
		fr = FR30
	}

	return fr, hr

}
