package barcode

import (
	"image"
	"image/color"
)

type borderedBarcode struct {
	barcode Barcode
	width   int
}

func (bb *borderedBarcode) At(x, y int) color.Color {
	bounds := bb.barcode.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	if x < bb.width || x >= w+bb.width || y < bb.width || y >= h+bb.width {
		return color.White
	}
	return bb.barcode.At(x-bb.width, y-bb.width)
}

func (bb *borderedBarcode) Bounds() image.Rectangle {
	b := bb.barcode.Bounds()
	return image.Rect(0, 0, b.Dx()+2*bb.width, b.Dy()+2*bb.width)
}

func (bb *borderedBarcode) ColorModel() color.Model {
	return bb.barcode.ColorModel()
}

func GetImageBorder(barcode Barcode, size int) image.Image {
	return &borderedBarcode{barcode, size}
}
