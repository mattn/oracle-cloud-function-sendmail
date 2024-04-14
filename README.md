# oracle-cloud-function-sendmail

## Usage

```
echo -n '{"to": "User Name <user@example.com>", "subject": "Hello from Oracle Cloud Function", "body": "Hello World!"}' | fn invoke my-app sendmail
```

## Requirements

* fnproject

## Installation

```
fn cf f my-app sendmail OCI_EMAIL_DELIVERY_SMTP_SERVER 'smtp.email.<region>.oci.oraclecloud.com'
fn cf f my-app sendmail OCI_EMAIL_DELIVERY_USER_OCID '<your-ocid>'
fn cf f my-app sendmail OCI_EMAIL_DELIVERY_USER_PASSWORD '<your-password>'
fn cf f my-app sendmail OCI_EMAIL_DELIVERY_APPROVED_SENDER 'Your Name <you@your-mail-server.com>'
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
