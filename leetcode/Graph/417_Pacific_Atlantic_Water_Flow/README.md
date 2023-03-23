# Idea

Create 2 matrices for Pacific and Atlantic oceans for visited cells.
Traverse the input matrix from each of the four edges (left, right, top, and bottom) and mark the cells as visited that can be reached from the Pacific and Atlantic oceans.
Pass previous height into the DFS function to prevent the direction from going back.
Traverse each cell in the input matrix, and add its coordinates to the resulting array if it exists in both visited matrices.
