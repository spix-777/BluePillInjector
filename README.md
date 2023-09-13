# Blue Pill Injector

## Overview

The "Blue Pill" is a Go program designed to overwrite files in the current directory with a hexadecimal-encoded message and subsequently delete them. It begins its operation by requesting sudo privileges to ensure unrestricted access to the files.

**WARNING**: This program is extremely dangerous and can result in total data loss in the directory where it is executed. It is vital to handle this tool with utmost caution and to be absolutely certain of your actions when using it. Always ensure that there is no critical data in the current directory before running the program.

## Dependencies

- Go (1.x or newer)
- github.com/cheggaaa/pb/v3

Ensure you have the necessary Go package installed using the following command before running or building the program:

## sh
go get github.com/cheggaaa/pb/v3

## Usage

To run the program, use the following command:

sudo go run main.go

Upon execution, the program will prompt the user for confirmation, warning them of the impending file rewrites and deletions. If the user consents, the program will then proceed to:

Find all files in the current directory.
Overwrite each file's contents with a repeated hex message: "You will only find the blue pills ...".
Delete each file.
A progress bar shows the progress of the operation, which ends with a "blue pill" message and graphic indicating completion.

## Development

The codebase is structured into several functions, each with specific roles:

## Banner
This function is responsible for displaying the ASCII banner at the initiation of the script. The color of the banner is blue, achieved through ANSI escape codes.

## Main
This function outlines the main flow of the program, which includes:

Ensuring sudo privileges.
Collecting user consent through a prompt.
Initializing and updating a progress bar as files are processed.
Calling the rewriteFileWithHexMessage function to handle file modification and deletion.
rewriteFileWithHexMessage
This function is in charge of the overwriting and truncation process of each file. It performs the following steps:

Opening the target file.
Reading all bytes from it.
Creating a hex-encoded message from the predefined string.
Overwriting all the bytes in the file with a repeated hex message.
Truncating the file to the new length.
Contribution

If you wish to contribute, please do so with extreme caution to avoid introducing functionalities that could potentially increase the risk of data loss. Always test thoroughly in safe environments before suggesting changes.

## License

The project currently does not hold a license, implying all rights remain with the original author. Consider adding a license to clarify the permissions around the project.

## Disclaimer

This tool is provided without any warranty and is to be used at the user's own risk. The authors cannot be held responsible for any harm to data or systems that arises from the use of this tool. Always ensure to have backups of vital data.