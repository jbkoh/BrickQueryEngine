package tripleconn

import (
    "fmt"
    "net/http"
    "io/ioutil"
)


type FusekiConnector struct{
    url string
}

func (fuseki FusekiConnector) RunQuery(qStr string) string{
    req, _ := http.NewRequest("GET", fuseki.url, nil)
    q := req.URL.Query()
    q.Add("query", qStr)
    req.URL.RawQuery = q.Encode()
    
    req.Header.Add("Accept","application/sparql-results+json")
    
    client := &http.Client{}
    resp,_ := client.Do(req)
    fmt.Println(resp.Status)
    
    htmlData, _ := ioutil.ReadAll(resp.Body) 
    htmlStr := string(htmlData)
    return htmlStr
}

func NewFusekiConnector(baseurl string) *FusekiConnector{
    f := new(FusekiConnector)
    f.url = baseurl
    return f
}
