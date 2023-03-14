# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18.1-alpine

# Creates an app directory to hold your app’s source code
# WORKDIR /app

# Copies everything from your root directory into /app
# COPY . .

#COPY go.mod ./
#COPY go.sum ./

# Installs Go dependencies
# RUN go mod download

# Tells Docker which network port your container listens on
EXPOSE 8080

# Specifies the executable command that runs when the container starts
# CMD [ “/godocker” ]