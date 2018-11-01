1. **Remove dups** write code to removed duplicates from an unsorted linked
   list. 
    * How would you solve this problem is a temporary buffer is not allowed?
    * Try using a hash table.
    * Without extra space, it will take `O(n^2)` time. Try using two pointers,
      where the second one searches ahead of the list.
