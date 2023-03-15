# Additional explanation for the question

Just the node itself to be the maximum

        3
       / \
     -3  -4

> output: 3

only 2 nodes (root+right)


        3
       / \
     -3   4

> output: 7 (3->4)

3 nodes (left+root+right)

        3
       / \
      2   4

> output: 9 (2->3->4)

3 nodes (root+right+right)

        3
       / \
     -2   4
           \
            5

> output: 12 (3->4->5)

3+ nodes

        1
       / \
      2   3
           \
            5

> output: 11 (2->1->3->5)

3 nodes

        1
       / \
      2   3
         / \
        4   5

> output: 12 (4->3->5) because the path `2->1->3->4->5` is invalid



# Idea

Use a recursive function that computes the maximum path sum for the left and right subtrees

for eaxample

        1 <- curr
       / \
      2   3
         / \
        4   5

The maximum of left subtree is 2, and the maximum of right subtree is 8 (3+5), then combine them all together, as a result, we got 11 (2 + 1 + 8) at the current node.

Note: When the recursive function goes to `3`, it finds the maximum by combining 3+4 and 3+5 as unsplit path and return 8 back to its parent `1`.

One thing to note is that when the function reaches `3`, it obtains the maximum value of 12 (4 -> 3 -> 5) and save it to the maximum variable before it turns maximum to its parent.







