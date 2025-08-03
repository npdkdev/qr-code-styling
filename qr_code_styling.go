// Package qr_code_styling provides a Go implementation of the qr-code-styling library
// for generating QR codes with custom styling and image embedding capabilities.
package qr_code_styling

import (
	"qr-code-styling/pkg/constants"
	"qr-code-styling/pkg/core"
	"qr-code-styling/pkg/types"
)

// QRCodeStyling is the main interface for generating styled QR codes
type QRCodeStyling = core.QRCodeStyling

// New creates a new QRCodeStyling instance with the given options
func New(options *types.Options) *QRCodeStyling {
	return core.New(options)
}

// NewWithDefaults creates a new QRCodeStyling instance with default options
func NewWithDefaults() *QRCodeStyling {
	return core.New(nil)
}

// Export constants for external use
var (
	// Dot Types
	DotTypeDots          = constants.DotTypeDots
	DotTypeRounded       = constants.DotTypeRounded
	DotTypeClassy        = constants.DotTypeClassy
	DotTypeClassyRounded = constants.DotTypeClassyRounded
	DotTypeSquare        = constants.DotTypeSquare
	DotTypeExtraRounded  = constants.DotTypeExtraRounded

	// Draw Types
	DrawTypeCanvas = constants.DrawTypeCanvas
	DrawTypeSVG    = constants.DrawTypeSVG

	// Shape Types
	ShapeTypeSquare = constants.ShapeTypeSquare
	ShapeTypeCircle = constants.ShapeTypeCircle

	// Error Correction Levels
	ErrorCorrectionLevelL = constants.ErrorCorrectionLevelL
	ErrorCorrectionLevelM = constants.ErrorCorrectionLevelM
	ErrorCorrectionLevelQ = constants.ErrorCorrectionLevelQ
	ErrorCorrectionLevelH = constants.ErrorCorrectionLevelH

	// Modes
	ModeNumeric      = constants.ModeNumeric
	ModeAlphanumeric = constants.ModeAlphanumeric
	ModeByte         = constants.ModeByte
	ModeKanji        = constants.ModeKanji

	// Gradient Types
	GradientTypeLinear = constants.GradientTypeLinear
	GradientTypeRadial = constants.GradientTypeRadial

	// File Extensions
	FileExtensionSVG  = constants.FileExtensionSVG
	FileExtensionPNG  = constants.FileExtensionPNG
	FileExtensionJPEG = constants.FileExtensionJPEG
	FileExtensionWEBP = constants.FileExtensionWEBP
)

// Export types for external use
type (
	Options               = types.Options
	QROptions             = types.QROptions
	ImageOptions          = types.ImageOptions
	DotsOptions           = types.DotsOptions
	CornersSquareOptions  = types.CornersSquareOptions
	CornersDotOptions     = types.CornersDotOptions
	BackgroundOptions     = types.BackgroundOptions
	DownloadOptions       = types.DownloadOptions
	Gradient              = types.Gradient
	ColorStop             = types.ColorStop
	DrawArgs              = types.DrawArgs
	GetNeighbor           = types.GetNeighbor
	ExtensionFunction     = types.ExtensionFunction

	DotType               = constants.DotType
	CornerDotType         = constants.CornerDotType
	CornerSquareType      = constants.CornerSquareType
	DrawType              = constants.DrawType
	ShapeType             = constants.ShapeType
	ErrorCorrectionLevel  = constants.ErrorCorrectionLevel
	Mode                  = constants.Mode
	TypeNumber            = constants.TypeNumber
	GradientType          = constants.GradientType
	FileExtension         = constants.FileExtension
)