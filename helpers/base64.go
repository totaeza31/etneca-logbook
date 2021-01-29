package helpers

import (
	"bytes"
	"encoding/base64"
	"etneca-logbook/models"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func SavePicture(b64 string, id string) (string, error) {
	idx := strings.Index(b64, ";base64,")
	if idx < 0 {
		panic("InvalidImage")
	}
	ImageType := b64[11:idx]
	unbased, ers := base64.StdEncoding.DecodeString(b64[idx+8:])
	if ers != nil {
	}
	r := bytes.NewReader(unbased)
	switch ImageType {
	case "png":
		im, ers := png.Decode(r)
		if ers != nil {
			panic("Bad png")
		}
		pathFile := "public/image/" + id + ".png"

		f := <-openFile(pathFile)

		err := <-encodePng(f, im)
		if err != nil {
			return "", err
		} else {
			return id + ".png", nil
		}
	case "jpeg":
		im, ers := jpeg.Decode(r)
		if ers != nil {
			panic("Bad jpeg")
		}
		pathFile := "public/image/" + id + ".jpeg"

		f := <-openFile(pathFile)

		err := <-encodeJpeg(f, im)
		if err != nil {
			return "", err
		} else {
			return id + ".jpeg", nil
		}
	case "gif":
		im, ers := gif.Decode(r)
		if ers != nil {
			panic("Bad gif")
		}
		pathFile := "public/image/" + id + ".gif"

		f := <-openFile(pathFile)

		err := <-encodeGif(f, im)
		if err != nil {
			return "", err
		} else {
			return id + ".gif", nil
		}

	}
	return "", nil
}

func DeleteFile(fileName string) <-chan error {
	e := make(chan error)
	go func() {
		defer close(e)
		err := os.Remove("public/image/" + fileName)
		e <- err
	}()
	return e
}

func openFile(fileName string) <-chan *os.File {
	r := make(chan *os.File)
	go func() {
		defer close(r)
		fe, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
		r <- fe
	}()
	return r
}

func encodePng(f *os.File, im image.Image) <-chan error {
	r := make(chan error)

	go func() {
		defer close(r)
		err := png.Encode(f, im)

		r <- err
	}()
	return r
}

func encodeJpeg(f *os.File, im image.Image) <-chan error {
	r := make(chan error)
	go func() {
		defer close(r)
		err := png.Encode(f, im)

		r <- err
	}()
	return r
}

func encodeGif(f *os.File, im image.Image) <-chan error {
	r := make(chan error)

	go func() {
		defer close(r)
		err := gif.Encode(f, im, nil)

		r <- err
	}()
	return r
}

func UnmarshalData(bytes []byte, model models.Boats) <-chan models.Boats {
	models := make(chan models.Boats)
	go func() {
		defer close(models)
		bson.Unmarshal(bytes, &model)
		models <- model
	}()
	return models
}
