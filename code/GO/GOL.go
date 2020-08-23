//Referred awk code from gist.github/timm
//Author:Devarsh Shah 
package main
import (
  "fmt"
  "math/rand"
  "time"
  "os"
)
func life(rows int ,cols int ,some float64 , generations int){
   var length int=rows*cols
   var c =make([]int,length)
   for i:=0;i<length;i++{
      if(rand.Float64()< some){
      c[i]=1}else{
      c[i]=0}
   }
   live(c,rows,generations,length)

}
func live(c []int, rows int, generations int, length int){
  
  if(generations<1){
   os.Exit(0)
   }
   var neighbours int
   fmt.Printf("\033[1;1H")
   fmt.Printf("\033[2J") 
   fmt.Printf("Generation %d\n",generations)
   var count int =0
   for x:=0;x<len(c);x++{
   if(c[x]==1){
   fmt.Print("*")}else{
   fmt.Print(" ")}
   count+=1
   if((count%rows)==0){
   fmt.Println()
   }
   }
   var b =make([]int,length)
   for x:=0;x<len(c);x++{
   neighbours=c[((x-1)%length+length)%length]+c[((x+1)%length+length)%length]+c[((x-rows-1)%length+length)%length]+c[((x-rows)%length+length)%length]+c[((x-rows+1)%length+length)%length]+c[((x+rows-1)%length+length)%length]+c[((x+rows)%length+length)%length]+c[((x+rows+1)%length+length)%length]
    b[x]=c[x]
    if(c[x]==0){
    if(neighbours==3){
    b[x]=1
    }else{
    b[x]=0
    }}else{
    if(neighbours==2 || neighbours==3){
    b[x]=1}else{
    b[x]=0}}
    
    }
    generations--
    time.Sleep(time.Second/2)
    live(b,rows,generations,length)
  }
func main() {
life(50,20,0.3,20)
}