# 6. ZigZag Conversion

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:

(you may want to display this pattern in a fixed font for better legibility)
```
P   A   H   N
A P L S I I G
Y   I   R
```
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

### Example 1:

```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```

### Example 2:

```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"

Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
```

### Translate:
```
字符串"PAYPALISHIRING"根据给定行数，以一个Z字形显示出来，
如下所示：（您可能希望以固定字体显示此模式以获得更好的可读性）
其实例子太小了，大家看着不太方便看，所以我举一个超级大例子
Z             Z
Z           Z Z
Z         Z	  Z
Z       Z	  Z
Z     Z		  Z
Z   Z		  Z
Z Z			  Z
Z			  Z
这就是一个放倒了的 Z字母

制作多个图形得出公式  (数字间隔)M = 2N - 2
```