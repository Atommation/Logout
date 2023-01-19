package main

import(
  "fmt"
  "time"
)

func loading(ltime int){
  for i:=0; i<=ltime; i++ {
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

func lightning(){
  var list1 [27]int = [27]int{12,12,11,11,10,10,9,9,8,8,7,7,9,8,8,7,7,6,6,5,5,7,7,6,6,5,5}
  var list2 [27]int = [27]int{12,11,11,10,10,9,9,8,8,11,11,10,7,7,6,6,5,5,8,8,7,4,3,3,2,2,1}

  for i:=0; i<27; i++ {
    j := list1[i]
    for j>=1 {
      fmt.Printf(" ")
      j--
    }

    k := list2[i]
    for k>=1 {
      fmt.Printf("*")
      k--
    }

    time.Sleep(time.Second/20)
    fmt.Printf("\n")
  }
  time.Sleep(time.Second*10)
}

func main(){

  time.Sleep(time.Second)
  fmt.Printf("\n\n王某人你真是个小天才")
  time.Sleep(time.Second)
  fmt.Printf("\n\n在学校装躺平回家偷着卷的屑")
  time.Sleep(time.Second)
  fmt.Printf("\n\n平日里小事装友好，考试前找你要笔记你倒不给")
  time.Sleep(time.Second*2)
  fmt.Printf("\n\n\n你就学吧")
  time.Sleep(time.Second)
  fmt.Printf("\n\n越学越聪明是吧")
  time.Sleep(time.Second)
  fmt.Printf("\n\n祝你……………………\n\n\n")
  time.Sleep(time.Second*2)

  loading(4)
  fmt.Printf("\n")

  lightning()
  time.Sleep(time.Second*30)
}
