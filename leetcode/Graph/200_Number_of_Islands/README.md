# Additional explanation for the question

This is an island

    0 0 0
    0 1 0
    0 0 0

This is also an island

    1 1 1
    1 0 1
    1 1 1

This is also an island

    0 1 0 1
    1 1 0 1
    1 0 1 1
    1 1 1 0


# Idea

Use DFS to explore the grid, traversing each node with an outer and inner loop.
Identify islands by encountering `1` and recursively call DFS if adjacent elements are `1` until reaching the water `0` or out of the grid range.
Mark visited elements by changing their value to `0` when checking the adjancent elements. (By doing this, we won't count the same island twice, and don't need extra memory for marking visited elements)

# Demonstration

Graph

     0  1  0  0
     1  1  0  0
     0  0  1  0
     0  0  0  1

Encounter 1, the first island, check adjancent elemnts to identify the whole area of the island

     0 [1] 0  0
     1  1  0  0
     0  0  1  0
     0  0  0  1

Mark it as visited element by changing to 0

     0 [0] 0  0
     1  1  0  0
     0  0  1  0
     0  0  0  1

Call DFS for adjancent elements

    [0] 0 [0] 0
     1 [1] 0  0
     0  0  1  0
     0  0  0  1

Found another 1, mark it as visited element

     0  0  0  0
     1 [0] 0  0
     0  0  1  0
     0  0  0  1

Check adjancent elements again

     0 [0] 0  0
    [1] 0 [0] 0
     0 [0] 1  0
     0  0  0  1

Found another 1, mark it as visited element

     0  0  0  0
    [0] 0  0  0
     0  0  1  0
     0  0  0  1

Check adjancent elements

    [0] 0  0  0
     0 [0] 0  0
    [0] 0  1  0
     0  0  0  1

they are all 0, the first island is identified, `island = 1`

     0  0  0  0
     0  0  0  0
     0  0  1  0
     0  0  0  1


Found the next island

     0  0  0  0
     0  0  0  0
     0  0 [1] 0
     0  0  0  1

Mark it as visited element and check adjancent elements

     0  0  0  0
     0  0  0  0
     0  0 [0] 0
     0  0  0  1

     0  0  0  0
     0  0 [0] 0
     0 [0] 0 [0]
     0  0 [0] 1

they are all 0, the second island is identified, `island = 2`

Found the last island

     0  0  0  0
     0  0  0  0
     0  0  0  0
     0  0  0 [1]

Mark it as visited and check adjancent elements

     0  0  0  0
     0  0  0  0
     0  0  0  0
     0  0  0 [0]

     0  0  0  0
     0  0  0  0
     0  0  0 [0]
     0  0 [0] 0

they are all 0, the third island is identified, `island = 3`

All elements have been traversed

     0  0  0  0
     0  0  0  0
     0  0  0  0
     0  0  0  0

end
