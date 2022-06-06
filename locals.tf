locals {
    default_tags = {
        provisioner = "terraform"
        name        = var.service_plan_name
    }
    tags = merge(local.default_tags, var.service_plan.custom_tags)
}
