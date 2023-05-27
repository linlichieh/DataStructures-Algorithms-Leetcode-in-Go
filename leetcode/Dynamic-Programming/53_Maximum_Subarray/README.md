# Extended explantion

The purpose of this problem is to find the maximum sum of consecutive order in the array

For example

    Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
    Output: 6

possible continuous subarrays

* ...
* `[4]`: 4
* `[4, -1, 2]`: 5
* `[4, -1, 2, 1]`: 6
* `[4, -1, 2, 1, -5, 4]`: 5
* ...

the maximum one should be `[4, -1, 2, 1]`, the answer is 6

# Idea

Use Kadane's algorithm

* Kadane's algorithm is used to solve the maximum subarray problem.
* Kadane's algorithm finds the maximum subarray sum by iteratively deciding at each position if it's more beneficial to extend the current subarray or start a new one, keeping track of the maximum sum found.

Iterate through the array, maintaining a running sum of the elements.
In each iteration, if at any point the running sum becomes negative, reset it to zero and start accumulating from the next element.
Also, keep track of the maximum sum encountered so far. By the end of the iteration, the maximum sum would be the maximum subarray sum.

# Demonstration

input

    [-1, 5, -2, 4, -7, 1]

initialisation

    sum = 0
    maximum = -1

> Set index 0 `-1` as default value is for the case with all negative values in the array

index = 0 (`-1`)

    sum = -1 (0 - 1)
    maximum = -1

index = 1 (`5`)

    sum = 0  // reset sum because is negative
    sum = 5 (0 + 5)
    maximum = 5

index = 2 (`-2`)

    sum = 3 (5 - 2)
    maximum = 5

index = 3 (`4`)

    sum = 7 (3 + 4)
    maximum = 7

index = 4 (`-7`)

    sum = 0 (7 - 7)
    maximum = 7

index = 5 (`1`)

    sum = 1 (0 + 1)
    maximum = 7

the maximum subarray sum is `7`

