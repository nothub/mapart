# Update conversion

Loading old maps on new servers can cause some issues.

The samples in here were created on 1.12.2, then subsequently loaded on a 1.16.5 and 1.21 server.

Additionally, a new map was created on each subsequent version.

## Dimension type issue

The server does not automatically convert maps to a new format!

Pre-1.16 maps will contain a byte value for the `dimension` entry.
Newer maps will contain a string (resource location) value for the `dimension` entry.

A 1.16 server can load and display a 1.12 map but will not update the format.

A 1.21 server will fail to load a 1.12 map, it will when trying to parse the old `dimension` (byte) value:

`java.lang.IllegalArgumentException: Invalid map dimension: null`

In a server log, the problem might look like this:

```
[12:57:53] [Server thread/ERROR]: Not a string
[12:57:53] [Server thread/ERROR]: Error loading saved data: map_0
java.lang.IllegalArgumentException: Invalid map dimension: null
	at net.minecraft.world.level.saveddata.maps.MapItemSavedData.lambda$load$1(MapItemSavedData.java:169) ~[paper-1.21.jar:1.21-124-df3b654]
...
```
