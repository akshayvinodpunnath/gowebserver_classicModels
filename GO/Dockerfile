FROM golang

#setting work directory
WORKDIR $GOPATH/src/github.com/akshayvinodpunnath/webserver

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

#RUN apt-get update && apt-get install -y git curl libmcrypt-dev default-mysql-client

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN ls -lrt $GOPATH/src/github.com/akshayvinodpunnath/webserver

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["webserver"]
