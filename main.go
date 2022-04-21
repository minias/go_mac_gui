package main

import "github.com/webview/webview"

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("TDI GroupWare")
	w.SetSize(1024, 768, webview.HintNone)
	//w.Navigate("http://kb-b3gk4oz6uss2k.koreasouth.cloudapp.azure.com:5601/app/kibana#/dev_tools/console?_g=()")
	w.Navigate("https://groupware.tdi9.com/login")
	w.Run()
}
