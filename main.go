package main

import (
  "math/rand"
  "time"
  rl "github.com/gen2brain/raylib-go/raylib"
)

var opened = true
var frame = 0

func draw(randomNumber int, randWidth int32, randHeight int32, randWidth2 int32, randHeight2 int32, texture *rl.Texture2D, bg *rl.Sound, scream *rl.Sound) {

  screenWidth := int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
  // screenHeight := int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
  
  var expanseWidth int32
  var expanseHeight int32
  expanseWidth = screenWidth / 8
  expanseHeight = screenWidth / 15

  if !rl.IsSoundPlaying(*bg) {
    rl.PlaySound(*bg)
  }

  if rl.IsMouseButtonDown(rl.MouseButtonLeft) {

    if randomNumber == 0 && rl.GetMouseX() > randWidth - expanseWidth && rl.GetMouseX() < randWidth - expanseWidth + expanseWidth && rl.GetMouseY() > randHeight - expanseHeight && rl.GetMouseY() < randHeight - expanseHeight + expanseHeight {
      opened = false
      return
    } else if randomNumber == 1 && rl.GetMouseX() > randWidth2 - expanseWidth && rl.GetMouseX() < randWidth2 - expanseWidth + expanseWidth && rl.GetMouseY() > randHeight2 - expanseHeight && rl.GetMouseY() < randHeight2 - expanseHeight + expanseHeight {
      opened = false
      return
    } else {
      if frame != 0 {
        rl.PlaySound(*scream)
      }
    }

  }

  rl.BeginDrawing()

    rl.ClearBackground(rl.White)
    rl.DrawTexture(*texture, 0, 0, rl.White)

    rl.DrawRectangle(randWidth - expanseWidth, randHeight - expanseHeight, expanseWidth, expanseHeight, rl.Lime)
    rl.DrawText("Yes", (randWidth - expanseWidth) + (expanseWidth / 2) - (rl.MeasureText("Yes", 40) / 2), (randHeight - expanseHeight) + (expanseHeight / 2), 40, rl.Black)
    rl.DrawRectangle(randWidth2 - expanseWidth, randHeight2 - expanseHeight, expanseWidth, expanseHeight, rl.Red)
    rl.DrawText("No", (randWidth2 - expanseWidth) + (expanseWidth / 2) - (rl.MeasureText("No", 40) / 2), (randHeight2 - expanseHeight) + (expanseHeight / 2), 40, rl.Black)

    rl.DrawRectangle(screenWidth / 2 - (expanseWidth * 2), expanseHeight, expanseWidth * 4, expanseHeight * 2, rl.LightGray)
    rl.DrawText("Are you still playing?", screenWidth / 2 - (rl.MeasureText("Are you still playing?", 60) / 2), expanseHeight * 2, 60, rl.Black)

  rl.EndDrawing()

  if (rl.WindowShouldClose()) {
    opened = false
  }

  frame = frame + 1

}

func main() {

  rand.Seed(time.Now().UnixNano())

	rl.InitWindow(800, 800, "wise tree 2")

  screenWidth := int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
  screenHeight := int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))

  expanseWidth := int32(screenWidth / 8)
  expanseHeight := int32(screenWidth / 15)

  // Textures

  image := rl.LoadImage("cokey.png")
  rl.ImageResize(image, screenWidth, screenHeight)

  texture := rl.LoadTextureFromImage(image)

  rl.UnloadImage(image)

  // Audio

  rl.InitAudioDevice()

  bg := rl.LoadSound("roddy_rich_new_beat.wav")
  scream := rl.LoadSound("mosquito.wav")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

    var randWidth int32
    var randHeight int32
    var randWidth2 int32
    var randHeight2 int32

    randWidth = expanseWidth + int32(rand.Intn(int((screenWidth / 2) - expanseWidth)))
    randHeight = (screenHeight / 2) + int32(rand.Intn(int(screenHeight / 2 - expanseHeight)))
    randWidth2 = (screenWidth / 2) + int32(rand.Intn(int(screenWidth / 2 - expanseWidth)))
    randHeight2 = (screenHeight / 2) + int32(rand.Intn(int(screenHeight / 2 - expanseHeight)))

    var rnd int
    rnd = rand.Intn(2)

    var timer int
    timer = timer + 30
    timer = timer + rand.Intn(90)

    rl.ToggleBorderlessWindowed()
    for opened {
      draw(rnd, randWidth, randHeight, randWidth2, randHeight2, &texture, &bg, &scream)
    }
    rl.ToggleBorderlessWindowed()

    rl.StopSound(bg)
    rl.StopSound(scream)

    if !rl.IsWindowHidden() {
      rl.ClearWindowState(rl.FlagWindowTopmost)
      rl.SetWindowState(rl.FlagWindowHidden)
    }

    if !rl.WindowShouldClose() {
      time.Sleep(time.Duration(timer) * time.Second)
    }

    rl.ClearWindowState(rl.FlagWindowHidden)
    rl.SetWindowState(rl.FlagWindowTopmost)
    opened = true
    frame = 0

	}

  rl.UnloadTexture(texture)
  rl.UnloadSound(bg)
  rl.UnloadSound(scream)

  rl.CloseAudioDevice()

	rl.CloseWindow()

}
