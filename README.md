# Placebo

Placebo is a command-line tool designed for creating and managing fake patient data. It's particularly useful for testing purposes in healthcare applications, offering functionalities to generate CSV files with fake patient data and to send HL7 messages simulating different patient scenarios.

## Features

- **CSV File Generation**: Create CSV files with fake patient data.
- **HL7 Message Sending**: Send HL7 messages with automatically generated fake patient data.

## Installation

Run the provided installer script in the root folder to have this tool installed.

*Note*: Please run the script with elevated permissions

```
$ sudo ./installer.sh
```

## Usage

### Create a CSV File

To create a CSV file with fake patient data, use the following command:

    placebo --file csv [number of patients]

- This command creates a random CSV file with a fake patient at `/tmp/`.
- Adding a number to the `csv` command will produce multiple fake patients.
- Example: `placebo --file csv 4` creates a CSV file with 4 fake patients.

### Send HL7 Messages

To send an HL7 message with automatically generated fake patient data, use the `placebo --send hl7` command. This feature supports various healthcare scenarios through different sub-commands.

    placebo --send hl7 [sub command]

- **Basic Usage**: Sends an HL7 message to the default address `127.0.0.1:9700`.
- **Sub-commands**:
  - `post_admit`: Generates an ADT^A01 event that admits a patient.
    - Usage: `placebo --send hl7 post_admit`
  - `post_discharge`: Generates ADT^A01 and ADT^A03 events to admit and then discharge a patient.
    - Usage: `placebo --send hl7 post_discharge`
  - `pre_admit`: Generates an ADT^A05 event that establishes preadmit information.
    - Usage: `placebo --send hl7 pre_admit`
  - `referral`: Generates a REF^I12 event that inform for a patient referral info.
    - Usage: `placebo --send hl7 referral`

Example:

    placebo --send hl7 post_discharge

This command sends an HL7 message that admits and then discharges a patient.

#### Helpful Auxiliary Send Commands

|Flag | Description | Usage |
| --- | --- | --- |
|`last` | Open last sent hl7 message in an interactive prompt. | `placebo --send hl7 last` |

## Error Handling

If an error occurs during file creation or HL7 message sending, the tool will output an error message.

## Contributing

(Provide instructions for how to contribute to the project, if applicable.)

## License

(Specify the license under which this tool is released.)

## Contact

(Provide contact information or a link to the project repository.)

