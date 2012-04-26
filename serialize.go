// Object serialization/de-serialization using JSON
// In this example a list of users is stored as a map
// username -> user structure
// that can be saved to JSON and restored
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

// Type we want to serialize
// Careful: field names need to start with a capital letter to be exported
// to JSON otherwise they are ignored. The corresponding JSON names can be
// specified by adding a tag to each field declaration.
type User struct {
    Password    string  `json:"password"`
    Age         int     `json:"age"`
}

type UserList map[string] User

func (ul UserList) Save(filename string) {
    js, _ := json.Marshal(ul)
    ioutil.WriteFile(filename, js, 0644)
}

func (ul UserList) Load(filename string) {
    js, _ := ioutil.ReadFile(filename)
    json.Unmarshal(js, &ul)
}


// Demonstration: list of users created in one function
// then saved to JSON and restored in another function.
func create_and_save() {
    ul := make(UserList)
    ul["alice"] = User{"ALICE", 42}
    ul["bob"] = User{"ROBERT", 21}
    ul.Save("users.json")
}
func restore_and_display() {
    ul := make(UserList)
    ul.Load("users.json")
    fmt.Println(ul)
}

func main() {
    create_and_save()
    restore_and_display()
}

