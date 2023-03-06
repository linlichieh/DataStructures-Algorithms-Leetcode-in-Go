# Idea

1. Use a sliding window with two pointers, one to keep track of the start of the window and another to keep track of the end of the window.
1. Initialize both pointers to 0. Loop through the input string using the second pointer, and at each position, check if the current character has already been seen in the window.
1. If it has, move the start of the window to the right until the duplicate character is no longer in the window.
1. Update the length of the longest substring seen so far at each iteration.
1. Continue looping through the string until the second pointer reaches the end of the string.


# Demonstration

input

    a b c c c d a b

Initialise

    j
    a b c c c d a b
    i
    map: {}
    longest: 0

Loop 1

      j
    a b c c c d a b
    i
    map: {a:true}
    longest: 1

Loop 2

        j
    a b c c c d a b
    i
    map: {a:true, b:true}
    longest: 2

Loop 3

          j
    a b c c c d a b
    i
    map: {a:true, b:true, c: true}
    longest: 3

Loop 4

          j                             // found duplicate `c` in the map
    a b c c c d a b
      i                                 // i++
    map: {b:true, c: true}              // remove a
    longest: 3

          j                             // found duplicate `c` in the map
    a b c c c d a b
        i                               // i++
    map: {c: true}                      // remove b
    longest: 3

          j                             // found duplicate `c` in the map
    a b c c c d a b
          i                             // i++
    map: {}                             // remove c
    longest: 3

            j
    a b c c c d a b
          i
    map: {c:true}
    longest: 3

Loop 5

            j                           // found duplicate `c` in the map
    a b c c c d a b
            i                           // i++
    map: {}                             // remove c
    longest: 3

              j
    a b c c c d a b
            i
    map: {c:true}
    longest: 3

Loop 6

                j
    a b c c c d a b
            i
    map: {c:true, d:true}
    longest: 3

Loop 7

                  j
    a b c c c d a b
            i
    map: {c:true, d:true, a:true}
    longest: 3

Loop 8

                   j
    a b c c c d a b
            i
    map: {c:true, d:true, a:true, b:true}
    longest: 4

Loop ends
