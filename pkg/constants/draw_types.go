package constants

// DrawType represents the output format type
type DrawType string

const (
	DrawTypeCanvas DrawType = "canvas"
	DrawTypeSVG    DrawType = "svg"
)

// DrawTypes provides a map of all available draw types
var DrawTypes = map[string]DrawType{
	"canvas": DrawTypeCanvas,
	"svg":    DrawTypeSVG,
}