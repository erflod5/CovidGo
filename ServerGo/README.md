```
docker build -t servergo .
docker run -d --name appGo -p 8002:8002 servergo
```