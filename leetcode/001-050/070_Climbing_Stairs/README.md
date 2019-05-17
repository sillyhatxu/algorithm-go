# 70. Climbing Stairs

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

### Note: 

Given n will be a positive integer.

### Example 1:

```
Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

### Example 2:

```
Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

### Translate:

> 爬楼梯，每次只能爬一节楼梯或两节楼梯
>
> 输出爬楼梯的方案数

> input 1 output 1
> 
> input 2 output 2
> 
> input 3 output 3
> 
> input 4 output 5
> 
> input 5 output 8
> 
> input 6 output 13
> 
> input 7 output 21
> 
> input 8 output 34
> 
> input 9 output 55
> 
> input 10 output 89
> 
> input 11 output 144
> 
> ...
> 
> input 35 output 14930352
> 
> next = previous + previous'previous
> eg: input 3 = input(2) + input(1) = 2 + 1 = 3
>     input 4 = input(3) + input(2) = 3 + 2 = 5
>     input 5 = input(4) + input(3) = 5 + 3 = 8