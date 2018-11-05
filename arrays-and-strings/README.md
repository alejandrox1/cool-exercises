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

4. **URLify** Write a method to replace all spaces in a string with `%20`. You
   may assume that the string has sufficient space at the end to hold the
   additional characters, and that you are given the "true" length of the
   string.

    Input: `MR John Smith   `, 13
    Output: `MR%20John%20Smith`

    * It is often easiest to modify strings by going from the end of the string
      to the beginning.
    * Try using a hash table that maps from a `runningSum` value to the number
      elements with the `runningSum`.
