# Exercises

1. **Is unique** Implement an algorithm to determine if a string has all unique
   characters. What if you cannot use additional data structures?
    * Try a hash table
    * Could a bit vector be useful?
    * Can you solve it in `O(N log N)` time?

    To run benchmarks: `go test -bench=.`

2. **Check permutation** Give two strings, write a method to decide if one is a
   permutation of the other.
    * If they are permutations of each other then the character count should be
      the same.
    * If you sort and compare the strings then you get `O(n log n)` time. If
      you use a hash table then you'll take up some space but you get `O(n)`.

3. **Sum of pairs** Find a pair of numbers in an array that add up to a 
   given sum.
