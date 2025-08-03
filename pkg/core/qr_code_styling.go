package core

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/fogleman/gg"
	"github.com/skip2/go-qrcode"
	_ "golang.org/x/image/webp"

	"qr-code-styling/pkg/constants"
	"qr-code-styling/pkg/types"
)

// QRCodeStyling represents the main QR code styling generator
type QRCodeStyling struct {
	options   *types.Options
	qrCode    *qrcode.QRCode
	extension types.ExtensionFunction
}

// New creates a new QRCodeStyling instance
func New(userOptions *types.Options) *QRCodeStyling {
	qr := &QRCodeStyling{
		options: MergeOptions(DefaultOptions(), userOptions),
	}
	qr.update()
	return qr
}

// Update updates the QR code with new options
func (qr *QRCodeStyling) Update(userOptions *types.Options) error {
	qr.options = MergeOptions(qr.options, userOptions)
	return qr.update()
}

func (qr *QRCodeStyling) update() error {
	if qr.options.Data == "" {
		return nil
	}

	// Convert error correction level
	var level qrcode.RecoveryLevel
	switch qr.options.QROptions.ErrorCorrectionLevel {
	case constants.ErrorCorrectionLevelL:
		level = qrcode.Low
	case constants.ErrorCorrectionLevelM:
		level = qrcode.Medium
	case constants.ErrorCorrectionLevelQ:
		level = qrcode.High
	case constants.ErrorCorrectionLevelH:
		level = qrcode.Highest
	default:
		level = qrcode.High
	}

	// Generate QR code
	code, err := qrcode.New(qr.options.Data, level)
	if err != nil {
		return fmt.Errorf("failed to generate QR code: %w", err)
	}

	qr.qrCode = code
	return nil
}

// GetRawData generates and returns the QR code as raw data
func (qr *QRCodeStyling) GetRawData(extension constants.FileExtension) ([]byte, error) {
	if qr.qrCode == nil {
		return nil, errors.New("QR code is empty")
	}

	switch extension {
	case constants.FileExtensionSVG:
		return qr.generateSVG()
	case constants.FileExtensionPNG:
		return qr.generatePNG()
	case constants.FileExtensionJPEG:
		return qr.generateJPEG()
	case constants.FileExtensionWEBP:
		return qr.generateWEBP()
	default:
		return nil, fmt.Errorf("unsupported extension: %s", extension)
	}
}

// generateSVG creates styled SVG output
func (qr *QRCodeStyling) generateSVG() ([]byte, error) {
	var buf bytes.Buffer
	canvas := svg.New(&buf)

	size := qr.qrCode.Bitmap()
	moduleCount := len(size)
	cellSize := float64(qr.options.Width-2*qr.options.Margin) / float64(moduleCount)

	canvas.Start(qr.options.Width, qr.options.Height)

	// Draw background
	canvas.Rect(0, 0, qr.options.Width, qr.options.Height, "fill:"+qr.options.BackgroundOptions.Color)

	// Draw QR modules with styling
	for row := 0; row < moduleCount; row++ {
		for col := 0; col < moduleCount; col++ {
			if size[row][col] {
				x := float64(qr.options.Margin) + float64(col)*cellSize
				y := float64(qr.options.Margin) + float64(row)*cellSize
				qr.drawStyledModule(canvas, x, y, cellSize, row, col, moduleCount)
			}
		}
	}

	// Add image if specified
	if qr.options.Image != nil && *qr.options.Image != "" {
		if err := qr.addImageToSVG(canvas); err != nil {
			return nil, fmt.Errorf("failed to add image: %w", err)
		}
	}

	canvas.End()

	// Apply extension if specified
	if qr.extension != nil {
		// Note: Extension would need to be adapted for Go SVG library
	}

	return buf.Bytes(), nil
}

