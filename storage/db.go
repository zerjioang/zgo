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
	"errors"
	"log"
	"strings"
	"time"

	"github.com/zerjioang/zgo/cache"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

// Generator returns a new struct of given object as interface{}
type Generator func() interface{}

// ORMDatabase is a GORM powered database access layer struct
type ORMDatabase struct {
	// Db is an orm based database connection access
	Db       *gorm.DB
	Metadata map[string]DbItem
	// Cache
	Cache *cache.Cache
}

// compilation time interface implementation check
var _ Repository = (*ORMDatabase)(nil)

var (
	QueryCachedErr = errors.New("query cached")
)

func (s *ORMDatabase) Create(ctx context.Context, obj DbItem) error {
	// make sure the new object to be created has a valid id
	_ = obj.Id()
	tx := s.Db.WithContext(ctx).Create(obj)
	return CheckResult(tx, false)
}

func (s *ORMDatabase) CreateNoWarnDuplicate(ctx context.Context, obj DbItem) error {
	// make sure the new object to be created has a valid id
	_ = obj.Id()
	tx := s.Db.WithContext(ctx).Create(obj)
	return CheckResult(tx, true)
}

// CreateIfNot attempts to register the given object in the database if not exists
func (s *ORMDatabase) CreateIfNot(ctx context.Context, obj DbItem) error {
	// make sure the new object to be created has a valid id
	tx := s.Db.WithContext(ctx).Where(obj).First(&obj)
	if err := CheckResult(tx, false); err != nil {
		// handler error
		if err.Error() == "record not found" {
			// write the item to the database and return any error occured
			return s.Create(ctx, obj)
		}
	}
	// since we did not get any errors, means that the object we are trying to include in the
	// database already exists
	return nil
}

func (s *ORMDatabase) ReadByKey(cacheKey string, ctx context.Context, gen Generator, out interface{}) (interface{}, error) {
	return s.withCache(cacheKey, gen, 10*time.Minute, func(dst interface{}) error {
		tx := s.Db.WithContext(ctx).First(dst, "id", out)
		return CheckResult(tx, false)
	})
}

// ReadOne returns object row in database as unique item
func (s *ORMDatabase) ReadOne(cacheKey string, ctx context.Context, gen Generator) (interface{}, error) {
	return s.withCache(cacheKey, gen, 10*time.Minute, func(dst interface{}) error {
		tx := s.Db.WithContext(ctx).Where(dst).First(&dst)
		return CheckResult(tx, false)
	})
}

// ReadAll makes a SELECT * style operation with given model and reads all fields
func (s *ORMDatabase) ReadAll(cacheKey string, ctx context.Context, tx *gorm.DB, gen Generator) (interface{}, error) {
	return s.withCache(cacheKey, gen, 10*time.Minute, func(dst interface{}) error {
		if tx != nil {
			// reuse passed tx Db context
			tx = tx.Find(dst) // find product with integer primary key
			return CheckResult(tx, false)
		}
		tx = s.Db.WithContext(ctx).Find(dst) // find product with integer primary key
		return CheckResult(tx, false)
	})
}

// ReadAllWithFields makes a SELECT query and ONLY retrieves selected column names
func (s *ORMDatabase) ReadAllWithFields(key string, ctx context.Context, tx *gorm.DB, genObj func() interface{}, columns ...string) (interface{}, error) {
	return s.withCache(key, genObj, 10*time.Minute, func(dst interface{}) error {
		if tx != nil {
			// reuse passed tx Db context
			tx = tx.Select(columns).Find(dst)
		} else {
			tx = s.Db.WithContext(ctx).Select(columns).Find(dst)
		}
		return CheckResult(tx, false)
	})
}

// ReadAllWithFields makes a SELECT query and ONLY retrieves selected column names
func (s *ORMDatabase) withCache(key string, gen Generator, d time.Duration, f func(dst interface{}) error) (interface{}, error) {
	// 1 first check if requested data is in the cache
	// note that, key value must be unique and must always be paired with method parameters
	data, found := s.Cache.Get(key)
	if found {
		// cache HIT
		return data, nil
	}
	// cache MISS
	// we need to generate destination obj to unmarshal data by GORM
	obj := gen()
	if err := f(obj); err != nil {
		return nil, err
	}
	// if no error in database query, add result to cache
	s.Cache.Set(key, obj, d)
	return obj, nil
}

