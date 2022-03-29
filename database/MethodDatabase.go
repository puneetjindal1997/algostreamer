package database

import "os"

func (mgr *manager) GetAllAppIds() (appIds []int, err error) {
	resp := mgr.connection.Table(os.Getenv("APPIDTABLE")).Select("application_id").Find(&appIds)
	err = resp.Error
	return appIds, err
}
