Docker
=========

Role for setting up the files and the containers.

Requirements
------------

- Ansible installed.
- If `.env` has to be used, create three dirs: `frontend`, `go` and `express_js` and then copy the respective `.env` under them. 

Role Variables
--------------

To change the directory on the remote servers where script should be copied, edit the `script_dest` in `main.yml` under `vars` directory.

Configure Hosts
----------------
<Will show connection to AWS EC2>

    - hosts: servers
      roles:
         - { role: username.rolename, x: 42 }

License
-------

BSD

Author Information
------------------

Author: Arnav Raina (github.com/arnavrneo)
