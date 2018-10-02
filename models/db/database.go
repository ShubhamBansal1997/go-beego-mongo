/*
* @Author: Shubham Bansal
* @Date:   2018-10-02 15:20:32
* @Last Modified by:   Shubham Bansal
* @Last Modified time: 2018-10-02 15:22:29
*/
package db

import "gopkg.in/mgo.v2"

type Database struct {
	s *mgo.Session
	name String
	session *mgo.Database
}

func (db *Database) Connect() {
	db.s = service.Session()
	session := *db.s.DB(db.name)
	db.session = &session
}

func newDBSession(name String) *Database {
	var db = Database{
		name:name,
	}
	db.Connect()
	return &db
}