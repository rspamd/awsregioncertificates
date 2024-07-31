## AWS Region Certificates for Go

[![GHA](https://github.com/rspamd/awsregioncertificates/actions/workflows/ci.yml/badge.svg)](https://github.com/rspamd/awsregioncertificates/actions/workflows/ci.yml)
[![GHA](https://github.com/rspamd/awsregioncertificates/actions/workflows/generate_test.yml/badge.svg)](https://github.com/rspamd/awsregioncertificates/actions/workflows/generate_test.yml)

Embeds region certificates from AWS (run `go generate ./...` to refresh them) - AFAIU they can only be downloaded as part of documentation.

Can be used to validate EC2 [instance identity documents](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html) using the `rsa`/`base64-encoded signature`.

## Credits

Special thanks to [@vstakhov](https://github.com/vstakhov) for sponsoring this work.
