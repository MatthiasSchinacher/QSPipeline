package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Println("Welcome to QSPipline :-) - ",time.Now());
    start_qspipeline()
    
    //a := [...]int{10,3,19,2,4,16,13,14,15,8,5,6,12,11,7,0,9,1,17,18}
    var a [50000000]int
    for i:=0;i<50000000;i++ {
        a[i] = rand.Intn(100000000)    
    }
    
    fmt.Println("sorting ...       - ",time.Now())
    qs(a[0:len(a)],pipeline)

    fmt.Println("checking ...      - ",time.Now())
    for i:=1; i<len(a); i++ {
        if a[i-1] > a[i] {
            fmt.Println("ERROR: i=",i," a[i-1]=",a[i-1]," a[i]=",a[i])
            break
        }
    }
    fmt.Println("shutting down ... - ",time.Now())
    stop_qspipeline()
    fmt.Println("done.             - ",time.Now())
}