package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TestSetup struct {
	Width int
	Height int
	Resize bool
	Quality int
}

type TestResult struct {
	Test TestSetup
	TransformTime time.Duration
	SaveTime time.Duration
}

func (tr TestResult) String() string {
	return fmt.Sprintf("TestSetup: %v, Transform: %v, Save: %v", tr.Test, tr.TransformTime.Milliseconds(), tr.SaveTime.Milliseconds())
}

type OveralResult struct {
	InputFile string
	InputReadTime time.Duration
	TotalTime time.Duration
	tests []TestResult
}

func (or OveralResult) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Input: %v, InputRead: %v, Total: %v\n", or.InputFile, or.InputReadTime.Milliseconds(), or.TotalTime.Milliseconds())
	for _, t := range or.tests {
		fmt.Fprintf(&b, "\t%v\n", t)
	}
	return b.String()
}

type ImgTester interface {
	Test(inputFiles []string, tests []TestSetup) ([]OveralResult, error)
}

func (ts TestSetup) String() string {
	return fmt.Sprintf("%v-%v-%v-%v",ts.Width, ts.Height, ts.Resize, ts.Quality)
}

func FileName(srcFile, dirName string, test TestSetup) string {
	base := filepath.Base(srcFile)
	dir := dirName
	if dir == "" {
		dir = filepath.Dir(srcFile)
	}
	ext := filepath.Ext(base)
	trimmed := strings.TrimSuffix(base, ext)
	dstFile := fmt.Sprintf("%s-%s%s",trimmed,test,ext)
	return filepath.Join(dir,dstFile)
}

func createTests() []TestSetup {
	tests := make([]TestSetup, 5, 5)
	tests[0] = TestSetup{Width: 400, Height: 400, Resize: false, Quality: 90}
	tests[1] = TestSetup{Width: 1200, Height: 628, Resize: false, Quality: 90}
	tests[2] = TestSetup{Width: 1200, Height: 1200, Resize: false, Quality: 90}
	tests[3] = TestSetup{Width: 1080, Height: 1350, Resize: false, Quality: 90}
	tests[4] = TestSetup{Width: 1200, Height: -1, Resize: true, Quality: 90}
	return tests
}

func main() {

	tests := createTests()

	files, _ := os.ReadDir("./testimages")
	path, _ := filepath.Abs("./testimages")
	imgFiles := make([]string,len(files), len(files))
	for i, entry := range files {
		imgFiles[i] = filepath.Join(path,entry.Name())
	}
	homeDir, _ := os.UserHomeDir()
	println(homeDir)
	bGen, err := NewBImgGen(filepath.Join(homeDir, "testimgs", "bimg"))
	iGen, err := NewImagingGen(filepath.Join(homeDir, "testimgs", "imaging"))

	if err != nil {
		println(err)
		return
	}
	results, err := bGen.Test(imgFiles, tests);
	println("Results of BImg")
	if err != nil {
		println(err)
	} else {
		for _,res := range results {
			fmt.Println(res)
		}
	}
	results, err = iGen.Test(imgFiles, tests)
	println("Results of IImg")
	if err != nil {
		println(err)
	} else {
		for _,res := range results {
			fmt.Println(res)
		}
	}


}
