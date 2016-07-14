# JabMUD

An experiment with eJabberd as a scalable communication backbone. 
I wanted to learn more about eJabberd / XMPP and toy around with it,
creating the basis for a text mud seemed like a good way to go at
the time.

JabMUD runs as an external component hosted through XEP-0114,
allowing me to write this in Go instead of Erlang. I might
end up regretting that...

## eJabberd Docker Container

Instead of installing ejabber + services locally you can run it
through a docker container, see ```docker\docker-build.sh```

### ejabberd.yml configuration

ejabberd needs to be told to open a listener for our component:

```yaml
listen:
  -
    port: 5275
    module: ejabberd_service
    access: all
    shaper_rule: fast
    hosts:
      "jabmud.localhost":
        password: "secret"
```

### Opening Ports

These are the ports that the docker-run script opens up.

    -p 4560:4560 \
    -p 5222:5222 \
    -p 5269:5269 \
    -p 5275:5275 \
    -p 5280:5280 \
    -p 5443:5443 \


### Creating jabmud-user

    -e "EJABBERD_USERS=jabmud_user@localhost:password"
