/**
// * @Author xuzhipeng
// * @Description
// * @Date 2021/11/27 9:24 下午
// **/
//package pkg
//

package pkg

func Umound(info *FileInfo) {
	if info.VolObj != nil {
		info.VolObj.Unmount()
	}
}
