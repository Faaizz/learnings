# Networking Foundations

## Basics
Notes from Networking Foundations: Networking Basics video course on LinkedIn Learning.

### Network Addresses
- Media Access Control (MAC) Address: 48 bits long. 24 bits for vendor code, 24 bits for hardware id. 
  Typically written as 12 hexadecimal digits. E.g. 0e:98:4d:ac:4b:21.
- Internet Protocol (IP) v4 Address: 32 bits long. 
  Typically written as 4 decimal octets. E.g. 192.168.0.1
  Network address is separated from host address using subnet masks. 
  E.g. 192.168.0.0/24 indicates that the first 24 bits identify the network, leaving 8 bits for host addressing.
- Internet Protocol (IP) v6 Address: 128 bits long.
  Typically written as 32 hexadecimal quartets. E.g. 0b23:1e4a:9003:a34d:0000:0000:0000:0000. The zeros can be omitted.
  Network prefix (used to identify the network) is separated from host address.
  E.g. 0b23:1e4a:9003:a34d:0000:0000:0000:0000/64 indicates that the first 64 bits identify the network, leaving 64 bits for host addressing.


## References
- [Networking Foundations: Networking Basics](https://www.linkedin.com/learning/networking-foundations-networking-basics/network-interface-cards)
