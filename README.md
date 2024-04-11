# url.ly - a url-shortener service

This project aims to implement a decoupled & scalable architecture for url shortener service.

# Architecture

This project uses:
- `Nextjs` as its frontend;
- `Golang` & `Expressjs` as the backend servers configured with Nginx for load-balancing;
- `MongoDB` as non-relational db for authentication
- `Redis` as in-memory key-value db for url keys management

- For scalability and portability, both the frontend and backend are containerized using Docker and can be spun-up easily using docker-compose.
- Ansible playbooks and roles are used for automating the service setup & infrastructure creation of single (or multiple) EC2 on AWS. They are used for:
    - Creation of a security group (with relevant properties) for EC2 instances
    - Creation of EC2 instances & attaching of the defined security group
    - Installation of Docker on those instances
    - Running of the url-shortener service through user input