// drawStyledModule draws a single QR module with the specified style
func (qr *QRCodeStyling) drawStyledModule(canvas *svg.SVG, x, y, size float64, row, col, moduleCount int) {
	switch qr.options.DotsOptions.Type {
	case constants.DotTypeSquare:
		canvas.Rect(int(x), int(y), int(size), int(size), "fill:"+qr.options.DotsOptions.Color)
	case constants.DotTypeRounded:
		radius := size / 4
		canvas.Roundrect(int(x), int(y), int(size), int(size), int(radius), int(radius), "fill:"+qr.options.DotsOptions.Color)
	case constants.DotTypeDots:
		radius := size / 2
		centerX := x + size/2
		centerY := y + size/2
		canvas.Circle(int(centerX), int(centerY), int(radius), "fill:"+qr.options.DotsOptions.Color)
	case constants.DotTypeClassy:
		// Implement classy style - slightly rounded corners
		radius := size / 8
		canvas.Roundrect(int(x), int(y), int(size), int(size), int(radius), int(radius), "fill:"+qr.options.DotsOptions.Color)
	case constants.DotTypeClassyRounded:
		// Implement classy-rounded style
		radius := size / 3
		canvas.Roundrect(int(x), int(y), int(size), int(size), int(radius), int(radius), "fill:"+qr.options.DotsOptions.Color)
	case constants.DotTypeExtraRounded:
		// Implement extra-rounded style
		radius := size / 2.5
		canvas.Roundrect(int(x), int(y), int(size), int(size), int(radius), int(radius), "fill:"+qr.options.DotsOptions.Color)
	default:
		canvas.Rect(int(x), int(y), int(size), int(size), "fill:"+qr.options.DotsOptions.Color)
	}
}

// addImageToSVG adds an image to the center of the QR code
func (qr *QRCodeStyling) addImageToSVG(canvas *svg.SVG) error {
	if qr.options.Image == nil {
		return nil
	}

	imageSize := float64(qr.options.Width) * qr.options.ImageOptions.ImageSize
	x := (float64(qr.options.Width) - imageSize) / 2
	y := (float64(qr.options.Height) - imageSize) / 2

	// For now, we'll embed as a reference. In a full implementation,
	// you might want to fetch and embed the image as base64
	if strings.HasPrefix(*qr.options.Image, "http") {
		canvas.Image(int(x), int(y), int(imageSize), int(imageSize), *qr.options.Image)
	} else {
		// Local file or base64 data
		canvas.Image(int(x), int(y), int(imageSize), int(imageSize), *qr.options.Image)
	}

	return nil
}

// generatePNG creates PNG output using canvas-like approach
func (qr *QRCodeStyling) generatePNG() ([]byte, error) {
	dc := gg.NewContext(qr.options.Width, qr.options.Height)

	// Set background
	dc.SetHexColor(strings.TrimPrefix(qr.options.BackgroundOptions.Color, "#"))
	dc.Clear()

	size := qr.qrCode.Bitmap()
	moduleCount := len(size)
	cellSize := float64(qr.options.Width-2*qr.options.Margin) / float64(moduleCount)

	// Set foreground color
	dc.SetHexColor(strings.TrimPrefix(qr.options.DotsOptions.Color, "#"))

	// Draw QR modules
	for row := 0; row < moduleCount; row++ {
		for col := 0; col < moduleCount; col++ {
			if size[row][col] {
				x := float64(qr.options.Margin) + float64(col)*cellSize
				y := float64(qr.options.Margin) + float64(row)*cellSize
				qr.drawStyledModuleCanvas(dc, x, y, cellSize)
			}
		}
	}

	// Add image if specified
	if qr.options.Image != nil && *qr.options.Image != "" {
		if err := qr.addImageToCanvas(dc); err != nil {
			return nil, fmt.Errorf("failed to add image: %w", err)
		}
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		return nil, fmt.Errorf("failed to encode PNG: %w", err)
	}

	return buf.Bytes(), nil
}

// drawStyledModuleCanvas draws a module on canvas context
func (qr *QRCodeStyling) drawStyledModuleCanvas(dc *gg.Context, x, y, size float64) {
	switch qr.options.DotsOptions.Type {
	case constants.DotTypeSquare:
		dc.DrawRectangle(x, y, size, size)
		dc.Fill()
	case constants.DotTypeRounded:
		radius := size / 4
		dc.DrawRoundedRectangle(x, y, size, size, radius)
		dc.Fill()
	case constants.DotTypeDots:
		radius := size / 2
		centerX := x + size/2
		centerY := y + size/2
		dc.DrawCircle(centerX, centerY, radius)
		dc.Fill()
	case constants.DotTypeClassy:
		radius := size / 8
		dc.DrawRoundedRectangle(x, y, size, size, radius)
		dc.Fill()
	case constants.DotTypeClassyRounded:
		radius := size / 3
		dc.DrawRoundedRectangle(x, y, size, size, radius)
		dc.Fill()
	case constants.DotTypeExtraRounded:
		radius := size / 2.5
		dc.DrawRoundedRectangle(x, y, size, size, radius)
		dc.Fill()
	default:
		dc.DrawRectangle(x, y, size, size)
		dc.Fill()
	}
}

