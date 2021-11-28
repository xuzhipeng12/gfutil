/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 8:24 下午
 **/
package cmd

import (
	"gfutil/pkg"
	"github.com/spf13/cobra"
	"strings"
)

var (
	recursive bool
	force     bool
	view      bool
)

var cmdCp = &cobra.Command{
	Use:   "cp [src] [desc]",
	Short: "Copy files",
	Example: `Copy files between volume and local, volume and volum
    Examples: 
    #Copy a file from gluster volume:
    gfutil cp gfs://volume/testdir/testfile.txt /data/testdir/ 
    #Copy a file to gluster:
    gfutil cp /data/testdir/testfile.txt gfs://volume/testdir/
    #Ccopy a dir from gluster volume:
    gfutil cp gfs://volume/testdir/a.txt gfs://volume/testdir/b.txt
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		o := pkg.CopyOptions{
			Hosts:     strings.Split(hosts, ","),
			Recursive: recursive,
			Force:     force,
			View:      view,
		}
		o.CP(args...)
	},
}

func init() {
	cmdCp.PersistentFlags().BoolVarP(&recursive, "recursive", "r", false, "Recursive copies the contents of directories")
	cmdCp.PersistentFlags().BoolVarP(&force, "force", "f", false, "Overwrite files and directories if exist")
	cmdCp.PersistentFlags().BoolVarP(&view, "view", "v", false, "View copy detail")
	cmdRoot.AddCommand(cmdCp)
}
