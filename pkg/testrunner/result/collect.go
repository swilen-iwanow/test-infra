// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package result

import (
	"fmt"
	"path/filepath"

	"github.com/gardener/gardener/pkg/client/kubernetes"
	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	"github.com/gardener/test-infra/pkg/testrunner"
	"github.com/gardener/test-infra/pkg/util"
	log "github.com/sirupsen/logrus"
)

// Collect collects results of all testruns and writes them to a file.
// It returns whether there are failed testruns or not.
func Collect(config *Config, tmClient kubernetes.Interface, namespace string, runs testrunner.RunList) (bool, error) {
	testrunsFailed := false
	for _, run := range runs {
		cfg := *config
		cfg.OutputDir = filepath.Join(config.OutputDir, util.RandomString(3))
		err := Output(&cfg, tmClient, namespace, run.Testrun, run.Metadata)
		if err != nil {
			return false, err
		}

		if cfg.OutputDir != "" && cfg.ESConfigName != "" {
			err = IngestDir(cfg.OutputDir, cfg.ESConfigName)
			if err != nil {
				log.Errorf("cannot persist file %s: %s", cfg.OutputDir, err.Error())
			} else {
				err := MarkTestrunsAsIngested(tmClient, run.Testrun)
				if err != nil {
					log.Warn(err.Error())
				}
			}
		}

		if run.Testrun.Status.Phase == tmv1beta1.PhaseStatusSuccess {
			log.Infof("Testrun %s finished successfully", run.Testrun.Name)
		} else {
			testrunsFailed = true
			log.Errorf("Testrun %s failed with phase %s", run.Testrun.Name, run.Testrun.Status.Phase)
		}
		fmt.Print(util.PrettyPrintStruct(run.Testrun.Status))
		printStatusTable(run.Testrun.Status.Steps)
	}

	return testrunsFailed, nil
}
