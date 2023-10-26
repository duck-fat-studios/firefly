package firefly

import (
	"fmt"
	"reflect"
	"testing"
)

type EncodeData struct {
	Frame          Frame
	expectedResult EncodedFrame
}
type EncodeDataWithError struct {
	Frame         Frame
	expectedError fireflyError
}

type DecodeData struct {
	Raw            [10]byte
	expectedResult Frame
}

type DecodeDataWithError struct {
	Raw [10]byte
	expectedError fireflyError
}

func TestEncode(t *testing.T) {

	t.Log("Passing Conditions:")

	testDataNoError := []EncodeData{
		{Frame{FR24, 1, 2, 3, 4}, EncodedFrame{240, 127, 127, 1, 1, 1, 2, 3, 4, 247}},
		{Frame{FR25, 1, 2, 3, 4}, EncodedFrame{240, 127, 127, 1, 1, 65, 2, 3, 4, 247}},
		{Frame{FR29, 1, 2, 3, 4}, EncodedFrame{240, 127, 127, 1, 1, 129, 2, 3, 4, 247}},
		{Frame{FR30, 12, 22, 32, 14}, EncodedFrame{240, 127, 127, 1, 1, 204, 22, 32, 14, 247}},
	}
	testDataWithError := []EncodeDataWithError{
		{Frame{FR24, 200, 2, 3, 4}, ffe[_ERROR_FRAME_HOUR_OUT_OF_RANGE]},
		{Frame{FR30, 12, 61, 32, 14}, errorMinuteOutOfRange},
		{Frame{FR30, 12, 2, 62, 14}, errorSecondOutOfRange},
		{Frame{FR30, 12, 1, 32, 64}, errorFrameOutOfRange},
		{Frame{FR24, 12, 1, 32, 24}, errorFrameOutOfRange},
		{Frame{FR25, 12, 1, 32, 25}, errorFrameOutOfRange},
		{Frame{FR29, 12, 1, 32, 29}, errorFrameOutOfRange},
		{Frame{FR30, 12, 1, 32, 30}, errorFrameOutOfRange},
	}

	for _, dataum := range testDataNoError {
		result, err := dataum.Frame.Encode()

		if err != nil {
			t.Errorf("Frame %v threw an error: %s", dataum.Frame, err)
		}

		if !reflect.DeepEqual(result, dataum.expectedResult) {
			t.Errorf("[ENCODE] FAIL ENCODE. Expected: %v, received: %v", result, dataum.expectedResult)
		} else {
			t.Logf("[ENCODE] PASS ENOCDE. Expected: %v, received: %v", result, dataum.expectedResult)
		}
	}
	for _, dataum := range testDataWithError {
		_, err := dataum.Frame.Encode()

		if err == nil {
			t.Errorf("[ENCODE] FAIL ERROR CHECK failed. Expected a non nil error. %v", dataum.Frame)
			break
		}

		if err.Error() == dataum.expectedError.Error() {
			t.Logf("[ENCODE] PASS ENCODE ERROR CHECK. Expected: %v, recevied: %v with data %v", dataum.expectedError.Error(), err.Error(), dataum.Frame)
		} else {
			t.Errorf("[ENCODE] FAIL ENCODE ERROR CHECK. Expected: %v, received: %v with data %v", dataum.expectedError.Error(), err.Error(), dataum.Frame)
		}
	}
}

func TestDecode(t *testing.T) {

	fmt.Println(0b10000000 ^ 17)

	testDataNoError := []DecodeData{
		{[10]byte{240, 127, 127, 1, 1, 1, 2, 3, 4, 247},
			Frame{FR24, 1, 2, 3, 4},
		},
		{[10]byte{240, 127, 127, 1, 1, 87, 29, 25, 24, 247},
			Frame{FR25, 23, 29, 25, 24},
		},
		{[10]byte{240, 127, 127, 1, 1, 145, 0, 15, 28, 247},
			Frame{FR29, 17, 0, 15, 28},
		},
		{[10]byte{240, 127, 127, 1, 1, 204, 22, 32, 14, 247},
			Frame{FR30, 12, 22, 32, 14},
		},
	}

	testDataError := []DecodeDataWithError{
		{[10]byte{241, 127, 127, 1, 1, 255, 2, 3, 4, 247}, errorFormatUnknown},
		{[10]byte{240, 127, 127, 1, 1, 255, 2, 3, 4, 24}, errorFormatNoEndOfSYSEX},
	}

	for _, datum := range testDataNoError {

		result, err := Decode(datum.Raw)

		if err != nil {
			t.Errorf("Decoding threw an error %s", err)
		}

		if !reflect.DeepEqual(result, datum.expectedResult) {
			t.Errorf("[DECODE] failed. Expected: %v, recevied: %v", datum.expectedResult, result)
		} else {
			t.Logf("[DECODE] PASS. Expected: %v, recevied: %v", datum.expectedResult, result)
		}
	}

	for _, datum := range testDataError {

		_, err := Decode(datum.Raw)

		if err == nil {
			t.Errorf("[DECODE] FAIL ERROR CHECK. Expected non nil error. Data:%v", datum.Raw)
		}
		
		if err.Error() == datum.expectedError.Error() {
			t.Logf("[DECODE] ERROR CHECK PASS. Expected: %v, recevied: %v", err.Error(), datum.expectedError.Error())
		} else {
			t.Errorf("[DECODE] ERROR CHECK FAIL. Expected %v, recevied %v", err.Error(), datum.expectedError.Error())
		}
	}
}
