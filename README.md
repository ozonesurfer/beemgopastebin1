beemgopastebin1
===============

A simple golang (Go) version of Pastebin using Beego, MongoDB and my model repository

# Dependencies

In order to build the website, you will first need to issue:

<code>go get github.com/ozonesurfer/pastemgomodel</code>
(that might require labix.org/v2/mgo)
<code>go get github.com/astaxie/beego<br>
go get github.com/beego/bee</code>

Build "bee" ("bee.exe" in Wondows) and put it in your PATH.

#Building

Add this repository to your GOPATH environment variable, go to src/beegopastebin1, then issue:

<code>bee run</code>
