// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiCloud SSO admin users.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSsoFortigateCloudAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSsoFortigateCloudAdminCreate,
		Read:   resourceSystemSsoFortigateCloudAdminRead,
		Update: resourceSystemSsoFortigateCloudAdminUpdate,
		Delete: resourceSystemSsoFortigateCloudAdminDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSsoFortigateCloudAdminCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSsoFortigateCloudAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoFortigateCloudAdmin resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSsoFortigateCloudAdmin(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSsoFortigateCloudAdmin resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSsoFortigateCloudAdmin")
	}

	return resourceSystemSsoFortigateCloudAdminRead(d, m)
}

func resourceSystemSsoFortigateCloudAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSsoFortigateCloudAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoFortigateCloudAdmin resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSsoFortigateCloudAdmin(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoFortigateCloudAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSsoFortigateCloudAdmin")
	}

	return resourceSystemSsoFortigateCloudAdminRead(d, m)
}

func resourceSystemSsoFortigateCloudAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSsoFortigateCloudAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSsoFortigateCloudAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoFortigateCloudAdminRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemSsoFortigateCloudAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoFortigateCloudAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSsoFortigateCloudAdmin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoFortigateCloudAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemSsoFortigateCloudAdminName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminVdom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemSsoFortigateCloudAdminVdomName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSsoFortigateCloudAdminVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSsoFortigateCloudAdmin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemSsoFortigateCloudAdminName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("accprofile", flattenSystemSsoFortigateCloudAdminAccprofile(o["accprofile"], d, "accprofile", sv)); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vdom", flattenSystemSsoFortigateCloudAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
			if !fortiAPIPatch(o["vdom"]) {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdom"); ok {
			if err = d.Set("vdom", flattenSystemSsoFortigateCloudAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
				if !fortiAPIPatch(o["vdom"]) {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSsoFortigateCloudAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSsoFortigateCloudAdminName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandSystemSsoFortigateCloudAdminVdomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSsoFortigateCloudAdminVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSsoFortigateCloudAdmin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSsoFortigateCloudAdminName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("accprofile"); ok {
		t, err := expandSystemSsoFortigateCloudAdminAccprofile(d, v, "accprofile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemSsoFortigateCloudAdminVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
