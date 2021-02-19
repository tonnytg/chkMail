## Go Module Valid Email

This is a simple module for Go to validate a Email regex.

### How install?

`go get github.com/tonnytg/gckmail`


### Doc
`func ChkReg(n string) bool`


### Example

Create a simple project and import this module:
```
package main  
  
import (  
   "fmt"  
 "github.com/tonnytg/gckmail")  
  
func main()  {  
   mail := "11111111111111" // invalid mail
  if m := chkmail.ChkReg(mail); m {  
      fmt.Printf("Email %v, it is valid!", mail)  
   } else {  
      fmt.Println("Email invalid")  
   }  
}
```