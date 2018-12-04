FROM python:3.7-alpine

# Install pip dependencies
WORKDIR /kb/module
COPY requirements.txt /kb/module/requirements.txt
RUN apk --update add --virtual build-dependencies python-dev build-base && \
    pip install --upgrade --no-cache-dir pip -r requirements.txt && \
    apk del build-dependencies

# Run the app
COPY . /kb/module
RUN chmod -R a+rw /kb/module
EXPOSE 5000
ENTRYPOINT ["sh", "/usr/local/bin/entrypoint.sh"]  # from the kbase_module package