// addImageToCanvas adds an image to the canvas
func (qr *QRCodeStyling) addImageToCanvas(dc *gg.Context) error {
	if qr.options.Image == nil {
		return nil
	}

	var img image.Image
	var err error

	if strings.HasPrefix(*qr.options.Image, "http") {
		// Download image
		resp, err := http.Get(*qr.options.Image)
		if err != nil {
			return fmt.Errorf("failed to download image: %w", err)
		}
		defer resp.Body.Close()

		img, _, err = image.Decode(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to decode downloaded image: %w", err)
		}
	} else if strings.HasPrefix(*qr.options.Image, "data:") {
		// Base64 encoded image
		parts := strings.Split(*qr.options.Image, ",")
		if len(parts) != 2 {
			return errors.New("invalid base64 image format")
		}

		data, decodeErr := base64.StdEncoding.DecodeString(parts[1])
		if decodeErr != nil {
			return fmt.Errorf("failed to decode base64 image: %w", decodeErr)
		}

		img, _, err = image.Decode(bytes.NewReader(data))
		if err != nil {
			return fmt.Errorf("failed to decode base64 image: %w", err)
		}
	} else {
		return errors.New("unsupported image source")
	}

	// Calculate image position and size
	imageSize := float64(qr.options.Width) * qr.options.ImageOptions.ImageSize
	x := (float64(qr.options.Width) - imageSize) / 2
	y := (float64(qr.options.Height) - imageSize) / 2

	dc.DrawImageAnchored(img, int(x+imageSize/2), int(y+imageSize/2), 0.5, 0.5)

	return nil
}

// generateJPEG creates JPEG output
func (qr *QRCodeStyling) generateJPEG() ([]byte, error) {
	dc := gg.NewContext(qr.options.Width, qr.options.Height)

	// Set background
	dc.SetHexColor(strings.TrimPrefix(qr.options.BackgroundOptions.Color, "#"))
	dc.Clear()

	size := qr.qrCode.Bitmap()
	moduleCount := len(size)
	cellSize := float64(qr.options.Width-2*qr.options.Margin) / float64(moduleCount)

	// Set foreground color
	dc.SetHexColor(strings.TrimPrefix(qr.options.DotsOptions.Color, "#"))

	// Draw QR modules
	for row := 0; row < moduleCount; row++ {
		for col := 0; col < moduleCount; col++ {
			if size[row][col] {
				x := float64(qr.options.Margin) + float64(col)*cellSize
				y := float64(qr.options.Margin) + float64(row)*cellSize
				qr.drawStyledModuleCanvas(dc, x, y, cellSize)
			}
		}
	}

	// Add image if specified
	if qr.options.Image != nil && *qr.options.Image != "" {
		if err := qr.addImageToCanvas(dc); err != nil {
			return nil, fmt.Errorf("failed to add image: %w", err)
		}
	}

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, dc.Image(), &jpeg.Options{Quality: 100}); err != nil {
		return nil, fmt.Errorf("failed to encode JPEG: %w", err)
	}

	return buf.Bytes(), nil
}

// generateWEBP creates WEBP output  
func (qr *QRCodeStyling) generateWEBP() ([]byte, error) {
	// For now, return PNG data since webp encoding is not straightforward
	// In a full implementation, you'd use a proper webp encoder
	return qr.generatePNG()
}

// ApplyExtension applies a custom extension function
func (qr *QRCodeStyling) ApplyExtension(extension types.ExtensionFunction) {
	qr.extension = extension
}

// DeleteExtension removes the current extension
func (qr *QRCodeStyling) DeleteExtension() {
	qr.extension = nil
}

// SaveToFile saves the QR code to a file
func (qr *QRCodeStyling) SaveToFile(filename string, extension constants.FileExtension) error {
	data, err := qr.GetRawData(extension)
	if err != nil {
		return err
	}

	return saveToFile(filename, data)
}

// saveToFile is a helper function to save data to file
func saveToFile(filename string, data []byte) error {
	// This would implement file saving logic
	// For now, return an error indicating it needs implementation
	return errors.New("file saving not implemented in this example")
}