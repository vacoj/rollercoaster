---
draft: false
date: "2017-10-02T15:15:50"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-A
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-A fill:#3cde8c,stroke:#552523,stroke-width:2px",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-A -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
]
---
			