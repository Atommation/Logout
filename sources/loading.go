package main

import(
  "fmt"
  "time"
)

func main(){
  for {
    for i:=0; i<=60; i++ {
      fmt.Printf("\\")
    }
    time.Sleep(time.Second/10)
    for i:=0; i<=60; i++ {
      fmt.Printf("\b")
    }
    time.Sleep(time.Second/10)
    for i:=0; i<=60; i++ {
      fmt.Printf("-")
    }
    time.Sleep(time.Second/10)
    for i:=0; i<=60; i++ {
      fmt.Printf("\b")
    }
    time.Sleep(time.Second/10)
    for i:=0; i<=60; i++ {
      fmt.Printf("/")
    }
    time.Sleep(time.Second/10)
    for i:=0; i<=60; i++ {
      fmt.Printf("\b")
    }
    time.Sleep(time.Second/10)
    for i:=0; i<=60; i++ {
      fmt.Printf("|")
    }
    time.Sleep(time.Second/10)
    for i:=0; i<=60; i++ {
      fmt.Printf("\b")
    }
    time.Sleep(time.Second/10)
  }
}
