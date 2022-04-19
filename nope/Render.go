package main

import (
  "image"
  "image/color"
  UI "github.com/illua1/Game-UI"
)

var(
  Screen UI.Rendering = UI.NewLayer(
      UI.NewMaxSize(
        UI.NewImage(FonImage),
      ),
      UI.NewBound(
        UI.NewToBorder(
          UI.NewImage(FonImage),
          UI.ToLeftBorder,
          UI.ToTopBorder,
        ),
        300,
        300,
      ),
      UI.NewBound(
        UI.NewToBorder(
          UI.NewImage(FonImage),
          UI.ToRightBorder,
          UI.ToCentre,
        ),
        300,
        300,
      ),
      UI.NewBound(
        UI.NewToBorder(
          UI.NewFill(255,0,0,50),
          UI.ToCentre,
          UI.ToBottomBorder,
        ),
        700,
        100,
      ),
      UI.NewBound(
        UI.NewToBorder(
          UI.NewLayer(
            UI.NewFill(
              59, 100, 200, 150,
            ),
            UI.NewZone(
              UI.NewBound(
                UI.NewToBorder(
                  UI.NewSet(
                    true,
                    []int{100},
                    UI.NewBorder(
                      UI.NewImage(FonImage),
                      10,
                    ),
                    UI.NewBorder(
                      UI.NewImage(FonImage),
                      10,
                    ),
                    UI.NewBorder(
                      UI.NewImage(FonImage),
                      10,
                    ),
                    UI.NewBorder(
                      UI.NewImage(FonImage),
                      10,
                    ),
                    UI.NewBorder(
                      UI.NewImage(FonImage),
                      10,
                    ),
                    UI.NewBorder(
                      UI.NewImage(FonImage),
                      10,
                    ),
                    UI.NewBorder(
                      UI.NewImage(FonImage),
                      10,
                    ),
                  ),
                UI.ToLeftBorder,
                UI.ToCentre,
                ),
                100,
                100,
              ),
            ),
            UI.NewText(
              UI.TextToRender{
                Location : image.Point{30,30},
                Font : FonFont,
                Str : "ABCD",
                Color : color.RGBA{255,0,255,255},
              },
            ),
          ),
          UI.ToCentre,
          UI.ToTopBorder,
        ),
        600,
        100,
      ),
    )
)