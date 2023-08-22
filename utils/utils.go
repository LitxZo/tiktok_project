package utils

import (
	"bytes"
	"fmt"
	"os"

	"github.com/disintegration/imaging"
	"github.com/spf13/viper"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func AppendNewErr(initErr, newErr error) error {
	if initErr == nil {
		return newErr
	}
	return fmt.Errorf("%v, %w", initErr, newErr)
}

func GetCover(inFileName, coverName string, frameNum int) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		return "", err
	}
	err = imaging.Save(img, viper.GetString("Server.staticPath")+"/"+"image"+"/"+coverName)
	if err != nil {
		return "", err
	}
	return viper.GetString("Server.staticUrl") + "/" + "image" + "/" + coverName, nil
}
