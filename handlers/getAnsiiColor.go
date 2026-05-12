package main

func getAnsiColor(colorName string) (string, bool) {
    m := map[string]string{
		"red": "\033[31m",
		"green": "\033[32m",
		"yellow": "\033[33m",
		"blue": "\033[34m",
		"magenta": "\033[35m",
		"cyan": "\033[36m",
		"white": "\033[37m",
	}

	val, ok := m[colorName]

	return val, ok
}

