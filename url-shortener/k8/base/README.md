# K8 Setup

This directory aims for the container orchestration of frontend, backend server and nginx as the reverse proxy between the formers. The `frontend` and `servers` directories contain the required manifest files for setting up the project on a kubernetes cluster (this project was created on minikube).

## Pre-requisites

The specified secrets (env variables) values need to be substituted in the target backend server secrets manifest files under `secrets` directory for the backend to work.

## Order of setup 

Let's choose expressjs as our backend server.

1) Create all the services under `servers/services/` directory and then proceed forward.

### nginx setup

2) Create the `nginx-deploy.yaml` deployment for the nginx reverse proxy.
3) Take note of the `nginx service` external ip address.

### frontend setup

4) Now, its turn to deploy our frontend. Before that, edit the `frontend-deploy.yaml` and add the `nginx service` endpoint as an env value for `NEXT_PUBLIC_BACKEND_URL` and suffixing the endpoint with "/api". For example, if the `nginx service` endpoint is `http://192.168.58.2:32015`, the `NEXT_PUBLIC_BACKEND_URL` will be `http://192.168.58.2:32015/api`. Then create the `frontend-deploy.yaml` deployment.
5) Create the `frontend-service.yaml` and note the `frontend service` ip address.
6) `sh` into `nginx-pod`. Edit the `default.conf` and replace the `ORIGIN` header by the `frontend-service` endpoint (without `/`). Then reload nginx.

### backend setup

7) Edit the `ejs-secret.yaml`, add the specified secret values in base64 format. For `ORIGINS` field,
   use the `base64` value of `frontend service` ip (without `/`) and for `NGINX_ENDPOINT`, use the value of `nginx service` ip (without `/`) in `base64` format and
   then create the secret.
8) Create the `ejs-deploy.yaml` (edit the `REDIS_PORT` env value according to your redis db port.)

The frontend can be accessed and nginx reverse proxies the requests to the backend. The k8 cluster is now configured.

