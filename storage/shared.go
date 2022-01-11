//
// Copyright zerjioang. 2021 All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package storage

import (
	"context"
	"database/sql"
)

// Repository specifies the method declared for database layer interaction
type Repository interface {
	Create(ctx context.Context, obj DbItem) error
	ReadByKey(ctx context.Context, obj DbItem, key interface{}) error
	FindOne(ctx context.Context, obj DbItem, query string, params ...string) error
	Update(ctx context.Context, obj DbItem) error
	Delete(ctx context.Context, obj DbItem) error
}

type DbItem interface {
	SQLItemParser
	SetId(id string) error
	Id() ID
	SetDeleted() error
}

// SQLItemParser interface implements a custom sql.Rows to custom Struct data loader
// instead of relying on GORM reflection methods
type SQLItemParser interface {
	LoadStruct(rows *sql.Rows) (interface{}, error)
}
