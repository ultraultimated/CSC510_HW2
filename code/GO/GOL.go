package main

//need to uncomment math/rand for lifeRandom
import (
   "fmt"
   //"math/rand"
   "os"
   "time"
)


//Function to initialize random matrix with cells
/*
func lifeRandom(rows int, cols int, threshold float64, generations int) {
   var length int = rows * cols
   var initialState = make([]int, length)
   for i := 0; i < length; i++ {
      if rand.Float64() < threshold {
         initialState[i] = 1
      } else {
         initialState[i] = 0
      }
   }
   live(initialState, rows, generations, length)

}
*/

//Function to initialize custom matrix with cells

func lifeCustom(rows int, cols int, generations int) {
    initialState := [30]int{
      0, 0, 0, 0, 0,0,
      0, 1, 0, 0, 0,0,
      0, 0, 1, 1, 1,0,
      0, 0, 0, 0, 0,0,
      0, 0, 0, 0, 0,0}
   live(initialState,rows,generations,30)
}

//Function to generate neighbours based on GOL's 4 rules
//Change currentState [30]int to currentState[]int
func live(currentState [30]int, rows int, generations int, length int) {

   if generations < 1 {
      os.Exit(0)
   }
   var neighbours int
   fmt.Printf("\033[1;1H")
   fmt.Printf("\033[2J")
   fmt.Printf("Generation %d\n", generations)
   var count int = 0
   for x := 0; x < len(currentState); x++ {
      if currentState[x] == 1 {
         fmt.Print("*")
      } else {
         fmt.Print(" ")
      }
      count += 1
      if (count % rows) == 0 {
         fmt.Println()
      }
   }
   //Comment var nextState for lifeCustom
   //var nextState = make([]int, length)
   var nextState [30]int
   for x := 0; x < len(currentState); x++ {
      //-7%2=-1 in golang, so ((x-1)%length+length)%length gives -7%2=1
      
      neighbours = currentState[((x-1)%length+length)%length] + currentState[((x+1)%length+length)%length] + currentState[((x-rows-1)%length+length)%length] + currentState[((x-rows)%length+length)%length] + currentState[((x-rows+1)%length+length)%length] + currentState[((x+rows-1)%length+length)%length] + currentState[((x+rows)%length+length)%length] + currentState[((x+rows+1)%length+length)%length]
      nextState[x] = currentState[x]
      if currentState[x] == 0 {
         if neighbours != 3 {
            nextState[x] = 0
         } else 
         {
            nextState[x] = 1
         }
      } else 
      {
         if neighbours == 3 || neighbours == 2 {
            nextState[x] = 1
         } else 
         {
            nextState[x] = 0
         }
      }

   }
   generations--
   time.Sleep(time.Second / 2)
   live(nextState, rows, generations, length)
}

//Main function 
func main() {
   //lifeRandom(50, 20, 0.3, 20)
   lifeCustom(6,5,20)
}
