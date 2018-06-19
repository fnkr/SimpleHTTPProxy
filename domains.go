package main

func checkHost(host string) bool {
	for _, regex := range config.Hosts {
		if regex.MatchString(host) {
			return true
		}
	}
	return false
}
