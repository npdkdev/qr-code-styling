package core

import (
	"qr-code-styling/pkg/constants"
	"qr-code-styling/pkg/types"
)

// DefaultOptions provides the default configuration for QR code generation
func DefaultOptions() *types.Options {
	roundSize := true
	return &types.Options{
		Type:   constants.DrawTypeCanvas,
		Shape:  constants.ShapeTypeSquare,
		Width:  300,
		Height: 300,
		Data:   "",
		Margin: 0,
		QROptions: types.QROptions{
			TypeNumber:           constants.TypeNumberAuto,
			Mode:                 nil,
			ErrorCorrectionLevel: constants.ErrorCorrectionLevelQ,
		},
		ImageOptions: types.ImageOptions{
			SaveAsBlob:         true,
			HideBackgroundDots: true,
			ImageSize:          0.4,
			CrossOrigin:        nil,
			Margin:             0,
		},
		DotsOptions: types.DotsOptions{
			Type:      constants.DotTypeSquare,
			Color:     "#000",
			RoundSize: &roundSize,
		},
		BackgroundOptions: types.BackgroundOptions{
			Round: 0,
			Color: "#fff",
		},
	}
}

// MergeOptions merges user options with default options
func MergeOptions(defaults *types.Options, userOptions *types.Options) *types.Options {
	if userOptions == nil {
		return defaults
	}

	result := *defaults // Copy defaults

	// Override with user options if provided
	if userOptions.Type != "" {
		result.Type = userOptions.Type
	}
	if userOptions.Shape != "" {
		result.Shape = userOptions.Shape
	}
	if userOptions.Width != 0 {
		result.Width = userOptions.Width
	}
	if userOptions.Height != 0 {
		result.Height = userOptions.Height
	}
	if userOptions.Margin != 0 {
		result.Margin = userOptions.Margin
	}
	if userOptions.Data != "" {
		result.Data = userOptions.Data
	}
	if userOptions.Image != nil {
		result.Image = userOptions.Image
	}

	// Merge QR options
	if userOptions.QROptions.TypeNumber != 0 {
		result.QROptions.TypeNumber = userOptions.QROptions.TypeNumber
	}
	if userOptions.QROptions.Mode != nil {
		result.QROptions.Mode = userOptions.QROptions.Mode
	}
	if userOptions.QROptions.ErrorCorrectionLevel != "" {
		result.QROptions.ErrorCorrectionLevel = userOptions.QROptions.ErrorCorrectionLevel
	}

	// Merge image options
	if userOptions.ImageOptions.ImageSize != 0 {
		result.ImageOptions.ImageSize = userOptions.ImageOptions.ImageSize
	}
	if userOptions.ImageOptions.Margin != 0 {
		result.ImageOptions.Margin = userOptions.ImageOptions.Margin
	}
	if userOptions.ImageOptions.CrossOrigin != nil {
		result.ImageOptions.CrossOrigin = userOptions.ImageOptions.CrossOrigin
	}

	// Merge dots options
	if userOptions.DotsOptions.Type != "" {
		result.DotsOptions.Type = userOptions.DotsOptions.Type
	}
	if userOptions.DotsOptions.Color != "" {
		result.DotsOptions.Color = userOptions.DotsOptions.Color
	}
	if userOptions.DotsOptions.Gradient != nil {
		result.DotsOptions.Gradient = userOptions.DotsOptions.Gradient
	}
	if userOptions.DotsOptions.RoundSize != nil {
		result.DotsOptions.RoundSize = userOptions.DotsOptions.RoundSize
	}

	// Merge background options
	if userOptions.BackgroundOptions.Round != 0 {
		result.BackgroundOptions.Round = userOptions.BackgroundOptions.Round
	}
	if userOptions.BackgroundOptions.Color != "" {
		result.BackgroundOptions.Color = userOptions.BackgroundOptions.Color
	}
	if userOptions.BackgroundOptions.Gradient != nil {
		result.BackgroundOptions.Gradient = userOptions.BackgroundOptions.Gradient
	}

	// Merge corner options
	if userOptions.CornersSquareOptions != nil {
		result.CornersSquareOptions = userOptions.CornersSquareOptions
	}
	if userOptions.CornersDotOptions != nil {
		result.CornersDotOptions = userOptions.CornersDotOptions
	}

	return &result
}