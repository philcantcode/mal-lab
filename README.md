# Contents

| Project | Description |
|----------|----------|
| /payloads/benign_go_exe | Basic benign payload that pops open a message box using WIN32 API.  |
| /payloads/calc_exe | DLL written in Go that pops the calculator open |
| /payloads/direct_syscall_notepad_go | Uses direct syscalls to CreateProcess opening notepad.exe | 
| /file_parsing/pe_parser | Golang pipeline that parses PE files iteratively by parsing bytes. |
| /code_loading/load_dll_c | C program that loads a DLL |
| /code_loading_load_dll_go | Go program that loads a DLL |
| /anti_analysis/metadata_edit_go | Go program that sets icon and windows metadata during build. |
| /anti_analysis/file_float | Python script that adds 200mb nulls to the end of a file to bloat it in order to bypass AV checks that might be limited to smaller files. |

# Golang Setup

- Every folder is a stand-alone Golang project.
- Every folder needs to be initialised with:

1. `go mod init github.com/philcantcode/mal-lab/<parent-folder>/<sub-folder>`
2. `go mod tidy`
3. In the root `go.work` add `./<parent-folder>/<sub-folder>` 

## Compiling

- Ensure that the `CGO_ENABLE` envirionrment variable is set to `1`
- Check go environment variables by typing `go env`
- Ensure that the `CC` env variable is set in the windows environment panel, point it to MinGW
- Don't compile with Cygwin as it will require Cygwin DLLs to be present on the target
- Compile with mingw GCC instead, it uses the Win32API calls instead

Ref: https://sourceforge.net/projects/mingw-w64/files/latest/download
Ref: Pre-compiled MinGW for win: https://winlibs.com/

## Golang Things

- To import `unsafe` without actually using it and prevent FMT from removing it, do `import _ "unsafe"`

# AV Detection

| Action | Description | AV Response | Mitigation |
|----------|----------|----------|----------|
| Load DLL | Loaded a DLL with calc.exe payload | None | N/A |
| File Bloating | Adds 200mb of nulls to an executable payload | If 4GB appended, flagged as Trojan | Try padding with legit strings instead of nulls |
| Direct Syscalls to Open Process | Opens a process using direct syscalls | None - Would have expected it to not allow direct syscalls outside of `ntdll.dll` | N/A |
