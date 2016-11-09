package main

import (
    "fmt"
//    "gopkg.in/mgo.v2"
    "io/ioutil"
    "strings"
    "./tripleconn"
    "./dbconn"
)

func main() {
    
    // Init Mongodb connection
    mongoConn := dbconn.NewMongoConn("127.0.0.1:27017")
    
    // Read query files and store queries into mongodb.
    files, _ := ioutil.ReadDir("./queries/")
    for _, file := range files{
        filename := file.Name()
        qBuf, _ := ioutil.ReadFile("./queries/"+filename)
        qStr := string(qBuf)
        qName := strings.Split(filename, ".")[0]
        mongoConn.WriteQueryToDb(qName, qStr)
        fmt.Println(qName)
    }
    
    // choose sample query
    queryName := "test_query"
    
    // Init Fuseki Connector. Send SPARQL query string to the connector.
    fuseki := tripleconn.NewFusekiConnector("http://localhost:3030/brick")
    qStr := mongoConn.ReadQueryFromDb(queryName)
    qResult := fuseki.RunQuery(qStr)
    fmt.Println(qStr)
    fmt.Println(qResult)
    
    /* // Run server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
    */
}