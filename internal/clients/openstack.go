/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"

	"github.com/pkg/errors"
	apitypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/provider-openstack/apis/v1alpha1"
)

// GetAuthInfo returns the necessary authentication information that is necessary
// to use when the controller connects to the OpenStack API in order to reconcile
// the managed resource.
func GetAuthInfo(ctx context.Context, c client.Client, mg resource.Managed) (projectID string, g gophercloud.ProviderClient, err error) {
	switch {
	case mg.GetProviderConfigReference() != nil:
		return UseProviderConfig(ctx, c, mg)
	default:
		return "", gophercloud.ProviderClient{}, errors.New("neither providerConfigRef nor providerRef is given")
	}
}

// UseProviderConfig to return an OpenStack API client.
func UseProviderConfig(ctx context.Context, c client.Client, mg resource.Managed) (g gophercloud.ProviderClient, err error) {
	pc := &v1alpha1.ProviderConfig{}
	t := resource.NewProviderConfigUsageTracker(c, &v1alpha1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return "", g, err
	}
	if err := c.Get(ctx, apitypes.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return "", g, err
	}
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return "", g, errors.Wrap(err, "cannot get credentials")
	}

	g, err = openstack.AuthOptionsFromEnv()
	if err != nil {
		return "", g, errors.Wrap(err, "failed to initialize auth options")
	}
	return g, nil
}
