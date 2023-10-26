package firefly

type Framerate uint8

// Enum describing the different available framerates.
const (
	FR24 Framerate = 0b00000000
	FR25 Framerate = 0b01000000
	FR29 Framerate = 0b10000000
	FR30 Framerate = 0b11000000
)

// Returns a string representation of a given Framerate.
func (f Framerate) String() string {

	switch f {
	case FR24:
		return "24 FPS"

	case FR25:
		return "25 FPS"

	case FR29:
		return "29.97 Drop FPS"

	case FR30:
		return "30 FPS"

	default:
		return "Unknown Framerate"
	}
}
