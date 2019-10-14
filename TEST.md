# Testing stuff

A sample corefile is located in `Corefile_autoipv6ptr`. The referenced presets file is located in `Presets_autoipv6ptr`.

Queries:

1) dig @::1 -p 533 -x 2001:DB8:300:b001:5054:ff:fe4b:db44 -> test.mydomain.tld
2) dig @::1 -p 533 -x 2001:DB8:300:b001:5054:ff:fe4b:db45 -> 20010db80300b001505400fffe4bdb45.servers.mydomain.tld
