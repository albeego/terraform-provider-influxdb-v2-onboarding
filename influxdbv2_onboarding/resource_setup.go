package influxdbv2_onboarding

import (
	"encoding/json"
	"net/http"
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/influxdata/influxdb-client-go/v2"
)

type Setup struct {
	Allowed bool `json:"allowed"`
}

func ResourceSetup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSetupCreate,
		Read:   resourceSetupRead,
		Delete: resourceSetupDelete,
		Update: resourceSetupUpdate,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_USERNAME", "administrator"),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFLUXDB_V2_PASSWORD", "Administrator1."),
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"org": {
				Type:     schema.TypeString,
				Required: true,
			},
			"retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bucket_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"server_url": {
				Type: 	schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSetupCreate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	serverUrl := influx.ServerURL()
	d.Set("server_url", serverUrl)
	err := resourceSetupRead(d, meta)
	if err != nil {
		return fmt.Errorf("error getting status of influxdbv2 instance: %v", err)
	}
	if d.Get("allowed").(bool) {
		result, err := influx.Setup(context.Background(), d.Get("username").(string), d.Get("password").(string), d.Get("org").(string), d.Get("bucket").(string), d.Get("retention_period").(int))
		if err != nil {
			return fmt.Errorf("error setup endpoint: %v", err)
		}

		fmt.Println(result)

		d.Set("token", *result.Auth.Token)
		d.Set("user_id", result.User.Id)
		d.Set("bucket_id", result.Bucket.Id)
		d.Set("org_id", result.Org.Id)
		d.Set("auth_id", result.Auth.Id)
		id := ""
		id = influx.ServerURL()
		d.SetId(id)
	}

	return nil
}

func resourceSetupRead(d *schema.ResourceData, meta interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, d.Get("server_url").(string)+"/api/v2/setup", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	setup := &Setup{}
	if err := json.NewDecoder(resp.Body).Decode(setup); err != nil {
		return err
	}

	d.Set("allowed", setup.Allowed)
	return nil
}

func resourceSetupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSetupUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
