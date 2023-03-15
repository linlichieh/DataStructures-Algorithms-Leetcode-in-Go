# Idea

In a preorder traversal, the first Preorder element represents the root of the tree. This means that everything to the left of 7
in the inorder traversal is part of the left subtree, and everything to the right of 7 is part of the right subtree.

# Demonstration

        7
       / \
      5   13
     / \    \
    8   6    21
       /    /
      4    9

pre: [7, 5, 8, 6, 4, 13, 21, 9]
in: [8, 5, 4, 6, 7, 13, 9, 21]


element 7

    [    7     5     8     6     4     13     21     9    ]
         *                      mid

    [    8     5     4     6     7     13     9     21    ]
    [          left         ]   mid   [       right       ]

left (element 5)

    [    7     5     8     6     4     13     21     9    ]
               *                mid

    [    8     5     4     6     7     13     9     21    ]
     [ left ]      [ right  ]   mid


right (element 13)

    [    7     5     8     6     4     13     21     9    ]
                                mid    *

    [    8     5     4     6     7     13     9     21    ]
                                mid          [ right  ]

and so on ...

simple demonstation version:

* 7 (preorder)
* left [8, 5, 4, 6] (inorder)
    * 5 (preorder) [5, 8, 6, 4]
    * left [8]
        * 8 (preorder)
        * left []
        * right []
    * right [4,6]
        * 6 (preorder) [6, 4]
        * left [4]
            * 4 (preorder)
            * left []
            * right []
        * right []
* right: [13, 9, 21] (inorder)
    * 13 (preorder) [13, 21, 9]
    * left []
    * right [9,21]
        * 21 (preorder)
            * left [9]
                * 9 (preorder)
                * left []
                * right []
            * right []
