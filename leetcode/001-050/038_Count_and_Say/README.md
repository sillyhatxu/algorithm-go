# 38. Count and Say

The count-and-say sequence is the sequence of integers with the first five terms as following:

```
1.     1
2.     11
3.     21
4.     1211
5.     111221
```

`1` is read off as `"one 1"` or `11`.
`11` is read off as `"two 1s"` or `21`.
`21` is read off as `"one 2, then one 1"` or `1211`.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

### Example 1:

```
Input: 1
Output: "1"
```

### Example 2:

```
Input: 4
Output: "1211"
```

### Translate:

```
我觉得这题描述有问题，题意应该是

n=1时，输出字符串1；
n=2时，数上次字符串中各个数值的个数，因为上个数字字符串中有1个1，所以输出11；
n=3时，由于上个字符串是11，有2个1，所以输出21；
n=4时，由于上个数字的字符串是21，有1个2和1个1，所以输出1211，依次类推...... 

例子1： 
输入: 1 
输出: "1" 

例子2： 
输入: 4 
输出: "1211"
```