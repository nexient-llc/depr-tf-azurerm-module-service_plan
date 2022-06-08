resource_group = {
    name = "deb-test-devops"
    location = "eastus"
}

service_plan = {
    os_type = "Linux"
    sku_name = "P1v2"
    custom_tags = {}
    worker_count = 2
}

service_plan_name = "demo-eastus-dev-000-lsp-000"