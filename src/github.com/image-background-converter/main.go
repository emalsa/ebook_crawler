package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
)

func main() {
	// Rename files by order
	renameFiles()
	// White background
	addWhiteBackground()
	// Add white rectangle on top & bottom
	addTopAndBottomBar()
	// Optimize
	compressImages()
	// Create PDF
	createPDF()
	fmt.Println("Finished!")
}

func renameFiles() {
	fmt.Println("Rename files by order.")
	dir := "/Users/daniele/misc/original/"
	f, _ := os.Open(dir)
	fis, _ := f.Readdir(-1)
	_ = f.Close()

	// Remove not used files (.DS_STORE)
	j := 0
	for _, n := range fis {
		if n.Name() != ".DS_Store" && !n.IsDir() {
			fis[j] = n
			j++
		}
	}
	fis = fis[:j]

	// Sort
	sort.Sort(ByModTime(fis))

	// Rename
	var i = 1
	var newFilename string
	var oldFilename string
	for _, fi := range fis {
		fmt.Println(fi.Name())
		newFilename = dir + strconv.Itoa(i) + ".png"
		oldFilename = dir + fi.Name()
		err := os.Rename(oldFilename, newFilename)
		if err != nil {
			fmt.Println(err)
			return
		}
		i++
	}

	fmt.Println("Finished rename files by order.")
	fmt.Println("--------------------")
}

func addWhiteBackground() {
	fmt.Println("Add white background.")
	cmd := exec.Command("/usr/local/bin/mogrify", "-path", "/Users/daniele/misc/original/converted", "-format", "png", "-fill", "#FFFFFF", "-opaque", "#E1DFDA", "-fuzz", "1%", "/Users/daniele/misc/original/*.png")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(err.Error()))
		fmt.Printf("combined out:\n%s\n", string(stdout))
		return
	}
	fmt.Println("Finished adding white background.")
	fmt.Println("--------------------")
}

func addTopAndBottomBar() {
	fmt.Println("Add white bar on top and bottom.")
	cmd1 := exec.Command("/usr/local/bin/mogrify", "-path", "/Users/daniele/misc/original/converted", "-format", "png", "-fill", "white", "-draw", "rectangle 0,0,1440,70", "-fill", "white", "-draw", "rectangle 0,1820,1440,1920", "/Users/daniele/misc/original/converted/*.png")
	stdout1, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		fmt.Println(err1.Error())
		fmt.Printf("combined out:\n%s\n", string(stdout1))
		return
	}
	fmt.Println("Finished adding white bar on top and bottom.")
	fmt.Println("--------------------")
}

func compressImages() {
	fmt.Println("Start compressing images.")
	cmd := exec.Command("/bin/bash", "-c", "/usr/local/bin/pngquant --quality 5-10 /Users/daniele/misc/original/converted/*png -f --ext .png")
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("combined out:\n%s\n", string(stdout))
		return
	}
	fmt.Println("Finished compressing.")
	fmt.Println("--------------------")
}

func createPDF() {
	fmt.Println("Start creating PDF.")
	cmd := exec.Command("/bin/bash", "-c", "convert -density 450 $(ls /Users/daniele/misc/original/converted/*.png | sort -n) /Users/daniele/misc/original/converted/output_Book.pdf")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("combined out:\n%s\n", string(stdout))
		return
	}

	fmt.Println("Finished creating PDF.")
	fmt.Println("--------------------")
}

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().Before(fis[j].ModTime())
}
