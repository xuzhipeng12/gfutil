/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 9:24 下午
 **/
package pkg

import (
	"fmt"
	"gfutil/config"
	"gfutil/tools"
	"github.com/xuzhipeng12/gogfapi/gfapi"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

type FileInfo struct {
	Volname  string
	Fullname string
	Name     string
	IsDir    bool
	IsVolume bool
	VolObj   *gfapi.Volume
	Perm     os.FileMode
	Size     int64
	ModTime  time.Time
}

func GetFileInfo(fullName string, hosts []string) *FileInfo {
	isVolume := strings.HasPrefix(fullName, config.GlusterFilePrefix)
	fi := FileInfo{}
	if isVolume {
		splitedPath := strings.Split(strings.Replace(fullName, config.GlusterFilePrefix, "", 1), "/")
		fi.Volname = splitedPath[0]
		fi.Fullname = strings.Replace(fullName, config.GlusterFilePrefix+fi.Volname, "", 1)
		fi.Name = path.Base(fi.Fullname)
		vol := &gfapi.Volume{}
		if err := vol.Init(fi.Volname, hosts...); err != nil {
			fmt.Fprintln(config.ErrOut, "Glusterfs volume init fail , try again")
			os.Exit(2)
		}
		if err := vol.Mount(); err != nil {
			fmt.Fprintln(config.ErrOut, "Glusterfs volume mount fail , try again")
			os.Exit(2)
		}
		fi.VolObj = vol
		fi.IsDir, _ = tools.GfIsDir(vol, fi.Fullname)
		fi.IsVolume = true
	} else {
		fi.Fullname = fullName
		fi.Name = path.Base(fi.Fullname)
		fi.IsDir, _ = tools.IsDir(fi.Fullname)
		fi.IsVolume = false

	}
	return &fi
}

func GetFiles(info *FileInfo) []*FileInfo {
	filesInfos := make([]*FileInfo, 0)
	if info.IsDir {
		// if is glusterfs file type
		if info.IsVolume {
			if d, ok := info.VolObj.Open(info.Fullname); ok == nil {
				files, _ := d.Readdir(0)
				for _, file := range files {
					if file.Name() == "." || file.Name() == ".." {
						continue
					}
					fi := FileInfo{
						Fullname: path.Join(info.Fullname, file.Name()),
						Volname:  info.Volname,
						Name:     file.Name(),
						IsDir:    file.IsDir(),
						IsVolume: info.IsVolume,
						VolObj:   info.VolObj,
						Perm:     file.Mode(),
						Size:     file.Size(),
						ModTime:  file.ModTime(),
					}
					filesInfos = append(filesInfos, &fi)
				}
			}
		} else { //  if local file
			if files, ok := ioutil.ReadDir(info.Fullname); ok == nil {
				for _, file := range files {
					if file.Name() == "." || file.Name() == ".." {
						continue
					}
					fi := FileInfo{
						Fullname: path.Join(info.Fullname, file.Name()),
						Volname:  info.Volname,
						Name:     file.Name(),
						IsDir:    file.IsDir(),
						IsVolume: false,
						VolObj:   info.VolObj,
						Perm:     file.Mode(),
						Size:     file.Size(),
						ModTime:  file.ModTime(),
					}
					//if ok, _ := tools.IsDir(fi.Fullname); ok {
					//
					//}
					filesInfos = append(filesInfos, &fi)
				}
			}
		}
	} else {
		filesInfos = append(filesInfos, info)
	}
	return filesInfos
}

func AppendFileInfo(slice []FileInfo, data ...FileInfo) []FileInfo {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]FileInfo, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
