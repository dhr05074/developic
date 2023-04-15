terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 0.14"
}

provider "aws" {
  region = "ap-northeast-2"
  profile = "side_project"
}

data "aws_caller_identity" "caller" {}