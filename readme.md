## Go Module Valid Email

This is a simple module for Go to validate a Email regex.

### How install?

`go get github.com/tonnytg/chkmail`


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

### Benchmark

go test -bench=.

    goos: darwin
    goarch: arm64
    pkg: github.com/tonnytg/chkmail
    BenchmarkChkReg-16                171973              7017 ns/op
    BenchmarkChkRegMultiple-16         58987             20420 ns/op
    PASS
    ok      github.com/tonnytg/chkmail      4.061s
