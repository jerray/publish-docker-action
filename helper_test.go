package main

import (
	"reflect"
	"testing"
)

func Test_resolveInputs(t *testing.T) {
	cases := []struct {
		github     GitHub
		inputs     *Inputs
		hasError   bool
		tags       []string
		repository string
	}{
		{
			github: GitHub{
				Repository: "username/repo",
				Commit:     "45ba489c4f97b5f854ebaba6454b51fa",
				Ref:        "refs/heads/master",
			},
			inputs: &Inputs{
				Tags: []string{"my_tag"},
			},
			hasError:   false,
			tags:       []string{"username/repo:my_tag"},
			repository: "username/repo",
		},
		{
			github: GitHub{
				Repository: "username/repo",
				Commit:     "45ba489c4f97b5f854ebaba6454b51fa",
				Ref:        "refs/heads/master",
			},
			inputs: &Inputs{
				Repository: "someone/repo",
				Tags:       []string{"my_tag"},
			},
			hasError:   false,
			tags:       []string{"someone/repo:my_tag"},
			repository: "someone/repo",
		},
		{
			github: GitHub{
				Repository: "username/repo",
				Commit:     "45ba489c4f97b5f854ebaba6454b51fa",
				Ref:        "refs/pull/master",
			},
			inputs: &Inputs{
				Tags:             []string{},
				AllowPullRequest: false,
			},
			hasError:   true,
			tags:       []string{},
			repository: "username/repo",
		},
	}
	for _, c := range cases {
		err := resolveInputs(c.github, c.inputs)
		if c.repository != c.inputs.Repository {
			t.Errorf("expect repository is %s, actual is %s", c.repository, c.inputs.Repository)
		}
		if c.hasError && err == nil {
			t.Errorf("expect error but it never occurs")
		}
		if !c.hasError && err != nil {
			t.Errorf("unexpected error occurs: %v", err)
		}
		if !reflect.DeepEqual(c.tags, c.inputs.Tags) {
			t.Errorf("expect tag list is %v, actual is %s", c.tags, c.inputs.Tags)
		}
	}
}

func Test_resolveRef(t *testing.T) {
	cases := []struct {
		ref        string
		expectType string
		expectName string
	}{
		{"refs/heads/master", RefTypeBranch, "master"},
		{"refs/heads/20190927/tests", RefTypeBranch, "20190927/tests"},
		{"refs/tags/v1.0.0", RefTypeTag, "v1.0.0"},
		{"refs/tags/v2.0.0-rc1", RefTypeTag, "v2.0.0-rc1"},
		{"refs/pull/master", RefTypePull, "master"},
		{"refs/pull/2019/09/27/pull", RefTypePull, "2019/09/27/pull"},
		{"refs/unknown/master", "", "master"},
		{"", "", ""},
	}
	for _, c := range cases {
		typ, name := resolveRef(GitHub{Ref: c.ref})
		if typ != c.expectType {
			t.Errorf("expect ref type is %s, actual is %s", c.expectType, typ)
		}
		if name != c.expectName {
			t.Errorf("expect ref name is %s, actual is %s", c.expectName, name)
		}
	}
}

func Test_resolveAutoTag(t *testing.T) {
	cases := []struct {
		typ     string
		name    string
		inputs  *Inputs
		expects []string
	}{
		{RefTypeBranch, "master", &Inputs{AutoTag: false, Tags: []string{}}, []string{}},
		{RefTypeBranch, "master", &Inputs{AutoTag: false, Tags: []string{"master"}}, []string{"master"}},
		{RefTypeBranch, "master", &Inputs{AutoTag: false, Tags: []string{"a", "b", "c"}}, []string{"a", "b", "c"}},
		{RefTypeBranch, "master", &Inputs{AutoTag: true, Tags: []string{"master"}}, []string{"latest"}},
		{RefTypeBranch, "master", &Inputs{AutoTag: true, Tags: []string{"feature", "develop"}}, []string{"latest"}},
		{RefTypeBranch, "master", &Inputs{AutoTag: true, Tags: []string{}}, []string{"latest"}},
		{RefTypeBranch, "develop", &Inputs{AutoTag: true, Tags: []string{}}, []string{"develop"}},
		{RefTypeTag, "v1.0.0", &Inputs{AutoTag: true, Tags: []string{}}, []string{"1", "1.0", "1.0.0"}},
		{RefTypePull, "master", &Inputs{AutoTag: true, Tags: []string{}}, []string{"pr-master"}},
	}
	for _, c := range cases {
		resolveAutoTag(c.typ, c.name, c.inputs)
		if !reflect.DeepEqual(c.expects, c.inputs.Tags) {
			t.Errorf("expect tag list is %v, actual is %v", c.expects, c.inputs.Tags)
		}
	}
}

func Test_resolveSemanticVersionTag(t *testing.T) {
	cases := []struct {
		input  string
		expect []string
	}{
		// invalid semantic version
		{"1", []string{"1"}},
		{"master", []string{"master"}},
		{"20190921-actions", []string{"20190921-actions"}},

		// valid semantic version
		{"0.0.0", []string{"0", "0.0", "0.0.0"}},
		{"1.0.0", []string{"1", "1.0", "1.0.0"}},
		{"2.7.10", []string{"2", "2.7", "2.7.10"}},
		{"v1.0.0", []string{"1", "1.0", "1.0.0"}},
		{"v2.3.6", []string{"2", "2.3", "2.3.6"}},
		{"3.0.0-rc1", []string{"3-rc1", "3.0-rc1", "3.0.0-rc1"}},
		{"v3.0.0-alpha.1", []string{"3-alpha.1", "3.0-alpha.1", "3.0.0-alpha.1"}},
		{"4.0.0-alpha.1.2", []string{"4-alpha.1.2", "4.0-alpha.1.2", "4.0.0-alpha.1.2"}},
	}
	for _, c := range cases {
		actual := resolveSemanticVersionTag(c.input)
		if !reflect.DeepEqual(c.expect, actual) {
			t.Errorf("expect tag list is %v, actual is %s", c.expect, actual)
		}
	}
}
