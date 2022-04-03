// requires Node.js v15
crypto = require('crypto');
const fs = require('fs');
require("./wasm_exec");

const go = new Go()
const inputBytes = fs.readFileSync(__dirname + '/testdata/input.json');
input = JSON.parse(inputBytes).inputs[0];

const policyBytes = fs.readFileSync(__dirname + '/testdata/rps.yaml');
const bytes = fs.readFileSync(__dirname + '/main.wasm');

WebAssembly.instantiate(new Uint8Array(bytes), go.importObject)
    .then(obj => {
        go.run(obj.instance);
        const result = Authz(policyBytes.toString(), JSON.stringify(input));
        console.log("result", result)
    })
