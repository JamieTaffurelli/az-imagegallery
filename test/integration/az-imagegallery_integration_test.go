package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestIT_OutputsAreCorrect(t *testing.T) {
	t.Parallel()

	tfOptions := &terraform.Options{
		TerraformDir: "../../",
		Vars: map[string]interface{}{
			"resource_group_name":       "testrg",
			"location":                  "francecentral",
			"image_gallery_name":        "testimggal",
			"image_gallery_description": "Image Gallery Description",
			"images": map[string]interface{}{
				"image": map[string]interface{}{
					"name":        "windows",
					"os_type":     "Windows",
					"description": "Windows description",
					"publisher":   "Windows",
					"offer":       "2016-datacenter",
					"sku":         "WindowsServer",
				},
			},
			"tags": map[string]string{},
		},
		NoColor: true,
	}

	defer terraform.Destroy(t, tfOptions)

	terraform.InitAndApply(t, tfOptions)

	assert.Equal(t, "/subscriptions/"+os.Getenv("AZURE_SUBSCRIPTION_ID")+"/resourceGroups/testrg/Microsoft.Compute/sharedImageGalleries/testimggal", terraform.Output(t, tfOptions, "shared_image_gallery_id"), "Shared Image Gallery Resource ID should be correct")
	assert.Equal(t, os.Getenv("AZURE_SUBSCRIPTION_ID")+"-TESTIMGGAL", terraform.Output(t, tfOptions, "shared_image_gallery_unique_name"), "Shared Image Gallery unique name should be correct")
	assert.Equal(t, "/subscriptions/"+os.Getenv("AZURE_SUBSCRIPTION_ID")+"/resourceGroups/testrg/Microsoft.ManagedIdentity/userAssignedIdentities/testimggal", terraform.Output(t, tfOptions, "user_assigned_identity_id"), "User Assigned Identity Resource ID should be correct")
}
