package libs

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	config "github.com/spf13/viper"
	"io"
	"iris/commons"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

func UploadFile(key string, Ctx iris.Context) (bool, string) {
	file, info, err := Ctx.FormFile(key)
	log.Println("file info err is", file, info, err)
	if file == nil {
		return true, ""
	}
	filePath := ""
	if err != nil {
		return false, "Error while uploading: <b>" + err.Error() + "</b>"
	}
	// log.Println(config.GetInt64("UploadSize"))
	// log.Println(config.GetInt("site.AdminId"))
	var minSize int64 = 0
	if info.Size > minSize {
		// log.Println(config.GetInt64("UploadSize"))
		if info.Size > config.GetInt64("site.UploadSize")*1024*1024 { // author 少加了site viper 看配置
			return false, "Error while uploading: UploadSize ToMax"
		}
		fname := strconv.Itoa(commons.GenerateRangeNum(100, 9999)) + "_" + info.Filename

		fileSuffix := path.Ext(fname)

		fileSuffixExists := false
		//CanFileSuffix := [...]string{".jpg", ".png", ".jpge", ".gif"}
		CanFileSuffix := strings.Split(config.GetString("site.UploadSuffixExists"), ",")
		for _, v := range CanFileSuffix {
			if v == fileSuffix {
				fileSuffixExists = true
			}
		}

		if fileSuffixExists == false {
			return false, "fileSuffix error: <b>" + fileSuffix + "</b>"
		}

		filePath = "./uploads/headico/" + fname
		out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return false, "Error while uploading: <b>" + err.Error() + "</b>"
		}
		defer out.Close()
		io.Copy(out, file)
		filePath = filePath[1:]
	}
	defer file.Close()
	return true, filePath
}

func UploadToQiniu(localFile string) (string, error) {
	config.SetConfigName("qiniu")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	accessKey := config.GetString("default.accessKey")
	secretKey := config.GetString("default.secretKey")
	bucket := config.GetString("default.bucket")

	tokens := strings.Split(localFile, "attachments/")
	key := tokens[len(tokens)-1]

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	//putExtra.NoCrc32Check = true
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return ret.Key, nil

}
