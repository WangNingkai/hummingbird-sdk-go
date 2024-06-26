/*******************************************************************************
 * Copyright 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package server

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	h "github.com/wangningkai/hummingbird-sdk-go"

	"github.com/winc-link/edge-driver-proto/drivercommon"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommonRPCServer struct {
	drivercommon.UnimplementedCommonServer
}

func NewCommonRPCServer() *CommonRPCServer {
	return &CommonRPCServer{}
}

// Ping tests whether the service is working
func (crs *CommonRPCServer) Ping(context.Context, *emptypb.Empty) (*drivercommon.Pong, error) {
	return &drivercommon.Pong{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// Version obtains version information from the target service.
func (crs *CommonRPCServer) Version(context.Context, *emptypb.Empty) (*drivercommon.VersionResponse, error) {
	return &drivercommon.VersionResponse{
		Version: h.SdkVersion,
	}, nil
}

func decoder(payload string, v interface{}) error {
	d := json.NewDecoder(strings.NewReader(payload))
	d.UseNumber()
	return d.Decode(v)
}
