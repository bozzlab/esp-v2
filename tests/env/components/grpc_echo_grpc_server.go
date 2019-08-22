// Copyright 2019 Google Cloud Platform Proxy Authors
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

package components

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	grpcEchoPath = "../../bin/grpc_echo_server"
)

type GrpcEchoGrpcServer struct {
	*Cmd
}

func NewGrpcEchoGrpcServer(port uint16) (*GrpcEchoGrpcServer, error) {
	cmd := exec.Command(grpcEchoPath, fmt.Sprintf("127.0.0.1:%v", port))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return &GrpcEchoGrpcServer{
		Cmd: &Cmd{
			name: "grpcEchoGRPCServer",
			Cmd:  cmd,
		},
	}, nil
}
