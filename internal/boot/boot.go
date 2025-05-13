package boot

func Run() (err error) {
	if err = InitConfig(); err != nil {
		return
	}

	if err = NewService(); err != nil {
		return
	}

	return nil
}
