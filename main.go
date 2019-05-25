package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    fmt.Println("Starting the application...")
    path := "http://" + os.Getenv("TEST4_SERVICE_HOST")  + ":" + os.Getenv("TEST4_SERVICE_PORT") 
    fmt.Println("Trying to get from", path)
    response, err := http.Get(path)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    jsonData := map[string]string{"ID": "1", "message": "Raboy"}
    jsonValue, _ := json.Marshal(jsonData)
    response, err = http.Post(path, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    fmt.Println("Terminating the application...")
}
