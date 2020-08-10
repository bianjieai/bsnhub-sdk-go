module github.com/bianjieai/bsnhub-sdk-go

go 1.14

require (
	github.com/bianjieai/irita-sdk-go v0.0.0-20200810074639-cce20cfcf98a
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.0
	github.com/tendermint/tendermint v0.33.4
)

replace github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703
