// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nsmgr_proxy

import (
	"context"
	"net"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"google.golang.org/grpc"

	"github.com/networkservicemesh/sdk/pkg/networkservice/chains/client"
	"github.com/networkservicemesh/sdk/pkg/networkservice/chains/endpoint"
	"github.com/networkservicemesh/sdk/pkg/networkservice/common/authorize"
	"github.com/networkservicemesh/sdk/pkg/networkservice/common/connect"
	"github.com/networkservicemesh/sdk/pkg/networkservice/common/interdomainurl"
	"github.com/networkservicemesh/sdk/pkg/networkservice/common/swap"
	"github.com/networkservicemesh/sdk/pkg/networkservice/core/adapters"
	"github.com/networkservicemesh/sdk/pkg/tools/addressof"
	"github.com/networkservicemesh/sdk/pkg/tools/grpcutils"
	"github.com/networkservicemesh/sdk/pkg/tools/token"
)

type NSMgrProxy interface {
	networkservice.NetworkServiceServer
	Register(s *grpc.Server)
}

type nsmgrProxy struct {
	endpoint.Endpoint
}

func (n *nsmgrProxy) Register(s *grpc.Server) {
	grpcutils.RegisterHealthServices(s, n)
	networkservice.RegisterNetworkServiceServer(s, n)
}

func NewServer(ctx context.Context, name string, externalIP net.IP, generatorFunc token.GeneratorFunc, options ...grpc.DialOption) NSMgrProxy {
	result := new(nsmgrProxy)

	result.Endpoint = endpoint.NewServer(ctx,
		name,
		authorize.NewServer(),
		generatorFunc,
		interdomainurl.NewServer(),
		swap.NewServer(externalIP),
		connect.NewServer(
			ctx,
			client.NewClientFactory(name,
				addressof.NetworkServiceClient(
					adapters.NewServerToClient(result)),
				generatorFunc),
			options...,
		),
	)
	return &nsmgrProxy{Endpoint: result}
}
