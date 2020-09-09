BASEDIR = $(shell pwd)

all: http-01-basics http-02-rest http-03-sse

http-01-basics:
	cd ${BASEDIR}/01-basics && go build -o ${BASEDIR}/http-01-basics

http-02-rest:
	cd ${BASEDIR}/02-rest && go build -o ${BASEDIR}/http-02-rest

http-03-sse:
	cd ${BASEDIR}/03-sse && go build -o ${BASEDIR}/http-03-sse

clean:
	rm ${BASEDIR}/http-01-basics
	rm ${BASEDIR}/http-02-rest
	rm ${BASEDIR}/http-03-sse