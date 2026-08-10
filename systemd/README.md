# systemd

The following is provided for reference only. You will need to adjust or tailor specifics to you set up as follows. These examples assume the following:

* There is a system-level user called `geocoder`.
* The source code (for `sfomuseum/geocoder`) is stored in `/usr/local/src/geocoder`.
* The data file that the geocoder service will consume is stored at `/usr/local/data/geocoder/wof.db`
* The geocoder-server listend for requests on `0.0.0.0:3000`

## Set up

Make sure there is a `geocoder` user:

```
$> sudo useradd -M geocoder
```

And then:

```
$> sudo ln -s /usr/local/src/geocoder/systemd/geocoder-server.service.example /lib/systemd/system/geocoder-server.service
$> sudo systemctl daemon-reload
$> sudo systemctl enable geocoder-server.service
$> sudo systemctl start geocoder-server.service
$> sudo systemctl status geocoder-server.service
```

For testing:

```
$> curl -s 'http://localhost:3000/api/query?query=SFO&placetype=airport'
```