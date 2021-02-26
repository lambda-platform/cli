sudo rm -R models
echo "Stoping service"
sudo systemctl stop goweb2
echo "Service Stoped"
sudo rm -R ./app/models/form
sudo rm -R ./app/models/grid
echo "INIT START"
sudo mkdir ./app/models/form
sudo mkdir ./app/models/grid
sudo chmod -R 755 ./app/models/form
sudo chmod -R 755 ./app/models/grid
echo "Ready"
echo "Building "
go build main.go
echo "Build completed"
sudo systemctl start goweb2
echo "Service started"