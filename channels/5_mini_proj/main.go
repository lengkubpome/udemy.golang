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
)

type Photos []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func main() {

	photos := Photos{}
	err := getJson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	fmt.Println(len(photos))

	for _, v := range photos {

		img, err := downloadImage(v.ThumbnailURL)
		if err != nil {
			log.Fatal(err)
		}

		format, err := decodeImage(img)
		if err != nil {
			log.Fatal(err)
		}
		fileName := fmt.Sprintf("%d.%s", v.ID, format) // (directory) / (filename)
		err = saveImage(fileName, img)
		if err != nil {
			log.Println(err)
		}
	}

}

func saveImage(fileName string, img []byte) error {

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	// f <-- img
	_, err = io.Copy(f, bytes.NewReader(img)) // สร้างไฟล์ไว้ในคอมเรา
	if err != nil {
		return err
	}
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
