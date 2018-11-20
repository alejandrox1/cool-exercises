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

5. **Palindrome permutation** Given a string, write a function to check if it
   is a permutation of a palindrome. A palindrome is a word or phrase that is
   the same forwards and bacwards. A permutation is a rearrangemen of letters.
   The palindrome does not need to be limited to just dictionary words.
    * To decide if a string is a permutation of a palindrome, we need to know
if it can be writen the same forwards ad backwards. So we need to have an even
number of all characters if the length of the string is an even number,
otherwise at most one character can have an odd count (the middle character).

6. **One away** There are three types of edits that can be performed on a
   string: insert a character, remove a character, or replace a character.
Given teo strings, write a function to check if they are one edit or zero edits
away.
    * Can you do all checks in one pass?

7. **String compression** Implement a method to do basic string compression
   using the counts of repeated characters. For example, `aabcccccaaa` would
become `a2b1c5a3`. If the compressed string does not become smallaer, then
return the original. You can assume the string has only upper and lower case
letters (a-z).
