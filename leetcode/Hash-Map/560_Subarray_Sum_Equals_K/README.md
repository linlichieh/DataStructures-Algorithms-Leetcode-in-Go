# Hint

Use a hash to track the cumulative sum and its frequency

* A cumulative sum is a sequence of partial sums of a given sequence. For example, for the sequence [1, 2, 3], the cumulative sum would be [1, 3, 6].
* In this problem, we're interested in finding subarrays whose sum equals k. One way to do this is by calculating the cumulative sum up to each index and then checking if there's a previous cumulative sum such that currentSum - previousSum = k.

time complexity: O(n)

# Step-by-Step Approach

k=3, nums:

    [1, -1, 2, -1, 2, -3, 3]

Initialisation

    sum = 0
    result = 0
    sumFreq = {0: 1}

> to handle edge case, 0's default is 1

index 0

         [1, -1, 2, -1, 2, -3, 3]
    sum:  1
    diff: -2                            // (1-k), can't find it in the sumFreq, so do nothing
    sumFreq: {0: 1, 1: 1}               // add key 1

index 1

         [1, -1, 2, -1, 2, -3, 3]
    sum:  1   0
    diff: 3                             // can't find it in the sumFreq, so do nothing
    sumFreq: {0: 2, 1: 1}               // sum is 0, key's count + 1

index 2

         [1, -1, 2, -1, 2, -3, 3]
    sum:  1   0  2
    diff: -1                            // can't find it in the sumFreq, so do nothing
    sumFreq: {0: 2, 1: 1, 2: 1}         // add key 2

index 3

         [1, -1, 2, -1, 2, -3, 3]
    sum:  1   0  2   1
    diff: -2                            // can't find it in the sumFreq, so do nothing
    sumFreq: {0: 2, 1: 2, 2: 1}         // sum is 1, key's count + 1

index 4

         [1, -1, 2, -1, 2, -3, 3]
    sum:  1   0  2   1  3
    diff: 0                             // find it in the sumFreq, count += 2
    sumFreq: {0: 2, 1: 2, 2: 1, 3: 1}   // add key 3
    count: 2

index 5

         [1, -1, 2, -1, 2, -3, 3]
    sum:  1   0  2   1  3   0
    diff: -3                            // can't find it in the sumFreq, so do nothing
    sumFreq: {0: 3, 1: 2, 2: 1, 3: 1}   // sum is 0, key's count + 1
    count: 2

index 6

         [1, -1, 2, -1, 2, -3, 3]
    sum:  1   0  2   1  3   0  3
    diff: 0                             // find it in the sumFreq, count += 3
    sumFreq: {0: 3, 1: 2, 2: 1, 3: 2}   // sum is 3, key's count + 1
    count: 5

Result is 5
