package main

import (
    "fmt"
)

type qsintslice struct {
    a []int      // the actual slice to sort
    c chan int   // the back- reporting channel
}

const quit = 1

func qsworker(admin chan int, pipeline chan qsintslice) int {
    bufferi := 0
    buffers := 10
    buffer := make([]qsintslice,buffers)
    for {
        if bufferi > buffers-2 { // we need to be able to put in 2 more
            tmp := make([]qsintslice,10)
            buffer = append(buffer,tmp...)
            buffers += 10
        }
        var s qsintslice
        var sort bool = false
        var tmpb bool = true
        
        // first we try to empty the local buffer
        //   => but the respective "case" would be undefined (no value to send), when the buffer is empty
        //      => potential additional select afterwards is also a must for when bufferi == 0
        for tmpb && bufferi > 0 {
            select {
                case msg1 := <-admin:
                    fmt.Println("  getting msg from admin channel (A): ",msg1)
                    if msg1 == quit {
                        fmt.Println("  worker shutting down ...")
                        return 0
                    }
                case pipeline <- buffer[bufferi-1]:
                    bufferi--
                case s = <-pipeline:
                    //fmt.Println("  getting slice from pipeline channel (A) - len: ",len(s.a))
                    sort = true
                    tmpb = false
            }
        }
        if !sort {
            select {
                case msg1 := <-admin:
                    fmt.Println("  getting msg from admin channel (B): ",msg1)
                    if msg1 == quit {
                        fmt.Println("  worker shutting down ...")
                        return 0
                    }
                case s = <-pipeline:
                    //fmt.Println("  getting slice from pipeline channel (B) - len: ",len(s.a))
                    sort = true
            }
        }
        if sort {
            n := len(s.a)
            if n < 2 {
                s.c <- -1
            } else if n < 10 {
                // bubble sort with optimization from wikipedia
                for n > 0 {
                    newn := 0
                    for i := 1 ; i < n; i++ {
                        if s.a[i-1] > s.a[i] {
                            swap := s.a[i-1]
                            s.a[i-1] = s.a[i]
                            s.a[i] = swap
                            newn = i
                        }
                    }
                    n = newn
                }
                s.c <- -1
            } else {
                // TODO: additional branch for n < 100 (or other value), were we complete qs internally without additional placment to the pipeline
                // quicksort from wikipedia
                i := 0
                j := n - 2
                nm1 := n-1
                pivot := s.a[nm1]
                for i < j {
                    for s.a[i] < pivot && i < nm1 {
                        i++
                    }
                    for s.a[j] >= pivot && j > 0 {
                        j--
                    }
                    if i < j {
                        swap := s.a[i]
                        s.a[i] = s.a[j]
                        s.a[j] = swap
                    }
                }
                if s.a[i] > pivot {
                    swap := s.a[i]
                    s.a[i] = s.a[nm1]
                    s.a[nm1] = swap
                }

                divider := i
                if divider > 1 {
                    s.c <- 1
                    buffer[bufferi] = qsintslice{s.a[0:divider],s.c}
                    bufferi++
                }
                if n > divider + 2 {
                    s.c <- 1
                    buffer[bufferi] = qsintslice{s.a[divider+1:n],s.c}
                    bufferi++
                }
                s.c <- -1 
            }
        }
    }
}

func qs(a []int, pipeline chan qsintslice) int {
   c := make(chan int, 20)
   pipeline <- qsintslice{a,c}
   cnt := 1
   for cnt > 0 {
       cnt += <-c
   }

   return 0
}


var admin chan int 
var pipeline chan qsintslice 

func start_qspipeline() {
	admin = make(chan int, 10)
	pipeline = make(chan qsintslice, 100)
    go qsworker(admin, pipeline)
    go qsworker(admin, pipeline)
    go qsworker(admin, pipeline)
    go qsworker(admin, pipeline)
}

func stop_qspipeline() {
    admin <- quit
    admin <- quit
    admin <- quit
    admin <- quit
}
