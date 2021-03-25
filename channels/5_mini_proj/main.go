package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/trace"
	"sync"
	"time"
)

type Photos []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Image struct {
	filePath string
	img      []byte
}

const MAX_DOWNLOAD = 2000

func main() {
	defer func() {
		fmt.Println("Main program exit successfully")
	}()

	log.SetFlags(log.Ltime)
	photos := Photos{}
	err := getJson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	fmt.Println(len(photos[0:MAX_DOWNLOAD]))

	dir := "myDownloadImage" + time.Now().Format("15_04_05")
	if _, err = os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModeDir)
	}

	// ตรวจสอบจำนวนของจำนวนการสร้าง trace ของ go rountine
	f, err := os.Create(dir + ".trace.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	

	chImg := make(chan Image, len(photos[0:MAX_DOWNLOAD]))
	counter := sync.WaitGroup{}
	token := make(chan struct{}, 2000) // limit download
	for _, v := range photos[0:MAX_DOWNLOAD] {
		v := v
		counter.Add(1)
		go func() {
			defer counter.Done()

			//  take a token
			token <- struct{}{}
			img, err := downloadImage(v.ThumbnailURL) // use token to work
			<-token

			if err != nil {
				// log.Fatal(err)
				log.Println(err)
				return
			}

			format, err := decodeImage(img)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Download : %v\n", v.ID)
			absoluteFileName := filepath.Join(dir, fmt.Sprintf("%d.%s", v.ID, format)) // (directory) / (filename)

			chImg <- Image{filePath: absoluteFileName, img: img}

		}()

	}
	go func() {
		counter.Wait()
		close(chImg)
	}()

	for v := range chImg {
		err := saveImage(v.filePath, v.img)
		if err != nil {
			log.Println(err)
		}

	}

}

func saveImage(fileName string, img []byte) error {

	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("saveImage - cannot create file : %v ", err)
	}
	defer f.Close()
	// f <-- img
	_, err = io.Copy(f, bytes.NewReader(img)) // สร้างไฟล์ไว้ในคอมเรา
	if err != nil {
		return fmt.Errorf("saveImage - cannot create file : %v ", err)
	}
	log.Printf("\tSave : %v\n", fileName)
	return nil
}

func decodeImage(img []byte) (string, error) {
	// analyze file type
	_, format, err := image.Decode(bytes.NewReader(img))
	return format, err
}

func downloadImage(url string) ([]byte, error) {
	errMsg := func(err error) error {
		return fmt.Errorf("downloadImage : %v", err)
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, errMsg(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errMsg(err)
	}
	return body, nil

}

func getJson(url string, structType interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	switch v := structType.(type) {
	case *Photos:
		decoder := json.NewDecoder(res.Body)
		photos := structType.(*Photos)
		decoder.Decode(&photos)
		return nil

	default:
		return fmt.Errorf("getJson : not support type %v", v)
	}
}
