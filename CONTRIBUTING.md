# Contributing to Placebo

Thanks for your interest in improving Placebo! This project generates **synthetic
(fake) patient data** for testing healthcare integrations. Contributions of all
kinds are welcome.

## Ground rules

- **Never commit real patient data (PHI/PII).** This tool exists specifically so
  that nobody has to. Sample HL7 messages and fixtures must be fabricated.
- Be respectful. See our [Code of Conduct](./CODE_OF_CONDUCT.md).

## Development setup

Requires Go 1.21 or newer.

```
git clone https://github.com/hl7x/placebo.git
cd placebo
go build ./...
go test ./...
```

## Before you open a pull request

Please make sure the following pass locally, since CI enforces them:

```
gofmt -l .        # should print nothing
go build ./...
go test ./...
```

## Pull request guidelines

- Keep PRs focused; open an issue first for larger changes.
- Add or update tests for behavior changes.
- Write a clear description of what changed and why.

## Reporting bugs / requesting features

Open an issue using the appropriate template. For anything security-related,
see [SECURITY.md](./SECURITY.md) instead of opening a public issue.