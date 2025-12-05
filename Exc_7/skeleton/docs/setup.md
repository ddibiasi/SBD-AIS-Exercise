# Setup Docker Swarm 

## Creating a Docker Swarm 

On the manager node:
```docker swarm init --advertise-addr <ipv4-address>```

On the other nodes:
```docker swarm join --token <token>```

To gather more information about the swarm one can run: 
* ```docker info```: gives a long list of information, look for ```Swarm: active```
* ```docker node ls```: lists all nodes connected to the swarm, ```*``` indicates the node you are connected to

## Deploying the stack to the swarm 

All the commands were executed on the manager node. 

Because a swarm consists of multiple Docker Engines, a registry is required to distribute images to all of them.
With the following command a registry is started as a service on the swarm: 
```docker service create --name registry --publish published=5000,target=5000 registry:2```.

The service can be deleted again with ```docker service rm registry```.
To see a list of all running services type ```docker service ls```.

Then, I tested the application with ```docker compose up``` and stopped/removed it again with ```docker compose down```.

To distribute the app's image across the swarm, it needs to be pushed to the registry with this command: ```docker compose push```.

To create the stack with the name *order* I executed the following: ```docker stack deploy --compose-file docker-compose.yml order```

With ```docker stack services order``` you can check whether it's running and with ```docker stack rm order``` you can delete the stack again.