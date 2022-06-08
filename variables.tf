variable "resource_group" {
  description = "resource group to create the service plan in"
  type = object({
    name     = string
    location = string
  })
}

variable "service_plan" {
  description = "Object containing the details for a service plan"
  type = object({
    os_type  = string
    sku_name = string
    custom_tags = map(string)
    worker_count = number
  })
  default = {
    os_type = "Linux"
    sku_name = "P1v2"
    worker_count = 1
    custom_tags = {}
  }
}

variable "service_plan_name" {
  description = "service plan name"
  type        = string
}
