```
 █     █░ ▄▄▄       ██ ▄█▀▓█████  █    ██  ██▓███  
▓█░ █ ░█░▒████▄     ██▄█▒ ▓█   ▀  ██  ▓██▒▓██░  ██▒
▒█░ █ ░█ ▒██  ▀█▄  ▓███▄░ ▒███   ▓██  ▒██░▓██░ ██▓▒
░█░ █ ░█ ░██▄▄▄▄██ ▓██ █▄ ▒▓█  ▄ ▓▓█  ░██░▒██▄█▓▒ ▒
░░██▒██▓  ▓█   ▓██▒▒██▒ █▄░▒████▒▒▒█████▓ ▒██▒ ░  ░
░ ▓░▒ ▒   ▒▒   ▓▒█░▒ ▒▒ ▓▒░░ ▒░ ░░▒▓▒ ▒ ▒ ▒▓▒░ ░  ░
  ▒ ░ ░    ▒   ▒▒ ░░ ░▒ ▒░ ░ ░  ░░░▒░ ░ ░ ░▒ ░     
  ░   ░    ░   ▒   ░ ░░ ░    ░    ░░░ ░ ░ ░░       
    ░          ░  ░░  ░      ░  ░   ░              
```    
# WakeUp

Wakeup is a Go Library for determining if the host is a VM. 

## Installation

```bash
go get github.com/RickConsole/wakeup
```

## Example

```go
package main

import "github.com/RickConsole/wakeup"

func main() {
  wakeup.Check() //Returns true (for vm) or false (for real machine)
}
```

## Extra Functions

```go
wakeup.GetMAC() //Gets Victims MAC Address
```
```go
wakeup.GetVendor() //Acquires Vendor from MAC Address Prefix
```
