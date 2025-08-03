package constants

// DotType represents the different styles of QR code dots
type DotType string

const (
	DotTypeDots          DotType = "dots"
	DotTypeRounded       DotType = "rounded"
	DotTypeClassy        DotType = "classy"
	DotTypeClassyRounded DotType = "classy-rounded"
	DotTypeSquare        DotType = "square"
	DotTypeExtraRounded  DotType = "extra-rounded"
)

// DotTypes provides a map of all available dot types
var DotTypes = map[string]DotType{
	"dots":           DotTypeDots,
	"rounded":        DotTypeRounded,
	"classy":         DotTypeClassy,
	"classyRounded":  DotTypeClassyRounded,
	"square":         DotTypeSquare,
	"extraRounded":   DotTypeExtraRounded,
}