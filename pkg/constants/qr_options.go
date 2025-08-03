package constants

// ErrorCorrectionLevel represents QR code error correction levels
type ErrorCorrectionLevel string

const (
	ErrorCorrectionLevelL ErrorCorrectionLevel = "L" // Low: 7% error correction
	ErrorCorrectionLevelM ErrorCorrectionLevel = "M" // Medium: 15% error correction  
	ErrorCorrectionLevelQ ErrorCorrectionLevel = "Q" // Quartile: 25% error correction
	ErrorCorrectionLevelH ErrorCorrectionLevel = "H" // High: 30% error correction
)

// Mode represents QR code encoding modes
type Mode string

const (
	ModeNumeric      Mode = "Numeric"
	ModeAlphanumeric Mode = "Alphanumeric"
	ModeByte         Mode = "Byte"
	ModeKanji        Mode = "Kanji"
)

// TypeNumber represents QR code version (1-40, 0 for auto-detect)
type TypeNumber int

const (
	TypeNumberAuto TypeNumber = 0 // Auto-detect
)

// ShapeType represents the overall shape of the QR code
type ShapeType string

const (
	ShapeTypeSquare ShapeType = "square"
	ShapeTypeCircle ShapeType = "circle"
)

// GradientType represents gradient types for styling
type GradientType string

const (
	GradientTypeLinear GradientType = "linear"
	GradientTypeRadial GradientType = "radial"
)

// FileExtension represents supported output file formats
type FileExtension string

const (
	FileExtensionSVG  FileExtension = "svg"
	FileExtensionPNG  FileExtension = "png"
	FileExtensionJPEG FileExtension = "jpeg"
	FileExtensionWEBP FileExtension = "webp"
)

// QRTypes provides a map of type numbers from 0-40
var QRTypes = func() map[int]TypeNumber {
	types := make(map[int]TypeNumber)
	for i := 0; i <= 40; i++ {
		types[i] = TypeNumber(i)
	}
	return types
}()