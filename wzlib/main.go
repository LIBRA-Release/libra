package main

import (
	util "github.com/wzlib/wzutil"
	"C"
)

func init(){

	util.Init("agent",true)
}

//export ValidateToken
func ValidateToken(_om string) (*C.char){
	
	uid,err:=util.ValidateTokenExp(_om)
	if err!=nil{
		util.ErrLog(err)
		return nil
	}
	return C.CString(uid)
}

func main() {}
