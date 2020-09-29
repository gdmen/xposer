echo "Server 'pong' is being updated. Rebooting instance..." >> ~/pong.log

echo "@reboot ~/xposer-cron.sh" > tmpcron
crontab tmpcron
rm tmpcron
