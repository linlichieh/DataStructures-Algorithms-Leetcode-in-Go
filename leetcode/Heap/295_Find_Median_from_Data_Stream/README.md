# Hint

* Brute force solution
    * Add num into a list and sort it.
    * Loop the list from the beginning, insert the number before the number that is greater
    * The worse case is `O(N)` for each num insertion
* Heap solution
    * min-heap + max-heap
        * It's really tricky. In order get the median numbers in `O(1)`, we need to
            - store the smaller half of the numbers into max-heap
            - store the larger half of the numbers into min-heap
         * Examples: `[3, 7, 6, 5, 1, 4]`
            - max-heap: `[1, 3, 4]`
            - min-heap: `[5, 6, 7]`
            - the median will be `(4 + 5)/2`
                - Get 4 from the maximum of the max-heap
                - Get 5 from the minimum of the min-heap
    * all elements in the max-heap are less than all elements in the min-heap
    * balance the size for both heaps
    * FindMedian is `O(1)`
    * AddNum is `O(log n)`

# Step-by-Step Approach

1. Initialisation
    * A MaxHeap (left half of numbers).
    * A MinHeap (right half of numbers).
2. AddNum Method
    * Insert into left if it's less than or equal to its top or if left is empty; otherwise, insert into right.
    * Balance the heaps:
        * If left has 2 more elements than right, move top of left to right.
        * If right has more elements, move its top to left.
3. FindMedian Method:
    * If both heaps have the same size, return the average of their tops.
    * Otherwise, return the top of left.
