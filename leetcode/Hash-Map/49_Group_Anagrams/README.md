# Idea

There are 2 solutions.

1. sort the characters for each string and use the map to group the input O(m * nlogn)
2. count the characters for each string (26 alphabets) O(m * n * 26) => O(m*n)

> m is the lengh of string slice, n is the number of each string

The solution 2 is adopted by my implementation due to the time complexity
