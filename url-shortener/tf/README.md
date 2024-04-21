# Infrastructure Creation

The `services` directory contains the terraform code for creating the project Infrastructure on the AWS. It creates:
- the security group that allows access on the following ports:
    - `22`: for `ssh` access
    - `3000`: the server where the frontend will be running

- EC2 instance with the above security group attached.
