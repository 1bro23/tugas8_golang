package main

import "fmt"
import "runtime"
import "time"
import "math/rand"

func main(){
  rand.Seed(time.Now().Unix())
  runtime.GOMAXPROCS(2)

  var message = make(chan string)

  go kirimdata(message)
  terimadata(message)
}

func kirimdata(ch chan<-string){
  for i:=0;true;i++{
    pesan:="Apa Kabar Teman Teman"
    ch<-pesan
    time.Sleep(time.Duration(rand.Int()%10+1)*time.Second)
  }
}

func terimadata(ch<-chan string){
  for{
    select{
    case data:=<-ch:
      fmt.Print(data,"\n")
    case <-time.After(time.Second*5):
      fmt.Println("Timeout, tidak ada aktivitas selama 5 detik")
    }
  }
}
