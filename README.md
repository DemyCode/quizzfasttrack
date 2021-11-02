# Project FastTrack GoLang

This project is a simple quizz that is implemented with a REST API.
For a simple test you can run the following command :

go run client/main.go questions

This command will fetch the questions from the server 34.140.239.165:8080
which is currently hosted on Google Cloud. (At least I hope it's still running)

To go more in depth you can change in config.go the variable HOST to "localhost"
then run in another window

go run server/main.go

## AUTHORS

Mehdi Bekhtaoui
mehdi.bekhtaoui@epita.fr
github.com/DemyCode

Under the supervision of :
Patrick Potocki
patrik@fasttrack-solutions.com

## MAINTENANCE

This project won't be maintained in the future, feel free to make changes as you please