package main

import (
	"../Justgood"
	"fmt"
)

var print = fmt.Println

func main() {
	c := Justgood.Justgood("YOUR_APIKEY")
	resp := c.Lineqr("DESKTOPWIN\t5.21.3\tWindows\t10", "IMJUSTGOOD", "")
	qrlink  := resp.Result["qr"]
	barcode := resp.Result["barcode"]
	callback := resp.Result["callback"]
	md, ok := callback.(map[string]interface{})
	if ok == true {
		print(qrlink)
		print(barcode)
		resp     = c.LineqrGetPin(md["pin"].(string))
		if resp.Status == 200 {
    			pin := resp.Result["pin"]
    			print(pin)
		}
    		resp = c.LineqrGetToken(md["token"].(string))
    		apps    := resp.Result["app"]
		cert    := resp.Result["cert"]
		token   := resp.Result["token"]
		result  := fmt.Sprintf("Application : %s\n\n", apps)
		result += fmt.Sprintf("Certified : %s\n\n", cert)
		result += fmt.Sprintf("Authtoken : %s\n\n", token)
		print(result)
	}
}
