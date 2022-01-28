package barcode

import (
	"image"
	"image/color"
)

type borderedBarcode struct {
	barcode Barcode
	size    int
}

func (bb *borderedBarcode) At(x, y int) color.Color {
	bounds := bb.barcode.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	if x < bb.size || x >= w+bb.size || y < bb.size || y >= h+bb.size {
		return color.White
	}
	return bb.barcode.At(x-bb.size, y-bb.size)
}

func (bb *borderedBarcode) Bounds() image.Rectangle {
	b := bb.barcode.Bounds()
	return image.Rect(0, 0, b.Dx()+2*bb.size, b.Dy()+2*bb.size)
}

func (bb *borderedBarcode) ColorModel() color.Model {
	return bb.barcode.ColorModel()
}

func GetImageBorder(barcode Barcode, size int) image.Image {
	return &borderedBarcode{barcode, size}
}
