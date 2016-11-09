package dbconn

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type SparqlQuery struct {
        Name string
        Query string
}

type MongoConn struct{
    c *mgo.Collection
}

func NewMongoConn (urlStr string) *MongoConn{
    m := new(MongoConn)
    session, err := mgo.Dial(urlStr)
    if err != nil{
        panic(err)
    }
//    defer session.Close() //How should I terminate the session?
    m.c = session.DB("Brick").C("QueryDictionary")
    return m
}

func (mongoConn MongoConn) ReadQueryFromDb(qName string) string{
    result := SparqlQuery{}
    _ = mongoConn.c.Find(bson.M{"_id":qName}).One(&result)
    return result.Query
}

func (mongoConn MongoConn) WriteQueryToDb(qName string, qStr string){
    qEntity := SparqlQuery{Name:qName, Query:qStr}
    _,_ = mongoConn.c.UpsertId(qName, qEntity)

}