# Script typed in Google Cloud startup script

sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go

git clone https://github.com/DemyCode/quizzfasttrack.git
cd quizzfasttrack
go build server/main.go
./main