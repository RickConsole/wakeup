
     _---~~(~~-_.
    _{        )   )
  ,   ) -~~- ( ,-' )_
 (  `-,_..`., )-- '_,)
( ` _)  (  -~( -_ `,  }
(_-  _  ~_-~~~~`,  ,' )
  `~ -^(    __;-,((()))
        ~~~~ {_ -_(())
               `\  }
                 { }   Wakeup
   
# WakeUp

Wakeup is a Go Library for determining if the host is a VM. 

## Installation

```bash
go get github.com/RIckConsole/wakeup
```

## Example

```go
package main

import "github.com/RickConsole/wakeup"

func main() {
  wakeup.CheckUsingSysinfo() //Returns true (for real machine) or false (for VM)
}
```

## Extra Functions

```go
wakeup.GetMAC() //Gets Victims MAC Address
```
```go
wakeup.GetVendor() //Acquires Vendor from MAC Address Prefix
```
```go
wakeup.CheckUsingMAC() 
//Gets client MAC, requests vendor name using mac from maccaddress.io, if its a vm vendor, will return false
```
