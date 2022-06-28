package util

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/YuleZhang/douyin/pkg/constants"
	"github.com/tencentyun/cos-go-sdk-v5"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 将[uid, filename, timestamp]合并后进行md5加密
func MD5FileNameEncoder(uid int64, filename string) (string, error) {
	md := md5.New()
	encodeStr := fmt.Sprintf("%d", uid) + filename + fmt.Sprintf("%d", time.Now().Unix())
	io.WriteString(md, encodeStr)
	return hex.EncodeToString(md.Sum(nil)), nil
}

// 读取视频文件，并抽取关键帧字节流返回
func ReadFrameAsJpeg(filePath string) (*bytes.Buffer, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	return reader, err
}

// 文件上传腾讯云oss，返回完整路径
func FileUploadToOss(fileName string, reader io.Reader) (string, error) {
	u, _ := url.Parse(constants.OssAdress)
	b := &cos.BaseURL{BucketURL: u}
	co := cos.NewClient(b, &http.Client{
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(constants.OssSecretID),
			SecretKey: os.Getenv(constants.OssSecretKey),
		},
	})

	resp, err := co.Object.Put(context.Background(), fileName, reader, nil)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	return constants.OssAdress + co.Object.GetObjectURL(fileName).Path, nil
}
