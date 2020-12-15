# Python virtualenv cheatsheet:

## If you need python2 to be the default python within a virtual env:

* On Ubuntu 20.04 only python3 is installed and python is not a symlink to anything. Python2 can be installed:
  `sudo apt install python2`

* Python versions must still be specified on the command line after installation - no symlink for "python", even after installation of python2. Take note of this for later:

   ```
   wlieberz@workstation4:venv$ which python
   wlieberz@workstation4:venv$ which python2
   /usr/bin/python2
   wlieberz@workstation4:venv$ which python3
   /usr/bin/python3
   
   ``` 


* Now you can use the virtualenv command with a custom path (-p) to python. 
  * In this case we create an env dir called "esxi-ansible-python2" within the current directory.
    `virtualenv -p /usr/bin/python2 esxi-ansible-python2`

* Activate this venv as usual:
  `source esxi-ansible-python2/bin/activate`

* The prompt changes to indicate you are in a virtual env:
  `(esxi-ansible-python2) wlieberz@workstation4:venv$`

* Within the env bare-word python now points to python2, as desired:
   ```
   (esxi-ansible-python2) wlieberz@workstation4:venv$ python --version
   Python 2.7.18rc1

   ```


* Now would be a good time to upgrade pip within the venv:
  `python -m pip install --upgrade pip`

* Install a specific version of a package via pip:
  `pip install 'ansible==2.8.5'`


* You can always get out of your env with:
  `deactivate`

* You can remove an env with:
  `rm -rf my-old-venv`


