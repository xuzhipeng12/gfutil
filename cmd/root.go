/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 5:45 下午
 **/
package cmd

import (
	"fmt"
	"gfutil/config"
	"github.com/spf13/cobra"
	"os"
)

var (
	hosts    string
	helpFlag bool
)
var cmdRoot = &cobra.Command{
	Use:   "gfutil",
	Short: "gfutil is a tool for operating gluster volume files \nUsage: gfutils [command] [args...] [options...]",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		if hosts == "" {
			envHosts := os.Getenv(config.GlusterfsHostsEnvariable)
			if envHosts == "" {
				fmt.Println("Require hosts not set.  Use -h set hosts or set os env " + config.GlusterfsHostsEnvariable)
				os.Exit(1)
			}
			hosts = envHosts
		}
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	cmdRoot.PersistentFlags().StringVarP(&hosts, "hosts", "h", "", "Hosts accepts one or more hostname(s) and/or IP(s)")
	cmdRoot.PersistentFlags().BoolVarP(&helpFlag, "help", "", false, "Help default flag")
	err := cmdRoot.Execute()
	if err != nil {
		fmt.Println(err)
	}

}
