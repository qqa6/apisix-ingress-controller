---
title: Upgrade Guide
---

<!--
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
-->

## Validate Compatibility

Apache APISIX Ingress project is a continuously actively developed project.
In order to make it better, some broken changes will be added when the new version is released.
For users, how to upgrade safely becomes very important.

The policy directory of this project contains these compatibility check strategies,
you can use the [`conftest`](https://github.com/open-policy-agent/conftest) tool to check the compatibility.

Here's a quick example.

```yaml
apiVersion: apisix.apache.org/v2beta2
kind: ApisixRoute
metadata:
 name: httpbin-route
spec:
 http:
 - name: rule1
   match:
     hosts:
     - httpbin.org
     paths:
       - /ip
     exprs:
     - subject:
         scope: Header
         name: X-Foo
       op: Equal
       value: bar
   backend:
     serviceName: httpbin
     servicePort: 80
```

It uses the `spec.http.backend` field that has been removed.
Save as httpbin-route.yaml.

Use conftest for compatibility check.

```bash
$ conftest test httpbin-route.yaml
FAIL - httpbin-route.yaml - main - ApisixRoute/httpbin-route: rule1 field http.backend has been removed, use http.backends instead.

2 tests, 1 passed, 0 warnings, 1 failure, 0 exceptions
```

Incompatible parts will generate errors.
