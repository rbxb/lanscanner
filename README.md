# lanscanner
Scan local network for HTTP servers.

```shell
$ go get github.com/rbxb/lanscanner
$ go install github.com/rbxb/lanscanner/cmd/lanscanner
```

#### `-ip`
List/range of IPs. (*required*)

#### `-port`
List/range of ports. (*required*)

#### `-rip`
Range IPs? (true)

#### `-rport`
Range ports? (false)

#### `-savepath`
Responses save path. (./responses)

#### `-delay`
Request delay duration in milliseconds. (100)

#### `-timeout`
Request timeout duration in seconds. (4)

## Usage

```shell
$ lanscanner -ip 192.168.1.0 -ip 192.168.1.20 -port 8000 -port 8080
```
This will scan all IPs in the range of `192.168.1.0` to `192.168.1.20` inclusively and check only the ports `8000` and `8080` on each IP.  
The responses will be saved in `./responses`.  

To disable IP ranging, use the `-rip` flag.  
```shell
$ lanscanner -ip 192.168.1.0 -ip 192.168.1.20 -port 8000 -port 8080 -rip=false
```
This will scan only the IPs `192.168.1.0` and `192.168.1.20`.  