# Frontend-Servers

This directory contains the `docker-compose` file for running up both the frontend, the backend and the proxy server containers all at once.

The container images are available on docker-hub, so the `compose.yaml` can be edited to pull the image instead of creating them locally. The respective links are:
- ![Frontend](https://hub.docker.com/r/arnavneo/nextjs-frontend)
- ![Express Backend](https://hub.docker.com/r/arnavneo/ejserver)
- ![Nginx Proxy](https://hub.docker.com/r/arnavneo/nginx-proxy)
- ![Go Backend](https://hub.docker.com/r/arnavneo/goserver)
