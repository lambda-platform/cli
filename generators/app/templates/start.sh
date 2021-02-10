sudo rm -R ./models
echo "INIT START"
go run init/init.go
sudo chmod -R 777 ./models
echo "Ready"
fresh