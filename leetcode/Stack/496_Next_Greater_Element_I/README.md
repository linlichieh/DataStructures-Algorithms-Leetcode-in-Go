# Idea

This solution works by using the stack to keep track of the numbers that are waiting for their next greater number.
It ensures that whenever a number is pushed onto the stack, all the numbers with their next greater number found are already popped off the stack

* Initialize a map for nums1, where the keys are the values of nums1 and the default value is `-1`.
* Iterate through the nums2 array:
    * Loop the the stack is not empty and the top of the stack is less than the current number
        * pop the stack as key1
        * update the value of map to current number by key1 (the next greater element of the popped number is the current number)
    * Push the current number onto the stack if it's in the list of nums1

# Time complexity

The time complexity is O(N+M)

By iterating nums2, we get `O(M)` complexity, but there is an additional loop of stack `O(N)` inside the nums2 loop. However, since the stack is amortized, the worst-case scenario is `O(N)` at most. This is why we get `O(N+M)` complexity instead of `O(N*M)`.

# Demonstration

nums1 = [2, 0, 1]

nums2 = [2, 0, 3, 4]

Initialise a map based on nums1

    m[2] = -1
    m[0] = -1
    m[1] = -1

Start to iterate through the nums2

n = 2

    stack = []

    2 exists in the nums1, add 2 into the stack


n = 0

    stack = [2], 2 > 0, do nothing

    0 exists in the nums1, add 0 into the stack

n = 3

    stack = [2, 0], top value 0 < 3:
        - pop 0
        - m[0] = 3
    stack = [2], top value 2 < 3:
        - pop 2
        - m[2] = 3

    3 doesn't exist in the nums1, skip

n = 4

    stack = []

    4 doesn't exist in the nums1, skip

nums2 ends, then return the result based on the map of nums1

    m[2] = 3
    m[0] = 3
    m[1] = -1

result `{3, 3, -1}`








