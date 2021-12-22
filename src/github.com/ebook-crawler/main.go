package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {

	var pageNumberCount int
	fmt.Println("Enter the number of pages for this books.")
	fmt.Scan(&pageNumberCount)

	//fmt.Printf("",pageNumberCount,"f")
	//fmt.Printf("%d\n",pageNumberCount)

	fmt.Printf("Processing now %d pages", pageNumberCount)
	// Add x pages to be sure we don't miss at the end of the book.
	pageNumberCount = pageNumberCount + 5
	initMousePosition()
	robotgo.Sleep(1)
	activateWindowTab()
	robotgo.Sleep(1)
	goToScreenshotButton()
	robotgo.Sleep(1)

	for i := 1; i <= pageNumberCount; i++ {
		takeScreenshot()
		robotgo.MilliSleep(500)
		nextPage()
		robotgo.Sleep(2)
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
	robotgo.MilliSleep(250)
	robotgo.KeyUp("right")
}

func takeScreenshot() {
	robotgo.Click()
}
