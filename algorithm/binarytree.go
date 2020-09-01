package main

import (
    "encoding/json"
    "fmt"
    //"os"
)

type All struct{
	It []Item `json:"item"`
}
type Item struct{
	Name string `json:"name"`
	Picture string `json:"picture"`
}

func main() {

    byt := []byte(`{"item":[{"name":"chandler","picture":"111"},{"name":"oscar","picture":"222"}]}`)
    fmt.Println(byt)

    var dat All

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat.It[0].Name)
	It1:=Item{Name:"stephen",Picture:"333"}

	dat.It=append(dat.It,It1)

   fmt.Println(dat)
}