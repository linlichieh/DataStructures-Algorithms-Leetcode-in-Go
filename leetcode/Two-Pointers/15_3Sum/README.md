# Solutions

2-pointer + sorting-based solution : This solution involves sorting the input array in non-decreasing order and then using two pointers to find all unique triplets that sum up to zero.

### why is sorting necessary

Without sorting the input array, we would need to check all pairs of elements in the input array, which would result in a time complexity of O(N^2)

Because the input array is sorted, we know that if nums[j] + nums[k] < target, we need to increase j to get closer to the target value, and if nums[j] + nums[k] > target, we need to decrease k to get closer to the target value. We can do this efficiently using the two-pointer approach without checking all pairs of elements in the input array.

### How about hash-based solution?

The array still need to be sorted, the implementation is almost the same as 2 sum but 1 more outer loop.

The main difference betwwen 2-pointer solution and hash-based solution is about space complexity.

2-pointer doesn't need additional space, but hash-based approach does. That's why 2-pointer is the first choice over hash-based approach

2-pointer solution

* Time `O(N^2)` in the worst case, and `O(N log N)` in the average case
* Space `O(1)`

Hash-based solution

* Time `O(N^2)`
* Space `O(N)`
