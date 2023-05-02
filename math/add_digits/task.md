# Intuition
The task is to find the sum of digits of the number `num` until a single-digit number is obtained. The initial idea is to iterate through all the digits of the number and sum them up, repeating this process until a single-digit number is achieved.

# Approach
We can use a loop to solve this problem. In the outer loop, we continue until the result becomes a single-digit number. In the inner loop, we go through each digit of the number, sum them up, and perform integer division by 10 to move on to the next digit. After the inner loop finishes, we update the number `num` with the sum of its digits and reset the value of `result`.

# Complexity
- Time complexity: `O(n)`, where `n` is the number of digits in the number `num`. The time complexity is linear since we iterate through each digit of the number.

- Space complexity: `O(1)`, as we only use a constant amount of additional space for storing the result and loop variables.
