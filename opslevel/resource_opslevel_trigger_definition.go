package opslevel

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/opslevel/opslevel-go/v2022"
)

func resourceTriggerDefinition() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a webhook action",
		Create:      wrap(resourceTriggerDefinitionCreate),
		Read:        wrap(resourceTriggerDefinitionRead),
		Update:      wrap(resourceTriggerDefinitionUpdate),
		Delete:      wrap(resourceTriggerDefinitionDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Trigger Definition",
				ForceNew:    false,
				Required:    true,
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "The owner of the Trigger Definition",
				ForceNew:    false,
				Required:    true,
			},
			"action_id": {
				Type:        schema.TypeString,
				Description: "The action that will be triggered by the Trigger Definition",
				ForceNew:    false,
				Optional:    true,
			},
		},
	}
}

func resourceTriggerDefinitionCreate(d *schema.ResourceData, client *opslevel.Client) error {
	input := opslevel.CustomActionsTriggerDefinitionCreateInput{
		Name:   d.Get("name").(string),
		Owner:  *opslevel.NewID(d.Get("owner").(string)),
		Action: opslevel.NewID(d.Get("action_id").(string)),
	}

	resource, err := client.CreateTriggerDefinition(input)
	if err != nil {
		return err
	}
	d.SetId(resource.Id.(string))

	return resourceTriggerDefinitionRead(d, client)
}

func resourceTriggerDefinitionRead(d *schema.ResourceData, client *opslevel.Client) error {
	id := d.Id()

	resource, err := client.GetTriggerDefinition(*opslevel.NewIdentifier(id))
	if err != nil {
		return err
	}

	if err := d.Set("name", resource.Name); err != nil {
		return err
	}

	if err := d.Set("owner", resource.Owner.Id.(string)); err != nil {
		return err
	}

	if err := d.Set("action_id", resource.Action.Id.(string)); err != nil {
		return err
	}

	return nil
}

func resourceTriggerDefinitionUpdate(d *schema.ResourceData, client *opslevel.Client) error {
	input := opslevel.CustomActionsTriggerDefinitionUpdateInput{
		Id: d.Id(),
	}

	if d.HasChange("name") {
		input.Name = opslevel.NewString(d.Get("name").(string))
	}
	if d.HasChange("owner") {
		input.Owner = opslevel.NewID(d.Get("owner").(string))
	}
	if d.HasChange("action_id") {
		input.Action = opslevel.NewID(d.Get("action_id").(string))
	}

	_, err := client.UpdateTriggerDefinition(input)
	if err != nil {
		return err
	}

	return resourceTriggerDefinitionRead(d, client)
}

func resourceTriggerDefinitionDelete(d *schema.ResourceData, client *opslevel.Client) error {
	id := d.Id()
	err := client.DeleteTriggerDefinition(*opslevel.NewIdentifier(id))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
