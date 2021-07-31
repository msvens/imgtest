package main

import (
	"github.com/disintegration/imaging"
	"image"
	"os"
	"time"
)

type ImagingGen struct {
	imgDir string
}

func NewImagingGen(dir string) (ImgTester, error) {

	if err := os.MkdirAll(dir, 0744); err != nil {
		return nil, err
	}
	println("BImgGen created with dir: ", dir)
	return &ImagingGen{dir}, nil
}

func (gen *ImagingGen) Test(inputFiles []string, tests []TestSetup) ([]OveralResult, error) {

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

func (gen *ImagingGen) testImg(inputFile string, tests []TestSetup) (OveralResult,error) {

	var res = OveralResult{InputFile: inputFile, tests: make([]TestResult, len(tests), len(tests))}

	start := time.Now()

	src, err := imaging.Open(inputFile)

	res.InputReadTime = time.Now().Sub(start)

	if err != nil {
		return res, err
	}

	for i, test := range tests {
		res.tests[i] = TestResult{}
		res.tests[i].Test = test
		start1 := time.Now()
		var dstImage image.Image

		if test.Resize {
			dstImage = imaging.Resize(src, test.Width, 0, imaging.Lanczos)

		} else {
			dstImage = imaging.Fill(src, test.Width, test.Height, imaging.Center, imaging.Lanczos)
		}
		res.tests[i].TransformTime = time.Now().Sub(start1)
		start1 = time.Now()
		imaging.Save(dstImage, FileName(inputFile, gen.imgDir, test), imaging.JPEGQuality(test.Quality))
		res.tests[i].SaveTime = time.Now().Sub(start1)
	}

	res.TotalTime = time.Now().Sub(start)
	return res,nil

}




