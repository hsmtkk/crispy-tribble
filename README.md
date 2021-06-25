# crispy-tribble
dhcp client inform

# memo

got error

```shell
$ go run . イーサネット
# github.com/insomniacslk/dhcp/dhcpv4/client4
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:81:13: undefined: unix.Socket
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:81:25: undefined: unix.AF_INET
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:81:39: undefined: unix.SOCK_RAW
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:81:54: undefined: unix.IPPROTO_RAW
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:85:8: undefined: unix.SetsockoptInt
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:85:31: undefined: unix.SOL_SOCKET
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:85:48: undefined: unix.SO_REUSEADDR
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:89:8: undefined: unix.SetsockoptInt
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:89:31: undefined: unix.IPPROTO_IP
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:89:48: undefined: unix.IP_HDRINCL
..\..\..\go\pkg\mod\github.com\insomniacslk\dhcp@v0.0.0-20210621130208-1cac67f12b1e\dhcpv4\client4\client.go:89:48: too many errors
```
