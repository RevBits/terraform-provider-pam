package pam

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePamSourceRead(d *schema.ResourceData, meta interface{}) error {
        config := meta.(*Config)
        secretValue, err := config.GetPamSecretValue()
        if err != nil{
                return err
        }
        d.SetId("pam_secret_value")
        d.Set("secret_value", secretValue)
        return nil
}

func dataSourcePam() *schema.Resource {
        return &schema.Resource{
                Read: resourcePamSourceRead,
                Schema: map[string]*schema.Schema{
                        "secret_value": &schema.Schema{
                                Type:     schema.TypeString,
                                Computed: true,
                        },
                },
        }
}


func providerConfig(d *schema.ResourceData) (interface{}, error) {
        c := &Config{
               Domain: d.Get("domain").(string),
               SecretID: d.Get("secret_id").(string),
               ApiKey: d.Get("api_key").(string),
        }
        return c, nil
}

func Provider() *schema.Provider {
        return &schema.Provider{
                DataSourcesMap: map[string]*schema.Resource{
                        "pam_provider": dataSourcePam(),
                },
                Schema: map[string]*schema.Schema{
                        "domain": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "secret_id": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "api_key": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
                ConfigureFunc: providerConfig,
        }
}
