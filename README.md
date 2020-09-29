# Xfinity Xposer

ssh-keygen -t rsa -b 4096 -m PEM -f key.pem
openssl rsa -in key.pem -pubout -outform PEM -out key.pub

e.g.

```
./bin/keygen -d raspberry -k id_garymenezes.pem

./bin/pong -k id_garymenezes.pub

./bin/ping -j bin/xposer.jwt -u http://localhost:8080/ping
```
