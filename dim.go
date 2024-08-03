package main

func dimById(id byte) string {
	switch id {
	case byte(0):
		return "minecraft:overworld"
	case byte(int8(-1)):
		return "minecraft:the_nether"
	case byte(1):
		return "minecraft:the_end"
	default:
		return "minecraft:overworld"
	}
}
