FROM ubuntu:18.04 AS runtime
# GoReleaser will automatically generate the binary in the root directory
COPY /kusionup .