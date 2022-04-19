package main

import(
  "math"
	"github.com/hajimehoshi/ebiten/v2"
  volume "github.com/illua1/go-helpful/Volume"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type Serface_3d struct {
  Matrixes matrix.Matrix[float64, [3]float64, [3][3]float64]
  Src *ebiten.Image
}

func (s3d *Serface_3d) Render(screen *ebiten.Image, wm *matrix.Matrix[float64, [3]float64, [3][3]float64], location [2]float64) {
  var op = &ebiten.DrawImageOptions{}
  mat_invert := wm.Mull(s3d.Matrixes)
  mat_invert.Slise(0,0,2,2).FillTo(op.GeoM.SetElement)
  op.GeoM.Translate(location[0], location[1])
  screen.DrawImage(s3d.Src, op)
}

type Boxe_3d struct {
  Matrixes matrix.Matrix[float64, [3]float64, [3][3]float64]
  Body volume.Boxe[int]
  Serfaces []Serface_3d
}

func NewBoxe(x, y, z int, angle [3]float64, img []*ebiten.Image) Boxe_3d {
  return Boxe_3d{
    Matrixes : matrix.Rotate3x3_YXZ[float64](angle),
    Body : volume.NewBoxe(-x/2, -y/2, -z/2, x/2, y/2, z/2),
    Serfaces : []Serface_3d{
      Serface_3d{
        Matrixes : matrix.Matrix3x3[float64](),
        Src : img[0],
      },
      Serface_3d{
        Matrixes : matrix.Matrix3x3[float64](),
        Src : img[0],
      },
      Serface_3d{
        Matrixes : matrix.Rotate3x3_x[float64](math.Pi/2),
        Src : img[0],
      },
      Serface_3d{
        Matrixes : matrix.Rotate3x3_x[float64](-math.Pi/2),
        Src : img[0],
      },
      Serface_3d{
        Matrixes : matrix.Rotate3x3_y[float64](math.Pi/2),
        Src : img[0],
      },
      Serface_3d{
        Matrixes : matrix.Rotate3x3_y[float64](-math.Pi/2),
        Src : img[0],
      },
    },
  }
}

func (b3d *Boxe_3d) Render(screen *ebiten.Image, wm *matrix.Matrix[float64, [3]float64, [3][3]float64], wm_invert *matrix.Matrix[float64, [3]float64, [3][3]float64], location [2]float64) {
  for i, p := range b3d.Body.FaceCentres() {
    location_i := wm.MulVector(
      matrix.Vector[float64, [3]float64]{
        A:[3]float64{
          float64(p.X),
          float64(p.Y),
          float64(p.Z),
        },
      },
    )
    location[0] += location_i.A[0]
    location[1] += location_i.A[1]
    var mat = wm_invert.Mull(b3d.Matrixes)
    b3d.Serfaces[i].Render(screen, &mat, location)
  }
}

type RO_3d_Ground struct {
  Body Boxe_3d
  Location [2]int
  WorldMatrix, WorldMatrix_invert *matrix.Matrix[float64, [3]float64, [3][3]float64]
}

func (r3d RO_3d_Ground) Render(screen *ebiten.Image, centre_location [2]int) {
  location_i := r3d.WorldMatrix.MulVector(
    matrix.Vector[float64, [3]float64]{
      A:[3]float64{
        float64(r3d.Location[0]),
        float64(r3d.Location[1]),
        0.0,
      },
    },
  )
  r3d.Body.Render(screen, r3d.WorldMatrix, r3d.WorldMatrix_invert, [2]float64{
    float64(centre_location[0])+location_i.A[0],
    float64(centre_location[1])+location_i.A[1],
  })
}

func NewBlock(x, y int, camera Camera) RO_3d_Ground {
  return RO_3d_Ground{
    Body : NewBoxe(250, 250, 250, [3]float64{}, []*ebiten.Image{Block_plit_face}),
    Location : [2]int{x,y},
    WorldMatrix : &camera.Matrix,
    WorldMatrix_invert : &camera.Matrix_Invert,
  }
}