package request

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/spf13/viper"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"
)

type ImageValidator struct {
	Name      string `json:"name"`
	ImageData string `json:"image_data"`
}

func (iv *ImageValidator) Validate() error {
	encodedImage := iv.ImageData[strings.Index(iv.ImageData, ",")+1:]
	rawImage, _ := base64.StdEncoding.DecodeString(encodedImage)
	switch strings.Trim(iv.ImageData[5:strings.Index(iv.ImageData, ",")], ";base64") {
	case "image/png":
		r := bytes.NewReader(rawImage)
		if pngImage, err := png.Decode(r); err != nil {
			return errors.New("Invalid image data. Please try again.")
		} else {
			iv.ImageData = strconv.Itoa(int(time.Now().Unix())) + "_" + iv.Name + ".png"
			if file, err := os.OpenFile(viper.GetString("app.public_path")+"/images/"+iv.Name, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
				return err
			} else {
				png.Encode(file, pngImage)
			}
		}
	case "image/jpeg":
		if jpegImage, err := jpeg.Decode(bytes.NewReader(rawImage)); err != nil {
			return errors.New("Invalid image data. Please try again.")
		} else {
			strconv.Itoa(int(time.Now().Unix()))
			iv.ImageData = strconv.Itoa(int(time.Now().Unix())) + "_" + iv.Name + ".jpg"
			if file, err := os.OpenFile(viper.GetString("app.public_path")+"/images/"+iv.ImageData, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
				return err
			} else {
				jpeg.Encode(file, jpegImage, &jpeg.Options{Quality: 75})
			}
		}
	case "image/gif":
		if gifImage, err := gif.Decode(bytes.NewReader(rawImage)); err != nil {
			return errors.New("Invalid image data. Please try again.")
		} else {
			iv.ImageData = strconv.Itoa(int(time.Now().Unix())) + "_" + iv.Name + ".gif"
			if file, err := os.OpenFile(viper.GetString("app.public_path")+"/images/"+iv.ImageData, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
				return err
			} else {
				gif.Encode(file, gifImage, &gif.Options{})
			}
		}
	default:
		return errors.New("Unsupported image format.")
	}

	return nil
}
