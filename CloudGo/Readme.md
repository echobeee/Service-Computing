# CloudGo
-------------------------------
## 框架使用
-------------------------------
[gorilla/mux][1]、[codegangsta/negroni][2]、[unrolled/render][3]
## 运行程序
-------------------------------
> $ go run ./cloudgo/main.go -p9090
### 测试命令curl
> $ curl -v http://localhost:3000/hello/echo
``` script
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 3000 (#0)
> GET /hello/echo HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.52.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Thu, 11 Nov 2017 10:18:08 GMT
< Content-Length: 23
< 
{
  "name": "echo"
}
* Curl_http_done: called premature == 0
* Connection #0 to host localhost left intact
```
### ab测试
> $ ab -n 100 -c 10 http://localhost:3000/hello/ab
``` script
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            3000

Document Path:          /hello/ab
Document Length:        23 bytes

Concurrency Level:      10
Time taken for tests:   0.066 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      14600 bytes
HTML transferred:       2300 bytes
Requests per second:    1611.78 [#/sec] (mean)   
Time per request:       5.824 [ms] (mean)
Time per request:       0.661 [ms] (mean, across all concurrent requests)
Transfer rate:          215.56 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   4.3      1      19
Processing:     0    4   3.8      2      20
Waiting:        0    3   3.6      2      18
Total:          0    6   5.5      4      20

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      8
  75%      8
  80%      9
  90%     20
  95%     20
  98%     20
  99%     20
 100%     20 (longest request)
```
#### 结果分析
- 吞吐率为1611.78 reqs/s，即每秒处理1611.78个请求
- 用户平均请求等待时间为0.661 ms 

[1]: https://github.com/gorilla/mux
[2]: https://github.com/codegangsta/negroni
[3]: https://github.com/unrolled/render
