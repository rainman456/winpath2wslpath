# winpath2wslpath
Convert directory paths on windows to use on  WSL linux like ubuntu or kali automatically 

To build go build -o wslpath wslpath.go

USAGE wslpath "C:\Users\USER\Desktop\blablasomepath"

Example Output:

Windows Path: C:\Users\USER\Desktop\output_directory

WSL Path:     /mnt/c/Users/USER/Desktop/output_directory
