package main

import (
	"image/png"
	"os"

	"github.com/ggiox/go-barcode/pkg/barcode"
	"github.com/ggiox/go-barcode/pkg/barcode/qr"
)

func main() {
	// Create the barcode
	qrCode, _ := qr.Encode("88888888ccc8888", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// Create a border in image
	img := barcode.GetImageBorder(qrCode, 10)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, img)
}
