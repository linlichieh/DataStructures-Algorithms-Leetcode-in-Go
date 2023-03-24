# Idea

The solution can be solved by both of BFS and DFS.

* In DFS, swapping the left and right subtrees recursively while traversing all nodes is easy to implement and is straighforward.
* In BFS, it requires more memory to maintain the queue of nodes to visit.

The solution is based on postorder traversal, but the key difference is swapping the left and right nodes before returning the root node.
