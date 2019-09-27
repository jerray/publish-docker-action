package main

func push(cmd Commander, inputs Inputs) error {
	for _, tag := range inputs.Tags {
		err := cmd.Run("docker", "push", tag)
		if err != nil {
			return err
		}
	}
	return nil
}
