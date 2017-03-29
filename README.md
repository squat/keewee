# keewee

[![Build Status](https://travis-ci.org/squat/keewee.svg?branch=master)](https://travis-ci.org/squat/keewee)

keewee makes it easy to launch and secure a self-hosted [KeeWeb](https://github.com/keeweb/keeweb) instance.

To secure your instance, keewee only allows access with encrypted connections. keewee uses [Let's Encrypt](https://letsencrypt.org/) to automatically obtain and renew TLS certificates for your KeeWeb instance. That means that when your certificates expire you aren't locked out of your passwords; instead, the certificates are kept up to date for you. It also means that there is no need to reconfigure and redeploy your password manager every few months.

## Running keewee

keewee is containerized and can be run with either rkt or docker. Just make sure to specify the host name for your instance:

```sh
docker run -p 443:443 -v keewee:/cache squat/keewee --host=example.com
```

*note*: this invocation uses a named volume mounted at `/cache` to ensure your certificates are cached between container restarts so that new certificates aren't issued all the time. If you don't care about caching your certificates between restarts, you can omit the `-v keewee:/cache` flag.

To keep an accessible cache of your certificates on your host, mount a host directory as the cache volume:

```sh
docker run -p 443:443 -v ./path/to/cache:/cache squat/keewee --host=example.com
```

Configuring the settings on your KeeWeb instance is also simplified with keewee. You no longer need to build your configuration file into your KeeWeb container; just update your configuration on disk!

Let's say I have a KeeWeb configuration file that looks like:

```json
{
    "settings": {
        "gdriveClientId": "<MY-CLIENT-ID>",
        "theme": "wb"
    }
}
```

I can then configure keewee by mounting that file as a volume at `/static/config.json` like:

```sh
docker run -p 443:443 -v keewee:/cache -v ./path/to/config.json:/static/config.json squat/keewee --host=example.com
```
