output "service_plan_id" {
  description = "service plan ID"
  value       = azurerm_service_plan.main.id
}

output "service_plan_kind" {
  description = "Kind of the service plan"
  value = azurerm_service_plan.main.kind
}

output "location" {
  description = "Location"
  value = azurerm_service_plan.main.location
}

output "service_plan_name" {
  description = "Name of the service plan"
  value = azurerm_service_plan.main.name
}

output "os_type" {
  description = "OS type of the service plan"
  value = azurerm_service_plan.main.os_type
}

output "is_reserved" {
  description = "Boolean Reserved flag"
  value = azurerm_service_plan.main.reserved
}

output "worker_count" {
  description = "The number of workers in the plan"
  value = azurerm_service_plan.main.worker_count
}

output "rg_name" {
  description = "Resource group of the Service plan"
  value = azurerm_service_plan.main.resource_group_name
}

output "sku_name" {
  description = "The SKU for the sevice plan"
  value = azurerm_service_plan.main.sku_name
}
