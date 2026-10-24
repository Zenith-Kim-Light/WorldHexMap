package gui

type Zoom struct {
	scaleKey   int
	scaleValue float64
}

var zoomScales = map[int]float64{
	-4: 0.15,
	-3: 0.25,
	-2: 0.5,
	-1: 0.75,
	0:  1.0,
	1:  1.5,
	2:  2.0,
	3:  3.0,
	4:  4.0,
}

type ZoomDirection float64

var ZoomNone ZoomDirection = 0.0
var ZoomIn ZoomDirection = -1.0
var ZoomOut ZoomDirection = +1.0

func InitZoom() Zoom {
	return Zoom{scaleKey: 0, scaleValue: 1.}
}

func (z *Zoom) Scale(direction ZoomDirection) {
	scaleKey := z.scaleKey + int(direction)
	if scaleValue, ok := zoomScales[scaleKey]; ok == true {
		z.scaleKey = scaleKey
		z.scaleValue = scaleValue
	}
}

func (z *Zoom) getScale() float64 {
	// если не заполнили вернуть значение по умолчанию
	if z.scaleValue == 0. {
		return 1.
	}
	return z.scaleValue
}
