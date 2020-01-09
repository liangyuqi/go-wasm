
import "../lib/wasm_exec.js";

declare var Go
const go = new Go();
declare var WebAssembly

// WebAssembly.instantiateStreaming() 方法直接从流式底层源编译和实例化WebAssembly模块。这是加载wasm代码一种非常有效的优化方式
WebAssembly.instantiateStreaming(
  fetch("fib.wasm"),
  go.importObject
).then(result => {
  go.run(result.instance);
});

// @ts-ignore
window.callFib = function() {
  let paramInput = document.getElementById("param")
  // @ts-ignore
  let n = parseInt(paramInput.value || "0")
  // 传入输入参数和回调函数
  // 回调函数负责呈现结果
  // @ts-ignore
  window.fib(n, function(result) {
      var resultDom = document.getElementById("result")
      // @ts-ignore
      resultDom.value = result
  })
}




// 不引入 lib写法
// var importObj = {
// 	env: {
// 		memoryBase: 0,
// 		tableBase: 0,
// 		memory: new WebAssembly.Memory({initial:0,maximum:0}),
// 		table: new WebAssembly.Table({initial:2,maximum:2,element:'anyfunc'}),
// 		abort: ()=>{}
// 	}
// }

// // 封装的异步loader
// function load(path) {
// 	return fetch(path)
// 		// 获取二进制buffer
// 		.then(res => res.arrayBuffer())
// 		// 编译&实例化，导入js对象
// 		.then(bytes => WebAssembly.instantiate(bytes, importObj))
// 		// 返回实例
// 		.then(res => res.instance)
// }