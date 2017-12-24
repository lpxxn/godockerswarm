# Dockerfile

```
docker build . -t goweb:1.0

docker run -p 8001:8000 de85
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
docker service ps mygoweb_web
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



