#/bin/bash
echo "####################################" >> lambda.log
echo "########Starting new build log#######" >> lambda.log
echo "####################################" >>  lambda.log
echo "Stop currently running !!!"
TIMESTAMP=`date "+%Y-%m-%d %H:%M:%S"`
build_cmd="go build -o lambda-starter"
delete_binary="pkill -9 lambda-starter"
echo "Stop currently running instances"
pkill -9 lambda-starter
echo "Lambda build process started!!"
echo "$TIMESTAMP Lambda build process started!!" >> lambda.log
go build -o lambda-starter >> lambda.log
echo "Set executable permission on compiled binary"
chmod +x lambda-starter
echo "running $build_cmd"
echo "$TIMESTAMP $build_cmd" >> lambda.log
echo "$TIMESTAMP Starting server" >> lambda.log
echo "$TIMESTAMP Starting server"
./lambda-starter & nohup >> lambda.log
