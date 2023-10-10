# Hint

Use a Map to keep track of the unique elements you've encountered

# Step-by-Step Approach

1. Initialize a Map: Use a Map to store unique elements from the array
2. Loop Over Array: Iterate over each element in the array. For each element:
    * Check if Exists: Check if the element already exists in the Set or Map.
    * Add to Set: If it doesn't, add it to the Map.
    * Duplicate Found: If it does, return true (since you found a duplicate).
3. End of Loop: If the loop ends without returning, return false.


