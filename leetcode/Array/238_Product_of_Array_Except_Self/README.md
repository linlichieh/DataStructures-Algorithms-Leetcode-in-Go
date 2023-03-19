# Additional explanation for the question

each value in the input array nums will be multipled every single number except for the value itself as return array

    Input: nums = [1,2,3,4]
    Output: [24,12,8,6]

    24 = 2*3*4
    12 = 1*3*4
    8=1*2*4
    6=1*2*3

> division operation isn't allowed, if it's allowed, it would be easier: Multiply all elements in the given array, and then iterate over each element. Divide the cumulative multiplication by the current element in each iteration.

# Idea

Use two loops to calculate the prefix and suffix products for each element in the array.
In the first loop, we calculate the prefix product up to each element, and store it in a separate output array.
In the second loop, we calculate the suffix product from each element, and multiply it with the corresponding prefix product to get the product of all other elements

# Demonstration


first loop to fill prefix multiplication

    input       1   2   3   4
    --------------------------
    output      1   1   2   6

second loop to get postfix multiplication

start from 4, post = 1 (default)

    input       1   2   3   4
    --------------------------
    output      1   1   2   6
                            |----- output[3](=6) * post(=1)

then 3, post = 4

    input       1   2   3   4
    --------------------------
    output      1   1   8   6
                        |--------- output[2](=2) * post(=4)

next 2, post = 12 (4*3)

    input       1   2   3   4
    --------------------------
    output      1   12  8   6
                    |------------- output[1](=1) * post(=12)

last 1, post = 24 (4*3*2)

    input       1   2   3   4
    --------------------------
    output      24  12  8   6
                |----------------- output[0](=1) * post(=24)

end
