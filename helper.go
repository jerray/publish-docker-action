package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	RefTypeBranch = "branch"
	RefTypeTag    = "tag"
	RefTypePull   = "pull"
)

func resolveInputs(github GitHub, inputs *Inputs) error {
	if inputs.Repository == "" {
		inputs.Repository = strings.ToLower(github.Repository)
	}

	if inputs.Registry != "" && !strings.HasPrefix(inputs.Repository, inputs.Registry) {
		inputs.Repository = strings.Join([]string{inputs.Registry, inputs.Repository}, "/")
	}

	typ, name := resolveRef(github)

	if typ == RefTypePull && !inputs.AllowPullRequest {
		return fmt.Errorf("if you want to build a pull request, please set `with.allow_pull_request` to `true`")
	}

	resolveAutoTag(typ, name, inputs)

	for i, t := range inputs.Tags {
		inputs.Tags[i] = strings.Join([]string{inputs.Repository, t}, ":")
	}

	return nil
}

func resolveRef(github GitHub) (string, string) {
	var typ, name string
	refs := strings.SplitN(github.Ref, "/", 3)
	if len(refs) == 3 {
		switch refs[1] {
		case "heads":
			typ = RefTypeBranch
		case "tags":
			typ = RefTypeTag
		case "pull":
			typ = RefTypePull
		}
		name = refs[2]
	}
	return typ, name
}

func resolveAutoTag(typ, name string, inputs *Inputs) {
	if !inputs.AutoTag {
		return
	}

	tags := make([]string, 0)

	name = strings.Replace(name, "/", "-", -1)
	switch typ {
	case RefTypeBranch:
		if name == "master" {
			name = "latest"
		}
		tags = append(tags, name)
	case RefTypeTag:
		tags = append(tags, resolveSemanticVersionTag(name)...)
	case RefTypePull:
		tags = append(tags, strings.Join([]string{"pr", name}, "-"))
	}

	inputs.Tags = tags
}

func resolveSemanticVersionTag(name string) []string {
	re := regexp.MustCompile(`^v?(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	matches := re.FindAllStringSubmatch(name, -1)
	if matches == nil {
		return []string{name}
	}
	subs := matches[0]
	suffixes := make([]string, 0)

	prerelease := subs[4]
	if prerelease != "" {
		suffixes = append(suffixes, "-", prerelease)
	}

	// metadata := subs[5]
	// if metadata != "" {
	// 	suffixes = append(suffixes, ".", metadata)
	// }

	tags := make([]string, 0)
	suffix := strings.Join(suffixes, "")
	for n := 2; n <= 4; n++ {
		vs := make([]string, n-1)
		copy(vs, subs[1:n])
		tags = append(tags, strings.Join([]string{strings.Join(vs, "."), suffix}, ""))
	}
	return tags
}
