# Idea

Implement the 2-pointer approach: one pointer for the base element to let the other pointer (j) compare with the element two positions before i.
If they are equal, then j is the value to be skipped. Otherwise, swap the two values and move both i and j forward until reaching the end of the array.

# Steps

Input

    1 1 1 2 2 3

Init

        j
    1 1 1 2 2 3
        i

arr[j] == arr[i-2], then j+1

          j
    1 1 1 2 2 3
        i

arr[j] != arr[i-2], swap i and j; also, move both i and j forward

          j
    1 1 2 1 2 3
        i

            j
    1 1 2 1 2 3
          i

arr[j] != arr[i-2], again

            j
    1 1 2 2 1 3
          i

              j
    1 1 2 2 1 3
            i

arr[j] != arr[i-2], again

              j
    1 1 2 2 3 1
            i

                j
    1 1 2 2 3 1
              i

loop ends


