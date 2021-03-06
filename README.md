# Azure Shared Image Gallery module

This repo contains a [Terraform](https://www.terraform.io/) module that defines a Shared Image Gallery.
=======
# Azure Shared Image Gallery module

This repo contains a [Terraform](https://www.terraform.io/) module that defines a Shared Image Gallery.
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.2.3 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | ~> 3.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_shared_image.images](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/shared_image) | resource |
| [azurerm_shared_image_gallery.images](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/shared_image_gallery) | resource |
| [azurerm_user_assigned_identity.images](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/user_assigned_identity) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_image_gallery_description"></a> [image\_gallery\_description](#input\_image\_gallery\_description) | Description of Shared Image Gallery | `string` | n/a | yes |
| <a name="input_image_gallery_name"></a> [image\_gallery\_name](#input\_image\_gallery\_name) | Name of Shared Image Gallery | `string` | n/a | yes |
| <a name="input_images"></a> [images](#input\_images) | Images to deploy to Shared Image Gallery | <pre>map(object({<br>    name        = string<br>    os_type     = string<br>    description = string<br>    publisher   = string<br>    offer       = string<br>    sku         = string<br>  }))</pre> | n/a | yes |
| <a name="input_location"></a> [location](#input\_location) | Location to deploy resources | `string` | n/a | yes |
| <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name) | Name of the Resource Group to deploy to | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | Tags to apply to resources | `map(string)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_images"></a> [images](#output\_images) | The images deployed to the Shared Image Gallery |
| <a name="output_shared_image_gallery_id"></a> [shared\_image\_gallery\_id](#output\_shared\_image\_gallery\_id) | Resource ID of the Shared Image Gallery |
| <a name="output_shared_image_gallery_unique_name"></a> [shared\_image\_gallery\_unique\_name](#output\_shared\_image\_gallery\_unique\_name) | The unique name of the Shared Image Gallery |
| <a name="output_user_assigned_identity_client_id"></a> [user\_assigned\_identity\_client\_id](#output\_user\_assigned\_identity\_client\_id) | Client ID of the User Assigned Identity |
| <a name="output_user_assigned_identity_id"></a> [user\_assigned\_identity\_id](#output\_user\_assigned\_identity\_id) | The resource ID of the User Assigned Identity |
| <a name="output_user_assigned_identity_principal_id"></a> [user\_assigned\_identity\_principal\_id](#output\_user\_assigned\_identity\_principal\_id) | Object ID of the User Assigned Identity |
| <a name="output_user_assigned_identity_tenant_id"></a> [user\_assigned\_identity\_tenant\_id](#output\_user\_assigned\_identity\_tenant\_id) | Tenant ID of the User Assigned Identity |
<!-- END_TF_DOCS -->
