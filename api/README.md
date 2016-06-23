# SMTP API

This is the SMTP API with fqdn go.micro.api.smtp for email delivery.

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
http[domain:micro API port]/smtp/send
{
    "recipient":[
        "user@domain.com", 
        "user2@domain.com"
    ], 
    "subject": "Mail subject", 
    "body": "<table style=\"width:100%;\"><tr><td>lets</td><td>see</td></tr><tr><td>the</td><td>markup</td></tr></table>"
}

{}
```
