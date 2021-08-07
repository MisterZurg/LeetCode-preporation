# Count and Say

The **count-and-say** sequence is a sequence of digit strings defined by the recursive formula:
* ```countAndSay(1) = "1"```
* ```countAndSay(n)``` is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.

For example, the saying and conversion for digit string ```"3322251"```:
> ```"3322251"```
> 
> two 3's, three 2's, one 5, and one 1
> 2 + 3 
>