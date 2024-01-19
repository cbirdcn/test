package main

import(
"fmt"
)

func f2() (r int) {
    t := 5
   defer func() {
    t = t + 5
    fmt.Println(t)
   }()

   defer func() {
    t = t + 1
    }()
   return t
}

func main(){
	fmt.Println(f2())
}
