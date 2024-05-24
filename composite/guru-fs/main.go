package main

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}
	file1.getName()

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")

	// folder3 := &Folder{
	// 	name: "Folder3",
	// }
	// folder3.add(file1)
	// folder3.add(file2)
	// folder3.add(file3)
	// folder3.add(folder1)
	// folder3.add(folder2)
	// folder3.search("File")
	// for _, f := range folder3.components {
	// 	fmt.Printf("%v", f)
	// }
}
