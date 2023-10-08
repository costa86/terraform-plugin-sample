package hashicups

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
)

func flattenOrderItemsData(orderItems *[]hc.OrderItem) []interface{} {
	if orderItems != nil {
		ois := make([]interface{}, len(*orderItems))
		for i, orderItem := range *orderItems {
			oi := make(map[string]interface{})
			oi["coffee_id"] = orderItem.Coffee.ID
			oi["coffee_name"] = orderItem.Coffee.Name
			oi["coffee_teaser"] = orderItem.Coffee.Teaser
			oi["coffee_description"] = orderItem.Coffee.Description
			oi["coffee_price"] = orderItem.Coffee.Price
			oi["coffee_name"] = orderItem.Coffee.Name
			oi["quantity"] = orderItem.Quantity
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func dataSourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*hc.Client)
	var diags diag.Diagnostics
	orderID := strconv.Itoa(d.Get("id").(int))
	order, err := c.GetOrder(orderID)

	if err != nil {
		return diag.FromErr(err)
	}
	orderItems := flattenOrderItemsData(&order.Items)

	if err := d.Set("items", orderItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(orderID)
	return diags
}

func dataSourceOrder() *schema.Resource {
	coffeeSchema := map[string]*schema.Schema{
		"coffee_id":          {Type: schema.TypeInt, Computed: true},
		"coffee_name":        {Type: schema.TypeString, Computed: true},
		"coffee_teaser":      {Type: schema.TypeString, Computed: true},
		"coffee_description": {Type: schema.TypeString, Computed: true},
		"coffee_price":       {Type: schema.TypeFloat, Computed: true},
		"coffee_image":       {Type: schema.TypeString, Computed: true},
		"quantity":           {Type: schema.TypeInt, Computed: true},
	}
	orderSchema := map[string]*schema.Schema{
		"id": {Type: schema.TypeInt,
			Required: true},
		"items": {Type: schema.TypeList, Computed: true, Elem: &schema.Resource{Schema: coffeeSchema}}}

	return &schema.Resource{
		ReadContext: dataSourceOrderRead,
		Schema:      orderSchema,
	}
}
