sudo service goweb stop
sudo rm -R ./app/graph
sudo rm -R ./app/models/form
sudo rm -R ./app/models/grid
echo "INIT START"
sudo mkdir ./app/graph
sudo mkdir ./app/models/form
sudo mkdir ./app/models/grid
sudo chmod -R 777 ./app/graph
sudo chmod -R 777 ./app/models/form
sudo chmod -R 777 ./app/models/grid
go run ./bootstrap/init/init.go
sudo chmod -R 777 ./app/graph
sudo chmod -R 777 ./app/models/form
sudo chmod -R 777 ./app/models/grid
go build main.go
sudo service goweb start
sudo chmod -R 755 ./app/graph
sudo chmod -R 755 ./app/models/form
sudo chmod -R 755 ./app/models/grid
echo "Service started"
