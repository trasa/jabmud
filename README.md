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

## ejabberd.yml configuration

ejabberd needs to be told to open a listener for our component:

```yaml
listen:
  -
    port: 5275
    module: ejabberd_service
    access: all
    shaper_rule: fast
    ip: "127.0.0.1"
    hosts:
      "jabmud.localhost":
        password: "secret"
```


