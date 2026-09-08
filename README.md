# Placebo

[![CI](https://github.com/hl7x/placebo/actions/workflows/placebo.yml/badge.svg)](https://github.com/hl7x/placebo/actions/workflows/placebo.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

Placebo is a command-line tool designed for creating and managing fake patient data. It's particularly useful for testing purposes in healthcare applications, offering functionalities to generate CSV files with fake patient data and to send HL7 messages simulating different patient scenarios.

In addition to all of that, it has some robust features that help aid with reading HL7 messages! Useful if you're not used to reading pipes and carets.

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

    placebo <command> [subcommand] [arguments] [options]

|Command | Description |
| --- | --- |
|`file` | Create a file of fake patient data. |
|`send` | Send an HL7 message built from fake patient data. |
|`listen` | Receive and print HL7 messages. |
|`read` | Break an HL7 message down into a readable structure. |
|`help` | Show help for a command, e.g. `placebo help send`. |

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

    placebo file csv [number of patients]

- This command creates a random CSV file with a fake patient at `/tmp/`.
- Adding a number to the `csv` subcommand will produce multiple fake patients.
- Example: `placebo file csv 4` creates a CSV file with 4 fake patients.

To create a HL7 message file with fake patient data:

    placebo file hl7

- This command creates a random HL7 file with a fake patient at `/tmp/`.

### Send HL7 Messages

<p align="center"> <b>Supported Segments</b> </p>
<p align="center"> <b>MSH | EVN | PID | PD1 | ROL | DB1 | ARV | NK1 | PV1 | PV2 | GT1 | IN1 | AL1 | DG1 | ORC | OBR | NTE | OBX</b> </p>

To send an HL7 message with automatically generated fake patient data, use the `placebo send hl7` command. This feature supports various healthcare scenarios through different subcommands.

    placebo send hl7 [subcommand]

- **Basic Usage**: Sends an HL7 message to the default address `127.0.0.1:9700`.
    - `placebo send hl7`
- **Preset ADT scenarios**:
  - `admit`: Generates an ADT^A01 event that admits a patient.
    - Usage: `placebo send hl7 admit`
  - `transfer`: Generates an ADT^A02 event that transfers a patient.
    - Usage: `placebo send hl7 transfer`
  - `discharge`: Generates an ADT^A03 event that discharges a patient.
    - Usage: `placebo send hl7 discharge`
  - `register`: Generates an ADT^A04 event that registers a patient.
    - Usage: `placebo send hl7 register`
  - `pre-admit`: Generates an ADT^A05 event that establishes preadmit information.
    - Usage: `placebo send hl7 pre-admit`

Example:

    placebo send hl7 discharge

This command sends an HL7 message that discharges a patient.

#### Helpful Auxiliary `send` Subcommands

|Subcommand | Description | Usage |
| --- | --- | --- |
|`file` | Edit an existing HL7 file and send it when you are done. | `placebo send hl7 file /path/to/hl7_message.txt` |
|`last` | Open last sent hl7 message in an interactive prompt. | `placebo send hl7 last` |
|`sugarpill` | Construct a hl7 message with assistance using an easy to read interactive prompt. | `placebo send hl7 sugarpill` |

### Receive HL7 Messages

To receive and print an HL7 message sent to the listening port:

    placebo listen hl7

### Choosing a Port

`send` and `listen` use port `9700` by default. Override it with the `--port` option, which may be given before the command or after the subcommand:

    placebo send hl7 --port 8500
    placebo --port 8500 listen hl7

It can also be set permanently with the `PLACEBO_PORT` environment variable:

    export PLACEBO_PORT=8500

### MLLP Transport

`placebo` speaks MLLP (Minimal Lower Layer Protocol), the framing every HL7 interface engine expects on a TCP connection. Messages go out wrapped as `<VT>message<FS><CR>` (`0x0B` ... `0x1C 0x0D`) with segments terminated by carriage returns, so `placebo` interoperates with Mirth, Rhapsody, Cloverleaf, Iguana, and Epic Bridges rather than only with itself.

- **Sending**: `placebo send hl7` frames the message, then waits up to 10 seconds for an acknowledgement and prints it. A receiver that never acknowledges is not treated as a failure; the message is already delivered.
- **Listening**: `placebo listen hl7` reads framed messages (a sender may put several on one connection) and replies to each with an `MSH` + `MSA|AA` acknowledgement. Without that reply a real sender would hold the connection open waiting.

A message that arrives without framing is still printed, along with a warning, so a misconfigured sender is obvious rather than silent.

### Read HL7 Message

For a better help at reading HL7 messages, you can tap into the `sugarpill` feature and have the file presented in a more readible structure.

    placebo read sugarpill /path/to/hl7_message.txt

Example:

`hl7_message.txt` content:
```
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|202405290800||ADT^A01|12345|P|2.3|
```

`placebo read sugarpill hl7_message.txt` output:

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
