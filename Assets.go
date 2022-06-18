package main

import (
	_ "embed"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	//go:embed .\BuildData\Assets\Textures\Nil_Texture.png
	Nil_Texture_png   string
	Nil_Texture, _, _ = ebitenutil.NewImageFromReader(strings.NewReader(Nil_Texture_png))
	//go:embed .\BuildData\Assets\Fonts\calibri-bold.ttf
	Nil_Font_otf string
	Nil_Font     font.Face = nul_font_init(Nil_Font_otf)
)

func nul_font_init(str string) font.Face {
	nil_font, err := opentype.Parse([]byte(str))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ffff, err := opentype.NewFace(
		nil_font,
		&opentype.FaceOptions{
			Size:    24,
			DPI:     72,
			Hinting: font.HintingFull,
		})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return ffff
}

var (
	Exist, _ = os.Executable()
	MainPath = NewPath(Exist)
	Assets   = MainPath.Dir("Assets")
	Textures = Assets.Dir("Textures")
	//FonImage = TextureOpen(Textures, "geno-online-v1c.jpg", Nil_Texture)
	Players                = Textures.Dir("Players")
	First                  = Players.Dir("First")
	Main                   = First.Dir("Main")
	Player_First_body_1_45 = TextureOpen(Main, "body_1_45.png", Nil_Texture)
	Ground                 = Textures.Dir("Ground")
	First_b                = Ground.Dir("First")
	Block_plit_face        = TextureOpen(First_b, "plit_face.png", Nil_Texture)
	Fonts                  = Assets.Dir("Fonts")
	//FonFont = FontOpen(Fonts, "Ranika.otf", Nil_Font)
)

func TextureOpen(p Path, f string, empty_texture *ebiten.Image) *ebiten.Image {
	path, err := p.FileCheck(NewFile(f))
	if err != nil {
		log.Print(err)
		return empty_texture
	}
	texture, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Print(err)
		return empty_texture
	}
	return texture
}

func FontOpen(p Path, f string, empty_texture font.Face) font.Face {
	path, err := p.FileCheck(NewFile(f))
	if err != nil {
		log.Print(err)
		return empty_texture
	}
	file, err := os.ReadFile(path)
	font_data, err := opentype.Parse(file)
	if err != nil {
		log.Print(err)
		return empty_texture
	}
	face, err := opentype.NewFace(
		font_data,
		&opentype.FaceOptions{
			Size:    24,
			DPI:     72,
			Hinting: font.HintingFull,
		})
	if err != nil {
		log.Print(err)
		return nil
	}
	return face
}

type Path string

func NewPath(s string) Path {
	return Path(filepath.Dir(s))
}

func (p Path) Dir(name string) Path {
	return Path(filepath.Join(
		string(p),
		name,
	))
}

func (p Path) FileCheck(f File) (string, error) {
	path := filepath.Join(string(p), f.String())
	_, err := os.Stat(path)
	return path, err
}

type File struct {
	Name, Extension string
}

func NewFile(s string) File {
	return File{
		Name:      strings.TrimSuffix(s, filepath.Ext(s)),
		Extension: filepath.Ext(s),
	}
}

func (f File) String() string {
	return f.Name + f.Extension
}
