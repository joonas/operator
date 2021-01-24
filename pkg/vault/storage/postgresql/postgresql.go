/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package postgresql

import (
	"context"
	"fmt"
	"strings"

	api "kubevault.dev/operator/apis/kubevault/v1alpha1"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var postgresqlStorageFmt = `
storage "postgresql" {
%s
}
`

type Options struct {
	api.PostgreSQLSpec

	// the connection string to use to authenticate and connect to PostgreSQL.
	ConnectionURL string
}

func NewOptions(kubeClient kubernetes.Interface, ns string, s api.PostgreSQLSpec) (*Options, error) {
	var (
		url string
	)

	if s.ConnectionURLSecret != "" {
		sr, err := kubeClient.CoreV1().Secrets(ns).Get(context.TODO(), s.ConnectionURLSecret, metav1.GetOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get connection url secret(%s)", s.ConnectionURLSecret)
		}

		if value, ok := sr.Data["connection_url"]; ok {
			url = string(value)
		} else {
			return nil, errors.Errorf("connection_url not found in secret(%s)", s.ConnectionURLSecret)
		}

	}
	return &Options{
		s,
		url,
	}, nil
}

func (o *Options) Apply(pt *core.PodTemplateSpec) error {
	return nil
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/postgresql.html
//
// GetGcsConfig creates postgresql storae config from GcsSpec
func (o *Options) GetStorageConfig() (string, error) {
	params := []string{}
	if o.ConnectionURLSecret != "" {
		params = append(params, fmt.Sprintf(`connection_url = "%s"`, o.ConnectionURL))
	}
	if o.Table != "" {
		params = append(params, fmt.Sprintf(`table = "%s"`, o.Table))
	}
	if o.MaxParallel != 0 {
		params = append(params, fmt.Sprintf(`max_parallel = "%d"`, o.MaxParallel))
	}

	storageCfg := fmt.Sprintf(postgresqlStorageFmt, strings.Join(params, "\n"))
	return storageCfg, nil
}
