# Use the official golang image as the base image
FROM golang:latest

RUN curl -fsSL https://miktex.org/download/key | tee /usr/share/keyrings/miktex-keyring.asc > /dev/null
RUN echo "deb [signed-by=/usr/share/keyrings/miktex-keyring.asc] https://miktex.org/download/debian bookworm universe" | tee /etc/apt/sources.list.d/miktex.list
	
# Install texlive and other LaTeX dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
	miktex \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
	
RUN miktexsetup finish
RUN initexmf --set-config-value [MPM]AutoInstall=1

RUN miktex packages update-package-database && miktex packages update

RUN mpm --admin --update && mpm --update
RUN mpm --admin --update && mpm --update

RUN cd ~/.miktex/texmfs/install/miktex/config && miktex-makefmt --engine=pdftex --dest-name=pdflatex --no-dump pdflatex.ini
	
# Set the working directory inside the container
WORKDIR /go/src/app

COPY ./select.sh .

# Run your application
ENTRYPOINT ["./select.sh"]

