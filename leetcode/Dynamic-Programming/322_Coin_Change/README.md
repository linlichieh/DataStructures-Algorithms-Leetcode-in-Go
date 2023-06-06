# Idea

There are a few approaches to solve this problem

* Bottom-Up Approach (recommended)
* DFS

### Bottom-Up Approach

1. Initialise an array `dp` with a length of `amount + 1`
1. Set default value 0 to the index 0 and set default value `amount + 1` to the rest of the elements
    * `dp[0] = 0`
    * `dp[1] = dp[2] = ... = dp[7] = 8`
1. Loop the `dp` from index 1 as `i`
    * Loop the coins
    * subtract coin from `i` (not `dp[i]`)
    * if abvoe result is `>= 0`, then set `dp[i]` to the smaller number out of `dp[i]` and `1 + dp[i-coin]`
    * after this process, we get to build the base value from the bottom
    * then wel will start to use cache without calculating again
1. Repeat the process during the loop until the end
1. The value of the last element of `dp` should be the fewest number of coins that makes up the amount

### DFS

* Use amount to subtract each coin
    * 7 - 1 = 6
    * 7 - 3 = 4
    * 7 - 4 = 3
    * 7 - 5 = 2
* Each remainder to subtract each coin, for example the last one `2`. stop going deep if the result is less than 0
    * 2 - 1 = 1
        * 1 - 1 = 0 (the number can be made up by coins)
        * 1 - 3 = -2 (stop)
        * 1 - 4 = -3 (stop)
        * 1 - 5 = -4 (stop)
    * 2 - 3 = -1 (stop)
    * 2 - 4 = -2 (stop)
    * 2 - 5 = -3 (stop)
    * save each result in the cache so the it won't be re-calculated again
    * we also want to save the minimum counter to the cache
* Repeat this process until the result is 0, which counter should be the number that makes up the amount

# Demonstration 1

Input

    [1, 3, 4, 5], amount = 7

Initialisation

    dp[0] = 0
    dp[1] = 8
    dp[2] = 8
    dp[3] = 8
    dp[4] = 8
    dp[5] = 8
    dp[6] = 8
    dp[7] = 8

### i = 1 (subtract each coin)

`1 - 1`

    dp[1] = min(dp[1], 1 + dp[0])   // index 0 = i - 1
    dp[1] = 1

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 8
    dp[3] = 8
    dp[4] = 8
    dp[5] = 8
    dp[6] = 8
    dp[7] = 8

`1 - 3`, less than 0, skip

`1 - 4`, less than 0, skip

`1 - 5`, less than 0, skip

### i = 2

`2 - 1`

    dp[2] = min(dp[2], 1 + dp[1])   // index 1 = i - 1
    dp[2] = 2

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 8
    dp[4] = 8
    dp[5] = 8
    dp[6] = 8
    dp[7] = 8

`2 - 3`, less than 0, skip

`2 - 4`, less than 0, skip

`2 - 5`, less than 0, skip

### i = 3

`3 - 1`

    dp[3] = min(dp[3], 1 + dp[2])
    dp[3] = 3

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 3
    dp[4] = 8
    dp[5] = 8
    dp[6] = 8
    dp[7] = 8

`3 - 3`

    dp[3] = min(dp[3], 1 + dp[0])
    dp[3] = 1

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 8
    dp[5] = 8
    dp[6] = 8
    dp[7] = 8

`3 - 4`, less than 0, skip

`3 - 5`, less than 0, skip

### i = 4

`4 - 1`

    dp[4] = min(dp[4], 1 + dp[3])
    dp[4] = 2

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 2
    dp[5] = 8
    dp[6] = 8
    dp[7] = 8

`4 - 3`

    dp[4] = min(dp[4], 1 + dp[1])

same value, no need to update `dp[4]`

`4 - 4`

    dp[4] = min(dp[4], 1 + dp[0])
    dp[4] = 1

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 1
    dp[5] = 8
    dp[6] = 8
    dp[7] = 8

`4 - 5`, less than 0, skip

### i = 5

