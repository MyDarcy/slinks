port=9090
pid=$(netstat -nlp | grep :$port | awk '{print $7}' | awk -F"/" '{ print $1 }');
if [  -n  "$pid"  ];  then
    kill  -9  $pid;
fi

cd /home/golang/src/slinks
go build
nohup ./slinks &> allStandard.log&
