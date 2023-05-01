# Solutions comparison

Candidate solutions

* Brute Force
* Compare one by one
* Priority Queue (min-heap)
* Merge lists one by one
* Merge with Divide And Conquer

Symbol explanation

* `n` = average size of each list
* `k` =  the number of linked lists
* `N` = `n` * `k` (the total number of nodes)

### Solution - Brute Force

* Algorithm
    * combining all the linked lists into a single list and then sorting the merged list
* Time
    * `O(NlogN)` = `O(nk log nk)`
        * traversing all the linked lists and collect the values of the nodes into an array `O(N)`
        * Sorting `O(N log N)`
        * total: `O(N)` + `O(N log N)` = `O(N log N)`
* Space
    * `O(N)`
        * create a new linked list to store the merged and sorted list `O(N)`

### Solution - Compare one by one

* Algorithm
    * Iterate `k` sorted lists, then compare every head of each list and get the node with the smallest value and put it into the final linked list
* Time
    * `O(kN)` = `O(nk^2)`
        * Explanation 1
            * Almost every node in final linked costs `O(k)` (`k-1` times comparison)
            * There are `N` nodes in the final linked list.
            * Therefore, it costs `O(kN)` in total
        * Explanation 2
            * The worst case occurs when all the input lists have a length of `n`, and we need to iterate over all `k` lists in each iteration to find the minimum element.
            * In each iteration, we need to compare the first element of each list to find the minimum element, which takes `O(k)` time. Since we need to perform this operation `nk` times to construct the merged list, the total time complexity is `O(nk^2)`.
* Space
    * `O(1)`
        * We only need a constant amount of extra space to store the merged list. We do not need to create any new data structures, as we can modify the pointers of the existing nodes to create the merged list.
    * `O(N)`
        * Creating a new linked list costs `O(N)` space.

### Solution - Priority Queue (min-heap)

* Algorithm
    * Add the first `k` nodes from each lists
    * Iterate k sorted lists
        * Extract the minimum node from the queue (the node with the smallest value)
        * Append the extracted node to the merged list
        * If the extracted node has a next node in its original list, add the next node to the queue
        * Until the heap is empty
* Time
    * `O(N log k)` = `O(nk log k)`
        * add the first node of each list to the heap in `O(k)` time
        * then do iteration
            * extract the minimum node and add the next node from the same list to the heap in `O(logk)` time
            * repeat above for each of the `N` nodes
        * total: `O(k)` + `N` * `O(logk)` = `O(N log k)`
        * Note: we only have `k` elements in the heap at any given time, because in each iteration we only add 1 element into heap
* Space
    * `O(k)`
        * We only need a priority queue to store the `k` nodes, one from each of the `k` lists, at any given time.
    * `O(N)`
        * Creating a new linked list costs `O(N)` space.

### Solution - Merge lists one by one

* Algorithm
    * merging two lists at a time until we merge all `k` lists into one sorted list
* Time
    * `O(kN)` = `O(nk^2)`
        * All lists have the same size `n`. At the first step of the algorithm we merge 2 lists with `O(n)` and get a list with size `2n`.
        * At the second step we merge a `2n` list with a `n` list, in the worse case we have to visit `2n` nodes and get a list with size `3n`.
        * And so on. At the end we have `n + 2n + ... + kn = n(1+2 + .. + k) = n * ((k + 1)*k) / 2` = `n * (k^2 + k)/2` = `nk^2/2` + `nk/2` = `nk^2`
            * e.g. 1+2+3+4+5 = ((1+5) * 5) / 2
        * So in average we have here `O(nk^2)`.
* Space
    * `O(1)`
        * modify the pointers of the existing nodes to create the merged list, without creating any new data structures.
    * `O(N)`
        * Creating a new linked list costs `O(N)` space.

### Solution - Merge with Divide And Conquer

* Algorithm
    * recursively merging the linked lists in a divide-and-conquer manner, until we merge all `k` lists into one sorted list
    * The idea behind the solution is to use a divide-and-conquer approach to merge the `k` sorted lists.  We first divide the `k` lists into two halves, recursively sort each half, and then merge the two halves.  This process is repeated until we have a single merged list.
* Time
    * `O(N log k)` = `O(nk log k)`
        * We divide the `k` lists into half at each level. Then we will have `O(logk)` levels in total.
        * Each level takes `O(nk)` time in total
        * Since we have `O(logk)` levels and each level take `O(nk)` time, the total time complexity is `O(nk logk)`.
* Space
    * O(1)
        * We can merge two sorted linked lists in `O(1)` space.

The demonstration of time complexity

    assuming, k = 8, n = 4

                n n n n n n n n n
                 /             \
            n n n n           n n n n     => each pair take O(8n), each level takes O(8n) = O(kn)
             /   \             /   \
          n n     n n       n n     n n   => each pair take O(4n), each level takes O(8n) = O(kn)
          / \     / \       / \     / \
         n   n   n   n     n   n   n   n  => each pair takes O(2n), each level takes O(8n) = O(kn)

> The total time complexity is `O(kn) + O(kn) + O(kn)` = `O(kn) * the number of levels` = `O(kn) * O(log k)` = `O(nk log k)`


### Observation in terms of performance on LeetCode

* Undoubtedly, the Brute Force method is the slowest.
* There is no significant difference between `Merge with Divide And Conquer` and `Priority Queue`.


> * chatGPT
> * https://leetcode.com/problems/merge-k-sorted-lists/editorial/
