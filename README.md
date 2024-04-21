# url.ly - a url-shortener service

This project aims to implement a decoupled & scalable architecture for url shortener service.

# Architecture

This project uses:
- `Nextjs` as its frontend;
- `Golang` & `Expressjs` as the backend servers configured with Nginx for load-balancing;
- `MongoDB` as non-relational db for authentication
- `Redis` as in-memory key-value db for url keys management

- For scalability and portability, both the frontend and backend are containerized using Docker and can be spun-up easily using docker-compose and can be orchestrated using the Kubernetes (the manifest files have been provided).
- Terraform is used for Infrastructure creation and Ansible playbooks and roles are used for configuring EC2 instances created on AWS providing:
    - Installation of Docker on those instances
    - Running of the url-shortener service through user input


