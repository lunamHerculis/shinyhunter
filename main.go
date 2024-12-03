package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	factor = 1

	displayID = 0

	keyA         = "g"
	keyB         = "f"
	keyUp        = "w"
	keyDown      = "s"
	keyLeft      = "a"
	keyRight     = "s"
	keyStart     = "t"
	keySelect    = "r"
	keyLShoulder = "q"
	keyRShoulder = "e"
)

var (
	nonShinyHex = ""
	locationX   = 0
	locationY   = 0
	counter     = 0
)

func main() {
	fmt.Println("starting in 5...")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 4...")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 3...")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 2...")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 1...")
	time.Sleep(1 * time.Second)
	fmt.Println("catch em all!")

	locationX, locationY = robotgo.Location()
	fmt.Printf("locationX: %d, locationY: %d\n", locationX, locationY)
	nonShinyHex = getMousePixelColor()
	fmt.Printf("nonShinyHex: %s\n", nonShinyHex)
	time.Sleep(1 * time.Second)

startingPos:
	counter++

	softReset()

	fullCycle()

	time.Sleep(3 * time.Second / factor)

	color := getMousePixelColor()
	if color == nonShinyHex {
		fmt.Printf("no shiny, target hex: %s, color: %s, counter: %d\n", nonShinyHex, color, counter)
		goto startingPos
	} else {
		fmt.Printf("holy shit a shiny!!!, target hex: %s, color: %s, counter: %d\n", nonShinyHex, color, counter)
		return
	}
}

func getMousePixelColor() string {
	c := robotgo.GetPixelColor(locationX, locationY, displayID)
	return c
}

func softReset() {
	robotgo.KeyDown(keyStart)
	robotgo.KeyDown(keySelect)
	robotgo.KeyDown(keyA)
	robotgo.KeyDown(keyB)
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp(keyStart)
	robotgo.KeyUp(keySelect)
	robotgo.KeyUp(keyA)
	robotgo.KeyUp(keyB)
	time.Sleep(10 * time.Millisecond)
}

func fullCycle() {
	mashButton(keyA, 15*time.Second/factor)
	time.Sleep(100 * time.Millisecond / factor)
	mashButton(keyB, 9*time.Second/factor)
	time.Sleep(500 * time.Millisecond / factor)
	keyStroke(keyStart)
	time.Sleep(1000 * time.Millisecond / factor)
	keyStroke(keyA)
	time.Sleep(1000 * time.Millisecond / factor)
	keyStroke(keyA)
	time.Sleep(700 * time.Millisecond / factor)
	keyStroke(keyA)
}

func mashButton(key string, duration time.Duration) {
	startTime := time.Now()
	for {
		if time.Now().After(startTime.Add(duration)) {
			break
		}
		time.Sleep(100 * time.Millisecond)
		keyStroke(key)
	}
}

func keyStroke(key string) {
	robotgo.KeyDown(key)
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp(key)
	time.Sleep(10 * time.Millisecond)
}
