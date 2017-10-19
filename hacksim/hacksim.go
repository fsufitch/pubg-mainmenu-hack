package main

import (
	"bufio"
	"fmt"
	"os"
)

const badHostsEntry = "104.239.228.225 front.battlegroundsgame.com"

func waitForEnter() {
	fmt.Println("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	hostsPath := getHostsPath()
	fmt.Println("Using hosts file:", hostsPath)

	// ===== Read original hosts
	f, err := os.Open(hostsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "opening hosts file:", err)
		waitForEnter()
		os.Exit(1)
	}

	lines := []string{}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading hosts file:", err)
		waitForEnter()
		return
	}
	f.Close()

	// ===== Write compromised hosts
	fmt.Println("Compromissing hosts with string:", badHostsEntry)

	f, err = os.OpenFile(hostsPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		fmt.Fprintln(os.Stderr, "opening hosts file:", err)
		waitForEnter()
		os.Exit(1)
	}

	for _, line := range lines {
		fmt.Println("=====", line)
		_, err = fmt.Fprintln(f, line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "writing to hosts file:", err)
			waitForEnter()
			os.Exit(1)
		}
	}
	fmt.Println("=====", badHostsEntry)
	_, err = fmt.Fprintln(f, badHostsEntry)
	if err != nil {
		fmt.Fprintln(os.Stderr, "writing to hosts file:", err)
		waitForEnter()
		os.Exit(1)
	}

	f.Close()

	waitForEnter()

	// ===== Restore original hosts
	fmt.Println("Restoring original hosts lines")

	f, err = os.OpenFile(hostsPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		fmt.Fprintln(os.Stderr, "opening hosts file:", err)
		waitForEnter()
		os.Exit(1)
	}

	for _, line := range lines {
		fmt.Println("=====", line)
		_, err = fmt.Fprintln(f, line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "writing to hosts file:", err)
			waitForEnter()
			os.Exit(1)
		}
	}

	f.Close()

	fmt.Println("Restored original hosts")
	waitForEnter()
}
