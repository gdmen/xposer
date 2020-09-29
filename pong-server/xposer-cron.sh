until ~/pong -k ~/key.pub > ~/pong.log; do
  echo "Server 'pong' crashed with exit code $?. Respawning process..." >&2
  sleep 1
done
