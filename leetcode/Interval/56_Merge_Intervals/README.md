# Additional explanation for the question

The first element represents the start of the interval and the second element represents the end of the interval.

* Interval [1,3] starts at 1 and ends at 3.
* Interval [2,6] starts at 2 and ends at 6.
* Interval [8,10] starts at 8 and ends at 10.
* Interval [15,18] starts at 15 and ends at 18.

The interval [1,3] and [2,6] overlap with each other because 2 (from the second interval) lies within the first interval.
In this case, these two intervals can be merged into one interval [1,6].

# Idea

1. Sort the intervals based on the starting value.
2. Initialize a merged list with the first interval.
3. Iterate through the intervals, starting from the second interval
    1. If the current interval's starting value is less than or equal to the ending value of the last interval in the merged list, it means they overlap. In this case, merge them.
    2. If they don't overlap, simply add the current interval to the merged list.
