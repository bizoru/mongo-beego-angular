package db

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"reflect"
)

func Cursor(session *mgo.Session, Collection string) *mgo.Collection {
	mongoDB := beego.AppConfig.String("mongo_db")
	c := session.DB(mongoDB).C(Collection)
	return c
}

func GetSession() (*mgo.Session,error) {

	mongoHost := beego.AppConfig.String("mongo_host")
	mongoUser := beego.AppConfig.String("mongo_user")
	mongoPassword := beego.AppConfig.String("mongo_pass")
	mongoDatabase := beego.AppConfig.String("mongo_db")

	info := &mgo.DialInfo{
		Addrs:    []string{mongoHost},
		Timeout:  60 * time.Second,
		Database: mongoDatabase,
		Username: mongoUser,
		Password: mongoPassword,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println("Helo this is an error!")
		panic(err)
	} else{
		session.SetMode(mgo.Monotonic, true)
	}

	return session,err
}



func GetAll(session *mgo.Session, collection string) []bson.M {
	c :=  Cursor(session, collection)
	defer session.Close()
	var records []bson.M
	err := c.Find(bson.M{}).All(&records)
	if err != nil {
		fmt.Println(err)
	}

	return records
}

func getType(i interface{}) interface{}{
	return reflect.New(reflect.TypeOf(i))
}