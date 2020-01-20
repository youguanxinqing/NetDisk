#bash

# developt cmd: 
# when-changed -r -v . sh start.sh 8080

# port
port=$1
# kill process
lsof -i tcp:$port | grep main | awk '{print "kill -9 " $2}' | sh 
# running
go run main.go &