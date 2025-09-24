package main


	// "bytes"
	// "encoding/base64"
	//"encoding/base64"
	// "crypto/hmac"
	// "crypto/sha256"
	// "fmt"

	//"golang.org/x/text/message"


import	"ecommerce/cmd"

//"golang.org/x/text/number"


func main() {
	cmd.Serve()

	//data := []byte("Hello") 
    // s := "HelloWorld"
	// data := []byte(s)
	// fmt.Println(s)
	// fmt.Println(data)
    
	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// b64Str := enc.EncodeToString(data)

	// fmt.Println(b64Str)

    // dec , err := enc.DecodeString(b64Str)
    // if err != nil {
	// 	fmt.Println(err) 
	// 	return
	// }
	// fmt.Println(dec)

	//SHA-256
    // hash := sha256.Sum256(data)
	// fmt.Println(hash)




	// key := []byte("a4bc166839d6ff3c83eb9d1cbb0ddda2a65f74eeea9a07413357200dfad1d808d708e7d5")
	// message := []byte("Hello World") 
    

	// h := hmac.New(sha256.New , key) 
	// h.Write(message) 

	// mac := h.Sum(nil) 

	// // fmt.Printf("HMAC: %x\n",mac)

	// message2 := []byte("Hello World") 
    
	// d := hmac.New(sha256.New, key)
	// d.Write(message2)
	// calculatedMAC := d.Sum(nil)
    

	// if hmac.Equal(calculatedMAC , mac) {
    //    fmt.Println("HMAC is valid")
	// }else {
    //    fmt.Println("HMAC is NOT valid!")
    // }

}