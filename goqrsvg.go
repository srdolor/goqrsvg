// Package goqrsvg is an API that makes QR Code to SVG conversions.
package goqrsvg

import (
	"errors"
	"image/color"

	"github.com/ajstarks/svgo/float"
	"github.com/boombuler/barcode"
)

// QrSVG holds the data related to the size, location,
// and block size of the QR Code. Holds unexported fields.
type QrSVG struct {
	qr        barcode.Barcode
	qrWidth   int
	blockSize float64
	startingX float64
	startingY float64
}

// NewQrSVG contructs a QrSVG struct. It takes a QR Code in the form
// of barcode.Barcode and sets the "pixel" or block size of QR Code in
// the SVG file.
func NewQrSVG(qr barcode.Barcode, blockSize float64) QrSVG {
	return QrSVG{
		qr:        qr,
		qrWidth:   qr.Bounds().Max.X,
		blockSize: blockSize,
		startingX: 0,
		startingY: 0,
	}
}

// WriteQrSVG writes the QR Code to SVG.
func (qs *QrSVG) WriteQrSVG(s *svg.SVG) error {
	if qs.qr.Metadata().CodeKind == "QR Code" {
		currY := qs.startingY

		for x := 0; x < qs.qrWidth; x++ {
			currX := qs.startingX
			for y := 0; y < qs.qrWidth; y++ {
				if qs.qr.At(y, x) == color.Black {
					s.Rect(float64(currX), float64(currY), qs.blockSize, qs.blockSize)
				}
				currX += qs.blockSize
			}
			currY += qs.blockSize
		}
		return nil
	}
	return errors.New("can not write to SVG: Not a QR code")
}

// SetStartPoint sets the top left start point of QR Code.
// This takes an X and Y value and then adds four white "blocks"
// to create the "quiet zone" around the QR Code.
func (qs *QrSVG) SetStartPoint(x, y float64) {
	qs.startingX = x + (qs.blockSize * 4)
	qs.startingY = y + (qs.blockSize * 4)
}

// StartQrSVG creates a start for writing an SVG file that
// only contains a barcode. This is similar to the svg.Start() method.
// This fucntion should only be used if you only want to write a QR code
// to the SVG. Otherwise use the regular svg.Start() method to start your
// SVG file.
func (qs *QrSVG) StartQrSVG(s *svg.SVG) {
	width := (float64(qs.qrWidth) * qs.blockSize) + (qs.blockSize * 8)
	qs.SetStartPoint(0, 0)
	s.Start(width, width)
}
