# Idea

Traverse the graph using DFS, create a copy of each node, and use a hash map to keep track of visited nodes.

# Demonstration

graph

            1
          /   \
         2     4
          \   /
            3

DFS

    1                                                           create 1 in the newNodes map; 1 has neighbors (2,4)
    |-----> 2                                                   create 2 in the newNodes map; 2 has neighbors (1,3)
    |       |-----> 1 (return)                                  call 1 but get return immediately due to its presence in the map
    |       |-----> 3                                           create 3 in the newNodes map; 3 has neighbors (2,4)
    |               |-----> 2 (return)                          call 2 but get return immediately due to its presence in the map
    |               |-----> 4                                   create 4 in the newNodes map; 4 has neighbors (1,3)
    |                       |-----> 1 (return)                  call 1 but get return immediately due to its presence in the map
    |                       |-----> 3 (return)                  call 3 but get return immediately due to its presence in the map
    |
    |-----> 4 (return)                                          call 4 but get return immediately due to its presence in the map
