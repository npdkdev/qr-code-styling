package main

import (
	"fmt"
	"os"

	qr "qr-code-styling"
)

func main() {
	fmt.Println("QR Code Styling Go Demo")
	fmt.Println("========================")

	// Create a basic QR code
	qrCode := qr.NewWithDefaults()
	
	// Update with demo data
	err := qrCode.Update(&qr.Options{
		Data:   "https://github.com/qr-code-styling-go",
		Width:  300,
		Height: 300,
		DotsOptions: qr.DotsOptions{
			Color: "#1e40af",
			Type:  qr.DotTypeRounded,
		},
		BackgroundOptions: qr.BackgroundOptions{
			Color: "#f8fafc",
		},
	})
	if err != nil {
		fmt.Printf("Error updating QR code: %v\n", err)
		return
	}

	// Generate SVG
	fmt.Println("Generating SVG...")
	svgData, err := qrCode.GetRawData(qr.FileExtensionSVG)
	if err != nil {
		fmt.Printf("Error generating SVG: %v\n", err)
		return
	}

	// Save SVG to file
	err = os.WriteFile("demo_qr.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving SVG: %v\n", err)
		return
	}

	fmt.Printf("✅ SVG saved to demo_qr.svg (%d bytes)\n", len(svgData))

	// Generate PNG
	fmt.Println("Generating PNG...")
	pngData, err := qrCode.GetRawData(qr.FileExtensionPNG)
	if err != nil {
		fmt.Printf("Error generating PNG: %v\n", err)
		return
	}

	// Save PNG to file
	err = os.WriteFile("demo_qr.png", pngData, 0644)
	if err != nil {
		fmt.Printf("Error saving PNG: %v\n", err)
		return
	}

	fmt.Printf("✅ PNG saved to demo_qr.png (%d bytes)\n", len(pngData))
	fmt.Println("\n🎉 Demo completed successfully!")
}