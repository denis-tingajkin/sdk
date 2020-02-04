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

package dnscontext

import (
	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// ConfigFromConnection returns default getter for DNS configs. Builds config from the connection.
func ConfigFromConnection() GetDNSConfigsFunc {
	return func(conn *networkservice.Connection) []*networkservice.DNSConfig {
		domain := conn.GetNetworkService()
		ip := conn.GetContext().GetIpContext().GetDstIpAddr()
		return []*networkservice.DNSConfig{
			{
				DnsServerIps:  []string{ip},
				SearchDomains: []string{domain},
			},
		}
	}
}
