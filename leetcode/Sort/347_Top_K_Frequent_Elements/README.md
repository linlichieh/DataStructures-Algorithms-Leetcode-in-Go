# Idea

### sorting solution `O(nlogn)`

1. Hash map for counting the frequency of each element `O(n)`
1. Sort the frequency `O(nlogn)`
1. Get K most frequent elements `O(k)`

> Due to the fact that the algorithm's time complexity must be better than `O(n log n)`, so solution 1 is out.

### max-heap solution `O(klogn)`

1. Hash map for counting the frequency of each element `O(n)`
1. Build a max-heap `O(n)`
1. Pop k elements `O(klogn)`

> better than min-heap (assuming that n is much larger than k)

### min-heap solution `O(nlogk)`

1. Hash map for counting the frequency of each element `O(n)`
1. Build a min-heap, and add each element into the heap and remove the smallest if heap > k. Once done, only top k most frequent elements will remain in the heap. `O((n-k)logk) = O(nlogk)`
1. Pop k elements `O(k)`

### bucket solution `O(n)`

1. Hash map for counting the frequency of each element `O(n)`
1. Create buckets with a length of n+1, each bucket will contain all elements that have the same frequency. Iterate over the hash map and add them to the corresponding bucket. `O(n)`
1. Starting from the end of the bucket list, iterate over each bucket and append its elements to the result list until k elements have been added. `O(n)` (in the worst case)

cons

* use more memory
