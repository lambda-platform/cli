sudo rm -R ./app/models/form
sudo rm -R ./app/models/grid
echo "INIT START"
sudo mkdir ./app/models/form
sudo mkdir ./app/models/grid
sudo chmod -R 755 ./app/models/form
sudo chmod -R 755 ./app/models/grid
go run ./bootstrap/init/init.go
sudo chmod -R 755 ./app/models/form
sudo chmod -R 755 ./app/models/grid
echo "Ready"
fresh
