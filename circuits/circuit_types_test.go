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

func TestClient_ListCircuitTypes(t *testing.T) {
	client, err := circuits.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("ListCircuitTypes", vcrConf)
	client.SetHTTPClient(vcr.Client)
	opts := models.ListCircuitTypesRequest{}
	res, err := client.ListCircuitTypes(opts)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, res)
	assert.GreaterOrEqual(t, res.Count, 0)
}

func TestClient_GetCircuitType(t *testing.T) {
	client, err := circuits.NewClient(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HTTPClient()
	vcr := govcr.NewVCR("GetCircuitType", vcrConf)
	client.SetHTTPClient(vcr.Client)

	// First list circuit types to get a valid ID
	opts := models.ListCircuitTypesRequest{}
	opts.Limit = 1
	listRes, err := client.ListCircuitTypes(opts)
	if err != nil {
		t.Fatal(err)
	}

	if listRes.Count > 0 && len(listRes.Results) > 0 {
		typeID := listRes.Results[0].ID
		res, err := client.GetCircuitType(typeID)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotNil(t, res)
		assert.Equal(t, typeID, res.ID)
		t.Log(res.Name)
	} else {
		t.Skip("No circuit types found in Netbox instance")
	}
}
