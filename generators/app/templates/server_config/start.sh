sudo rm -R models
echo "Stoping service"
sudo systemctl stop goweb2
echo "Service Stoped"
go run init/init.go
echo "Building "
go build main.go
echo "Build completed"
sudo systemctl start goweb2
echo "Service started"