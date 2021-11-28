/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 9:24 下午
 **/
package pkg

import (
	"errors"
	"fmt"
	"gfutil/config"
	"gfutil/tools"
	"github.com/xuzhipeng12/gogfapi/gfapi"
	"os"
	"path"
	"strings"
)

type CopyOptions struct {
	Recursive    bool
	Force        bool
	View         bool
	Hosts        []string
	SrcfileInfo  *FileInfo
	DestfileInfo *FileInfo
}

func (o *CopyOptions) CP(args ...string) {
	o.SrcfileInfo = GetFileInfo(args[0], o.Hosts)
	o.DestfileInfo = GetFileInfo(args[1], o.Hosts)
	defer Umound(o.SrcfileInfo)
	defer Umound(o.DestfileInfo)
	if !o.Recursive && o.SrcfileInfo.IsDir {
		fmt.Fprintln(config.Out, o.SrcfileInfo.Fullname, "is directory, the recursive (-r) option makes copy all directories recursively.")
		os.Exit(2)
	} else {
		o.cp(o.SrcfileInfo, o.DestfileInfo)
	}

}
func (o *CopyOptions) cp(src *FileInfo, dest *FileInfo) {
	for _, f := range GetFiles(src) {
		var absPath string
		absPath = strings.Replace(f.Fullname, path.Dir(o.SrcfileInfo.Fullname), "", 1)
		absPath = path.Join(path.Dir(o.DestfileInfo.Fullname), absPath)
		d := &FileInfo{
			Volname:  dest.Volname,
			Fullname: absPath,
			Name:     dest.Name,
			IsDir:    dest.IsDir,
			IsVolume: dest.IsVolume,
			VolObj:   dest.VolObj,
			Perm:     f.Perm,
		}
		if !f.IsDir {
			if f.IsVolume && d.IsVolume {
				o.copyVolumeToVolume(f, d)
			} else if f.IsVolume {
				o.copyFromVolume(f, d)
			} else if d.IsVolume {
				o.copyToVolume(f, d)
			} else {
				fmt.Fprintln(config.ErrOut, "Please use the file path in glusterfs format like  gfs://volume/testdir/")
				os.Exit(2)
			}
			if o.View {
				fmt.Fprintln(config.Out, "Copied: ", d.Fullname)
			}
		} else {
			if d.IsVolume {
				d.VolObj.MkdirAll(absPath, d.Perm)
			} else {
				os.MkdirAll(absPath, d.Perm)
			}
			if o.View {
				fmt.Fprintln(config.Out, "Madedir:", absPath)
			}
		}
		if o.Recursive && f.IsDir {
			o.cp(f, dest)
		}
	}
}

func (o *CopyOptions) openLocalFileToWrite(fileInfo *FileInfo, fileName string) *os.File {
	file := &os.File{}
	err := errors.New("")
	var fileCreate = func(dest string) (*os.File, error) {
		if !o.Force && tools.CheckExist(dest) {
			fmt.Fprintln(config.Out, dest, "local dest file exist,  force（-f） to overwrite.")
			os.Exit(2)
		}
		return os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, fileInfo.Perm)
	}
	if !tools.CheckExist(path.Dir(fileInfo.Fullname)) {
		os.MkdirAll(path.Dir(fileInfo.Fullname), fileInfo.Perm)
	}
	file, err = fileCreate(fileInfo.Fullname)
	if err != nil {
		fmt.Fprintln(config.ErrOut, "Open local file error: ", err.Error())
		os.Exit(2)
	}
	return file
}
func (o *CopyOptions) openVolumeFileToWrite(volume *gfapi.Volume, fileInfo *FileInfo, fileName string) (*gfapi.File, error) {
	file := &gfapi.File{}
	err := errors.New("")
	var fileCreate = func(dest string) (*gfapi.File, error) {
		if !o.Force && tools.GfCheckExist(fileInfo.VolObj, dest) {
			fmt.Fprintln(config.Out, dest, " dest file exist,  force（-f） to overwrite.")
			os.Exit(2)
		}
		return volume.OpenFile(dest, os.O_WRONLY|os.O_CREATE, fileInfo.Perm)
	}
	if !tools.GfCheckExist(fileInfo.VolObj, path.Dir(fileInfo.Fullname)) {
		fileInfo.VolObj.MkdirAll(path.Dir(fileInfo.Fullname), fileInfo.Perm)
	}
	file, err = fileCreate(fileInfo.Fullname)
	if err != nil {
		fmt.Fprintln(config.ErrOut, "Open gluster volume file error ", err.Error())
	}
	return file, err
}
func (o *CopyOptions) copyToVolume(SrcfileInfo *FileInfo, DestfileInfo *FileInfo) {
	srcFile, err := os.Open(SrcfileInfo.Fullname)
	if err != nil {
		fmt.Fprintln(config.ErrOut, "Open src file error:", err)
		return
	}
	defer srcFile.Close()
	destFile, err := o.openVolumeFileToWrite(DestfileInfo.VolObj, DestfileInfo, SrcfileInfo.Name)
	if err != nil {
		fmt.Fprintln(config.ErrOut, "Open dest file error:", err)
		return
	}
	defer destFile.Close()
	copyError := tools.CopyFile(srcFile, destFile)
	if copyError != nil {
		fmt.Fprintln(config.ErrOut, "Copy file error: ", err)
	}
}

func (o *CopyOptions) copyFromVolume(SrcfileInfo *FileInfo, DestfileInfo *FileInfo) {
	srcFile, err := SrcfileInfo.VolObj.Open(SrcfileInfo.Fullname)
	if err != nil {
		fmt.Fprintln(config.ErrOut, "Open src file error:", err)
		return
	}
	defer srcFile.Close()
	destFile := o.openLocalFileToWrite(DestfileInfo, SrcfileInfo.Name)
	defer destFile.Close()
	copyError := tools.CopyFile(srcFile, destFile)
	if copyError != nil {
		fmt.Fprintln(config.ErrOut, "Copy file error:", err)
	}
}

func (o *CopyOptions) copyVolumeToVolume(SrcfileInfo *FileInfo, DestfileInfo *FileInfo) {
	srcFile, err := o.SrcfileInfo.VolObj.Open(SrcfileInfo.Fullname)
	if err != nil {
		fmt.Fprintln(config.ErrOut, "Open src file error:", err)
	}
	defer srcFile.Close()
	destFile, err := o.openVolumeFileToWrite(DestfileInfo.VolObj, DestfileInfo, SrcfileInfo.Name)
	if err != nil {
		fmt.Fprintln(config.ErrOut, "Open dest file error:", err)
	}
	defer destFile.Close()
	copyError := tools.CopyFile(srcFile, destFile)
	if copyError != nil {
		fmt.Fprintln(config.ErrOut, "copy file error:", err)
	}
}
