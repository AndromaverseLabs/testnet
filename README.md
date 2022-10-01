# Androma testnet
**androma** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli) and the simple command "ignite scaffold chain androma --address-prefix andr"

## Get started
Make sure your system is properly updated
```
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get install -y build-essential curl wget jq
sudo su -c "echo 'fs.file-max = 65536' >> /etc/sysctl.conf"
sudo sysctl -p
```

Install go
```
# First remove any existing old Go installation
sudo rm -rf /usr/local/go

# Install the latest version of Go using this helpful script 
curl https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash

sudo mv $HOME/.go /usr/local/

# Update environment variables to include go
cat <<'EOF' >>$HOME/.profile
export GOROOT=/usr/local/.go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/.go/bin:$HOME/go/bin
EOF
source $HOME/.profile

go version
# Should return go version go1.16.4 linux/amd64
```

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
Monitor using:
```
sudo journalctl -u andromad -f
```
