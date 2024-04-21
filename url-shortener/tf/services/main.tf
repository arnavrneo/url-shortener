terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  profile = "arnav" # enter relevant aws profile
}

# security group creation
resource "aws_security_group" "urlSG" {
  name = "urlshortener-sg"

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port = 22
    to_port = 22
    protocol = "tcp"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port = 3000
    to_port = 3000
    protocol = "tcp"
  }

  tags = {
    Name = "proj-sg"
  }
}

# ec2 instance creation
resource "aws_instance" "urlShortener" {
  ami = "ami-080e1f13689e07408"
  instance_type = "t2.micro"
  vpc_security_group_ids = [aws_security_group.urlSG.id]
}

resource "aws_ec2_tag" "ec2_tag" {
  key         = "Ansible"
  resource_id = aws_instance.urlShortener.id
  value       = "urlShortenerServer"
}


data "aws_caller_identity" "current" {}

output "caller_user" {
  description = "AWS Calling user: "
  value = data.aws_caller_identity.current.user_id
}

output "public_ip" {
  description = "Public IP: "
  value = aws_instance.urlShortener.public_ip
}

output "public_dns" {
  description = "Public DNS: "
  value = aws_instance.urlShortener.public_dns
}