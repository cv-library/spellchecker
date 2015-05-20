FROM debian:jessie

MAINTAINER Tim Retout <t.retout@cv-library.co.uk>

RUN apt-get update && apt-get install -y \
	libhunspell-1.3-0 \
	myspell-en-gb \
	unzip \
	wget

RUN 	mkdir -p /app/static \
	&& wget https://github.com/cv-library/jquery-spellchecker/releases/download/v0.3.0/jquery.spellchecker-0.3.0.zip \
	&& unzip jquery.spellchecker-0.3.0.zip -d /app/static \
	&& rm jquery.spellchecker-0.3.0.zip

COPY spellchecker /app/
COPY static /app/static/

EXPOSE 80

WORKDIR /app
CMD ["/app/spellchecker"]
