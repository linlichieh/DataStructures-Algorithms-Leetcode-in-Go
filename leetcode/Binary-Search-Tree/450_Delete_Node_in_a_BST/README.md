# Idea

`insert` and `lookup` a node is easy, but removing a node is quite complex.

There are 3 cases to consider:

* no child: just delete
* 1 child: replace the targetted node with its child node
* 2 children: there are 2 options to do that (so that the tree can continue to follow the rules of BST).
    * find the minimum value in right subtree, assign it to the node we want to delete (most common)
    * find the maximum value in left subtree, assign it to the node we want to delete

# Demonstration

no child

             100                                100
            /    \                             /
          75     125 <- remove      =>       75

1 child

                    100                         100
                    /                           /
        remove -> 75              =>          65
                 /                            /\
               65                           60  70
               /\
             60  70

1 child

                    3                         3
                   / \                       / \
        remove -> 1   4           =>        2   4
                   \
                    2

1 child

           3                             3
          / \                           / \
         1   5 <- remove     =>        1   4
            /
           4

2 children (leftmost in right subtree is 70)

                        75                      75
                       /  \                    /  \
           remove -> 65    85       =>       70    85
                     /\                     /
                   60  70                  60

2 children (leftmost in right subtree is 115)

               100 <- remove                     115
              /   \                             /  \
            75     125               =>       75    125
                  /   \                               \
                115    150                            150


2 children (leftmost in right subtree is 80)

                         100                            100
                         /  \                          /   \
             remove -> 75    125                     80     125
                     /    \          =>            /    \
                   65      85                    65      85
                  /  \    /  \                  /  \       \
                60   70  80  95               60   70       95

2 children (leftmost in right subtree is 135)

                100                             100
               /   \                           /   \
             75     125 <- remove            75     135
                   /   \             =>            /   \
                115     150                      115   150
                       /   \                             \
                     135   175                           175

2 children (leftmost in right subtree is 34)

                2                                   2
              /   \                               /   \
             0     33 <- remove                 0      34
                  /  \                                /  \
                25    40                            25    40
                     /  \           =>                   /  \
                   34    45                            36    45
                     \                                /  \
                      36                            35    39
                      /\
                    35  39

