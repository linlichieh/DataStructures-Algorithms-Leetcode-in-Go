# Idea

The main idea is to subtract the frequency of the most common characters from the window length and then check if it satisfies the given rule.
If it doesn't, move foward the window.

# Demonstration 1

Example

* input: `A B A C D E F`
* k = 2
* output: 4

char A

    r
    A B A C D E F
    l

    map: {A:1}
    mostFreq: 1
    longest: A

char B

      r
    A B A C D E F
    l

    map: {A:1, B:1}
    mostFreq: 1
    longest: AB

char A

        r
    A B A C D E F
    l

    map: {A:2, B:1}
    mostFreq: 2
    longest: ABA

char C

          r
    A B A C D E F
    l

    map: {A:2, B:1, C:1}
    mostFreq: 2
    longest: ABAC

char D

            r
    A B A C D E F
    l

    map: {A:2, B:1, C:1, D:1}
    mostFreq: 2
    longest: ABAC

length (5) - mostFreq (2) > k (2), moving the left forward

            r
    A B A C D E F
      l

    map: {A:1, B:1, C:1, D:1}
    mostFreq: 2                         // still 2, don't need to update here
    longest: ABAC

char E

              r
    A B A C D E F
      l

    map: {A:1, B:1, C:1, D:1, E:1}
    mostFreq: 2
    longest: ABAC

length (5) - mostFreq (2) > k (2), moving the left forward

              r
    A B A C D E F
        l

    map: {A:1, C:1, D:1, E:1}
    mostFreq: 2
    longest: ABAC

char F

                r
    A B A C D E F
        l

    map: {A:1, C:1, D:1, E:1, F:1}
    mostFreq: 2
    longest: ABAC

length (5) - mostFreq (2) > k (2), moving the left forward

                r
    A B A C D E F
          l

    map: {C:1, D:1, E:1, F:1}
    mostFreq: 2
    longest: ABAC

end

# Demonstration 2

Example

* input: `A B A C C C C`
* k = 2
* output = 6

char A

    r
    A B A C C C C
    l

    map: {A:1}
    mostFreq: 1
    longest: A

char B

      r
    A B A C C C C
    l

    map: {A:1, B:1}
    mostFreq: 1
    longest: AB

char A

        r
    A B A C C C C
    l

    map: {A:2, B:1}
    mostFreq: 2
    longest: ABA

char C

          r
    A B A C C C C
    l

    map: {A:2, B:1, C:1}
    mostFreq: 2
    longest: ABAC

char C

            r
    A B A C C C C
    l

    map: {A:2, B:1, C:2}
    mostFreq: 2
    longest: ABAC

length (5) - mostFreq (2) > k (2), moving the left forward

            r
    A B A C C C C
      l

    map: {A:1, B:1, C:2}
    mostFreq: 2
    longest: ABAC

char C

              r
    A B A C C C C
      l

    map: {A:1, B:1, C:3}
    mostFreq: 3
    longest: BACCC              // update

char C

                r
    A B A C C C C
      l

    map: {A:1, B:1, C:4}
    mostFreq: 4
    longest: BACCCC             // update
