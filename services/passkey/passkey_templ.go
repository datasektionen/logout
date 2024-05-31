// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package passkey

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/datasektionen/logout/pkg/templates"
	"github.com/go-webauthn/webauthn/protocol"
)

func LoginPasskey(ca *protocol.CredentialAssertion) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = templ.JSONScript("credential-assertion", ca).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = base64ArrayBufferScript().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <script type=\"module\">\n\t\t\tlet decodebase64url = data =>\n\t\t\t\tUint8Array.from(window.atob(data.replace(/-/g, \"+\").replace(/_/g, \"/\")), v => v.charCodeAt(0));\n\t\t\tlet encodebase64url = data =>\n\t\t\t\tbase64ArrayBuffer(data).replace(/\\+/g, \"-\").replace(/\\//g, \"_\");\n\n\t\t\tlet ca = JSON.parse(document.querySelector(\"#credential-assertion\").textContent);\n\t\t\tca.publicKey.challenge = decodebase64url(ca.publicKey.challenge);\n\t\t\tfor (let c of ca.publicKey.allowCredentials) {\n\t\t\t\tc.id = decodebase64url(c.id);\n\t\t\t}\n\t\t\tlet cred = await navigator.credentials.get(ca);\n\t\t\tlet res = await fetch(\"/login/passkey\" + window.location.search, {\n\t\t\t\tmethod: \"post\",\n\t\t\t\theaders: { \"Content-Type\": \"application/json\" },\n\t\t\t\tbody: JSON.stringify({\n\t\t\t\t\tid: cred.id,\n\t\t\t\t\trawId: encodebase64url(cred.rawId),\n\t\t\t\t\ttype: cred.type,\n\t\t\t\t\tauthenticatorAttachment: cred.authenticatorAttachment,\n\t\t\t\t\tresponse: {\n\t\t\t\t\t\tauthenticatorData: encodebase64url(cred.response.authenticatorData),\n\t\t\t\t\t\tclientDataJSON: encodebase64url(cred.response.clientDataJSON),\n\t\t\t\t\t\tsignature: encodebase64url(cred.response.signature),\n\t\t\t\t\t\tuserHandle: encodebase64url(cred.response.userHandle),\n\t\t\t\t\t},\n\t\t\t\t}),\n\t\t\t});\n\t\t\tif (res.status == 200) window.location = \"/\";\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = templates.Layout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func AddPasskey(cc *protocol.CredentialCreation) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var4 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form onsubmit=\"addPasskey(event)\"><input id=\"passkey-name\" placeholder=\"passkey name\" name=\"name\" type=\"text\"> <button>Add</button></form>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.JSONScript("credential-creation", cc).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = base64ArrayBufferScript().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <script>\n\t\t\tlet decodebase64url = data =>\n\t\t\t\tUint8Array.from(window.atob(data.replace(/-/g, \"+\").replace(/_/g, \"/\")), v => v.charCodeAt(0));\n\t\t\tlet encodebase64url = data =>\n\t\t\t\tbase64ArrayBuffer(data).replace(/\\+/g, \"-\").replace(/\\//g, \"_\");\n\n\t\t\tasync function addPasskey(event) {\n\t\t\t\tevent.preventDefault();\n\t\t\t\tlet name = document.querySelector(\"#passkey-name\").value;\n\t\t\t\tlet cc = JSON.parse(document.querySelector(\"#credential-creation\").textContent);\n\t\t\t\tcc.publicKey.challenge = decodebase64url(cc.publicKey.challenge);\n\t\t\t\tcc.publicKey.user.id = decodebase64url(cc.publicKey.user.id);\n\t\t\t\tlet cred = await navigator.credentials.create(cc);\n\t\t\t\tlet res = await fetch(\"/passkey/add?name=\" + encodeURIComponent(name), {\n\t\t\t\t\tmethod: \"post\",\n\t\t\t\t\theaders: { \"Content-Type\": \"application/json\" },\n\t\t\t\t\tbody: JSON.stringify({\n\t\t\t\t\t\tid: cred.id,\n\t\t\t\t\t\ttype: cred.type,\n\t\t\t\t\t\tauthenticatorAttachment: cred.authenticatorAttachment,\n\t\t\t\t\t\tresponse: {\n\t\t\t\t\t\t\tattestationObject: encodebase64url(cred.response.attestationObject),\n\t\t\t\t\t\t\tclientDataJSON: encodebase64url(cred.response.clientDataJSON),\n\t\t\t\t\t\t},\n\t\t\t\t\t}),\n\t\t\t\t});\n\t\t\t\tif (res.status == 200) window.location = \"/account\";\n\t\t\t}\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = templates.Layout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func base64ArrayBufferScript() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n/*\nMIT LICENSE\nCopyright 2011 Jon Leighton\nPermission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the \"Software\"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:\nThe above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.\n*/\n\nfunction base64ArrayBuffer(arrayBuffer) {\n\tvar base64    = ''\n\tvar encodings = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'\n\n\tvar bytes         = new Uint8Array(arrayBuffer)\n\tvar byteLength    = bytes.byteLength\n\tvar byteRemainder = byteLength % 3\n\tvar mainLength    = byteLength - byteRemainder\n\n\tvar a, b, c, d\n\tvar chunk\n\n\t// Main loop deals with bytes in chunks of 3\n\tfor (var i = 0; i < mainLength; i = i + 3) {\n\t\t// Combine the three bytes into a single integer\n\t\tchunk = (bytes[i] << 16) | (bytes[i + 1] << 8) | bytes[i + 2]\n\n\t\t// Use bitmasks to extract 6-bit segments from the triplet\n\t\ta = (chunk & 16515072) >> 18 // 16515072 = (2^6 - 1) << 18\n\t\tb = (chunk & 258048)   >> 12 // 258048   = (2^6 - 1) << 12\n\t\tc = (chunk & 4032)     >>  6 // 4032     = (2^6 - 1) << 6\n\t\td = chunk & 63               // 63       = 2^6 - 1\n\n\t\t// Convert the raw binary segments to the appropriate ASCII encoding\n\t\tbase64 += encodings[a] + encodings[b] + encodings[c] + encodings[d]\n\t}\n\n\t// Deal with the remaining bytes and padding\n\tif (byteRemainder == 1) {\n\t\tchunk = bytes[mainLength]\n\n\t\ta = (chunk & 252) >> 2 // 252 = (2^6 - 1) << 2\n\n\t\t// Set the 4 least significant bits to zero\n\t\tb = (chunk & 3)   << 4 // 3   = 2^2 - 1\n\n\t\tbase64 += encodings[a] + encodings[b] + '=='\n\t} else if (byteRemainder == 2) {\n\t\tchunk = (bytes[mainLength] << 8) | bytes[mainLength + 1]\n\n\t\ta = (chunk & 64512) >> 10 // 64512 = (2^6 - 1) << 10\n\t\tb = (chunk & 1008)  >>  4 // 1008  = (2^6 - 1) << 4\n\n\t\t// Set the 2 least significant bits to zero\n\t\tc = (chunk & 15)    <<  2 // 15    = 2^4 - 1\n\n\t\tbase64 += encodings[a] + encodings[b] + encodings[c] + '='\n\t}\n\t\n\treturn base64\n}\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
