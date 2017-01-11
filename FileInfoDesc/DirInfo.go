/* 

Language : Go
The program is written using Go language for specifically Ethos OS.
Ethos Operating System Description: http://www.ethos-os.org/
Program : Display files present in a folder and also display their sizes

*/

package main

//Import specific Ethos packages in order to its functions
import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
	"ethos/efmt"	
)


func main () {

	path := "/programs/"

	status := ethosLog.RedirectToLog("DirInfo")
	
	// Check if the above call was executed in correct manner.
	if status != syscall.StatusOk {
		efmt.Fprintf(syscall.Stderr, "Error opening %v: %v\n", path, status)
		syscall.Exit(syscall.StatusOk)
	}

	// Create a new variable of type DirR to store certain results
	var result = DirR { count: 0 , size: 0 }
	fname := ".."

	// Get a file descriptor reference to the folder present in 'path' variable
	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		efmt.Fprintf(syscall.Stderr, "Error Path Opening %v: %v\n", path, status)
		syscall.Exit(syscall.StatusOk)
	}

	efmt.Printf("\nFiles at path : server/rootfs%s \n",path)
	efmt.Printf("\n%-18s %-18s \n\n","File Name","Size(bytes)")

	for {

		//Retrieve the next file name present in folder at location 'path'
		name, stat := ethos.GetNextName(fd,fname)

		if stat != syscall.StatusOk {
			break
			efmt.Fprintf(syscall.Stderr, "Error Getting Name %v: %v\n", name, stat)
			syscall.Exit(syscall.StatusOk)
		}

		fname = string(name)
		result.count = result.count + 1
		filepath := path + fname

		// Retrieve file|folder information made available by the system
		finfo, sta := ethos.GetFileInformationPath(filepath)

		efmt.Printf("%-18s %-18v \n", string(name), finfo.Size)

		if sta != syscall.StatusOk {
			efmt.Fprintf(syscall.Stderr, "Error getting file info for %v: %v", finfo, sta)
			syscall.Exit(syscall.StatusOk)
		}

		result.size = result.size + finfo.Size
		
	}

	// Display Results
	efmt.Printf("%15s \n","Result")
	efmt.Printf("Total files in directory : %v\n", result.count)
	efmt.Printf("Size of directory : %v bytes\n", result.size)

	efmt.Print("Project Done !!!\n")

}
