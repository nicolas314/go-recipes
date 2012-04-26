// Random password generator
// Generate a blunt random password
package main
import (
    "fmt"
    "os"
    "strconv"
)

// Generate a password of length sz
func GeneratePassword(sz int) string {
    // Character set for password chars: extend if needed
    letterSet := "1234567890" +
                 "abcdefghijklmnopqrstuvwxyz" +
                 "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

    // Get sz random bytes
    pick := make([]byte, sz)
    f,_:=os.Open("/dev/urandom")
    _,_ = f.Read(pick)
    f.Close()

    // Pick sz characters at random
    passwd:=""
    for i:=0 ; i<sz ; i++ {
        passwd=passwd+string(letterSet[int(pick[i]) % len(letterSet)])
    }
    return passwd
}

func main() {
    if len(os.Args)!=2 {
        fmt.Println("use: generate-password SIZE")
        return
    }
    sz, err := strconv.Atoi(os.Args[1])
    if err!=nil {
        fmt.Println(err)
        return
    }
    fmt.Println(GeneratePassword(sz))
}
