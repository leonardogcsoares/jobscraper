package backend

import (
	"os"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// LastItem TODO
var LastItem *mgo.Collection

// Start TODO
func Start() error {
	var err error
	os.Setenv("DATABASE_URL", "mongodb://admin:admin@ds017185.mlab.com:17185/jobscraper")
	session, err = mgo.Dial(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	LastItem = session.DB("jobscraper").C("lastItem")

	return nil
}

// Close TODO
func Close() {
	session.Close()
}
