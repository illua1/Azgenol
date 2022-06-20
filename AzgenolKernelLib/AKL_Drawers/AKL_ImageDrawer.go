package AKL_Drawers

import (
	"math"

	values "github.com/illua1/go-helpful"

	"github.com/hajimehoshi/ebiten/v2"
)

type ImageDrawer struct {
	*ebiten.Image
	ebiten.GeoM
}

func NewImageDrawer[Value values.Values](x_size, y_size Value, img *ebiten.Image) (ret ImageDrawer) {
	ret.Image = img
	x, y := img.Size()
	ret.GeoM.Translate(
		(-float64(x) / 2),
		(-float64(y) / 2),
	)
	ret.GeoM.Scale(
		float64(x_size)/float64(x),
		float64(y_size)/float64(y),
	)
	return
}

func NewImageDrawerR[Value values.Values](x_size, y_size Value, rot int, img *ebiten.Image) (ret ImageDrawer) {
	ret.Image = img
	x, y := img.Size()

	var r = ebiten.GeoM{}
	r.Rotate(float64(rot) * math.Pi / 2)
	var x_size_f, y_size_f = r.Apply(float64(x_size), float64(y_size))

	ret.GeoM.Translate(
		(-float64(x) / 2),
		(-float64(y) / 2),
	)
	ret.GeoM.Scale(
		x_size_f/float64(x),
		y_size_f/float64(y),
	)
	return
}

func (iDrawe ImageDrawer) Draw(screen *ebiten.Image, GeomContext ebiten.GeoM) {
	screen.DrawImage(iDrawe.Image, &ebiten.DrawImageOptions{GeoM: iDrawe.GetGeoM(GeomContext)})
}

func (iDrawe ImageDrawer) ToImageDrawer(GeomContext ebiten.GeoM) ImageDrawer {
	iDrawe.GeoM = iDrawe.GetGeoM(GeomContext)
	return iDrawe
}

func (iDrawe ImageDrawer) GetGeoM(GeomContext ebiten.GeoM) ebiten.GeoM {
	iDrawe.GeoM.Concat(GeomContext)
	return iDrawe.GeoM
}

func (iDrawe ImageDrawer) Flip(x, y bool) ImageDrawer {
	var mirror = ebiten.GeoM{}
	if x {
		mirror.Scale(-1, 1)
	}
	if y {
		mirror.Scale(1, -1)
	}
	return iDrawe.ToImageDrawer(mirror)
}

func (iDrawe ImageDrawer) Rot(angle float64) ImageDrawer {
	var mirror = ebiten.GeoM{}
	mirror.Rotate(math.Pi / 2 * angle)
	return iDrawe.ToImageDrawer(mirror)
}
