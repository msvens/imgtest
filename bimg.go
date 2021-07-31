package main

import (
	"github.com/h2non/bimg"
	"os"
	"time"
)

func CreateImageDir(dir string) error {
	var err error
	if err = os.MkdirAll(dir, 0744); err != nil {
		return err
	}
	return nil
}

type BImgGen struct {
	imgDir string
}

func NewBImgGen(dir string) (ImgTester, error) {

	if err := os.MkdirAll(dir, 0744); err != nil {
		return &BImgGen{}, err
	}
	println("BImgGen created with dir: ", dir)
	return &BImgGen{dir}, nil
}

func (gen *BImgGen) Test(inputFiles []string, tests []TestSetup) ([]OveralResult, error){

	res := make([]OveralResult, len(inputFiles), len(inputFiles))

	for i,fname := range inputFiles {
		tr, err := gen.testImg(fname, tests)
		res[i] = tr
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (gen *BImgGen) testImg(inputFile string, tests []TestSetup) (OveralResult,error) {

	var res = OveralResult{InputFile: inputFile, tests: make([]TestResult, len(tests), len(tests))}

	start := time.Now()

	buffer, err := bimg.Read(inputFile)

	res.InputReadTime = time.Now().Sub(start)

	if err != nil {
		return res, err
	}

	var options bimg.Options


	img := bimg.NewImage(buffer)

	for i, test := range tests {
		start1 := time.Now()
		res.tests[i] = TestResult{}
		res.tests[i].Test = test
		if test.Resize {
			options = bimg.Options{Quality: test.Quality, Width: test.Width}
		} else {
			options = bimg.Options{Quality: test.Quality, Crop: true, Gravity: bimg.GravityCentre,
				Width: test.Width, Height: test.Height}
		}
		buff, err := bimg.Resize(img.Image(), options)
		res.tests[i].TransformTime = time.Now().Sub(start1)
		if err != nil {
			return res, err
		}
		start1 = time.Now()
		bimg.Write(FileName(inputFile, gen.imgDir, test), buff)
		res.tests[i].SaveTime = time.Now().Sub(start1)
	}

	res.TotalTime = time.Now().Sub(start)

	return res, nil



	

}




