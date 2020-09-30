if [ "$1" = "" ]; then
  echo "usage: ./deploy-pong.sh <user>@<host> <public key path>"
fi
ssh $1 "rm ~/pong; exit;"
scp ./bin/pong ./pong/conf.json ./pong/xposer-cron.sh ./pong/xposer-setup.sh $1:.
scp $2 $1:./key.pub
ssh $1 "./xposer-setup.sh; sleep 5 && sudo reboot; exit;"
