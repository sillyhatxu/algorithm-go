# 11. Container With Most Water

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

### Note: 

You may not slant the container and n is at least 2.

![](https://lh3.googleusercontent.com/r_LsiwktuN0QrPbfujA7M5J1KDQt2W24d6ZipakS1QhW0uUlsWxFVlzFuLrUFYmneIthaLsjqSovc_1HaT31K8mErqOzDL6ViaaS_ej2j0Jj91LQDsUTOFMUbihePWM4H6q6zigZIVVYRRee2hSJYC3mHRGRwzVH39wHS_WBuURSfHI3TqJO89yDO29SSYnH9lk-JiDL2QjUGPvioVT2v4Ky0MBkkqmim-kAPV9VTJb6mctrrV7xtgNhtRyUr47FyRA_YEis4wvbV4t6IwtEy3ZkAqw5hva83fZWC66PllREa8X0wRF9_hKjRW35V9F-Ky5ciTVMMZt_Br7FWM2tj17_QMQSV23BMmK7MdbOsZ9JZbrmg5anKaSmPluLIEULdT5Rd9sVmI0ClJHvPPceiPBHdwPOkcwNya8EAZYHxKWu2-BeCxkS5BkKOpi7HblnqJ2_rJxF7ZvtsbLdQnPK6l9MWx3fiLQyPMGFoJKx2B-GLSlOV5slzMKKrD-Y-xLs0g3E0nXaBenW3zLfVT2Lo9xb--PQXrKzUzgkdptqGh39a4UFxHxptCSz5ssHtR3vUf9E2X2BYSVh5BNzlVZ159MGaqPVL5pu38dZQ2tHKITzKRCPti31v2LTTTE1nIswbRaE5RnkaqU5ekccy3x-jiy6wu4N9dY=w801-h383-no)

`The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.`

### Example:

```
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```

### Translate:

> 给定一个非负数的数组，数组中每个数字a1都是一个从零到a1的立柱，
>
> 按数组顺序，选出两个立柱装水，使其体积最大