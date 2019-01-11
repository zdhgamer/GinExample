package upload

import (
	"../setting"
	"path"
	"strings"
	"../util"
	"../file"
	"mime/multipart"
	"log"
	"../logging"
	"os"
	"fmt"
)

//获取在服务器上的路径
func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + setting.AppSetting.ImageSaveUrl+"/"+name
}

//重新生成新的md5文件名
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

//获取存贮路径
func GetImageSavePath() string {
	return setting.AppSetting.ImageSavePath
}

//获取存贮url
func GetImageSaveUrl() string {
	return setting.AppSetting.ImageSaveUrl
}

//获取文件全部存贮路径
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImageSavePath()
}

//检查文件是否是图片，通过后缀匹配
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

//判断文件大小是否符合标准大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Info(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

//检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
