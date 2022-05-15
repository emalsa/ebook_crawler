package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"strconv"
)

func main() {
	var pageNumberCount int
	args := os.Args

	if len(args) > 1 {
		pageNumberCount, _ = strconv.Atoi(args[1])
	} else {
		fmt.Println("Enter the number of pages for this books.")
		_, err := fmt.Scan(&pageNumberCount)
		if err != nil {
			fmt.Println("Problem with directory name input:")
			return
		}
	}

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
func activateWindowTab() {
	robotgo.MoveSmooth(500, 30)
	robotgo.Click()
}
func goToScreenshotButton() {
	robotgo.MoveSmooth(620, 320)
}
func initMousePosition() {
	robotgo.MoveSmooth(0, 0)
}
func nextPage() {
	robotgo.KeyDown("right")
	robotgo.MilliSleep(350)
	robotgo.KeyUp("right")
}
func takeScreenshot() {
	robotgo.Click()
}
