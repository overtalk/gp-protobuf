#!/bin/bash

protoc -I . --doc_out=docs/v1/ --doc_opt=markdown,introduction.md proto/v1/*.proto
