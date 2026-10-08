# Placebo

[![CI](https://github.com/hl7x/placebo/actions/workflows/placebo.yml/badge.svg)](https://github.com/hl7x/placebo/actions/workflows/placebo.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

Placebo is a command-line tool designed for creating and managing fake patient data. It's particularly useful for testing purposes in healthcare applications, offering functionalities to generate CSV files with fake patient data and to send HL7 messages simulating different patient scenarios.

In addition to all of that, it has some robust features that help aid with reading HL7 messages! Useful if you're not used to reading pipes and carets.


<img width="1200" height="600" alt="new_demo" src="https://github.com/user-attachments/assets/9c593c97-f3e9-46f3-8153-90bba651de8b" />


## Features

- **CSV File Generation**: Create CSV files with automatically generated fake patient data.
- **HL7 Message Sending**: Send HL7 messages with automatically generated fake patient data.
- **Multiple Message Types**: Build admits, orders, results, appointments, referrals and document notifications, each carrying the segments its type is defined around.
- **HL7 Message Reading**: Feed `placebo` an hl7 file and get readible structure of the hl7 message.

## Installation

Tagged releases are built by [GoReleaser](https://goreleaser.com) and published on the
[releases page](https://github.com/hl7x/placebo/releases) for Linux, macOS, and Windows on both
`amd64` and `arm64`. Pick whichever of the following fits your machine.

### Homebrew (macOS, Linux)

```
$ brew install hl7x/tap/placebo
```

Upgrade later with `brew upgrade placebo`.

### Linux packages (deb, rpm, apk)

Each release ships `.deb`, `.rpm`, and `.apk` packages. Download the one matching your
distribution and architecture from the [releases page](https://github.com/hl7x/placebo/releases/latest),
then install it:

```
$ sudo dpkg -i placebo_<version>_amd64.deb      # Debian, Ubuntu
$ sudo rpm -i placebo-<version>.x86_64.rpm      # Fedora, RHEL, openSUSE
$ sudo apk add --allow-untrusted placebo_<version>_x86_64.apk   # Alpine
```

### Pre-built binary

Archives are named `placebo_<version>_<os>_<arch>.tar.gz` (`.zip` on Windows). To grab one from
the terminal:

```
$ VERSION=0.1.0
$ curl -sSL -o placebo.tar.gz \
    https://github.com/hl7x/placebo/releases/download/v${VERSION}/placebo_${VERSION}_darwin_arm64.tar.gz
$ tar -xzf placebo.tar.gz
$ sudo mv placebo /usr/local/bin/
```

Substitute `linux`/`darwin`/`windows` and `amd64`/`arm64` as needed. Every release also includes a
`checksums.txt` you can verify against:

```
$ sha256sum --check --ignore-missing checksums.txt
```

*Note*: macOS binaries are unsigned. If Gatekeeper quarantines a binary you downloaded manually,
clear it with `xattr -dr com.apple.quarantine /usr/local/bin/placebo`. The Homebrew cask does this
for you.

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

Confirm whichever route you took with:

```
$ placebo version
```

## Usage

    placebo <command> [subcommand] [arguments] [options]

|Command | Description |
| --- | --- |
|`file` | Create a file of fake patient data. |
|`send` | Send an HL7 message built from fake patient data. |
|`listen` | Receive and print HL7 messages. |
|`read` | Break an HL7 message down into a readable structure. |
|`version` | Show the version of placebo you are running, e.g. `placebo version` or `placebo --version`. |
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

    placebo file hl7 [message type] [scenario]

- This command creates a random HL7 file with a fake patient at `/tmp/`.
- With nothing further it builds an `ADT^A01` admit.
- Name a [message type](#message-types) to build another kind of message, and a scenario to pick the trigger event within it.
- Example: `placebo file hl7 oru` writes a lab result, `placebo file hl7 siu reschedule` writes a rescheduled appointment.
- `placebo file hl7 types` lists every message type and scenario.

### Send HL7 Messages

<p align="center"> <b>Supported Segments</b> </p>
<p align="center"> <b>MSH | EVN | PID | PD1 | ROL | DB1 | ARV | NK1 | PV1 | PV2 | GT1 | IN1 | AL1 | DG1 | SCH | RF1 | TXA | ORC | OBR | NTE | OBX</b> </p>

To send an HL7 message with automatically generated fake patient data, use the `placebo send hl7` command.

    placebo send hl7 [message type] [scenario]
    placebo send hl7 [subcommand]

- **Basic Usage**: Sends an HL7 message to the default address `127.0.0.1:9700`.
    - `placebo send hl7`
- With nothing further, an `ADT^A01` admit is opened in your editor before it is sent.
- A scenario on its own is read as an ADT event, so `placebo send hl7 discharge` still means `ADT^A03`.

Example:

    placebo send hl7 discharge

This command sends an HL7 message that discharges a patient.

### Message Types

A message type decides which segments the message carries. An order message is built from `ORC` and `OBR`; an appointment is built from `SCH`. Only the segments that type is defined around are sent, so a lab result does not arrive with an insurance segment stapled to it.

Name the type first, then the scenario:

    placebo send hl7 oru
    placebo send hl7 siu reschedule
    placebo file hl7 ref referral

Leaving the scenario off uses the first one listed for that type. Run `placebo send hl7 types` for this table at the terminal.

|Message Type | Description | Segments |
| --- | --- | --- |
|`ADT` | Patient administration | MSH EVN PID PD1 ROL DB1 ARV NK1 PV1 PV2 GT1 IN1 AL1 DG1 |
|`ORM` | Order | MSH PID PV1 ORC OBR DG1 NTE |
|`ORU` | Observation result | MSH PID PV1 ORC OBR OBX NTE |
|`SIU` | Scheduling | MSH SCH PID PV1 PV2 NTE |
|`REF` | Patient referral | MSH RF1 PID PV1 DG1 NTE |
|`MDM` | Document notification | MSH EVN PID PV1 TXA OBX |

#### Scenarios

|Message Type | Scenario | Event | Description |
| --- | --- | --- | --- |
|`ADT` | `admit` | ADT^A01 | Admit a patient |
|`ADT` | `transfer` | ADT^A02 | Transfer a patient |
|`ADT` | `discharge` | ADT^A03 | Discharge a patient |
|`ADT` | `register` | ADT^A04 | Register a patient |
|`ADT` | `pre-admit` | ADT^A05 | Establish preadmit information |
|`ADT` | `update` | ADT^A08 | Update patient information |
|`ADT` | `cancel-admit` | ADT^A11 | Cancel an admit |
|`ADT` | `cancel-discharge` | ADT^A13 | Cancel a discharge |
|`ORM` | `order` | ORM^O01 | Place a new order |
|`ORU` | `result` | ORU^R01 | Report an observation result |
|`SIU` | `schedule` | SIU^S12 | Book a new appointment |
|`SIU` | `reschedule` | SIU^S13 | Reschedule an appointment |
|`SIU` | `cancel` | SIU^S15 | Cancel an appointment |
|`SIU` | `no-show` | SIU^S26 | Record a patient who did not arrive |
|`REF` | `referral` | REF^I12 | Refer a patient |
|`REF` | `modify` | REF^I13 | Modify a referral |
|`REF` | `cancel` | REF^I14 | Cancel a referral |
|`MDM` | `document` | MDM^T02 | Notify of a new document |
|`MDM` | `status-change` | MDM^T04 | Notify of a document status change |
|`MDM` | `addendum` | MDM^T06 | Notify of a document addendum |

Order and result messages carry a generated lab test with a value placed against its own reference range, so results come back flagged normal, high or low rather than always reading the same.

#### Helpful Auxiliary `send` Subcommands

|Subcommand | Description | Usage |
| --- | --- | --- |
|`file` | Edit an existing HL7 file and send it when you are done. | `placebo send hl7 file /path/to/hl7_message.txt` |
|`last` | Open last sent hl7 message in an interactive prompt. | `placebo send hl7 last` |
|`sugarpill` | Construct a hl7 message with assistance using an easy to read interactive prompt. | `placebo send hl7 sugarpill` |
|`types` | List every message type and scenario placebo can build. | `placebo send hl7 types` |

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
