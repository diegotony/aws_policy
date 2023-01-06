module "name" {
    source = "../"
    groups = ["developers"]
    policies = ["arn:aws:iam::123456789:policy/test_policy"]
}
