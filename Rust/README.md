```
docker build -t rustapp .
docker run -it -d --name rustServer -p 8000:8000 rustapp
```