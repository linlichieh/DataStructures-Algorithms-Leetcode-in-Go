# Idea

### DFS

Using the DFS approach can solve this problem for small grids.
However, this solution could be extremely slow when the grid is huge.
It will take a long time to solve a big grid because it still goes through the cells that have been visited before.

### 2D Cache

Given the patten below, you can calculate a cell by adding up its below and right cells
, and to avoid the cells that have been visited before by storing each calculation into a 2D cache

Example 1: grid (no obstacle)

        0 0 0
        0 0 0
        0 0 0

calculatetotal number of paths

        6 3 1
        3 2 1
        1 1 1

> `c[0][0] = 2` is the asnwer

Example 2: grid (has obstacle)

    grid

        0 0 0
        0 1 0
        0 0 0

    total number of paths

        2 1 1
        1 0 1           <- treat obstacle as 0 if encountered
        1 1 1

> `c[0][0] = 2` is the asnwer

The formula is `grid[i][j] = grid[i-1][j] + grid[i][j-1]`

Note that, to avoid recalculating the result, use a cache grid to store the computed values

Time: O(Row*Col)

Space: O(Row*Col)


## 1D Cache (Recommended)

To improve the space complexity of 2D Cache solution, use only 1D Cache by calculating these cells bottom-up

Time: O(Row*Col)

Space: O(Col)


# Demonstration of 1D Cache solution

grid

     0  0  0  0
     0  0  1  0
     0  0  0  0 <-

initialise an array as the last row, and the last column should be `1` as default

     0 0 0 1

### row(index: 2)

grid

     0  0  0  0
     0  0  1  0
     0  0  0 [0]

cache

     0  0  0 [1]

skip, as it's the last column

grid

     0  0  0  0
     0  0  1  0
     0  0 [0] 0

cache

     0  0 [0] 1

cache[j] += cache[j+1]

     0  0 [1] 1

grid

     0  0  0  0
     0  0  1  0
     0 [0] 0  0

cache

     0 [0] 1  1

cache[j] += cache[j+1]

     0 [1] 1  1

grid

     0  0  0  0
     0  0  1  0
    [0] 0  0  0

cache

    [0] 1  1  1

cache[j] += cache[j+1]

    [1] 1  1  1

### row(index: 1)

grid

     0  0  0  0
     0  0  1 [0]
     0  0  0  0

cache

     1  1  1 [1]

skip, as it's the last column

grid

     0  0  0  0
     0  0 [1] 0
     0  0  0  0

cache

     1  1 [1] 1

change to `0` as it's obstacle

     1  1 [0] 1

grid

     0  0  0  0
     0 [0] 1  0
     0  0  0  0

cache

     1 [1] 0  1

cache[j] += cache[j+1]

     1 [1] 0  1

grid

     0  0  0  0
    [0] 0  1  0
     0  0  0  0

cache

    [1] 1  0  1

cache[j] += cache[j+1]

    [2] 1  0  1


### row(index: 0)

grid

     0  0  0 [0]
     0  0  1  0
     0  0  0  0

cache

     2  1  0 [1]

skip, as it's the last column

grid

     0  0 [0] 0
     0  0  1  0
     0  0  0  0

cache

     2  1 [0] 1

cache[j] += cache[j+1]

     2  1 [1] 1

grid

     0 [0] 0  0
     0  0  1  0
     0  0  0  0

cache

     2 [1] 1  1

cache[j] += cache[j+1]

     2 [2] 1  1

grid

    [0] 0  0  0
     0  0  1  0
     0  0  0  0

cache

    [2] 2  1  1

cache[j] += cache[j+1]

    [4] 2  1  1

end

### result

cache

    4  2  1  1

the result is the first cell of the cache array, which is `4`












