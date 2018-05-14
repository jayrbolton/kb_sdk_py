FROM jjeffryes/kbasepythonsdk:latest

# -----------------------------------------
# App-specific dependencies


# End app-specific dependencies
# -----------------------------------------

COPY ./ /kb/module
WORKDIR /kb/module
RUN mkdir -p /kb/module/work
RUN chmod -R a+rw /kb/module
RUN make all
ENTRYPOINT [ "./scripts/entrypoint.sh" ]
CMD [ ]
