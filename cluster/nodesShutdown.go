// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"fmt"
	"github.com/LLoyd66/elastigo/api"
	"net/url"
	"strconv"
	"strings"
)

// NodesShutdown allows the caller to shutdown between one and all nodes in the cluster
// delay is a integer representing number of seconds
// passing "" or "_all" for the nodes parameter will shut down all nodes
// see http://www.elasticsearch.org/guide/reference/api/admin-cluster-nodes-shutdown/
func NodesShutdown(delay int, nodes ...string) error {
	shutdownUrl := fmt.Sprintf("/_cluster/nodes/%s/_shutdown", strings.Join(nodes, ","))
	if delay > 0 {
		var values url.Values = url.Values{}
		values.Add("delay", strconv.Itoa(delay))
		shutdownUrl += "?" + values.Encode()
	}
	_, err := api.DoCommand("POST", shutdownUrl, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
