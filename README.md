# taraxa-contracts-go-clients
Taraxa contracts go clients

# Usage
```
```

See demo.go for more examples


#### Prerequisites
##### solc
```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

##### abigen
```
go install github.com/ethereum/go-ethereum@latest
cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum@v<X.Y.Z>/
make
make devtools
```

## !!! To work with latest contracts interfaces
### taraxa-evm
Update taraxa-evm submodule by running:
```
git submodule update submodules/taraxa-evm
```

or update all submodules:

```
git submodule update --init --recursive
```

Generate abi & client:
```
solc --abi --overwrite --optimize submodules/taraxa-evm/taraxa/state/contracts/dpos/solidity/dpos_contract_interface.sol --output-dir tara/dpos_contract_client/dpos_interface/

abigen --abi=tara/dpos_contract_client/dpos_interface/DposInterface.abi --pkg=dpos_interface --out=tara/dpos_contract_client/dpos_interface/dpos_interface.go
```

### bridge
Update bridge submodule by running:
```
git submodule update submodules/bridge
```

or update all submodules:

```
git submodule update --init --recursive
```

Generate abi & client:
```
solc --abi --overwrite --optimize submodules/taraxa-evm/taraxa/state/contracts/dpos/solidity/dpos_contract_interface.sol --output-dir tara/dpos_contract_client/dpos_interface/

abigen --abi=tara/dpos_contract_client/dpos_interface/DposInterface.abi --pkg=dpos_interface --out=tara/dpos_contract_client/dpos_interface/dpos_interface.go
```
