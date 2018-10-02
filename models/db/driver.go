/*
* @Author: Shubham Bansal
* @Date:   2018-10-02 15:22:34
* @Last Modified by:   Shubham Bansal
* @Last Modified time: 2018-10-02 17:53:24
*/
package db

import "github.com/astaxie/beego"

var maxPool int

func init() {
	var err error
	maxPool, err = beego.AppConfig.Int("DBMaxPool")
	if err!= nil {
		// todo : panic!!!
		// panic(err)
		println(err)
	}
	// init method to start db
	checkAndInitServiceConnection()
}

func checkAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = beego.AppConfig.String("DBPath")
		err := service.New()
		if err != nil {
			// todo: panic!!!
			// panic(err)
			println(err)
		}
	}
}