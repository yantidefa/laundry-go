package config

import "github.com/jinzhu/gorm"

func DB() (*gorm.DB, error) {
	dbConn, err := gorm.Open("mysql", "root:0303@/laundry?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
