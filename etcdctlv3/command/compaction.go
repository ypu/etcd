// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"strconv"

	"github.com/coreos/etcd/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/Godeps/_workspace/src/google.golang.org/grpc"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

// NewCompactionCommand returns the CLI command for "compaction".
func NewCompactionCommand() cli.Command {
	return cli.Command{
		Name: "compaction",
		Action: func(c *cli.Context) {
			compactionCommandFunc(c)
		},
	}
}

// compactionCommandFunc executes the "compaction" command.
func compactionCommandFunc(c *cli.Context) {
	if len(c.Args()) != 1 {
		panic("bad arg")
	}

	rev, err := strconv.ParseInt(c.Args()[0], 10, 64)
	if err != nil {
		panic("bad arg")
	}

	conn, err := grpc.Dial(c.GlobalString("endpoint"))
	if err != nil {
		panic(err)
	}
	kv := pb.NewKVClient(conn)
	req := &pb.CompactionRequest{Revision: rev}

	kv.Compact(context.Background(), req)
}
