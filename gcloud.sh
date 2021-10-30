# Script typed in Google Cloud startup script

sudo add-apt-repository ppa:longsleep/golang-backports -y
sudo apt update -y
sudo apt install golang-go -y
sudo apt install git -y

cd /home/verycols
git clone https://github.com/DemyCode/quizzfasttrack.git
cd quizzfasttrack
go build server/main.go
./main