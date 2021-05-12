unique binary search treees 

For example, F(3,7), the number of unique BST tree with the number 3 as its root. To construct an unique BST out of the entire sequence [1, 2, 3, 4, 5, 6, 7] with 3 as the root, which is to say, we need to construct a subtree out of its left subsequence [1, 2] and another subtree out of the right subsequence [4, 5, 6, 7], and then combine them together (i.e. cartesian product). Now the tricky part is that we could consider the number of unique BST out of sequence [1,2] as 
G(2), and the number of of unique BST out of sequence [4, 5, 6, 7] as G(4). For G(n), it does not matter the content of the sequence, but the length of the sequence. Therefore,  F(3,7) = G(2) ⋅ G(4). To generalise from the example, we could derive the following formula:

F(i,n) = G(i−1)⋅ G(n−i)

---
∑ G(i−1) ⋅ G(n−i)

