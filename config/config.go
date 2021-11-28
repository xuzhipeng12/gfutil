/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 6:20 下午
 **/
package config

import "os"

const (
	Version                  = "gfutil v0.1"
	GlusterfsHostsEnvariable = "GLUSTER_HOSTS"
	GlusterFilePrefix        = "gfs://"
	CopyFileBufferSize       = 1024 * 1024 * 10
)

var (
	// In think, os.Stdin
	In = os.Stdin
	// Out think, os.Stdout
	Out = os.Stdout
	// ErrOut think, os.Stderr
	ErrOut = os.Stderr
)
