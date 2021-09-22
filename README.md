# gobasics
My personal repository to learn mind-blowing go concepts like pointers, interfaces, and channels.

## Important Concepts
On building a local package with "remotely"-callable methods:
1. create the new files in one directory. One package can span multiple files!
2. set the package name == directory name
3. In each of these package files, method names initiated with a capital letter will be a public method. Otherwise, the method will be considered as private
4. Make sure you have go mod init'ed previously, else, do it
5. Done. Call the method by this format "(package name or alias).(MethodName)"
