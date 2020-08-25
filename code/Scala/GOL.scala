object GOL {
  def main(args: Array[String]): Unit = {
    lifeCustom(6,5,10)
   //lifeRandom(6,5,0.3,10)
  }
  
  // Function to initialize Custom matrix with cells
 def lifeCustom(rows:Int, cols:Int, generation:Int){
    val initialState = Array(0,0,0,0,0,0, 
                             0,1,1,1,0,0, 
                             0,1,0,0,0,0, 
                             0,0,0,0,0,0, 
                             0,0,0,0,0,0)
    live(initialState, rows, cols, generation)
  }
 
  // Function to initialize Random matrix with cells
  def lifeRandom(rows:Int, cols:Int, threshold:Double, generation:Int){
    initialState = Array.fill(rows*cols){0}
    var size: Int = rows*cols;
    for(index <- 0 to size-1){
       if (scala.util.Random.nextFloat < threshold){
        initialState(index) = 1
       }
       else{
       initialState(index)=0
       }
    }
    live(initialState, rows, cols, generation)
  }

  def live(currentState: Array[Int], rows:Int, cols:Int, generation:Int)
  {
   /*
    INPUT: number of rows, number of columns, threshold value, total generation to iterate
    OUTPUT: next stable state of matrix based on 4 rules of Game of life
    */
    if(generation<1){
        System.exit(0)
      }
    //print("\u001b[2J")
    println("\nGeneration: "+generation)
    var cnt: Int = 0;
    var ind: Int = 0;
    for(ind <- 0 to currentState.length-1){
      if(currentState(ind)==1){
        print("*")
      }
      else{
        print(" ")
      }
      cnt=cnt+1
      if (cnt%rows==0){
        println("")
      }
    }
    var nextState = currentState.clone();
    for (ind <- 0 to currentState.length-1){
      var size: Int = rows*cols;
      // scala returns -ve modulo e.g. -7%2 = -1 so below logic will return +ve modulo
      var neighbours: Int = currentState((((ind-1)%size)+size)%size)+
                            currentState((((ind+1)%size)+size)%size)+
                            currentState((((ind-rows-1)%size)+size)%size)+
                            currentState((((ind-rows)%size)+size)%size)+
                            currentState((((ind-rows+1)%size)+size)%size)+
                            currentState((((ind+rows-1)%size)+size)%size)+
                            currentState((((ind+rows)%size)+size)%size)+
                            currentState((((ind+rows+1)%size)+size)%size);
      nextState(ind) = currentState(ind)
      if (currentState(ind)==0){
              if (neighbours==3){
              nextState(ind) = 0;
              }
              else{
              nextState(ind)= 1 ;
              }
      }
      else{
              var temp:int = 0
              if (neighbours==3 || neighbours==2){
              temp=1
              }
              nextState(ind) = temp
      }
  }
  var new_generation: Int = generation-1;
  live(nextState,rows,cols, new_generation)
  }
}
