package main

func build(cmd Commander, inputs Inputs) error {
	args := []string{
		"build",
		"--file", inputs.Dockerfile,
	}

	if inputs.Cache != "" {
		args = append(args, "--cache-from", inputs.Cache)
	}

	for _, v := range inputs.BuildArgs {
		args = append(args, "--build-arg", v)
	}

	for _, tag := range inputs.Tags {
		args = append(args, "--tag", tag)
	}

	args = append(args, inputs.Path)

	return cmd.Run("docker", args...)
}
