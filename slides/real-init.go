func main() {
	log.Println("excalibur: initializing system")
	// our shell is called `gsh`
	program := "/gsh"

	c := exec.Command(program)
	c.Stdout, c.Stderr, c.Stdin = os.Stdout, os.Stderr, os.Stdin

	err := c.Start()
	if err != nil {
		log.Println("excalibur: could not boot system!")
	}

	for {
		log.Println("excalibur: just stayin' alive")
		time.Sleep(1 * time.Minute)
	}
}
