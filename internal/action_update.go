package internal

func Update() error {
	for _, repo := range GetRepositories() {
		err := repo.Update()
		if err != nil {
			return err
		}
	}

	return nil
}
