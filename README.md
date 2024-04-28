# Cron Parser

The Cron Parser is a command-line tool written in Go that parses a cron string and expands each field to show the times at which it will run.

## Features

- Parses a cron string in the standard format with five time fields (minute, hour, day of month, month, and day of week) plus a command.
- Supports special characters such as `*`, `,`, `-`, and `/`.
- Allows specifying range using numerical values (0-6).
- Does not require any external dependencies.

## Installation

To install this Cron Parser, you need to have Go installed on your system. Then, you can clone the repository and build the binary:

```
git clone https://github.com/shivanivarshney/deliveroo.git
cd deliveroo
go build cron_parser.go 

```
This will create an executable binary named `cron_parser` in the current directory.


## Usage

To use the Cron Parser, run the following command: 

`./cron_parser "*/15 0 1,15 * 1-5 /usr/bin/find"`

## Output

This will output the expanded cron fields in a table format:

```
minute 0 15 30 45
hour 0
day of month 1 15
month 1 2 3 4 5 6 7 8 9 10 11 12
day of week 1 2 3 4 5
command /usr/bin/find

```






