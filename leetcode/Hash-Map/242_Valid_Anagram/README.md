# Hint

* solution 1: Use a frequency counter to count the occurrences of each character in both strings, and then comparing these counts.
* solution 2: Use a frequency counter to count the occurrences of each character in the first string.
    * Then subtract the occurrences of each character while iterating through the second string.
    * If all occurrences are 0 after the process, it indicates that the second string is the anagram of the first string.

# Step-by-Step Approach

1. Create Map(s): Create a map to count the occurrences of characters (runes).
    * If you create one map, increment the count for characters in the first string and decrement for characters in the second string.
    * If counts are all zero at the end, the strings are anagrams.
    * Alternatively, you can create two separate maps for both strings and compare them at the end.
2. Iterate Over Strings: Use for loop with range to iterate over each string's characters, updating the counts in the map(s).
3. Compare Maps: If you used two maps, compare them to ensure they have the same counts for each character. If you used one map, ensure all counts are zero.
