# Placebo

[![CI](https://github.com/hl7x/placebo/actions/workflows/placebo.yml/badge.svg)](https://github.com/hl7x/placebo/actions/workflows/placebo.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

Placebo is a command-line tool designed for creating and managing fake patient data. It's particularly useful for testing purposes in healthcare applications, offering functionalities to generate CSV files with fake patient data and to send HL7 messages simulating different patient scenarios.

In addition to all of that, it has some robust features that help aid with reading HL7 messages! Useful if you're not use to reading pipes and carets.

## Features

- **CSV File Generation**: Create CSV files with automatically generated fake patient data.
- **HL7 Message Sending**: Send HL7 messages with automatically generated fake patient data.
- **HL7 Message Reading**: Feed `placebo` an hl7 file and get readible structure of the hl7 message.

## Installation

### With Go

```
$ go install github.com/hl7x/placebo/cmd/placebo@latest
```

### From source

Run the provided installer script in the root folder to have this tool installed.

*Note*: Please run the script with elevated permissions

```
$ sudo ./install.sh
```

## Usage

### Generated Header Fields

|Header | Description |
| --- | --- |
*FirstName* | Patient first name |
*LastName* | Patient last name | 
*MRN* | Patient Identifier |
*PatientId* | Patient account identifier |
*VisitId* | Encounter/visit identifier |
*Phone* | Patient phone number |
*DOB* | Patient date of birth |
*Street* | Patient street address |
*StructureNumber* | Patient address number |
*State* | Patient address state |
*City* | Patient address city |
*PostalCode* | Patient zip code |
*ArrivalDate* | Patient encounter arrival date |
*DischargeDate* | Patient encounter discharge date |
*Appointment* | Patient future appointment date |


### Create a File

To create a CSV file with fake patient data, use the following command:

    placebo --file csv [number of patients]

- This command creates a random CSV file with a fake patient at `/tmp/`.
- Adding a number to the `csv` command will produce multiple fake patients.
- Example: `placebo --file csv 4` creates a CSV file with 4 fake patients.

To create a HL7 message file with fake patient data:

    placebo --file hl7

- This command creates a random HL7 file with a fake patient at `/tmp/`.

### Send HL7 Messages

<p align="center"> <b>Supported Segments</b> </p>
<p align="center"> <b>MSH | EVN | PID | PD1 | ROL | DB1 | ARV | NK1 | PV1 | PV2 | GT1 | IN1 | AL1 | DG1 | ORC | OBR | NTE | OBX</b> </p>

To send an HL7 message with automatically generated fake patient data, use the `placebo --send hl7` command. This feature supports various healthcare scenarios through different sub-commands.

    placebo --send hl7 [sub command]

- **Basic Usage**: Sends an HL7 message to the default address `127.0.0.1:9700`.
    - `placebo --send hl7`
- **Sub-commands**:
  - `post_admit`: Generates an ADT^A01 event that admits a patient.
    - Usage: `placebo --send hl7 post-admit`
  - `post_discharge`: Generates ADT^A01 and ADT^A03 events to admit and then discharge a patient.
    - Usage: `placebo --send hl7 post-discharge`
  - `pre_admit`: Generates an ADT^A05 event that establishes preadmit information.
    - Usage: `placebo --send hl7 pre-admit`
  - `referral`: Generates a REF^I12 event that inform for a patient referral info.
    - Usage: `placebo --send hl7 referral`

Example:

    placebo --send hl7 post-discharge

This command sends an HL7 message that admits and then discharges a patient.

#### Helpful Auxiliary `send` Commands

|Flag | Description | Usage |
| --- | --- | --- |
|`last` | Open last sent hl7 message in an interactive prompt. | `placebo --send hl7 last` |
|`sugarpill` | Construct a hl7 message with assistance using an easy to read interactive prompt. | `placebo --send hl7 sugarpill` |

### Read HL7 Message

For a better help at reading HL7 messages, you can tap into the `sugarpill` feature and have the file presented in a more readible structure.

    placebo --read sugarpill /path/to/hl7_message.txt

Example:

`hl7_message.txt` content:
```
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|202405290800||ADT^A01|12345|P|2.3|
```

`placebo --read sugarpill hl7_message.txt` output:

```
{
 "MSH": {
  "Encode": "^~\\\u0026",
  "SendingApplication": "SENDAPP",
  "SendingFacility": "PLACEBO",
  "ReceivingApplication": "RECVAPP",
  "ReceivingFacility": "LAB",
  "DateTimeOfMessage": "202405290800",
  "Security": "",
  "MessageType": {
   "MessageCode": "ADT",
   "TriggerEvent": "A01"
  },
  "MessageControlID": "12345",
  "ProcessingID": "P",
  "VersionID": "2.3"
 },
}
```
