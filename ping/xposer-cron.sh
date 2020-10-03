until ~/ping -j ~/xposer.jwt -u $1 >> ~/ping.log; do
  echo "'ping' crashed with exit code $?. Respawning process..." | tee -a ~/ping.log >&2
  sleep 1
done
