package main

func fixCmd(inputs []string) {
	// TODO
}

func dimById(id int) string {
	switch id {
	case 0:
		return "minecraft:overworld"
	case -1:
		return "minecraft:the_nether"
	case 1:
		return "minecraft:the_end"
	default:
		return "minecraft:overworld"
	}
}
