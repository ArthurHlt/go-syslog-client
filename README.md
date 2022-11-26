# go-syslog-client

A syslog client for golang which give possibility to send syslog message to a remote syslog server in multiple ways:
- HTTP
- UDP
- TCP
- TCP over TLS
- Multiples syslog servers at same time

## Usage

```go
package main

import (
    "github.com/ArthurHlt/go-syslog-client"
)

func main() {
    writer, err := syslog.NewWriter(
    // data will be sent in http to this url with POST params
	"https://my.http.server.com/receive",
	// data will be sent http with gzip compression
	"https://my.http.server.com/receive?in_gzip=true",
	// in tcp mode
	"tcp://my.syslog.server.com:514",
	// in tcp mode with a pool of 10 connections (there is no pool by default)
    "tcp://my.syslog.server.com:514?pool_size=10",
	// tcp over tls without verifying server certificate (not recommended)
	"tcp+tls://my.syslog.server.com:514?verify=false",
	// tcp over tls with verifying server certificate
	"tcp+tls://my.syslog.server.com:514?cert=path/to/a/ca/file",
	// in udp mode
    "udp://my.syslog.server.com:514",
    )
	if err != nil {
        panic(err)
    }
    // write a message
    writer.Write([]byte(`<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] BOMAn application event log entry...`))
    // close the writer
    writer.Close()
}
```
