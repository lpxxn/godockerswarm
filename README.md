# Dockerfile

```
docker build . -t goweb:1.0

docker run -p 8100:8000 [image id] 
```

# Tag image to your private registry

private registry tutorials
```
http://www.cnblogs.com/li-peng/p/6511331.html
```

my private register is lpxxn.com:5000

```
docker tag goweb:1.0 lpxxn.com:5000/goweb:1.0

docker push lpxxn.com:5000/goweb:1.0
```

# before we can use the docker stack deploy command we'll first run:
```
docker swarm init
```

# docker deploy
```
docker stack deploy -c docker-compose.yml mygoweb
```

# Get service ID
```
docker service ls
```
# list your service by name 
```
docker service ps mygoweb
```

# Scale the app
you can scale the app by changing the replicas value in docker-compose.yml, save the change, and re-running the docker stack deploy command


```
docker stack deploy -c docker-compose.yml mygoweb
```

## take down app and swarm
- take the app down 
```
docker stack rm mygoweb
```

## take down the swarm.
```
docker swarm leave --force
```

