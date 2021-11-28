/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 10:24 下午
 **/
package tools

import (
	"errors"
	"fmt"
	"gfutil/config"
	"github.com/xuzhipeng12/gogfapi/gfapi"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetFriendlyFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

// 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

//检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return !os.IsNotExist(err)
}

// 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

//如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

//新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func IsDir(src string) (bool, error) {
	if !CheckExist(src) {
		return false, os.ErrNotExist
	}
	filneInfo, _ := os.Stat(src)
	return filneInfo.IsDir(), nil
}

func GfIsDir(volume *gfapi.Volume, src string) (bool, error) {
	if GfCheckExist(volume, src) {
		fileInfo, _ := volume.Stat(src)
		return fileInfo.IsDir(), nil
	}
	return false, os.ErrNotExist

}
func GfCheckExist(volume *gfapi.Volume, src string) bool {
	_, err := volume.Stat(src)
	return !os.IsNotExist(err)
}

func CopyFile(srcFile interface{}, destFile interface{}) error {
	buf := make([]byte, config.CopyFileBufferSize)
	for {
		var n = 0
		var err = errors.New("")
		if sof, ok := srcFile.(*os.File); ok {
			n, err = sof.Read(buf)
		} else {
			sgf := srcFile.(*gfapi.File)
			n, err = sgf.Read(buf)
		}
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if sof, ok := destFile.(*os.File); ok {
			if _, err := sof.Write(buf[:n]); err != nil {
				return err
			}
		} else {
			sgf := destFile.(*gfapi.File)
			if _, err := sgf.Write(buf[:n]); err != nil {
				return err
			}
		}
	}
	return nil
}
