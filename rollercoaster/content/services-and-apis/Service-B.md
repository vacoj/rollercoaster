---
draft: false
date: "2017-10-02T15:03:19"
tags: ["MessageBus-C",
"IdentityProvider",
"MessageBus-B",
]
title: Service-B
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-B fill:#76e40a,stroke:#809549,stroke-width:2px",
"Service-B -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
]
---
			