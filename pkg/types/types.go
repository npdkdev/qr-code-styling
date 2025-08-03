package types

import (
	"qr-code-styling/pkg/constants"
)

// ColorStop represents a gradient color stop
type ColorStop struct {
	Offset float64 `json:"offset"`
	Color  string  `json:"color"`
}

// Gradient represents gradient styling options
type Gradient struct {
	Type       constants.GradientType `json:"type"`
	Rotation   *float64               `json:"rotation,omitempty"`
	ColorStops []ColorStop            `json:"colorStops"`
}

// QROptions represents QR code generation options
type QROptions struct {
	TypeNumber           constants.TypeNumber           `json:"typeNumber,omitempty"`
	Mode                 *constants.Mode                `json:"mode,omitempty"`
	ErrorCorrectionLevel constants.ErrorCorrectionLevel `json:"errorCorrectionLevel"`
}

// ImageOptions represents image embedding options
type ImageOptions struct {
	SaveAsBlob         bool    `json:"saveAsBlob"`
	HideBackgroundDots bool    `json:"hideBackgroundDots"`
	ImageSize          float64 `json:"imageSize"`
	CrossOrigin        *string `json:"crossOrigin,omitempty"`
	Margin             int     `json:"margin"`
}

// DotsOptions represents dot styling options
type DotsOptions struct {
	Type      constants.DotType `json:"type"`
	Color     string            `json:"color"`
	Gradient  *Gradient         `json:"gradient,omitempty"`
	RoundSize *bool             `json:"roundSize,omitempty"`
}

// CornersSquareOptions represents corner square styling options
type CornersSquareOptions struct {
	Type     *constants.CornerSquareType `json:"type,omitempty"`
	Color    *string                     `json:"color,omitempty"`
	Gradient *Gradient                   `json:"gradient,omitempty"`
}

// CornersDotOptions represents corner dot styling options
type CornersDotOptions struct {
	Type     *constants.CornerDotType `json:"type,omitempty"`
	Color    *string                  `json:"color,omitempty"`
	Gradient *Gradient                `json:"gradient,omitempty"`
}

// BackgroundOptions represents background styling options
type BackgroundOptions struct {
	Round    float64   `json:"round"`
	Color    string    `json:"color"`
	Gradient *Gradient `json:"gradient,omitempty"`
}

// Options represents all QR code styling options
type Options struct {
	Type                 constants.DrawType        `json:"type"`
	Shape                constants.ShapeType       `json:"shape"`
	Width                int                       `json:"width"`
	Height               int                       `json:"height"`
	Margin               int                       `json:"margin"`
	Data                 string                    `json:"data"`
	Image                *string                   `json:"image,omitempty"`
	QROptions            QROptions                 `json:"qrOptions"`
	ImageOptions         ImageOptions              `json:"imageOptions"`
	DotsOptions          DotsOptions               `json:"dotsOptions"`
	CornersSquareOptions *CornersSquareOptions     `json:"cornersSquareOptions,omitempty"`
	CornersDotOptions    *CornersDotOptions        `json:"cornersDotOptions,omitempty"`
	BackgroundOptions    BackgroundOptions         `json:"backgroundOptions"`
}

// DownloadOptions represents file download options
type DownloadOptions struct {
	Name      string                    `json:"name"`
	Extension constants.FileExtension  `json:"extension"`
}

// DrawArgs represents arguments for drawing functions
type DrawArgs struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Size      float64 `json:"size"`
	Rotation  float64 `json:"rotation"`
}

// GetNeighbor is a function type to check neighboring dots
type GetNeighbor func(x, y int) bool

// ExtensionFunction is a function type for custom SVG extensions
type ExtensionFunction func(svg interface{}, options *Options)