/*
* @Author: Shubham Bansal
* @Date:   2018-10-02 15:16:43
* @Last Modified by:   Shubham Bansal
* @Last Modified time: 2018-10-02 15:20:26
*/
package db

import "gopkg.in/mgo.v2"

type Service struct {
	baseSession *mgo.Session
	queue chan int
	URL string
	Open int
}

var service Service

func(s *Service) New() error {
	var err error
	s.queue = make(cha int, maxPool)
	fo i:=0; i < maxPool; i = i + 1 {
		s.queue < -1
	}
	s.Open = 0
	s.baseSession , err = mgo.Dial(s.URL)
	return err
}

func(s *Service) Session() *mgo.Session {
	<-s.queue
	s.Open++
	return s.baseSession.Copy()
}

func(s *Service) Close(c *Collection) {
	c.db.s.Close()
	s.queue <- 1
	s.Open -
}