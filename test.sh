#!/usr/bin/env bash

# Licensed to Elasticsearch B.V. under one or more contributor
# license agreements. See the NOTICE file distributed with
# this work for additional information regarding copyright
# ownership. Elasticsearch B.V. licenses this file to you under
# the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#	http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

# Script to test the output of crd-ref-docs

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

TEMP_DIR=$(mktemp -d -t crd-ref-docs-XXXXX)
trap '[[ $TEMP_DIR ]] && rm -rf "$TEMP_DIR"' EXIT


(
    cd "$SCRIPT_DIR"
    go run main.go --log-level=ERROR --source-path="${SCRIPT_DIR}/test" --renderer=asciidoctor --templates-dir="${SCRIPT_DIR}/templates/asciidoctor" --output-path="${TEMP_DIR}/out.asciidoc"
    DIFF=$(diff -a -y --suppress-common-lines "${SCRIPT_DIR}/test/expected.asciidoc" "${TEMP_DIR}/out.asciidoc") || true
    if [ "$DIFF" ]; then
        echo "ERROR: outputs differ"
        echo ""
        echo "$DIFF"
        exit 1
    else
        echo "OK"
    fi
)
