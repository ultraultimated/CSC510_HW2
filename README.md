[![Run on Repl.it](https://repl.it/badge/github/ultraultimated/CSC510_HW2)](https://repl.it/github/ultraultimated/CSC510_HW2)
[![Build Status](https://travis-ci.org/ultraultimated/CSC510_HW2.svg?branch=master)](https://travis-ci.org/ultraultimated/CSC510_HW2)

#  Conway's Game of Life in Go, Scala and Ruby

## Introduction
The conway's game of life (known as Life) is a game invented by mathematician John Conway in 1970. The Game of Life is played on an infinite two-dimensional rectangular grid of cells. Each cell can be either alive or dead. The status of each cell changes each turn of the game depending on the statuses of that cell's 8 neighbors. Neighbors of a cell are cells that touch that cell, either horizontal, vertical, or diagonal from that cell.
At each step in time, the following transitions occur:
1) Any live cell with fewer than two live neighbours dies, as if by underpopulation.
2) Any live cell with two or three live neighbours lives on to the next generation.
3) Any live cell with more than three live neighbours dies, as if by overpopulation.
4) Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

<p align="center">
  <img src="https://github.com/ultraultimated/CSC510_HW2/blob/master/GOL.png">
</p>

## Run Instructions:

Click on repl.it badge. Edit .replit file

- Golang

```
language = "go"
run = "go run code/GO/GOL.go"
```

- Ruby

```
language = "ruby"
run = "ruby code/Ruby/GOL.rb"
```

- Scala

```
language = "scala"
run = "scalac -classpath . -d . code/Scala/GOL.scala && scala -classpath . GOL"
```

## Experimental Design:
We have introduced two bugs for each language i.e Go, Ruby and Scala. 
* <b>Compilation Bug</b>: This bug will be more language specific and our intention will be to analyze how the compiler error messages help while debugging syntatctic bugs.
* <b>Logical Bug</b>: This bug will be more specific to implementation and our intention will be to analyze how readable and friendly the language is.

During the experinents, following steps will be performed:
1) Click on the repl badge.
2) To run the program, follow the steps as mentioned [above](#run-instructions).
3) Fill out the post evaluation form [here](https://docs.google.com/forms/d/e/1FAIpQLSf9GUf4OV2xPpL0iZznrvSeS0LzEJndorxmxxVVGwjSoXwShA/viewform?usp=sf_link)


## Evaluation Method:
The evaluation method is a 2 part process. 
1. <b> Post Evaluation form </b> - This form evaluates the familiarity of user with different languages.
2. <b> Manual Observation during debugging </b>
    * Was the user able to find the compilation and logical bug?
    * Time taken to find the bug?
    * Number of times the code was run?

The final evaluation will be done by congregating the results of manual observation and post evaluation form.

## GROUP:
1) Bin Patel (bpatel24)
2) Devarsh Shah (dshah3)
3) Karan Dave (kdave)
4) Neel Parikh (nnparik2)
5) Ritesh Gorse (rghorse)

## Reference for Test Cases:
1) Conway's Game of Life: http://pi.math.cornell.edu/~lipa/mec/lesson6.html 


