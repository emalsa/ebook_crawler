package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
)

func main() {
	var directoryName string
	args := os.Args

	// Argument or with cli
	if len(args) > 1 {
		directoryName = args[1]
	} else if directoryName == "" {
		fmt.Println("Enter the directory name:")
		_, err := fmt.Scan(&directoryName)
		if err != nil {
			fmt.Println("Problem with directory name input:")
			return
		}
	}

	// Create zip archive
	createZip(directoryName)

	// Create dir
	createDir(directoryName)

	// Rename files by order
	renameFiles(directoryName)

	// White background
	addWhiteBackground(directoryName)

	// Add white rectangle on top & bottom
	addTopAndBottomBar(directoryName)

	// Optimize
	compressImages(directoryName)

	// Create PDF
	createPDF(directoryName)
	fmt.Println("Finished!")
}

func createZip(directoryName string) {
	fmt.Println("Add zip archive.")
	cmd := exec.Command("/bin/bash", "-c", "-path", "rm -r /Users/daniele/misc/"+directoryName+"/converted/ ; cd /Users/daniele/misc/"+directoryName+" && zip -r ../"+directoryName+".zip *")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(err.Error()))
		fmt.Printf("combined out:\n%s\n", string(stdout))
		return
	}
	fmt.Println("Finished creating archive.")
	fmt.Println("--------------------")
}
func createDir(directoryName string) {
	fmt.Println("Create dir */converted.")
	path := "/Users/daniele/misc/" + directoryName + "/converted"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModeDir|0755)
	}
	fmt.Println("Finished creating dir */converted.")
	fmt.Println("--------------------")

}
func renameFiles(directoryName string) {
	fmt.Println("Rename files by order.")

	dir := "/Users/daniele/misc/" + directoryName + "/"
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
	var prefixNumber string
	for _, fi := range fis {
		fmt.Println(fi.Name())
		prefixNumber = fmt.Sprintf("%04d", i)
		newFilename = dir + prefixNumber + ".png"
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
func addWhiteBackground(directoryName string) {
	fmt.Println("Add white background.")
	cmd := exec.Command("/usr/local/bin/mogrify", "-path", "/Users/daniele/misc/"+directoryName+"/converted", "-format", "png", "-fill", "#FFFFFF", "-opaque", "#E1DFDA", "-fuzz", "1%", "/Users/daniele/misc/"+directoryName+"/*.png")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(err.Error()))
		fmt.Printf("combined out:\n%s\n", string(stdout))
		return
	}
	fmt.Println("Finished adding white background.")
	fmt.Println("--------------------")
}
func addTopAndBottomBar(directoryName string) {
	fmt.Println("Add white bar on top and bottom.")
	cmd1 := exec.Command("/usr/local/bin/mogrify", "-path", "/Users/daniele/misc/"+directoryName+"/converted", "-format", "png", "-fill", "white", "-draw", "rectangle 0,0,1440,70", "-fill", "white", "-draw", "rectangle 0,1820,1440,1920", "/Users/daniele/misc/"+directoryName+"/converted/*.png")
	stdout1, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		fmt.Println(err1.Error())
		fmt.Printf("combined out:\n%s\n", string(stdout1))
		return
	}
	fmt.Println("Finished adding white bar on top and bottom.")
	fmt.Println("--------------------")
}
func compressImages(directoryName string) {
	fmt.Println("Start compressing images.")
	cmd := exec.Command("/bin/bash", "-c", "/usr/local/bin/pngquant --quality 15-20 /Users/daniele/misc/"+directoryName+"/converted/*png -f --ext .png")
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("combined out:\n%s\n", string(stdout))
		return
	}
	fmt.Println("Finished compressing.")
	fmt.Println("--------------------")
}
func createPDF(directoryName string) {
	fmt.Println("Start creating PDF.")
	cmd := exec.Command("/bin/bash", "-c", "convert -density 450 $(ls /Users/daniele/misc/"+directoryName+"/converted/*.png | sort -n) /Users/daniele/misc/"+directoryName+"/converted/"+directoryName+".pdf")
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
