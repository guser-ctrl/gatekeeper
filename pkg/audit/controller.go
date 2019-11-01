/*
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

package audit

import (
	"context"

	opa "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddToManager adds audit manager to the Manager
func AddToManager(m manager.Manager, opa *opa.Client, reporter StatsReporter) error {
	am, err := New(context.Background(), m.GetConfig(), opa, reporter)
	if err != nil {
		return err
	}
	return m.Add(am)
}
