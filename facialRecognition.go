package main
import (
	"gocv.io/x/gocv"
	"image/color"
	"log"
	"sync"
)
func main() {
	var wg sync.WaitGroup
	route := "/home/mb/Documentos/go_projects/facialRecognition/haarcascades/"
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("facial recognition")
	img := gocv.NewMat()
	files := []string{
		"haarcascade_frontalface_default.xml",
		"haarcascade_eye.xml",
	}
	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the device")
			continue
		}
		for index := range files {
			wg.Add(1)
			detected(&wg, route + files[index], img)
		}
		wg.Wait()

		window.IMShow(img)
		gocv.WaitKey(1)
	}
}
func detected (wg *sync.WaitGroup, route string, img gocv.Mat){
	color := color.RGBA{0, 0, 255, 0}
	file := route
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(file)
	detect := classifier.DetectMultiScale(img)
	for _, r := range detect {
		//fmt.Println("Eyes", r)
		gocv.Rectangle(&img, r, color, 3)
	}
	wg.Done()
}
