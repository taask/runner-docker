FROM docker:18.09

RUN wget https://storage.googleapis.com/gvisor/releases/nightly/latest/runsc \
	&& wget https://storage.googleapis.com/gvisor/releases/nightly/latest/runsc.sha512 \
	&& sha512sum -c runsc.sha512 \
	&& chmod a+x runsc \
	&& mv runsc /usr/local/bin