#builder
FROM golang:1.25 AS builder
WORKDIR /app

# copi the proyect to the container
COPY . . 

#execute the build script that generate /app/ordersystem
RUN chmod +x /app/scripts/build-application.sh && sh /app/scripts/build-application.sh


#Runtime
FROM alpine:latest

#copy the complied binary from the builder stage into the container
COPY --from=builder /app/ordersystem /ordersystem

#expose the port 3000 so the api can be accessed from outside the container
EXPOSE 3000

#default comand when the container starts
ENTRYPOINT ["/ordersystem"]
