

#golang 单元测试&&性能测试


## 一:单元测试

### 1.为什么要做单元测试和性能测试
- 减少bug
- 快速定位bug
- 减少调试时间
- 提高代码质量

### 2.golang的单元测试
- 单元测试代码的go文件必须以_test.go结尾
- 单元测试的函数名必须以Test开头，是可导出公开的函数
- 测试函数的签名必须接收一个指向testing.T类型的指针，并且不能返回任何值
 
### 3.golang单元测试组
- 有好几个不同的输入以及输出组成的一组单元测试 
    
### 4.测试覆盖率 
- ``go test -cover`` 直接输出覆盖率
- ``go test -run TestAll`` 指定要运行的单元测试函数
- ``go test -cover -coverprofile=c.out`` 
- ``go tool cover -html=c.out -o coverage.html``  输出具体的代码覆盖情况


## 二：性能测试

- 执行所有函数 ``go test -bench=".*"``

- 执行指定函数 ``go test -bench="BenchmarkWithPool"``

- 生成内存和cpu分析文件 ``go test -bench="BenchmarkWithPool" -cpuprofile cpu.out -memprofile mem.out``

- 分析cpu文件 ``go tool pprof cpu.out``  

- 分析内存文件 ``go tool pprof mem.out`` 

- 输出png图 图中红色方块面积越大 证明申请的内存越大或者消耗cpu越多



[源码链接](https://github.com/quguolin/goBase/tree/master/unitTestAndBenchmark)
