variable "resource_group_name" {
  type        = string
  description = "Name of the Resource Group to deploy to"
}

variable "location" {
  type        = string
  description = "Location to deploy resources"
}
variable "image_gallery_name" {
  type        = string
  description = "Name of Shared Image Gallery"
}

variable "image_gallery_description" {
  type        = string
  description = "Description of Shared Image Gallery"
}

variable "images" {
  type = map(object({
    name        = string
    os_type     = string
    description = string
    publisher   = string
    offer       = string
    sku         = string
  }))

  validation {
    condition = alltrue(
      [
        for image in var.images : contains(["Linux", "Windows"], image.os_type)
      ]
    )
    error_message = "OS type must be Windows or Linux."
  }
  description = "Images to deploy to Shared Image Gallery"
}

variable "tags" {
  type        = map(string)
  description = "Tags to apply to resources"
}
