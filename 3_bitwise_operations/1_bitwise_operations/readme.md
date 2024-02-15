# Bitwise operations

operations:
- OR |
- AND &
- XOR ^
- Left Shift <<
- Right Shift >>
- NOT ^x (e.d. 1 ^ x)

```
x ? y = y ? x
(x ? y) ? z = x ? (y ? z)

x ? 0
x & 0 = 0
x | 0 = x
x ^ 0 = x

x & x = x
x | x = x
x ^ x = 0

xx <--- 00
00001100 << 2   (12) 
      ||
00110000        (48)

00 -->xx
00001100 >> 2   (12)
      ||
00000011        (3)

12 << 2 = 48 = 12 * 4
12 >> 2 = 3  = 12 / 4 

n << i = n * 2^i
n >> i = n / 2^i

2^i = 1 << i
```



