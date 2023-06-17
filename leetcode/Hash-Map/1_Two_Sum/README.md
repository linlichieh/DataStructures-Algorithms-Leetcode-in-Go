# Idea

1. Initialise a map to store the value as key and index as value for each element in the array for quick lookup later
1. Iterate the given array
    1. Calculate the complement by subtracting the value from target
    1. Check if the complement exists in the map
        * If yes, return the current index in the loop and the index of complement, which is the value of complement as key in the map
        * If no, add the value as key and index as value to map

# Demonstration

Input

    [3, 2, 4]   target = 6

Initialisation

    [3, 2, 4]   target = 6

    map[]

i=0

    [3, 2, 4]   target = 6
     *

    complement = 6 - 3
    3 can't be found in the map

    // add value and index to map
    map[3: 0]

i=1

    [3, 2, 4]   target = 6
        *

    complement = 6 - 2
    4 can't be found in the map

    // add value and index to map
    map[3: 0, 2: 1]

i=2

    [3, 2, 4]   target = 6
           *

    complement = 6 - 4
    2 exists in the map

    map[3: 0, 2: 1]

the complement can be found in the map, return both indices

    return [2, 1]

