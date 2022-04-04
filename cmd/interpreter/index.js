globalThis.require = require;
globalThis.fs = require("fs");
globalThis.TextEncoder = require("util").TextEncoder;
globalThis.TextDecoder = require("util").TextDecoder;
globalThis.performance = {
    now() {
        const [sec, nsec] = process.hrtime();
        return sec * 1000 + nsec / 1000000;
    },
};

const crypto = require("crypto");

globalThis.crypto = {
    getRandomValues(b) {
        crypto.randomFillSync(b);
    },
};
require("./wasm_exec");
const assert = require('assert');
const fs = require('fs');

const testBytes = fs.readFileSync(__dirname + '/testdata/test.json');
const tc = JSON.parse(testBytes);
const input = tc.inputs[0];
const output = Object.fromEntries(Object.entries(tc.wantOutputs[0].actions).map(([k, v]) => [k, v.effect]));

console.log("expected", output);

const policyBytes = fs.readFileSync(__dirname + '/testdata/rps.yaml');
const bytes = fs.readFileSync(__dirname + '/main.wasm');

const go = new Go()

WebAssembly.instantiate(new Uint8Array(bytes), go.importObject)
    .then(obj => {
        go.run(obj.instance);
        const result = Authz(policyBytes.toString(), JSON.stringify(input));
        console.log("actual", result);
        assert.deepEqual(result, output);
        process.exit(1)
    })
