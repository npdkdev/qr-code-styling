package main

import (
	"fmt"
	"os"

	qr "qr-code-styling"
)

func main() {
	// Create a basic QR code with default styling
	basicQR := qr.NewWithDefaults()
	
	// Update with custom data
	basicQR.Update(&qr.Options{
		Data: "https://www.example.com",
	})

	// Generate SVG
	svgData, err := basicQR.GetRawData(qr.FileExtensionSVG)
	if err != nil {
		fmt.Printf("Error generating SVG: %v\n", err)
		return
	}

	// Save SVG to file
	err = os.WriteFile("basic_qr.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving SVG: %v\n", err)
		return
	}

	// Generate PNG
	pngData, err := basicQR.GetRawData(qr.FileExtensionPNG)
	if err != nil {
		fmt.Printf("Error generating PNG: %v\n", err)
		return
	}

	// Save PNG to file
	err = os.WriteFile("basic_qr.png", pngData, 0644)
	if err != nil {
		fmt.Printf("Error saving PNG: %v\n", err)
		return
	}

	fmt.Println("Basic QR codes generated successfully!")
}