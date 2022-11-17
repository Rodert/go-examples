[toc]

## 文档地址

<https://blog.csdn.net/qq_40374604/category_12089625.html>


## csv 预览

![在这里插入图片描述](https://img-blog.csdnimg.cn/46e6323b1e7a4c4db9a446b96d77f3c9.png)


## 结果

![在这里插入图片描述](https://img-blog.csdnimg.cn/45fb6735e2ad49e49f91853a11dbdf82.png)


## 特殊情况

1. csv解析，下面的列大于首行的列；

> 以读txt的方式按行读取，再切割拼装数组。

2. 编码格式异常，`bare " in non-quoted-field`；

> 经测试不是bom头的问题，解决办法：粘贴复制到一个新的csv，重新识别。

> 猜测原因：csv是被其他文件改了文件后缀得来的。





