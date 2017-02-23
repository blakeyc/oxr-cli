# oxr-cli

>Command line utility to extract data from https://openexchangerates.org/ API.

**Note**: You need to register an account to get an app_id from Open Exchange Rates.

## Install

## Commands

### Params

`--app_id` Your application ID from Open Exchange Rates API.

`--job` Name of the job to run (latest|historical).

`--output` Path to the file you wish to write results to.

`--dates` Date of historical exchange rates, can supply multiple dates.

`--base` Currency to use as the base, defaults to USD

### Latest
    oxr --job latest --app_id YOUR_APP_ID --output PATH_TO_FILE

### Historical
    oxr --job historical --dates YYYY-MM-DD --app_id YOUR_APP_ID --output PATH_TO_FILE
