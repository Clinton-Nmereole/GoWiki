# GoWiki
A program that allows the user to get information from Wikipedia without leaving the terminal. The user simply needs to enter the word they wish to search after the executable. 
GoWiki is a command-line tool written in Go that allows you to retrieve information about a subject from Wikipedia using its webscraping. Simply provide a search string as a parameter, and GoWiki will fetch relevant information and display it in your terminal.

## Features

- Retrieve information about a subject from Wikipedia.
- Display summary, page URL, and other relevant details.
- Lightweight and easy-to-use command-line interface.

## Installation

Before installing GoWiki, make sure you have Go (Golang) installed on your system.

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/gowiki.git
   ```
   
2. Navigate to the project directory:

   ```bash
   cd gowiki
   ```
   
3. go build

   ```bash
   go build
   ```

4. Move the binary to a directory in your system's PATH (optional):

   ```bash
   sudo mv gowiki /usr/local/bin
   ```


## Usage

Run GoWiki from the command line by providing a search string as an argument:

```bash
gowiki "Albert Einstein"
```
Replace "Albert Einstein" with the subject you want to search for on Wikipedia.

## Example

```bash
gowiki "OpenAI"
```
## License

This project is licensed under the MIT License - see the LICENSE file for details.

Disclaimer: This project is developed and maintained by github.com/Clinton-Nmereol. It is not affiliated with or endorsed by Wikipedia.
