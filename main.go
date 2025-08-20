package main

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {
	cfg := new(config)
	startRepl(cfg)
}
