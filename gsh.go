func main() {
	log.Println("welcome to the best shell ever")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("sh: ")
		text, _ := reader.ReadString('\n')
		cmd := strings.Split(strings.TrimRight(text, "\n"), " ")
		c := exec.Command(cmd[0], cmd[1:]...)
		c.Stdout, c.Stderr, c.Stdin = os.Stdout, os.Stderr, os.Stdin
		err := c.Start()
		if err != nil {
			fmt.Println("error:", err)
		}
	}
}
