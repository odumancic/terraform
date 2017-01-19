package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/atlassian/go-vtm"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceSSLServerKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSLServerKeyCreate,
		Read:   resourceSSLServerKeyRead,
		Update: resourceSSLServerKeyUpdate,
		Delete: resourceSSLServerKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"private": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: diffCurrentAndRemoteSSLKey,
			},

			"public": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func diffCurrentAndRemoteSSLKey(k, old, new string, d *schema.ResourceData) bool {
	shaOfContent := sha256.Sum256([]byte(new))
	base64Content := base64.StdEncoding.EncodeToString(shaOfContent[:])
	log.Printf("[DEBUG] [%s] %s == %s is %s", d.Id(), base64Content, old, base64Content == old)
	return base64Content == old
}

func resourceSSLServerKeyCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceSSLServerKeySet(d, meta)
	if err != nil {
		return err
	}

	return resourceSSLServerKeyRead(d, meta)
}

func resourceSSLServerKeyRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client

	r, resp, err := c.GetSSLServerKey(d.Id())
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	if d.Get("name") == nil {
		d.Set("name", d.Id())
	}

	d.Set("note", string(*r.Basic.Note))
	d.Set("private", string(*r.Basic.Private))
	d.Set("public", string(*r.Basic.Public))
	d.Set("request", string(*r.Basic.Request))

	return nil
}

func resourceSSLServerKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceSSLServerKeySet(d, meta)
	if err != nil {
		return err
	}

	return resourceSSLServerKeyRead(d, meta)
}

func resourceSSLServerKeyDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewSSLServerKey(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}

func resourceSSLServerKeySet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewSSLServerKey(d.Get("name").(string))

	setString(&r.Basic.Note, d, "note")
	setString(&r.Basic.Private, d, "private")
	setString(&r.Basic.Public, d, "public")
	setString(&r.Basic.Request, d, "request")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

	return nil
}
