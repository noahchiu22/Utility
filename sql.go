package util

import (
	"gorm.io/gorm"
)

// 檢查是否重複
func CheckIsRepeat(db *gorm.DB, tableName string, condition map[string]interface{}, exclusion map[string]interface{}) (isRepeat bool, err error) {
	count := int64(0)

	result := db.Debug().Table(tableName).Not(exclusion).Where(condition).Count(&count)

	if result.Error != nil {
		err = result.Error
		return
	}

	if count > 0 {
		isRepeat = true
		return
	}

	return
}

// 取得最大序或是最大編號
func FindMaxSN[SN string | int](db *gorm.DB, tableName, column string, condition map[string]interface{}) (maxSN SN, err error) {
	result := db.Debug().Table(tableName).
		Select(column).
		Where(condition).
		Order(column + " DESC").
		Limit(1).
		Find(&maxSN)

	if result.Error != nil {
		err = result.Error
		return
	}

	return
}
