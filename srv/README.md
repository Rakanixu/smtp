# SMTP srv

This is the SMTP service with fqdn go.micro.srv.smtp for email delivery.

## Getting Started

### Prerequisites

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev -advertise=127.0.0.1
```

### Run Service

```
$ go run main.go
```

### Usage

#### Send
```
micro query go.micro.srv.smtp SMTP.Send '{"recipient":["user@domain.com", "add@domain.es"], "subject": "Mail subject", "body": "<table style=\"width:100%;\"><tr><td>lets</td><td>see</td></tr><tr><td>the</td><td>markup</td></tr></table>"}'
```
