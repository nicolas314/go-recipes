// A tool to compute various crypto hashes on files
// Demonstrates: crypto hashes and buffered I/O
package main

import (
    "bufio"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "errors"
    "fmt"
    "hash"
    "os"
)

func Hash(hashname string, filename string) (sum string, err error) {
    f, err := os.Open(filename)
    if err!=nil {
        return
    }

    var h hash.Hash

    switch hashname {
        case "md5":
        h = md5.New()
        case "sha1":
        h = sha1.New()
        case "sha2", "sha256":
        h = sha256.New()
        case "sha5", "sha512":
        h = sha512.New()
        default:
        err = errors.New("unknown hash: "+hashname)
        return
    }

    var nr int
    buf := make([]byte, h.BlockSize())
    bf  := bufio.NewReader(f)
    for {
        nr, _ = bf.Read(buf)
        h.Write(buf[0:nr])
        if nr<len(buf) {
            break
        }
    }
    f.Close()
    sum = fmt.Sprintf("%0x", h.Sum(nil))
    return
}

func main() {
    if len(os.Args)<3 {
        fmt.Println("use: md5|sha1|sha2|sha5 filenames...")
        return
    }
    for i:=2 ; i<len(os.Args) ; i++ {
        digest, err := Hash(os.Args[1], os.Args[i])
        if err!=nil {
            fmt.Println(err)
        } else {
            fmt.Printf("%s  %s\n", digest, os.Args[i])
        }
    }
}
