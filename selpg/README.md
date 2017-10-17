# illustration for HW1:selpg
------
###设计思路
1.命令行检查 对结构体内的值进行判断，检查参数个数，参数值是否符合要求

2.读写部分 通过参数判断是通过键盘读入还是文件读入，读取时根据-f的存在与否确定分页方式。

3.管道部分 使用了 os/exec 包来建立用于进程间通信的管道。

------

###usage
* ./selpg -startPage -endPage  [f | -linesPerPage] [ -ddest] [-fileName]
* example：./selpg -s3 -e5 -l8 test.txt

