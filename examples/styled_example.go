package main

import (
	"fmt"
	"os"

	qr "qr-code-styling"
)

func main() {
	// Create a styled QR code with custom options
	image := "https://upload.wikimedia.org/wikipedia/commons/5/51/Facebook_f_logo_%282019%29.svg"
	
	styledQR := qr.New(&qr.Options{
		Width:  300,
		Height: 300,
		Type:   qr.DrawTypeSVG,
		Data:   "https://www.facebook.com/",
		Image:  &image,
		QROptions: qr.QROptions{
			TypeNumber:           0, // Auto-detect
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

	// Generate styled SVG
	svgData, err := styledQR.GetRawData(qr.FileExtensionSVG)
	if err != nil {
		fmt.Printf("Error generating styled SVG: %v\n", err)
		return
	}

	// Save to file
	err = os.WriteFile("styled_qr.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving styled SVG: %v\n", err)
		return
	}

	// Also generate as PNG for comparison
	pngData, err := styledQR.GetRawData(qr.FileExtensionPNG)
	if err != nil {
		fmt.Printf("Error generating styled PNG: %v\n", err)
		return
	}

	err = os.WriteFile("styled_qr.png", pngData, 0644)
	if err != nil {
		fmt.Printf("Error saving styled PNG: %v\n", err)
		return
	}

	fmt.Println("Styled QR codes generated successfully!")
}