`5 - 1`

    dp[5] = min(dp[5], 1 + dp[4])
    dp[5] = 2

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 1
    dp[5] = 2
    dp[6] = 8
    dp[7] = 8

`5 - 3`

    dp[5] = min(dp[5], 1 + dp[2])

bigger than `dp[5]`, no need to update `dp[5]`

`5 - 4`

    dp[5] = min(dp[5], 1 + dp[1])

same value, no need to update `dp[5]`

`5 - 5`

    dp[5] = min(dp[5], 1 + dp[0])
    dp[5] = 1

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 1
    dp[5] = 1
    dp[6] = 8
    dp[7] = 8

### i = 6

`6 - 1`

    dp[6] = min(dp[6], 1 + dp[5])
    dp[6] = 2

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 1
    dp[5] = 1
    dp[6] = 2
    dp[7] = 8

`6 - 3`

    dp[6] = min(dp[6], 1 + dp[3])

same value, no need to update `dp[6]`

`6 - 4`

    dp[6] = min(dp[6], 1 + dp[5])

same value, no need to update `dp[6]`

`6 - 5`

    dp[6] = min(dp[6], 1 + dp[1])

same value, no need to update `dp[6]`

### i = 7

`7 - 1`

    dp[7] = min(dp[7], 1 + dp[6])
    dp[7] = 3

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 1
    dp[5] = 1
    dp[6] = 2
    dp[7] = 3

`7 - 3`

    dp[7] = min(dp[7], 1 + dp[4])
    dp[7] = 2

dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 1
    dp[5] = 1
    dp[6] = 2
    dp[7] = 2

`7 - 4`

    dp[7] = min(dp[7], 1 + dp[3])

same value, no need to update `dp[7]`

`7 - 5`

    dp[7] = min(dp[7], 1 + dp[2])

bigger than `dp[7]`, no need to update `dp[7]`

the final dp

    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    dp[3] = 1
    dp[4] = 1
    dp[5] = 1
    dp[6] = 2
    dp[7] = 2

`dp[7]` is 2, which is not the same as default value, it means that it can be made up by coins


# Demonstration 2

Input

    [2, 4], amount = 5

Initialisation

    dp[0] = 0
    dp[1] = 6
    dp[2] = 6
    dp[3] = 6
    dp[4] = 6
    dp[5] = 6

# i = 1

`1 - 2`, less than 0, skip

`1 - 4`, less than 0, skip

# i = 2

`2 - 2`

    dp[2] = min(dp[2], 1 + dp[0])
    dp[2] = 1

dp

    dp[0] = 0
    dp[1] = 6
    dp[2] = 1
    dp[3] = 6
    dp[4] = 6
    dp[5] = 6

`2 - 4`, less than 0, skip

# i = 3

`3 - 2`

    dp[3] = min(dp[3], 1 + dp[1])

bigger than `dp[3]`, no need to update `dp[3]`

`3 - 4`, less than 0, skip

# i = 4

`4 - 2`

    dp[4] = min(dp[4], 1 + dp[2])
    dp[4] = 2

dp

    dp[0] = 0
    dp[1] = 6
    dp[2] = 1
    dp[3] = 6
    dp[4] = 2
    dp[5] = 6

`4 - 4`

    dp[4] = min(dp[4], 1 + dp[0])
    dp[4] = 1

dp

    dp[0] = 0
    dp[1] = 6
    dp[2] = 1
    dp[3] = 6
    dp[4] = 1
    dp[5] = 6

# i = 5

`5 - 2`

    dp[5] = min(dp[5], 1 + dp[3])

bigger than `dp[5]`, no need to update `dp[5]`

`5 - 4`

    dp[5] = min(dp[5], 1 + dp[1])

bigger than `dp[5]`, no need to update `dp[5]`

the final dp

    dp[0] = 0
    dp[1] = 6
    dp[2] = 1
    dp[3] = 6
    dp[4] = 1
    dp[5] = 6

`dp[5]` is 6, which is the same as default value, it indicates that it cannot be made up by any combination of the coins
