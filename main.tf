resource "azurerm_service_plan" "main" {
  resource_group_name = var.resource_group.name
  location            = var.resource_group.location

  name          =  var.service_plan_name
  os_type       =  var.service_plan.os_type
  sku_name      =  var.service_plan.sku_name
  worker_count  =  var.service_plan.worker_count

  tags          = local.tags
}
