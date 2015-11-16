# mibob
connect your boblight to your milights

## Compile
`make` will compile mibob.  
To cross-compile for linux run: `make linux`  
For the RaspberryPi run: `make pi` or `make pi2`

## mibob config
`mibob.config`
```
{
	"bridge" : "192.168.2.69"
}

```
**The config file must be located in the same path as mibob**

## boblight config
```
[device]
name     milight
output   [path to mibob]/mibob
channels        3
type     popen
interval        20000

[color]
name            red
rgb                     FF0000

[color]
name            green
rgb                     00FF00

[color]
name            blue
rgb                     0000FF

[light]
name            right
color           red     milight 1
color           green   milight 2
color           blue    milight 3
hscan           10 90
vscan           10 90
```

if you are not sure about the config parameters read up on them here: https://code.google.com/p/boblight/wiki/boblightconf

## Limitations
The current version has a few limitations
- boblight color will be sent to all milight groups
- in order to set a different color than the one send by boblight, you need to disable boblight
- from time to time to lamps wont react anymore, send the all group on command to fix this
