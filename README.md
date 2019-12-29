# skiing

## question

[original question](http://geeks.redmart.com/2015/01/07/skiing-in-singapore-a-coding-diversion/)

Given the elevation map of a mountain, find the longest path with the highest elevation. You can start from any cell and move up, down, left or right but you can only move to a cell of lower elevation.

Example input  
```
4 8 7 3  
2 5 9 3  
6 3 2 5  
4 4 1 6  
```

Example output  
```
length: 5  
height: 8
```

Explanation  
There are 2 longest path, 8->5->3->2->1 (elevation of 7) and 9->5->3->2->1 (elevation of 8).

## solution

Create 2 tables, length and height.
```
length  

0 0 0 0  
0 0 0 0  
0 0 0 0  
0 0 0 0  
```
```
height  

4 8 7 3  
2 5 9 3  
6 3 2 5  
4 4 1 6  
```

* Each cell in length is the longest path length starting from this cell. It is 1 + the maximum value of its valid neighbors. A neighbor is valid if it is of lower elevation.  
* Each cell in height tracks the lowest possible elevation for the longest path.  
* Whenever a cell is updated, its neighbors values might need to be updated again if the neighbors are on a higher elevation, so we use a queue to store these neighbors.  
* The solution will be the highest value in the length table and its corresponding lowest value in the height table.

example end tables
```
length  

2 5 2 1  
1 4 5 1  
4 3 2 3  
1 4 1 4  
```
```
height  

2 1 3 3  
2 1 1 3  
1 1 1 1  
4 1 1 1  
```

Thus the answer is from cell {1,2} with length 5 and elevation 9-1=8.

## run program

### if go is installed

```
go run ./cmd/main.go
```

## if docker is installed

```
docker build . -t skiing
docker run --rm skiing
```
