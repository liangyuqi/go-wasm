
import "../lib/wasm_exec.js";

declare var Go
const go = new Go();
declare var WebAssembly

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
