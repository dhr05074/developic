locals {
  front_usernames = ["jaehwan"]
}

module "front_user" {
  count = length(local.front_usernames)
  source  = "terraform-aws-modules/iam/aws//modules/iam-user"

  create_iam_access_key = false
  create_iam_user_login_profile = false

  name = local.front_usernames[count.index]
}

module "front_group" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-group-with-policies"

  name = "front"

  group_users = module.front_user[*].iam_user_name

  attach_iam_self_management_policy = true

  custom_group_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonS3FullAccess",
    "arn:aws:iam::aws:policy/CloudFrontFullAccess",
  ]
}