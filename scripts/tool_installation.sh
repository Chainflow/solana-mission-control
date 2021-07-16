#!/bin/bash

set -e

cd $HOME

echo "--------- Cloning solana-monitoring-tool -----------"

git clone https://github.com/Chainflow/solana-mission-control.git

cd solana-mission-control

cp example.config.toml config.toml

cd $HOME

echo "------ Updatig config fields with exported values -------"


sed -i '/rpc_endpoint =/c\rpc_endpoint = "'"$RPC_ENDPOINT"'"' ~/solana-prometheus/config.toml

sed -i '/network_rpc =/c\network_rpc = "'"$NETWORK_RPC"'"' ~/solana-prometheus/config.toml

sed -i '/validator_name =/c\validator_name = "'"$VALIDATOR_NAME"'"'  ~/solana-prometheus/config.toml

sed -i '/pub_key =/c\pub_key = "'"$PUB_KEY"'"'  ~/solana-prometheus/config.toml

sed -i '/vote_key =/c\vote_key = "'"$VOTE_KEY"'"'  ~/solana-prometheus/config.toml

if [ ! -z "${TELEGRAM_CHAT_ID}" ] && [ ! -z "${TELEGRAM_BOT_TOKEN}" ];
then 
    sed -i '/tg_chat_id =/c\tg_chat_id = '"$TELEGRAM_CHAT_ID"''  ~/solana-prometheus/config.toml

    sed -i '/tg_bot_token =/c\tg_bot_token = "'"$TELEGRAM_BOT_TOKEN"'"'  ~/solana-prometheus/config.toml

    sed -i '/enable_telegram_alerts =/c\enable_telegram_alerts = 'true''  ~/solana-prometheus/config.toml
else
    echo "---- Telgram chat id and/or bot token are empty --------"
fi

echo "------ Building and running the code --------"

cd solana-mission-control

go build -o solana-mc
mv solana-mc $HOME/go/bin

echo "----------- Setupsolana-mc service------------"

echo "[Unit]
Description=Solana-mc
After=network-online.target

[Service]
User=$USER
ExecStart=$HOME/go/bin/solana-mc
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target" | sudo tee "/lib/systemd/system/solana_mc.service"

echo "--------------- Start Solana-Mession-Control service ----------------"


sudo systemctl daemon-reload

sudo systemctl enable solana_mc.service

sudo systemctl start solana_mc.service

echo "** Done **"