// Testing 3DES on a set of known vectors
package main

import (
    "crypto/des"
    "encoding/hex"
    "fmt"
)

type testVector struct {
    key string
    plaintext string
    encrypted string
}

func test3DES(tv * testVector) int {
    key,_ := hex.DecodeString(tv.key)
    plaintext,_ := hex.DecodeString(tv.plaintext)

    c, err := des.NewTripleDESCipher(key)
    if err!=nil {
        fmt.Println(err)
        return 1
    }

    encrypted := make([]byte, len(plaintext))
    c.Encrypt(encrypted, plaintext)

    hexEncrypted := hex.EncodeToString(encrypted)
    fmt.Printf("exp:[%s]\n", tv.encrypted)
    fmt.Printf("got:[%s]\n", hexEncrypted)
    errors:=0
    if tv.encrypted != hexEncrypted {
        errors++
    }
    return errors

}

func main() {
    tv := []testVector {
        {
            "000000000000000000000000000000000000000000000000",
            "0000000000000000",
            "8ca64de9c1b123a7",
        },
        {
            "ffffffffffffffffffffffffffffffffffffffffffffffff",
            "ffffffffffffffff",
            "7359b2163e4edc58",
        },
        {
            "300000000000000030000000000000003000000000000000",
            "1000000000000001",
            "958e6e627a05557b",
        },
        {
            "111111111111111111111111111111111111111111111111",
            "1111111111111111",
            "f40379ab9e0ec533",
        },
        {
            "0123456789abcdef0123456789abcdef0123456789abcdef",
            "1111111111111111",
            "17668dfc7292532d",
        },
        {
            "111111111111111111111111111111111111111111111111",
            "0123456789abcdef",
            "8a5ae1f81ab8f2dd",
        },
        {
            "000000000000000000000000000000000000000000000000",
            "0000000000000000",
            "8ca64de9c1b123a7",
        },
        {
            "fedcba9876543210fedcba9876543210fedcba9876543210",
            "0123456789abcdef",
            "ed39d950fa74bcc4",
        },
        {
            "7ca110454a1a6e577ca110454a1a6e577ca110454a1a6e57",
            "01a1d6d039776742",
            "690f5b0d9a26939b",
        },
        {
            "0131d9619dc1376e0131d9619dc1376e0131d9619dc1376e",
            "5cd54ca83def57da",
            "7a389d10354bd271",
        },
        {
            "07a1133e4a0b268607a1133e4a0b268607a1133e4a0b2686",
            "0248d43806f67172",
            "868ebb51cab4599a",
        },
        {
            "3849674c2602319e3849674c2602319e3849674c2602319e",
            "51454b582ddf440a",
            "7178876e01f19b2a",
        },
        {
            "04b915ba43feb5b604b915ba43feb5b604b915ba43feb5b6",
            "42fd443059577fa2",
            "af37fb421f8c4095",
        },
        {
            "0113b970fd34f2ce0113b970fd34f2ce0113b970fd34f2ce",
            "059b5e0851cf143a",
            "86a560f10ec6d85b",
        },
        {
            "0170f175468fb5e60170f175468fb5e60170f175468fb5e6",
            "0756d8e0774761d2",
            "0cd3da020021dc09",
        },
        {
            "43297fad38e373fe43297fad38e373fe43297fad38e373fe",
            "762514b829bf486a",
            "ea676b2cb7db2b7a",
        },
        {
            "07a7137045da2a1607a7137045da2a1607a7137045da2a16",
            "3bdd119049372802",
            "dfd64a815caf1a0f",
        },
        {
            "04689104c2fd3b2f04689104c2fd3b2f04689104c2fd3b2f",
            "26955f6835af609a",
            "5c513c9c4886c088",
        },
        {
            "37d06bb516cb754637d06bb516cb754637d06bb516cb7546",
            "164d5e404f275232",
            "0a2aeeae3ff4ab77",
        },
        {
            "1f08260d1ac2465e1f08260d1ac2465e1f08260d1ac2465e",
            "6b056e18759f5cca",
            "ef1bf03e5dfa575a",
        },
        {
            "584023641aba6176584023641aba6176584023641aba6176",
            "004bd6ef09176062",
            "88bf0db6d70dee56",
        },
        {
            "025816164629b007025816164629b007025816164629b007",
            "480d39006ee762f2",
            "a1f9915541020b56",
        },
        {
            "49793ebc79b3258f49793ebc79b3258f49793ebc79b3258f",
            "437540c8698f3cfa",
            "6fbf1cafcffd0556",
        },
        {
            "4fb05e1515ab73a74fb05e1515ab73a74fb05e1515ab73a7",
            "072d43a077075292",
            "2f22e49bab7ca1ac",
        },
        {
            "49e95d6d4ca229bf49e95d6d4ca229bf49e95d6d4ca229bf",
            "02fe55778117f12a",
            "5a6b612cc26cce4a",
        },
        {
            "018310dc409b26d6018310dc409b26d6018310dc409b26d6",
            "1d9d5c5018f728c2",
            "5f4c038ed12b2e41",
        },
        {
            "1c587f1c13924fef1c587f1c13924fef1c587f1c13924fef",
            "305532286d6f295a",
            "63fac0d034d9f793",
        },
        {
            "010101010101010101010101010101010101010101010101",
            "0123456789abcdef",
            "617b3a0ce8f07100",
        },
        {
            "1f1f1f1f0e0e0e0e1f1f1f1f0e0e0e0e1f1f1f1f0e0e0e0e",
            "0123456789abcdef",
            "db958605f8c8c606",
        },
        {
            "e0fee0fef1fef1fee0fee0fef1fef1fee0fee0fef1fef1fe",
            "0123456789abcdef",
            "edbfd1c66c29ccc7",
        },
        {
            "000000000000000000000000000000000000000000000000",
            "ffffffffffffffff",
            "355550b2150e2451",
        },
        {
            "ffffffffffffffffffffffffffffffffffffffffffffffff",
            "0000000000000000",
            "caaaaf4deaf1dbae",
        },
        {
            "0123456789abcdef0123456789abcdef0123456789abcdef",
            "0000000000000000",
            "d5d44ff720683d0d",
        },
        {
            "fedcba9876543210fedcba9876543210fedcba9876543210",
            "ffffffffffffffff",
            "2a2bb008df97c2f2",
        },
    }

    total_errs:=0
    for _,t := range tv {
        err:=test3DES(&t)
        if err>0 {
            fmt.Printf("failed\n")
            total_errs+=err
        } else {
            fmt.Printf("ok\n")
        }
    }
    fmt.Printf("total errors: %d\n", total_errs)
}
