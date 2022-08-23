// Copyright 2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	helmbase "github.com/pulumi/pulumi-go-helmbase"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/middleware"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pp "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

const (
	ProviderName  = "kubernetes-coredns"
	ComponentName = ProviderName + ":index:CoreDNS"
)

func Provider() p.Provider {
	return schema.Wrap(&middleware.Scaffold{
		ConstructFn: func(pctx p.Context, typ string, name string,
			ctx *pulumi.Context, inputs pp.ConstructInputs, opts pulumi.ResourceOption) (pulumi.ComponentResource, error) {
			state := &State{}
			_, err := helmbase.Construct(ctx, state, typ, name, &Args{}, inputs, opts)
			if err != nil {
				return nil, err
			}
			return state, nil
		},
	}).WithDisplayName("Kubernetes CoreDNS").
		WithKeywords([]string{
			"pulumi",
			"kubernetes",
			"coredns",
			"kind/component",
			"category/infrastructure",
		}).
		WithPublisher("Pulumi").
		WithLanguageMap(map[string]any{
			"csharp": map[string]map[string]string{
				"namespaces": {
					"kubernetes-coredns": "KubernetesCoreDNS",
				},
				"packageReferences": {
					"Pulumi":            "3.*",
					"Pulumi.Kubernetes": "3.*",
				},
			},
			"go": map[string]any{
				"generateResourceContainerTypes": true,
				"importBasePath":                 "github.com/pulumi/pulumi-kubernetes-coredns/sdk/go/kubernetes-coredns",
			},
			"nodejs": map[string]map[string]string{
				"dependencies": {
					"@pulumi/kubernetes": "^3.7.1",
				},
				"devDependencies": {
					"typescript": "^3.7.0",
				},
			},
			"python": map[string]map[string]string{
				"requires": {
					"pulumi":            ">=3.0.0,<4.0.0",
					"pulumi-kubernetes": ">=3.7.1,<4.0.0",
				},
			},
		})
}

// Serve launches the gRPC server for the resource provider.
func Serve(version string, schema []byte) {
	if err := provider.ComponentMain(ProviderName, version, schema, Construct); err != nil {
		cmdutil.ExitError(err.Error())
	}
}

// Construct is the RPC call that initiates the creation of a new component resource. It
// creates, registers, and returns the resulting object.
func Construct(ctx *pulumi.Context, typ, name string, inputs pp.ConstructInputs,
	opts pulumi.ResourceOption) (*pp.ConstructResult, error) {
	return helmbase.Construct(ctx, &State{}, typ, name, &Args{}, inputs, opts)
}
