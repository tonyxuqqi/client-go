module rawkv

go 1.15

require github.com/tikv/client-go/v2 v2.0.0

replace github.com/tikv/client-go/v2 => ../../

replace github.com/pingcap/kvproto v0.0.0-20210806074406-317f69fb54b4 => github.com/tonyxuqqi/kvproto v0.0.0-20210908222950-9062956df1a6
