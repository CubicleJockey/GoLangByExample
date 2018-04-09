/*
Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.
*/

package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type JsonResponse struct {
    Page int
    Fruits []string
}

type JsonResponseTyped struct {
    Page int `json:"page"`
    Fruits []string `json:"fruits"`
}


func main() {
    booleanJsonAsBytes, _ := json.Marshal(true)
    fmt.Println(booleanJsonAsBytes)
    fmt.Println(string(booleanJsonAsBytes))

    integerJsonBytes, _ := json.Marshal(1)
    fmt.Println(integerJsonBytes)
    fmt.Println(string(integerJsonBytes))

    floatJsonBytes, _ := json.Marshal(2.34)
    fmt.Println(floatJsonBytes)
    fmt.Println(string(floatJsonBytes))

    stringJsonBytes, _ := json.Marshal("gopher")
    fmt.Println(stringJsonBytes)
    fmt.Println(string(stringJsonBytes))

    fruitSlice := []string{"apple", "peach", "pear"}
    fruitSliceJsonBytes, _ := json.Marshal(fruitSlice)
    fmt.Println("Fruit Slice:", fruitSlice)
    fmt.Println("Fruit Slice Json Byte Array:", fruitSliceJsonBytes)
    fmt.Println("Fruit Slice Json as String:", string(fruitSliceJsonBytes))

    fruitMapping := map[string]int{"apple": 5, "lettuce": 7}
    fruitMappingJsonBytes, _ := json.Marshal(fruitMapping)
    fmt.Println("Fruit Mapping:", fruitMapping)
    fmt.Println("Fruit Mapping Json Byte Array:", fruitMappingJsonBytes)
    fmt.Println("Fruit Mapping Json as String:", string(fruitMappingJsonBytes))

    fruits := []string{"apple", "peach", "pear"}

    response := JsonResponse { 1, fruits }
    responseJsonBytes, _ := json.Marshal(response)
    fmt.Println("Response Object: ", response)
    fmt.Println("Response Json Byte Array:", responseJsonBytes)
    fmt.Println("Response Json as String:", string(responseJsonBytes))

    typeResponse := JsonResponseTyped{1, fruits}
    typeResponseJsonBytes, _ := json.Marshal(response)
    fmt.Println("Typed Response Object: ", typeResponse)
    fmt.Println("Typed Response Json Byte Array:", typeResponseJsonBytes)
    fmt.Println("Typed Response Json as String:", string(typeResponseJsonBytes))

    stringAsBytes := []byte(`{"num":6.13,"strings":["a","b"]}`)
    var data map[string]interface{}

    fmt.Println("Convert bytes: ", stringJsonBytes, " to data object.")
    if err := json.Unmarshal(stringAsBytes, &data); err != nil {
        panic(err)
    }
    fmt.Println("Data object:", data)

    number := data["num"].(float64)
    fmt.Println("Key to 'num' property:", number)

    stringItems := data["strings"].([]interface{})

    items1 := stringItems[0].(string)
    fmt.Println("Item1:", items1)

    item2 := stringItems[1].(string)
    fmt.Println("Item2:", item2)

    fruitsJson := `{"page": 1, "fruits": ["apple", "peach"]}`
    fmt.Println("Json:", fruitsJson)

    jsonResponseTyped := JsonResponseTyped {}

    json.Unmarshal([]byte(fruitsJson), &jsonResponseTyped)
    fmt.Println(jsonResponseTyped)
    fmt.Println("Page:", jsonResponseTyped.Page)
    fmt.Println("Item1:", jsonResponseTyped.Fruits[0])
    fmt.Println("Item2:", jsonResponseTyped.Fruits[1])

    fmt.Println()
    fmt.Println("Encoding to Stdout")
    encode := json.NewEncoder(os.Stdout)
    mappedData := map[string]int{"apple": 5, "lettuce": 7}
    encode.Encode(mappedData)


}