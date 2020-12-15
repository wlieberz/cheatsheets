
* Avoid shell wildcard expansion errors with hard quotes:

  `find . -iname '*puppet*'`


* The grep opt -H returns name of file which matched:
  So, find all files in the current dir which contain "API_BASE_URL":

  `find . -type f -exec grep -H "API_BASE_URL" {} \;`

* Get octal permissions of each dir in current working dir:
  * Uses bash's tilde expansion feature.
  
  `find ~+ -type d -exec stat -c '%a %n' {} \;`

