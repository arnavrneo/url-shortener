# K8 Setup

`frontend` and `servers` directories contain the required manifest files for setting up the project on a kubernetes cluster.

## Pre-requisites

The specified secrets (env variables) values need to be substituted in the target backend server secrets manifest files under `secrets` directory.

## Order of setup 

Let's choose expressjs as our backend server (goserver under dev).

### nginx setup

1) Create the `nginx-deploy.yaml` deployment for the nginx reverse proxy.
2) Create the `nginx-service.yaml` to expose the nginx pod to external access.
3) Take note of the `nginx service` external ip address.

### frontend setup

4) Now, its turn to deploy our frontend. Before that, edit the `frontend-deploy.yaml` and add the `nginx service` endpoint as an env value for `NEXT_PUBLIC_BACKEND_URL` and suffixing the endpoint with "/api". For example, if the `nginx service` endpoint is `http://192.168.58.2:32015`, the `NEXT_PUBLIC_BACKEND_URL` will be `http://192.168.58.2:32015/api`. Then create the `frontend-deploy.yaml` deployment.
5) Expose the `frontend` for external access by creating `frontend-service.yaml` service and note the `frontend service` ip address.

### backend setup

6) Edit the `ejs-secret.yaml`, add the specified secret values in base64 format. For `ORIGINS` field,
   use the `base64` value of `frontend service` ip and for `NGINX_ENDPOINT`, use the value of `nginx service` ip in `base64` format and
   then create the secret.
7) Create the `ejs-service.yaml` as well as the `ejs-deploy.yaml` (edit the `REDIS_PORT` env value according to your redis db port.)

The frontend can be accessed and nginx reverse proxies the requests. The k8 cluster is now created.

