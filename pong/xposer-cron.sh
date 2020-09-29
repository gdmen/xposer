until ~/pong -k ~/key.pub >> ~/pong.log; do
  echo "Server 'pong' crashed with exit code $?. Respawning process..." | tee -a ~/pong.log >&2
  sleep 1
done
