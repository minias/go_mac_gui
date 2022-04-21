package main

import "github.com/webview/webview"

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("비밥-앱서비스 ES")
	w.SetSize(1024, 768, webview.HintNone)
	w.Navigate("http://kb-b3gk4oz6uss2k.koreasouth.cloudapp.azure.com:5601/app/kibana#/dev_tools/console?_g=()")
	w.Run()
}
