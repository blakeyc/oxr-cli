# oxr-cli

[![Build Status](https://travis-ci.org/blakeyc/oxr-cli.svg?branch=master)](https://travis-ci.org/blakeyc/oxr-cli)

>Command line utility to extract data from https://openexchangerates.org/ API.

**Note**: You need to register an account to get an app_id from Open Exchange Rates.

## Install

Download binaries directly from the releases page, or alternatively clone and build from
source, see below for details.

TODO: Publish to homebrew etc...

## Usage

### Commands

    oxr [flags]
    oxr --help
    oxr --version

### Flags

`--app_id` Your application ID from Open Exchange Rates API.

`--job` Name of the job to run (latest|historical).

`--output` Path to the file you wish to write results to.

`--dates` Date of historical exchange rates, can supply multiple dates.

`--base` Currency to use as the base, defaults to USD

`--fields` Define the list of fields to output in file (base,currency,rate,timestamp,date)


### Example's

#### Latest
    oxr --job latest --app_id YOUR_APP_ID --output PATH_TO_FILE

#### Latest, Selecting Fields
    oxr --job latest --app_id YOUR_APP_ID --output PATH_TO_FILE --fields base,currency,rate

#### Historical
    oxr --job historical --dates 2017-01-01 --app_id YOUR_APP_ID --output PATH_TO_FILE

#### Historical Range of Dates
    oxr --job historical --dates 2017-01-01,2017-01-02 --app_id YOUR_APP_ID --output PATH_TO_FILE

## Build

To run or build from source clone the repo and run `make build` this will also install the dependencies listed below.
The executables will be output to the `./dist` directory for darwin, linux and windows.

### Dependencies

    go get github.com/mvillalba/go-openexchangerates/oxr
