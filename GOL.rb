

def live(now, rows, generations, size)
    if(generations < 1)
        exit 
    end
    print "Generations "+generations.to_s.rjust(3, "0")+"\n"

    for x in 1..(now.length()+1)
        if(now[x] == 1)
            print "*"
        else
            print " "
        end
        if((x%rows) == 0)
            print "\n"
        end
    end
    b = Array.new(size)
    for x in 0..(now.length())
        neighbors = now[((x-1)%size+size)%size]+now[((x+1)%size+size)%size]+
                        now[((x-rows-1)%size+size)%size]+now[((x-rows)%size+size)%size]+
                        now[((x-rows+1)%size+size)%size]+now[((x+rows-1)%size+size)%size]+
                        now[((x+rows)%size+size)%size]+now[((x+rows+1)%size+size)%size]
        b[x] = now[x]
        if(now[x] == 0)
            if(neighbors == 3)
                b[x] = 1
            else
                b[x] = 0
            end
            if(neighbors==2 || neighbors==3)
                b[x] = 1
            else
                b[x] = 0
            end
        end
    end
    generations -= 1
    live(b, rows, generations, size)
end


def life (rows, cols, some, generations)
    size = rows * cols
    now = Array.new(size)
    
    for i in 0..size
        if(rand < some)  
            now[i] = 1
        else
            now[i] = 0
        end
    end
    live(now, rows, generations, size)
end


life(50, 20, 0.619, 20)