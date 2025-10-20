package main

import (
	"ecommerce/cmd"
	"log/slog"
	"ecommerce/logger"
)

func main() {
   logs := logger.SetupLogger()
   slog.SetDefault(logs)
   slog.Info("Application Started")
   cmd.Serve()
}



////////////////////////////////////////////////////////////////////////////////////////////
    
// import "fmt"

// // "bytes"
// // "encoding/base64"
// //"encoding/base64"
// // "crypto/hmac"
// // "crypto/sha256"
// // "fmt"

// //"golang.org/x/text/message"

// //import	"ecommerce/cmd"

// //"golang.org/x/text/number"


// // Struct can implement interface
// // Sudu Darona (Pure Abustract Class)
// // Only Signature of a function
// type User interface {
//     PrintDetails()
// 	ReceiveAmmount(ammount float64) float64
// }

// type BankUser interface {
// 	WithDrawMoney(ammout float64) float64
// }

// type user struct {
//     Name  string 
// 	Age   int 
// 	Money float64 
// }

// // receiver function = method
// func (obj user) PrintDetails() {
// 	fmt.Println("Name ",obj.Name) 
// 	fmt.Println("Age ",obj.Age)
// 	fmt.Println("Money",obj.Money)
// }

// func (obj user) ReceiveAmmount(ammount float64) float64 {
// 	obj.Money = obj.Money + ammount
// 	return obj.Money
// }

// func (obj user) WithDrawMoney(ammout float64) float64 {
// 	obj.Money = obj.Money - ammout 
// 	return obj.Money
// }
   
//  var usr1 User = user {
	// 	Name: "Rupak Biswas",
	// 	Age: 26,
	// 	Money: 100.10,
	// }
	// usr1.PrintDetails()
	// usr1.ReceiveAmmount(100)
	// usr1.PrintDetails()

	// var usr2 BankUser = user{
	// 	Name: "Pronob Bala",
	// 	Age: 26,
	// 	Money: 100.10,
	// }

	// usr2.WithDrawMoney(90)

	// obj , flag := usr2.(user) 
	// if flag {
	// 	fmt.Println(obj)
	// } else {
	// 	fmt.Println("VALO NA")
	// }
	    


    // data := []byte("Hello") 
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

/////////////////////////////////////////////////////////////////////////////////////////////////