## Get started
Make sure your system is properly updated
```
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get install -y build-essential curl wget jq
sudo su -c "echo 'fs.file-max = 65536' >> /etc/sysctl.conf"
sudo sysctl -p
```

Install toolchain and ensure accurate time synchronization 
```
sudo apt-get install make build-essential gcc git jq chrony -y
```

## Install go
First remove any existing old Go installation
```
sudo rm -rf /usr/local/go
```

Install the latest version of Go using this helpful script 
```
curl https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash
```
Make go available for your complete system
```
sudo mv $HOME/.go /usr/local/
```

Update environment variables to include go
```
cat <<'EOF' >>$HOME/.profile
export GOROOT=/usr/local/.go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/.go/bin:$HOME/go/bin
EOF
source $HOME/.profile
```

Check if everything went ok
```
go version
```
Should return go version go1.16.4 linux/amd64

## Install Ignite
Install the latest version of ignite
```
git clone https://github.com/ignite/cli --depth=1
cd cli 
make install
```

## Create a copy of the github repo
```
git clone https://github.com/AndromaverseLabs/testnet
cd testnet
cd Chain
```

## Build the binary
```
ignite chain build
```

## Prepare the node
```
andromad config chain-id androma-1
andromad init [moniker] --chain-id androma-1
andromad keys add [walletname] 
OR if you already have a wallet
andromad keys add [walletname] --recover
curl https://raw.githubusercontent.com/AndromaverseLabs/testnet/main/genesis.json > ~/.androma/config/genesis.json
sudo ufw allow 26656
sudo tee /etc/systemd/system/andromad.service > /dev/null <<EOF
[Unit]
Description=Androma Daemon
After=network-online.target
[Service]
User=$USER
ExecStart=$(which andromad) start
Restart=always
RestartSec=3
LimitNOFILE=65536
[Install]
WantedBy=multi-user.target
EOF
sudo mv /etc/systemd/system/andromad.service /lib/systemd/system/
sudo systemctl daemon-reload
sudo -S systemctl enable andromad
sudo service andromad start
```
## Check node info:
```
curl -s localhost:26657/status | jq .result.sync_info.catching_up
#true output is syncing - false is synced
curl -s localhost:26657/status | jq .result.sync_info.latest_block_height
#this output is your last block synced
curl -s "http://:26657/status?" | jq .result.sync_info.latest_block_height
#this output the public node last block synced
```
## Monitor using:
```
sudo journalctl -u andromad -f
```
## Get Testnet Tokens:
```
Seek out a validator role on the discord to get access to the testnet faucet
https://linktr.ee/andromaverse
```
## Create Validator:
```
andromad tx staking create-validator \
 --amount=10andr \
 --pubkey=$(andromad tendermint show-validator) \
 --moniker=<your-moniker> \
 --chain-id=test-chain-androma-1 \
 --commission-rate="0.05" \
 --commission-max-rate="0.10" \
 --commission-max-change-rate="0.05" \
 --min-self-delegation="1" \
 --gas=auto \
 --gas-adjustment=1.5 \ 
 --gas-prices=0.025andr \
 --from=(key of your wallet address)
```

