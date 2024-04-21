# Frontend-Servers

This directory contains the `docker-compose` file for running up both the frontend, the backend and the proxy server containers all at once.

The container images are available on docker-hub, so the `compose.yaml` can be edited to pull the image instead of building them locally. The respective links are:
- <a href="https://hub.docker.com/r/arnavneo/nextjs-frontend">Frontend Image</a>)
- <a href="https://hub.docker.com/r/arnavneo/ejserver">Express Image</a>
- <a href="https://hub.docker.com/r/arnavneo/nginx-proxy">Nginx Proxy</a>
- <a href="https://hub.docker.com/r/arnavneo/goserver">Go Image</a>
