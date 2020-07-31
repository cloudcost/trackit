//   Copyright 2020 MSolution.IO
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package indexes

import (
	"context"

	"github.com/trackit/trackit/db"
	"github.com/trackit/trackit/es/indexes/common"
	"github.com/trackit/trackit/models"
)

func updateOutdatedIndexes(ctx context.Context) error {
	for _, data := range versioningData {
		outdatedMappingsForTemplate, err := models.OutdatedEsMappings(db.Db, data.Name, data.Version)
		if err != nil {
			return err
		}

		for _, outdatedIndex := range outdatedMappingsForTemplate {
			err := updateOutdatedIndex(ctx, outdatedIndex, data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func updateOutdatedIndex(ctx context.Context, ev *models.EsVersioning, template common.VersioningData) error {
	ev.CurrentVersion = template.Version
	ev.Update(db.Db)
	return nil
}
