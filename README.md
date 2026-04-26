# ASCII Art Web

## Description
ASCII Art Web is a web-based application that allows users to generate ASCII art from standard text input. The application features a clean, responsive, and modern user interface where you can dynamically select from different ASCII fonts (such as Standard, Shadow, or Thinkertoy) to stylize your text and display it immediately in your browser.

## Authors
- Maleek

## Usage: how to run
To run the web server locally, ensure you have [Go](https://golang.org/) installed on your machine.
Navigate to the root directory of the project in your terminal and execute the following command:

```bash
go run ./cmd/server/
```

The application will start, typically listening on port `8080`. 
Open your web browser and navigate to: `http://localhost:8080` (or whatever specific port your environment logs) to begin using the service.

## Implementation details: algorithm
The core ASCII generator (`internal/ascii/`) operates through the following steps to construct the art:

1. **Input Processing**: The user's input string and selected font name are received from the form submission. The text is split into distinct lines or segments if the user input contains line breaks (`\r\n` or `\n`).
2. **Font File Ingestion**: The algorithm reads the corresponding font text file from the `fonts/` directory. 
3. **Newline Normalization**: To prevent cross-platform issues and handle irregular files, all `\r\n` characters in the raw font data are converted to standard `\n`, allowing for a predictable split.
4. **Slicing**: The full font string is split by newlines, resulting in a large array (`fontSlice`) where each index precisely corresponds to a line of the formatted font data.
5. **Character Mapping & Horizontal Construction**:
   - Every printable ASCII character is strictly 8 lines tall.
   - The algorithm evaluates the starting index of each character within `fontSlice` using its standard decimal ASCII value.
   - It loops exactly 8 times (creating the item row by row). In each iteration, it strings together the corresponding horizontal slice for every character in the user's word.
   - A `strings.Builder` is used to accumulate the characters efficiently and concatenate them perfectly across horizontal rows.
6. **Serving**: The finalized `strings.Builder` output is returned to the router and served to the frontend as the payload in the AJAX response.
