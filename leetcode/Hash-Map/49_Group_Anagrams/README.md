# Idea

There are 2 solutions.

1. sort the characters for each string and use the map to group the input O(m * nlogn)
2. count the characters for each string (26 alphabets) O(m * n * 26) => O(m*n)
    * make a string from characters in order of ASCII

> m is the lengh of string slice, n is the number of each string

Note: Adding up the total of code point to group the anagram doesn't work e.g. `duh` and `ill` are the same
