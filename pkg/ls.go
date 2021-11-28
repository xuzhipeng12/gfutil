/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 11:05 下午
 **/
package pkg

import (
	"fmt"
	"gfutil/config"
	"gfutil/tools"
	"os"
)

type LsOptions struct {
	Hosts       []string
	LongFormat  bool
	SrcfileInfo *FileInfo
}

func (l *LsOptions) LS(args ...string) {
	l.SrcfileInfo = GetFileInfo(args[0], l.Hosts)
	defer Umound(l.SrcfileInfo)
	if l.SrcfileInfo.IsVolume {
		fileFinfos := GetFiles(l.SrcfileInfo)
		for _, f := range fileFinfos {
			if l.LongFormat {
				fmt.Fprintf(config.Out, "%-11s %-7s %s  ", f.Perm, tools.GetFriendlyFileSize(f.Size), f.ModTime.Format("2006-01-02 15:04:05"))
			}
			if f.IsDir {
				fmt.Fprintf(config.Out, tools.SprintDirectory(f.Name))
			} else {
				fmt.Fprintf(config.Out, tools.SprintFile(f.Name))
			}
			fmt.Fprintf(config.Out, " ")
			if l.LongFormat {
				fmt.Fprintln(config.Out)
			}
		}
		fmt.Fprintln(config.Out)
	} else {
		fmt.Fprintln(config.ErrOut, "Error, input glusterfs format path")
		os.Exit(2)
	}
}
