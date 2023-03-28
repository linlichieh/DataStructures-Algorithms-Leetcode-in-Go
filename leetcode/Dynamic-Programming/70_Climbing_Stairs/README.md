# Idea

 We can break down the problem into subproblems where the solution to the n-th subproblem is the sum of solutions to the (n-1)-th and (n-2)-th subproblems.
 To avoid recomputing the solutions to subproblems, we can store the solutions in an array.

    |-------------------------|
    | n  | num of combination |
    |--------------------------
    | 0  | 0                  |
    |--------------------------
    | 1  | 1                  |
    |--------------------------
    | 2  | 2                  |
    |--------------------------
    | 3  | 3                  |
    |--------------------------
    | 4  | 5                  |
    |--------------------------
    | 5  | 8                  |
    |--------------------------
    | 6  | 13                 |
    |-------------------------|

For example, `n = 6` is the total of `n = 5` + `n = 4`

