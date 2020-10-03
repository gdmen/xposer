if [ "$1" = "" ]; then
  echo "usage: ./deploy.sh <user>@<host> <jwt path>"
fi
ssh $1 "rm ~/ping; exit;"
scp ./bin/ping ./bin/xposer.jwt ./ping/xposer-cron.sh ./ping/xposer-setup.sh $1:.
ssh $1 "./xposer-setup.sh $2 && sudo reboot; exit;"
