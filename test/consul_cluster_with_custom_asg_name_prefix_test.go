package test

import (
	"testing"
)

func TestConsulClusterWithCustomASGNamePrefixUbuntu16Ami(t *testing.T) {
	t.Parallel()
	terraformVars := map[string]interface{}{
		"asg_name_prefix": "consul-asg",
	}
	runConsulClusterTestWithVars(t, "ubuntu16-ami", ".", "../examples/consul-ami/consul.json", "ubuntu", terraformVars, "")
}
