# Overview:

Notes on a method to temporarily bypass a restrictive corporate firewall to install packages with yum or with apt. This has been tested on CentOS 7 and Ubuntu 18.04. 

For this to work you need a host with unrestricted internet access which can ssh to the restricted host. We will call these: internet-host & restricted-host.

High-level: 

1) Run proxy.py on internet-host listening on port 8899.

2) Reverse ssh tunnel from internet-host to restricted-host, connecting the restricted-hosts's local port 8899 to internet-host's 8899.

3) On restricted-host, configure yum to use a proxy: http://127.0.0.1:8899.

# Details:

## 1 Run proxy.py:

### Install proxy.py in a virtualenv in your home directory:

* On internet-host:

```bash

cd ~
mkdir env
cd env
python3 -m pip install virtualenv --user
python3 -m virtualenv proxy-py
source proxy-py/bin/activate
pip install --upgrade proxy.py

```

### Run proxy.py:

* On internet-host:

Ensure your virtualenv for proxy-py is active, then run the proxy:

```bash

source proxy-py/bin/activate
proxy --hostname 127.0.0.1 --port 8899
```

## 2 Reverse ssh tunnel:

* On internet-host:

```bash

ssh -R 8899:localhost:8899 <Your User>@<your restricted-host>
```

## 3 Configure host to use proxy:

On your restricted-host:

* Testing:

Assuming curl'ing google.com usually fails, this should now work:

`curl -x "http://127.0.0.1:8899" https://www.google.com`

#### Configure yum for RHEL-based hosts:

Add to `/etc/yum.conf`:

```bash
proxy=http://127.0.0.1:8899
```

#### Configure apt for Debian-based hosts:

* Add file:

`/etc/apt/apt.conf.d/proxy.conf`

* With contents:

```
Acquire::http::Proxy "http://127.0.0.1:8899/";
Acquire::https::Proxy "http://127.0.0.1:8899/";
```


Now, you should be able to:

`yum install foo`

or 

`apt install foo`


* Optional: comment out the changes you made to the yum / apt conf files when done. 

Thanks for reading. 