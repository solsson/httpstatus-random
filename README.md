# httpstatus-random

Web server that sits on port 8080 and responds with random status codes.

https://hub.docker.com/r/solsson/httpstatus-random/

```
docker build -t httpstatus-random .
docker run -d --rm -p 8080:8080 --name httpstatus-random httpstatus-random
curl http://localhost:8080/[1-20] -I
docker kill httpstatus-random
```
