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
	"database/sql"
	"github.com/zerjioang/zgo/timer"
	"github.com/zerjioang/zgo/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

type ID string

func (i ID) String() string {
	return string(i)
}

type Item struct {
	ID           ID `gorm:"primaryKey" json:"id,omitempty"`
	ItemMetadata `gorm:"embedded"`
}

type ItemMetadata struct {
	// UNIX Epoch timestamp millis
	UpdatedAt int64 `gorm:"autoUpdateTime"  json:"updated_at,omitempty"` // Use unix milli seconds as updating time
	// UNIX Epoch timestamp millis
	CreatedAt int64           `gorm:"autoCreateTime" json:"created_at,omitempty"` // Use unix seconds as creating time
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

var _ callbacks.BeforeCreateInterface = (*ItemMetadata)(nil)

func (mt *ItemMetadata) LoadStruct(rows *sql.Rows) (interface{}, error) {
	panic("LoadStruct method needs to be implemented")
}

func (mt *ItemMetadata) BeforeCreate(*gorm.DB) (err error) {
	mt.CreatedAt = timer.Now()
	return
}

func (mt *ItemMetadata) BeforeUpdate(*gorm.DB) (err error) {
	mt.UpdatedAt = timer.Now()
	return
}

func (mt *ItemMetadata) SetDeleted() error {
	return mt.DeletedAt.Scan(timer.Now())
}

// DbItem interface implementation methids

// SetId sets current object id to given id string
// The id string must contain a valid UINT value
// otherwise an error is returned
func (s *Item) SetId(id string) error {
	s.ID = ID(id)
	return nil
}

func (s *Item) Id() ID {
	if s.ID == "" {
		// create new id
		s.ID = ID(uuid.New())
	}
	return s.ID
}
