export let decodebase64url = (data: string) =>
    Uint8Array.from(window.atob(data.replace(/-/g, "+").replace(/_/g, "/")), v => v.charCodeAt(0));
export let encodebase64url = (data: Uint8Array) =>
    base64ArrayBuffer(data).replace(/\+/g, "-").replace(/\//g, "_");

// Taken from https://gist.github.com/jonleighton/958841 and slightly modified.
/*
MIT LICENSE
Copyright 2011 Jon Leighton
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

function base64ArrayBuffer(arrayBuffer: ArrayBuffer) {
    let base64 = "";
    let encodings = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

    let bytes = new Uint8Array(arrayBuffer);
    let byteLength = bytes.byteLength;
    let byteRemainder = byteLength % 3;
    let mainLength = byteLength - byteRemainder;

    // Main loop deals with bytes in chunks of 3
    for (let i = 0; i < mainLength; i = i + 3) {
        // Combine the three bytes into a single integer
        let chunk = (bytes[i] << 16) | (bytes[i + 1] << 8) | bytes[i + 2];

        // Use bitmasks to extract 6-bit segments from the triplet
        let a = (chunk & 16515072) >> 18; // 16515072 = (2^6 - 1) << 18
        let b = (chunk & 258048) >> 12; // 258048   = (2^6 - 1) << 12
        let c = (chunk & 4032) >> 6; // 4032     = (2^6 - 1) << 6
        let d = chunk & 63;               // 63       = 2^6 - 1

        // Convert the raw binary segments to the appropriate ASCII encoding
        base64 += encodings[a] + encodings[b] + encodings[c] + encodings[d];
    }

    // Deal with the remaining bytes and padding
    if (byteRemainder == 1) {
        let chunk = bytes[mainLength];

        let a = (chunk & 252) >> 2; // 252 = (2^6 - 1) << 2

        // Set the 4 least significant bits to zero
        let b = (chunk & 3) << 4; // 3   = 2^2 - 1

        base64 += encodings[a] + encodings[b] + '==';
    } else if (byteRemainder == 2) {
        let chunk = (bytes[mainLength] << 8) | bytes[mainLength + 1];

        let a = (chunk & 64512) >> 10; // 64512 = (2^6 - 1) << 10
        let b = (chunk & 1008) >> 4; // 1008  = (2^6 - 1) << 4

        // Set the 2 least significant bits to zero
        let c = (chunk & 15) << 2; // 15    = 2^4 - 1

        base64 += encodings[a] + encodings[b] + encodings[c] + '=';
    }

    return base64;
}
