package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	fastForwardSpeed = 1

	displayID = 0

	layoutUS = "2006-01-02 15:04:05"

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
	saveCounter = 0
	realCounter = 0
)

func main() {
	fmt.Println("starting in 7...")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 6...")
	time.Sleep(1 * time.Second)
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
	nonShinyHex = getPixelColorAtPresetLocation()
	fmt.Printf("nonShinyHex: %s\n", nonShinyHex)
	time.Sleep(1 * time.Second)

	for {

		softReset()
		toStarter()

		// after 100 runs, wait and save
		if saveCounter > 100 {

			newAdditionalWaitTime := 20 * saveCounter // roughly 1.2 frames * saveCounter
			fmt.Printf("new additional wait time in seconds: %.3f\n", float64(newAdditionalWaitTime)/1000)
			time.Sleep(time.Duration(newAdditionalWaitTime) * time.Millisecond)

			fmt.Println("saving...")
			save()

			saveCounter = 0
			continue
		}

		realCounter++
		saveCounter++

		newAdditionalWaitTime := 20 * saveCounter // roughly 1.2 frames * saveCounter
		fmt.Printf("new additional wait time in seconds: %.3f\n", float64(newAdditionalWaitTime)/1000)
		time.Sleep(time.Duration(newAdditionalWaitTime) * time.Millisecond)

		takeStarter()
		goIntoPokedex()

		color := getPixelColorAtPresetLocation()
		if color == nonShinyHex {
			fmt.Printf("no shiny, target hex: %s, color: %s, realCounter: %d, saveCounter: %d, time: %s\n", nonShinyHex, color, realCounter, saveCounter, time.Now().Format(layoutUS))
			continue
		}

		fmt.Printf("holy shit a shiny!!!, target hex: %s, color: %s, realCounter: %d,  saveCounter: %d\n time: %s\n", nonShinyHex, color, realCounter, saveCounter, time.Now().Format(layoutUS))
		return
	}
}

func getPixelColorAtPresetLocation() string {
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

func toStarter() {
	mashButton(keyA, 12*time.Second/fastForwardSpeed)
	mashButton(keyB, 3*time.Second/fastForwardSpeed)
}

func save() {
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyStart)
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyDown)
	time.Sleep(500 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyDown)
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyDown)
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyDown)
	time.Sleep(500 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyA)
	time.Sleep(2000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyA)
	time.Sleep(2000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyA)
	time.Sleep(9000 * time.Millisecond / fastForwardSpeed)
}

func takeStarter() {
	mashButton(keyA, 2*time.Second/fastForwardSpeed)
	mashButton(keyB, 2*time.Second/fastForwardSpeed)
}

func goIntoPokedex() {
	color := getPixelColorAtPresetLocation()
	fmt.Printf("color at pixel before pressing start: %s\n", color)
	keyStroke(keyStart)
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyDown)
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyA)
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyUp)
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyUp)
	time.Sleep(1000 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyA)
	time.Sleep(700 * time.Millisecond / fastForwardSpeed)
	keyStroke(keyA)
	time.Sleep(3 * time.Second / fastForwardSpeed)
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
