# About playbooks

### `1_create_ec2.yaml`

Creates a security group & an ubuntu-ami backed EC2 instance using the specified properties.

> [!NOTE]
> Enter that `aws_profile` which is configured using the aws cli.

### `2_configure_ec2.yaml`

Updates the machine as well as installs the docker required in the next step for running the application.


> [!IMPORTANT]
> The above two playbooks are now replaced by the `terraform`. Either run the above playbooks or use the provided terraform code for infrastructure setup.

### `3_docker_book.yaml`

This provides functionalities for setting up the application files, running the application in the docker as well as the ability to remove the containers on the remote servers.


# Pre-requisites

- Change the `profile` inside the `inventory/aws_ec2.yaml` to the one that has been configured in system's environment or using the aws cli.
- Make sure to have the relevant `key pair` (here `ansible.pem`) for the ec2 instance to be present inside the `group_vars` directory as well as the `playbook` root directory in order for the playbook to configure the instances.
