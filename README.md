# BSN Hub Go SDK

BSN Hub GO SDK makes a simple package of API provided by BSN Hub, which provides great convenience for users to quickly develop applications based on BSN Hub.

## install

### Requirement

- Go version above 1.14.0

### Use Go Mod

```text
require (
    github.com/bianjieai/bsnhub-sdk-go latest
)
```

## Usage

### Init Client

The initialization SDK code is as follows:

```go
import (
    "github.com/bianjieai/irita-sdk-go/types"
    "github.com/bianjieai/irita-sdk-go/types/store"
    oraclesdk "github.com/bianjieai/bsnhub-sdk-go"
    ...
)

options := []types.Option{
    types.KeyDAOOption(store.NewMemory(nil)),
    types.TimeoutOption(10),
}

cfg, err := types.NewClientConfig(nodeURI, chainID, options...)
if err != nil {
    panic(err)
}

client := oraclesdk.NewIRITAClient(cfg)
oracleClient := oraclesdk.OracleClient(client)

...
```

For more usage, please refer to [irita-sdk-go](https://github.com/bianjieai/irita-sdk-go/blob/master/README.md).
