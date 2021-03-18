package diceyml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var basic_validate_yml = `version: 2.0
envs:
  TERMINUS_APP_NAME: PMP
  TERMINUS_TRACE_ENABLE: true
  TRACE_SAMPLE: 1

services:
  pmp-ui:
    ports:
    - 80
    expose:
    - 80
    deployments:
      replicas: 1
    depends_on:
    - pmp-backend
    resources:
      cpu: -0.1
      mem: -256
    health_check:
      exec:
        cmd: echo 1
    traffic_security:
      mode: https
    endpoints:
    - domain: dubbo-test.*
    - domain: dubbo-cc.dev.terminus.io
      path: /123
      policies:
        cors:
          allow_origins: any
        rate_limit:
          qps: 100

  pmp-backend:
    deployments:
      replicas: 1
    ports:
    - 5080
    resources:
      cpu: 0.1
      mem: -512
    health_check:
      exec:
        cmd: echo 1

addons:
#  mysql:
#    plan: mysql:small
#    options:
#      create_dbs: pmp
  pmp-redis1:
    plan: redis:small
    image: redis:alpine
  pmp-zk:
    plan: zookeeper:medium
`

func TestBasicValidate(t *testing.T) {
	d, err := New([]byte(basic_validate_yml), false)
	assert.Nil(t, err)
	es := BasicValidate(d.Obj())
	assert.Equal(t, 3, len(es), "%v", es)

}
