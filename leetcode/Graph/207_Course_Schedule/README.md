# Additional explanation for the question

    Input: numCourses = 4, prerequisites = [[2,0],[1,0],[3,1],[3,2]]
    Output: true
    Explanation:
    There are a total of 4 courses to take. There is a way to take these courses in the following order:
    - Take course 0 (no prerequisite).
    - Take course 1 or 2 (both have a prerequisite of course 0).
    - If course 1 is taken, then take course 2 (or vice versa).
    - Finally, take course 3 (which has prerequisites of course 1 and 2).

    Input: numCourses = 3, prerequisites = [[0,1],[1,2],[2,0]]
    Output: false
    Explanation:
    There is a cycle in the prerequisite structure (0 -> 1 -> 2 -> 0), so it is not possible to finish all courses.

    Input: numCourses = 5, prerequisites = [[1,0],[2,0],[3,1],[4,3]]
    Output: true
    Explanation:
    You can complete courses in the following sequence: 0, 1, 3, 4, and then 2. There are other valid sequences as well.

    Input: numCourses = 3, prerequisites = []
    Output: true
    Explanation:
    There are no prerequisites for any course, so you can take them in any order.

    Input: numCourses = 3, prerequisites = [[2,1],[1,0]]
    Output: true
    Explanation:
    Take the courses in the sequence: 0, 1, and then 2.

# Hint

You are given a directed graph where each node represents a course, and an edge represents a prerequisite relationship between two courses.
You need to determine if it's possible to take all courses without running into a circular dependency of prerequisites.

* DFS - Use DFS to detect cycles in the directed graph. If there's a cycle, it's impossible to finish all courses. If no cycles are detected, then it's possible.
    * might have issues with deep recursion that leads to stack overflow for very large graphs if your platform doesn't support it
* BFS - The idea is to iteratively remove nodes (courses) with zero in-degrees (no prerequisites) until no nodes remain. For each taken course, reduce the in-degree of its dependent courses.
    * If we can reduce the indegree of all courses to 0, it means we can finish all courses.
    * If there are still nodes left at the end (courses with prerequisites), it means there's a cycle.
    * less intuitive for cycle detection than DFS
    * NOTE: In-degree of a node (course): The number of courses that need to be taken before you can take this course.

> Both DFS and BFS solve the problem in O(V+E) time complexity, where V is the number of courses and E is the number of prerequisites.
> Both approaches also use O(V+E) space

# Step-by-Step Approach (DFS)

1. Convert prerequisites into an adjacency list representation of a directed graph.
2. Use an array to track states of each course: 0 for Unvisited, 1 for Visiting, and 2 for Visited.
3. DFS Cycle Detection
    * Implement a DFS function that traverses courses.
    * Mark a course as Visiting when starting and as Visited when done.
    * If a Visiting course is encountered again, a cycle exists.
4. Course Processing
    * Iterate over all courses. Run DFS for unvisited courses.
    * If DFS detects a cycle, return false.
5. Result
    * If all courses are processed without cycles, return true.

# Step-by-Step Approach (BFS)

1. Initialization
    * Convert prerequisites to an adjacency list.
    * Calculate in-degrees for all courses.
2. Initial Queue
    * Initialize a queue with courses having no prerequisites (in-degree of 0).
3. Recursive BFS
    * For the current course:
        1. Process its dependent courses, decrementing their in-degrees.
        2. If a dependent course now has an in-degree of 0, add it to the queue.
    * Recursively call BFS on the remaining courses in the queue.
4. Result:
    * Check if the processed courses count matches numCourses to determine the absence of cycles.

> My understanding: we start from the bottom nodes, that's why we enqueue nodes whose in-degree is 0.
