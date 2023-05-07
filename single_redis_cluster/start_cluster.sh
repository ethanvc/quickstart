#!/bin/bash

GetConf() {
  local port_number="$1"
  cat <<EOF
port $port_number
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
appendonly yes
EOF
}

RunCommandWithRetry(){
  command_to_run=$1
  max_attempts=$2
  sleep_duration=$3

  attempt=1
  while [ $attempt -le $max_attempts ]; do
      echo "Attempt $attempt of $max_attempts for command $command_to_run"
      $command_to_run

      exit_code=$?
      if [ $exit_code -eq 0 ]; then
          echo "Command succeeded!"
          break
      else
          echo "ERR: Command failed with exit code: $exit_code"
          if [ $attempt -lt $max_attempts ]; then
              echo "Sleeping for $sleep_duration seconds before trying again..."
              sleep $sleep_duration
          fi
      fi

      attempt=$((attempt + 1))
  done

  if [ $attempt -gt $max_attempts ]; then
      echo "ERR: Reached maximum number of attempts. Exiting."
  fi
}

main() {
  mkdir /etc/redis
  local svr_list=""
  for i in $(seq 1 ${CLUSTER_INSTANCE_COUNT}); do
    local filename="/etc/redis/redis_${i}.conf"
    svr_list="$svr_list 127.0.0.1:$((i + 5999))"
    echo "Creating $filename..."
    GetConf $((i + 5999)) >"$filename"
    local working_dir="/data/redis/redis_${i}"
    mkdir -p "$working_dir"
    ( cd "$working_dir" && redis-server "$filename" > "/tmp/redis_svr_log_${i}.log" & )
  done

  RunCommandWithRetry "redis-cli --cluster create $svr_list --cluster-replicas 1 --cluster-yes" 10 5

  # pause forever
  tail -f /dev/null
}

main
