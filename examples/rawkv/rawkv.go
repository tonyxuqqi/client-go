// Copyright 2021 TiKV Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"

	"github.com/tikv/client-go/v2/config"
	"github.com/tikv/client-go/v2/rawkv"
)

func main() {
	var tenant_id uint32
	tenant_id = 1
	ctx := context.WithValue(context.TODO(), "tenant-id", tenant_id)
	cli, err := rawkv.NewClient(ctx, []string{"172.16.7.237:23711"}, config.DefaultConfig().Security)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	fmt.Printf("cluster ID: %d\n", cli.ClusterID())

	key := []byte("usertable:user1489991169483799204")
	val := []byte("PingCAP")

	for i := 0; i < 10000; i++ {
		// put key into tikv
		err = cli.Put(ctx, key, val)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Successfully put %s:%s to tikv\n", key, val)

	// get key from tikv
	val, err = cli.Get(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("found val: %s for key: %s\n", val, key)

	// delete key from tikv
	err = cli.Delete(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %s deleted\n", key)

	// get key again from tikv
	val, err = cli.Get(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("found val: %s for key: %s\n", val, key)
}
