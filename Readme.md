# About

**IN PROGRESS**

imgtest is a simple test/benchmark for golang image processing frameworks. Frameworks tested

- [bimg](https://github.com/h2non/bimg) - uses libvips as the underlying processing engine
- [imaging](https://github.com/disintegration/imaging) - native go processing engine

## Tests
The purpose of the tests is to primarily test resize and crop functionality. The typical scenario
is when you have a src image and you need to generate a set of scaled/cropped images 
(e.g. thumbnails, landscape, portraits) from that image. The suite currently contains 5 tests

1. 400x400 resize, keep aspect ratio (crop)
2. 1200x628 resize, keep aspect ratio (crop)
3. 1200x1200 resize, keep aspect ratio (crop)
4. 1080x1350 resize, keep aspect ratio (crop)
5. 1200x0 resize on width

Test 1-4 uses a combination of resize and crop with a center gravity strategy
All tests uses a cubic tranformation filter

All resulting images are saved as jpeg quality 90

## Results

Both libraries does a good job with the simple processing in these tests and the processing times
are fairly simlar. For bimg read/writes are significantly faster as they dont involve any encode/decode
of the src/dest image (that is done in the actual resize function). Thus it is not meaningful to compare
the transform steps and IO steps (only the total)

### Detailed Results

## Conclusions

On average bimg performs slightly faster so if that is the only metric you are after **bimg** is the library
to go with. On the other hand, if you are looking for something with zero compile hassle and no external
dependencies **imaging** is a good choice. That said, these are my subjective pros and cons with each library

### Bimg
Built ontop of libvips which is an extremly well tested and maintained image library. Hence bimg can deal with
virtually any image format, has numerous transformation capabilities and maintains image meta information (e.g. EXIF). bimg
also offers an intutive and simple API
The one downside with bimg is that it uses libvips and as a developer/system owner you have to manually ensure that libvips
is installed on the target platform. Thus if you choose to go with bimg be prepared for some extra work and dont expect things
to just "work"

### Imaging
Imaging is a totally golang native library that uses gos inbuilt image APIs for processing. If you are primarily looking for 
various resize/crop/blur/etc operations with a simple to use API Imaging is a good alternative. You have a variety of filters to choose
from and if offers good performance on those. The big downside of Imagining is that it is limited to the image formats that
go handles natively as well as it doesnt handle image meta information (e.g. exif information). libvips is also image framework that
will be around while Imaging is a smaller project with a more unclear future. That said, if you are looking for solid a solid
resize like library that will work out of the box on any platform Imaging is a good choice


