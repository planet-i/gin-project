package export

import "github.com/planet-i/gin-project/pkg/setting"

// 获取导出Excel文件完整的访问URL
func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

// 获取导出Excel文件的路径
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// 获取导出Excel文件的完整路径
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
