# finance-gogo

This small project collects portfolio information from CVS files (from Google
Sheet) and calculates data based on it.

## Usage

Before you start, see the `Requirements` below in this document.

Create a file config.yml with the sources you want to be parsed.

```
# config.yml
sources:
- name: Document0
  url: https://docs.google.com/spreadsheets/d/e/2PACX-1vT2B_5vb-chkU9_D2wppEmKO3BS4mn6YGOaXwTWCaD2kRp8LlmkaZadPEJX1j31W9MzSLai1dHpV6Qo/pub?gid=1223520038&single=true&output=csv
```

Next set the environment variable CONFIG_PATH to that file.


```shell
$ export CONFIG_PATH=config.yml
```

Next you can run  the following command to fetch and group the information
specified in the config file.

```
$ finance-gogo group

[
  {
    "Name": "Document0",
    "Entries": {
      "A0RPWH": {
        "Paid": 14158.289999999995,
        "Costs": 144.32,
        "Amount": 268.24499999999995
      },
      "A2H58J": {
        "Paid": 970.98,
        "Costs": 9.83,
        "Amount": 229.89699999999996
      }
    }
  }
]
```

## Requirements

You need to create a Google Sheet with the following columns:
- source: comdirect or trade republic for example
- date: format `dd.mm.yyyy`
- paid: paid amount for the whole transaction
- cost column0
- cost column1
- cost column2
- wkn number
- amount
- rate
- addition column0

| comdirect | 02.06.2019 | 75 | 1.11 | 0 | 0 | A0RPWH | 1.59 | 46.56 | 73.89 |

Then export the current tab via `File` and `Publish to the web`. Switch
`Entire Document` to your tab name. Change `Web page` to
`Comma-separated values (.csv)`.
