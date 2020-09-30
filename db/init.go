package db

import (
	"github.com/l552121229/golang-tools/db/dbhandle"

	_ "github.com/l552121229/golang-tools/db/mysql"
)

func init() {
	conf, err := dbhandle.GetConfig()
	if err != nil {
		panic("init database failed." + err.Error())
	}
	Default, err = conf.Get("DB.type")
	if err != nil {
		panic("get default database type failed." + err.Error())
	}
	InitDB(Default)
}
