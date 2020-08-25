def live(currentState, rows, generations, length)
    if(generations < 1)
        exit 
    end
    print "\nGenerations "+generations.to_s.rjust(3, "0")+"\n"

    count = 0

    for x in 0..(currentState.length)
        if(currentState[x] == 1)
            print "*"
        else
            print " "
        end
      
        count++
        if((count % rows) == 0)
            print "\n"
        end
    end
    nextState = Array.new(length)
    for x in 0..(currentState.length())
        neighbors = currentState[((x-1)%length+length)%length]+currentState[((x+1)%length+length)%length]+
                        currentState[((x-rows-1)%length+length)%length]+currentState[((x-rows)%length+length)%length]+
                        currentState[((x-rows+1)%length+length)%length]+currentState[((x+rows-1)%length+length)%length]+
                        currentState[((x+rows)%length+length)%length]+currentState[((x+rows+1)%length+length)%length]
        nextState[x] = currentState[x]
        if(currentState[x] == 0)
            if(neighbors == 3)
                nextState[x] = 1
            else
                nextState[x] = 0
            end
        else
            if(neighbors==2 || neighbors==3)
                nextState[x] = 1
            else
                nextState[x] = 0
            end
        end
    end
    generations += 1
    sleep 2
    live(nextState, rows, generations, length)
end


def lifeRandom (rows, cols, threshold, generations)
    length = rows * cols
    initialState = Array.new(length)
    
    for i in 0..length
        if(rand < threshold)  
            initialState[i] = 1
        else
            initialState[i] = 0
        end
    end
    live(initialState, rows, generations, length)
end

def lifeCustom (rows, cols, generations)
    initialState = [0, 0, 0, 0, 0, 0,
                    0 ,1, 0, 0, 0, 0,
                    0, 0, 1, 1, 1, 0,
                    0, 0, 0, 0, 0, 0,
                    0, 0, 0, 0, 0, 0]
    live(initialState, rows, generations)
end


lifeRandom(50, 20, 0.3, 20)
# lifeCustom(6, 5, 20)
