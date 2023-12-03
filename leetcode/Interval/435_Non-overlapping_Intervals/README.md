# Hint

There are 2 ways to solve this problem:

1. Sorting by start point
    * Sort the intervals based on their starting points and then processing them to minimize overlap
2. Sorting by end point
    * Greedy Algorithm - pick the interval that ends the earliest (i.e., has the smallest end point), because this leaves the most room for the rest of the intervals
    * Similar to the first approach, but sort by end point

Cases to consider:

Case 1 `[1,4] [2,3] [3,5]`, we should remove `[1,4]`

         |-----|---------|
         2     3         5
    |---------------|
    1               4

Case 2 `[1,5] [2,3] [4,7]`, we should remove `[1,5]`

         |-----|    |-------------|
         2     3    4             7
    |--------------------|
    1                    5

Case 3 `[1,4] [2,3] [5,7]`, we should remove `[2,3]`

         |-----|
         2     3
    |---------------|    |----------|
    1               4    5          7

### Sorting by start point VS Sorting by end point

* Sorting by End Point (Preferred Approach)
    * The core idea behind sorting by end point is to prioritize intervals that finish early, thereby leaving more room for subsequent intervals and reducing the likelihood of overlap.
    * By choosing the interval that ends first, you maximize the possibility of accommodating more intervals afterward without overlaps.
    * This approach directly targets the problem's objective: minimizing the number of intervals you need to remove.
* Sorting by Start Point
    * If you sort by start point, you might end up choosing intervals that begin early but last longer, which could overlap with many other intervals.
    * This approach doesn't necessarily minimize the number of intervals to be removed since an early starting interval could overlap with multiple upcoming intervals.
    * For example, consider intervals [1,5] and [2,3]. If you sort by start point, you'd pick [1,5] first, but this overlaps with more intervals than choosing [2,3].

While sorting by start point can work in many scenarios, it's not always the most optimal for minimizing the number of intervals to be removed.
In certain cases, especially where intervals have long durations that overlap with many subsequent intervals, this approach might lead to suboptimal results compared to sorting by end point.
The method of sorting by end point is generally more effective for this problem as it prioritizes keeping intervals that finish early,
thereby leaving more room for subsequent intervals and reducing potential overlaps.


# Step-by-Step Approach

### Sorting by start point

1. Sort Intervals: Arrange the intervals in ascending order based on their start point.
1. Iterate and Check for Overlap: For each interval, compare its start point with the end point of the previously kept interval. If there's an overlap, increment a counter.
1. Choose the Smaller End Point on Overlap: When overlapping occurs, keep the interval with the smaller end point to potentially reduce future overlaps.
1. Count Overlaps: The counter reflects the number of intervals that need to be removed to eliminate all overlaps.
1. Return the Counter: The final count is the minimum number of intervals to remove.

### Sorting by end point

1. Sort the Intervals: First, sort the intervals by their end point in ascending order.
1. Iterate and Select: Keep track of the end point of the last added interval.
    * As you iterate through the sorted intervals, check if the current interval starts after the end point of the last added interval.
    * If it does, it means there's no overlap, and you should update the end point to be the end point of this current interval.
    * If there's an overlap, you'll need to skip this interval (this counts as a removal).
1. Count Removals: Every time you skip an interval due to an overlap, increment a counter. This counter represents the minimum number of intervals that need to be removed.
