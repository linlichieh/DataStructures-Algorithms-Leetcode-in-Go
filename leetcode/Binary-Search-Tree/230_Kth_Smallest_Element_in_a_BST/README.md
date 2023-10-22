# Hint

Perform an in-order traversal of the BST.
While traversing, keep track of the number of nodes you've visited. When the count reaches 'k', you've found your kth smallest element.

Follow-up question: (I didn't implement this)
- Maintain a Size Field: For each node in the BST, maintain the count of nodes in its left subtree. This size attribute will help you decide at any node whether the kth smallest is on the left subtree, the current node, or the right subtree.
    * If the left subtree has 'k-1' nodes, then the current node is the kth smallest.
    * If the left subtree has more than 'k' nodes, then the kth smallest is in the left subtree.
    * If the left subtree has 'l' nodes where 'l' is less than 'k-1', then the kth smallest is the (k-l-1)th smallest node in the right subtree.
- Balanced BST: If insertions and deletions are frequent, ensuring that the BST is balanced (like an AVL tree or Red-Black tree) will be beneficial for maintaining efficiency.
- Caching: If you've recently queried the kth smallest, consider caching recent lookups. If the tree hasn't been modified since the last query, you can return the result from the cache.

# Step-by-Step Approach

* Initialisation
    * initialise `count` for tracking of the Nth smallest values
    * initialise `result` for storing kth smallest value
* in-order traversal
    * recursively call itself with left node
    * count++
        * if counter equals to k, assign node's value to `result` and do early return
    * recursively call itself with right node
* After in-order traversal, return `result`
