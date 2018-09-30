package main

import (
  "fmt"
  "os"
  "strings"
    "net/http"
    "io/ioutil"
  flag "github.com/ogier/pflag"
)

// flags
var (
   user  string
)

func main() {
  httpGet()
    fmt.Println("Hello, World")
 flag.Parse() // if user does not supply flags, print usage
 // we can clean this up later by putting this into its own function
  if flag.NFlag() == 0 {
     fmt.Printf("Usage: %s [options]\n", os.Args[0])
     fmt.Println("Options:")
     flag.PrintDefaults()
     os.Exit(1)
  }
  users := strings.Split(user, ",")
  fmt.Printf("Searching user(s): %s\n", users)
}

func init() {
 flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func httpGet() {
    response, err := http.Get("http://golang.org/")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }
}

func httpPost() {

    url := "http://restapi3.apiary.io/notes"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
