package main

import (
	"log"
  "math"
  "image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
  values "github.com/illua1/go-helpful"
  sort "github.com/illua1/go-helpful/Sort"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
  volume "github.com/illua1/go-helpful/Volume"
  //UI "github.com/illua1/Game-UI"
)

type Programm struct {
  //WR WorldRender
}

var(
  r [3]float64
  
  
  Player_I = NewPlayer()
  
  cube = NewCube(200, 200, 10, 0, 0, 0)
  cube2 = NewCube(500, 500, 10, 0, 0, -150)
  
)

func (g *Programm) Update() error {
	
  if ebiten.IsKeyPressed(ebiten.KeyW) {
    Player_I.Body.AccelerateAdd(0, -300, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyS) {
    Player_I.Body.AccelerateAdd(0, 300, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyA) {
    Player_I.Body.AccelerateAdd(-300, 0, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyD) {
    Player_I.Body.AccelerateAdd(300, 0, 0)
  }
  
  
  
  player.Body.AccelerateUpdate()
  player.Body.VelocityUpdate()
  player.Body.LocationUpdate()
  
  Player_I.Body.AccelerateAdd(0, 0, -100)
  Player_I.Body.AccelerateUpdate()
  Player_I.Body.VelocityUpdate()
  Player_I.Body.LocationUpdate()
  
  BoxesColiseTest(&cube.Body, &Player_I.Body)
  BoxesColiseTest(&cube2.Body, &Player_I.Body)
  
  {
    r[0] = math.Pi/4
    r[2] = math.Pi/4
    /*
    x, _ := ebiten.CursorPosition()
    r[2] = float64(x)/100
    */
  }
  for i := range player.Camera.Location {
    player.Camera.Location[i] = Lerp(player.Camera.Location[i], player.Body.Location[i], 0.06)
  }
  player.Camera.Matrix = matrix.Rotate3x3_YXZ[float64](r)
  player.Camera.MatrixInvert = player.Camera.Matrix.Invert()
  
  return nil
}

func (g *Programm) Draw(screen *ebiten.Image) {
  x,y := screen.Size()
  op := &ebiten.DrawImageOptions{}
  ScreenGeom := ebiten.GeoM{}
  ScreenGeom.Scale(float64(sort.MaxF(x, y))/1000, float64(sort.MaxF(x, y))/1000)
  ScreenGeom.Translate(float64(x/2), float64(y/2))
  GeomMatrix := ebiten.GeoM{}
  player.Camera.Matrix.Slise(0,0,2,2).FillTo(GeomMatrix.SetElement)
  GeomMatrix.Translate(GeomMatrix.Apply(-player.Camera.Location[0], -player.Camera.Location[1]))
  /*
  {
    var FloorGeom ebiten.GeoM
    for i := range floor {
      FloorGeom.Reset()
      
      {
        var m = player.Camera.MatrixInvert.Mull(floor[i].Matrix)
        m.Slise(0,0,2,2).FillTo(FloorGeom.SetElement)
      }
      
      ImageCentre(floor[i].Img, 100, 100, &FloorGeom)
      
      {
        location_3d := matrix.Vector[float64, [3]float64]{floor[i].Location}
        location_3d.Sub(matrix.Vector[float64, [3]float64]{player.Camera.Location})
        location := player.Camera.Matrix.MulVector(location_3d)
        FloorGeom.Translate(location.A[0], location.A[1])
      }
      
      FloorGeom.Concat(ScreenGeom)
      
      op.GeoM = FloorGeom
      screen.DrawImage(floor[i].Img, op)
    }
  }
  */
  {
    op.GeoM = ScreenGeom
    cube2.Draw(screen, op, player.Camera.Matrix, &player.Camera)
  }
  {
    op.GeoM = ScreenGeom
    cube.Draw(screen, op, player.Camera.Matrix, &player.Camera)
  }
  {
    var PlayerGeom ebiten.GeoM
    ImageCentre(Player_First_body_1_45, 100, 100, &PlayerGeom)
    
    PlayerGeom.Translate(GeomMatrix.Apply(player.Body.Location[0], player.Body.Location[1]))
    
    PlayerGeom.Concat(ScreenGeom)
    
    op.GeoM = PlayerGeom
    screen.DrawImage(Player_First_body_1_45, op)
  }
  /*
  {
    serface := NewSerface(NewImageDrawer(Block_plit_face))
    serface.Src.SetSize(200, 200)
    serface.SetMatrix(player.Camera.Matrix)
    
    op.GeoM = ScreenGeom
    serface.Draw(screen, op)
    
  }
  
  {
    serface := NewSerface(NewImageDrawer(Block_plit_face))
    serface.Src.SetSize(200, 200)
    serface.SetMatrix(player.Camera.Matrix.Invert().Mull(matrix.Rotate3x3_x[float64](math.Pi/2)))
    
    op.GeoM = ScreenGeom
    serface.Draw(screen, op)
    
  }
  */
  {
    op.GeoM = ScreenGeom
    Player_I.Draw(screen, op, player.Camera.Matrix, &player.Camera)
  }
}

func (g *Programm) Layout(outsideWidth, outsideHeight int) (int, int) {
  return outsideWidth, outsideHeight
}

func main() {
  
	ebiten.SetWindowTitle("Azgenol")
  ebiten.SetWindowResizable(true)
  
  prog := Programm{}
  
	if err := ebiten.RunGame(&prog); err != nil {
		log.Fatal(err)
	}
}

type Player struct {
  Body Boxe
  Camera Camera
}

type Camera struct {
  Location [3]float64
  Matrix matrix.Matrix[float64, [3]float64, [3][3]float64]
  MatrixInvert matrix.Matrix[float64, [3]float64, [3][3]float64]
}

type Boxe struct{
  Body volume.Boxe[float64]
  Physics
}

func NewBoxe(x, y, z float64, physics Physics)Boxe{
  x, y, z = x/2, y/2, z/2
  return Boxe{
    volume.NewBoxe[float64](-x, -y, -z, x, y, z),
    physics,
  }
}

func(boxe *Boxe)Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, worldMatrix matrix.Matrix[float64, [3]float64, [3][3]float64]){
  points := boxe.Body.Points()
  edges := boxe.Body.Edges()
  
  for i := range edges {
    var p1 = worldMatrix.MulVector(matrix.Vector[float64, [3]float64]{
      A:[3]float64{
        float64(points[edges[i][0]].X),
        float64(points[edges[i][0]].Y),
        float64(points[edges[i][0]].Z),
      },
    })
    var p2 = worldMatrix.MulVector(matrix.Vector[float64, [3]float64]{
      A:[3]float64{
        float64(points[edges[i][1]].X),
        float64(points[edges[i][1]].Y),
        float64(points[edges[i][1]].Z),
      },
    })
    x1, y1 := op.GeoM.Apply(p1.A[0], p1.A[1])
    x2, y2 := op.GeoM.Apply(p2.A[0], p2.A[1])
    ebitenutil.DrawLine(screen, x1, y1, x2, y2, color.RGBA{255,255,255,255})
  }
}

type Physics struct {
  Location [3]float64
  Velocity [3]float64
  Accelerate [3]float64
}

func NewPhysics(px, py, pz float64)Physics{
  return Physics{
    Location : [3]float64{px, py, pz},
  }
}

func (physics *Physics)LocationUpdate(){
  for i := range physics.Location {
    physics.Location[i] += physics.Velocity[i]
  }
}

func (physics *Physics)VelocityUpdate(){
  for i := range physics.Velocity {
    physics.Velocity[i] += physics.Accelerate[i]
  }
  velocity := matrix.Vector[float64, [3]float64]{physics.Velocity}
  velocity.Scale(0.05)
  physics.Velocity = velocity.A
}

func (physics *Physics)AccelerateUpdate(){
  accelerate := matrix.Vector[float64, [3]float64]{physics.Accelerate}
  accelerate.Scale(0.3)
  physics.Accelerate = accelerate.A
}

func (physics *Physics)AccelerateAdd(x, y, z float64){
  physics.Accelerate[0] += x
  physics.Accelerate[1] += y
  physics.Accelerate[2] += z
}

func NewCamera(angle [3]float64)Camera{
  return Camera{
    Matrix : matrix.Rotate3x3_YXZ[float64](angle),
  }
}

var(
  player = Player{
    Camera : NewCamera(
      [3]float64{math.Pi/4, 0, math.Pi/4},
    ),
  }
)

var(
  floor []Plit
)

func init() {
  floor = make([]Plit, 400)
  var pointer int
  for x := -10; x < 10; x++ {
    for y := -10; y < 10; y++ {
      floor[pointer] = Plit{
        Img : Block_plit_face,
        Location : [3]float64{float64(x)*100, float64(y)*100, 0},
        Matrix : matrix.Rotate3x3_YXZ[float64](
            [3]float64{
              math.Pi*float64(y)/10,
              -math.Pi*float64(x)/10,
              0,
            },
          ),
      }
      pointer++
    }
  }
}

type Plit struct {
  Img *ebiten.Image
  Location [3]float64
  Matrix matrix.Matrix[float64, [3]float64, [3][3]float64]
}

type ImageDrawer struct {
  GeoM ebiten.GeoM
  Img *ebiten.Image
}

func NewImageDrawer(img *ebiten.Image)ImageDrawer{
  return ImageDrawer{ebiten.GeoM{}, img}
}

func (imageDrawe *ImageDrawer)SetSize(Size_x, Size_y float64){
  x, y := imageDrawe.Img.Size()
  imageDrawe.GeoM.Translate(
    (-float64(x)/2),
    (-float64(y)/2),
  )
  imageDrawe.GeoM.Scale(
    Size_x/float64(x),
    Size_y/float64(y),
  )
}

func (imageDrawe ImageDrawer)Draw(screen *ebiten.Image, Op *ebiten.DrawImageOptions){
  imageDrawe.GeoM.Concat(Op.GeoM)
  Op.GeoM = imageDrawe.GeoM
  screen.DrawImage(imageDrawe.Img, Op)
}

type Serface struct {
  Src ImageDrawer
  Matrix matrix.Matrix[float64, [3]float64, [3][3]float64]
}

func (serface Serface)Draw(screen *ebiten.Image, Op *ebiten.DrawImageOptions){
  serface.Matrix.Slise(0,0,2,2).FillTo(Op.GeoM.SetElement)
  serface.Src.Draw(screen, Op)
}

func (serface *Serface)SetMatrix(matrix matrix.Matrix[float64, [3]float64, [3][3]float64]){
  serface.Matrix = matrix
}

func NewSerface(imageDraver ImageDrawer)Serface {
  return Serface{
    Src : imageDraver,
    Matrix : matrix.Matrix3x3[float64](),
  }
}

func ImageCentre(img *ebiten.Image, x_, y_ float64, edit *ebiten.GeoM) {
  var geom ebiten.GeoM
  x, y := img.Size()
  geom.Translate(
    (-float64(x)/2),
    (-float64(y)/2),
  )
  geom.Scale(
    x_/float64(x),
    y_/float64(y),
  )
  geom.Concat(*edit)
  *edit = geom
}

type Cube struct {
  Body Boxe
}

func NewCube(sx, sy, sz, px, py, pz float64)Cube{
  return Cube{NewBoxe(sx, sy, sz, NewPhysics(px, py, pz))}
}

func (cube *Cube)Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, worldMatrix matrix.Matrix[float64, [3]float64, [3][3]float64], camera *Camera){

  var matrixes = [6]matrix.Matrix[float64, [3]float64, [3][3]float64]{
    matrix.Matrix3x3[float64](),
    matrix.Matrix3x3[float64](),
    matrix.Rotate3x3_x[float64](math.Pi/2),
    matrix.Rotate3x3_x[float64](-math.Pi/2),
    matrix.Rotate3x3_y[float64](math.Pi/2),
    matrix.Rotate3x3_y[float64](-math.Pi/2),
  }
  
  var sizes = cube.Body.Body.FaceArea()
  loca_ := matrix.Vector[float64, [3]float64]{cube.Body.Location}
  loca_.Sub(matrix.Vector[float64, [3]float64]{camera.Location})
  global_location := worldMatrix.MulVector(loca_)
  var geom_location = op.GeoM
  
  is_driw := [6]bool{
    worldMatrix.A[2][2] < 0,
    worldMatrix.A[2][2] >= 0,
    worldMatrix.A[2][1] < 0,
    worldMatrix.A[2][1] >= 0,
    worldMatrix.A[2][0] < 0,
    worldMatrix.A[2][0] >= 0,
  }
  
  {
    sx := op.GeoM.Element(0,0)*global_location.A[0] + op.GeoM.Element(1,0)*global_location.A[1]
    sy := op.GeoM.Element(0,1)*global_location.A[0] + op.GeoM.Element(1,1)*global_location.A[1]
  
    geom_location.Translate(sx, sy)
  }
  
  for i, location := range cube.Body.Body.FaceCentres() {
    if is_driw[i]{
      var location_ = worldMatrix.MulVector(matrix.Vector[float64, [3]float64]{
        A:[3]float64{
          float64(location.X),
          float64(location.Y),
          float64(location.Z),
        },
      })
      op.GeoM = geom_location
      {
        lx := op.GeoM.Element(0,0)*location_.A[0] + op.GeoM.Element(1,0)*location_.A[1]
        ly := op.GeoM.Element(0,1)*location_.A[0] + op.GeoM.Element(1,1)*location_.A[1]
        op.GeoM.Translate(lx, ly)
      }
      serface := NewSerface(NewImageDrawer(Block_plit_face))
      {
        sx := op.GeoM.Element(0,0)*sizes[i][0] + op.GeoM.Element(1,0)*sizes[i][1]
        sy := op.GeoM.Element(0,1)*sizes[i][0] + op.GeoM.Element(1,1)*sizes[i][1]
        serface.Src.SetSize(sx*2, sy*2)
      }
      serface.SetMatrix(worldMatrix.Invert().Mull(matrixes[i]))
      
      serface.Draw(screen, op)
    }
  }
  op.GeoM = geom_location
  cube.Body.Draw(screen, op, worldMatrix)
}

type Stand struct {
  Body Boxe
  Img ImageDrawer
}

type Player_ Stand

func NewPlayer()Player_{
  var img = NewImageDrawer(Player_First_body_1_45)
  img.SetSize(100,100)
  return Player_(Stand{NewBoxe(25,25,25, NewPhysics(0,0,100)), img})
}

func(player *Player_)Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, worldMatrix matrix.Matrix[float64, [3]float64, [3][3]float64], camera *Camera){
  
  
  loca_ := matrix.Vector[float64, [3]float64]{player.Body.Location}
  loca_.Sub(matrix.Vector[float64, [3]float64]{camera.Location})
  global_location := worldMatrix.MulVector(loca_)
  
  
  {
    sx := op.GeoM.Element(0,0)*global_location.A[0] + op.GeoM.Element(1,0)*global_location.A[1]
    sy := op.GeoM.Element(0,1)*global_location.A[0] + op.GeoM.Element(1,1)*global_location.A[1]
  
    op.GeoM.Translate(sx, sy)
  }
  var geom = op.GeoM
  player.Img.Draw(screen, op)
  op.GeoM = geom
  player.Body.Draw(screen, op, worldMatrix)
}

func BoxesColiseTest(a, b *Boxe)bool{
  
  var FacePoints1 = a.Body.FaceCentres()
  var FacePoints2 = b.Body.FaceCentres()
  
  p1 := a.Location
  v1 := a.Velocity
  p2 := b.Location
  v2 := b.Velocity
  
  x_dist := (FacePoints1[0].Z + p1[2]) - (FacePoints2[0].Z + p2[2])
  
  x_velocity := v1[2] - v2[2]
  
  t := (x_dist / x_velocity)
  
  //log.Print(x_dist, x_velocity, t)
  
  if (t < 1)&&(t > 0) {
    
    var sizes1 = a.Body.FaceArea()
    var sizes2 = b.Body.FaceArea()
    
    if (sizes1[0][0]+sizes2[0][0]) > matrix.Abc(p1[0] - p2[0])&&(sizes1[0][1]+sizes2[0][1]) > matrix.Abc(p1[1] - p2[1]) {
      a.Velocity[2] = -a.Velocity[2]
      b.Velocity[2] = 100
      //b.Velocity[2] = 200
      //a.Accelerate[2] = -a.Accelerate[2]
      //b.Accelerate[2] = -b.Accelerate[2]
      
      return true
    }
  }
  return false
}