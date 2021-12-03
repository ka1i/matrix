package usage

import "fmt"

//App Usage ...
func Usage() {
	fmt.Println(`Usage:
     ------- < Commands Arguments > -------
argument:
  -f, file          Required config source.

optional:
  -h, help          Show this help message.
  -v, version       Show the app version.

For more help, you can check the 
author account: https://github.com/ka1i for the detailed information`)
}
