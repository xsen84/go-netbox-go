// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package circuits_test

import (
	"os"
	"testing"

	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"

	"github.com/xsen84/go-netbox-go/circuits"
	"github.com/xsen84/go-netbox-go/models"
)

func TestClient_ListCircuitProviders(t *testing.T) {
	client, err := circuits.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListCircuitProviders", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListCircuitProvidersRequest{}
	res, err := client.ListCircuitProviders(opts)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, res)
	assert.GreaterOrEqual(t, res.Count, 0)
}

func TestClient_GetCircuitProvider(t *testing.T) {
	client, err := circuits.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("GetCircuitProvider", vcrConf)
	client.SetHTTPClient(vcr.Client)

	// First list providers to get a valid ID
	opts := models.ListCircuitProvidersRequest{}
	opts.Limit = 1
	listRes, err := client.ListCircuitProviders(opts)
	if err != nil {
		t.Fatal(err)
	}

	if listRes.Count > 0 && len(listRes.Results) > 0 {
		providerID := listRes.Results[0].ID
		res, err := client.GetCircuitProvider(providerID)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotNil(t, res)
		assert.Equal(t, providerID, res.ID)
		t.Log(res.Name)
	} else {
		t.Skip("No circuit providers found in Netbox instance")
	}
}
