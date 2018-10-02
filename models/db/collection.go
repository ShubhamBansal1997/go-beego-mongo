/*
* @Author: Shubham Bansal
* @Date:   2018-10-02 15:30:01
* @Last Modified by:   Shubham Bansal
* @Last Modified time: 2018-10-02 17:52:38
*/
package db

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

type Collection struct {
	db *Database
	name string
	Session *mgo.Collection
}

func (c *Collection) Connect() {
	session := *c.db.session.C(c.name)
	c.Session = &session
}

func NewCollectionSession(name string) *Collection {
	var c = Collection{
		db: newDBSession(beego.AppConfig.String("DBName")),
		name: name,
	}
	c.Connect()
	return &c
}

func (c *Collection) Close() {
	service.Close(c)
}