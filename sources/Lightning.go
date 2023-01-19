package main

import(
  "fmt"
  "time"
)

func main(){
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
