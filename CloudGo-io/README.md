# CloudGo-inout
----------------------------------
## 任务要求
1. 支持静态文件服务
2. 支持简单 js 访问
3. 提交表单，并输出一个表格
4. 对 /unknown 给出开发中的提示，返回码 5xx
------------------------------------
## 实现结果
1.访问静态文件
```script
$ go run main.go
[negroni] listening on :8080
```
用浏览器进入localhost:8080能看到到assets/index.html网页
```script
[negroni] 2017-11-21T14:10:13+08:00 | 200 | 	 2.357753ms | localhost:8080 | GET / 
[negroni] 2017-11-21T14:10:05+08:00 | 200 |      56.332µs | localhost:8080 | GET /main.css
[negroni] 2017-11-21T14:11:50+08:00 | 200 |      25.299µs | localhost:8080 | GET /check.js
```
可以看见程序相应了客户端的请求，返回了主页，以及静态目录的css文件和js文件

2.访问localhost:8080/api,javascript代码通过获取到Go返回的JSON对象,一串JSON序列化输出数据
可以看到console记录下返回的数据
```script
echobee
20
```

3.提交表格 返回表单
```script
[negroni] 2017-11-21T14:34:18+08:00 | 200 |      500.4µs | localhost:8080 | POST /
```

4.对于未知的路径,访问localhost:8080/unknown浏览器会看见501 not implement的字样

  而服务端也会显示出现了501错误：
  ```script
  [negroni] 2017-11-21T14:36:38+08:00 | 501 |      1.0035ms | localhost:8080 | GET /unknown
  ```
