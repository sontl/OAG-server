package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Images map[string]*Image
)

type Image struct {
	ImageId    string
	Score      int64
	PlayerName string
}

func init() {
	Images = make(map[string]*Image)
	Images["hjkhsbnmn123"] = &Image{"hjkhsbnmn123", 100, "astaxie"}
	Images["mjjkxsxsaa23"] = &Image{"mjjkxsxsaa23", 101, "someone"}
}

func AddImage(image Image) (ImageId string) {
	image.ImageId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Images[image.ImageId] = &image
	return image.ImageId
}

func GetImage(ImageId string) (image *Image, err error) {
	if v, ok := Images[ImageId]; ok {
		return v, nil
	}
	return nil, errors.New("ImageId Not Exist")
}

func GetAllImages() map[string]*Image {
	return Images
}

func UpdateImage(ImageId string, Score int64) (err error) {
	if v, ok := Images[ImageId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ImageId Not Exist")
}

func DeleteImage(ImageId string) {
	delete(Images, ImageId)
}
