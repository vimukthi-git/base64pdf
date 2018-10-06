# base64pdf
Utilities for converting base64 to pdf

## Usage

- install
    - With go `go get github.com/vimukthi-git/base64pdf` 
    - with brew `brew install base64pdf`

- Help just run `base64pdf`

- if you have already a json file `testdata/input.json` and a base64 string in it at path `data.extra_data`
    running,
    ```
    base64pdf -f testdata/input.json -p data.extra_data
    ```
    will create a pdf `output.pdf` with base64 string converted to pdf