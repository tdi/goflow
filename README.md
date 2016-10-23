[![Build Status](https://travis-ci.org/tdi/goflow.svg?branch=master)](https://travis-ci.org/tdi/goflow)

# goflow

Version: 0.4

A simple NetFlow v5 checker.

## Installation

`go get github.com/tdi/goflow`

or you can use pre-built images from [gobuilder.me/github.com/tdi/goflow](https://gobuilder.me/github.com/tdi/goflow) page. 

## CLI usage

`goflow [-h] -H HOSTNAME -p PORT`

defaults: 127.0.0.1:2055 (UDP)

    Duration    SrcIP:SrcPORT           DstIP:DstPort        Proto   NPackets NOctets

     19s          199.96.57.6:443       192.168.1.31:37371   TCP     6      1982
     20s       192.168.1.31:50477        212.2.108.110:443   TCP     4       427
     19s          199.96.57.6:443       192.168.1.31:37372   TCP     6      1982
     20s       192.168.1.31:50474        212.2.108.110:443   TCP    10      1005
      7s       192.168.1.31:37380          199.96.57.6:443   TCP     2       104
      8s          199.96.57.6:443       192.168.1.31:37380   TCP     2       104
     19s       192.168.1.31:44605       173.194.112.75:443   TCP     2       112
     20s       192.168.1.31:48343        173.194.44.48:443   TCP     4       426
     20s       192.168.1.31:41577       173.194.112.76:443   TCP     7       592
     19s       192.168.1.31:37372          199.96.57.6:443   TCP     5       484
     20s       192.168.1.31:46295         212.2.108.88:443   TCP     4       427
     19s       192.168.1.31:37371          199.96.57.6:443   TCP     7       714



## AUTHOR

Copyright (c) Dariusz Dwornikowski




