package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsd0298a1ab9e9d77f039737f4b2018d10a3c02702 = "{{ define \"test.tmpl\" }}\r\n<html>\r\n<form action=\"http://localhost:8080/test\">\r\n        <input type=\"submit\" value=\"Refresh\" />\r\n    </form>\r\n<h1> {{.title}}</h1>\r\n<p>\r\n    Description: FSIP - Freund SIP Server, FSA - Freund SIP Audio\r\n</p>\r\n\r\n    <p1> {{.Device}}</p1>\r\n    <p2> {{.IP}}</p2>\r\n    <script>\r\n        let heartbeat = () => {\r\n          fetch('/ping')\r\n          .then(response => {\r\n            if(!response.ok) {\r\n              throw new Error('Network response was not ok.');\r\n            }\r\n          }).then(myJson => {\r\n            console.log(JSON.stringify(myJson));\r\n          }).catch(error => console.error('Error:', error));\r\n          ;\r\n          setTimeout(heartbeat, 3000);\r\n        }\r\n        heartbeat();\r\n    </script>\r\n\r\n</html>\r\n{{ end }}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"test.tmpl"}}, map[string]*assets.File{
	"/test.tmpl": &assets.File{
		Path:     "/test.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1543405499, 1543405499308455800),
		Data:     []byte(_Assetsd0298a1ab9e9d77f039737f4b2018d10a3c02702),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1543415475, 1543415475553077800),
		Data:     nil,
	}}, "")
