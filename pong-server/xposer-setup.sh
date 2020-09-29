echo "@reboot ~/xposer-cron.sh" >> tmpcron
crontab tmpcron
rm tmpcron
