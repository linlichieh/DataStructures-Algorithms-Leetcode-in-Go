# Idea

* Brute force
* Sliding window
* Binary Search along with prefix sums
    * It's difficult to understand, skip it for now
* (NO) Sort first then subtract num from bottom by target
    * it doesn't work as the problem asks for a contiguous subarray

### Brute force

Use outer and inner loop to iterate each element by adding them up until they are `>=` target.

Repeat the process and update the minimal length of subarray till the end

time complexity: `O(n^2)`

### Sliding window

Use 2 pointers - left and right.
Sum up the nums between left and right.
If the total is less than target, move the right pointer forward; otherwise, move the left forward.
Update the minimal subarray length during the process.
Repeat the process until the end.

time complexity: `O(n)`

space complexity: `O(1)`


# Demonstration

Input

    [2, 3, 1, 2, 4, 3] target: 7

Initialisation

    [2, 3, 1, 2, 4, 3] target: 7
     L
     R

    min = 7 (len(nums) + 1)
    total = 0

R = 0

    [2, 3, 1, 2, 4, 3] target: 7
     L
     R

    min = 7
    total = 2 (0 + 2)

R = 1

    [2, 3, 1, 2, 4, 3] target: 7
     L
        R

    min = 7
    total = 5 (2 + 3)

R = 2

    [2, 3, 1, 2, 4, 3] target: 7
     L
           R

    min = 7
    total = 6 (5 + 1)

R = 3

    [2, 3, 1, 2, 4, 3] target: 7
     L
              R

    min = 7
    total = 8 (6 + 2)

total `>=` target

    [2, 3, 1, 2, 4, 3] target: 7
     L
              R

    min = 4 (updated)
    total = 6 (8 - nums[0])
    left++

    [2, 3, 1, 2, 4, 3] target: 7
        L
              R

R = 4

    [2, 3, 1, 2, 4, 3] target: 7
        L
                 R

    min = 4
    total = 10 (6 + 4)

total `>=` target

    [2, 3, 1, 2, 4, 3] target: 7
        L
                 R

    min = 4
    total = 7 (10 - nums[1])
    left++

    [2, 3, 1, 2, 4, 3] target: 7
           L
                 R

total `>=` target

    [2, 3, 1, 2, 4, 3] target: 7
           L
                 R

    min = 3 (updated)
    total = 6 (7 - nums[2])
    left++

    [2, 3, 1, 2, 4, 3] target: 7
              L
                 R
R = 5

    [2, 3, 1, 2, 4, 3] target: 7
              L
                    R

    min = 3
    total = 9 (6 + 3)

total `>=` target

    [2, 3, 1, 2, 4, 3] target: 7
              L
                    R

    min = 3
    total = 7 (9 - 2)
    left++

    [2, 3, 1, 2, 4, 3] target: 7
                 L
                    R

total `>=` target

    [2, 3, 1, 2, 4, 3] target: 7
                 L
                    R

    min = 2 (updated)
    total = 3 (7 - 4)
    left++

    [2, 3, 1, 2, 4, 3] target: 7
                    L
                    R

end
