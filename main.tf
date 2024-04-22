terraform {

  # terraform_cloudからなんやかんやするように設定
  # cloud {
  #   organization = "aki"
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "aki"

    workspaces {
      name = "go-todo-app"
    }
  }

  # awsでなんやかんやしたいので、awsのプロバイダを設定
  # https://registry.terraform.io/providers/hashicorp/aws/latest/docs
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# AWSプロバイダーの設定
provider "aws" {
  region = "ap-northeast-1"
  # localで実行する場合は、シークレットキーなど設定する必要がある

  # Assuming an IAM Role の設定
  # https://registry.terraform.io/providers/hashicorp/aws/latest/docs#assuming-an-iam-role
  # provider "aws" {
  #   assume_role {
  #     role_arn     = "arn:aws:iam::123456789012:role/ROLE_NAME"
  #     session_name = "SESSION_NAME"
  #     external_id  = "EXTERNAL_ID"
  #   }
  # }
}

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}
