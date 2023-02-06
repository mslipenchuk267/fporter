package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define Subcommands & flags
	cpCmd := flag.NewFlagSet("cp", flag.ExitOnError)     // subcommand def
	cpIn := cpCmd.String("in", "", "Input Directory")    // flag def
	cpOut := cpCmd.String("out", "", "Output Directory") // flag def

	mvCmd := flag.NewFlagSet("mv", flag.ExitOnError)
	mvIn := mvCmd.String("in", "", "Input Directory")
	mvOut := mvCmd.String("out", "", "Output Directory")

	rpCmd := flag.NewFlagSet("rp", flag.ExitOnError)
	rpIn := rpCmd.String("in", "", "Input Directory")
	rpRgx := rpCmd.String("rgx", "", "Regex Schema")

	// Execute User's File Manipulation Procedure
	switch os.Args[1] {
	case "cp":
		fmt.Println("Invoking Copy Procedure")
		cpCmd.Parse(os.Args[2:])
		fmt.Println("\tInput Dir:", *cpIn)
		fmt.Println("\tOutput Dir:", *cpOut)
		fmt.Println("\tTrailing Args:", cpCmd.Args())
	case "mv":
		fmt.Println("Invoking Move Procedure")
		mvCmd.Parse(os.Args[2:])
		fmt.Println("\tInput Dir:", *mvIn)
		fmt.Println("\tOutput Dir:", *mvOut)
		fmt.Println("\tTrailing Args:", mvCmd.Args())
	case "rp":
		fmt.Println("Invoking Replace Procedure")
		rpCmd.Parse(os.Args[2:])
		fmt.Println("\tInput Dir:", *rpIn)
		fmt.Println("\tRegex:", *rpRgx)
		fmt.Println("\tTrailing Args:", rpCmd.Args())
	default:
		fmt.Println("Arg 1 Error: Please specify valid procedure: Copy (cp), Move (mv), Replace (rp)")
		os.Exit(1)
	}

}
