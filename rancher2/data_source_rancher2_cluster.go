package rancher2

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceRancher2Cluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRancher2ClusterRead,
		// Schema: clusterFields(),
		Schema: map[string]*schema.Schema{
			"clusters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Required Admin username for AKS",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Required Agent dns prefix for AKS",
						},
					},
				},
			},
		},
	}
}

func dataSourceRancher2ClusterRead(d *schema.ResourceData, meta interface{}) error {

	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}

	cluster, err := client.Cluster.List(nil)

	out := make([]interface{}, len(cluster.Data))

	for i, cluster := range cluster.Data {
		clusterData := map[string]interface{}{
			"id":   cluster.ID,
			"name": cluster.Name,
		}
		out[i] = clusterData
	}

	d.SetId("test-data-rancher-cluster")
	d.Set("clusters", out)
	if err != nil {
		return err
	}

	return nil
}
