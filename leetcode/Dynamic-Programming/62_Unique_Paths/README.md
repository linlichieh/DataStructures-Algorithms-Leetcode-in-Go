# Idea

Find the pattern `dp[i][j] = dp[i-1][j] + dp[i][j-1]`

    70  35  15  5  1
    35  20  10  4  1
    15  10   6  3  1
     5   4   3  2  1
     1   1   1  1  1

The `dp[0][0]` is the total unique paths of `(m * n)` grid.

The first approach is to initialise a cache table (m*n) for storing the result of `dp[i][j]`.

Then traverse each element using DFS, the cache can significantly improve the performance.

* Time Complexity: `O(m * n)`
* Space Complexity: `O(m * n)`

Base on above solution, the memory complexity can be improved by just using 1D array for cache.

Initialise all elements of the cache array to 1. This is because there is only one way to reach any cell in the first row and any cell in the first column.

Iterate through the remaining cells of the grid, starting from the second to last row (m - 2) and moving upwards, because the last row has been initialised as default cache.
For each row, start from the last column (n - 2) and move left, because the last column is always 1 and can be skipped.

Then replace each col in the cache with the sum of `dp[j]` and `dp[j+1]`

* Time Complexity: `O(m * n)`
* Space Complexity: `O(n)`

# Demonstration

input: m = 4, n = 3

Initialisation 1D array (cache):

    1 1 1

    you can also think it's the last row like this:

    10 4 1
     6 3 1
     3 2 1
     1 1 1      <--

the second to the last row, and loop column from right to left

    1 [1] 1     1 + 1
    1 [2] 1     = 2

    [1] 2 1     1 + 2
    [3] 2 1     = 3

    3 2 1


    you can also think it's the third row like this:

    10 4 1
     6 3 1
     3 2 1      <--
     1 1 1

then

    3 [2] 1     2 + 1
    3 [3] 1     = 3

    [3] 3 1     3 + 3
    [6] 3 1     = 6

    6 3 1


    you can also think it's the second row like this:

    10 4 1
     6 3 1      <--
     3 2 1
     1 1 1

lastly

    6 [3] 1     3 + 1
    6 [4] 1     = 4

     [6] 4 1    6 + 4
    [10] 4 1    = 10

    10 4 1


    you can also think it's the first row like this:

    10 4 1      <--
     6 3 1
     3 2 1
     1 1 1

end, the final cache looks like

    10 4 1

output `cache[0]`, that is the number of possible unique paths
