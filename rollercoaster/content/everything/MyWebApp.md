---
draft: false
date: "2017-10-02T15:15:50"
tags: ["Service-A",
"ServiceAPI",
"File_Share",
"LoggingDB",
"IdentityProvider",
]
title: MyWebApp
categories: ["everything"]
depmap: [ "graph LR",
"style MyWebApp fill:#21babc,stroke:#6f1f35,stroke-width:2px",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
			