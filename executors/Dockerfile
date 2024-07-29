FROM ubuntu:20.04

RUN apt-get update && apt-get install -y \
    g++ \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -m coderunner

WORKDIR /app
RUN chown coderunner:coderunner /app

USER coderunner

CMD ["/bin/bash"]
