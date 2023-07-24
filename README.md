# httphasher

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Introduction

httphasher is a command-line interface (CLI) program designed to fetch multiple URLs simultaneously using a specified number of parallel processes.
By providing the number of parallel processes and a list of URLs, the program efficiently retrieves the HTTP responses from the URLs and calculates the MD5 hash of each response. The output is then presented to the user, displaying the URL alongside its corresponding MD5 hash.

## Features

- Fetches multiple URLs in parallel: The program efficiently fetches the content of multiple URLs concurrently, utilizing the specified number of parallel processes, to minimize response time.
- Calculates MD5 hash: After receiving the HTTP response from each URL, the program calculates the MD5 hash of the response data. MD5 hashing provides a unique and fixed-size representation of the response content.

## Installation

1. Make sure you have Go installed. You can download it from the official website: https://golang.org/
2. Clone this repository
3. Navigate to the root directory of the project using CLI.
4. Build the executable for the CLI program using the following command: go build cmd/main.go
   This will create an executable binary named "main" in the same directory.
5. Now, you are ready to use the "httphasher" CLI program. Simply run it with the desired number of parallel processes and a list of URLs as arguments, like in the example below:
- ./main -parallel=3 example1.com example2.net example3.org
6. For units testing. simply run this command:
- go test ./... 
