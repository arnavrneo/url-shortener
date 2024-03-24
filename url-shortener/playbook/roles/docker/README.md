Docker
=========

Role for setting up the files and the containers.

Requirements
------------

No extra requirements other than the `ansible` itself installed.

Role Variables
--------------

To change the directory on the remote servers where script should be copied, edit the `script_dest` in `main.yml` under `vars` directory.

Example Playbook
----------------

Including an example of how to use your role (for instance, with variables passed in as parameters) is always nice for users too:

    - hosts: servers
      roles:
         - { role: username.rolename, x: 42 }

License
-------

BSD

Author Information
------------------

Author: Arnav Raina (github.com/arnavrneo)
