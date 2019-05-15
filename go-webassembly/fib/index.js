"use strict";
exports.__esModule = true;
require("../lib/wasm_exec.js");
var go = new Go();
WebAssembly.instantiateStreaming(fetch("fib.wasm"), go.importObject).then(function (result) {
    go.run(result.instance);
});
function callFib() {
    var paramInput = document.getElementById("param");
    // @ts-ignore
    var n = parseInt(paramInput.value || "0");
    // 传入输入参数和回调函数
    // 回调函数负责呈现结果
    // @ts-ignore
    window.fib(n, function (result) {
        var resultDom = document.getElementById("result");
        // @ts-ignore
        resultDom.value = result;
    });
}
