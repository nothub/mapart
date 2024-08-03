package main

func dimById(id byte) (string, bool) {
	switch id {
	case 0x00: //  0
		return "minecraft:overworld", true
	case 0xff: // -1
		return "minecraft:the_nether", true
	case 0x01: //  1
		return "minecraft:the_end", true
	default:
		return "", false
	}
}
