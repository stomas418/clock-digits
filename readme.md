# Four digits to clock problem 
I made this solution to a problem because I was asked to solve it in a technical test for a job and I couldn't solve it in the given time

## The problem

Given four digits, count how many valid times can be displayed on a digital clock (in 24-hour format) using those digits. The earliest time is 00:00 and the latest time is 23:59. Write a function that, given four integers A, B, C and D (describing the four digits), returns the number of valid times that can be displayed on a digital clock. 

Examples: 
1. Given A = 1, B = 8, C = 3, D = 2, the function should return 6. The valid times are: 12:38, 13:28, 18:23, 18:32, 21:38 and 23:18. 
2. Given A = 2, B = 3, C = 3, D = 2, the function should return 3. The valid times are: 22:33, 23:23 and 23:32.
3. Given A = 6, B = 2, C = 4, D = 7, the function should return O. It is not possible to display any valid time using the given digits. 

Assume that: A, B, C and D are integers within the range [0..9).

## The solution

My approach to solving the problem is a simple algorithm for each digit, ordered from lowest to highest, that works as follows:
1. Put it as the first number and check if it could be the first digit on a digital clock format (either 0, 1 or 2). If it can't be the first number then we just return the result up to that point, else move to step two.
2. For every other digit, check if it can be the second digit ( `[0, 9]` if first is 0 or 1, or `[0, 3]` if first is 2). If it can, move to step two.
3. For the other 2 numbers, check if either can be the third digit (`[0, 5]`). If one of them fits, then the other will be the fourth digit as any digit fits the range. Move to step four.
4. Check if that exact combination of 4 digits has been already used, if not then add 1 to the result.