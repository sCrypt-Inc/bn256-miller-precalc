package main

import (
    "fmt"
    "os"
)

func main() {

    argsWithoutProg := os.Args[1:]

    alpha := &curvePoint{
      	bigFromBase16(argsWithoutProg[0]),
      	bigFromBase16(argsWithoutProg[1]),
      	bigFromBase16("1"),
      	bigFromBase16("1"),
    }

    beta := &twistPoint{
        &gfP2{
            bigFromBase16(argsWithoutProg[3]),
            bigFromBase16(argsWithoutProg[2]),
        },
        &gfP2{
            bigFromBase16(argsWithoutProg[5]),
            bigFromBase16(argsWithoutProg[4]),
        },
        &gfP2{
            bigFromBase16("0"),
            bigFromBase16("1"),
        },
        &gfP2{
            bigFromBase16("0"),
            bigFromBase16("1"),
        },
    }

    pool := new(bnPool)

    millerBetaAlpha := miller(beta, alpha, pool)
    fmt.Println(millerBetaAlpha.x.x.x)
    fmt.Println(millerBetaAlpha.x.x.y)
    fmt.Println(millerBetaAlpha.x.y.x)
    fmt.Println(millerBetaAlpha.x.y.y)
    fmt.Println(millerBetaAlpha.x.z.x)
    fmt.Println(millerBetaAlpha.x.z.y)
    fmt.Println(millerBetaAlpha.y.x.x)
    fmt.Println(millerBetaAlpha.y.x.y)
    fmt.Println(millerBetaAlpha.y.y.x)
    fmt.Println(millerBetaAlpha.y.y.y)
    fmt.Println(millerBetaAlpha.y.z.x)
    fmt.Println(millerBetaAlpha.y.z.y)
}
