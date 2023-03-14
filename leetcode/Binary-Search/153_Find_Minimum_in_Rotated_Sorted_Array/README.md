# Idea

To find the pivot point, check if the middle element is greater than the last element.
If it is, the pivot point must be to the right of the middle element.
Otherwise, it must be to the left of the middle element.

List all the case by making comparisons between mid, left and right

    1 2 3 4 5       mid > left, go left         mid < right, go left
    l   m   r

    5 1 2 3 4       mid < left, go left         mid < right, go left
    l   m   r

    4 5 1 2 3       mid = left, go left/right   mid = rihgt, go left/right
    l   m   r

    3 4 5 1 2       mid > left,  go right       mid > right, go right
    l   m   r

    2 3 4 5 1       mid > left,  go right       mid > right, go right
    l   m   r

Summary

* if mid is greater than left, it would go both right and left.
* if mid is greater than right, it can only go one way which means that mid and right should be used to compare and decide which way to go.
