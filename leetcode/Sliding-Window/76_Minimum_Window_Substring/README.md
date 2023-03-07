# Idea

Use the sliding window approach by iterating through the characters of string `s`. Increase the counter and add into the queue for each character if it's in the string `t`.
Once the window contains all characters of string `t`, find and update the shortest substring, and pop the first element from the queue in order to reset the window range and decrease the counter.
Until the current window no longer fulfills the criteria, then continue iterating through the rest or the characters and repeat the process until reaching the end.

# Demonstration

Example

* string `s`: `ADCECDZIBZBECDCAGFB`
* string `t`: `ABCC`
* output: `BECDCA`

char A

    ADCECDZIBZBECDCAGFB
    *

    h = {A:1}
    matched = 1
    queue = [A]         // for better to understand, use the char to demonstrate instead of the index in the string
    shorest = ""

char D (do nothing)

    ADCECDZIBZBECDCAGFB
     *

    h = {A:1}
    matched = 1
    queue = [A]
    shorest = ""

char C

    ADCECDZIBZBECDCAGFB
      *

    h = {A:1, C:1}
    matched = 1         // C still doesn't match the critiri, it needs 2
    queue = [A, C]
    shorest = ""

char E (do nothing)

    ADCECDZIBZBECDCAGFB
       *

    h = {A:1, C:1}
    matched = 1
    queue = [A, C]
    shorest = ""

char C

    ADCECDZIBZBECDCAGFB
        *

    h = {A:1, C:2}
    matched = 2         // C has 2, matched++
    queue = [A, C, C]
    shorest = ""

char D (do nothing)

    ADCECDZIBZBECDCAGFB
         *

    h = {A:1, C:2}
    matched = 2
    queue = [A, C, C]
    shorest = ""

char Z (do nothing)

    ADCECDZIBZBECDCAGFB
          *

    h = {A:1, C:2}
    matched = 2
    queue = [A, C, C]
    shorest = ""

char I (do nothing)

    ADCECDZIBZBECDCAGFB
           *

    h = {A:1, C:2}
    matched = 2
    queue = [A, C, C]
    shorest = ""

char B

    ADCECDZIBZBECDCAGFB
            *

    h = {A:1, C:2, B:1}
    matched = 3             // matched++ due to B
    queue = [A, C, C, B]
    shorest = ""

matched = 3 (window matches all critiria), pop the queue and update the shortest

    ADCECDZIBZBECDCAGFB
    p       *

    h = {C:2, B:1}
    matched = 2
    queue = [C, C, B]
    shorest = "ADCECDZIB"           // update

char Z (matched = 2, continue to iterate) (do nothing)

    ADCECDZIBZBECDCAGFB
             *

    h = {C:2, B:1}
    matched = 2
    queue = [C, C, B]
    shorest = "ADCECDZIB"

char B

    ADCECDZIBZBECDCAGFB
              *

    h = {C:2, B:2}
    matched = 2
    queue = [C, C, B, B]
    shorest = "ADCECDZIB"

char E (do nothing)

    ADCECDZIBZBECDCAGFB
               *

    h = {C:2, B:2}
    matched = 2
    queue = [C, C, B, B]
    shorest = "ADCECDZIB"

char C

    ADCECDZIBZBECDCAGFB
                *

    h = {C:3, B:2}
    matched = 2
    queue = [C, C, B, B, C]
    shorest = "ADCECDZIB"

char D (do nothing)

    ADCECDZIBZBECDCAGFB
                 *

    h = {C:3, B:2}
    matched = 2
    queue = [C, C, B, B, C]
    shorest = "ADCECDZIB"

char C

    ADCECDZIBZBECDCAGFB
                  *

    h = {C:4, B:2}
    matched = 2
    queue = [C, C, B, B, C, C]
    shorest = "ADCECDZIB"

char A

    ADCECDZIBZBECDCAGFB
                   *

    h = {C:4, B:2, A:1}
    matched = 3                         // matched++ due to A
    queue = [C, C, B, B, C, C, A]
    shorest = "ADCECDZIB"

matched = 3 (window matches all critiria), pop the queue and update the shortest

    ADCECDZIBZBECDCAGFB
      p            *

    h = {C:3, B:2, A:1}
    matched = 3
    queue = [C, B, B, C, C, A]
    shorest = "ADCECDZIBZB"         // "CECDZIBZBECDCA" isn't the shortest

matched = 3 (window matches all critiria), pop the queue and update the shortest

    ADCECDZIBZBECDCAGFB
        p          *

    h = {C:2, B:2, A:1}
    matched = 3
    queue = [B, B, C, C, A]
    shorest = "ADCECDZIBZB"         // "CDZIBZBECDCA" isn't the shortest

matched = 3 (window matches all critiria), pop the queue and update the shortest

    ADCECDZIBZBECDCAGFB
            p      *

    h = {C:2, B:1, A:1}
    matched = 3
    queue = [B, C, C, A]
    shorest = "BZBECDCA"            // update

matched = 3 (window matches all critiria), pop the queue and update the shortest

    ADCECDZIBZBECDCAGFB
              p    *

    h = {C:2, A:1}
    matched = 2
    queue = [C, C, A]
    shorest = "BECDCA"              // update

char G (matched = 2, continue to iterate) (do nothing)

    ADCECDZIBZBECDCAGFB
                    *

    h = {C:2, A:1}
    matched = 2
    queue = [C, C, A]
    shorest = "BECDCA"

char F (do nothing)

    ADCECDZIBZBECDCAGFB
                     *

    h = {C:2, A:1}
    matched = 2
    queue = [C, C, A]
    shorest = "BECDCA"

char B (do nothing)

    ADCECDZIBZBECDCAGFB
                      *

    h = {C:2, B:1, A:1}
    matched = 3
    queue = [C, C, A, B]
    shorest = "BECDCA"

matched = 3 (window matches all critiria), pop the queue and update the shortest

    ADCECDZIBZBECDCAGFB
                p     *

    h = {C:1, B:1, A:1}
    matched = 2
    queue = [C, A, B]
    shorest = "BECDCA"     // "CDCAGFB" isn't the shortest

out of range

    ADCECDZIBZBECDCAGFB
                       *

    h = {C:1, B:1, A:1}
    matched = 2
    queue = [C, A, B]
    shorest = "BECDCA"     // "CDCAGFB" isn't the shortest

end
