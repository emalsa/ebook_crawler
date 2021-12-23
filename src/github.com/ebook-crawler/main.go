package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	var pageNumberCount int
	fmt.Println("Enter the number of pages for this books.")
	fmt.Scan(&pageNumberCount)
	fmt.Printf("Processing now %d pages", pageNumberCount)
	//Add x pages to be sure we don't miss something at end of the book.

	initMousePosition()
	robotgo.MilliSleep(500)
	activateWindowTab()
	robotgo.MilliSleep(500)
	goToScreenshotButton()
	robotgo.MilliSleep(500)

	for i := 0; i <= pageNumberCount; i++ {
		takeScreenshot()
		robotgo.Sleep(1)
		nextPage()
		robotgo.Sleep(1)
	}
	fmt.Println("End")
}

func goToScreenshotButton() {
	robotgo.MoveSmooth(630, 315)
}
func initMousePosition() {
	robotgo.MoveSmooth(0, 0)
}
func activateWindowTab() {
	robotgo.MoveSmooth(610, 110)
	robotgo.Click()
}
func nextPage() {
	robotgo.KeyDown("right")
	robotgo.MilliSleep(350)
	robotgo.KeyUp("right")
}
func takeScreenshot() {
	robotgo.Click()
}
