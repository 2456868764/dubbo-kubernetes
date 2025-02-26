/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package snp

import (
	"github.com/apache/dubbo-kubernetes/api/mesh"
	core_runtime "github.com/apache/dubbo-kubernetes/pkg/core/runtime"
	"github.com/apache/dubbo-kubernetes/pkg/snp/server"
	"github.com/pkg/errors"
)

func Setup(rt core_runtime.Runtime) error {
	if !rt.Config().KubeConfig.IsKubernetesConnected {
		return nil
	}
	snp := server.NewSnp(rt.Config(), rt.KubeClient().DubboClientSet())
	snp.CertStorage = rt.CertStorage()
	snp.CertClient = rt.CertClient()
	mesh.RegisterServiceNameMappingServiceServer(rt.GrpcServer().SecureServer, snp)
	mesh.RegisterServiceNameMappingServiceServer(rt.GrpcServer().PlainServer, snp)
	if err := rt.Add(snp); err != nil {
		return errors.Wrap(err, "Add Snp Component failed")
	}
	return nil
}
