package main

import "fmt"
import "time"

type Ball struct{ hits int }

func main() {
  table := make(chan *Ball)
  go player("ping", table)
  go player("pong", table)

//  table <- new(Ball)  // game on; toss the ball
  time.Sleep(10 * time.Second)
//  <- table // game over; grab the ball
}

func player(name string, table chan *Ball) {
  for {
    fmt.Println(name, "is waiting a ball...")
    ball := <-table     // waiting until a Ball is happened
    fmt.Println(name, "get ball", ball)
//    ball.hits ++
//    fmt.Println(name, ball.hits)
//    time.Sleep(100 * time.Millisecond)
//    table <- ball
  }
}
