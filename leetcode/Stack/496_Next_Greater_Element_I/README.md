# Time complexity

The time complexity is O(N+M)

By iterating nums2, we get `O(M)` complexity, but there is an additional loop of stack `O(N)` inside the nums2 loop. However, since the stack is amortized, the worst-case scenario is `O(N)` at most. This is why we get `O(N+M)` complexity instead of `O(N*M)`.
