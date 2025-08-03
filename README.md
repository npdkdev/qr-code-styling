# QR Code Styling - Go

A Go implementation of the popular [qr-code-styling](https://github.com/kozakdenys/qr-code-styling) JavaScript library for generating QR codes with custom styling and image embedding capabilities.

## Features

- ✅ **Multiple Output Formats**: SVG, PNG, JPEG, WebP
- ✅ **Custom Styling**: Various dot types (square, rounded, classy, dots, etc.)
- ✅ **Image Embedding**: Add logos or images to QR codes
- ✅ **Gradient Support**: Linear and radial gradients for dots and background
- ✅ **Corner Customization**: Custom styling for corner squares and dots
- ✅ **Error Correction**: Full support for L, M, Q, H error correction levels
- ✅ **Shape Options**: Square and circle QR code shapes
- ✅ **Extensions**: Support for custom SVG extensions

## Installation

```bash
go get github.com/your-username/qr-code-styling
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "os"
    qr "qr-code-styling"
)

func main() {
    // Create a basic QR code
    qrCode := qr.NewWithDefaults()
    
    // Update with your data
    qrCode.Update(&qr.Options{
        Data: "https://www.example.com",
    })

    // Generate SVG
    svgData, err := qrCode.GetRawData(qr.FileExtensionSVG)
    if err != nil {
        panic(err)
    }

    // Save to file
    os.WriteFile("qr_code.svg", svgData, 0644)
    fmt.Println("QR code generated successfully!")
}
```

### Advanced Styling

```go
package main

import (
    "os"
    qr "qr-code-styling"
)

func main() {
    image := "https://example.com/logo.png"
    
    qrCode := qr.New(&qr.Options{
        Width:  300,
        Height: 300,
        Type:   qr.DrawTypeSVG,
        Data:   "https://www.facebook.com/",
        Image:  &image,
        QROptions: qr.QROptions{
            ErrorCorrectionLevel: qr.ErrorCorrectionLevelQ,
        },
        ImageOptions: qr.ImageOptions{
            HideBackgroundDots: true,
            ImageSize:          0.4,
            Margin:             20,
        },
        DotsOptions: qr.DotsOptions{
            Color: "#4267b2",
            Type:  qr.DotTypeRounded,
        },
        BackgroundOptions: qr.BackgroundOptions{
            Color: "#e9ebee",
        },
    })

    data, _ := qrCode.GetRawData(qr.FileExtensionSVG)
    os.WriteFile("styled_qr.svg", data, 0644)
}
```

### Gradient Styling

```go
package main

import (
    "os"
    qr "qr-code-styling"
)

func main() {
    qrCode := qr.New(&qr.Options{
        Width:  400,
        Height: 400,
        Type:   qr.DrawTypeSVG,
        Data:   "https://github.com/example/repo",
        DotsOptions: qr.DotsOptions{
            Type: qr.DotTypeRounded,
            Gradient: &qr.Gradient{
                Type:     qr.GradientTypeLinear,
                Rotation: func() *float64 { r := 45.0; return &r }(),
                ColorStops: []qr.ColorStop{
                    {Offset: 0, Color: "#ff6b6b"},
                    {Offset: 1, Color: "#4ecdc4"},
                },
            },
        },
        BackgroundOptions: qr.BackgroundOptions{
            Gradient: &qr.Gradient{
                Type: qr.GradientTypeRadial,
                ColorStops: []qr.ColorStop{
                    {Offset: 0, Color: "#ffffff"},
                    {Offset: 1, Color: "#f0f0f0"},
                },
            },
        },
    })

    data, _ := qrCode.GetRawData(qr.FileExtensionSVG)
    os.WriteFile("gradient_qr.svg", data, 0644)
}
```

## API Reference

### QRCodeStyling

The main struct for generating styled QR codes.

#### Constructor Functions

- `New(options *Options) *QRCodeStyling` - Create with custom options
- `NewWithDefaults() *QRCodeStyling` - Create with default options

#### Methods

- `Update(options *Options) error` - Update QR code with new options
- `GetRawData(extension FileExtension) ([]byte, error)` - Generate QR code data
- `ApplyExtension(extension ExtensionFunction)` - Apply custom SVG extension
- `DeleteExtension()` - Remove current extension

### Options Structure

```go
type Options struct {
    Type                 DrawType              // "canvas" or "svg"
    Shape                ShapeType             // "square" or "circle"
    Width                int                   // QR code width
    Height               int                   // QR code height
    Margin               int                   // Margin around QR code
    Data                 string                // Data to encode
    Image                *string               // Image URL or base64
    QROptions            QROptions             // QR generation options
    ImageOptions         ImageOptions          // Image embedding options
    DotsOptions          DotsOptions           // Dot styling options
    CornersSquareOptions *CornersSquareOptions // Corner square styling
    CornersDotOptions    *CornersDotOptions    // Corner dot styling
    BackgroundOptions    BackgroundOptions     // Background styling
}
```

### Dot Types

- `DotTypeSquare` - Square dots
- `DotTypeRounded` - Rounded corners
- `DotTypeDots` - Circular dots
- `DotTypeClassy` - Slightly rounded
- `DotTypeClassyRounded` - More rounded
- `DotTypeExtraRounded` - Heavily rounded

### File Extensions

- `FileExtensionSVG` - SVG format
- `FileExtensionPNG` - PNG format
- `FileExtensionJPEG` - JPEG format
- `FileExtensionWEBP` - WebP format

### Error Correction Levels

- `ErrorCorrectionLevelL` - Low (7%)
- `ErrorCorrectionLevelM` - Medium (15%)
- `ErrorCorrectionLevelQ` - Quartile (25%)
- `ErrorCorrectionLevelH` - High (30%)

## Examples

Check the `examples/` directory for complete working examples:

- `basic_example.go` - Basic QR code generation
- `styled_example.go` - Advanced styling with images
- `gradient_example.go` - Gradient styling examples

## Dependencies

This library uses the following Go packages:

- `github.com/skip2/go-qrcode` - QR code generation
- `github.com/ajstarks/svgo` - SVG generation
- `github.com/fogleman/gg` - 2D graphics rendering
- `golang.org/x/image` - Image processing

## Migration from TypeScript

This Go version maintains API compatibility with the original TypeScript library where possible. Key differences:

1. **Constructor**: Use `New()` or `NewWithDefaults()` instead of `new QRCodeStyling()`
2. **Error Handling**: Go methods return errors instead of throwing exceptions
3. **File Operations**: Use `GetRawData()` and save manually instead of `download()`
4. **Pointers**: Optional fields use pointers (`*string`, `*int`, etc.)

## License

MIT License - see LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
