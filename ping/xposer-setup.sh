echo "'ping' is being updated. Rebooting instance..." >> ~/ping.log

echo "@reboot ~/xposer-cron.sh $1" > tmpcron
crontab tmpcron
rm tmpcron
