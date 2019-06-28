# 自己手动写java虚拟机学习项目

## Java命令常用选项及用途
选项|用途
:----:|:-----:
-version|输出版本信息，然后退出
-?/-help|输出帮助信息，然后退出
-cp/-classpath|指定用户类路径
-Dproperty=value|设置Java系统属性
-Xms<size>|设置初始堆空间大小
-Xmx<size>|设置最大堆空间大小
-Xss<size>|设置线程栈空间大小

## 类路径
1、启动类路径（bootstrap classpath）默认jre/lib

2、扩展类路径（extension classpath）默认jre/lib/ext

3、用户类路径（user classpath）

        用户类路径默认是当前目录，也就是“.”。可以通过CLASSPATH环境变量设置。-classpath选项
    优先级更高，可以覆盖CLASSPATH环境变量设置。
        -classpath可以指定目录，或者jar文件和zip文件。
        java -cp path/com.xieong...
        java -cp path/lib.jar
        java -cp path/lib.zip
        可以同时指定多个目录和文件，用分号隔开
        java -cp path/com.xieong.controller;path/lib.jar;path/lib.zip ...
        可以用通配符 * 指定某个目录下的所有jar文件
        java -cp path/com.xieong.controller;path/* ...

