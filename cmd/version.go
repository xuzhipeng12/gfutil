/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 7:23 下午
 **/
package cmd

import (
	"fmt"
	"gfutil/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "gfutil version information",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Version)
	},
}

func init() {
	cmdRoot.AddCommand(versionCmd)
}