func (s *ORMDatabase) GetItems(cacheKey string, ctx context.Context, order string, filter DbItem, limit uint, dst Generator) (interface{}, error) {
	return s.withCache(cacheKey, dst, 10*time.Minute, func(dst interface{}) error {
		tx := s.Db.WithContext(ctx)
		if order != "" {
			tx = tx.Order(order)
		}
		if limit > 0 {
			tx = tx.Limit(int(limit))
		}
		tx = tx.Find(dst) // find product with integer primary key
		return CheckResult(tx, false)
	})
}

func (s *ORMDatabase) FindOne(cacheKey string, ctx context.Context, gen Generator, query string, params ...string) (interface{}, error) {
	return s.withCache(cacheKey, gen, 10*time.Minute, func(dst interface{}) error {
		tx := s.Db.WithContext(ctx).First(dst, query, params)
		return CheckResult(tx, false)
	})
}

func (s *ORMDatabase) FindByKeyWithFields(ctx context.Context, obj DbItem, columns ...string) error {
	tx := s.Db.WithContext(ctx).Select(columns).First(obj)
	return CheckResult(tx, false)
}

func (s *ORMDatabase) Update(ctx context.Context, obj DbItem) error {
	tx := s.Db.WithContext(ctx).Model(obj).Updates(obj)
	return CheckResult(tx, false)
}

func (s *ORMDatabase) SoftDelete(ctx context.Context, obj DbItem) error {
	_ = obj.SetDeleted()
	tx := s.Db.WithContext(ctx).Model(obj).Updates(obj)
	return CheckResult(tx, false)
}

func (s *ORMDatabase) Delete(ctx context.Context, obj DbItem) error {
	tx := s.Db.WithContext(ctx).Delete(obj)
	return CheckResult(tx, false)
}

// Exists returns if the item exists in the database or not
func (s *ORMDatabase) Exists(ctx context.Context, obj DbItem) (bool, error) {
	tx := s.Db.WithContext(ctx).Where(obj).First(&obj)
	if tx.Error != nil {
		return tx.Error.Error() == "record not found", tx.Error
	}
	return true, nil
}

// FindMatches returns a list of matching elements
func (s *ORMDatabase) FindMatches(ctx context.Context, obj DbItem, dst interface{}) error {
	tx := s.Db.WithContext(ctx).Find(dst, obj)
	return CheckResult(tx, false)
}

func (s *ORMDatabase) Set(key string, obj DbItem) {
	s.Metadata[key] = obj
}

func (s *ORMDatabase) Get(key string) DbItem {
	v, _ := s.Metadata[key]
	return v
}

func (s *ORMDatabase) DeleteTable(obj DbItem) error {
	// get the table name given the passed struct
	stmt := &gorm.Statement{DB: s.Db}
	if err := stmt.Parse(obj); err != nil {
		return err
	}
	// execute the query
	// we assume GORM provides the right table name an no injections are possible
	tx := s.Db.Raw("DELETE FROM " + stmt.Schema.Table)
	return CheckResult(tx, false)
}

// Close closes database connection
func (s *ORMDatabase) Close() error {
	if s != nil && s.Db != nil {
		sqlDB, err := s.Db.DB()
		if err != nil {
			return err
		}
		// Close
		return sqlDB.Close()
	}
	return nil
}

func CheckResult(tx *gorm.DB, noWarnDuplicate bool) error {
	if tx == nil {
		return errors.New("could not get a valid response from database")
	}
	// ID           // returns inserted data's primary key
	// Error        // returns error
	// RowsAffected // returns inserted records count
	if tx.Error != nil && tx.Error == QueryCachedErr {
		return nil
	}
	if tx.Error != nil {
		log.Println(tx.Error)
		switch tx.Error.(type) {
		case *pgconn.PgError:
			err := tx.Error.(*pgconn.PgError)
			if strings.Index(err.Message, "duplicate key value violates") != -1 {
				// duplicate key error detected
				if noWarnDuplicate {
					return nil
				}
				// return duplicate error
				return errors.New(err.Severity + ": requested operation is not allowed to be executed.")
			}
			return tx.Error
		default:
			return tx.Error
		}
	}
	return nil
}
