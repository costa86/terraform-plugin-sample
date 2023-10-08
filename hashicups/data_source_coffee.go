package hashicups

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCoffeesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: time.Second * 10}
	var diags diag.Diagnostics
	req, err := http.NewRequest(http.MethodGet, "http://localhost:19090/coffees", nil)

	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}

	defer r.Body.Close()

	coffees := make([]map[string]interface{}, 0)

	err = json.NewDecoder(r.Body).Decode(&coffees)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("coffees", coffees); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func dataSourceCoffees() *schema.Resource {
	ingredientSchema := map[string]*schema.Schema{"ingredient_id": {
		Type:     schema.TypeInt,
		Computed: true,
	}}

	coffeeSchema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"teaser": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"price": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"image": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ingredients": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: ingredientSchema,
			},
		},
	}

	coffeeListSchema := map[string]*schema.Schema{
		"coffees": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: coffeeSchema,
			},
		},
	}

	return &schema.Resource{
		ReadContext: dataSourceCoffeesRead,
		Schema:      coffeeListSchema,
	}
}
