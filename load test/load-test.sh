max="$1"
date
echo "url: $2
rate: $max calls / second"
START=$(date +%s);

get () {
  curl --http3 -s -w ",%{http_code}" "$1" 2>&1 | tr '\r\n' '\\n' | awk -v date="$(date +'%r')" '{print $0",", date}' >> perf-test.log
}

while true
do
  echo $(($(date +%s) - START)) | awk '{print int($1/60)":"int($1%60)}'
  sleep 1

  for i in `seq 1 $max`
  do
    get $2 &
    curl_pid=$!
    #echo -e "CURL PID = $curl_pid"
    kill -0 "$curl_pid"
  done
done
