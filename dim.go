package main

func dimById(id byte) string {
	switch id {
	case 0b00000000:
		return "minecraft:overworld"
	case 0b11111111:
		return "minecraft:the_nether"
	case 0b00000001:
		return "minecraft:the_end"
	default:
		return "minecraft:overworld"
	}
}
