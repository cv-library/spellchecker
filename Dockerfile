FROM debian:jessie

MAINTAINER Tim Retout <t.retout@cv-library.co.uk>

RUN apt-get update && apt-get install -y \
	libhunspell-1.3-0 \
	myspell-en-gb \
	unzip \
	wget

COPY spellchecker /app/
COPY static /app/static/

EXPOSE 80

WORKDIR /app
CMD ["/app/spellchecker"]
