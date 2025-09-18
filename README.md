# mapart

A tool to fumble around with Minecraft [map files](https://minecraft.wiki/w/Map_item_format#map_%3C#%3E.dat_format_(Java)).

---

Latest supported version: `1.21.8` / `4440`

For usage help see: [USAGE.txt](USAGE.txt)

## Map Format Infos

| Data Version | MC Versions     | Changes                                                                                                                                                                                                       |
|--------------|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| -            | <= 1.8.9        | initial map format                                                                                                                                                                                            |
| 169          | 1.9 - 1.10.2    | add field: trackingPosition (byte)                                                                                                                                                                            |
| 819          | 1.11 - 1.12.2   | add field: unlimitedTracking (byte)                                                                                                                                                                           |
| 1519         | 1.13            | add field: banners (but we ignore this anyway)<br>remove field: width<br>remove field: height                                                                                                                 |
| 1628         | 1.13.1 - 1.13.2 | change dimension type to int<br>add field: frames (but we ignore this anyways)                                                                                                                                |
| 1952         | 1.14 - 1.14.4   | add field: locked (byte)                                                                                                                                                                                      |
| 2566         | 1.16 - 1.16.4   | change dimension type to string (resource location)                                                                                                                                                           |
| 2586         | 1.16.5 - 1.19.4 | add field: UUIDMost (long)<br>add field: UUIDLeast (long)                                                                                                                                                     |
| 3463         | 1.20 - 1.20.5   | remove field: UUIDMost (long)<br>remove field: UUIDLeast (long)                                                                                                                                               |
| 3839         | 1.20.6 - 1.21   | add field: UUIDMost (long)<br>add field: UUIDLeast (long)                                                                                                                                                     |
| 3955         | 1.21.1 - 1.21.4 | remove field: UUIDMost (long)<br>remove field: UUIDLeast (long)                                                                                                                                               |
| 4325         | \>= 1.21.5      | remove field: locked (byte)<br>remove field: scale (byte)<br>remove field: trackingPosition (byte)<br>remove field: unlimitedTracking (byte)<br>remove field: banners (array)<br>remove field: frames (array) |
