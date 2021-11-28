/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/26 11:18 下午
 **/
package cmd

import (
	"fmt"
	"gfutil/pkg"
	"github.com/spf13/cobra"
	"strings"
)

var (
	longFormat bool
)
var cmdLs = &cobra.Command{
	Use:   "ls [src]",
	Short: "List files",
	Example: `List files 
    Examples: 
    #from gluster volume:
    gfutil ls gfs://volume/testdir/
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		l := pkg.LsOptions{
			LongFormat: longFormat,
			Hosts:      strings.Split(hosts, ","),
		}
		l.LS(args...)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Close Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
}

func init() {
	cmdLs.PersistentFlags().BoolVarP(&longFormat, "long-format", "l", false, "List information in a longer format")
	cmdRoot.AddCommand(cmdLs)
}
