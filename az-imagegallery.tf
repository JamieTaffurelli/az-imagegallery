resource "azurerm_shared_image_gallery" "images" {
  name                = var.image_gallery_name
  resource_group_name = var.resource_group_name
  location            = var.location
  description         = var.image_gallery_description
  tags                = var.tags
}

resource "azurerm_shared_image" "images" {
  for_each            = var.images
  name                = each.value.name
  gallery_name        = azurerm_shared_image_gallery.images.name
  resource_group_name = var.resource_group_name
  location            = var.location
  os_type             = each.value.os_type
  description         = each.value.description
  specialized         = false

  identifier {
    publisher = each.value.publisher
    offer     = each.value.offer
    sku       = each.value.sku
  }
  tags = var.tags
}

resource "azurerm_user_assigned_identity" "images" {
  name                = var.image_gallery_name
  resource_group_name = var.resource_group_name
  location            = var.location
  tags                = var.tags
}