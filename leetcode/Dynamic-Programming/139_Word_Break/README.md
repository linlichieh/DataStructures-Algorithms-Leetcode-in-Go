# Idea

Create a map to store the wordDict

Initialise a boolean slice `dp` with the size `len(s)`+1. `dp[i]` represent whether the substring `s[0:i]` can be segmented into words form the wordDict

`dp[0]` is true because an empty string can always be segmented as well

For example, `s = "leetcode" wordDict = ["leet", "code"]`, "leetcode" can be segmented into "", "leet" and "code"


# Demonstration 1

input

    s = "leetcode"
    wordDict = ["leet", "code"]

initialisation

    index   0  1  2  3  4  5  6  7  8
    string  l  e  e  t  c  o  d  e
    dp      T  F  F  F  F  F  F  F  F

> Note that string is 1-based array and dp is 0-based array, which reserve index `0` as true for an empty array

the first segment matched

    j i
    --------------------
    0 1 l
    --------------------
    0 2 le
    1 2  e
    --------------------
    0 3 lee
    1 3  ee
    2 3   e
    --------------------
    0 4 leet (matched, break the loop)
    --------------------

> `i=0, j=4` -> `"leetcode"[0:4]` -> `"leet"` (start at `0` with length 4)

then update dp at index `i` which is 4 (as opposed to 3)

    index   0  1  2  3  4  5  6  7  8
    string  l  e  e  t  c  o  d  e
    dp      T  F  F  F  T  F  F  F  F
                        *
> It is easy to get confused because the dp array is 1-based, while the string is 0-based.

the second segment matched

    j i
    --------------------
    0 5 leetc
    1 5  eetc
    2 5   etc
    3 5    tc
    4 5     c
    --------------------
    0 6 leetco
    1 6  eetco
    2 6   etco
    3 6    tco
    4 6     co
    5 6      o
    --------------------
    0 7 leetcod
    1 7  eetcod
    2 7   etcod
    3 7    tcod
    4 7     cod
    5 7      od
    6 7       d
    --------------------
    0 8 leetcode
    1 8  eetcode
    2 8   etcode
    3 8    tcode
    4 8     code (matched, break the loop)
    --------------------

then update dp at index 8

    index   0  1  2  3  4  5  6  7  8
    string  l  e  e  t  c  o  d  e
    dp      T  F  F  F  T  F  F  F  T
                                    *
The last element of dp is true, it indicates that the string can be segmented

# Demonstration 2

input

    s = "aaaaaaa"
    wordDict = ["aaaa", "aaa"]

initialisation

    index   0  1  2  3  4  5  6  7
    string  a  a  a  a  a  a  a
    dp      T  F  F  F  F  F  F  F

the first segment matched

    j i
    --------------------
    0 1 a
    --------------------
    0 2 aa
    1 2  a
    --------------------
    0 3 aaa (matched, break the loop)

then update dp at index `i` which is 3

    index   0  1  2  3  4  5  6  7
    string  a  a  a  a  a  a  a
    dp      T  F  F  T  F  F  F  F

the second segment matched

    j i
    --------------------
    0 4 aaaa (matched, break the loop)

then update dp at index `i` which is 4

    index   0  1  2  3  4  5  6  7
    string  a  a  a  a  a  a  a
    dp      T  F  F  T  T  F  F  F

the third segment matched

    j i
    --------------------
    0 5 aaaaa
    1 5  aaaa
    2 5   aaa
    3 5    aa
    4 5     a
    --------------------
    0 6 aaaaaa
    1 6  aaaaa
    2 6   aaaa
    3 6    aaa (matched, break the loop)

then update dp at index `i` which is 6

    index   0  1  2  3  4  5  6  7
    string  a  a  a  a  a  a  a
    dp      T  F  F  T  T  F  T  F

the forth segment matched

    j i
    --------------------
    0 7 aaaaaaa
    1 7  aaaaaa
    2 7   aaaaa
    3 7    aaaa (matched, break the loop)

then update dp at index `i` which is 7

    index   0  1  2  3  4  5  6  7
    string  a  a  a  a  a  a  a
    dp      T  F  F  T  T  F  T  T

The last element of dp is true, it indicates that the string can be segmented
