GOARCH=wasm GOOS=js go build -o fib.wasm fib.go

# 对应的操作系统名为js，对应的CPU类型为wasm。目前还无法通过go run的方式直接运行输出的wasm文件，因此我们需要通过go build的方式生成wasm目标文件，然后通过Node环境执行。