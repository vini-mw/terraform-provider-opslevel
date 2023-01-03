package opslevel

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/opslevel/opslevel-go/v2022"
)

func resourceWebhookAction() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a webhook action",
		Create:      wrap(resourceWebhookActionCreate),
		Read:        wrap(resourceWebhookActionRead),
		Update:      wrap(resourceWebhookActionUpdate),
		Delete:      wrap(resourceWebhookActionDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the external action.",
				ForceNew:    false,
				Required:    true,
			},
			"liquid_template": {
				Type:        schema.TypeString,
				Description: "Template that can be used to generate a Webhook payload.",
				ForceNew:    false,
				Required:    true,
			},
			"webhook_url": {
				Type:        schema.TypeString,
				Description: "The URL of the webhook action.",
				ForceNew:    false,
				Required:    true,
			},
			"http_method": {
				Type:         schema.TypeString,
				Description:  "The http method used to call the webhook action.",
				ForceNew:     false,
				Required:     true,
				ValidateFunc: validation.StringInSlice(opslevel.AllCustomActionsHttpMethodEnum(), false),
			},
			"headers": {
				Type:        schema.TypeMap,
				Description: "HTTP headers to be passed along with your Webhook when triggered.",
				ForceNew:    false,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceWebhookActionCreate(d *schema.ResourceData, client *opslevel.Client) error {

	input := opslevel.CustomActionsWebhookActionCreateInput{
		Name:           d.Get("name").(string),
		LiquidTemplate: d.Get("liquid_template").(string),
		WebhookURL:     d.Get("webhook_url").(string),
		HTTPMethod:     opslevel.CustomActionsHttpMethodEnum(d.Get("http_method").(string)),
		Headers:        convertHeadersMap(d.Get("headers")),
	}

	resource, _ := client.CreateWebhookAction(input)
	//if err != nil {
	//	return err
	//}
	d.SetId(resource.Id.(string))

	return nil
}

func resourceWebhookActionRead(d *schema.ResourceData, client *opslevel.Client) error {
	//id := d.Id()

	//resource, err := client.GetWebhookAction(*opslevel.NewIdentifier(id))
	//if err != nil {
	//	return err
	//}
	//
	//if err := d.Set("name", resource.Name); err != nil {
	//	return err
	//}
	//
	//if err := d.Set("liquid_template", resource.LiquidTemplate); err != nil {
	//	return err
	//}
	//
	//if err := d.Set("webhook_url", resource.WebhookURL); err != nil {
	//	return err
	//}
	//
	//if err := d.Set("http_method", string(resource.HTTPMethod)); err != nil {
	//	return err
	//}

	return nil
}

func resourceWebhookActionUpdate(d *schema.ResourceData, client *opslevel.Client) error {
	input := opslevel.CustomActionsWebhookActionUpdateInput{
		Id: d.Id(),
	}

	if d.HasChange("name") {
		input.Name = opslevel.NewString(d.Get("name").(string))
	}
	if d.HasChange("liquid_template") {
		input.WebhookURL = opslevel.NewString(d.Get("liquid_template").(string))
	}
	if d.HasChange("webhook_url") {
		input.WebhookURL = opslevel.NewString(d.Get("webhook_url").(string))
	}
	if d.HasChange("http_method") {
		input.HTTPMethod = opslevel.CustomActionsHttpMethodEnum(d.Get("http_method").(string))
	}
	if d.HasChange("headers") {
		input.Headers = convertHeadersMap(d.Get("headers"))
	}

	_, err := client.UpdateWebhookAction(input)
	if err != nil {
		return err
	}

	return nil
}

func resourceWebhookActionDelete(d *schema.ResourceData, client *opslevel.Client) error {
	id := d.Id()
	err := client.DeleteWebhookAction(*opslevel.NewIdentifier(id))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
