/*
	Maps are Go’s built-in associative data type (sometimes called hashes or dictionaries in other languages).
*/
package main

import "fmt"

func main(){

    const KEY1 = "key1"
    const KEY2 = "key2"
    const NUMBERMESSAGES = "Number of items in mapping: "
    const MAPPINGMESSAGE = "Mapping: "

    //map[key-type]val-type
    mapping := make(map[string]int)

    mapping[KEY1] = 7
    mapping[KEY2] = 13

    fmt.Println(MAPPINGMESSAGE, mapping)

    fmt.Println("Getting first value...")
    value1 := mapping[KEY1]
    fmt.Println("Value #1", value1)

    fmt.Println("Getting second value...")
    value2 := mapping[KEY2]
    fmt.Println("Value #2", value2)

    fmt.Println(NUMBERMESSAGES, len(mapping))

    fmt.Println("Deleting item #2...")
    delete(mapping, KEY2)
    fmt.Println("Item deleted.")

    fmt.Println(NUMBERMESSAGES, len(mapping))
    fmt.Println(MAPPINGMESSAGE, mapping)

    //Optional second parameter on get value
    getDeletedValue, isDeletedValuePresent := mapping[KEY2]

    fmt.Println("Was deleted value present? ", isDeletedValuePresent)
    fmt.Println("Missing values value was: ", getDeletedValue)

    getExistingValue, isNoneDeletedValuePresent := mapping[KEY1]

    fmt.Println("Was none deleted value present? ", getExistingValue)
    fmt.Println("None missing values value was: ", isNoneDeletedValuePresent)

    //declare and initialize
    thingy := map[string]int {"item1" : 1, "item2" : 2}
    fmt.Println("Declared and Initalized at once: ", thingy)
}
