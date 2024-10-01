package util

import (
	"os"
)

// 檢查 檔案/資料夾 是否存在
func IsDirExist(path, target string) (bool, error) {
	filenames, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	// 跑過每個檔案
	for _, filename := range filenames {
		if filename.Name() == target { // 如果有個檔案的檔名跟目標檔明一樣
			return true, nil
		}
	}

	return false, nil
}

// 檢查並建立資料夾
func CheckAndMakeDir(pathArr []string) (path string, err error) {
	// root
	path = "./"
	// 跑過每個資料夾名稱
	for _, dir := range pathArr {
		// 檢查當前路徑是否有資料夾
		exist, _ := IsDirExist(path, dir)

		// 更新路徑
		path += dir

		// 如果當前路徑沒有資料夾，創建一個
		if !exist {
			if err = os.Mkdir(path, 0777); err != nil {
				Log("新增資料夾失敗", path, err.Error())
				return
			}
		}

		path += "/"
	}

	return
}
