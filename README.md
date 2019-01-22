# Sstr
Go语言中封装的字符串库,包含一些官方库中没有的功能

![](https://img.shields.io/badge/sstr-0.1-green.svg)

#### Sstr基本使用样例
```
s := Sstr{"Hello,World!"}
fmt.Println(s.val())
```

#### RemoveAllSpaces
功能解释:

去除字符串中所有空格

使用样例:
```
s := Sstr{"  1  1 2 3 2  "}
s.RemoveAllSpaces()
```

#### KeepNumbers
功能解释:

保留字符串中的数字

使用样例:
```
s := Sstr{"abc123"}
s.KeepNumbers()
```

#### KeepUpperChar
功能解释:

保留字符串中的大写字符

使用样例:
```
s := Sstr{"ASHdwadwa213"}
s.KeepUpperChar()
```

#### KeepLowerChar
功能解释:

保留字符串中的小写字符

使用样例:
```
s := Sstr{"ASHdwadwa213"}
s.KeepLowerChar()
```

#### KeepSpecificChar
功能解释:

保留特定字符 可变参数 第一个保留数字 第二个保留大写字母 第三个保留小写字母

(可以随意组合 保留true 否则false 三个参数必须齐全) 

使用样例:
```
s1 := Sstr{"ASBabc123"}
s1.KeepSpecificChar(true,false,false)

s2 := Sstr{"ASBabc123"}
s2.KeepSpecificChar(false,true,false)

s3 := Sstr{"ASBabc123"}
s3.KeepSpecificChar(false,false,true)
```