# Whisper
Server/host discovery in a network


## Building

```
make all
```

It will create two binaries, `bin/client` and `bin/server`

## Running

Run the `client` executable on a machine that should be discovered, and the `server` on a machine to monitor hosts on the same network.
Once running, clients send a periodic ping in broadcast mode, with a signature payload, and the server picks it up.

To query the the server for hosts the network, send  GET request to it's at http://{server-IP}:46790
The server will respond with a payload of listed hosts.

e.g.

```
{
   "192.168.1.10!:55460":{
      "Host":{
         "IP":"192.168.1.101",
         "Port":55460,
         "Zone":""
      },
      "LastPing":"2016-06-27T19:10:14.675370062+02:00",
      "Payload":{
         "os":"darwin",
         "arch":"amd64"
      }
   }
}
```
