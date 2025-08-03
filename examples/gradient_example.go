package main

import (
	"fmt"
	"os"

	qr "qr-code-styling"
)

func main() {
	// Create a QR code with gradient styling
	gradientQR := qr.New(&qr.Options{
		Width:  400,
		Height: 400,
		Type:   qr.DrawTypeSVG,
		Data:   "https://github.com/kozakdenys/qr-code-styling",
		DotsOptions: qr.DotsOptions{
			Type: qr.DotTypeRounded,
			Gradient: &qr.Gradient{
				Type:     qr.GradientTypeLinear,
				Rotation: func() *float64 { r := 45.0; return &r }(), // 45 degrees
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

	// Generate gradient SVG
	svgData, err := gradientQR.GetRawData(qr.FileExtensionSVG)
	if err != nil {
		fmt.Printf("Error generating gradient SVG: %v\n", err)
		return
	}

	// Save to file
	err = os.WriteFile("gradient_qr.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving gradient SVG: %v\n", err)
		return
	}

	fmt.Println("Gradient QR code generated successfully!")
}