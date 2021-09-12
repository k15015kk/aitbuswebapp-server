package server

func Init() error {
	router, err := Router()

	if err != nil {
		return err
	}

	router.Run(":8080")

	return nil
}
