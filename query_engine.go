package main

import (
    "os"
    "fmt"
    //"html"
    //"log"
    "net/http"
    //"net/url"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    //"container/list"
    "io/ioutil"
    "strings"
    //"bytes"
)

type SparqlQuery struct {
        Name string
        Query string
}

type FusekiInterface struct{
    url string
}

func (fuseki FusekiInterface) runQuery(qStr string) string{
   // data := url.Values{}
    //data.Add("query", qStr)
    
    req, _ := http.NewRequest("GET", fuseki.url, nil)
    //req, _ := http.NewRequest("GET", fuseki.url, bytes.NewBufferString(data.Encode()))
    q := req.URL.Query()
    q.Add("query", qStr)
    req.URL.RawQuery = q.Encode()
    
    req.Header.Add("Accept","application/sparql-results+json")
    client := &http.Client{}
    resp,_ := client.Do(req)
    fmt.Println(resp.Status)
    
    htmlData, _ := ioutil.ReadAll(resp.Body) //<--- here!
    fmt.Println(os.Stdout, string(htmlData))
    fmt.Println("=========================")
    fmt.Println(qStr)
    fmt.Println("=========================")
    fmt.Println(req.URL.String())
    
    
//    _ = json.NewDecoder(resp.Body).Decode(&result)
    
    return "TEST"
}

func NewFusekiInterface(baseurl string) *FusekiInterface{
    f := new(FusekiInterface)
    f.url = baseurl
    return f
}

func readQueryFromDb(c *mgo.Collection, qName string) string{
    result := SparqlQuery{}
    _ = c.Find(bson.M{"_id":qName}).One(&result)
    return result.Query
}

func main() {
    
    session, err := mgo.Dial("127.0.0.1:27017")
    if err !=nil{
        panic(err)
    }
    defer session.Close()
    
    c := session.DB("Brick").C("QueryDictionary")
//    err = c.Insert(q)
    
    // Read query files and Make a SParqlQuery list
    files, err := ioutil.ReadDir("./queries/")
    for _, file := range files{
        filename := file.Name()
        qBuf, _ := ioutil.ReadFile("./queries/"+filename)
        qStr := string(qBuf)
        qName := strings.Split(filename, ".")[0]
        fmt.Println(qName)
        qEntity := SparqlQuery{Name:qName, Query:qStr}
        _,_ = c.UpsertId(qName, qEntity)
    }
    
    queryName := "test_query"
    fmt.Println(queryName)
    
//    fuseki := FusekiInterface{url:"http:localhost:3030/brick"}
    fuseki := NewFusekiInterface("http://localhost:3030/brick")
    qStr := readQueryFromDb(c, queryName)
    fmt.Println(fuseki.runQuery(qStr))
    

    
    /*
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
*/
}