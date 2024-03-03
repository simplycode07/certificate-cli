#!/bin/bash
go run . -t example/template.json -c example/certificate_blank.png -n example/names.csv
ranger example/certificatesGenerated/
