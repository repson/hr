# README

HR is a tool used to export a system's user information. The command will be able to export user names, IDs, home directories, and shells as either JSON or CSV. This command will not include information about system users (users with IDs under 1000).

By default, the command will display the information as JSON to stdout, but the -format flag will allow a person to specify csv as the export type. Additionally, a file can be specified by using the -path flag.

~~~
$ hr -format=csv -path=path/to/users.csv
$ hr -path=path/to/users.json
$ hr
[
  {
    "name": "user1",
    "id": 1002,
    "home": "/home/user1",
    "shell": "/bin/bash"
  },
  {
    "name": "user2",
    "id": 1003,
    "home": "/home/user2",
    "shell": "/bin/zsh"
  },
]
$ hr -format=csv
name,id,home,shell
user1,1002,/home/user1,/bin/bash
user2,1003,/home/user2,/bin/zsh
~~~

Hints:

Some Go libraries used to implement this tool:

- encoding/json: The built-in JSON encoding/decoding library.
- encoding/csv: The built-in CSV encoding/decoding library.
- strconv: The standard library package for converting strings to other types.
- strings: A library for working with strings. Useful for changing a string to all lowercase letters.

The information we want can be found in the /etc/passwd file. The lines in this file look like this:

user1:x:1002:1003::/home/user1:/bin/bash
